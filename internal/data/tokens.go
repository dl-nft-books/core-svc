package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type Token struct {
	Id           int64                 `db:"id" structs:"-" json:"-"`
	Account      string                `db:"account" structs:"account"`
	TokenId      int64                 `db:"token_id" structs:"token_id"`
	BookId       int64                 `db:"book_id" structs:"book_id" json:"book_id"`
	PaymentId    int64                 `db:"payment_id" structs:"payment_id" json:"payment_id"`
	MetadataHash string                `db:"metadata_hash" structs:"metadata_hash" json:"metadata_hash"`
	Signature    string                `db:"signature" structs:"signature" json:"signature"`
	Status       resources.TokenStatus `db:"status" structs:"status" json:"status"`
}

type TokensQ interface {
	New() TokensQ

	FilterByAccount(account ...string) TokensQ
	FilterByStatus(status ...resources.TokenStatus) TokensQ
	FilterById(id ...int64) TokensQ
	FilterByTokenId(tokenId ...int64) TokensQ
	FilterByBookId(bookId ...int64) TokensQ
	FilterByPaymentId(paymentId ...int64) TokensQ
	FilterByMetadataHash(metadataHash ...string) TokensQ

	Get() (*Token, error)
	Select() ([]Token, error)

	Sort(sort pgdb.Sorts) TokensQ
	Page(page pgdb.OffsetPageParams) TokensQ

	Insert(token Token) (id int64, err error)
	Transaction(fn func(q TokensQ) error) error

	UpdateStatus(newStatus resources.TokenStatus) TokensQ
	UpdateTokenId(newTokenId int64) TokensQ
	UpdateOwner(newOwner string) TokensQ
	UpdateMetadataHash(newMetadataHash string) TokensQ
	Update(id int64) error
}
