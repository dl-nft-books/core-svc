package data

type PromocodesBooksQ interface {
	New() PromocodesBooksQ

	Insert(promocodeId int64, bookIds ...int64) error
	DeleteBooks(promocodeId int64, bookIds ...int64) error

	UpdatePromocodes(bookId int64, promocodeIds ...int64) error
	UpdateBooks(promocodeId int64, bookIds ...int64) error

	SelectPromocodes() (promocodes []int64, err error)
	GetPromocodes() (*int64, error)
	SelectBooks() (promocodes []int64, err error)
	GetBooks() (*int64, error)
}
