package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"time"
)

const (
	timestampFormat = "2006-01-02"
	base            = 10
)

var ConversionFromStringToBigIntErr = errors.New("failed to convert string to the big int format")

type Payment struct {
	Id                int64     `db:"id" structs:"-" json:"-"`
	ContractId        int64     `db:"contract_id" structs:"contract_id" json:"contract_id"`
	ContractAddress   string    `db:"contract_address" structs:"contract_address"`
	PayerAddress      string    `db:"payer_address" structs:"payer_address"`
	TokenAddress      string    `db:"token_address" structs:"token_address"`
	TokenSymbol       string    `db:"token_symbol" structs:"token_symbol"`
	TokenName         string    `db:"token_name" structs:"token_name"`
	TokenDecimals     uint8     `db:"token_decimals" structs:"token_decimals"`
	Amount            string    `db:"amount" structs:"amount"`
	PriceToken        string    `db:"price_token" structs:"price_token"`
	PriceMinted       string    `db:"price_minted" structs:"price_minted"`
	BookUrl           string    `db:"book_url" structs:"book_url"`
	PurchaseTimestamp time.Time `db:"purchase_timestamp" structs:"purchase_timestamp"`
}

type PaymentsQ interface {
	New() PaymentsQ

	FilterById(id ...int64) PaymentsQ
	FilterByPayer(payer ...string) PaymentsQ
	FilterByTokenAddress(tokenAddress ...string) PaymentsQ
	FilterByContractAddress(contractAddress ...string) PaymentsQ
	FilterByContractId(contractId ...int64) PaymentsQ

	Get() (*Payment, error)
	Select() ([]Payment, error)

	Sort(sort pgdb.Sorts) PaymentsQ
	Page(page pgdb.OffsetPageParams) PaymentsQ

	Insert(payment Payment) (id int64, err error)
	Delete(id int64) error
}

func (p *Payment) Resource() (*resources.Payment, error) {
	return &resources.Payment{
		Key: resources.NewKeyInt64(p.Id, resources.PAYMENT),
		Attributes: resources.PaymentAttributes{
			Amount:            p.Amount,
			PayerAddress:      p.PayerAddress,
			PaymentTokenPrice: p.PriceToken,
			MintedTokenPrice:  p.PriceMinted,
			PurchaseTimestamp: p.PurchaseTimestamp.Format(timestampFormat),
			BookUrl:           p.BookUrl,
			Erc20Data: resources.Erc20Data{
				Address:  p.TokenAddress,
				Name:     p.TokenName,
				Symbol:   p.TokenSymbol,
				Decimals: int32(p.TokenDecimals),
			},
		},
	}, nil
}
