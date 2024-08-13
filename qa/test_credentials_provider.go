package qa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/credentials"
)

type TestCredentialsProvider struct {
	Token string
}

func (TestCredentialsProvider) Name() string {
	return "test"
}

func (t TestCredentialsProvider) Configure(ctx context.Context, cfg *config.Config) (credentials.CredentialsProvider, error) {
	fun := func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.Token))
		return nil
	}
	return credentials.NewCredentialsProvider(fun), nil
}
