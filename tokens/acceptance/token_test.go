package acceptance

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/tokens"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"

	"testing"
)

func TestAccTokenResource(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_token" "this" {
				lifetime_seconds = 6000
				comment = "Testing token"
			}`,
			ExpectNonEmptyPlan: true,
		},
		{
			ExpectNonEmptyPlan: true,
			Callback: func(ctx context.Context,
				client *common.DatabricksClient, id string) error {
				return tokens.NewTokensAPI(context.Background(), client).Delete(id)
			},
		},
	})
}
