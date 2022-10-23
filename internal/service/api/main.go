package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log             *logan.Entry
	copus           types.Copus
	listener        net.Listener
	booksDB         *pgdb.DB
	generatorDB     *pgdb.DB
	ethMinterConfig *config.EthMinterConfig
	coingeckoConfig *config.CoingeckoConfig
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
		booksDB:         cfg.BookDB().DB,
		generatorDB:     cfg.GeneratorDB().DB,
		ethMinterConfig: cfg.EthMinter(),
		coingeckoConfig: cfg.Coingecko(),
	}
}

func Run(ctx context.Context, cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
