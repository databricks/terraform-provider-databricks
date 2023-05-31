---
subcategory: "Unity Catalog"
---
# databricks_volume (Resource)
Volumes are a new Unity Catalog (UC) capability for accessing, storing, governing, and organizing files. Volumes unlock new processing capabilities for data governed by the Unity Catalog, including support for most machine learning and data science workloads. You can use volumes to store and access files in any format; data can be structured, semi-structured, or unstructured.

With Volumes, files managed centrally in Unity Catalogare are organized under a 3-level namespace: `<catalog>.<schema>.<volume>`.

This resource does Creates, Read, Update, Delete and List Unity Catalog volumes.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_storage_credential" "external" {
  name = "creds"
  aws_iam_role {
    role_arn = "role"
  }
}

resource "databricks_external_location" "some" {
  name            = "external-location"
  url             = "some-url"
  credential_name = databricks_storage_credential.external.id
}

resource "databricks_volume" "this" {
  name = "quickstart_volume"
  catalog_name = databricks_catalog.sandbox.name
  schema_name = databricks_schema.things.name 
  owner = "volume_owner"
  volume_type = "EXTERNAL"
  storage_location   = databricks_external_location.some.url 
  comment = "this volume is managed by terraform"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of volume
* `catalog_name` - Name of parent catalog
* `schema_name` - Name of parent Schema relative to parent Catalog
* `volume_type` - URL of storage location. Currently only `EXTERNAL` is supported
* `owner` - (Optional) Name of the volume owner
* `storage_location` - (Optional) If `EXTERNAL` volume type is used, then location of that volume. 
* `comment` - (Optional) User-supplied free-form text.

## Import

This resource can be imported by `full_name` which is comprised under a 3-level namespace `<catalog>.<schema>.<volume>`

```bash
$ terraform import databricks_volume.this <catalog_name>.<schema_name>.<name>
```
