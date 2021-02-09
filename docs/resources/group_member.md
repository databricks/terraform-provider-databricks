---
subcategory: "Security"
---
# databricks_group_member Resource

This resource allows you to attach [users](user.md) and [groups](group.md) as group members.

## Example Usage

After the following example, Bradley would have direct membership in group B and transitive membership in group A.

```hcl
resource "databricks_group" "a" {
    display_name = "A"
}

resource "databricks_group" "b" {
    display_name = "B"
}

resource "databricks_group_member" "ab" {
    group_id = databricks_group.a.id
    member_id = databricks_group.b.id
}

resource "databricks_user" "bradley" {
    user_name = "bradley@example.com"
}

resource "databricks_group_member" "bb" {
    group_id = databricks_group.b.id
    member_id = databricks_user.bradley.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) This is the id of the [group](group.md) resource.
* `member_id` - (Required) This is the id of the [group](group.md) or [user](user.md).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the `databricks_group_member` object which is in the format `<group_id>|<member_id>`.

## Import

-> **Note** Importing this resource is not currently supported.
