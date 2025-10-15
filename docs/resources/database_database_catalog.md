---
subcategory: "Database Instances"
---
# databricks_database_database_catalog Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Database Catalogs are databases inside a Lakebase Database Instance which are synced into a Postgres Catalog inside Unity Catalog.

### Use Cases

Database Catalogs can be used to achieve governance of Postgres databases in Unity Catalog. It also allows access to underlying data via REST APIs and Feature Stores.

### Permissions

Users with CAN_USE permission on a Database Instance can create Database Catalogs.

## Example Usage
### Example

This example creates a Database Catalog based on an existing database in the Database Instance
```hcl
resource "databricks_database_database_catalog" "this" {
  name = "my_registered_catalog"
  database_instance_name = "my-database-instance"
  database_name = "databricks_postgres"
}
```

This example creates a Database Catalog along with a new database inside an existing Database Instance
```hcl
resource "databricks_database_database_catalog" "this" {
  name = "my_registered_catalog"
  database_instance_name = "my-database-instance"
  database_name = "new_registered_catalog_database"
  create_database_if_not_exists = true
}
```

This example creates a DatabaseInstance and then a Database Catalog inside it
```hcl
resource "databricks_database_instance" "instance" {
  name = "my-database-instance"
  capacity = "CU_1"
}

resource "databricks_database_database_catalog" "catalog" {
  name = "my_registered_catalog"
  database_instance_name = databricks_database_instance.instance.name
  database_name = "new_registered_catalog_database"
  create_database_if_not_exists = true
}
```

## Arguments
The following arguments are supported:
* `database_instance_name` (string, required) - The name of the DatabaseInstance housing the database
* `database_name` (string, required) - The name of the database (in a instance) associated with the catalog
* `name` (string, required) - The name of the catalog in UC
* `create_database_if_not_exists` (boolean, optional)
* `database_branch_id` (string, optional) - The branch_id of the database branch associated with the catalog
* `database_project_id` (string, optional) - The project_id of the database project associated with the catalog
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

## Attributes
In addition to the above arguments, the following attributes are exported:
* `uid` (string)

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_database_database_catalog.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_database_database_catalog "name"
```