package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
	"strconv"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	// Getting the create token request
	request, err := requests.NewCreateTokenRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	bookId, err := strconv.Atoi(request.Data.Relationships.Book.Data.ID)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to convert book id from string format to the integer one", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	paymentId, err := strconv.Atoi(request.Data.Relationships.Payment.Data.ID)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to convert book id from string format to the integer one", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	createdTokenId, err := helpers.GeneratorDB(r).Tokens().Insert(data.Token{
		Account:      request.Data.Attributes.Account,
		TokenId:      int64(request.Data.Attributes.TokenId),
		BookId:       int64(bookId),
		PaymentId:    int64(paymentId),
		MetadataHash: request.Data.Attributes.MetadataHash,
		Status:       request.Data.Attributes.Status,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new task with book id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(createdTokenId, resources.TOKENS),
	})
}
