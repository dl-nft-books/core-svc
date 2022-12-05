package api

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/handlers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/middlewares"

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
			helpers.CtxPricer(s.pricer),
			helpers.CtxApiRestrictions(s.apiRestrictions),
			helpers.CtxBooker(s.booker),
			helpers.CtxTracker(s.tracker),
			helpers.CtxDoormanConnector(s.doorman),
		),
	)

	r.Route("/integrations/generator", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.With(middlewares.CheckAccessToken).Post("/", handlers.CreateTask)
			r.Get("/", handlers.ListTasks)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTaskByID)
				r.With(middlewares.CheckAccessToken).Patch("/", handlers.UpdateTask)
			})
		})

		r.Route("/tokens", func(r chi.Router) {
			r.Get("/", handlers.ListTokens)
			r.With(middlewares.CheckAccessToken).Post("/", handlers.CreateToken)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTokenById)
				r.With(middlewares.CheckAccessToken).Patch("/", handlers.UpdateToken)
			})
		})

		r.Route("/signature", func(r chi.Router) {
			r.Get("/mint", handlers.SignMint)
		})
	})

	return r
}
