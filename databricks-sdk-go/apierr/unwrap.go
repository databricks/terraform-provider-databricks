package apierr

func inheritErr(err error, message string) *wrapError {
	return &wrapError{
		wrap:    err,
		message: message,
	}
}

type wrapError struct {
	wrap    error
	message string
}

func (e *wrapError) Error() string {
	return e.message
}

func (e *wrapError) Unwrap() error {
	return e.wrap
}

func ByStatusCode(statusCode int) (error, bool) {
	err, ok := statusCodeMapping[statusCode]
	return err, ok
}

// Unwrap error for easier client code checking
//
// See https://pkg.go.dev/errors#example-Unwrap
func (apiError *APIError) Unwrap() error {
	if apiError.unwrap != nil {
		return apiError.unwrap
	}
	byErrorCode, ok := errorCodeMapping[apiError.ErrorCode]
	if ok {
		return byErrorCode
	}
	byStatusCode, ok := ByStatusCode(apiError.StatusCode)
	if ok {
		return byStatusCode
	}
	// A nil error returned from e.Unwrap() indicates that e does not wrap
	// any error.
	return nil
}
