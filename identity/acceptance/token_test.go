package acceptance

import (
	"context"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"

	"testing"
)

func TestAccTokenResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_token" "this" {
				lifetime_seconds = 6000
				comment = "Testing token"
			}`,
		},
		{
			Callback: func(ctx context.Context, client *common.DatabricksClient, id string) error {
				return NewTokensAPI(context.Background(), client).Delete(id)
			},
		},
	})
}
