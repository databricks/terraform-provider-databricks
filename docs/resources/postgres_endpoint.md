---
subcategory: "Postgres"
---
# databricks_postgres_endpoint Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

These resources use a spec/status model: the `spec` field defines the resource configuration Terraform manages, and `status` reports the resource's current state.

Provider versions without drift detection do not detect changes made through the UI, CLI, or API. They apply updates based on changes to your Terraform configuration.

With a provider version that supports drift detection, `terraform plan` detects changes to managed `spec` fields made outside Terraform.

For help resolving provider inconsistency errors or recovering from a partially completed apply, see [Terraform state and drift detection for Lakebase](https://docs.databricks.com/aws/en/oltp/projects/terraform-state-and-drift). Use shared Terraform state storage when collaborating on the same resources.

### Upgrading to a Provider with Drift Detection

When upgrading from a provider version without drift detection to one that supports it, follow the steps below to resolve initial or persistent diffs.

Prefer a complete `spec` that describes the resource's full desired configuration. The examples below show the minimal valid specs for an existing resource, not the recommended end state.

#### Replace a Null Spec

Null specs can produce persistent differences against the object returned during refresh; repeated applies do not resolve those differences.

For an existing Endpoint, replace `spec = null` with a spec containing its current `endpoint_type`. For example, for a read-only Endpoint:

```hcl
spec = {
  endpoint_type = "ENDPOINT_TYPE_READ_ONLY"
}
```

Use `ENDPOINT_TYPE_READ_WRITE` for a read-write Endpoint. The type is required inside a non-null spec, is immutable, and must match the existing Endpoint. Review and apply the plan, then run `terraform plan` again; keeping the type in the spec avoids repeatedly planning to remove its value.

#### Reconcile an Initial Diff with an Existing Spec

An Endpoint created before Lakebase supported drift detection can show an initial in-place update even when its current values already match your configured spec. Review the plan against your intended configuration and retain the Endpoint's current `endpoint_type`. Apply the reviewed plan, then run `terraform plan` again.

#### Resolve a Persistent Diff

If differences remain after applying, follow the [shared reconciliation steps](https://docs.databricks.com/aws/en/oltp/projects/terraform-state-and-drift#reconcile-existing-resources-after-rollout) to choose between the values in your configuration and the current values returned by Lakebase, then review the resulting plan.

### Optional Fields and Durations

Prefer `null` for unused optional maps and nested objects, such as `settings` and `group`. Keep the root `spec` non-null and include `endpoint_type`.

Use whole-second durations for `suspend_timeout_duration`, such as `"300s"`. Fractional durations can cause a provider inconsistency error after apply.

### Clearing PostgreSQL Settings

Terraform currently does not support clearing the entire `settings.pg_settings` map. Removing it from `spec`, setting it to `null`, or using an empty map preserves existing PostgreSQL settings.

### Import limitations

See [Import](#import) for import commands.

Include the Endpoint's current `endpoint_type` in its spec when importing. This field is required and immutable.

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
### Managing Implicitly Created Read-Write Endpoint

A read-write endpoint named `primary` is implicitly created for every branch. Since Terraform is declarative, managing an already-existing resource requires `replace_existing = true`: it lets Terraform represent the implicitly created endpoint in Terraform state and immediately apply the provided configuration to it. Support for providing a custom `endpoint_id` will be available in later versions.

For this implicitly created read-write Endpoint, Terraform manages updates only. It does not control creation or deletion of the endpoint itself. Removing the resource from your Terraform configuration only removes it from Terraform state; the actual endpoint is unaffected, because its lifecycle is controlled by the parent branch. The only way to remove the actual endpoint is to delete the branch it belongs to. If you don't want to delete the parent branch and are concerned about the cost, use the `disabled` or `suspend_timeout_duration` fields in `spec`.

```hcl
resource "databricks_postgres_project" "this" {
  project_id      = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }

  lifecycle {
    prevent_destroy = true
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
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 4.0
    suspend_timeout_duration = "600s"
  }
  replace_existing = true
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
    suspend_timeout_duration = "600s" # Suspend after 10 minutes of inactivity
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
  replace_existing = true # for managing implicitly created read-write endpoint
}
```

### Endpoint with No Suspension

```hcl
resource "databricks_postgres_endpoint" "always_on" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type = "ENDPOINT_TYPE_READ_WRITE"
    no_suspension = true # Never suspend
  }
  replace_existing = true # for managing implicitly created read-write endpoint
}
```

### High Availability Endpoint

Configure a single endpoint with multiple compute instances for high availability.
One compute instance acts as the read-write primary, while the remaining secondary compute instances stand ready for automatic failover.

High availability requires scale-to-zero to be disabled.
Set `no_suspension = true` in `spec` as shown in the example below.

```hcl
resource "databricks_postgres_endpoint" "ha_primary" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    no_suspension            = true
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 4.0
    group = {
      min = 2
      max = 2
    }
  }
  replace_existing = true # for managing implicitly created read-write endpoint
}
```

### High Availability Endpoint with Readable Secondaries

Enable readable secondaries to offload read traffic to replica computes via a
dedicated read-only host, in addition to hot-standby failover. Only supported
on read-write endpoints with more than one compute. The secondaries are optionally 
exposed as read-only host via `enable_readable_secondaries`.

```hcl
resource "databricks_postgres_endpoint" "ha_readable" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.dev.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    no_suspension            = true
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 4.0
    group = {
      min                         = 2
      max                         = 2
      enable_readable_secondaries = true
    }
  }
  replace_existing = true # for managing implicitly created read-write endpoint
}
```

### Complete Example

```hcl
resource "databricks_postgres_project" "prod" {
  project_id      = "production"
  spec = {
    pg_version                 = 17
    display_name               = "Production Workloads"
    history_retention_duration = "2592000s" # 30 days

    default_endpoint_settings = {
      autoscaling_limit_min_cu = 1.0
      autoscaling_limit_max_cu = 8.0
      suspend_timeout_duration = "86400s" # 24 hours
    }
  }

  lifecycle {
    prevent_destroy = true
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
    no_suspension            = true # Never suspend
    group = {
      min                         = 2
      max                         = 2
      enable_readable_secondaries = true
    }
  }
  replace_existing = true
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
* `endpoint_id` (string, required) - The endpoint ID chosen when the endpoint was created. It is the final segment of `name`
* `parent` (string, required) - The branch containing this endpoint (API resource hierarchy).
  Format: projects/{project_id}/branches/{branch_id}
* `replace_existing` (boolean, optional) - If true, update the endpoint if it already exists instead of returning an error. The request's
  `spec.endpoint_type` must match the existing endpoint's type, and the other `spec` fields are applied
  as an update. Use this to configure the read-write endpoint that was created together with the branch
* `spec` (EndpointSpec, optional) - The desired configuration of the endpoint. Every field under `spec` except `endpoint_type` can be
  changed after creation. See `status` for the configuration that is currently in effect.
  
  Read the resource's current values from `status`, rather than `spec`.
  `spec` tracks field ownership for infrastructure-as-code tools and may omit
  unmanaged fields
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### EndpointGroupSpec
* `max` (integer, required) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer, required) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1
* `enable_readable_secondaries` (boolean, optional) - Whether to allow read-only connections to read-write endpoints. Only relevant for read-write endpoints where
  `max` is greater than 1. When enabled, the endpoint also reports a read-only host in `status.hosts`

### EndpointGroupStatus
* `max` (integer, required) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer, required) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1

### EndpointSettings
* `pg_settings` (object, optional) - Postgres server parameters to set on the compute endpoint, keyed by parameter name. Each value is
  given as a string

### EndpointSpec
* `endpoint_type` (string, required) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units. The maximum value is 64.
  The difference between the minimum and maximum Compute Units (max - min) must not exceed 16.
  Set it equal to `autoscaling_limit_min_cu` for a fixed-size compute. When omitted at creation, the
  project's default endpoint settings apply
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units. Minimum value is 0.5, and it must not exceed
  `autoscaling_limit_max_cu`. When omitted at creation, the project's default endpoint settings apply
* `disabled` (boolean, optional) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action. To accept connections again, set this field back to false
* `group` (EndpointGroupSpec, optional) - Settings for optional HA configuration of the endpoint. If unspecified, the endpoint defaults
  to non HA settings, with a single compute backing the endpoint (and no readable secondaries
  for Read/Write endpoints). Automatic suspension is not available for an endpoint with more than one
  compute, so set `no_suspension` to true when configuring one
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  The value must be true when provided.
  Mutually exclusive with `suspend_timeout_duration`
* `settings` (EndpointSettings, optional) - Postgres server parameters for the compute endpoint. On update, an empty `pg_settings` map leaves
  the current parameters unchanged
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified, the value must be between 60s and 604800s (1 minute to 1 week). A suspended endpoint
  keeps its hostnames and configuration, and the next connection attempt restarts it. When neither
  this field nor `no_suspension` is set at creation, the project's default endpoint settings apply.
  Mutually exclusive with `no_suspension`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the compute endpoint was created
* `name` (string) - Output only. The full resource path of the endpoint.
  Format: projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
* `status` (EndpointStatus) - The configuration and state of the endpoint that are currently in effect, as reported by the
  system, including the hostnames to connect to
* `uid` (string) - A system-generated ID that identifies the endpoint. Unlike `name`, it is not derived from the
  user-chosen endpoint ID. The endpoint's connection hostnames are based on this ID
* `update_time` (string) - A timestamp indicating when the compute endpoint was last updated

### EndpointGroupStatus
* `enable_readable_secondaries` (boolean) - Whether read-only connections to read-write endpoints are allowed. Only relevant if the endpoint has more than
  one compute (`max` greater than 1)

### EndpointHosts
* `host` (string) - The hostname to connect to this endpoint. For read-write endpoints, this is a read-write hostname which connects
  to the primary compute. For read-only endpoints, this is a read-only hostname which allows read-only operations
* `read_only_host` (string) - An optionally defined read-only host for the endpoint, without pooling. For read-only endpoints,
  this attribute is always defined and is equivalent to host. For read-write endpoints, this attribute is defined
  if the enclosing endpoint is a group with greater than 1 computes configured, and has readable secondaries enabled
* `read_only_pooled_host` (string) - The read-only hostname of the compute endpoint, with pooling. This attribute is always defined for read-only endpoints,
  and may be defined for read-write endpoints if configured with read replicas and allow read-only connections
* `read_write_pooled_host` (string) - The read-write hostname of the compute endpoint, with pooling. This attribute is only defined for read-write endpoints

### EndpointStatus
* `autoscaling_limit_max_cu` (number) - The maximum number of Compute Units. The maximum value is 64.
  The difference between the minimum and maximum Compute Units (max - min) must not exceed 16
* `autoscaling_limit_min_cu` (number) - The minimum number of Compute Units
* `current_state` (string) - The state the compute endpoint is in now. Possible values are: `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`
* `disabled` (boolean) - Whether to restrict connections to the compute endpoint.
  Enabling this option schedules a suspend compute operation.
  A disabled compute endpoint cannot be enabled by a connection or
  console action
* `endpoint_id` (string) - The endpoint ID, which is the final segment of the endpoint's resource name
* `endpoint_type` (string) - The endpoint type. A branch can only have one READ_WRITE endpoint. Possible values are: `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`
* `group` (EndpointGroupStatus) - Details on the HA configuration of the endpoint
* `hosts` (EndpointHosts) - Contains host information for connecting to the endpoint
* `last_active_time` (string) - A timestamp indicating when the compute endpoint was last active
* `pending_state` (string) - The state the compute endpoint is transitioning to, if a transition is in progress. Possible values are: `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`
* `settings` (EndpointSettings) - The Postgres server parameters currently set on the compute endpoint
* `suspend_timeout_duration` (string) - Duration of inactivity after which the compute endpoint is automatically suspended. Absent when
  automatic suspension is disabled

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