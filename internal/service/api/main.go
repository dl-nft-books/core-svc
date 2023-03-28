package api

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"
	doorman "gitlab.com/tokend/nft-books/doorman/connector"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"
	"net"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	// Base configs
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	db       *pgdb.DB

	// Custom configs
	ethMinterConfig *config.MintConfig
	apiRestrictions config.ApiRestrictions
	promocoder      config.PromocoderCfg
	ipfser          config.IpfserCfg
	// Connectors
	pricer  *pricer.Connector
	booker  *booker.Connector
	tracker *tracker.Connector
	doorman doorman.ConnectorI
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
		// Base configs
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		db:       cfg.DB(),

		// Custom configs
		ethMinterConfig: cfg.MintConfig(),
		apiRestrictions: cfg.ApiRestrictions(),
		promocoder:      cfg.PromocoderCfg(),
		ipfser:          cfg.IpfserCfg(),
		// Connectors
		pricer:  cfg.PricerConnector(),
		booker:  cfg.BookerConnector(),
		tracker: cfg.TrackerConnector(),
		doorman: cfg.DoormanConnector(),
	}
}

func Run(ctx context.Context, cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(errors.Wrap(err, "failed to run a service"))
	}
}
