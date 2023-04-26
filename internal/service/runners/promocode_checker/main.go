package promocode_checker

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/data/postgres"
	"github.com/dl-nft-books/core-svc/resources"
)

type PromocodeChecker struct {
	name        string
	logger      *logan.Entry
	promocodesQ data.PromocodesQ
	runnerData  config.RunnerData
}

func New(cfg config.Config) *PromocodeChecker {
	return &PromocodeChecker{
		name:        cfg.PromocoderCfg().Name,
		logger:      cfg.Log(),
		promocodesQ: postgres.NewPromocodesQ(cfg.DB()),
		runnerData:  cfg.PromocoderCfg().Runner,
	}
}

func (pc *PromocodeChecker) promocodeCheck(ctx context.Context) error {
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeExpired).FilterExpired().Update(); err != nil {
		return errors.Wrap(err, "Failed to update promocode expired state")
	}
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeFullyUsed).FilterFullyUsed().Update(); err != nil {
		return errors.Wrap(err, "Failed to update promocode fully used state")
	}
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeActive).FilterActive().Update(); err != nil {
		return errors.Wrap(err, "Failed to update promocode active state")
	}
	return nil
}

func (pc *PromocodeChecker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, pc.logger, pc.name,
		pc.promocodeCheck,
		pc.runnerData.NormalPeriod, pc.runnerData.MinAbnormalPeriod, pc.runnerData.MaxAbnormalPeriod,
	)
}
