/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package book_resources

import "gitlab.com/tokend/nft-books/generator-svc/resources"

type Media struct {
	resources.Key
	Attributes MediaAttributes `json:"attributes"`
}
type MediaResponse struct {
	Data     Media              `json:"data"`
	Included resources.Included `json:"included"`
}

type MediaListResponse struct {
	Data     []Media            `json:"data"`
	Included resources.Included `json:"included"`
	Links    *resources.Links   `json:"links"`
}
