---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignment Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This resource allows you to create, update, list, and delete tag assignments on Unity Catalog entities.

## Example Usage
### Basic tag assignment to a catalog

```hcl
resource "databricks_entity_tag_assignment" "catalog_tag" {
  entity_type = "catalogs"
  entity_name = "production_catalog"
  tag_key     = "environment"
  tag_value   = "production"
}

resource "databricks_entity_tag_assignment" "schema_tag" {
  entity_type = "schemas"
  entity_name = "production_catalog.sales_data"
  tag_key     = "owner"
  tag_value   = "sales-team"
}

resource "databricks_entity_tag_assignment" "table_tag" {
  entity_type = "tables"
  entity_name = "production_catalog.sales_data.customer_orders"
  tag_key     = "data_classification"
  tag_value   = "confidential"
}

resource "databricks_entity_tag_assignment" "column_tag" {
  entity_type = "columns"
  entity_name = "production_catalog.sales_data.customers.email_address"
  tag_key     = "pii"
  tag_value   = "email"
}

resource "databricks_entity_tag_assignment" "volume_tag" {
  entity_type = "volumes"
  entity_name = "production_catalog.raw_data.landing_zone"
  tag_key     = "purpose"
  tag_value   = "data_ingestion"
}
```

## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string, required) - The key of the tag
* `tag_value` (string, optional) - The value of the tag
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

## Attributes
In addition to the above arguments, the following attributes are exported:

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "entity_type,entity_name,tag_key"
  to = databricks_entity_tag_assignment.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_entity_tag_assignment "entity_type,entity_name,tag_key"
```