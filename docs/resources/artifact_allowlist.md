---
subcategory: "Unity Catalog"
---
# databricks_artifact_allowlist Resource

In Databricks Runtime 13.3 and above, you can add libraries and init scripts to the allowlist in UC so that users can leverage these artifacts on compute configured with shared access mode.

~> It is required to define all allowlist for an artifact type in a single resource, otherwise Terraform cannot guarantee config drift prevention.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_artifact_allowlist" "init_scripts" {
  artifact_type = "INIT_SCRIPT"
  artifact_matcher {
    artifact   = "/Volumes/inits"
    match_type = "PREFIX_MATCH"
  }
}
```

## Argument Reference

The following arguments are required:

* `artifact_type` - The artifact type of the allowlist. Can be `INIT_SCRIPT`, `LIBRARY_JAR` or `LIBRARY_MAVEN`. Change forces creation of a new resource.

One or more `artifact_matcher` blocks with the following arguments:

* `artifact` - The artifact path or maven coordinate.
* `match_type` - The pattern matching type of the artifact. Only `PREFIX_MATCH` is supported.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `metastore_id` - ID of the parent metastore.
* `created_at` -  Time at which this artifact allowlist was set.
* `created_by` -  Identity that set the artifact allowlist.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the artifact allow list in form of `metastore_id|artifact_type`.

## Import

This resource can be imported by name:

```hcl
import {
  to = databricks_artifact_allowlist.this
  id = "<metastore_id>|<artifact_type>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_artifact_allowlist.this "<metastore_id>|<artifact_type>"
```

## Related Resources

The following resources are used in the same context:

* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
