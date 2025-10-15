---
subcategory: "Database Instances"
---
# databricks_database_endpoint Resource
[![Development](https://img.shields.io/badge/Release_Stage-Development-red)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `endpoint_id` (string, required)
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units
* `branch_id` (string, optional)
* `disabled` (boolean, optional) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `pooler_mode` (string, optional) - . Possible values are: `TRANSACTION`
* `project_id` (string, optional)
* `settings` (DatabaseEndpointSettings, optional)
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended
* `type` (string, optional) - NOTE: if want type to default to some value set the server then an effective_type field
  OR make this field REQUIRED. Possible values are: `READ_ONLY`, `READ_WRITE`
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

### DatabaseEndpointSettings
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `pgbouncer_settings` (object, optional) - A raw representation of PgBouncer settings

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `current_state` (string) - . Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `host` (string) - The hostname of the compute endpoint. This is the hostname specified when connecting to a database
* `last_active_time` (string) - A timestamp indicating when the compute endpoint was last active
* `pending_state` (string) - . Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `start_time` (string) - A timestamp indicating when the compute endpoint was last started
* `suspend_time` (string) - A timestamp indicating when the compute endpoint was last suspended
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "project_id,branch_id,endpoint_id"
  to = databricks_database_endpoint.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_database_endpoint "project_id,branch_id,endpoint_id"
```