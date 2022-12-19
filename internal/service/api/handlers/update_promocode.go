package handlers

import (
	"errors"
	"github.com/spf13/cast"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func UpdatePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewUpdatePromocodeRequest(r)
	if err != nil {
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
		logger.WithError(err).Info("left usages should be lower or equal initial usages")
		ape.RenderErr(w, problems.BadRequest(errors.New("left usages should be lower or equal initial usages"))...)
	}

	promocodesQ := applyPromocodeUpdateFilters(r, helpers.DB(r).Promocodes().New(), *request)

	if err = promocodesQ.Update(promocodeId); err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func validateUsages(request requests.UpdatePromocodeRequest, promocode data.Promocode) bool {
	// if we update only initial_usages
	if request.Data.Attributes.InitialUsages != nil && promocode.LeftUsages > *request.Data.Attributes.InitialUsages {
		return false
	}
	// if we update only left_usages
	if request.Data.Attributes.LeftUsages != nil && promocode.InitialUsages < *request.Data.Attributes.LeftUsages {
		return false
	}
	// if we update initial_usages and left_usages
	if request.Data.Attributes.InitialUsages != nil && request.Data.Attributes.LeftUsages != nil &&
		*request.Data.Attributes.LeftUsages > *request.Data.Attributes.InitialUsages {
		return false
	}
	return true
}

func applyPromocodeUpdateFilters(r *http.Request, q data.PromocodesQ, request requests.UpdatePromocodeRequest) data.PromocodesQ {
	if request.Data.Attributes.State != nil {
		q = q.UpdateState(*request.Data.Attributes.State)
	}
	if request.Data.Attributes.LeftUsages != nil {
		q = q.UpdateLeftUsages(*request.Data.Attributes.LeftUsages)
	}
	if request.Data.Attributes.InitialUsages != nil {
		q = q.UpdateInitialUsages(*request.Data.Attributes.InitialUsages)
	}
	if request.Data.Attributes.ExpirationDate != nil {
		q = q.UpdateExpirationDate(*request.Data.Attributes.ExpirationDate)
	}
	if request.Data.Attributes.Discount != nil {
		q = q.UpdateDiscount(helpers.Trancate(*request.Data.Attributes.Discount, helpers.Promocodes(r).Decimal))
	}
	return q
}
