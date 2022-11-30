package models

import (
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type ListTasksRequest struct {
	Limit    *uint64               `structs:"page[limit]"`
	Cursor   *uint64               `structs:"page[cursor]"`
	Account  *string               `structs:"filter[account]"`
	Status   *resources.TaskStatus `structs:"filter[status]"`
	IpfsHash *string               `structs:"filter[ipfs_hash]"`
}
