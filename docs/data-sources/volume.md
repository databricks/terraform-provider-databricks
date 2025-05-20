---
subcategory: "Unity Catalog"
---
# databricks_volume Data Source

Retrieves details about [databricks_volume](../resources/volume.md) that was created by Terraform or manually. 
A volume can be identified by its three-level (fully qualified) name (in the form of: `catalog_name`.`schema_name`.`volume_name`) as input. This can be retrieved programmatically using [databricks_volumes](../data-sources/volumes.md) data source.

-> This data source can only be used with a workspace-level provider!

## Example Usage

* Retrieve details of all volumes in in a _things_ [databricks_schema](../resources/schema.md) of a  _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_volumes" "all" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

data "databricks_volume" "this" {
  for_each = data.databricks_volumes.all.ids
  name     = each.value
}
```

* Search for a specific volume by its fully qualified name

```hcl
data "databricks_volume" "this" {
  name = "catalog.schema.volume"
}
```

## Argument Reference

* `name` - (Required) a fully qualified name of [databricks_volume](../resources/volume.md): *`catalog`.`schema`.`volume`*


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this Unity Catalog Volume in form of `<catalog>.<schema>.<name>`.
* `volume_info` - `VolumeInfo` object for a Unity Catalog volume. This contains the following attributes:
  * `name` - Name of the volume, relative to parent schema.
  * `access_point` - the AWS access point to use when accessing s3 bucket for this volume's external location
  * `browse_only` - indicates whether the principal is limited to retrieving metadata for the volume through the BROWSE privilege when include_browse is enabled in the request. 
  * `catalog_name` - the name of the catalog where the schema and the volume are
  * `comment` - the comment attached to the volume
  * `created_at` - the Unix timestamp at the volume's creation
  * `created_by` - the identifier of the user who created the volume
  * `encryption_details` - encryption options that apply to clients connecting to cloud storage
  * `full_name` - the three-level (fully qualified) name of the volume
  * `metastore_id` - the unique identifier of the metastore
  * `name` - the name of the volume
  * `owner` - the identifier of the user who owns the volume
  * `schema_name` - the name of the schema where the volume is
  * `storage_location` - the storage location on the cloud
  * `updated_at` - the timestamp of the last time changes were made to the volume
  * `updated_by` - the identifier of the user who updated the volume last time
  * `volume_id` - the unique identifier of the volume
  * `volume_type` - whether the volume is `MANAGED` or `EXTERNAL`

## Related Resources

The following resources are used in the same context:

* [databricks_volume](../resources/volume.md) to manage volumes within Unity Catalog.
* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
