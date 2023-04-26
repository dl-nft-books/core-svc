package responses

import (
	"github.com/dl-nft-books/core-svc/resources"
)

func NewBuyVoucherResponse(transaction string) *resources.TransactionResponse {
	return &resources.TransactionResponse{
		Data: resources.Transaction{
			Key: resources.Key{
				ID:   "0",
				Type: resources.TRANSACTION,
			},
			Attributes: resources.TransactionAttributes{
				TxHash: transaction,
			},
		},
	}
}
