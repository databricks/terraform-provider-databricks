---
subcategory: "Tags"
---
# databricks_tag_policies Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to list all tag policies in the account.

The following resources are often used in the same context:

* [databricks_entity_tag_assignment](../resources/entity_tag_assignment.md) for assigning tags to supported Unity Catalog entities.
* [databricks_workspace_entity_tag_assignment](../resources/workspace_entity_tag_assignment.md) for assigning tags to supported workspace entities.
* [databricks_policy_info](../resources/policy_info.md) for defining ABAC policies using governed tags.
* [databricks_access_control_rule_set](../resources/access_control_rule_set.md) for managing account-level and individual tag policy permissions.

-> **Note** This resource can only be used with a workspace-level provider!


## Example Usage
Getting a list of all tag policies:

```hcl
data "databricks_tag_policies" "all" {}

output "all_tag_policies" {
  value = data.databricks_tag_policies.all.tag_policies
}
```

## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return in this request. Fewer results may be returned than requested. If
  unspecified or set to 0, this defaults to 1000. The maximum value is 1000; values above 1000 will be coerced down
  to 1000


## Attributes
This data source exports a single attribute, `tag_policies`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Timestamp when the tag policy was created
* `description` (string)
* `id` (string)
* `tag_key` (string)
* `update_time` (string) - Timestamp when the tag policy was last updated
* `values` (list of Value)

### Value
* `name` (string)
