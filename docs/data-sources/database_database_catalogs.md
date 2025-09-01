---
subcategory: "Database Instances"
---
# databricks_database_database_catalogs Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `instance_name` (string, required) - Name of the instance to get database catalogs for
* `page_size` (integer, optional) - Upper bound for items returned
* `workspace_id` (string, optional) - Workspace ID of the resource


## Attributes
This data source exports a single attribute, `database_catalogs`. It is a list of resources, each with the following attributes:
* `create_database_if_not_exists` (boolean)
* `database_instance_name` (string) - The name of the DatabaseInstance housing the database
* `database_name` (string) - The name of the database (in a instance) associated with the catalog
* `name` (string) - The name of the catalog in UC
* `uid` (string)