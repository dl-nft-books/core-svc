package handlers

import (
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"math"
	"net/http"
)

func CreatePromocode(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreatePromocodeRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	prString := uuid.NewString()
	promocode, err := helpers.DB(r).Promocodes().Insert(data.Promocode{
		Promocode:      prString,
		Discount:       math.Floor(request.Data.Attributes.Discount*100) / 100,
		InitialUsages:  request.Data.Attributes.InitialUsages,
		LeftUsages:     request.Data.Attributes.InitialUsages,
		ExpirationDate: request.Data.Attributes.ExpirationDate,
		State:          resources.PromocodeActive,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	promocodeResponse, err := responses.NewGetPromocodeResponse(promocode)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get promocode response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *promocodeResponse)
}
