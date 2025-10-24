---
subcategory: "Tags"
---
# databricks_tag_policy Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single tag policy by its tag key.

-> **Note** This resource can only be used with a workspace-level provider!

## Example Usage
Referring to a tag policy by its tag key:

```hcl
data "databricks_tag_policy" "example_tag_policy" {
  tag_key = "example_tag_key"
}
```

## Arguments
The following arguments are supported:
* `tag_key` (string, required)

## Attributes
The following attributes are exported:
* `create_time` (string) - Timestamp when the tag policy was created
* `description` (string)
* `id` (string)
* `tag_key` (string)
* `update_time` (string) - Timestamp when the tag policy was last updated
* `values` (list of Value)

### Value
* `name` (string)