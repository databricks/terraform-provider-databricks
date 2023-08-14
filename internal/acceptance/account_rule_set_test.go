package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/common"
)

// Application ID is mandatory in Azure today.
func getServicePrincipalResource(cloudEnv string) string {
	if cloudEnv == "azure" {
		return `
		resource "databricks_service_principal" "this" {
			application_id = "{var.RANDOM_UUID}"
			display_name = "SPN {var.RANDOM}"
		}
		`
	}
	return `
	resource "databricks_service_principal" "this" {
		display_name = "SPN {var.RANDOM}"
	}
	`
}

func TestMwsAccAccountRuleSetsFullLifeCycle(t *testing.T) {
	// This endpoint is restricted to basic auth today, used only by AWS account-level tests.
	// Remove this skip when this restriction is lifted in Azure & GCP.
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "aws" {
		t.Skip("Skipping test in Azure")
	}
	spResource := getServicePrincipalResource(cloudEnv)
	accountLevel(t, step{
		Template: spResource + `
		resource "databricks_group" "this" {
			display_name = "Group {var.RANDOM}"
		}
		resource "databricks_access_control_rule_set" "sp_rule_set" {
			name = "accounts/{env.DATABRICKS_ACCOUNT_ID}/servicePrincipals/${databricks_service_principal.this.application_id}/ruleSets/default"
			grant_rules {
				principals = [
					databricks_group.this.acl_principal_id
				]
				role = "roles/servicePrincipal.manager"
			}
		}`,
		Check: resourceCheck("databricks_access_control_rule_set.sp_rule_set",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				a, err := client.AccountClient()
				if err != nil {
					return err
				}
				ruleSetRes, err := a.AccessControl.GetRuleSet(ctx, iam.GetRuleSetRequest{
					Name: id,
					Etag: "",
				})
				if err != nil {
					return err
				}
				assert.Equal(t, len(ruleSetRes.GrantRules), 1)
				return nil
			}),
	})
}
