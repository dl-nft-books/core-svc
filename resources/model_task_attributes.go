/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TaskAttributes struct {
	// Id of book
	BookId int64 `json:"book_id"`
	// hash of file on IPFS
	FileIpfsHash string `json:"file_ipfs_hash"`
	// hash of metadata on IPFS
	MetadataIpfsHash string `json:"metadata_ipfs_hash"`
	// personal string on first page
	Signature string `json:"signature"`
	// task solution status
	Status TaskStatus `json:"status"`
	// Id of token
	TokenId int64  `json:"token_id"`
	Uri     string `json:"uri"`
}
