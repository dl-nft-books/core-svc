/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package book_resources

import "gitlab.com/tokend/nft-books/generator-svc/resources"

type Book struct {
	resources.Key
	Attributes BookAttributes `json:"attributes"`
}

type BookResponse struct {
	Data     Book               `json:"data"`
	Included resources.Included `json:"included"`
}

type BookListResponse struct {
	Data     []Book             `json:"data"`
	Included resources.Included `json:"included"`
	Links    *resources.Links   `json:"links"`
}
