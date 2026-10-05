package qa_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/tokens"
)

func TestTerraformTokenLifecycleWithHTTPFixtures(t *testing.T) {
	acceptance.HTTPFixtures(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/.well-known/databricks-config",
			Response:     map[string]string{"host_type": "WORKSPACE_HOST", "workspace_id": "12345"},
			ReuseRequest: true,
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedHeaders: map[string]string{
				"Authorization": "Bearer ...",
			},
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
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: tokens.TokenList{TokenInfos: []tokens.TokenInfo{
				{TokenID: "abc", Comment: "Testing token"},
			}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: tokens.TokenList{TokenInfos: []tokens.TokenInfo{
				{TokenID: "abc", Comment: "Testing token"},
			}},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.0/token/delete",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			ExpectedRequest: map[string]interface{}{"token_id": "abc"},
		},
	}, acceptance.Step{
		Template: `resource "databricks_token" "this" {
			lifetime_seconds = 6000
			comment = "Testing token"
		}`,
	})
}
