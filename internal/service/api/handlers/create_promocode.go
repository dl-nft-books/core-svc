package handlers

import (
	"github.com/dl-nft-books/book-svc/connector"
	"github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

func CreatePromocode(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreatePromocodeRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	address := r.Context().Value("address").(string)
	isMarketplaceManager, err := helpers.CheckMarketplacePerrmision(*helpers.Networker(r), address)
	if err != nil {
		helpers.Log(r).WithError(err).WithFields(logan.F{"account": address}).Debug("failed to check permissions")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if !isMarketplaceManager {
		helpers.Log(r).WithFields(logan.F{"account": address}).Debug("you don't have permission to create book")
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	prString := uuid.NewString()
	if request.Data.Attributes.Promocode != nil && *request.Data.Attributes.Promocode != "" {
		pr, err := helpers.DB(r).Promocodes().FilterByPromocode(*request.Data.Attributes.Promocode).Get()
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to check promocode existing")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if pr != nil {
			helpers.Log(r).Error("promocode is already exists")
			ape.RenderErr(w, problems.Forbidden())
			return
		}
		prString = *request.Data.Attributes.Promocode
	}
	books, err := getBooks(request.Data.Attributes.Books, helpers.Booker(r))
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to get books for promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var promocodeID int64
	if err = helpers.DB(r).Transaction(func() error {
		promocodeID, err = helpers.DB(r).Promocodes().Insert(data.Promocode{
			Promocode:      prString,
			Discount:       helpers.Trancate(request.Data.Attributes.Discount, helpers.Promocoder(r).Decimal),
			InitialUsages:  request.Data.Attributes.InitialUsages,
			Usages:         0,
			ExpirationDate: request.Data.Attributes.ExpirationDate,
			State:          resources.PromocodeActive,
		})
		if err != nil {
			return err
		}
		return helpers.DB(r).PromocodesBooks().Insert(promocodeID, books...)
	}); err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(promocodeID, resources.PROMOCODE),
	})
}

func getBooks(books *[]int64, booker *connector.Connector) ([]int64, error) {
	var bookIds []int64
	if books != nil && len(*books) > 0 {
		return *books, nil
	}
	booksResponse, err := booker.ListBooks(models.ListBooksParams{})
	if err != nil {
		return bookIds, errors.Wrap(err, "failed to get books")
	}
	for _, book := range booksResponse.Data {
		bookIds = append(bookIds, cast.ToInt64(book.ID))
	}
	return bookIds, nil
}
