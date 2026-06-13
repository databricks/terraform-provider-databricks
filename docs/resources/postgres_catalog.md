---
subcategory: "Postgres"
---
# databricks_postgres_catalog Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres catalog registers a Postgres database in Unity Catalog. This enables Lakehouse Federation queries against the Postgres data and is a prerequisite for creating synced tables.

### Hierarchy Context

Catalogs exist within the Lakebase Autoscaling resource hierarchy:
- A **catalog** belongs to a **branch** within a **project**
- A catalog registers exactly one Postgres database in Unity Catalog
- A Postgres database can only be registered with one catalog at a time
- **Synced tables** are created inside a catalog to replicate Delta tables into Postgres

### Use Cases

- **Unity Catalog governance**: Register Postgres databases in Unity Catalog for centralized access control and auditing
- **Lakehouse Federation**: Query Postgres data directly from Databricks SQL and notebooks via Unity Catalog
- **Synced table prerequisite**: Catalogs are required before creating synced tables that replicate Delta data into Postgres


## Example Usage
### Basic Catalog Creation

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

resource "databricks_postgres_catalog" "this" {
  catalog_id = "my_catalog"
  spec = {
    postgres_database          = "my_database"
    create_database_if_missing = true
    branch                     = databricks_postgres_branch.main.name
  }
}
```

### Catalog with Existing Database

If the Postgres database already exists, omit `create_database_if_missing`:

```hcl
resource "databricks_postgres_catalog" "this" {
  catalog_id = "existing_db_catalog"
  spec = {
    postgres_database = "existing_database"
    branch            = databricks_postgres_branch.main.name
  }
}
```


## Arguments
The following arguments are supported:
* `catalog_id` (string, required) - The part of the name, chosen by the user when the resource was created
* `spec` (CatalogCatalogSpec, optional) - The desired state of the Catalog
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### CatalogCatalogSpec
* `postgres_database` (string, required) - The name of the Postgres database inside the specified Lakebase project and branch to be associated with the UC catalog.
  This database must already exist, unless create_database_if_missing is set to true on creation.
  
  A database can only be registered with one UC catalog at a time.
  To re-register a database with a different catalog, the existing catalog must be deleted first.
  
  A child branch inherits the fact of parent's registration. This means the same-named database
  in a child branch cannot be registered with a second catalog
  while the parent's registration exists. To allow registering the database of a child branch,
  drop and recreate the database on the child branch.
  This removes the fact of parent's registration from this branch only.
  
  Doing Point In Time Restore (PITR) prior to the moment before the Postgres DB was registered
  in the Catalog drops the fact of registration of the database. So the user should avoid doing so
* `branch` (string, optional) - The resource path of the branch associated with the catalog.
  
  Format: projects/{project_id}/branches/{branch_id}
* `create_database_if_missing` (boolean, optional) - If set to true, the specified postgres_database is created on behalf of the calling user
  if it does not already exist. In this case, the calling user has a role created for
  them in Postgres if they do not already have one.
  
  Defaults to false, meaning that the request fails if the specified postgres_database does not already exist

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the catalog was created
* `name` (string) - Output only. The full resource path of the catalog.
  
  Format: "catalogs/{catalog_id}"
* `status` (CatalogCatalogStatus) - The observed state of the Catalog
* `uid` (string) - System-generated unique identifier for the catalog
* `update_time` (string) - A timestamp indicating when the catalog was last updated

### CatalogCatalogStatus
* `branch` (string) - The resource path of the branch associated with the catalog.
  
  Format: projects/{project_id}/branches/{branch_id}
* `postgres_database` (string) - The name of the Postgres database associated with the catalog
* `project` (string) - The resource path of the project associated with the catalog.
  
  Format: projects/{project_id}

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_catalog.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_catalog.this "name"
```