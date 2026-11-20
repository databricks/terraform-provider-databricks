---
subcategory: "Tags"
---
# databricks_tag_policy Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/tagpolicies)

Define tag policies to manage governed tags in your account.

The following resources are often used in the same context:

* [databricks_entity_tag_assignment](entity_tag_assignment.md) for assigning tags to supported Unity Catalog entities.
* [databricks_workspace_entity_tag_assignment](workspace_entity_tag_assignment.md) for assigning tags to supported workspace entities.
* [databricks_policy_info](policy_info.md) for defining ABAC policies using governed tags.
* [databricks_access_control_rule_set](access_control_rule_set.md) for managing account-level and individual tag policy permissions.

-> **Note** This resource can only be used with a workspace-level provider!


## Example Usage
```hcl
resource "databricks_tag_policy" "example_tag_policy" {
  tag_key     = "example_tag_key"
  description = "Example description."
  values = [
    {
      name = "example_value_1"
    },
    {
      name = "example_value_2"
    },
    {
      name = "example_value_3"
    }
  ]
}
```

## Arguments
The following arguments are supported:
* `tag_key` (string, required) - The tag key that this policy governs. Tag keys are case-sensitive
* `description` (string, optional) - A description of the governed tag, such as its purpose or the meaning of its values
* `values` (list of Value, optional) - The allowed values of the governed tag. Only these values can be assigned to the tag key, and each
  value must be unique. A policy without values defines a key-only governed tag
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### Value
* `name` (string, required) - The value that can be assigned to the tag key. Values are case-sensitive

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Timestamp when the tag policy was created
* `id` (string) - The unique identifier of the tag policy, generated when the policy is created
* `update_time` (string) - Timestamp when the tag policy was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "tag_key"
  to = databricks_tag_policy.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_tag_policy.this "tag_key"
```