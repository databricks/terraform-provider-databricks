---
subcategory: "Settings"
---

# databricks_automatic_cluster_update_workspace_setting Resource

-> **Note** This resource could be only used with workspace-level provider!

The `databricks_automatic_cluster_update_workspace_setting` resource allows you to control whether automatic cluster update is enabled for the current workspace. By default, it is turned off. Enabling this feature on a workspace requires that you add the Enhanced Security and Compliance add-on.

## Example Usage

```hcl
resource "databricks_automatic_cluster_update_workspace_setting" "this" {
  automatic_cluster_update_workspace {
    enabled                              = true
    restart_even_if_no_updates_available = true
    maintenance_window {
      week_day_based_schedule {
        day_of_week = "MONDAY"
        frequency   = "EVERY_WEEK"
        window_start_time {
          hours   = 1
          minutes = 0
        }
      }
    }
  }
}
```

## Argument Reference

The resource supports the following arguments:

- `automatic_cluster_update_workspace` (Required) block with following attributes
  - `enabled` - (Required) The configuration details.
  - `restart_even_if_no_updates_available` - (Optional) To force clusters and other compute resources to restart during the maintenance window regardless of the availability of a new update.
  - `maintenance_window` block that defines the maintenance frequency with the following arguments
    - `week_day_based_schedule` block with the following arguments
      - `day_of_week` - the day of the week in uppercase, e.g. `MONDAY` or `SUNDAY`
      - `frequency` - one of the `FIRST_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`, `FOURTH_OF_MONTH`, `FIRST_AND_THIRD_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `EVERY_WEEK`.
      - `window_start_time` block that defines the time of your maintenance window. The default timezone is UTC and cannot be changed.
        - `hours` - hour to perform update: 0-23
        - `minutes` - minute to perform update: 0-59

## Import

This resource can be imported by predefined name `global`:

```bash
terraform import databricks_automatic_cluster_update_workspace_setting.this global
```
