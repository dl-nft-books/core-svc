package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	paymentsTable = "payments"

	paymentsId              = "id"
	paymentsContractId      = "contract_id"
	paymentsContractAddress = "contract_address"
	paymentsPayerAddress    = "payer_address"
	paymentsTokenAddress    = "token_address"
	paymentsAmount          = "amount"
	paymentsPrice           = "price"
)

type paymentsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewPaymentsQ(database *pgdb.DB) external.PaymentsQ {
	return &paymentsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", paymentsTable)).From(paymentsTable),
	}
}

func (q *paymentsQ) New() external.PaymentsQ {
	return NewPaymentsQ(q.database.Clone())
}

func (q *paymentsQ) Page(page pgdb.OffsetPageParams) external.PaymentsQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *paymentsQ) Sort(sort pgdb.Sorts) external.PaymentsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *paymentsQ) FilterById(id ...int64) external.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsId: id})
	return q
}

func (q *paymentsQ) FilterByPayer(payer ...string) external.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsPayerAddress: payer})
	return q
}

func (q *paymentsQ) FilterByTokenAddress(tokenAddress ...string) external.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsTokenAddress: tokenAddress})
	return q
}

func (q *paymentsQ) FilterByContractId(contractId ...int64) external.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsContractId: contractId})
	return q
}

func (q *paymentsQ) FilterByContractAddress(contractAddress ...string) external.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsContractAddress: contractAddress})
	return q
}

func (q *paymentsQ) Select() (payments []external.Payment, err error) {
	err = q.database.Select(&payments, q.selector)
	return
}

func (q *paymentsQ) Get() (*external.Payment, error) {
	var payment external.Payment
	err := q.database.Get(&payment, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &payment, err
}

func (q *paymentsQ) Insert(payment external.Payment) (id int64, err error) {
	statement := squirrel.Insert(paymentsTable).
		Suffix("returning id").
		SetMap(structs.Map(&payment))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *paymentsQ) Delete(id int64) error {
	statement := squirrel.Delete(paymentsTable).Where(squirrel.Eq{paymentsId: id})
	return q.database.Exec(statement)
}
