package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const groupDataSourceTemplate = `
resource "databricks_user" "this" {
	user_name = "tf-{var.RANDOM}@example.com"
}
resource "databricks_service_principal" "this" {
	application_id = "{var.RANDOM_UUID}"
	display_name = "tf-spn-{var.RANDOM}"
	force = true
}
resource "databricks_group" "child" {
	display_name = "tf-child-{var.RANDOM}"
}
resource "databricks_group" "parent" {
	display_name = "tf-parent-{var.RANDOM}"
}
resource "databricks_group_member" "m01" {
	group_id  = databricks_group.parent.id
	member_id = databricks_group.child.id
}
resource "databricks_group_member" "m02" {
	group_id  = databricks_group.parent.id
	member_id = databricks_service_principal.this.id
}
resource "databricks_group_member" "m03" {
	group_id  = databricks_group.parent.id
	member_id = databricks_user.this.id
}
data "databricks_group" "this" {
	display_name = databricks_group.parent.display_name
	depends_on = [
		databricks_group_member.m01,
		databricks_group_member.m02,
		databricks_group_member.m03
	]
}`

func checkGroupDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_group.this"]
		require.True(t, ok, "data.databricks_group.this has to be there")
		attr := r.Primary.Attributes

		assert.Equal(t, s.Modules[0].Resources["databricks_user.this"].Primary.ID, attr["users.0"])
		assert.Equal(t, s.Modules[0].Resources["databricks_service_principal.this"].Primary.ID, attr["service_principals.0"])
		assert.Equal(t, s.Modules[0].Resources["databricks_group.child"].Primary.ID, attr["child_groups.0"])
		return nil
	}
}

func TestMwsAccGroupDataSplitMembers(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	accountLevel(t, LegacyStep{
		Template: groupDataSourceTemplate,
		Check:    checkGroupDataSourcePopulated(t),
	})
}

func TestAccGroupDataSplitMembers(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, LegacyStep{
		Template: groupDataSourceTemplate,
		Check:    checkGroupDataSourcePopulated(t),
	})
}
