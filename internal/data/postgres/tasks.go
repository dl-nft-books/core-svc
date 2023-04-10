package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

const (
	tasksTable = "tasks"

	tasksId               = "id"
	tasksAccount          = "account"
	tasksFileIpfsHash     = "file_ipfs_hash"
	tasksMetadataIpfsHash = "metadata_ipfs_hash"
	tasksTokenId          = "token_id"
	tasksChainId          = "chain_id"
	tasksStatus           = "status"
	tasksUri              = "book_uri"
	TasksCreatedAt        = "created_at"
)

type tasksQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewTasksQ(database *pgdb.DB) data.TasksQ {
	return &tasksQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tasksTable)).From(tasksTable),
		updater:  squirrel.Update(tasksTable),
	}
}

func (q *tasksQ) New() data.TasksQ {
	return NewTasksQ(q.database.Clone())
}

func (q *tasksQ) Page(page pgdb.OffsetPageParams) data.TasksQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *tasksQ) Sort(sort pgdb.Sorts) data.TasksQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id":         "id",
		"created_at": "created_at",
	})

	return q
}

func (q *tasksQ) Select(selector data.TaskSelector) (tasks []data.Task, err error) {
	return q.selectByQuery(applyTasksSelector(q.selector, selector))
}

func (q *tasksQ) GetById(id int64) (*data.Task, error) {
	var task data.Task

	err := q.database.Get(&task, q.selector.Where(squirrel.Eq{tasksId: id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &task, err
}

func (q *tasksQ) Insert(task data.Task) (id int64, err error) {
	statement := squirrel.Insert(tasksTable).
		Suffix("returning id").
		SetMap(structs.Map(&task))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *tasksQ) Delete(id int64) error {
	statement := squirrel.Delete(tasksTable).Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateBannerIpfsHash(newIpfsHash string) data.TasksQ {
	q.updater = q.updater.Set(tasksFileIpfsHash, newIpfsHash)
	return q
}

func (q *tasksQ) UpdateMetadataIpfsHash(newIpfsHash string) data.TasksQ {
	q.updater = q.updater.Set(tasksMetadataIpfsHash, newIpfsHash)
	return q
}

func (q *tasksQ) updateHash(fieldName, newIpfsHash string) data.TasksQ {
	q.updater = q.updater.Set(fieldName, newIpfsHash)
	return q
}

func (q *tasksQ) UpdateTokenId(newTokenId int64) data.TasksQ {
	q.updater = q.updater.Set(tasksTokenId, newTokenId)
	return q
}

func (q *tasksQ) UpdateUri(newUri string) data.TasksQ {
	q.updater = q.updater.Set(tasksUri, newUri)
	return q
}

func (q *tasksQ) UpdateStatus(newStatus resources.TaskStatus) data.TasksQ {
	q.updater = q.updater.Set(tasksStatus, newStatus)
	return q
}

func (q *tasksQ) Update(id int64) error {
	return q.database.Exec(q.updater.Where(squirrel.Eq{tasksId: id}))
}

func (q *tasksQ) Transaction(fn func(q data.TasksQ) error) (err error) {
	return q.database.Transaction(func() error {
		return fn(q)
	})
}

func (q *tasksQ) selectByQuery(query squirrel.Sqlizer) (subtasks []data.Task, err error) {
	err = q.database.Select(&subtasks, query)
	return subtasks, err
}

func applyTasksSelector(sql squirrel.SelectBuilder, selector data.TaskSelector) squirrel.SelectBuilder {
	if selector.Account != nil {
		sql = sql.Where(squirrel.Eq{tasksAccount: *selector.Account})
	}
	if selector.TokenId != nil {
		sql = sql.Where(squirrel.Eq{tasksTokenId: *selector.TokenId})
	}
	if selector.ChainId != nil {
		sql = sql.Where(squirrel.Eq{tasksChainId: *selector.ChainId})
	}
	if selector.IpfsHash != nil {
		sql = sql.Where(squirrel.Eq{tasksMetadataIpfsHash: *selector.IpfsHash})
	}
	if selector.Status != nil {
		sql = sql.Where(squirrel.Eq{tasksStatus: *selector.Status})
	}
	if selector.Period != nil {
		expirationDate := time.Now().UTC().Add(-(*selector.Period))
		sql = sql.Where(squirrel.LtOrEq{TasksCreatedAt: expirationDate})
	}
	if selector.OffsetParams != nil {
		return selector.OffsetParams.ApplyTo(sql, tasksId)
	}
	if selector.PageParams != nil {
		return selector.PageParams.ApplyTo(sql, tasksId)
	}

	return sql
}
