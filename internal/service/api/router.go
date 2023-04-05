package api

import (
	"github.com/dl-nft-books/core-svc/internal/data/postgres"
	"github.com/dl-nft-books/core-svc/internal/service/api/handlers"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/middlewares"
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

		r.Route("/promocodes", func(r chi.Router) {
			r.With(middlewares.CheckAccessToken).Get("/", handlers.ListPromocodes)
			r.With(middlewares.CheckAccessToken).Post("/", handlers.CreatePromocode)

			r.Get("/validate/{promocode}", handlers.ValidatePromocodeById)

			r.Route("/{id}", func(r chi.Router) {
				r.With(middlewares.CheckAccessToken).Get("/", handlers.GetPromocodeById)
				r.With(middlewares.CheckAccessToken).Delete("/", handlers.DeletePromocodeById)
				r.With(middlewares.CheckAccessToken).Patch("/", handlers.UpdatePromocodeById)
			})
		})
		r.Route("/nft-request", func(r chi.Router) {
			r.Post("/", handlers.CreateNftRequest)
			r.Get("/", handlers.ListNftRequests)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetNftRequestById)
				r.With(middlewares.CheckAccessToken).Patch("/", handlers.UpdateNftRequestById)
			})
		})
		r.Route("/signature", func(r chi.Router) {
			r.Get("/mint", handlers.SignMint)
			r.Get("/mint/nft", handlers.SignMintByNft)
		})
	})

	return r
}
