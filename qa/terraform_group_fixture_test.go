package qa_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const initialGroupTemplate = `
resource "databricks_group" "this" {
  display_name = "Initial Group"
}
`

const updatedGroupTemplate = `
resource "databricks_group" "this" {
  display_name = "Updated Group"
}
`

func TestTerraformGroupLifecycle(t *testing.T) {
	terraformtest.Run(t, terraformtest.TestCase{
		FixtureFile:  "testdata/TestTerraformGroupLifecycle.jsonl",
		HostMetadata: terraformtest.WorkspaceHostMetadata("12345"),
		Steps: []terraformtest.Step{
			{
				Name:     "create_group",
				Template: initialGroupTemplate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_group.this", "id", "group-id"),
					resource.TestCheckResourceAttr("databricks_group.this", "display_name", "Initial Group"),
				),
			},
			{
				Name:     "update_group",
				Template: updatedGroupTemplate,
				Check: resource.TestCheckResourceAttr(
					"databricks_group.this", "display_name", "Updated Group",
				),
			},
		},
	})
}
