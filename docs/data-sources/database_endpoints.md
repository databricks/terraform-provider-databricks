---
subcategory: "Database Instances"
---
# databricks_database_endpoints Data Source
[![Development](https://img.shields.io/badge/Release_Stage-Development-red)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `branch_id` (string, required)
* `project_id` (string, required)
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource


## Attributes
This data source exports a single attribute, `database_endpoints`. It is a list of resources, each with the following attributes:
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `branch_id` (string)
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `current_state` (string) - . Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_id` (string)
* `host` (string) - The hostname of the compute endpoint. This is the hostname specified when connecting to a database
* `last_active_time` (string) - A timestamp indicating when the compute endpoint was last active
* `pending_state` (string) - . Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `pooler_mode` (string) - . Possible values are: `TRANSACTION`
* `project_id` (string)
* `settings` (DatabaseEndpointSettings)
* `start_time` (string) - A timestamp indicating when the compute endpoint was last started
* `suspend_time` (string) - A timestamp indicating when the compute endpoint was last suspended
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended
* `type` (string) - NOTE: if want type to default to some value set the server then an effective_type field
  OR make this field REQUIRED. Possible values are: `READ_ONLY`, `READ_WRITE`
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

### DatabaseEndpointSettings
* `pg_settings` (object) - A raw representation of Postgres settings
* `pgbouncer_settings` (object) - A raw representation of PgBouncer settings