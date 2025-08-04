package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const groupsDataSourceTemplate = `
resource "databricks_user" "test_user" {
	user_name = "tf-groups-user-{var.RANDOM}@example.com"
}

resource "databricks_service_principal" "test_sp" {
	application_id = "{var.RANDOM_UUID}"
	display_name = "tf-groups-sp-{var.RANDOM}"
	force = true
}

resource "databricks_group" "test_group_1" {
	display_name = "tf-groups-test-{var.RANDOM}-alpha"
}

resource "databricks_group" "test_group_2" {
	display_name = "tf-groups-test-{var.RANDOM}-beta"
}

resource "databricks_group" "other_group" {
	display_name = "tf-other-{var.RANDOM}"
}

resource "databricks_group_member" "member_1_user" {
	group_id  = databricks_group.test_group_1.id
	member_id = databricks_user.test_user.id
}

resource "databricks_group_member" "member_1_sp" {
	group_id  = databricks_group.test_group_1.id
	member_id = databricks_service_principal.test_sp.id
}

resource "databricks_group_member" "member_2_group" {
	group_id  = databricks_group.test_group_2.id
	member_id = databricks_group.test_group_1.id
}

data "databricks_groups" "test_filter" {
	filter = "displayName co \"tf-groups-test-{var.RANDOM}\""
	depends_on = [
		databricks_group_member.member_1_user,
		databricks_group_member.member_1_sp,
		databricks_group_member.member_2_group
	]
}

data "databricks_groups" "exact_match" {
	filter = "displayName eq \"${databricks_group.test_group_1.display_name}\""
	depends_on = [
		databricks_group_member.member_1_user,
		databricks_group_member.member_1_sp
	]
}

data "databricks_groups" "empty_filter" {
	filter = "displayName co \"nonexistent-group-{var.RANDOM}\""
}`

func checkGroupsDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		// Check the filtered data source that should return 2 groups
		r, ok := s.Modules[0].Resources["data.databricks_groups.test_filter"]
		require.True(t, ok, "data.databricks_groups.test_filter has to be there")
		attr := r.Primary.Attributes

		// Should find 2 groups matching the filter
		assert.Equal(t, "2", attr["groups.#"])

		// Groups should be sorted alphabetically (alpha before beta)
		group1DisplayName := attr["groups.0.display_name"]
		group2DisplayName := attr["groups.1.display_name"]

		require.Contains(t, group1DisplayName, "alpha", "First group should be alpha (alphabetically first)")
		require.Contains(t, group2DisplayName, "beta", "Second group should be beta (alphabetically second)")

		// Verify group 1 (alpha) has the expected members
		assert.NotEmpty(t, attr["groups.0.users.0"], "Group 1 should have users")
		assert.NotEmpty(t, attr["groups.0.service_principals.0"], "Group 1 should have service principals")
		assert.Contains(t, attr["groups.0.acl_principal_id"], "groups/", "ACL principal ID should have groups/ prefix")

		// Verify group 2 (beta) has the expected child group
		assert.NotEmpty(t, attr["groups.1.child_groups.0"], "Group 2 should have child groups")
		assert.Contains(t, attr["groups.1.acl_principal_id"], "groups/", "ACL principal ID should have groups/ prefix")

		// Check the exact match data source that should return 1 group
		exactMatch, ok := s.Modules[0].Resources["data.databricks_groups.exact_match"]
		require.True(t, ok, "data.databricks_groups.exact_match has to be there")
		exactAttr := exactMatch.Primary.Attributes

		assert.Equal(t, "1", exactAttr["groups.#"])
		assert.Contains(t, exactAttr["groups.0.display_name"], "alpha", "Exact match should return the alpha group")

		// Check empty result data source
		emptyResult, ok := s.Modules[0].Resources["data.databricks_groups.empty_filter"]
		require.True(t, ok, "data.databricks_groups.empty_filter has to be there")
		emptyAttr := emptyResult.Primary.Attributes

		assert.Equal(t, "0", emptyAttr["groups.#"], "Filter matching no groups should return empty list")

		return nil
	}
}

func TestMwsAccGroupsDataSource(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	acceptance.AccountLevel(t, acceptance.Step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}

func TestAccGroupsDataSource(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}

func TestAccGroupsDataSourceOnAWS(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}

func TestAccGroupsDataSourceOnGCP(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupsDataSourceTemplate,
		Check:    checkGroupsDataSourcePopulated(t),
	})
}
