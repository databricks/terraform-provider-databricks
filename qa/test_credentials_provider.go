package qa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config"
)

type testCredentialsProvider struct {
	token string
}

func (testCredentialsProvider) Name() string {
	return "test"
}

func (t testCredentialsProvider) Configure(ctx context.Context, cfg *config.Config) (func(*http.Request) error, error) {
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
		return nil
	}, nil
}
