package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
<<<<<<< HEAD
	promocodeChecker "gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/promocode_checker"
=======
	taskCleaner "gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/cleaner"
>>>>>>> origin/feature/S3_cleaning
	taskProcessor "gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/task_processor"
)

func RunTaskProcessor(ctx context.Context, cfg config.Config) {
	taskProcessor.New(cfg).Run(ctx)
}

<<<<<<< HEAD
func RunPromocodeChecker(ctx context.Context, cfg config.Config) {
	promocodeChecker.New(cfg).Run(ctx)
=======
func RunTaskCleaner(ctx context.Context, cfg config.Config) {
	taskCleaner.New(cfg).Run(ctx)
>>>>>>> origin/feature/S3_cleaning
}
