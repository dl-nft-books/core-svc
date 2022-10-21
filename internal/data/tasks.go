package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type Task struct {
	Id        int64                `db:"id" structs:"-" json:"-"`
	Account   string               `db:"account" structs:"account"`
	Signature string               `db:"signature" structs:"signature"`
	IpfsHash  string               `db:"ipfs_hash" structs:"ipfs_hash"`
	TokenId   int64                `db:"token_id" structs:"token_id"`
	Status    resources.TaskStatus `db:"status" structs:"status"`
}

type TasksQ interface {
	New() TasksQ

	Get() (*Task, error)
	Select() ([]Task, error)

	Sort(sort pgdb.Sorts) TasksQ
	Page(page pgdb.OffsetPageParams) TasksQ

	FilterById(id ...int64) TasksQ
	FilterByIpfsHash(ipfsHash string) TasksQ

	Update(task Task, id int64) error
	Insert(task Task) (id int64, err error)
	Delete(id int64) error

	UpdateIpfsHash(newIpfsHash string, id int64) error
	UpdateTokenId(newTokenId, id int64) error
	UpdateStatus(newStatus resources.TaskStatus, id int64) error
}
