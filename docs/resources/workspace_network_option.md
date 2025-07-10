---
subcategory: "Settings"
---
# databricks_workspace_network_option Resource


## Example Usage


## Arguments
The following arguments are supported:
* `network_policy_id` (string, optional) - The network policy ID to apply to the workspace. This controls the network access rules
  for all serverless compute resources in the workspace. Each workspace can only be
  linked to one policy at a time. If no policy is explicitly assigned,
  the workspace will use 'default-policy'
* `workspace_id` (integer, optional) - The workspace ID

## Attributes
In addition to the above arguments, the following attributes are exported:

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = workspace_id
  to = databricks_workspace_network_option.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_network_option workspace_id
```
