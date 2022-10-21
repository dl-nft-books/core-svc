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

	tasksId        = "id"
	tasksSignature = "signature"
	tasksIpfsHash  = "ipfs_hash"
	tasksTokenId   = "token_id"
	tasksStatus    = "status"
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

func (q *tasksQ) FilterById(id ...int64) data.TasksQ {
	q.selector = q.selector.Where(squirrel.Eq{tasksId: id})
	return q
}

func (q *tasksQ) FilterByIpfsHash(ipfsHash string) data.TasksQ {
	q.selector = q.selector.Where(squirrel.Eq{tasksIpfsHash: ipfsHash})
	return q
}

func (q *tasksQ) Select() (tasks []data.Task, err error) {
	err = q.database.Select(&tasks, q.selector)
	return
}

func (q *tasksQ) Get() (*data.Task, error) {
	var task data.Task

	err := q.database.Get(&task, q.selector)
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

func (q *tasksQ) UpdateIpfsHash(newIpfsHash string, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksIpfsHash, newIpfsHash).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateTokenId(newTokenId, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksTokenId, newTokenId).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateStatus(newStatus resources.TaskStatus, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksStatus, newStatus).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}
