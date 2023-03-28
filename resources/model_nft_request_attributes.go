/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type NftRequestAttributes struct {
	// Id of network chain
	ChainId int64 `json:"chain_id"`
	// Address of a nft collection
	CollectionAddress string `json:"collection_address"`
	// Timestamp when the user have purchased a book
	CreatedAt time.Time `json:"created_at"`
	// Id of NFT
	NftId int64 `json:"nft_id"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// nft request status
	Status NftRequestStatus `json:"status"`
}
