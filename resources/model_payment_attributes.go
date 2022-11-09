/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PaymentAttributes struct {
	// Amount of tokens paid
	Amount  string `json:"amount"`
	BookUrl string `json:"book_url"`
	// Token metadata information
	Erc20Data Erc20Data `json:"erc20_data"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Price of a token in $
	Price string `json:"price"`
	// Timestamp when the user have purchased a book
	PurchaseTimestamp string `json:"purchase_timestamp"`
}
