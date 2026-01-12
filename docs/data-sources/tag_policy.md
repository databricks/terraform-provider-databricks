---
subcategory: "Tags"
---
# databricks_tag_policy Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single tag policy by its tag key.

The following resources are often used in the same context:

* [databricks_entity_tag_assignment](../resources/entity_tag_assignment.md) for assigning tags to supported Unity Catalog entities.
* [databricks_workspace_entity_tag_assignment](../resources/workspace_entity_tag_assignment.md) for assigning tags to supported workspace entities.
* [databricks_policy_info](../resources/policy_info.md) for defining ABAC policies using governed tags.
* [databricks_access_control_rule_set](../resources/access_control_rule_set.md) for managing account-level and individual tag policy permissions.

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