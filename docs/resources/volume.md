---
subcategory: "Unity Catalog"
---
# databricks_volume (Resource)

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

-> This resource can only be used with a workspace-level provider!

Volumes are Unity Catalog objects representing a logical volume of storage in a cloud object storage location. Volumes provide capabilities for accessing, storing, governing, and organizing files. While tables provide governance over tabular datasets, volumes add governance over non-tabular datasets. You can use volumes to store and access files in any format, including structured, semi-structured, and unstructured data.

A volume resides in the third layer of Unity Catalog’s three-level namespace. Volumes are siblings to tables, views, and other objects organized under a schema in Unity Catalog.

A volume can be **managed** or **external**.

A **managed volume** is a Unity Catalog-governed storage volume created within the default storage location of the containing schema. Managed volumes allow the creation of governed storage for working with files without the overhead of external locations and storage credentials. You do not need to specify a location when creating a managed volume, and all file access for data in managed volumes is through paths managed by Unity Catalog.

An **external volume** is a Unity Catalog-governed storage volume registered against a directory within an external location.

A volume can be referenced using its identifier: ```<catalogName>.<schemaName>.<volumeName>```, where:

* ```<catalogName>```: The name of the catalog containing the Volume.
* ```<schemaName>```: The name of the schema containing the Volume.
* ```<volumeName>```: The name of the Volume. It identifies the volume object.

The path to access files in volumes uses the following format:

```/Volumes/<catalog>/<schema>/<volume>/<path>/<file_name>```

Databricks also supports an optional ```dbfs:/``` scheme, so the following path also works:

```dbfs:/Volumes/<catalog>/<schema>/<volume>/<path>/<file_name>```

This resource manages Volumes in Unity Catalog.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox"
  comment = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.name
  name         = "things"
  comment      = "this schema is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_storage_credential" "external" {
  name = "creds"
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
}

resource "databricks_external_location" "some" {
  name            = "external_location"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.name
}

resource "databricks_volume" "this" {
  name             = "quickstart_volume"
  catalog_name     = databricks_catalog.sandbox.name
  schema_name      = databricks_schema.things.name
  volume_type      = "EXTERNAL"
  storage_location = databricks_external_location.some.url
  comment          = "this volume is managed by terraform"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of the Volume
* `catalog_name` - Name of parent Catalog. Change forces creation of a new resource.
* `schema_name` - Name of parent Schema relative to parent Catalog. Change forces creation of a new resource.
* `volume_type` - Volume type. `EXTERNAL` or `MANAGED`. Change forces creation of a new resource.
* `owner` - (Optional) Name of the volume owner.
* `storage_location` - (Optional) Path inside an External Location. Only used for `EXTERNAL` Volumes. Change forces creation of a new resource.
* `comment` - (Optional) Free-form text.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this Unity Catalog Volume in form of `<catalog>.<schema>.<name>`.
* `volume_path` - base file path for this Unity Catalog Volume in form of `/Volumes/<catalog>/<schema>/<name>`.

## Import

This resource can be imported by `full_name` which is the 3-level Volume identifier: `<catalog>.<schema>.<name>`

```bash
terraform import databricks_volume.this <catalog_name>.<schema_name>.<name>
```
