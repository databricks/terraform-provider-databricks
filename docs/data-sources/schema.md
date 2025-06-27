---
subcategory: "Unity Catalog"
---
# databricks_schema Data Source

Retrieves details about [databricks_schema](../resources/schema.md) that was created by Terraform or manually. 
A schema can be identified by its two-level (fully qualified) name (in the form of: `catalog_name`.`schema_name`) as input. This can be retrieved programmatically using [databricks_schemas](../data-sources/schemas.md) data source.

-> This data source can only be used with a workspace-level provider!

## Example Usage

* Retrieve details of all schemas in in a _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_schemas" "all" {
  catalog_name = "sandbox"
}

data "databricks_schema" "this" {
  for_each = data.databricks_schemas.all.ids
  name     = each.value
}
```

* Search for a specific schema by its fully qualified name:

```hcl
data "databricks_schema" "this" {
  name = "catalog.schema"
}
```

## Argument Reference

* `name` - (Required) a fully qualified name of [databricks_schema](../resources/schema.md): *`catalog`.`schema`*


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this Unity Catalog Schema in form of `<catalog>.<schema>`.
* `schema_info` - `SchemaInfo` object for a Unity Catalog schema. This contains the following attributes:
  * `browse_only` - indicates whether the principal is limited to retrieving metadata for the schema through the BROWSE privilege. 
  * `catalog_name` - the name of the catalog where the schema is.
  * `catalog_type` - the type of the parent catalog.
  * `comment` - the comment attached to the volume
  * `created_at` - time at which this schema was created, in epoch milliseconds.
  * `created_by` - username of schema creator.
  * `effective_predictive_optimization_flag` - information about actual state of predictive optimization. 
  * `enable_predictive_optimization` - whether predictive optimization should be enabled for this object and objects under it.
  * `full_name` - the two-level (fully qualified) name of the schema
  * `metastore_id` - the unique identifier of the metastore
  * `name` - Name of schema, relative to parent catalog.
  * `owner` - the identifier of the user who owns the schema
  * `properties` - map of properties set on the schema
  * `schema_id` - the unique identifier of the schema
  * `storage_location` - the storage location on the cloud.
  * `storage_root` - storage root URL for managed tables within schema.
  * `updated_at` - the timestamp of the last time changes were made to the schema
  * `updated_by` - the identifier of the user who updated the schema last time

## Related Resources

The following resources are used in the same context:

* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
