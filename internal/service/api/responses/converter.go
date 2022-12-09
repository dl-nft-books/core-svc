package responses

import (
	booker "gitlab.com/tokend/nft-books/book-svc/connector/models"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector/models"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func convertPaymentToResource(paymentResponse tracker.GetPaymentResponse) resources.Resource {
	return &resources.Payment{
		Key: resources.Key{
			ID:   paymentResponse.Data.Key.ID,
			Type: resources.ResourceType(paymentResponse.Data.Key.Type),
		},
		Attributes: resources.PaymentAttributes{
			Amount:            paymentResponse.Data.Attributes.Amount,
			BookUrl:           paymentResponse.Data.Attributes.BookUrl,
			Erc20Data:         resources.Erc20Data(paymentResponse.Data.Attributes.Erc20Data),
			MintedTokenPrice:  paymentResponse.Data.Attributes.MintedTokenPrice,
			PayerAddress:      paymentResponse.Data.Attributes.PayerAddress,
			PaymentTokenPrice: paymentResponse.Data.Attributes.PaymentTokenPrice,
			PurchaseTimestamp: paymentResponse.Data.Attributes.PurchaseTimestamp,
		},
		Relationships: nil,
	}
}

func convertBookToResource(bookResponse booker.GetBookResponse) resources.Resource {
	return &resources.Book{
		Key: resources.Key{
			ID:   bookResponse.Data.Key.ID,
			Type: resources.ResourceType(bookResponse.Data.Key.Type),
		},
		Attributes: resources.BookAttributes{
			Banner: resources.Media{
				Key: resources.Key{
					ID:   bookResponse.Data.Attributes.Banner.Key.ID,
					Type: resources.ResourceType(bookResponse.Data.Attributes.Banner.Key.Type),
				},
				Attributes: resources.MediaAttributes{
					Key:      bookResponse.Data.Attributes.Banner.Attributes.MimeType,
					MimeType: bookResponse.Data.Attributes.Banner.Attributes.MimeType,
					Name:     bookResponse.Data.Attributes.Banner.Attributes.Name,
					Url:      bookResponse.Data.Attributes.Banner.Attributes.Url,
				},
			},
			ContractAddress: bookResponse.Data.Attributes.ContractAddress,
			ContractName:    bookResponse.Data.Attributes.ContractName,
			ContractSymbol:  bookResponse.Data.Attributes.ContractSymbol,
			ContractVersion: bookResponse.Data.Attributes.ContractVersion,
			CreatedAt:       bookResponse.Data.Attributes.CreatedAt,
			DeployStatus:    resources.DeployStatus(bookResponse.Data.Attributes.DeployStatus),
			Description:     bookResponse.Data.Attributes.Description,
			File: resources.Media{
				Key: resources.Key{
					ID:   bookResponse.Data.Attributes.File.Key.ID,
					Type: resources.ResourceType(bookResponse.Data.Attributes.File.Key.Type),
				},
				Attributes: resources.MediaAttributes{
					Key:      bookResponse.Data.Attributes.File.Attributes.MimeType,
					MimeType: bookResponse.Data.Attributes.File.Attributes.MimeType,
					Name:     bookResponse.Data.Attributes.File.Attributes.Name,
					Url:      bookResponse.Data.Attributes.File.Attributes.Url,
				},
			},
			Price:   bookResponse.Data.Attributes.Price,
			Title:   bookResponse.Data.Attributes.Title,
			TokenId: bookResponse.Data.Attributes.TokenId,
		},
	}
}
