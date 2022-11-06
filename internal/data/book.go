package data

import (
	"encoding/json"
	"gitlab.com/tokend/nft-books/generator-svc/resources/book_resources"
	"time"
)

type BookQ interface {
	New() BookQ
	Get() (*Book, error)
	FilterActual() BookQ
	FilterByID(id int64) BookQ
}

type Book struct {
	ID              int64     `db:"id" structs:"-"`
	Title           string    `db:"title" structs:"title"`
	Description     string    `db:"description" structs:"description"`
	CreatedAt       time.Time `db:"created_at" structs:"created_at"`
	Price           string    `db:"price" structs:"price"`
	ContractAddress string    `db:"contract_address" structs:"contract_address"`
	ContractName    string    `db:"contract_name" structs:"contract_name"`
	ContractSymbol  string    `db:"contract_symbol" structs:"contract_symbol"`
	ContractVersion string    `db:"contract_version" structs:"contract_version"`
	Banner          string    `db:"banner" structs:"banner"`
	File            string    `db:"file" structs:"file"`
	Deleted         bool      `db:"deleted" structs:"-"`
	LastBlock       uint64    `db:"last_block" structs:"last_block"`
}

func (b *Book) Resource() (*book_resources.Book, error) {
	media, err := UnmarshalMedia(b.Banner, b.File)
	if err != nil {
		return nil, err
	}

	media[0].Key = book_resources.NewKeyInt64(b.ID, book_resources.BANNERS)
	media[1].Key = book_resources.NewKeyInt64(b.ID, book_resources.FILES)

	res := book_resources.Book{
		Key: book_resources.NewKeyInt64(b.ID, book_resources.BOOKS),
		Attributes: book_resources.BookAttributes{
			Title:           b.Title,
			Description:     b.Description,
			CreatedAt:       b.CreatedAt,
			Price:           b.Price,
			ContractAddress: b.ContractAddress,
			ContractName:    b.ContractName,
			ContractSymbol:  b.ContractSymbol,
			ContractVersion: b.ContractVersion,
			File:            media[0],
			Banner:          media[1],
		},
	}

	return &res, nil

}

func UnmarshalMedia(media ...string) ([]book_resources.Media, error) {
	var res []book_resources.Media
	var unmarshalledMedia *book_resources.Media

	for _, value := range media {
		err := json.Unmarshal([]byte(value), &unmarshalledMedia)
		if err != nil {
			return nil, err
		}

		res = append(res, *unmarshalledMedia)
		unmarshalledMedia = nil
	}
	return res, nil
}
