---
subcategory: "Postgres"
---
# databricks_postgres_database Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source retrieves a single Postgres database.


## Example Usage
### Retrieve Database by Name

```hcl
data "databricks_postgres_database" "this" {
  name = "projects/my-project/branches/main/databases/app"
}

output "postgres_database_name" {
  value = data.databricks_postgres_database.this.status.postgres_database
}

output "database_owner_role" {
  value = data.databricks_postgres_database.this.status.role
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the database.
  Format: projects/{project_id}/branches/{branch_id}/databases/{database_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the database was created
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