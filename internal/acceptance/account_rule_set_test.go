package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/common"
)

func TestMwsAccAccountRuleSetsFullLifeCycle(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_service_principal" "this" {
			display_name = "SPN {var.RANDOM}"
		}
		resource "databricks_group" "this" {
			display_name = "Group {var.RANDOM}"
		}
		resource "databricks_rule_set" "sp_rule_set" {
			name = "accounts/{env.DATABRICKS_ACCOUNT_ID}/servicePrincipals/${databricks_service_principal.this.application_id}/ruleSets/default"
			grant_rules {
				principals = [
					"groups/${databricks_group.this.display_name}"
				]
				role = "roles/servicePrincipal.manager"
			}
		}`,
		Check: resourceCheck("databricks_rule_set.sp_rule_set",
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
