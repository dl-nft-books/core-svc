package cli

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners"
	"os"
	"os/signal"
	"sync"

	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
)

var (
	app = kingpin.New("generator-svc", "service responsible for generating a book's pdf with a custom signature on it, handling status of uploading process, and storing tokens")

	// Run commands
	runCommand              = app.Command("run", "run command")
	apiCommand              = runCommand.Command("api", "run api")
	taskProcessorCommand    = runCommand.Command("task-processor", "run task processor")
	promocodeCheckerCommand = runCommand.Command("promocode-checker", "run promocode checker")
	allRunnersCommand       = runCommand.Command("all-runners", "run all runners")

	// Migration commands
	migrateCommand     = app.Command("migrate", "migrate command")
	migrateUpCommand   = migrateCommand.Command("up", "migrate database up")
	migrateDownCommand = migrateCommand.Command("down", "migrate database down")
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	cmd, err := app.Parse(args[1:])
	if err != nil {
		panic(errors.Wrap(err, "failed to parse arguments"))
	}

	// Creating context and sync.WaitGroup
	waitGroup := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	switch cmd {
	case apiCommand.FullCommand():
		run(waitGroup, ctx, cfg, api.Run)
		log.Info("started api...")
	case taskProcessorCommand.FullCommand():
		for i := uint64(0); i < cfg.TaskProcessorCfg().ProcessesNumber; i++ {
			run(waitGroup, ctx, cfg, runners.RunTaskProcessor)
			log.Infof("started task processor #%d", i+1)
		}
	case promocodeCheckerCommand.FullCommand():
		run(waitGroup, ctx, cfg, runners.RunPromocodeChecker)
		log.Info("started promocode checker...")
	case allRunnersCommand.FullCommand():
		for i := uint64(0); i < cfg.TaskProcessorCfg().ProcessesNumber; i++ {
			run(waitGroup, ctx, cfg, runners.RunTaskProcessor)
			log.Infof("started task processor #%d", i+1)
		}
		run(waitGroup, ctx, cfg, runners.RunPromocodeChecker)
		log.Info("started promocode checker...")
	case migrateUpCommand.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCommand.FullCommand():
		err = MigrateDown(cfg)
	}
	if err != nil {
		log.WithError(err).Error("failed to run command")
		cancel()
		return false
	}

	// We will stop channels gracefully
	graceful := make(chan os.Signal, 1)
	signal.Notify(graceful, os.Interrupt, os.Kill)

	waitGroupChannel := make(chan struct{})
	// Selectable wait group
	go func() {
		waitGroup.Wait()
		close(waitGroupChannel)
	}()

	select {
	case <-waitGroupChannel:
		log.Info("all services stopped")
		// Actually context should be already canceled.
		// Yet just to be sure we'll do it anyway.
		cancel()
	case <-graceful:
		log.Info("got signal to stop gracefully")
		cancel()
		<-waitGroupChannel
	}

	return true
}

func run(wg *sync.WaitGroup, ctx context.Context,
	cfg config.Config, runner func(context.Context, config.Config),
) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				cfg.Log().WithRecover(r).Error("service panicked")
			}
		}()

		runner(ctx, cfg)
	}()
}
