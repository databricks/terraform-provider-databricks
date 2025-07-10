---
subcategory: "Databricks SQL"
---
# databricks_alert Resource

This resource allows you to manage [Databricks SQL Alerts](https://docs.databricks.com/en/sql/user/alerts/index.html).  It supersedes [databricks_sql_alert](sql_alert.md) resource - see migration guide below for more details.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_directory" "shared_dir" {
  path = "/Shared/Queries"
}

# This will be replaced with new databricks_query resource
resource "databricks_query" "this" {
  warehouse_id = databricks_sql_endpoint.example.id
  display_name = "My Query Name"
  query_text   = "SELECT 42 as value"
  parent_path  = databricks_directory.shared_dir.path
}

resource "databricks_alert" "alert" {
  query_id     = databricks_query.this.id
  display_name = "TF new alert"
  parent_path  = databricks_directory.shared_dir.path
  condition {
    op = "GREATER_THAN"
    operand {
      column {
        name = "value"
      }
    }
    threshold {
      value {
        double_value = 42
      }
    }
  }
}
```

## Argument Reference

The following arguments are available:

* `query_id` - (Required, String) ID of the query evaluated by the alert.
* `display_name` - (Required, String) Name of the alert.
* `condition` - (Required) Trigger conditions of the alert. Block consists of the following attributes:
  * `op` - (Required, String Enum) Operator used for comparison in alert evaluation. (Enum: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `EQUAL`, `NOT_EQUAL`, `IS_NULL`)
  * `operand` - (Required, Block) Name of the column from the query result to use for comparison in alert evaluation:
    * `column` - (Required, Block) Block describing the column from the query result to use for comparison in alert evaluation:
      * `name` - (Required, String) Name of the column.
  * `threshold` - (Optional for `IS_NULL` operation, Block) Threshold value used for comparison in alert evaluation:
    * `value` - (Required, Block) actual value used in comparison (one of the attributes is required):
      * `string_value` - string value to compare against string results.
      * `double_value` - double value to compare against integer and double results.
      * `bool_value` - boolean value (`true` or `false`) to compare against boolean results.
  * `empty_result_state` - (Optional, String Enum) Alert state if the result is empty (`UNKNOWN`, `OK`, `TRIGGERED`)
* `custom_subject` - (Optional, String) Custom subject of alert notification, if it exists. This includes email subject, Slack notification header, etc. See [Alerts API reference](https://docs.databricks.com/en/sql/user/alerts/index.html) for custom templating instructions.
* `custom_body` - (Optional, String) Custom body of alert notification, if it exists. See [Alerts API reference](https://docs.databricks.com/en/sql/user/alerts/index.html) for custom templating instructions.
* `parent_path` - (Optional, String) The path to a workspace folder containing the alert. The default is the user's home folder.  If changed, the alert will be recreated.
* `seconds_to_retrigger` - (Optional, Integer) Number of seconds an alert must wait after being triggered to rearm itself. After rearming, it can be triggered again. If 0 or not specified, the alert will not be triggered again.
* `owner_user_name` - (Optional, String) Alert owner's username.
* `notify_on_ok` - (Optional, Boolean) Whether to notify alert subscribers when alert returns back to normal.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - unique ID of the Alert.
* `lifecycle_state` - The workspace state of the alert. Used for tracking trashed status. (Possible values are `ACTIVE` or `TRASHED`).
* `state` - Current state of the alert's trigger status (`UNKNOWN`, `OK`, `TRIGGERED`). This field is set to `UNKNOWN` if the alert has not yet been evaluated or ran into an error during the last evaluation.
* `create_time` - The timestamp string indicating when the alert was created.
* `update_time` - The timestamp string indicating when the alert was updated.
* `trigger_time` - The timestamp string when the alert was last triggered if the alert has been triggered before.

## Migrating from `databricks_sql_alert` resource

Under the hood, the new resource uses the same data as the `databricks_sql_alert`, but is exposed via a different API. This means that we can migrate existing alerts without recreating them.

-> It's also recommended to migrate to the `databricks_query` resource - see [databricks_query](query.md) for more details.

This operation is done in few steps:

* Record the ID of existing `databricks_sql_alert`, for example, by executing the `terraform state show databricks_sql_alert.alert` command.
* Create the code for the new implementation by performing the following changes:
  * the `name` attribute is now named `display_name`
  * the `parent` (if exists) is renamed to `parent_path` attribute and should be converted from `folders/object_id` to the actual path.
  * the `options` block is converted into the `condition` block with the following changes:
    * the value of the `op` attribute should be converted from a mathematical operator into a string name, like, `>` is becoming `GREATER_THAN`, `==` is becoming `EQUAL`, etc.
    * the `column` attribute is becoming the `operand` block
    * the `value` attribute is becoming the `threshold` block.  **Please note that the old implementation always used strings so you may have changes after import if you use `double_value` or `bool_value` inside the block.**
  * the `rearm` attribute is renamed to `seconds_to_retrigger`.

For example, if we have the original `databricks_sql_alert` defined as:

```hcl
resource "databricks_sql_alert" "alert" {
  query_id = databricks_sql_query.this.id
  name     = "My Alert"
  parent   = "folders/${databricks_directory.shared_dir.object_id}"
  options {
    column = "value"
    op     = ">"
    value  = "42"
    muted  = false
  }
}
```

we'll have a new resource defined as:

```hcl
resource "databricks_alert" "alert" {
  query_id     = databricks_query.this.id
  display_name = "My Alert"
  parent_path  = databricks_directory.shared_dir.path
  condition {
    op = "GREATER_THAN"
    operand {
      column {
        name = "value"
      }
    }
    threshold {
      value {
        double_value = 42
      }
    }
  }
}
```

### For Terraform version >= 1.7.0

Terraform 1.7 introduced the [removed](https://developer.hashicorp.com/terraform/language/resources/syntax#removing-resources) block in addition to the [import](https://developer.hashicorp.com/terraform/language/import) block introduced in Terraform 1.5.  Together they make import and removal of resources easier, avoiding manual execution of `terraform import` and `terraform state rm` commands.

So with Terraform 1.7+, the migration looks as the following:

* remove the old alert definition and replace it with the new one.
* Adjust references, like, `databricks_permissions`.
* Add `import` and `removed` blocks like this:

```hcl
import {
  to = databricks_alert.alert
  id = "<alert-id>"
}

removed {
  from = databricks_sql_alert.alert

  lifecycle {
    destroy = false
  }
}
```

* Run the `terraform plan` command to check possible changes, such as value type change, etc.
* Run the `terraform apply` command to apply changes.
* Remove the `import` and `removed` blocks from the code.

### For Terraform version < 1.7.0

* Remove the old alert definition and replace it with the new one.
* Remove the old resource from the state with the `terraform state rm databricks_sql_alert.alert` command.
* Import new resource with the `terraform import databricks_alert.alert <alert-id>` command.
* Adjust references, like, `databricks_permissions`.
* Run the `terraform plan` command to check possible changes, such as value type change, etc.

## Access Control

[databricks_permissions](permissions.md#sql-alert-usage) can control which groups or individual users can *Manage*, *Edit*, *Run* or *View* individual alerts.

```hcl
resource "databricks_permissions" "alert_usage" {
  sql_alert_id = databricks_alert.alert.id
  access_control {
    group_name       = "users"
    permission_level = "CAN_RUN"
  }
}
```

## Access Control

[databricks_permissions](permissions.md#sql-alert-usage) can control which groups or individual users can *Manage*, *Edit*, *Run* or *View* individual alerts.

```hcl
resource "databricks_permissions" "alert_usage" {
  sql_alert_id = databricks_alert.alert.id
  access_control {
    group_name       = "users"
    permission_level = "CAN_RUN"
  }
}
```

## Import

This resource can be imported using alert ID:

```hcl
import {
  to = databricks_alert.this
  id = "<alert-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_alert.this <alert-id>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_query](query.md) to manage [Databricks SQL Queries](https://docs.databricks.com/sql/user/queries/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage [Databricks SQL Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
