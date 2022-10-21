package api

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/handlers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"

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
			helpers.CtxDB(postgres.NewDB(s.db)),
			helpers.CtxMinter(*s.ethMinterConfig),
			helpers.CtxCoingecko(*s.coingeckoConfig),
		),
	)
	r.Route("/integrations/generator", func(r chi.Router) {
		r.Route("/price", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPrice)
			})
		})
		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", handlers.CreateTask)
		})
	})

	return r
}
