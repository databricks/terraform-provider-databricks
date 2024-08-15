package acceptance

import (
	"testing"
)

func TestUcAccCreateProviderDb2Open(t *testing.T) {
	unityWorkspaceLevel(t, LegacyStep{
		Template: `
		resource "databricks_provider" "this" {
			name = "terraform-test-provider"
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
