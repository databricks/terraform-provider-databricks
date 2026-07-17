---
subcategory: "Postgres"
---
# databricks_postgres_cdf_config Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres CDF (Change Data Feed) configuration replicates the tables of a single Postgres schema into a Unity Catalog schema. There is one CDF configuration per Postgres schema per database. The configuration is immutable once created: changing the catalog, schema, or source Postgres schema requires replacing the resource.

### Hierarchy Context

CDF configurations exist within the Lakebase Autoscaling resource hierarchy:
- A **CDF configuration** belongs to a **database** within a **branch** and **project**
- Each configuration replicates exactly one Postgres schema (`postgres_schema`) into one Unity Catalog schema (`catalog` + `schema`)
- The `cdf_config_id` forms the final segment of the resource name; it defaults to the Postgres schema name when omitted

### Use Cases

- **Analytics on operational data**: Continuously replicate Postgres tables into Unity Catalog for lakehouse analytics without impacting the operational database
- **Change data capture**: Stream row-level changes from Postgres into Delta tables for downstream processing


## Example Usage
### Basic CDF Configuration

Replicate a Postgres schema into a Unity Catalog schema.

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "main" {
  branch_id = "main"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_database" "app" {
  database_id = "app_db"
  parent      = databricks_postgres_branch.main.name
  spec = {
    postgres_database = "app_db"
  }
}

resource "databricks_postgres_cdf_config" "this" {
  parent          = databricks_postgres_database.app.name
  catalog         = "main"
  schema          = "app_replicated"
  postgres_schema = "public"
}
```


## Arguments
The following arguments are supported:
* `catalog` (string, required) - The Unity Catalog catalog that replicated tables are written into.
  Set at creation; the CdfConfig is immutable
* `parent` (string, required) - The parent database under which to create the CdfConfig.
  Format: projects/{project}/branches/{branch}/databases/{database}
* `postgres_schema` (string, required) - The Postgres schema this CdfConfig replicates from. Unique within the
  parent database. Set at creation; the CdfConfig is immutable
* `schema` (string, required) - The Unity Catalog schema that replicated tables are written into.
  Set at creation; the CdfConfig is immutable
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `cdf_config_id` (string) - The user-specified id; equals the final segment of `name`. Defaults to the
  Postgres schema name for configs without an explicit id
* `create_time` (string) - When the CdfConfig was created
* `name` (string) - Output only. The full resource name of the CdfConfig.
  Format: projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_cdf_config.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_cdf_config.this "name"
```