package credentials

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
)

var errTokenRequired = errors.New("token is required")

// NewPATCredentials returns a Credentials that can be used to authenticate with
// a Personal Access Token.
func NewPATCredentials(token string) (auth.Credentials, error) {
	if token == "" {
		return nil, errTokenRequired
	}
	return auth.CredentialsFn(func(_ context.Context) (map[string]string, error) {
		return map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token),
		}, nil
	}), nil
}
