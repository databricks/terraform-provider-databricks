---
subcategory: "Unity Catalog"
---
# databricks_volume_directory Resource

This resource allows creating and managing directories in Unity Catalog [volumes](volume.md) using the Files API.

-> This resource can only be used with a workspace-level provider!

Directories in Unity Catalog volumes provide a way to organize files and data within volumes. The Files API automatically creates parent directories as needed (similar to `mkdir -p`), making it easy to create nested directory structures.

The directory path uses the following format:

```
/Volumes/<catalog>/<schema>/<volume>/<path>
```


## Example Usage

### Basic Directory Creation

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

resource "databricks_volume" "this" {
  name         = "quickstart_volume"
  catalog_name = databricks_catalog.sandbox.name
  schema_name  = databricks_schema.things.name
  volume_type  = "MANAGED"
  comment      = "this volume is managed by terraform"
}

resource "databricks_volume_directory" "data" {
  directory_path = "${databricks_volume.this.volume_path}/data"
}
```

### Nested Directory Structure

```hcl
resource "databricks_volume_directory" "logs" {
  directory_path = "${databricks_volume.this.volume_path}/logs/2024/01"
}

resource "databricks_volume_directory" "raw_data" {
  directory_path = "${databricks_volume.this.volume_path}/raw/input"
}

resource "databricks_volume_directory" "processed_data" {
  directory_path = "${databricks_volume.this.volume_path}/processed/output"
}
```

### Directory with Files

```hcl
resource "databricks_volume_directory" "scripts" {
  directory_path = "${databricks_volume.this.volume_path}/scripts"
}

resource "databricks_file" "init_script" {
  source = "/local/path/to/init.sh"
  path   = "${databricks_volume_directory.scripts.id}/init.sh"
}
```

## Argument Reference

The following arguments are required:

* `directory_path` - (Required) The absolute path of the directory in a Unity Catalog volume. Must be in the format `/Volumes/<catalog>/<schema>/<volume>/<path>`. Changing this value will force recreation of the resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the directory resource, same as `directory_path`.

## Import

The resource `databricks_volume_directory` can be imported using the directory path:

```hcl
import {
  to = databricks_volume_directory.this
  id = "/Volumes/main/default/my_volume/my_directory"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_volume_directory.this /Volumes/main/default/my_volume/my_directory
```

## Related Resources

The following resources are often used in the same context:

* [databricks_file](file.md) to manage files in Unity Catalog volumes.
* [databricks_volume](volume.md) to manage [volumes within Unity Catalog](https://docs.databricks.com/en/connect/unity-catalog/volumes.html).
* [databricks_schema](schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](catalog.md) to manage catalogs within Unity Catalog.
