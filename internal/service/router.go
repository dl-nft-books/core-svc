package service

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/handlers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/helpers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxBooksQ(postgres.NewBooksQ(s.db)),
			helpers.CtxMinter(*s.ethMinterConfig),
			helpers.CtxCoingecko(*s.coingeckoConfig),
		),
	)
	r.Route("/integrations/price", func(r chi.Router) {
		r.Get("/{id}", handlers.GetPrice)
	})

	return r
}
