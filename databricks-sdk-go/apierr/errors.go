package apierr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
)

// Deprecated: Use [ErrorDetails] instead.
type ErrorDetail struct {
	Type     string            `json:"@type,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Domain   string            `json:"domain,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// APIError represents a standard Databricks API error.
type APIError struct {
	ErrorCode  string
	Message    string
	StatusCode int

	// ResponseWrapper contains request/response for the error.
	//
	// EXPERIMENTAL INTERNAL: This field is experimental and meant for internal
	// use only. It will change or be removed in the future.
	ResponseWrapper *common.ResponseWrapper

	errorDetails ErrorDetails

	// If non-nil, the underlying error that should be returned by calling
	// errors.Unwrap on this error.
	unwrap error

	// Details is the sublist of error details that can be unmarshalled into
	// the [ErrorDetail] type.
	//
	// Deprecated: Use [APIError.ErrorDetails] instead.
	Details []ErrorDetail
}

// ErrorDetails returns the error details of the APIError.
func (apiErr *APIError) ErrorDetails() ErrorDetails {
	return apiErr.errorDetails
}

// Error returns the error message string.
func (apiError *APIError) Error() string {
	return apiError.Message
}

// IsMissing tells if error is about missing resource
func IsMissing(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsMissing tells if it is missing resource.
func (apiError *APIError) IsMissing() bool {
	return errors.Is(apiError, ErrNotFound)
}

// IsTooManyRequests shows rate exceeded limits.
func (apiError *APIError) IsTooManyRequests() bool {
	return errors.Is(apiError, ErrTooManyRequests)
}

// GetErrorInfo returns all entries in the list of error details of type
// `ErrorInfo`.
//
// Deprecated: Use [APIError.ErrorDetails] instead.
func GetErrorInfo(err error) []ErrorDetail {
	var apiError *APIError
	if !errors.As(err, &apiError) {
		return nil
	}

	filteredDetails := []ErrorDetail{}
	for _, detail := range apiError.Details {
		if errorInfoType == detail.Type {
			filteredDetails = append(filteredDetails, detail)
		}
	}
	return filteredDetails
}

// IsRetriable returns true if error is retriable.
func (apiError *APIError) IsRetriable(ctx context.Context) bool {
	if apiError.IsTooManyRequests() {
		return true
	}
	// Retry when the API is unavailable. This includes "No webapps are available
	// to handle your request. Please try again later." from the API gateway.
	// Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
	if apiError.StatusCode == http.StatusServiceUnavailable {
		return true
	}
	// Handle transient errors for retries
	for _, regex := range allTransientErrors {
		if regex.Match([]byte(apiError.Message)) {
			logger.Debugf(ctx, "Attempting retry because of %s", regex.String())
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// GetAPIError returns the API error from the response. If the response is not
// an error, it returns nil.
func GetAPIError(ctx context.Context, resp common.ResponseWrapper) error {
	// Responses in the 2xx and 3xx ranges are not standard Databricks errors.
	if resp.Response.StatusCode >= 400 {
		apiError := parseErrorFromResponse(ctx, resp)
		applyOverrides(ctx, apiError, resp.Response)
		apiError.ResponseWrapper = &resp
		return apiError
	}

	// Return an error if the response indicates that the request tried to
	// access a private link workspace without proper access.
	requestUrl := resp.Response.Request.URL
	if isPrivateLinkRedirect(requestUrl) {
		apiError := privateLinkValidationError(requestUrl)
		apiError.ResponseWrapper = &resp
		return apiError
	}

	return nil // not an error
}

func parseErrorFromResponse(ctx context.Context, resp common.ResponseWrapper) *APIError {
	errorBody, err := io.ReadAll(resp.ReadCloser)
	if err != nil {
		return &APIError{
			ErrorCode:  "IO_READ",
			StatusCode: resp.Response.StatusCode,
			Message:    err.Error(),
		}
	}

	if len(errorBody) == 0 {
		return &APIError{
			Message:    http.StatusText(resp.Response.StatusCode),
			StatusCode: resp.Response.StatusCode,
		}
	}

	if err := standardErrorParser(ctx, resp.Response, errorBody); err != nil {
		return err
	}
	if err := stringErrorParser(ctx, resp.Response, errorBody); err != nil {
		return err
	}
	if err := htmlErrorParser(ctx, resp.Response, errorBody); err != nil {
		return err
	}

	// Unknown error response typically come from API gateways, load balancers,
	// and other middlewares. These responses are not expected to be standard
	// Databricks API errors.
	return &APIError{
		ErrorCode:  "UNKNOWN",
		StatusCode: resp.Response.StatusCode,
		Message:    string(errorBody),
	}
}

// standardErrorParser is the default error parser for Databricks API errors.
// It handles JSON error messages with error code, message, and details fields.
// It also provides compatibility with the old API 1.2 error format and SCIM API
// errors.
func standardErrorParser(ctx context.Context, resp *http.Response, responseBody []byte) *APIError {
	// Anonymous struct used to unmarshal JSON Databricks API error responses.
	var errorBody struct {
		ErrorCode  any               `json:"error_code,omitempty"` // int or string
		Message    string            `json:"message,omitempty"`
		API12Error string            `json:"error,omitempty"`
		RawDetails []json.RawMessage `json:"details,omitempty"`

		// The following fields are for scim api only. See RFC7644 section 3.7.3
		// https://tools.ietf.org/html/rfc7644#section-3.7.3
		ScimDetail string `json:"detail,omitempty"`
		ScimStatus string `json:"status,omitempty"`
		ScimType   string `json:"scimType,omitempty"`
	}
	if err := json.Unmarshal(responseBody, &errorBody); err != nil {
		logger.Tracef(ctx, "standardErrorParser: failed to unmarshal error response: %s", err)
		return nil
	}

	// Convert API 1.2 error (which used a different format) to the new format.
	if errorBody.API12Error != "" {
		errorBody.Message = errorBody.API12Error
	}

	// Handle SCIM error message details.
	if errorBody.Message == "" && errorBody.ScimDetail != "" {
		if errorBody.ScimDetail == "null" {
			errorBody.Message = "SCIM API Internal Error"
		} else {
			errorBody.Message = errorBody.ScimDetail
		}
		// Add more context from SCIM responses.
		errorBody.Message = fmt.Sprintf("%s %s", errorBody.ScimType, errorBody.Message)
		errorBody.Message = strings.Trim(errorBody.Message, " ")
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
	}

	apierr := &APIError{
		Message:    errorBody.Message,
		ErrorCode:  fmt.Sprintf("%v", errorBody.ErrorCode),
		StatusCode: resp.StatusCode,
	}

	// Parse the error details, dropping any that fail to unmarshal.
	details := []any{}
	for _, rd := range errorBody.RawDetails {
		details = append(details, unmarshalDetails(rd))

		// Deprecated: unmarshal ErrorDetail type for backwards compatibility
		// with the previous behavior.
		ed := ErrorDetail{}
		if json.Unmarshal(rd, &ed) == nil { // ignore errors
			apierr.Details = append(apierr.Details, ed)
		}
	}
	apierr.errorDetails = parseErrorDetails(details)

	return apierr
}

var stringErrorRegex = regexp.MustCompile(`^([A-Z_]+): (.*)$`)

// stringErrorParser parses errors of the form "STATUS_CODE: status message".
// Some account-level APIs respond with this error code, e.g.
// https://github.com/databricks/databricks-sdk-go/issues/998
func stringErrorParser(ctx context.Context, resp *http.Response, responseBody []byte) *APIError {
	matches := stringErrorRegex.FindSubmatch(responseBody)
	if len(matches) < 3 {
		logger.Tracef(ctx, "stringErrorParser: failed to match error response")
		return nil
	}
	return &APIError{
		Message:    string(matches[2]),
		ErrorCode:  string(matches[1]),
		StatusCode: resp.StatusCode,
	}
}

var htmlMessageRe = regexp.MustCompile(`<pre>(.*)</pre>`)

// htmlErrorParser parses HTML error responses. Some legacy APIs respond with
// an HTML page in certain error cases, like when trying to create a cluster
// before the worker environment is ready.
func htmlErrorParser(ctx context.Context, resp *http.Response, responseBody []byte) *APIError {
	messageMatches := htmlMessageRe.FindStringSubmatch(string(responseBody))
	// No messages with <pre> </pre> format found so return a APIError
	if len(messageMatches) < 2 {
		logger.Tracef(ctx, "htmlErrorParser: no <pre> tag found in error response")
		return nil
	}

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Message:    strings.Trim(messageMatches[1], " ."),
	}

	// this is most likely HTML... since un-marshalling JSON failed
	// Status parts first in case html message is not as expected
	statusParts := strings.SplitN(resp.Status, " ", 2)
	if len(statusParts) < 2 {
		apiErr.ErrorCode = "UNKNOWN"
	} else {
		apiErr.ErrorCode = strings.ReplaceAll(strings.ToUpper(strings.Trim(statusParts[1], " .")), " ", "_")
	}

	return apiErr
}
