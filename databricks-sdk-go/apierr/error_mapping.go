// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apierr

import (
	"errors"
)

var (
	ErrBadRequest             = errors.New("the request is invalid")
	ErrUnauthenticated        = errors.New("the request does not have valid authentication (AuthN) credentials for the operation")
	ErrPermissionDenied       = errors.New("the caller does not have permission to execute the specified operation")
	ErrNotFound               = errors.New("the operation was performed on a resource that does not exist")
	ErrResourceConflict       = errors.New("maps to all HTTP 409 (Conflict) responses")
	ErrTooManyRequests        = errors.New("maps to HTTP code: 429 Too Many Requests")
	ErrCancelled              = errors.New("the operation was explicitly canceled by the caller")
	ErrInternalError          = errors.New("some invariants expected by the underlying system have been broken")
	ErrNotImplemented         = errors.New("the operation is not implemented or is not supported/enabled in this service")
	ErrTemporarilyUnavailable = errors.New("the service is currently unavailable")
	ErrDeadlineExceeded       = errors.New("the deadline expired before the operation could complete")
	ErrInvalidState           = inheritErr(ErrBadRequest, "unexpected state")
	ErrInvalidParameterValue  = inheritErr(ErrBadRequest, "supplied value for a parameter was invalid")
	ErrResourceDoesNotExist   = inheritErr(ErrNotFound, "operation was performed on a resource that does not exist")
	ErrAborted                = inheritErr(ErrResourceConflict, "the operation was aborted, typically due to a concurrency issue such as a sequencer check failure")
	ErrAlreadyExists          = inheritErr(ErrResourceConflict, "operation was rejected due a conflict with an existing resource")
	ErrResourceAlreadyExists  = inheritErr(ErrResourceConflict, "operation was rejected due a conflict with an existing resource")
	ErrResourceExhausted      = inheritErr(ErrTooManyRequests, "operation is rejected due to per-user rate limiting")
	ErrRequestLimitExceeded   = inheritErr(ErrTooManyRequests, "cluster request was rejected because it would exceed a resource limit")
	ErrUnknown                = inheritErr(ErrInternalError, "this error is used as a fallback if the platform-side mapping is missing some reason")
	ErrDataLoss               = inheritErr(ErrInternalError, "unrecoverable data loss or corruption")

	statusCodeMapping = map[int]error{
		400: ErrBadRequest,
		401: ErrUnauthenticated,
		403: ErrPermissionDenied,
		404: ErrNotFound,
		409: ErrResourceConflict,
		429: ErrTooManyRequests,
		499: ErrCancelled,
		500: ErrInternalError,
		501: ErrNotImplemented,
		503: ErrTemporarilyUnavailable,
		504: ErrDeadlineExceeded,
	}

	errorCodeMapping = map[string]error{
		"INVALID_STATE":           ErrInvalidState,
		"INVALID_PARAMETER_VALUE": ErrInvalidParameterValue,
		"RESOURCE_DOES_NOT_EXIST": ErrResourceDoesNotExist,
		"ABORTED":                 ErrAborted,
		"ALREADY_EXISTS":          ErrAlreadyExists,
		"RESOURCE_ALREADY_EXISTS": ErrResourceAlreadyExists,
		"RESOURCE_EXHAUSTED":      ErrResourceExhausted,
		"REQUEST_LIMIT_EXCEEDED":  ErrRequestLimitExceeded,
		"UNKNOWN":                 ErrUnknown,
		"DATA_LOSS":               ErrDataLoss,
	}
)
