---
subcategory: "Unity Catalog"
---
# databricks_external_location Resource

-> **Private Preview** This feature is in [Private Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Unity Catalog’s security and governance model provides excellent support for External Tables. This can then be used to define external paths/tables where you can control access. 

## Example Usage

```hcl
resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_external_location" "some" {
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE TABLE", "READ FILES"]
  }
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of External Location, which must be unique within the [databricks_metastore](metastore.md). Change forces new resource.
* `url` - Path URL in cloud storage, of the form: `s3://bucket-host/[bucket-dir]` (AWS), `abfss://[user@]host/[path]` (Azure).
* `credential_name` - Name of the [databricks_storage_credential](storage_credential.md) to use with this External Location.
* `owner` - (Optional) Username/groupname of External Location owner. Currently this field can only be changed after the resource is created.
* `comment` - (Optional) User-supplied free-form text.

## Import

This resource can be imported via name:

```bash
$ terraform import databricks_external_location.this <name>
```
