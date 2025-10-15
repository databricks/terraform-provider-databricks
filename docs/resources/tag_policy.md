---
subcategory: "Tags"
---
# databricks_tag_policy Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

Define tag policies to manage governed tags in your account.

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
* `tag_key` (string, required)
* `description` (string, optional)
* `values` (list of Value, optional)
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

### Value
* `name` (string, required)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Timestamp when the tag policy was created
* `id` (string)
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
terraform import databricks_tag_policy "tag_key"
```