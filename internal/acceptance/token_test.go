package acceptance

import (
	"testing"
)

func TestAccTokenResource(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `resource "databricks_token" "this" {
			lifetime_seconds = 6000
			comment = "Testing token"
		}`,
	})
}
