---
subcategory: "Postgres"
---
# databricks_postgres_database Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres database is a logical database inside a Postgres branch. Each database is owned by exactly one Postgres role; that role holds the default privileges on the database's objects.

### Hierarchy Context

Databases exist within the Lakebase Autoscaling resource hierarchy:
- A **database** belongs to a **branch** within a **project**
- A **database** is owned by a **role** in the same branch
- A branch can contain multiple databases

### Use Cases

- **Application isolation**: Create one database per application so objects, schemas, and grants stay separate
- **Ownership management**: Transfer ownership between roles by updating `spec.role`
- **Renaming**: Change the underlying Postgres database name by updating `spec.postgres_database`


## Example Usage
### Database Owned by a Specific Role

Assign ownership to a role you manage alongside the database. The Postgres database will be created with the specified role as its owner.

```hcl
resource "databricks_postgres_role" "app_owner" {
  role_id = "app-owner"
  parent  = databricks_postgres_branch.main.name
  spec = {
    postgres_role = "app_owner"
  }
}

resource "databricks_postgres_database" "app" {
  database_id = "app"
  parent      = databricks_postgres_branch.main.name
  spec = {
    postgres_database = "app"
    role              = databricks_postgres_role.app_owner.name
  }
}
```

### Renaming a Database

Changing `spec.postgres_database` renames the underlying Postgres database without replacing the resource. The resource identifier (`database_id`) is separate from the Postgres database name, and stays intact in the example below.

```hcl
resource "databricks_postgres_database" "analytics" {
  database_id = "analytics"
  parent      = databricks_postgres_branch.main.name
  spec = {
    # Rename from "analytics_v1" to "analytics_v2" in Postgres by updating this field
    postgres_database = "analytics_v2"
  }
}
```

### Multiple databases in a branch

By default, Terraform creates resources in parallel if the dependency graph allows for that. However, Lakebase
doesn't allow the parallel management of resource inside a single branch. Only one of these resources can
be created at a time:

- Role
- Database
- Endpoint

If you try to create resources in parallel, you'll see a conflict error like:

> Your project already has conflicting operations in progress. Please wait until they are complete, and then try again.

Terraform serializes automatically when one resource references another, forming an edge in the dependency graph.
For example, if a database's `spec.role` points at a role, Terraform creates the role before the database.
For resources that don't reference each other, like two sibling databases in the same branch, add `depends_on` so
Terraform knows to wait for complete creation of the first resource, before starting the creation of the second one.

For example:

```hcl
resource "databricks_postgres_role" "schema_owner" {
  role_id = "schemamigrator"
  parent  = databricks_postgres_branch.test.name  # previously created branch, omitted for compactness
  spec = {
    postgres_role = "schemamigrator"
    membership_roles = ["DATABRICKS_SUPERUSER"]
  }
}

resource "databricks_postgres_database" "application1" {
  database_id = "application1"
  parent      = databricks_postgres_branch.test.name
  spec = {
    postgres_database = "application1"
    role              = databricks_postgres_role.schema_owner.name
  }
}

resource "databricks_postgres_database" "application2" {
  database_id = "application2"
  parent      = databricks_postgres_branch.test.name
  spec = {
    postgres_database = "application2"
    role              = databricks_postgres_role.schema_owner.name
  }
  
  depends_on = [ databricks_postgres_database.application1 ]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The branch containing this database.
  Format: projects/{project_id}/branches/{branch_id}
* `database_id` (string, optional) - The ID to use for the Database, which will become the final component of
  the database's resource name.
  This ID becomes the database name in postgres.
  
  This value should be 4-63 characters, and only use characters available in DNS names,
  as defined by RFC-1123
  
  If database_id is not specified in the request, it is generated automatically
* `spec` (DatabaseDatabaseSpec, optional) - The desired state of the Database
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DatabaseDatabaseSpec
* `postgres_database` (string, optional) - The name of the Postgres database.
  
  This expects a valid Postgres identifier as specified in the link below.
  https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
  Required when creating the Database.
  
  To rename, pass a valid postgres identifier when updating the Database
* `role` (string, optional) - The name of the role that owns the database.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
  
  To change the owner, pass valid existing Role name when updating the Database
  
  A database always has an owner

### DatabaseDatabaseStatus
* `postgres_database` (string, optional) - The name of the Postgres database
* `role` (string, optional) - The name of the role that owns the database.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the database was created
* `name` (string) - The resource name of the database.
  Format: projects/{project_id}/branches/{branch_id}/databases/{database_id}
* `status` (DatabaseDatabaseStatus) - The observed state of the Database
* `update_time` (string) - A timestamp indicating when the database was last updated

### DatabaseDatabaseStatus
* `database_id` (string) - The short identifier of the database, suitable for showing to the users.
  For a database with name `projects/my-project/branches/my-branch/databases/my-db`,
  the database_id is `my-db`.
  
  Use this field when building UI components that display databases to users (e.g., a drop-down
  selector). Prefer showing `database_id` instead of the full resource name from `Database.name`,
  which follows the `projects/{project_id}/branches/{branch_id}/databases/{database_id}` format
  and is not user-friendly

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_database.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_database.this "name"
```