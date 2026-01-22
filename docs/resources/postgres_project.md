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
  The ID must be 1-63 characters long, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens (RFC 1123).
  Examples:
  - With custom ID: `production` → name becomes `projects/production`
  - Without custom ID: system generates UUID → name becomes `projects/a7f89b2c-3d4e-5f6g-7h8i-9j0k1l2m3n4o`
* `spec` (ProjectSpec, optional) - The spec contains the project configuration, including display_name, pg_version (Postgres version), history_retention_duration, and default_endpoint_settings

### ProjectDefaultEndpointSettings
* `autoscaling_limit_max_cu` (number, optional) - The maximum number of Compute Units. Minimum value is 0.5
* `autoscaling_limit_min_cu` (number, optional) - The minimum number of Compute Units. Minimum value is 0.5
* `no_suspension` (boolean, optional) - When set to true, explicitly disables automatic suspension (never suspend).
  Should be set to true when provided
* `pg_settings` (object, optional) - A raw representation of Postgres settings
* `suspend_timeout_duration` (string, optional) - Duration of inactivity after which the compute endpoint is automatically suspended.
  If specified should be between 60s and 604800s (1 minute to 1 week)

### ProjectSpec
* `default_endpoint_settings` (ProjectDefaultEndpointSettings, optional)
* `display_name` (string, optional) - Human-readable project name. Length should be between 1 and 256 characters
* `history_retention_duration` (string, optional) - The number of seconds to retain the shared history for point in time recovery for all branches in this project. Value should be between 0s and 2592000s (up to 30 days)
* `pg_version` (integer, optional) - The major Postgres version number. Supported versions are 16 and 17

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the project was created
* `name` (string) - The resource name of the project. This field is output-only and constructed by the system.
  Format: `projects/{project_id}`
* `status` (ProjectStatus) - The current status of a Project
* `uid` (string) - System-generated unique ID for the project
* `update_time` (string) - A timestamp indicating when the project was last updated

### ProjectStatus
* `branch_logical_size_limit_bytes` (integer) - The logical size limit for a branch
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