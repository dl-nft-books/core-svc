/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TaskAttributes struct {
	BookId           int64      `json:"book_id"`
	FileIpfsHash     string     `json:"file_ipfs_hash"`
	MetadataIpfsHash string     `json:"metadata_ipfs_hash"`
	Signature        string     `json:"signature"`
	Status           TaskStatus `json:"status"`
	TokenId          int64      `json:"token_id"`
	Uri              string     `json:"uri"`
}
