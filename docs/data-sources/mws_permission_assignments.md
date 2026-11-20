---
subcategory: "Security"
---
# databricks_mws_permission_assignments Data Source

[API Documentation](https://docs.databricks.com/api/account/workspaceassignment/list)

This data source lists all workspace permission assignments for a given workspace, using the Databricks account API. It is the read-only, symmetric counterpart of the [databricks_mws_permission_assignment](../resources/mws_permission_assignment.md) resource, which manages a single assignment.

-> This data source can only be used with an account-level provider!

## Example Usage

In account context, listing all permission assignments for a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

data "databricks_mws_permission_assignments" "this" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
}

output "assignments" {
  value = data.databricks_mws_permission_assignments.this.permission_assignments
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required) The numeric ID of the workspace to list permission assignments for.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `permission_assignments` - A list of all workspace permission assignments. Each element contains the following attributes:
    * `error` - Error response associated with a workspace permission assignment, if any.
    * `permissions` - The permission levels of the principal. Possible values are `USER` and `ADMIN`.
    * `principal` - Information about the principal assigned to the workspace:
        * `principal_id` - The unique, opaque ID of the principal.
        * `display_name` - The display name of the principal.
        * `group_name` - The group name of the group. Present only if the principal is a group.
        * `service_principal_name` - The name of the service principal. Present only if the principal is a service principal.
        * `user_name` - The username of the user. Present only if the principal is a user.

If the workspace has no permission assignments, an empty list will be returned.

## Related Resources

The following resources are used in the same context:

* [databricks_mws_permission_assignment](../resources/mws_permission_assignment.md) to manage a single workspace permission assignment in the account context.
* [databricks_permission_assignment](../resources/permission_assignment.md) to manage a single workspace permission assignment in the workspace context.
