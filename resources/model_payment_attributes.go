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
	// Price of a minted token in $
	MintedTokenPrice string `json:"minted_token_price"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Price of a payment token (which the minted token was bought by) in $
	PaymentTokenPrice string `json:"payment_token_price"`
	// Timestamp when the user have purchased a book
	PurchaseTimestamp string `json:"purchase_timestamp"`
}
