package handlers

import (
	"fmt"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

var (
	InvalidUsagesError = errors.New("usages should be lower or equal initial usages")
)

func Inactive() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusForbidden),
		Status: fmt.Sprintf("%d", http.StatusForbidden),
		Detail: "promocode is inactive",
	}
}
