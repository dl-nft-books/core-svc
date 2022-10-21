package cli

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners"
	"os"
	"os/signal"
	"sync"

	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
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

	app := kingpin.New("generator-svc", "")

	runCmd := app.Command("run", "run command")
	serviceCmd := runCmd.Command("service", "run service")
	runnerCmd := runCmd.Command("runner", "run runner")

	migrateCmd := app.Command("migrate", "migrate command")
	migrateUpCmd := migrateCmd.Command("up", "migrate db up")
	migrateDownCmd := migrateCmd.Command("down", "migrate db down")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	// Creating context and sync.WaitGroup
	waitGroup := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	switch cmd {
	case serviceCmd.FullCommand():
		api.Run(cfg)
	case runnerCmd.FullCommand():
		run(waitGroup, ctx, cfg, runners.StartTaskRunner)
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
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
