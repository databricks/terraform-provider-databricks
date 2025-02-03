package sharing_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccCreateProviderDb2Open(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_provider" "this" {
			name = "terraform-test-provider-{var.RANDOM}"
			comment = "made by terraform"
			authentication_type = "TOKEN"
			recipient_profile_str = jsonencode({
				"shareCredentialsVersion":1,
				"bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
				"endpoint":"https://sharing.delta.io/delta-sharing/"
			}
			)
		}`,
	})
}
