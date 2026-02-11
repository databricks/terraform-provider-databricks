---
subcategory: "Postgres"
---
# databricks_postgres_endpoint Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres endpoint is a virtualized service that runs Postgres for your Lakebase Autoscaling projects. Each branch has one primary (read-write) endpoint. An endpoint is required to connect to a branch and access its data.

### Hierarchy Context

Endpoints exist within the Lakebase Autoscaling resource hierarchy:
- Each **endpoint** belongs to a **branch**
- The **branch** provides the data, while the **endpoint** provides the compute to access it
- Multiple endpoints can access the same branch data simultaneously

### Use Cases

- **Read-write primary compute**: Create a read-write endpoint as the primary compute for database operations
- **Read replicas**: Create read-only endpoints to horizontally scale read operations
- **Workload isolation**: Use separate endpoints for different applications or teams accessing the same branch
- **Cost optimization**: Configure autoscaling and suspension to match your workload patterns


## Example Usage
### Basic Read-Write Endpoint

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "dev" {
  branch_id = "dev-branch"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_endpoint" "primary" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type = "ENDPOINT_TYPE_READ_WRITE"
  }
}
```

### Read-Only Endpoint with Autoscaling

```hcl
resource "databricks_postgres_endpoint" "read_replica" {
  endpoint_id = "read-replica-1"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_ONLY"
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 4.0
  }
}
```

### Endpoint with Custom Autoscaling and Suspension

```hcl
resource "databricks_postgres_endpoint" "analytics" {
  endpoint_id = "analytics"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_ONLY"
    autoscaling_limit_min_cu = 1.0
    autoscaling_limit_max_cu = 8.0
    suspend_timeout_duration = "600s"  # Suspend after 10 minutes of inactivity
  }
}
```

### Disabled Endpoint

```hcl
resource "databricks_postgres_endpoint" "maintenance" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type = "ENDPOINT_TYPE_READ_WRITE"
    disabled      = true
  }
}
```

### Endpoint with No Suspension

```hcl
resource "databricks_postgres_endpoint" "always_on" {
  endpoint_id = "always-on"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    no_suspension = true  # Never suspend
  }
}
```

### Complete Example

```hcl
resource "databricks_postgres_project" "prod" {
  project_id = "production"
  spec = {
    pg_version                = 17
    display_name              = "Production Workloads"
    history_retention_duration = "2592000s"  # 30 days

    default_endpoint_settings = {
      autoscaling_limit_min_cu = 1.0
      autoscaling_limit_max_cu = 8.0
      suspend_timeout_duration = "300s"
    }
  }
}

resource "databricks_postgres_branch" "main" {
  branch_id = "main"
  parent    = databricks_postgres_project.prod.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_endpoint" "primary" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.main.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    autoscaling_limit_min_cu = 1.0
    autoscaling_limit_max_cu = 9.0
    no_suspension = true  # Never suspend
  }
}

resource "databricks_postgres_endpoint" "read_replica" {
  endpoint_id = "read-replica"
  parent      = databricks_postgres_branch.main.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_ONLY"
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 8.0
    suspend_timeout_duration = "600s"
  }
}
```


## Arguments
The following arguments are supported:
* `endpoint_id` (string, required) - The ID to use for the Endpoint. This becomes the final component of the endpoint's resource name.
  The ID is required and must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens.
  For example, `primary` becomes `projects/my-app/branches/development/endpoints/primary`
* `parent` (string, required) - The branch containing this endpoint (API resource hierarchy).
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (EndpointSpec, optional) - The spec contains the compute endpoint configuration, including autoscaling limits, suspend timeout, and disabled state
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

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
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided
* `settings` (EndpointSettings, optional)
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `name` (string) - Output only. The full resource path of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
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