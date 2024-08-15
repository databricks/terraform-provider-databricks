package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

// Application ID is mandatory in Azure today.
func getServicePrincipalResource(t *testing.T) string {
	if isAzure(t) {
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

func TestMwsAccAccountServicePrincipalRuleSetsFullLifeCycle(t *testing.T) {
	loadAccountEnv(t)
	spResource := getServicePrincipalResource(t)
	accountLevel(t, LegacyStep{
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
		Check: resourceCheckWithState("databricks_access_control_rule_set.sp_rule_set",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				id := state.ID
				eTag := state.Attributes["etag"]
				a, err := client.AccountClient()
				if err != nil {
					return err
				}
				ruleSetRes, err := a.AccessControl.GetRuleSet(ctx, iam.GetRuleSetRequest{
					Name: id,
					Etag: eTag,
				})
				if err != nil {
					return err
				}
				assert.Equal(t, len(ruleSetRes.GrantRules), 1)
				return nil
			}),
	})
}

func TestMwsAccAccountGroupRuleSetsFullLifeCycle(t *testing.T) {
	username := qa.RandomEmail()
	accountLevel(t, LegacyStep{
		Template: `
		resource "databricks_user" "this" {
			user_name = "` + username + `"
		}
		resource "databricks_group" "this" {
			display_name = "Group {var.RANDOM}"
		}
		resource "databricks_access_control_rule_set" "group_rule_set" {
			name = "accounts/{env.DATABRICKS_ACCOUNT_ID}/groups/${databricks_group.this.id}/ruleSets/default"
			grant_rules {
				principals = [
					databricks_user.this.acl_principal_id
				]
				role = "roles/group.manager"
			}
		}`,
		Check: resourceCheck("databricks_access_control_rule_set.group_rule_set",
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
