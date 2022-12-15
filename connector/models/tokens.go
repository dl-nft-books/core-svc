package models

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type (
	ListTokensRequest  requests.ListTokensRequest
	ListTokensResponse resources.TokenListResponse

	UpdateTokenParams struct {
		Id int64 `json:"-"`
		// Address of a user who purchased this token
		Owner        *string                `json:"owner,omitempty"`
		Status       *resources.TokenStatus `json:"status,omitempty"`
		TokenId      *int64                 `json:"token_id,omitempty"`
		MetadataHash *string                `json:"metadata_hash,omitempty"`
	}

	CreateTokenParams struct {
		// Address of a user who purchased this token
		Account string `json:"account"`
		// Hash of a metadata file
		MetadataHash string                `json:"metadata_hash"`
		Status       resources.TokenStatus `json:"status"`
		TokenId      int64                 `json:"token_id"`
		Signature    string                `json:"signature"`

		//relations
		BookId    int64
		PaymentId int64
	}

	TokenResponse resources.TokenResponse
)
