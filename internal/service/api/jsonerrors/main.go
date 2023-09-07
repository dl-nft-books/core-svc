package jsonerrors

import (
	"github.com/google/jsonapi"
)

var errorToDetails = map[string]string{
	ApiMaxTriesExceeded: "Max number of tries exceeded",
	TaskDelayNotPassed:  "Cannot create new task if delay period has not been passed after the previous creation",
	NotManagerAuthToken: "Only manager can perform the operation",
}

const (
	ApiMaxTriesExceeded = "api_max_tries_exceeded"
	TaskDelayNotPassed  = "task_delay_not_passed"
	NotManagerAuthToken = "not_manager_auth_token"
)

func WithDetails(err *jsonapi.ErrorObject, code string) *jsonapi.ErrorObject {
	err.Code = code
	err.Detail = errorToDetails[code]
	err.Meta = &map[string]interface{}{
		"error_code": code,
	}

	return err
}
