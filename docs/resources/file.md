---
subcategory: "Storage"
---
# databricks_file Resource

This resource allows uploading and downloading files in [databricks_volume](volume.md).

Notes:

* Currently the limit is 5GiB in octet-stream.
* Currently, only UC volumes are supported. The list of destinations may change.

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

resource "databricks_volume" "this" {
  name         = "quickstart_volume"
  catalog_name = databricks_catalog.sandbox.name
  schema_name  = databricks_schema.things.name
  volume_type  = "MANAGED"
  comment      = "this volume is managed by terraform"
}

resource "databricks_file" "this" {
  source = "/full/path/on/local/system"
  path   = "${databricks_volume.this.volume_path}/fileName"
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
  path = "${databricks_volume.this.volume_path}/fileName"
}
```

## Argument Reference

The following arguments are supported:

* `source` - The full absolute path to the file. Conflicts with `content_base64`.
* `content_base64` - Contents in base 64 format. Conflicts with `source`.
* `path` - The path of the file in which you wish to save. For example, `/Volumes/main/default/volume1/file.txt`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.
* `modification_time` - The last time stamp when the file was modified

## Import

The resource `databricks_file` can be imported using the path of the file:

```bash
terraform import databricks_file.this <path>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_workspace_file](./workspace_file.md)
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_volume](../resources/volume.md) to manage [volumes within Unity Catalog](https://docs.databricks.com/en/connect/unity-catalog/volumes.html).
