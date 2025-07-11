package finops_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var (
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
						1234567890098765
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
