---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignment Data Source
This data source allows you to get information about a tag assignment for a specific entity using the entity type, entity name, and tag key.

## Example Usage
### Get environment tag from a catalog

```hcl
data "databricks_entity_tag_assignment" "catalog_tag" {
  entity_type = "catalogs"
  entity_name = "production_catalog"
  tag_key     = "environment"
}

data "databricks_entity_tag_assignment" "schema_tag" {
  entity_type = "schemas"
  entity_name = "production_catalog.analytics_data"
  tag_key     = "cost_center"
}

data "databricks_entity_tag_assignment" "table_tag" {
  entity_type = "tables"
  entity_name = "production_catalog.sales_data.customer_orders"
  tag_key     = "owner"
}

data "databricks_entity_tag_assignment" "column_tag" {
  entity_type = "columns"
  entity_name = "production_catalog.customer_data.users.email_address"
  tag_key     = "pii_classification"
}

data "databricks_entity_tag_assignment" "volume_tag" {
  entity_type = "volumes"
  entity_name = "production_catalog.raw_data.landing_zone"
  tag_key     = "purpose"
}
```

## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string, required) - The key of the tag

## Attributes
The following attributes are exported:
* `entity_name` (string) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string) - The key of the tag
* `tag_value` (string) - The value of the tag