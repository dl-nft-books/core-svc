package handlers

import (
	"net/http"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateTokenRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		bookId    = cast.ToInt64(request.Data.Relationships.Book.Data.ID)
		paymentId = cast.ToInt64(request.Data.Relationships.Payment.Data.ID)
	)

	// checking if book exists + retrieving signature
	bookResponse, err := helpers.Booker(r).GetBookById(bookId)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to fetch book with id %v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bookResponse == nil {
		helpers.Log(r).Errorf("book with id %v was not found", bookId)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	token, err := helpers.DB(r).Tokens().FilterByMetadataHash(request.Data.Attributes.MetadataHash).Get()
	if err != nil {
		helpers.Log(r).Errorf("failed to get token")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if token != nil {
		helpers.Log(r).Errorf("token is already exists")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	createdTokenId, err := helpers.DB(r).New().Tokens().Insert(data.Token{
		Account:      request.Data.Attributes.Account,
		TokenId:      request.Data.Attributes.TokenId,
		BookId:       bookId,
		PaymentId:    paymentId,
		MetadataHash: request.Data.Attributes.MetadataHash,
		Status:       request.Data.Attributes.Status,
		ChainId:      request.Data.Attributes.ChainId,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new token with id of #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(createdTokenId, resources.TOKENS),
	})
}
