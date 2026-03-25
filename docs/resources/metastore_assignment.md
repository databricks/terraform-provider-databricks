---
subcategory: "Unity Catalog"
---
# databricks_metastore_assignment (Resource)

-> This resource can be used with an account or workspace-level provider.

A single [databricks_metastore](metastore.md) can be shared across Databricks workspaces, and each linked workspace has a consistent view of the data and a single set of access policies. You can only create a single metastore for each region in which your organization operates.

## Example Usage

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = "uc admins"
  region        = "us-east-1"
  force_destroy = true
}

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `metastore_id` - Unique identifier of the parent Metastore
* `workspace_id` - id of the workspace for the assignment
* `default_catalog_name` - (Deprecated) Default catalog used for this assignment. Please use [databricks_default_namespace_setting](default_namespace_setting.md) instead.
* `api` - (Optional) Specifies whether to use account-level or workspace-level API. Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this metastore assignment in form of `<workspace_id>|<metastore_id>`.

## Import

This resource can be imported by combination of workspace id and metastore id:

```hcl
import {
  to = databricks_metastore_assignment.this
  id = "<workspace_id>|<metastore_id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_metastore_assignment.this "<workspace_id>|<metastore_id>"
```
