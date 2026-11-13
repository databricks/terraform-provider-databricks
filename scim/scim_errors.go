package scim

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdkapierr "github.com/databricks/sdk-go/core/apierr"
)

// isKnownScimError reports whether the API sent one of messages as its error message.
// Each SDK formats that message into its error differently, so for API errors it is
// read from the response body with the existing client's precedence: error, message,
// then a SCIM detail. A detail with a scimType doesn't match, because the existing
// client puts the scimType before the detail.
func isKnownScimError(err error, messages []string) bool {
	body, ok := apiErrorBody(err)
	if !ok {
		return slices.Contains(messages, err.Error())
	}
	var e struct {
		Error    string `json:"error"`
		Message  string `json:"message"`
		Detail   string `json:"detail"`
		ScimType string `json:"scimType"`
	}
	if json.Unmarshal(body, &e) != nil {
		return false
	}
	switch {
	case e.Error != "":
		return slices.Contains(messages, e.Error)
	case e.Message != "":
		return slices.Contains(messages, e.Message)
	default:
		return e.ScimType == "" && slices.Contains(messages, e.Detail)
	}
}

// apiErrorBody returns the response body of an API error from either SDK.
func apiErrorBody(err error) ([]byte, bool) {
	if apiErr, ok := errors.AsType[*sdkapierr.APIError](err); ok {
		return apiErr.HTTPBody(), true
	}
	apiErr, ok := errors.AsType[*apierr.APIError](err)
	if ok && apiErr.ResponseWrapper != nil {
		return apiErr.ResponseWrapper.DebugBytes, true
	}
	return nil, false
}
