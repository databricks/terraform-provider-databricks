// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import "github.com/databricks/databricks-sdk-go/apierr"

var (
	// the request is invalid
	ErrBadRequest = apierr.ErrBadRequest
	// the request does not have valid authentication (AuthN) credentials for the operation
	ErrUnauthenticated = apierr.ErrUnauthenticated
	// the caller does not have permission to execute the specified operation
	ErrPermissionDenied = apierr.ErrPermissionDenied
	// the operation was performed on a resource that does not exist
	ErrNotFound = apierr.ErrNotFound
	// maps to all HTTP 409 (Conflict) responses
	ErrResourceConflict = apierr.ErrResourceConflict
	// maps to HTTP code: 429 Too Many Requests
	ErrTooManyRequests = apierr.ErrTooManyRequests
	// the operation was explicitly canceled by the caller
	ErrCancelled = apierr.ErrCancelled
	// some invariants expected by the underlying system have been broken
	ErrInternalError = apierr.ErrInternalError
	// the operation is not implemented or is not supported/enabled in this service
	ErrNotImplemented = apierr.ErrNotImplemented
	// the service is currently unavailable
	ErrTemporarilyUnavailable = apierr.ErrTemporarilyUnavailable
	// the deadline expired before the operation could complete
	ErrDeadlineExceeded = apierr.ErrDeadlineExceeded
	// unexpected state
	ErrInvalidState = apierr.ErrInvalidState
	// supplied value for a parameter was invalid
	ErrInvalidParameterValue = apierr.ErrInvalidParameterValue
	// operation was performed on a resource that does not exist
	ErrResourceDoesNotExist = apierr.ErrResourceDoesNotExist
	// the operation was aborted, typically due to a concurrency issue such as a sequencer check failure
	ErrAborted = apierr.ErrAborted
	// operation was rejected due a conflict with an existing resource
	ErrAlreadyExists = apierr.ErrAlreadyExists
	// operation was rejected due a conflict with an existing resource
	ErrResourceAlreadyExists = apierr.ErrResourceAlreadyExists
	// operation is rejected due to per-user rate limiting
	ErrResourceExhausted = apierr.ErrResourceExhausted
	// cluster request was rejected because it would exceed a resource limit
	ErrRequestLimitExceeded = apierr.ErrRequestLimitExceeded
	// this error is used as a fallback if the platform-side mapping is missing some reason
	ErrUnknown = apierr.ErrUnknown
	// unrecoverable data loss or corruption
	ErrDataLoss = apierr.ErrDataLoss
)
