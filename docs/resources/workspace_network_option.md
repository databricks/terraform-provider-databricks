---
subcategory: "Settings"
---
# databricks_workspace_network_option Resource
Workspace network options allow configuration of network settings for Databricks workspaces by selecting which network policy to associate with the workspace.

Each workspace is always associated with exactly one network policy that controls which network destinations can be accessed from the Databricks environment. By default, workspaces are associated with the `default-policy` network policy.

This resource has the following characteristics:

- You cannot create or delete a workspace's network option
- You can only update it to associate the workspace with a different policy
- This resource is used to change the network policy assignment for existing workspaces

-> **Note** This resource can only be used with an account-level provider!

## Example Usage
```hcl
resource "databricks_workspace_network_option"    "example_workspace_network_option" {
  workspace_id         = "9999999999999999"
  network_policy_id    = "default-policy"
}
```

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
  id = "workspace_id"
  to = databricks_workspace_network_option.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_network_option "workspace_id"
```