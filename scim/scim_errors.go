package scim

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdkapierr "github.com/databricks/sdk-go/core/apierr"
)

// isKnownScimError reports whether err is an API error with one of details as its SCIM
// detail and no scimType, or another error with one of details as its message. The
// detail is read from the response because each SDK formats it into its error message
// differently. A detail with a scimType doesn't match, because the existing client
// puts the scimType before the detail in the message that was compared before.
func isKnownScimError(err error, details []string) bool {
	if body, ok := apiErrorBody(err); ok {
		var scimErr struct {
			Detail   string `json:"detail"`
			ScimType string `json:"scimType"`
		}
		if json.Unmarshal(body, &scimErr) == nil && scimErr.Detail != "" {
			return scimErr.ScimType == "" && slices.Contains(details, scimErr.Detail)
		}
	}
	return slices.Contains(details, err.Error())
}

// apiErrorBody returns the response body of an API error from either SDK.
func apiErrorBody(err error) ([]byte, bool) {
	if apiErr, ok := errors.AsType[*sdkapierr.APIError](err); ok {
		return apiErr.HTTPBody(), true
	}
	if apiErr, ok := errors.AsType[*apierr.APIError](err); ok && apiErr.ResponseWrapper != nil {
		return apiErr.ResponseWrapper.DebugBytes, true
	}
	return nil, false
}
