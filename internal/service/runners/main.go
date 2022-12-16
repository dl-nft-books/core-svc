package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	taskProcessor "gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/task_processor"
)

func RunTaskProcessor(ctx context.Context, cfg config.Config) {
	taskProcessor.New(cfg).Run(ctx)
}

func RunTaskCleaner(ctx context.Context, cfg config.Config) {
	taskProcessor.New(cfg).RunCleaner(ctx)
}
