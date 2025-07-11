---
subcategory: "Settings"
---
# databricks_workspace_network_option Data Source


## Example Usage


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