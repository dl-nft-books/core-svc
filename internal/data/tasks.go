package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type Task struct {
	Id        int64                `db:"id" structs:"-" json:"-"`
	BookId    int64                `db:"book_id" structs:"-" json:"-"`
	TokenId   int64                `db:"token_id" structs:"token_id"`
	Account   string               `db:"account" structs:"account"`
	Signature string               `db:"signature" structs:"signature"`
	IpfsHash  string               `db:"ipfs_hash" structs:"ipfs_hash"`
	Status    resources.TaskStatus `db:"status" structs:"status"`
}

// TaskSelector is a structure for all applicable filters and params on tasksQ `Select`
type TaskSelector struct {
	PageParams pgdb.CursorPageParams
	Account    *string
	IpfsHash   *string
	Status     *resources.TaskStatus
}

type TasksQ interface {
	New() TasksQ

	GetById(id int64) (*Task, error)
	Select(selector TaskSelector) ([]Task, error)

	Sort(sort pgdb.Sorts) TasksQ
	Page(page pgdb.OffsetPageParams) TasksQ

	Update(task Task, id int64) error
	Insert(task Task) (id int64, err error)
	Delete(id int64) error
	Transaction(fn func(q TasksQ) error) error

	UpdateIpfsHash(newIpfsHash string, id int64) error
	UpdateTokenId(newTokenId, id int64) error
	UpdateStatus(newStatus resources.TaskStatus, id int64) error
}

func (t Task) Resource() resources.Task {
	return resources.Task{
		Key: resources.NewKeyInt64(t.Id, resources.TASKS),
		Attributes: resources.TaskAttributes{
			TokenId:   int32(t.TokenId),
			BookId:    int32(t.BookId),
			IpfsHash:  t.IpfsHash,
			Signature: t.Signature,
			Status:    t.Status,
		},
	}
}
