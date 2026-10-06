package qa_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const initialUserTemplate = `
resource "databricks_user" "this" {
  user_name    = "fixture@example.com"
  display_name = "Initial Name"
}
`

const updatedUserTemplate = `
resource "databricks_user" "this" {
  user_name    = "fixture@example.com"
  display_name = "Updated Name"
}
`

func TestTerraformUserLifecycle(t *testing.T) {
	terraformtest.Run(t, terraformtest.TestCase{
		FixtureFile:  "testdata/TestTerraformUserLifecycle.jsonl",
		HostMetadata: terraformtest.WorkspaceHostMetadata("12345"),
		Steps: []terraformtest.Step{
			{
				Name:     "create_user",
				Template: initialUserTemplate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_user.this", "id", "user-id"),
					resource.TestCheckResourceAttr("databricks_user.this", "display_name", "Initial Name"),
				),
			},
			{
				Name:     "update_user",
				Template: updatedUserTemplate,
				Check: resource.TestCheckResourceAttr(
					"databricks_user.this", "display_name", "Updated Name",
				),
			},
		},
	})
}
