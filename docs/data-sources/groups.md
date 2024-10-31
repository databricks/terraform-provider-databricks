---
subcategory: "Security"
---
# databricks_groups Data Source

Retrieves a list of all [databricks_group](../resources/group.md) display names associated with the workspace, or those matching the provided filter. Maximum 100 results.


## Example Usage

Get all groups:

```hcl
data "databricks_groups" "all" {}

output "all_groups" {
  value = data.databricksdatabricks_groups.all.display_names
}
```

Get groups whose displayName contains `foo` or displayName contains `bar`.:

```hcl
data "databricks_groups" "this" {
  filter = "displayName co \"foo\" or displayName co \"bar\""
}

output "foobar_groups" {
  value = data.databricks_pipelines.this.display_names
}
```


## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Query by which the results have to be filtered. See [API reference](https://docs.databricks.com/api/workspace/groups/list#filter).


## Attribute Reference

Data source exposes the following attributes:

* `display_names` - List of display names for [Groups](https://docs.databricks.com/data-engineering/delta-live-tables/index.html) matching the provided search criteria.


## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_group](../data-sources/group.md) to retrive specific [Group](https://docs.databricks.com/en/admin/users-groups/groups.html) information.
* [databricks_cluster](../resources/cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_directory](../resources/directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
* [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_user](../resources/user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](../resources/group.md) within the workspace.
