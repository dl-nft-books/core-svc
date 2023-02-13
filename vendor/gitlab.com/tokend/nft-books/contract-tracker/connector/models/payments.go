package models

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

type (
	// ListPaymentsParams is a helper struct to be included when calling ListPayments request
	ListPaymentsParams    requests.ListPaymentsRequest
	ListNftPaymentsParams requests.ListNftPaymentsRequest

	GetPaymentResponse      resources.PaymentResponse
	ListPaymentsResponse    resources.PaymentListResponse
	GetNftPaymentResponse   resources.NftPaymentResponse
	ListNftPaymentsResponse resources.NftPaymentListResponse
)
