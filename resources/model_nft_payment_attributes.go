/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftPaymentAttributes struct {
	// Url of book
	BookUrl string `json:"book_url"`
	// Floor nft price
	FloorPrice string `json:"floor_price"`
	// Price of a minted token in $
	MintedTokenPrice string `json:"minted_token_price"`
	// Address of a nft collection
	NftAddress string `json:"nft_address"`
	// Nft id
	NftId int32 `json:"nft_id"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Timestamp when the user have purchased a book
	PurchaseTimestamp string `json:"purchase_timestamp"`
}
