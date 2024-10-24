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
		user_name = "testuser1@databricks.com"
	}

	resource "databricks_user" "user2" {
		user_name = "testuser2@databricks.com"
	}

	data "databricks_users" "this" {
		user_name_contains = "testuser"
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
