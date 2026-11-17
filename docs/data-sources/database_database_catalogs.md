---
subcategory: "Database Instances"
---
# databricks_database_database_catalogs Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/database)



## Example Usage


## Arguments
The following arguments are supported:
* `instance_name` (string, required) - Name of the instance to get database catalogs for
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `database_catalogs`. It is a list of resources, each with the following attributes:
* `create_database_if_not_exists` (boolean) - If true, the database named in database_name is created in the instance when it does not already
  exist, with the calling user as its owner. Defaults to false
* `database_instance_name` (string) - The name of the DatabaseInstance housing the database
* `database_name` (string) - The name of the Postgres database in the instance that the catalog registers. The database must
  already exist unless create_database_if_not_exists is true. The system databases postgres,
  databricks_system, template0 and template1 cannot be registered
* `name` (string) - The name of the catalog in Unity Catalog
* `uid` (string) - An immutable UUID identifier for the catalog, assigned when the catalog is created