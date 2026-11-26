---
subcategory: "Postgres"
---
# databricks_postgres_project Resource
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

For an existing Project, replace `spec = null` with an empty object:

```hcl
spec = {}
```

This resolves the persistent null/object difference without changing the Project's configuration. Review and apply the plan, then run `terraform plan` again.

#### Reconcile an Initial Diff with an Existing Spec

A Project created before Lakebase supported drift detection can show an initial in-place update even when its current values already match your configured spec. Review the plan against your intended configuration. Apply the reviewed plan, then run `terraform plan` again.

#### Resolve a Persistent Diff

If differences remain after applying, follow the [shared reconciliation steps](https://docs.databricks.com/aws/en/oltp/projects/terraform-state-and-drift#reconcile-existing-resources-after-rollout) to choose between the values in your configuration and the current values returned by Lakebase, then review the resulting plan.

### Optional Fields and Durations

Prefer `null` for unused optional maps, lists, and nested objects. This recommendation does not apply to the root `spec`, which should remain non-null.

Use whole-second durations, such as `"172800s"` for `history_retention_duration` and `"300s"` for endpoint suspend timeouts. Fractional durations can cause a provider inconsistency error after apply.

### Clearing Fields

Terraform currently does not support clearing the following Project fields:

- `display_name`
- `default_endpoint_settings.pg_settings`
- `custom_tags`
- `budget_policy_id`

Removing these fields from `spec` or setting them to `null` preserves their current values. Empty maps or lists do not clear PostgreSQL settings or custom tags.

### Import limitations

See [Import](#import) for import commands.

Configuring `pg_version` on an imported Project is currently unsupported. Omit it from the imported resource's spec, even when it matches the existing PostgreSQL version.

For example:

```hcl
spec = {
  display_name = "My Project"
}
```

### Protecting Projects from Accidental Deletion

Use `prevent_destroy = true` to protect against accidental destruction or replacement.

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version = 17
  }

  lifecycle {
    prevent_destroy = true
  }
}
```

By default, deleting a Project soft-deletes it, retaining its data and preventing reuse of its `project_id` during the retention period. To recreate a Project with the same `project_id` after deletion, set `purge_on_delete = true` when creating the Project. With this option, deletion permanently removes the Project and its data.

For intentional deletion or recreation, remove `prevent_destroy` and review the plan.

> **Important:** Removing the entire resource block also removes its `prevent_destroy` protection.

### Overview

A Postgres project is the top-level container for your Lakebase Autoscaling resources, including branches, endpoints, databases, and roles.

### Hierarchy Context

Projects are the root of the Lakebase Autoscaling resource hierarchy:
- A **project** contains one or more **branches**
- Each **branch** contains **endpoints**, **databases**, and **roles**
- **Endpoints** provide compute resources for accessing branch data

### Use Cases

- **Multi-tenancy**: Isolate different applications or customers in separate projects
- **Application isolation**: Separate production, staging, and development workloads
- **Environment separation**: Create distinct environments with their own configuration and resource limits


## Example Usage
### Basic Project Creation

```hcl
resource "databricks_postgres_project" "this" {
  project_id      = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Application Project"
  }

  lifecycle {
    prevent_destroy = true
  }
}
```

### Project with Custom Configuration

```hcl
resource "databricks_postgres_project" "this" {
  project_id      = "analytics-project"
  spec = {
    pg_version                 = 16
    display_name               = "Analytics Workloads"
    history_retention_duration = "1209600s" # 14 days

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
```

### Project with High Availability Endpoint

Create a project whose read-write endpoint is configured with multiple compute instances for high availability. 
One compute instance acts as the read-write primary. 
The remaining secondary compute instances are ready for automatic failover if the primary becomes unavailable.

This example manages implicitly created resources for the root branch and read-write endpoint. For more details, see the documentation for `databricks_postgres_branch` and `databricks_postgres_endpoint`.

```hcl
resource "databricks_postgres_project" "ha" {
  project_id      = "ha-project"
  spec = {
    pg_version   = 17
    display_name = "HA Production Project"
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "databricks_postgres_branch" "production" {
  branch_id = "production"
  parent    = databricks_postgres_project.ha.name
  spec = {
    no_expiry    = true
    is_protected = true
  }
  replace_existing = true
}

resource "databricks_postgres_endpoint" "primary" {
  endpoint_id = "primary"
  parent      = databricks_postgres_branch.production.name
  spec = {
    endpoint_type            = "ENDPOINT_TYPE_READ_WRITE"
    no_suspension            = true
    autoscaling_limit_min_cu = 0.5
    autoscaling_limit_max_cu = 4.0
    group = {
      min                         = 2
      max                         = 2
      enable_readable_secondaries = false
    }
  }
  replace_existing = true
}
```

### Referencing in Other Resources

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
```


## Arguments
The following arguments are supported:
* `project_id` (string, required) - The project ID chosen when the project was created. It is the final segment of `name`
* `initial_branch_spec` (InitialBranchSpec, optional) - Configuration for the initial default branch created as part of project creation.
  Allows overriding branch protection. These settings only apply at creation time
  and do not affect resources created after project creation
* `initial_endpoint_spec` (InitialEndpointSpec, optional) - Configuration settings for the initial Read/Write endpoint created inside the initial branch for a newly
  created project. If omitted, the initial endpoint created will have default settings, without high availability
  configured. This field does not apply to any endpoints created after project creation. Use
  spec.default_endpoint_settings to configure default settings for endpoints created after project creation
* `purge_on_delete` (boolean, optional) - If true, permanently deletes the project and all of its data immediately. If false or unset, the
  project is soft-deleted: it stops accepting connections and is retained for 7 days, during which
  it can be restored with Undelete, and is then permanently deleted. Deleting an already
  soft-deleted project with `purge` set to true removes it immediately. See
  https://docs.databricks.com/aws/en/oltp/projects/manage-projects for details
* `spec` (ProjectSpec, optional) - The desired configuration of the project. The fields under `spec` are the ones that can be
  changed after creation, except `pg_version`, which is fixed when the project is created. See
  `status` for the configuration that is currently in effect.
  
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

### InitialBranchSpec
* `is_protected` (boolean, optional) - Whether the initial default branch should be protected from deletion

### InitialEndpointSpec
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units for the initial endpoint. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units for the initial endpoint. Minimum value is 0.5, and it must
  not exceed autoscaling_limit_max_cu
* `group` (EndpointGroupSpec, optional) - Settings for HA configuration of the endpoint
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided.
  Mutually exclusive with `suspend_timeout_duration`
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the initial endpoint is automatically suspended.
  If specified, should be between 60s and 604800s (1 minute to 1 week).
  Mutually exclusive with `no_suspension`

### ProjectCustomTag
* `key` (string, optional) - The key of the custom tag
* `value` (string, optional) - The value of the custom tag

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units. Minimum value is 0.5
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided.
  Mutually exclusive with `suspend_timeout_duration`
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week).
  Mutually exclusive with `no_suspension`

### ProjectSpec
* `budget_policy_id` (string, optional) - The desired budget policy to associate with the project.
  See status.budget_policy_id for the policy that is actually applied to the project
* `custom_tags` (list of ProjectCustomTag, optional) - Custom tags to associate with the project, used for billing and cost tracking.
  To update tags, provide the new tag list and include "spec.custom_tags" in the update_mask.
  To clear all tags, provide an empty list and include "spec.custom_tags" in the update_mask.
  To preserve existing tags, omit this field from the update_mask (or use wildcard "*" which auto-excludes empty tags)
* `default_branch` (string, optional) - The full resource path of the default branch of the project.
  Format: projects/{project_id}/branches/{branch_id}
  At creation time, the branch ID in this path names the initial branch, which is called
  `production` when the field is omitted. On update, the path must name an existing branch of this
  project, which then becomes the default branch
* `default_endpoint_settings` (ProjectDefaultEndpointSettings, optional) - Settings applied by default to compute endpoints created in this project. The initial endpoint
  also takes these values for any setting that initial_endpoint_spec does not provide
* `display_name` (string, optional) - Human-readable project name. Length should be between 1 and 256 characters
* `enable_pg_native_login` (boolean, optional) - Whether to enable PG native password login on all endpoints in this project. Defaults to false
* `history_retention_duration` (string, optional) - The number of seconds to retain the shared history for point in time recovery for all branches in this project. Value should be between 172800s (2 days) and 3024000s (35 days).
  Defaults to 604800s (7 days) when omitted
* `pg_version` (integer, optional) - The major Postgres version number. Defaults to 17 when omitted. The set of versions that new
  projects can use may change over time

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `delete_time` (string) - A timestamp indicating when the project was soft-deleted.
  Empty if the project is not deleted, otherwise set to a timestamp in the past
* `name` (string) - Output only. The full resource path of the project.
  Format: projects/{project_id}
* `purge_time` (string) - A timestamp indicating when the project is scheduled for permanent deletion.
  Empty if the project is not deleted, otherwise set to a timestamp in the future
* `status` (ProjectStatus) - The configuration and state of the project that are currently in effect, as reported by the
  system
* `uid` (string) - A system-generated UUID that identifies the project. Unlike `name`, it is not derived from the
  user-chosen project ID
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `budget_policy_id` (string) - The budget policy that is applied to the project
* `compute_last_active_time` (string) - The most recent time when any endpoint of this project was active
* `compute_policy_id` (string) - The compute policy that is applied to the project
* `custom_tags` (list of ProjectCustomTag) - The effective custom tags associated with the project
* `default_branch` (string) - The full resource path of the default branch of the project
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `enable_pg_native_login` (boolean) - Whether to enable PG native password login on all endpoints in this project
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `owner` (string) - The email of the project owner
* `pg_version` (integer) - The effective major Postgres version number
* `project_id` (string) - The project ID, which is the final segment of the project's resource name
* `synthetic_storage_size_bytes` (integer) - The current space occupied by the project in storage

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_project.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_project.this "name"
```