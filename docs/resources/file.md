---
subcategory: "Storage"
---
# databricks_file Resource

This resource allows uploading and downloading files in [databricks_volume](volume.md) up to 2GB in octet-stream.

## Example Usage

In order to manage a file on Unity Catalog Volumes with Terraform, you must specify the `source` attribute containing the full path to the file on the local filesystem.

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

resource "databricks_file" "this" {
  source = "/full/path/on/local/system"
  path   = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/{databricks_volume.this.name}/fileName"
}
```

You can also inline sources through `content_base64`  attribute.

```hcl
resource "databricks_file" "init_script" {
  content_base64 = base64encode(<<-EOT
    #!/bin/bash
    echo "Hello World"
    EOT
  )
  path   = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/{databricks_volume.this.name}/fileName"
}
```

## Argument Reference

The following arguments are supported:

* `source` - The full absolute path to the file. Conflicts with `content_base64`.
* `content_base64` - Contents in base 64 format. Conflicts with `source`.
* `path` - The path of the file in which you wish to save. Should start with `/Volumes`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.
* `modification_time` - The last time stamp when the file was modified


## Import

The resource `databricks_file` can be imported using the path of the file:

```bash
$ terraform import databricks_file.this <path>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_workspace_file](./workspace_file.md)
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
