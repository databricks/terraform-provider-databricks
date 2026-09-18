---
subcategory: "Billing"
---
# databricks_budget Resource

[API Documentation](https://docs.databricks.com/api/account/budgets)

This resource allows you to manage [Databricks Budgets](https://docs.databricks.com/en/admin/account-settings/budgets.html).

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

-> This resource can only be used with an account-level provider!

## Example Usage

```hcl
resource "databricks_budget" "this" {
  display_name = "databricks-workspace-budget"

  alert_configurations {
    time_period        = "MONTH"
    trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
    quantity_type      = "LIST_PRICE_DOLLARS_USD"
    quantity_threshold = "840"

    action_configurations {
      action_type = "EMAIL_NOTIFICATION"
      target      = "abc@gmail.com"
    }
  }

  filter {
    workspace_id {
      operator = "IN"
      values = [
        1234567890098765
      ]
    }

    tags {
      key = "Team"
      value {
        operator = "IN"
        values   = ["Data Science"]
      }
    }

    tags {
      key = "Environment"
      value {
        operator = "IN"
        values   = ["Development"]
      }
    }
  }
}
```

### Budgets for AI Gateway Resources

Budgets can also be scoped to track only spend through Unity AI Gateway endpoints by setting `resource_type` to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`.

This includes Databricks products that may use Unity AI Gateway endpoints, such as Databricks Genie.

#### AI Gateway Budget for all endpoints

Create a shared budget tracking all costs for Unity AI Gateway endpoints, and send an email when the budget threshold is exceeded.

```hcl
resource "databricks_budget" "ai_gateway_shared_budget" {
  display_name  = "aigw-shared-budget"
  resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

  alert_configurations {
    quantity_threshold = "10000"
    quantity_type      = "LIST_PRICE_DOLLARS_USD"
    trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
    time_period        = "MONTH"
    scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"

    action_configurations {
      action_type = "EMAIL_NOTIFICATION"
      target      = "abc@gmail.com"
    }
  }
}
```

#### Shared Genie budget

Genie budgets use the Unity AI Gateway resource type and the `databricks-product: genie` tag. Do not add other resource tags to a Genie budget.

A shared Genie budget for all users. Spend is tracked in aggregate.

```hcl
resource "databricks_budget" "genie_shared_budget" {
  display_name  = "genie-shared-budget"
  resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

  // Apply the filter on databricks-product tag
  filter {
    tags {
      key = "databricks-product"
      value {
        operator = "IN"
        values   = ["genie"]
      }
    }
  }

  alert_configurations {
    quantity_threshold = "2000"
    quantity_type      = "LIST_PRICE_DOLLARS_USD"
    trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
    time_period        = "MONTH"
    scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"

    action_configurations {
      action_type = "EMAIL_NOTIFICATION"
      target      = "abc@gmail.com"
    }
  }
}
```

#### Per-user budget overrides with block usage

A per-user threshold applies to each user in the budget's scope. Use `principal_overrides` to override the threshold for specific users, groups, or service principals. `BLOCK_USAGE` prevents further requests through Unity AI Gateway when the threshold is reached.

```hcl
// Find a group we want to apply
data "databricks_group" "genie-power-users" {
  display_name = "genie-power-users"
}

// Create a budget for Genie of $100 per user, and override
// the threshold to $300 per user for the genie-power-users group.
resource "databricks_budget" "genie_per_user_budget" {
  display_name  = "genie-tier-1-budget"
  resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

  filter {
    tags {
      key = "databricks-product"
      value {
        operator = "IN"
        values   = ["genie"]
      }
    }
  }

  alert_configurations {
    // Default the budget threshold to 100 per user
    quantity_threshold = "100"
    quantity_type      = "LIST_PRICE_DOLLARS_USD"
    trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
    time_period        = "MONTH"
    scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER"

    // Override the threshold to 300 for a power user group
    principal_overrides {
      principal_id       = databricks_group.genie-power-users.id
      override_threshold = "300"
    }

    action_configurations {
      action_type = "BLOCK_USAGE"
    }
  }
}
```

## Argument Reference

The following arguments are available:

* `display_name` - (Required) Name of the budget in Databricks Account.
* `resource_type` - (Optional, String Enum) The resource scope for this budget. Determines whether the budget tracks all resources or a specific resource. (Enum: `BUDGET_RESOURCE_TYPE_ALL_RESOURCES`, `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`)

### alert_configurations Configuration Block (Required)

* `time_period` - (Required, String Enum) The time window of usage data for the budget. (Enum: `MONTH`)
* `trigger_type` - (Required, String Enum) The evaluation method to determine when this budget alert is in a triggered state. (Enum: `CUMULATIVE_SPENDING_EXCEEDED`)
* `quantity_type` - (Required, String Enum) The way to calculate cost for this budget alert. This is what quantity_threshold is measured in. (Enum: `LIST_PRICE_DOLLARS_USD`)
* `quantity_threshold` - (Required, String) The threshold for the budget alert to determine if it is in a triggered state. The number is evaluated based on `quantity_type`.
* `scope_type` - (Optional, String Enum) How the alert threshold is evaluated. Determines whether spend is tracked in aggregate or per individual user. (Enum: `ALERT_CONFIGURATION_SCOPE_TYPE_SHARED`, `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`)
* `principal_overrides` - (Optional) Per-principal threshold overrides for this alert. Only applies to per-user alerts (`scope_type` = `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`); ignored for shared alerts. Consists of the following fields:
  * `principal_id` - (Optional, Integer) Account-level principal id (user, group, or service principal).
  * `override_threshold` - (Optional, String) Dollar amount that overrides the parent alert's `quantity_threshold` for this principal.
* `action_configurations` - (Required) List of action configurations to take when the budget alert is triggered. Consists of the following fields:
  * `action_type` - (Required, String Enum) The type of action to take when the budget alert is triggered. (Enum: `EMAIL_NOTIFICATION`, `BLOCK_USAGE`). Note: `BLOCK_USAGE` action type is only supported on AI Gateway-scoped budgets.
  * `target` - (Optional, String) For `EMAIL_NOTIFICATION` action type, this is the email address to send the notification to. Required if the `action_type` is `EMAIL_NOTIFICATION`. Does not apply to the `BLOCK_USAGE` action type, and this field must be omitted in that case.

### filter Configuration Block (Optional)

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

## Import

This resource can be imported by Databricks account ID and Budget:

```hcl
import {
  to = databricks_budget.this
  id = "<account_id>|<budget_configuration_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_budget.this "<account_id>|<budget_configuration_id>"
```

## Related Resources

The following resources are used in the context:

* [databricks_mws_workspaces](mws_workspaces.md) to set up Databricks workspaces.
