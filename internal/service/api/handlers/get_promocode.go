package handlers

import (
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
)

func GetPromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewPromocodeByIdRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch get promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocode, err := helpers.DB(r).Promocodes().FilterById(request.Id).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		logger.Error("promocode with such id is not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	promocodeResponse := responses.NewGetPromocodeResponse(*promocode)

	ape.Render(w, *promocodeResponse)
}
