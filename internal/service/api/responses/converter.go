package responses

import (
	booker "github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/core-svc/resources"
)

func convertBookToResource(bookResponse booker.GetBookResponse) resources.Resource {
	var networks []resources.BookNetwork
	for _, network := range bookResponse.Data.Attributes.Networks {
		networks = append(networks, resources.BookNetwork{
			Key: resources.Key{
				ID:   network.ID,
				Type: resources.BOOK_NETWORK,
			},
			Attributes: resources.BookNetworkAttributes{
				ChainId:         network.Attributes.ChainId,
				ContractAddress: network.Attributes.ContractAddress,
				DeployStatus:    resources.DeployStatus(network.Attributes.DeployStatus),
				TokenId:         network.Attributes.TokenId,
			},
		})
	}
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
			CreatedAt:   bookResponse.Data.Attributes.CreatedAt,
			Description: bookResponse.Data.Attributes.Description,
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
			Networks: networks,
		},
	}
}
