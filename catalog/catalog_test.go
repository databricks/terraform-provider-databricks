package catalog_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestUcAccCatalog(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
			%s
		}`, getPredictiveOptimizationSetting(t, true)),
	})
}

func TestUcAccCatalogIsolated(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "ISOLATED"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "OPEN"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	})
}

type checkResourceRecreate struct {
	address string
}

func (c checkResourceRecreate) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	var change *tfjson.ResourceChange
	for _, resourceChange := range req.Plan.ResourceChanges {
		if resourceChange.Address == c.address {
			change = resourceChange
			break
		}
	}
	if change == nil {
		addressesWithPlannedChanges := make([]string, 0, len(req.Plan.ResourceChanges))
		for _, change := range req.Plan.ResourceChanges {
			addressesWithPlannedChanges = append(addressesWithPlannedChanges, change.Address)
		}
		resp.Error = fmt.Errorf("address %s not found in resource changes; only planned changes for addresses %s", c.address, strings.Join(addressesWithPlannedChanges, ", "))
		return
	}
	if change.Change.Actions[0] != tfjson.ActionDelete {
		plannedActions := make([]string, 0, len(change.Change.Actions))
		for _, action := range change.Change.Actions {
			plannedActions = append(plannedActions, string(action))
		}
		resp.Error = fmt.Errorf("no delete is planned for %s; planned actions are: %s", c.address, strings.Join(plannedActions, ", "))
	}
}

func TestUcAccCatalogUpdate(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "account users"
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "{env.TEST_DATA_ENG_GROUP}"
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform - updated comment"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
		}`, getPredictiveOptimizationSetting(t, false)),
	}, acceptance.Step{
		// Adding options should cause the catalog to be recreated.
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform - updated comment"
			properties     = {
				purpose = "testing"
			}
			options = {
				user = "miles"
			}
			%s
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
		}`, getPredictiveOptimizationSetting(t, false)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				checkResourceRecreate{address: "databricks_catalog.sandbox"},
			},
		},
	})
}

// Create a connection to an HMS catalog and update authorized_paths.
func TestUcAccCatalogHmsConnectionUpdate(t *testing.T) {
	authorizedPath := fmt.Sprintf("s3://%s/path/to/authorized", qa.RandomName("hms-bucket-"))
	otherAuthorizedPath := fmt.Sprintf("s3://%s/path/to/authorized", qa.RandomName("hms-other-bucket-"))
	otherInfra := fmt.Sprintf(`
		resource "databricks_connection" "sandbox" {
			name = "hms_connection{var.STICKY_RANDOM}"
			connection_type = "HIVE_METASTORE"
			comment         = "created in TestUcAccCatalogHmsConnectionUpdate"
			options = {
				host     = "test.mysql.database.azure.com"
				port     = "3306"
				user     = "user"
				password = "password"
				database = "metastore{var.STICKY_RANDOM}"
				db_type  = "MYSQL"
				version  = "2.3"
			}
			properties = {
				purpose = "testing"
			}
		}
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
		}
		resource "databricks_external_location" "sandbox" {
		    name = "sandbox{var.STICKY_RANDOM}"
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
			url = "%s"
			credential_name = "${databricks_storage_credential.external.name}"
			skip_validation = true
		}
		resource "databricks_external_location" "sandbox-other" {
		    name = "sandbox-other{var.STICKY_RANDOM}"
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
			url = "%s"
			credential_name = "${databricks_storage_credential.external.name}"
			skip_validation = true
		}
	`, authorizedPath, otherAuthorizedPath)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s"
			}
			lifecycle {
			    prevent_destroy = true
			}
			depends_on = [databricks_external_location.sandbox, databricks_external_location.sandbox-other]
		}`, authorizedPath),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
			}
			lifecycle {
			    prevent_destroy = true
			}
		}`, authorizedPath, otherAuthorizedPath),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
				other_option = "value"
			}
			lifecycle {
			    prevent_destroy = true
			}
		}`, authorizedPath, otherAuthorizedPath),
		ExpectError: regexp.MustCompile("Instance cannot be destroyed"),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
			}
		}`, authorizedPath, otherAuthorizedPath),
	})
}
