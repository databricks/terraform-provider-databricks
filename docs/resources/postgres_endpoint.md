---
subcategory: "Postgres"
---
# databricks_postgres_endpoint Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `endpoint_id` (string, required) - The ID to use for the Endpoint. This becomes the final component of the endpoint's resource name.
  The ID must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens (RFC 1123).
  Examples:
  - With custom ID: `primary` → name becomes `projects/{project_id}/branches/{branch_id}/endpoints/primary`
  - Without custom ID: system generates slug → name becomes `projects/{project_id}/branches/{branch_id}/endpoints/ep-example-name-x1y2z3a4`
* `parent` (string, required) - The branch containing this endpoint (API resource hierarchy).
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (EndpointSpec, optional) - The spec contains the compute endpoint configuration, including autoscaling limits, suspend timeout, and disabled state

### EndpointSettings
* `pg_settings` (object, optional) - A raw representation of Postgres settings

### EndpointSpec
* `endpoint_type` (string, required) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units. Minimum value is 0.5
* `disabled` (boolean, optional) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `settings` (EndpointSettings, optional)
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `name` (string) - The resource name of the endpoint. This field is output-only and constructed by the system.
  Format: `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
* `status` (EndpointStatus) - Current operational status of the compute endpoint
* `uid` (string) - System-generated unique ID for the endpoint
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

### EndpointHosts
* `host` (string) - The hostname to connect to this endpoint. For read-write endpoints, this is a read-write hostname which connects
  to the primary compute. For read-only endpoints, this is a read-only hostname which allows read-only operations

### EndpointStatus
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `current_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `hosts` (EndpointHosts) - Contains host information for connecting to the endpoint
* `pending_state` (string) - Possible values are: `ACTIVE`, `IDLE`, `INIT`
* `settings` (EndpointSettings)
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended

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