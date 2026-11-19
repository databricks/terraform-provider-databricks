---
subcategory: "Database Instances"
---
# databricks_database_database_catalog Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/database)

This data source can be used to get a single Database Catalog.

For new workloads, Databricks recommends the `databricks_postgres_catalog` data source.


## Example Usage
Referring to a Database Catalog by name:

```hcl
data "databricks_database_database_catalog" "this" {
  name = "my-database-catalog"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the catalog in Unity Catalog
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_database_if_not_exists` (boolean) - If true, the database named in database_name is created in the instance when it does not already
  exist, with the calling user as its owner. Defaults to false
* `database_instance_name` (string) - The name of the DatabaseInstance housing the database
* `database_name` (string) - The name of the Postgres database in the instance that the catalog registers. The database must
  already exist unless create_database_if_not_exists is true. The system databases postgres,
  databricks_system, template0 and template1 cannot be registered
* `name` (string) - The name of the catalog in Unity Catalog
* `uid` (string) - An immutable UUID identifier for the catalog, assigned when the catalog is created