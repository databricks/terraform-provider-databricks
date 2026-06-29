---
subcategory: "Workspace"
---
# databricks_genie_space Resource

This resource allows you to manage [Databricks AI/BI Genie Spaces](https://docs.databricks.com/en/genie/index.html). A Genie space provides a no-code experience for business users to ask natural-language questions, powered by AI/BI on data registered in Unity Catalog. To manage Genie spaces you must have CAN USE permission on a Pro or Serverless SQL warehouse, and Databricks Assistant must be enabled.

-> This resource can only be used with a workspace-level provider!

## Example Usage

Genie space with an inline `serialized_space`:

```hcl
data "databricks_sql_warehouse" "starter" {
  name = "Starter Warehouse"
}

resource "databricks_genie_space" "sales" {
  title            = "Sales Genie"
  warehouse_id     = data.databricks_sql_warehouse.starter.id
  parent_path      = "/Shared/genie-spaces"
  description      = "Ask questions about sales data"
  serialized_space = jsonencode({ datasets = [], conversations = [] })
}
```

Loading `serialized_space` from a file (GitOps-friendly pattern):

```hcl
resource "databricks_genie_space" "sales" {
  title            = "Sales Genie"
  warehouse_id     = data.databricks_sql_warehouse.starter.id
  parent_path      = "/Shared/genie-spaces"
  serialized_space = file("${path.module}/spaces/sales.json")
}
```

Manage access with [databricks_permissions](permissions.md):

```hcl
resource "databricks_permissions" "sales_genie" {
  genie_space_id = databricks_genie_space.sales.space_id

  access_control {
    group_name       = "data-analysts"
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = "platform-admins"
    permission_level = "CAN_MANAGE"
  }
}
```

Attach tags via the existing tag-assignment resource:

```hcl
resource "databricks_workspace_entity_tag_assignment" "env" {
  entity_type = "geniespaces"
  entity_id   = databricks_genie_space.sales.space_id
  tag_key     = "env"
  tag_value   = "prod"
}
```

## Argument Reference

The following arguments are supported:

* `title` - (Required) Title of the Genie space.
* `warehouse_id` - (Required) ID of the SQL warehouse associated with the Genie space.
* `serialized_space` - (Required) JSON-encoded contents of the Genie space. Whitespace and key-order differences between configuration and the API representation are normalized and do not cause spurious diffs. For GitOps-friendly workflows, load the JSON from disk with the built-in `file()` function, e.g. `serialized_space = file("${path.module}/sales.json")`.
* `description` - (Optional) Description of the Genie space.
* `parent_path` - (Required) The workspace path of the folder containing the Genie space. Includes a leading slash and no trailing slash. If the folder does not exist it is created on first apply. Changing this attribute forces replacement of the resource.
* `provider_config` - (Optional) Configure the provider for management through an account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Genie space ID (same as `space_id`).
* `space_id` - The Genie space ID assigned by Databricks.
* `etag` - ETag returned by the API on every read. Informational only; the provider never sends it on update and operates on a last-writer-wins basis. Concurrency control can be added later if needed.

## Access Control

[databricks_permissions](permissions.md) can grant `CAN_READ`, `CAN_RUN`, `CAN_EDIT` or `CAN_MANAGE` on the Genie space by setting `genie_space_id`. The provider implicitly adds the current user as `CAN_MANAGE` on every ACL update (create/update) and on delete to avoid locking the calling principal out of the space.

-> The Databricks workspace UI shows the read-only level as `CAN VIEW`, while the Permissions API (and therefore this resource) uses `CAN_READ` for the same level. See [Genie Space ACLs](https://docs.databricks.com/aws/en/security/auth/access-control/#genie-space-acls) for the full ability matrix.

## Import

You can import a `databricks_genie_space` resource using its `space_id`:

```hcl
import {
  to = databricks_genie_space.this
  id = "<space-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_genie_space.this <space-id>
```

## Notes

* `serialized_space` accepts any valid JSON; the provider re-marshals both the configuration and the API response into a canonical form to suppress whitespace and key-order diffs.
* On delete the provider issues a trash (soft-delete) call. An already-trashed space is treated as successfully deleted, so re-running `terraform destroy` after a manual trash is idempotent.
