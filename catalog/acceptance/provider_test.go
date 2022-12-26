package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccCreateProviderDb2Open(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_provider" "this" {
			  name = "terraform-test-provider"
			  comment = "made by terraform"
			  authentication_type = "TOKEN"
			  recipient_profile_str = jsonencode({
					"shareCredentialsVersion":1,
					"bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
					"endpoint":"https://sharing.delta.io/delta-sharing/"}
				}
			  )
			}`,
		},
	})
}
