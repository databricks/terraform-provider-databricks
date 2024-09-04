package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const groupsDataSourceTemplate = `
resource "databricks_group" "group1" {
	display_name = "tf-group-{var.RANDOM}"
}
resource "databricks_group" "group2" {
	display_name = "tf-groupfoo-{var.RANDOM}"
}
resource "databricks_group" "group3" {
	display_name = "tf-groupbar-{var.RANDOM}"
}

#data "databricks_groups" "all" {}

data "databricks_groups" "this" {
	filter = "displayName co \"foo\" or displayName co \"bar\""
}
`

func checkGroupsDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r_all, ok := s.Modules[0].Resources["data.databricks_groups.all"]
		require.True(t, ok, "data.databricks_group.all has to be there")
		attr := r_all.Primary.Attributes
		assert.Equal(t, attr["display_names.#"], "3")

		r_filtered, ok := s.Modules[0].Resources["data.databricks_groups.this"]
		require.True(t, ok, "data.databricks_group.this has to be there")
		attr_filtered := r_filtered.Primary.Attributes
		assert.Equal(t, attr_filtered["display_names.#"], "2")

		return nil
	}
}

func TestAccDataSourceGroups(t *testing.T) {
	accountLevel(t, step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}
func TestMwsAccDataSourceGroups(t *testing.T) {
	workspaceLevel(t, step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}
