package mws_permission_assignments_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestUcAccPermissionAssignmentsDataSource creates a workspace permission
// assignment for a freshly created group and asserts the
// databricks_mws_permission_assignments data source lists it back. Runs against
// an account-level (Unity) provider; DUMMY_WORKSPACE_ID must point at a
// workspace the account can manage assignments for. The TestUcAcc prefix pairs
// with UnityAccountLevel (TEST_ENVIRONMENT_TYPE=UC_ACCOUNT); a TestMwsAcc name
// would be skipped by skipIfNotEnvironmentType in the non-UC account env.
func TestUcAccPermissionAssignmentsDataSource(t *testing.T) {
	acceptance.UnityAccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}

		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["USER"]
		}

		data "databricks_mws_permission_assignments" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			depends_on   = [databricks_mws_permission_assignment.this]
		}`,
		Check: func(s *terraform.State) error {
			group, ok := s.RootModule().Resources["databricks_group.this"]
			if !ok {
				return fmt.Errorf("databricks_group.this is missing from the state")
			}
			principalID := group.Primary.ID

			ds, ok := s.RootModule().Resources["data.databricks_mws_permission_assignments.this"]
			if !ok {
				return fmt.Errorf("data.databricks_mws_permission_assignments.this is missing from the state")
			}
			count, err := strconv.Atoi(ds.Primary.Attributes["permission_assignments.#"])
			if err != nil {
				return fmt.Errorf("permission_assignments count is not a number: %w", err)
			}
			for i := 0; i < count; i++ {
				id := ds.Primary.Attributes[fmt.Sprintf("permission_assignments.%d.principal.principal_id", i)]
				if id != principalID {
					continue
				}
				if perm := ds.Primary.Attributes[fmt.Sprintf("permission_assignments.%d.permissions.0", i)]; perm != "USER" {
					return fmt.Errorf("expected principal %s to have USER permission, got %q", principalID, perm)
				}
				return nil
			}
			return fmt.Errorf("assigned principal %s not returned by the data source (%d listed)", principalID, count)
		},
	})
}
