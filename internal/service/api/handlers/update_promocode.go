package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdatePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewUpdatePromocodeRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch update promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocodeId := cast.ToInt64(request.Data.ID)
	promocode, err := helpers.DB(r).Promocodes().FilterById(promocodeId).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !validateUsages(*request, *promocode) {
		logger.WithError(err).Info(InvalidUsagesError)
		ape.RenderErr(w, problems.BadRequest(InvalidUsagesError)...)
	}

	promocodesQ, err := applyPromocodeUpdateFilters(r, helpers.DB(r).Promocodes().New(), *request)

	if err != nil { // are already logged
		ape.RenderErr(w, problems.InternalError())
	}
	if promocodesQ == nil {
		ape.RenderErr(w, problems.Forbidden())
	}

	if err = (*promocodesQ).FilterUpdateById(promocodeId).Update(); err != nil {
		logger.WithError(err).Error("failed to update promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if request.Data.Attributes.Books != nil {
		if err = helpers.DB(r).PromocodesBooks().UpdateBooks(promocodeId, *request.Data.Attributes.Books...); err != nil {
			logger.WithError(err).Error("failed to update books in promocode")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func validateUsages(request requests.UpdatePromocodeRequest, promocode data.Promocode) bool {
	return request.Data.Attributes.InitialUsages == nil || promocode.Usages <= *request.Data.Attributes.InitialUsages
}

func applyPromocodeUpdateFilters(r *http.Request, q data.PromocodesQ, request requests.UpdatePromocodeRequest) (*data.PromocodesQ, error) {
	if request.Data.Attributes.State != nil {
		q = q.UpdateState(*request.Data.Attributes.State)
	}
	if request.Data.Attributes.Usages != nil {
		q = q.UpdateUsages(*request.Data.Attributes.Usages)
	}
	if request.Data.Attributes.InitialUsages != nil {
		q = q.UpdateInitialUsages(*request.Data.Attributes.InitialUsages)
	}
	if request.Data.Attributes.ExpirationDate != nil {
		q = q.UpdateExpirationDate(*request.Data.Attributes.ExpirationDate)
	}
	if request.Data.Attributes.Discount != nil {
		q = q.UpdateDiscount(helpers.Trancate(*request.Data.Attributes.Discount, helpers.Promocoder(r).Decimal))
	}
	if request.Data.Attributes.Promocode != nil {
		if *request.Data.Attributes.Promocode == "" {
			q = q.FilterByPromocode(uuid.NewString())
			return &q, nil
		}
		pr, err := helpers.DB(r).Promocodes().FilterByPromocode(*request.Data.Attributes.Promocode).Get()
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to check promocode existing")
			return nil, err
		}
		if pr != nil {
			helpers.Log(r).Error("promocode is already exists")
			return nil, nil
		}
		q = q.UpdatePromocode(*request.Data.Attributes.Promocode)
	}
	return &q, nil
}
