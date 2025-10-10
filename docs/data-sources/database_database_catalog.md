---
subcategory: "Database Instances"
---
# databricks_database_database_catalog Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single Database Catalog.


## Example Usage
Referring to a Database Catalog by name:

```hcl
data "databricks_database_database_catalog" "this" {
  name = "my-database-catalog"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the catalog in UC

## Attributes
The following attributes are exported:
* `create_database_if_not_exists` (boolean)
* `database_instance_name` (string) - The name of the DatabaseInstance housing the database
* `database_name` (string) - The name of the database (in a instance) associated with the catalog
* `name` (string) - The name of the catalog in UC
* `uid` (string)