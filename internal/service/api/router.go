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
			helpers.CtxBooksQ(postgres.NewBooksQ(s.booksDB)),
			helpers.CtxPaymentsQ(postgres.NewPaymentsQ(s.trackerDB)),
			helpers.CtxGeneratorDB(postgres.NewGeneratorDB(s.generatorDB)),
			helpers.CtxMinter(*s.ethMinterConfig),
			helpers.CtxPricer(s.pricer),
			helpers.CtxNetworkerConnector(*s.networker),
		),
	)
	r.Route("/integrations/generator", func(r chi.Router) {

		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", handlers.CreateTask)
			r.Get("/", handlers.ListTasks)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTaskByID)
			})
		})

		r.Route("/tokens", func(r chi.Router) {
			r.Get("/", handlers.ListTokens)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTokenById)
			})
		})

		r.Route("/signature", func(r chi.Router) {
			r.Get("/mint", handlers.SignMint)
		})
	})

	return r
}
