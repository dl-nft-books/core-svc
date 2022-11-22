package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/task_processor"
)

func RunTaskProcessor(ctx context.Context, cfg config.Config) {
	task_processor.New(cfg).Run(ctx)
}
