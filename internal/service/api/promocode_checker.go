package api

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"sync"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	minAbnormalDuration = time.Second
	maxAbnormalDuration = time.Minute * 30
)

const promocodeCheckName = "promocode_checker"

var promocodeCheckerInterval = 10 * time.Second

type promocodeChecker struct {
	logger      *logan.Entry
	promocodesQ data.PromocodesQ
}

func NewPromocodeChecker(db *pgdb.DB, l *logan.Entry) *promocodeChecker {
	return &promocodeChecker{
		logger:      l,
		promocodesQ: postgres.NewPromocodesQ(db),
	}
}

func (pc *promocodeChecker) promocodeCheck() error {
	err := pc.promocodesQ.New().UpdateState(resources.PromocodeExpired).UpdateWhereExpired()
	if err != nil {
		return err
	}
	return nil
}

func (pc *promocodeChecker) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	running.WithBackOff(
		ctx, pc.logger, promocodeCheckName,
		func(ctx context.Context) error {
			//ar.logger.Info("Check if promocodes has been expired")
			err := pc.promocodeCheck()
			if err != nil {
				return errors.Wrap(err, "failed to update promocodes state")
			}
			return nil
		}, promocodeCheckerInterval, minAbnormalDuration, maxAbnormalDuration,
	)
}
