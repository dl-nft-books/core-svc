package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	booksConnector "gitlab.com/tokend/nft-books/book-svc/connector/api"
	networkerConnector "gitlab.com/tokend/nft-books/network-svc/connector/api"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log             *logan.Entry
	copus           types.Copus
	listener        net.Listener
	ethMinterConfig *config.EthMinterConfig
	pricer          *pricer.Connector
	networker       *networkerConnector.Connector
	booker          *booksConnector.Connector
	apiRestrictions config.ApiRestrictions

	booksDB     *pgdb.DB
	generatorDB *pgdb.DB
	trackerDB   *pgdb.DB
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
		log:             cfg.Log(),
		copus:           cfg.Copus(),
		listener:        cfg.Listener(),
		ethMinterConfig: cfg.EthMinter(),
		pricer:          cfg.PricerConnector(),
		booker:          cfg.BooksConnector(),
		networker:       cfg.NetworkConnector(),
		apiRestrictions: cfg.ApiRestrictions(),

		booksDB:     cfg.BookDB().DB,
		generatorDB: cfg.GeneratorDB().DB,
		trackerDB:   cfg.TrackerDB().DB,
	}
}

func Run(ctx context.Context, cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(errors.Wrap(err, "failed to run a service"))
	}
}
