/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TaskAttributes struct {
	Banner *Media `json:"banner,omitempty"`
	// hash of banner on IPFS
	BannerIpfsHash string `json:"banner_ipfs_hash"`
	// Id of book
	BookId int64 `json:"book_id"`
	// Id of network chain
	ChainId *int64 `json:"chain_id,omitempty"`
	// hash of metadata on IPFS
	MetadataIpfsHash string `json:"metadata_ipfs_hash"`
	// task solution status
	Status TaskStatus `json:"status"`
	// Id of token
	TokenId int64  `json:"token_id"`
	Uri     string `json:"uri"`
}
