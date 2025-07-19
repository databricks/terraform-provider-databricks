---
subcategory: "Security"
---
# databricks_groups Data Source

Retrieves information about multiple [databricks_group](../resources/group.md) objects matching a filter criteria. This data source allows you to search for groups using SCIM filter expressions and returns detailed information about matching groups including their members, entitlements and instance profiles.

-> This data source can be used with an account or workspace-level provider.

## Example Usage

Retrieve all groups with "admin" in the name:

```hcl
data "databricks_groups" "admin_groups" {
  filter = "displayName co \"admin\""
}

resource "databricks_permissions" "cluster_permissions" {
  cluster_id = databricks_cluster.example.id

  dynamic "access_control" {
    for_each = data.databricks_groups.admin_groups.groups
    content {
      group_name       = access_control.value.display_name
      permission_level = "CAN_MANAGE"
    }
  }
}
```

Find groups by exact display name:

```hcl
data "databricks_groups" "specific_teams" {
  filter = "displayName eq \"team-data\" or displayName eq \"team-ml\""
}

locals {
  team_principals = [for group in data.databricks_groups.specific_teams.groups : group.acl_principal_id]
}
```

Find all groups without a filter (retrieve all groups):

```hcl
data "databricks_groups" "all_groups" {
  filter = ""
}

output "total_groups" {
  value = length(data.databricks_groups.all_groups.groups)
}
```

## Argument Reference

Data source allows you to filter groups by the following attributes:

* `filter` - (Optional) SCIM filter expression to match groups. Uses the same filter syntax as the Databricks SCIM API. Common examples:
  * `displayName co "admin"` - Groups containing "admin" in the display name
  * `displayName eq "admins"` - Groups with exact display name "admins"
  * `displayName sw "team-"` - Groups whose display name starts with "team-"
  * Empty string or omitted - Returns all groups

## Attribute Reference

Data source exposes the following attributes:

* `groups` - List of group objects matching the filter criteria. Each group object contains:
  * `display_name` - Display name of the group.
  * `external_id` - ID of the group in an external identity provider.
  * `acl_principal_id` - Identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), formatted as `groups/{display_name}`.
  * `members` - Set of all group member identifiers (users, service principals, and child groups combined).
  * `users` - Set of [databricks_user](../resources/user.md) identifiers that are members of this group.
  * `service_principals` - Set of [databricks_service_principal](../resources/service_principal.md) identifiers that are members of this group.
  * `child_groups` - Set of [databricks_group](../resources/group.md) identifiers that are child groups of this group.
  * `groups` - Set of parent [databricks_group](../resources/group.md) identifiers that this group belongs to.
  * `instance_profiles` - Set of [instance profile](../resources/instance_profile.md) ARNs associated with this group.

-> **Note**: The `groups` attribute returns a list sorted by `display_name` for consistent ordering. All member sets within each group are also sorted alphabetically.

## SCIM Filter Syntax

The `filter` parameter supports SCIM 2.0 filter expressions. Common operators include:

* `co` - Contains (case insensitive)
* `eq` - Equals (case sensitive)
* `ne` - Not equals
* `sw` - Starts with
* `ew` - Ends with
* `and` - Logical AND
* `or` - Logical OR

Examples:
* `displayName co "admin"` - Contains "admin"
* `displayName sw "team-" and displayName ew "-prod"` - Starts with "team-" and ends with "-prod"
* `displayName eq "admins" or displayName eq "developers"` - Exact match for either name

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_group](group.md) to retrieve information about a single group
* [databricks_group](../resources/group.md) to manage groups in Databricks workspace
* [databricks_group_member](../resources/group_member.md) to manage group membership
* [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace
* [databricks_user](../resources/user.md) to manage users that can be added to groups
* [databricks_service_principal](../resources/service_principal.md) to manage service principals that can be added to groups