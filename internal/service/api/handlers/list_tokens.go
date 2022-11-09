package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
)

func ListTokens(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewListTokensRequest(r)
	if err != nil {
		logger.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tokens, err := applyTokensQFilters(helpers.GeneratorDB(r).Tokens(), request).Select()
	if err != nil {
		logger.WithError(err).Error("unable to select tokens from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tokensListResponse, err := responses.NewTokenListResponse(r, request, tokens, helpers.PaymentsQ(r), helpers.GeneratorDB(r).Tasks())
	if err != nil {
		logger.WithError(err).Error("unable to form task list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *tokensListResponse)
}

func applyTokensQFilters(q data.TokensQ, request *requests.ListTokensRequest) data.TokensQ {
	if len(request.Account) > 0 {
		q = q.FilterByAccount(request.Account...)
	}
	if len(request.Status) > 0 {
		q = q.FilterByStatus(request.Status...)
	}

	q = q.Page(request.OffsetPageParams)
	q = q.Sort(request.Sorts)

	return q
}
