package user_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const dataSourceTemplate = `
	resource "databricks_user" "user1" {
		user_name = "tf-{var.STICKY_RANDOM}-1@databricks.com"
	}

	resource "databricks_user" "user2" {
		user_name = "tf-{var.STICKY_RANDOM}-2@databricks.com"
	}

	data "databricks_users" "this" {
		filter = "userName co \"testuser\""
		depends_on = [databricks_user.user1, databricks_user.user2]
	}
`

const dataSourceTemplateExtraAttributes = `
	resource "databricks_group" "admins" {
		display_name = "admins-{var.STICKY_RANDOM}"
	}

	resource "databricks_user" "user1" {
		user_name = "tf-{var.STICKY_RANDOM}-1@databricks.com"
	}

	resource "databricks_group_member" "membership" {
		group_id = databricks_group.admins.id
		member_id = databricks_user.user1.id
	}

	data "databricks_users" "this" {
		filter = "userName eq \"me-{var.STICKY_RANDOM}@example.com\""
		extra_attributes = "groups"
		depends_on = [databricks_group_member.membership]
	}
`

func checkUsersDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_users.this"]
		require.True(t, ok, "data.databricks_users.this has to be there")

		usersCount := ds.Primary.Attributes["users.#"]
		require.Equal(t, "2", usersCount, "expected two users")

		userIds := []string{
			ds.Primary.Attributes["users.0.id"],
			ds.Primary.Attributes["users.1.id"],
		}

		expectedUserIDs := []string{
			s.Modules[0].Resources["databricks_user.user1"].Primary.ID,
			s.Modules[0].Resources["databricks_user.user2"].Primary.ID,
		}

		assert.ElementsMatch(t, expectedUserIDs, userIds, "expected user ids to match")

		return nil
	}
}

func checkUsersDataSourceWithGroups(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_users.this"]
		require.True(t, ok, "data.databricks_users.this must be present")

		usersCount := ds.Primary.Attributes["users.#"]
		require.Equal(t, "1", usersCount, "expected one user")

		userPrefix := "users.0."

		groupsCountAttr := userPrefix + "groups.#"
		groupsCount, exists := ds.Primary.Attributes[groupsCountAttr]
		require.True(t, exists, "attribute groups.# should be present")
		require.Equal(t, "1", groupsCount, "expected one group membership")

		groupIdAttr := userPrefix + "groups.0.value"
		groupId, exists := ds.Primary.Attributes[groupIdAttr]
		require.True(t, exists, "attribute group.0.value should be present")

		expectedGroupId := s.Modules[0].Resources["databricks_group.admins"].Primary.ID
		assert.Equal(t, expectedGroupId, groupId, "group id should match the admins group id")
	}
}

func TestAccDataSourceDataUsers(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: dataSourceTemplate,
		Check:    checkUsersDataSourcePopulated(t),
	})
}

func TestWorkspaceDataSourceDataUsers(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: dataSourceTemplate,
		Check:    checkUsersDataSourcePopulated(t),
	})
}

func TestAccDataSourceUsers_WithGroups(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: dataSourceTemplateExtraAttributes,
		Check:    checkUsersDataSourceWithGroups(t),
	})
}

func TestWorkspaceDataSourceUsers_WithGroups(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: dataSourceTemplateExtraAttributes,
		Check:    checkUsersDataSourceWithGroups(t),
	})
}
