/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateNftRequestAttributes struct {
	// Id of network chain
	ChainId int64 `json:"chain_id"`
	// Address of a nft collection
	CollectionAddress string `json:"collection_address"`
	// Id of NFT
	NftId int64 `json:"nft_id"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
}
