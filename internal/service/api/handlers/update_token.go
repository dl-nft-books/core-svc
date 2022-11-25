package handlers

import (
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
)

func UpdateToken(w http.ResponseWriter, r *http.Request) {
	// Getting the update token request
	request, err := requests.NewUpdateTokenRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch update token request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tokensQ := helpers.GeneratorDB(r).Tokens()

	// Validating whether specified task exists
	tokenId := cast.ToInt64(request.Data.ID)
	token, err := tokensQ.FilterByTokenId(tokenId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to get a token with id of %v", tokenId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if token == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Clearing selector filters and applying updator ones
	tokensQ = applyTokenUpdateFilters(tokensQ.New(), *request)

	if err = tokensQ.Update(tokenId); err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to update token with id of #%v", tokenId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func applyTokenUpdateFilters(q data.TokensQ, request resources.UpdateTokenRequest) data.TokensQ {
	if request.Data.Attributes.TokenId != nil {
		q = q.UpdateTokenId(int64(*request.Data.Attributes.TokenId))
	}
	if request.Data.Attributes.Owner != nil {
		q = q.UpdateOwner(*request.Data.Attributes.Owner)
	}
	if request.Data.Attributes.Status != nil {
		q = q.UpdateStatus(*request.Data.Attributes.Status)
	}

	return q
}
