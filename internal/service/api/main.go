package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	networker "gitlab.com/tokend/nft-books/network-svc/connector/api"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	db       *pgdb.DB

	ethMinterConfig *config.EthMinterConfig
	apiRestrictions config.ApiRestrictions

	pricer    *pricer.Connector
	networker *networker.Connector
	booker    *booker.Connector
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		db:       cfg.DB(),

		ethMinterConfig: cfg.EthMinter(),
		apiRestrictions: cfg.ApiRestrictions(),

		pricer:    cfg.PricerConnector(),
		booker:    cfg.BookerConnector(),
		networker: cfg.NetworkConnector(),
	}
}

func Run(ctx context.Context, cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(errors.Wrap(err, "failed to run a service"))
	}
}
