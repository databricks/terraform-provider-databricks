---
subcategory: "FinOps"
---
# databricks_budget Resource

-> **Note** Initialize provider with `alias = "account"`, and `host` pointing to the account URL, like, `host = "https://accounts.cloud.databricks.com"`. Use `provider = databricks.account` for all account-level resources.

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

This resource allows you to manage [Databricks Budgets](https://docs.databricks.com/en/admin/account-settings/budgets.html).

## Example Usage

```hcl
resource "databricks_budget" "this" {
  display_name = "databricks-workspace-budget"

  alert_configurations {
    time_period         = "MONTH"
    trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED"
    quantity_type       = "LIST_PRICE_DOLLARS_USD"
    quantity_threshold  = "840"

    action_configurations {
      action_type = "EMAIL_NOTIFICATION"
      target      = "abc@gmail.com"
    }
  }

  filter {
    workspace_id {
      operator = "IN"
      values   = [
        1234567890098765
      ]
    }
    
    tags {
      key   = "Team"
      value {
        operator = "IN"
        values = ["Data Science"]
      }
    }

    tags {
      key   = "Environment"
      value {
        operator = "IN"
        values = ["Development"]
      }
    }
  }
}
```

## Argument Reference

The following arguments are available:

* `display_name` - (Required) Name of the budget in Databricks Account.

### alert_configurations Configuration Block (Required)

* `time_period` - (Required, String Enum) The time window of usage data for the budget. (Enum: `MONTH`)
* `trigger_type` - (Required, String Enum) The evaluation method to determine when this budget alert is in a triggered state. (Enum: `CUMULATIVE_SPENDING_EXCEEDED`)
* `quantity_type` - (Required, String Enum) The way to calculate cost for this budget alert. This is what quantity_threshold is measured in. (Enum: `LIST_PRICE_DOLLARS_USD`)
* `quantity_threshold` - (Required, String) The threshold for the budget alert to determine if it is in a triggered state. The number is evaluated based on `quantity_type`.
* `action_configurations` - (Required) List of action configurations to take when the budget alert is triggered. Consists of the following fields:
    * `action_type` - (Required, String Enum) The type of action to take when the budget alert is triggered. (Enum: `EMAIL_NOTIFICATION`)
    * `target` - (Required, String) The target of the action. For `EMAIL_NOTIFICATION`, this is the email address to send the notification to.

### filter Configuration Block

* `workspace_id` - (Optional) Filter by workspace ID (if empty, include usage all usage for this account). Consists of the following fields:
    * `operator` - (Required, String Enum) The operator to use for the filter. (Enum: `IN`)
    * `values` - (Required, List of numbers) The values to filter by.
    * `tags` - (Optional) List of tags to filter by. Consists of the following fields:
      * `key` - (Required, String) The key of the tag.
      * `value` - (Required) Consists of the following fields:
          * `operator` - (Required, String Enum) The operator to use for the filter. (Enum: `IN`)
          * `values` - (Required, List of strings) The values to filter by.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `budget_configuration_id` - The ID of the budget configuration.
* `account_id` - The ID of the Databricks Account.
* `alert_configurations` - The alert configurations of the budget.
* `filter` - The filters of the budget.
* `display_name` - The name of the budget in Databricks Account.

## Import

This resource can be imported by Databricks account ID and Budget.

```sh
terraform import databricks_budget.this <account_id>|<budget_configuration_id>
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_workspaces](mws_workspaces.md) to set up Databricks workspaces.
