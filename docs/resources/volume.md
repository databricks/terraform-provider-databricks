---
subcategory: "Unity Catalog"
---
# databricks_volume (Resource)

Volumes are a new Unity Catalog (UC) capability for accessing, storing, governing, and organizing files. Volumes unlock new processing capabilities for data governed by the Unity Catalog, including support for most machine learning and data science workloads. You can use volumes to store and access files in any format; data can be structured, semi-structured, or unstructured.

Volumes are organized under a 3-level namespace: `<catalog>.<schema>.<volume>`.

This resource manages Volumes in Unity Catalog.

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
  name            = "external-location"
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
* `catalog_name` - Name of parent Catalog
* `schema_name` - Name of parent Schema relative to parent Catalog
* `volume_type` - Volume type. `EXTERNAL` or `MANAGED`.
* `owner` - (Optional) Name of the volume owner.
* `storage_location` - (Optional) Path inside an External Location. Only used for `EXTERNAL` Volumes.
* `comment` - (Optional) Free-form text.

## Import

This resource can be imported by `full_name` which is the 3-level Volume identifier: `<catalog>.<schema>.<volume>`

```bash
terraform import databricks_volume.this <catalog_name>.<schema_name>.<name>
```
