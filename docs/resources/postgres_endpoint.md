---
subcategory: "Postgres"
---
# databricks_postgres_endpoint Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `endpoint_type` (string, required) - The endpoint type. There could be only one READ_WRITE endpoint per branch. Possible values are: `READ_ONLY`, `READ_WRITE`
* `parent` (string, required) - The branch containing this endpoint.
  Format: projects/{project_id}/branches/{branch_id}
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units
* `disabled` (boolean, optional) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_id` (string, optional) - The ID to use for the Endpoint, which will become the final component of
  the endpoint's resource name.
  
  This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/
* `pooler_mode` (string, optional) - Possible values are: `TRANSACTION`
* `settings` (EndpointSettings, optional)
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended

### EndpointSettings
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `pgbouncer_settings` (object, optional) - A raw representation of PgBouncer settings

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `current_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `effective_autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `effective_autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `effective_disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `effective_pooler_mode` (string) - Possible values are: `TRANSACTION`
* `effective_settings` (EndpointSettings)
* `effective_suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended
* `host` (string) - The hostname of the compute endpoint. This is the hostname specified when connecting to a database
* `last_active_time` (string) - A timestamp indicating when the compute endpoint was last active
* `name` (string) - The resource name of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
* `pending_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `start_time` (string) - A timestamp indicating when the compute endpoint was last started
* `suspend_time` (string) - A timestamp indicating when the compute endpoint was last suspended
* `uid` (string) - System generated unique ID for the endpoint
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_endpoint.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_endpoint.this "name"
```