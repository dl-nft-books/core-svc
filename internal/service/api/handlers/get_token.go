package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func GetTokenById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewTokenByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	token, err := helpers.DB(r).Tokens().FilterById(request.Id).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if token == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	tokenResponse, err := responses.NewGetTokenResponse(*token, helpers.Tracker(r))
	if err != nil {
		logger.WithError(err).Error("failed to get token response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *tokenResponse)
}
