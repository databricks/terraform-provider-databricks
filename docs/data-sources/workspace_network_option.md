---
subcategory: "Settings"
---
# databricks_workspace_network_option Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single workspace network option.

-> **Note** This data source can only be used with an account-level provider!

## Example Usage
Referring to a network policy by id:

```hcl
data "databricks_workspace_network_option" "this" {
  workspace_id = "9999999999999999"
}
```

## Arguments
The following arguments are supported:
* `workspace_id` (integer, required) - The workspace ID

## Attributes
The following attributes are exported:
* `network_policy_id` (string) - The network policy ID to apply to the workspace. This controls the network access rules
  for all serverless compute resources in the workspace. Each workspace can only be
  linked to one policy at a time. If no policy is explicitly assigned,
  the workspace will use 'default-policy'
* `workspace_id` (integer) - The workspace ID