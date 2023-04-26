package runners

import (
	"context"
	"github.com/dl-nft-books/core-svc/internal/config"
	taskCleaner "github.com/dl-nft-books/core-svc/internal/service/runners/cleaner"
	promocodeChecker "github.com/dl-nft-books/core-svc/internal/service/runners/promocode_checker"
)

func RunPromocodeChecker(ctx context.Context, cfg config.Config) {
	promocodeChecker.New(cfg).Run(ctx)
}

func RunTaskCleaner(ctx context.Context, cfg config.Config) {
	taskCleaner.New(cfg).Run(ctx)
}
