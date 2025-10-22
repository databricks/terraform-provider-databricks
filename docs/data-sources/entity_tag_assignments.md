---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignments Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source allows you to retrieve tag assignments that have been applied to a particular entity in Unity Catalog.

## Example Usage
### Get all tag assignments for a catalog

```hcl
data "databricks_entity_tag_assignments" "catalog_tags" {
  entity_type = "catalogs"
  entity_name = "production_catalog"
}

data "databricks_entity_tag_assignments" "schema_tags" {
  entity_type = "schemas"
  entity_name = "production_catalog.sales_data"
}

data "databricks_entity_tag_assignments" "table_tags" {
  entity_type = "tables"
  entity_name = "production_catalog.sales_data.customer_orders"
}

data "databricks_entity_tag_assignments" "column_tags" {
  entity_type = "columns"
  entity_name = "production_catalog.customer_data.users.email_address"
}

data "databricks_entity_tag_assignments" "volume_tags" {
  entity_type = "volumes"
  entity_name = "production_catalog.raw_data.landing_zone"
}
```

## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `max_results` (integer, optional) - Optional. Maximum number of tag assignments to return in a single page


## Attributes
This data source exports a single attribute, `tag_assignments`. It is a list of resources, each with the following attributes:
* `entity_name` (string) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string) - The key of the tag
* `tag_value` (string) - The value of the tag