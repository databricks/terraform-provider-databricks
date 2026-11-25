package u2m

import "errors"

// ErrMissingRefreshToken is returned when a token refresh is requested but the
// cached OAuth token does not include a refresh token.
var ErrMissingRefreshToken = errors.New("cached token has no refresh token")

// InvalidRefreshTokenError is returned from PersistentAuth's Token() and
// ForceRefreshToken() methods when a token refresh is attempted and the cached
// refresh token is invalid.
type InvalidRefreshTokenError struct {
	error
}
