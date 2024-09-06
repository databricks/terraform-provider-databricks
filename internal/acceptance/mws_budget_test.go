package acceptance

import "testing"

func TestMwsAccBudget(t *testing.T) {
	AccountLevel(t, Step{
		Template: `resource "databricks_mws_budget" "this" {
			display_name = "tf-{var.RANDOM}"
		
			alert_configurations {
				time_period         = "MONTH"
				trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED" 
				quantity_type       = "LIST_PRICE_DOLLARS_USD"
				quantity_threshold  = "840"
		
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
		}`,
	})
}
