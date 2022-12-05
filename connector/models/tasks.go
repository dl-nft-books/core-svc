package models

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type (
	ListTasksRequest  requests.ListTasksRequest
	ListTasksResponse resources.TaskListResponse

	UpdateTaskParams struct {
		Id      int64                 `json:"-"`
		Status  *resources.TaskStatus `json:"status,omitempty"`
		TokenId *int64                `json:"token_id,omitempty"`
	}

	CreateTaskParams struct {
		Account   string `json:"account"`
		BookId    int64  `json:"book_id"`
		Signature string `json:"signature"`
	}

	TaskResponse resources.TaskResponse
)
