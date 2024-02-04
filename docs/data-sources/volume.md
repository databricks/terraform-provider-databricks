---
subcategory: "Unity Catalog"
---
# databricks_volume Data Source

Retrieves details about [databricks_volume](../resources/volume.md) that was created by Terraform or manually. 
A volume can be identified by its three-level (fully qualified) name (in the form of: `catalog_name`.`schema_name`.`volume_name`) or by providing the its name, the schema and catalog names it belongs to, as inputs. This former could be retrieved programmatically using [databricks_volumes](../data-sources/volumes.md) data source.

## Example Usage

* Retrieve details of all volumes in in a _things_ [databricks_schema](../resources/schema.md) of a  _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_volumes" "all" {
  catalog_name = "sandbox"
  schema_name = "things"
}

data "databricks_volume" {
  for_each = data.datatbricks_volumes.all.ids
  full_name = each.value
}
```

* Search for a specific volume by name

```hcl
data "databricks_volume "this" {
  catalog_name = "sandbox"
  schema_name = "things"
  name = "volume"
}
```

## Argument Reference

* `full_name` - (Required if `catalog_name` is not specified) a full name of [databricks_volume](../resources/volume.md): *`catalog`.`schema`.`volume`*
* `catalog_name` - (Required if `full_name` is not specified) a name of [databricks_catalog](../resources/catalog.md) to which the target volume belongs
* `schema_name` - (Required if `catalog_name` is specified) a name of [databricks_schema](../resources/schema.md) to which the target volume belongs
* `name` - (Required if `catalog_name` is specified) a name of [databricks_volume](../resources/volume.md) 

## Attribute Reference

This data source exports the following attributes:

* `access_point` - the AWS access point to use when accessing s3 bucket for this volume's external location
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
