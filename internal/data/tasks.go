package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type Task struct {
	Id               int64                `db:"id" structs:"-" json:"-"`
	BookId           int64                `db:"book_id" structs:"book_id" json:"book_id"`
	TokenId          int64                `db:"token_id" structs:"token_id"`
	Account          string               `db:"account" structs:"account"`
	Signature        string               `db:"signature" structs:"signature"`
	FileIpfsHash     string               `db:"file_ipfs_hash" structs:"file_ipfs_hash"`
	MetadataIpfsHash string               `db:"metadata_ipfs_hash" structs:"metadata_ipfs_hash"`
	Uri              string               `db:"uri" structs:"uri"`
	Status           resources.TaskStatus `db:"status" structs:"status"`
	CreatedAt        time.Time            `db:"created_at" structs:"created_at"`
}

// TaskSelector is a structure for all applicable filters and params on tasksQ `Select`
type TaskSelector struct {
	PageParams   *pgdb.CursorPageParams
	OffsetParams *pgdb.OffsetPageParams
	Account      *string
	IpfsHash     *string
	Status       *resources.TaskStatus
}

type TasksQ interface {
	New() TasksQ

	GetById(id int64) (*Task, error)
	Select(selector TaskSelector) ([]Task, error)

	Sort(sort pgdb.Sorts) TasksQ
	Page(page pgdb.OffsetPageParams) TasksQ

	Insert(task Task) (id int64, err error)
	Delete(id int64) error
	Transaction(fn func(q TasksQ) error) error

	UpdateFileIpfsHash(newIpfsHash string) TasksQ
	UpdateMetadataIpfsHash(newIpfsHash string) TasksQ
	UpdateUri(newUri string) TasksQ
	UpdateTokenId(newTokenId int64) TasksQ
	UpdateStatus(newStatus resources.TaskStatus) TasksQ
	Update(id int64) error
}

func (t Task) Resource() resources.Task {
	return resources.Task{
		Key: resources.NewKeyInt64(t.Id, resources.TASKS),
		Attributes: resources.TaskAttributes{
			TokenId:          t.TokenId,
			BookId:           t.BookId,
			FileIpfsHash:     t.FileIpfsHash,
			MetadataIpfsHash: t.MetadataIpfsHash,
			Uri:              t.Uri,
			Signature:        t.Signature,
			Status:           t.Status,
		},
	}
}
