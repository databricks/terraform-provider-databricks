---
subcategory: "Postgres"
---
# databricks_postgres_databases Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source lists all Postgres databases in a branch.


## Example Usage
### List All Databases in a Branch

```hcl
data "databricks_postgres_databases" "all" {
  parent = "projects/my-project/branches/main"
}

output "database_names" {
  value = [for db in data.databricks_postgres_databases.all.databases : db.name]
}

output "postgres_database_names" {
  value = [for db in data.databricks_postgres_databases.all.databases : db.status.postgres_database]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Branch that owns this collection of databases.
  Format: projects/{project_id}/branches/{branch_id}
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `databases`. It is a list of resources, each with the following attributes:
* `create_time` (string) - A timestamp indicating when the database was created
* `database_id` (string) - The part of the name, chosen by the user when the resource was created
* `name` (string) - The resource name of the database.
  Format: projects/{project_id}/branches/{branch_id}/databases/{database_id}
* `parent` (string) - The branch containing this database.
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (DatabaseDatabaseSpec) - The desired state of the Database
* `status` (DatabaseDatabaseStatus) - The observed state of the Database
* `update_time` (string) - A timestamp indicating when the database was last updated

### DatabaseDatabaseSpec
* `postgres_database` (string) - The name of the Postgres database.
  
  This expects a valid Postgres identifier as specified in the link below.
  https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
  Required when creating the Database.
  
  To rename, pass a valid postgres identifier when updating the Database
* `role` (string) - The name of the role that owns the database.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
  
  To change the owner, pass valid existing Role name when updating the Database
  
  A database always has an owner

### DatabaseDatabaseStatus
* `database_id` (string) - Part of the resource name
* `postgres_database` (string) - The name of the Postgres database
* `role` (string) - The name of the role that owns the database.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}