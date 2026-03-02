---
subcategory: "Postgres"
---
# databricks_postgres_project Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

### Understanding Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics, which differ from typical Databricks Terraform resources:

- **No drift detection**: This implementation does not detect changes made outside Terraform. If a resource is modified manually outside of Terraform, the next `terraform apply` will not surface those differences. It only applies updates based on the Terraform configuration.

- **Global state requirement**: You must use shared state storage (e.g., an S3 bucket) to avoid accidentally overriding the resource configuration changes of other users when multiple users manage the same resources.

- **Spec vs Status behavior**:
  - The `spec` field represents your intended resource configuration - what you want the resource to be.
  - The `status` field represents the current system state - what the resource actually is.
  - Removing a field from `spec` removes your intended configuration for that field, but the server-side value persists. The server will continue using its current value unless you explicitly provide a new value in `spec`.

For branches and endpoints that reference projects, see the documentation for `databricks_postgres_branch` and `databricks_postgres_endpoint` resources.

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
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Application Project"
  }
}
```

### Project with Custom Settings

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "analytics-project"
  spec = {
    pg_version                = 16
    display_name              = "Analytics Workloads"
    history_retention_duration = "1209600s"  # 14 days

    default_endpoint_settings = {
      autoscaling_limit_min_cu    = 1.0
      autoscaling_limit_max_cu    = 8.0
      suspend_timeout_duration    = "300s"  # 5 minutes
    }
  }
}
```

### Referencing in Other Resources

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
```


## Arguments
The following arguments are supported:
* `project_id` (string, required) - The ID to use for the Project. This becomes the final component of the project's resource name.
  The ID is required and must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens.
  For example, `my-app` becomes `projects/my-app`
* `initial_endpoint_spec` (InitialEndpointSpec, optional) - Configuration settings for the initial Read/Write endpoint created inside the default branch for a newly
  created project. If omitted, the initial endpoint created will have default settings, without high availability
  configured. This field does not apply to any endpoints created after project creation. Use
  spec.default_endpoint_settings to configure default settings for endpoints created after project creation
* `spec` (ProjectSpec, optional) - The spec contains the project configuration, including display_name, pg_version (Postgres version), history_retention_duration, and default_endpoint_settings
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### EndpointGroupSpec
* `max` (integer, required) - The maximum number of computes in the endpoint group. Currently, this must be equal to min. Set to 1 for single
  compute endpoints, to disable HA. To manually suspend all computes in an endpoint group, set disabled to
  true on the EndpointSpec
* `min` (integer, required) - The minimum number of computes in the endpoint group. Currently, this must be equal to max. This must be greater
  than or equal to 1
* `enable_readable_secondaries` (boolean, optional) - Whether to allow read-only connections to read-write endpoints. Only relevant for read-write endpoints where
  size.max > 1

### InitialEndpointSpec
* `group` (EndpointGroupSpec, optional) - Settings for HA configuration of the endpoint

### ProjectCustomTag
* `key` (string, optional) - The key of the custom tag
* `value` (string, optional) - The value of the custom tag

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units. Minimum value is 0.5
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

### ProjectSpec
* `budget_policy_id` (string, optional) - The desired budget policy to associate with the project.
  See status.budget_policy_id for the policy that is actually applied to the project
* `custom_tags` (list of ProjectCustomTag, optional) - Custom tags to associate with the project. Forwarded to LBM for billing and cost tracking.
  To update tags, provide the new tag list and include "spec.custom_tags" in the update_mask.
  To clear all tags, provide an empty list and include "spec.custom_tags" in the update_mask.
  To preserve existing tags, omit this field from the update_mask (or use wildcard "*" which auto-excludes empty tags)
* `default_endpoint_settings` (ProjectDefaultEndpointSettings, optional)
* `display_name` (string, optional) - Human-readable project name. Length should be between 1 and 256 characters
* `history_retention_duration` (string, optional) - The number of seconds to retain the shared history for point in time recovery for all branches in this project. Value should be between 0s and 2592000s (up to 30 days)
* `pg_version` (integer, optional) - The major Postgres version number. Supported versions are 16 and 17

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - Output only. The full resource path of the project.
  Format: projects/{project_id}
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System-generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
* `budget_policy_id` (string) - The budget policy that is applied to the project
* `custom_tags` (list of ProjectCustomTag) - The effective custom tags associated with the project
* `default_endpoint_settings` (ProjectDefaultEndpointSettings) - The effective default endpoint settings
* `display_name` (string) - The effective human-readable project name
* `history_retention_duration` (string) - The effective number of seconds to retain the shared history for point in time recovery
* `owner` (string) - The email of the project owner
* `pg_version` (integer) - The effective major Postgres version number
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