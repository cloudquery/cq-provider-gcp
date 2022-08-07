package client

import (
	"errors"
	"net/http"

	"google.golang.org/api/googleapi"
)

func IgnoreErrorHandler(err error) bool {
	var gerr *googleapi.Error
	if ok := errors.As(err, &gerr); ok {
		if gerr.Code == http.StatusForbidden && len(gerr.Errors) > 0 {
			switch gerr.Errors[0].Reason {
			case "accessNotConfigured", "forbidden", "SERVICE_DISABLED":
				return true
			}
		}
		if gerr.Code == http.StatusNotFound && len(gerr.Errors) > 0 {
			return true
		}
	}
	return false
}
