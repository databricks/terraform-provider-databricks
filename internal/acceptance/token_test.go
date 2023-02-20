package acceptance

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/tokens"

	"testing"
)

func TestAccTokenResource(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_token" "this" {
			lifetime_seconds = 6000
			comment = "Testing token"
		}`,
		ExpectNonEmptyPlan: true,
	}, step{
		ExpectNonEmptyPlan: true,
		Callback: func(ctx context.Context,
			client *common.DatabricksClient, id string) error {
			return tokens.NewTokensAPI(context.Background(), client).Delete(id)
		},
	})
}
