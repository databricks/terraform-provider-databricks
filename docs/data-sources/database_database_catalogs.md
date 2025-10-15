---
subcategory: "Database Instances"
---
# databricks_database_database_catalogs Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `instance_name` (string, required) - Name of the instance to get database catalogs for
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource


## Attributes
This data source exports a single attribute, `database_catalogs`. It is a list of resources, each with the following attributes:
* `create_database_if_not_exists` (boolean)
* `database_branch_id` (string) - The branch_id of the database branch associated with the catalog
* `database_instance_name` (string) - The name of the DatabaseInstance housing the database
* `database_name` (string) - The name of the database (in a instance) associated with the catalog
* `database_project_id` (string) - The project_id of the database project associated with the catalog
* `name` (string) - The name of the catalog in UC
* `uid` (string)