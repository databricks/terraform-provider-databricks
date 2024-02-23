---
subcategory: "Settings"
---

# databricks_restrict_workspace_admins_setting Resource

-> **Note** This resource could be only used with workspace-level provider!

The `databricks_restrict_workspace_admins_setting` resource lets you control the capabilities of workspace admins.

With the status set to `ALLOW_ALL`, workspace admins can:

1. Create service principal personal access tokens on behalf of any service principal in their workspace.
2. Change a job owner to any user in the workspace.
3. Change the job run_as setting to any user in their workspace or a service principal on which they have the Service Principal User role.

With the status set to `RESTRICT_TOKENS_AND_JOB_RUN_AS`, workspace admins can:

1. Only create personal access tokens on behalf of service principals on which they have the Service Principal User role.
2. Only change a job owner to themselves.
3. Only change the job run_as setting to themselves a service principal on which they have the Service Principal User role.

-> **Note** Only account admins can update the setting. And the account admin must be part of the workspace to change the setting status.

## Example Usage

```hcl
resource "databricks_restrict_workspace_admins_setting" "this" {
  restrict_workspace_admins {
    status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
  }
}
```

## Argument Reference

The resource supports the following arguments:

* `restrict_workspace_admins` - (Required) The configuration details.
* `status` - (Required) The restrict workspace admins status for the workspace.
