package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const (
	tasksTable = "tasks"

	tasksId               = "id"
	tasksAccount          = "account"
	tasksSignature        = "signature"
	tasksFileIpfsHash     = "file_ipfs_hash"
	tasksMetadataIpfsHash = "metadata_ipfs_hash"
	tasksTokenId          = "token_id"
	tasksStatus           = "status"
	tasksUri              = "uri"
	tasksCreatedAt        = "created_at"
)

type tasksQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewTasksQ(database *pgdb.DB) data.TasksQ {
	return &tasksQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tasksTable)).From(tasksTable),
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

func (q *tasksQ) Update(task data.Task, id int64) error {
	statement := squirrel.Update(tasksTable).
		SetMap(structs.Map(&task)).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateFileIpfsHash(newIpfsHash string, id int64) error {
	return q.updateHash(tasksFileIpfsHash, newIpfsHash, id)
}

func (q *tasksQ) UpdateMetadataIpfsHash(newIpfsHash string, id int64) error {
	return q.updateHash(tasksMetadataIpfsHash, newIpfsHash, id)
}

func (q *tasksQ) updateHash(fieldName, newIpfsHash string, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(fieldName, newIpfsHash).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateTokenId(newTokenId, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksTokenId, newTokenId).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateUri(newUri string, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksUri, newUri).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateStatus(newStatus resources.TaskStatus, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksStatus, newStatus).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
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
	if selector.IpfsHash != nil {
		sql = sql.Where(squirrel.Eq{tasksMetadataIpfsHash: *selector.IpfsHash})
	}
	if selector.Status != nil {
		sql = sql.Where(squirrel.Eq{tasksStatus: *selector.Status})
	}
	if selector.OffsetParams != nil {
		return selector.OffsetParams.ApplyTo(sql, tasksId)
	}
	if selector.PageParams != nil {
		return selector.PageParams.ApplyTo(sql, tasksId)
	}

	return sql
}
