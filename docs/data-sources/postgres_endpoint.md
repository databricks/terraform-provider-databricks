---
subcategory: "Postgres"
---
# databricks_postgres_endpoint Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves a single Postgres endpoint.


## Example Usage
### Retrieve Endpoint by Name

```hcl
data "databricks_postgres_endpoint" "this" {
  name = "projects/my-project/branches/dev-branch/endpoints/primary"
}

output "endpoint_type" {
  value = data.databricks_postgres_endpoint.this.status.endpoint_type
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The full resource path of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `name` (string) - Output only. The full resource path of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
* `parent` (string) - The branch containing this endpoint (API resource hierarchy).
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (EndpointSpec) - The spec contains the compute endpoint configuration, including autoscaling limits, suspend timeout, and disabled state
* `status` (EndpointStatus) - Current operational status of the compute endpoint
* `uid` (string) - System-generated unique ID for the endpoint
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

### EndpointGroupSpec
* `enable_readable_secondaries` (boolean) - Whether to allow read-only connections to read-write endpoints. Only relevant for read-write endpoints where
  size.max > 1
* `max` (integer) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1

### EndpointGroupStatus
* `enable_readable_secondaries` (boolean) - Whether read-only connections to read-write endpoints are allowed. Only relevant if read replicas are configured
  by specifying size.max > 1
* `max` (integer) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1

### EndpointHosts
* `host` (string) - The hostname to connect to this endpoint. For read-write endpoints, this is a read-write hostname which connects
  to the primary compute. For read-only endpoints, this is a read-only hostname which allows read-only operations
* `read_only_host` (string) - An optionally defined read-only host for the endpoint, without pooling. For read-only endpoints,
  this attribute is always defined and is equivalent to host. For read-write endpoints, this attribute is defined
  if the enclosing endpoint is a group with greater than 1 computes configured, and has readable secondaries enabled

### EndpointSettings
* `pg_settings` (object) - A raw representation of Postgres settings

### EndpointSpec
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units. Minimum value is 0.5
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `group` (EndpointGroupSpec) - Settings for optional HA configuration of the endpoint. If unspecified, the endpoint defaults
  to non HA settings, with a single compute backing the endpoint (and no readable secondaries
  for Read/Write endpoints)
* `no_suspension` (boolean) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided
* `settings` (EndpointSettings)
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

### EndpointStatus
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `current_state` (string) - Possible values are: `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `group` (EndpointGroupStatus) - Details on the HA configuration of the endpoint
* `hosts` (EndpointHosts) - Contains host information for connecting to the endpoint
* `pending_state` (string) - Possible values are: `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`
* `settings` (EndpointSettings)
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended