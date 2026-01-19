---
subcategory: "Postgres"
---
# databricks_postgres_endpoints Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Branch that owns this collection of endpoints.
  Format: projects/{project_id}/branches/{branch_id}
* `page_size` (integer, optional) - Upper bound for items returned


## Attributes
This data source exports a single attribute, `endpoints`. It is a list of resources, each with the following attributes:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `name` (string) - The resource name of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
* `parent` (string) - The branch containing this endpoint.
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (EndpointSpec) - The desired state of an Endpoint
* `status` (EndpointStatus) - The current status of an Endpoint
* `uid` (string) - System generated unique ID for the endpoint
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

### EndpointSettings
* `pg_settings` (object) - A raw representation of Postgres settings

### EndpointSpec
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `READ_ONLY`, `READ_WRITE`
* `settings` (EndpointSettings)
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended

### EndpointStatus
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `current_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `READ_ONLY`, `READ_WRITE`
* `host` (string) - The hostname of the compute endpoint. This is the hostname specified when connecting to a database
* `last_active_time` (string) - A timestamp indicating when the compute endpoint was last active
* `pending_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `settings` (EndpointSettings)
* `start_time` (string) - A timestamp indicating when the compute endpoint was last started
* `suspend_time` (string) - A timestamp indicating when the compute endpoint was last suspended
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended