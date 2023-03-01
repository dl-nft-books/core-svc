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
			// Base configs
			helpers.CtxLog(s.log),
			helpers.CtxDB(postgres.NewDB(s.db)),

			// Custom configs
			helpers.CtxMinter(*s.ethMinterConfig),
			helpers.CtxApiRestrictions(s.apiRestrictions),
			helpers.CtxPromocoder(s.promocoder),
			helpers.CtxIpfser(s.ipfser),

			// Connectors
			helpers.CtxPricer(s.pricer),
			helpers.CtxBooker(s.booker),
			helpers.CtxTracker(s.tracker),
			helpers.CtxDoormanConnector(s.doorman),
		),
	)

	r.Route("/integrations/generator", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", handlers.CreateTask)
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

		r.Route("/promocodes", func(r chi.Router) {
			r.Get("/", handlers.ListPromocodes)
			r.Post("/", handlers.CreatePromocode)

			r.Get("/validate/{promocode}", handlers.ValidatePromocodeById)

			r.Patch("/rollback/{id}", handlers.RollbackPromocode)
			r.Patch("/use/{id}", handlers.UsePromocode)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPromocodeById)
				r.Delete("/", handlers.DeletePromocodeById)
				r.Patch("/", handlers.UpdatePromocodeById)
			})
		})

		r.Route("/signature", func(r chi.Router) {
			r.Get("/mint", handlers.SignMint)
			r.Get("/mint/nft", handlers.SignMintByNft)
		})
	})

	return r
}
