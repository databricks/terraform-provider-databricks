package finops_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var (
	// The filter's workspace must belong to the test account: the API rejects
	// budgets that reference any other workspace ID.
	budgetTemplate = `resource "databricks_budget" "this" {
			display_name = "tf-{var.RANDOM}"

			alert_configurations {
				time_period         = "MONTH"
				trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED"
				quantity_type       = "LIST_PRICE_DOLLARS_USD"
				quantity_threshold  = "%s"

				action_configurations {
					action_type = "EMAIL_NOTIFICATION"
					target      = "me@databricks.com"
				}
			}

			filter {
				tags {
					key   = "Environment"
					value {
						operator = "IN"
						values = ["Testing"]
					}
				}

				workspace_id {
					operator = "IN"
					values   = [
						{env.TEST_WORKSPACE_ID}
					]
				}
			}
		}`
)

func TestMwsAccBudgetCreate(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(budgetTemplate, "840"),
	})
}

func TestMwsAccBudgetUpdate(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(budgetTemplate, "840"),
	}, acceptance.Step{
		Template: fmt.Sprintf(budgetTemplate, "940"),
	})
}

const (
	// The API returns alerts and actions in an order unrelated to the configuration (see
	// issue #5978). Thresholds are high so that BLOCK_USAGE never blocks the workspace.
	budgetAlertsTemplate = `resource "databricks_group" "this" {
			display_name = "tf-budget-{var.STICKY_RANDOM}"
		}

		resource "databricks_budget" "this" {
			display_name  = "tf-{var.STICKY_RANDOM}"
			resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

			filter {
				workspace_id {
					operator = "IN"
					values   = [{env.TEST_WORKSPACE_ID}]
				}
			}
			%s
		}`
	sharedEmailAlert = `
			alert_configurations {
				time_period        = "MONTH"
				trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
				quantity_type      = "LIST_PRICE_DOLLARS_USD"
				quantity_threshold = "800000"
				scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"
				action_configurations {
					action_type = "EMAIL_NOTIFICATION"
					target      = "me@databricks.com"
				}
			}`
	sharedBlockAlert = `
			alert_configurations {
				time_period        = "MONTH"
				trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
				quantity_type      = "LIST_PRICE_DOLLARS_USD"
				quantity_threshold = "1500000"
				scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"
				action_configurations {
					action_type = "BLOCK_USAGE"
				}
				action_configurations {
					action_type = "EMAIL_NOTIFICATION"
					target      = "me@databricks.com"
				}
			}`
	perUserAlert = `
			alert_configurations {
				time_period        = "MONTH"
				trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
				quantity_type      = "LIST_PRICE_DOLLARS_USD"
				quantity_threshold = "%s"
				scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER"
				principal_overrides {
					principal_id       = databricks_group.this.id
					override_threshold = "%s"
				}
				action_configurations {
					action_type = "BLOCK_USAGE"
				}
			}`
)

func TestMwsAccBudgetAlerts(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	perUser := fmt.Sprintf(perUserAlert, "200000", "250000")
	created := sharedEmailAlert + sharedBlockAlert + perUser
	reordered := perUser + sharedEmailAlert + sharedBlockAlert
	updated := sharedBlockAlert + fmt.Sprintf(perUserAlert, "300000", "350000")
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(budgetAlertsTemplate, created),
	}, acceptance.Step{
		Template: fmt.Sprintf(budgetAlertsTemplate, reordered),
		PlanOnly: true,
	}, acceptance.Step{
		Template: fmt.Sprintf(budgetAlertsTemplate, updated),
	})
}
