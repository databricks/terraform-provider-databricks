package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/tokens"
)

func TestRunningRealTerraformWithFixtureBackend(t *testing.T) {
	t.Skip("fails on GitHub Actions")
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedRequest: tokens.TokenRequest{
				LifetimeSeconds: 6000,
				Comment:         "Testing token",
			},
			Response: tokens.TokenResponse{
				TokenValue: "xyz",
				TokenInfo: &tokens.TokenInfo{
					TokenID: "abc",
					Comment: "Testing token",
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/token/list",
			ReuseRequest: true,
			Response: tokens.TokenList{
				TokenInfos: []tokens.TokenInfo{
					{
						TokenID: "abc",
						Comment: "Testing token",
					},
				},
			},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.0/token/delete",
			ExpectedRequest: map[string]interface{}{"token_id": "abc"},
		},
	},
		func(ctx context.Context, client *common.DatabricksClient) {
			t.Setenv("DATABRICKS_HOST", client.Config.Host)
			t.Setenv("DATABRICKS_TOKEN", client.Config.Token)

			workspaceLevel(t, LegacyStep{
				Template: `resource "databricks_token" "this" {
					lifetime_seconds = 6000
					comment = "Testing token"
				}`,
			})
		})
}
