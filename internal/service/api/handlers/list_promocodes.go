package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
)

func ListPromocodes(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewListPromocodesRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch list promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	address := r.Context().Value("address").(string)
	isMarketplaceManager, err := helpers.CheckMarketplacePerrmision(*helpers.Networker(r), address)
	if err != nil {
		logger.WithError(err).WithFields(logan.F{"account": address}).Debug("failed to check permissions")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if !isMarketplaceManager {
		logger.WithFields(logan.F{"account": address}).Debug("you don't have permission to create book")
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	promocodes, err := applyPromocodesQFilters(helpers.DB(r).Promocodes(), request).Select()
	if err != nil {
		logger.WithError(err).Error("unable to select promocodes from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	promocodesListResponse, err := responses.NewPromocodeListResponse(r, request, promocodes)
	if err != nil {
		logger.WithError(err).Error("unable to form promocode list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, *promocodesListResponse)
}

func applyPromocodesQFilters(q data.PromocodesQ, request *requests.ListPromocodesRequest) data.PromocodesQ {
	if len(request.State) > 0 {
		q = q.FilterByState(request.State...)
	}

	q = q.Page(request.OffsetPageParams)
	q = q.Sort(request.Sorts)

	return q
}
