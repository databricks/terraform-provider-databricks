---
subcategory: "Workspace"
---
# databricks_genie_spaces Data Source

This data source allows you to retrieve information about [Databricks AI/BI Genie Spaces](https://docs.databricks.com/en/genie/index.html) in the workspace.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_genie_spaces" "all" {}

resource "databricks_permissions" "all_genie_spaces" {
  for_each = toset(data.databricks_genie_spaces.all.spaces[*].space_id)

  genie_space_id = each.value

  access_control {
    group_name       = "data-analysts"
    permission_level = "CAN_RUN"
  }
}
```

Filter by title (case-insensitive substring):

```hcl
data "databricks_genie_spaces" "sales" {
  title_contains = "sales"
}
```

## Argument Reference

The following arguments are supported:

* `title_contains` - (Optional) A **case-insensitive** substring used to filter Genie spaces by their `title`.
* `provider_config` - (Optional) Configure the provider for management through an account provider. This block consists of the following fields:
  * `workspace_id` - (Optional) Workspace ID which the data source belongs to.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `spaces` - A list of Genie spaces matching the specified criteria. Each element contains the following attributes:
  * `space_id` - The Genie space ID.
  * `title` - The Genie space title.
  * `description` - The Genie space description.
  * `warehouse_id` - The associated SQL warehouse ID.
  * `parent_path` - The workspace folder containing the Genie space.
  * `etag` - The current ETag for the Genie space.
  * `serialized_space` - Always empty for list responses; fetch a single space with [databricks_genie_space](../resources/genie_space.md) to get the full serialized form.

If no matches are found, an empty list is returned.
