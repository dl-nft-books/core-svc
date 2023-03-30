package data

import (
	"github.com/dl-nft-books/core-svc/resources"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type Task struct {
	Id               int64                `db:"id" structs:"-" json:"-"`
	BookId           int64                `db:"book_id" structs:"book_id" json:"book_id"`
	TokenId          int64                `db:"token_id" structs:"token_id"`
	ChainId          int64                `db:"chain_id" structs:"chain_id"`
	Account          string               `db:"account" structs:"account"`
	BannerIpfsHash   string               `db:"banner_ipfs_hash" structs:"banner_ipfs_hash"`
	MetadataIpfsHash string               `db:"metadata_ipfs_hash" structs:"metadata_ipfs_hash"`
	Uri              string               `db:"uri" structs:"uri"`
	Banner           string               `db:"banner" structs:"banner"`
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
	TokenId      *int64
	ChainId      *int64
	Period       *time.Duration
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

	UpdateBannerIpfsHash(newIpfsHash string) TasksQ
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
			BannerIpfsHash:   t.BannerIpfsHash,
			MetadataIpfsHash: t.MetadataIpfsHash,
			Uri:              t.Uri,
			Status:           t.Status,
		},
	}
}
