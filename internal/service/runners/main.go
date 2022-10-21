package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/taskRunner"
)

func StartTaskRunner(ctx context.Context, cfg config.Config) {
	taskRunner.New(cfg).Run(ctx)
}
