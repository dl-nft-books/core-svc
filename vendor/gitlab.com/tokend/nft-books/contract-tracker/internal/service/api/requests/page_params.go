package requests

import (
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

const (
	pageParamLimit  = "page[limit]"
	pageParamNumber = "page[number]"
	pageParamSort   = "sort"
)

func GetOffsetLinksWithSort(r *http.Request, params pgdb.OffsetPageParams, sorts pgdb.Sorts) *resources.Links {
	links := resources.Links{
		First: getOffsetLinkWithSort(r, 0, params.Limit, sorts),
		Next:  getOffsetLinkWithSort(r, params.PageNumber+1, params.Limit, sorts),
		Self:  getOffsetLinkWithSort(r, params.PageNumber, params.Limit, sorts),
	}

	if params.PageNumber != 0 {
		links.Prev = getOffsetLinkWithSort(r, params.PageNumber-1, params.Limit, sorts)
	}

	return &links
}

func getOffsetLinkWithSort(r *http.Request, pageNumber, limit uint64, sorts pgdb.Sorts) string {
	var (
		url   = r.URL
		query = url.Query()
	)

	query.Set(pageParamNumber, strconv.FormatUint(pageNumber, 10))
	query.Set(pageParamLimit, strconv.FormatUint(limit, 10))

	if _, ok := query[pageParamSort]; !ok {
		for _, sort := range sorts {
			query.Add(pageParamSort, string(sort))
		}
	}

	url.RawQuery = query.Encode()
	return url.String()
}
