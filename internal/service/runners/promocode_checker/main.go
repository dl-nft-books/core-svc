package promocode_checker

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type PromocodeChecker struct {
	name        string
	logger      *logan.Entry
	promocodesQ data.PromocodesQ
	runnerData  config.RunnerData
}

func New(cfg config.Config) *PromocodeChecker {
	return &PromocodeChecker{
		name:        cfg.PromocodesCfg().Name,
		logger:      cfg.Log(),
		promocodesQ: postgres.NewPromocodesQ(cfg.DB()),
		runnerData:  cfg.PromocodesCfg().Runner,
	}
}

func (pc *PromocodeChecker) promocodeCheck() error {
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeExpired).UpdateWhereExpired(); err != nil {
		return err
	}
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeFullyUsed).UpdateWhereFullyUsed(); err != nil {
		return err
	}
	if err := pc.promocodesQ.New().UpdateState(resources.PromocodeActive).UpdateWhereActive(); err != nil {
		return err
	}
	return nil
}

func (pc *PromocodeChecker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, pc.logger, pc.name,
		func(ctx context.Context) error {
			//pc.logger.Info("Check if promocodes has been expired")
			return pc.promocodeCheck()
		}, pc.runnerData.NormalPeriod, pc.runnerData.MinAbnormalPeriod, pc.runnerData.MaxAbnormalPeriod,
	)
}
