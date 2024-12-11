package acceptance

import (
	"fmt"
	"testing"
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
	LoadAccountEnv(t)
	if IsGcp(t) {
		skipf(t)("not available on GCP")
	}
	AccountLevel(t, Step{
		Template: fmt.Sprintf(budgetTemplate, "840"),
	})
}

func TestMwsAccBudgetUpdate(t *testing.T) {
	LoadAccountEnv(t)
	if IsGcp(t) {
		skipf(t)("not available on GCP")
	}
	AccountLevel(t, Step{
		Template: fmt.Sprintf(budgetTemplate, "840"),
	}, Step{
		Template: fmt.Sprintf(budgetTemplate, "940"),
	})
}
