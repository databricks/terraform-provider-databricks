---
subcategory: "Unity Catalog"
---
# databricks_share Resource

-> **Note** This resource could be only used with workspace-level provider!

Within a metastore, Unity Catalog provides the ability to create a share, which is a named object that contains a collection of tables in a metastore that you want to share as a group. A share can contain tables from only a single metastore. You can add or remove tables from a share at any time.

A `databricks_share` is contained within [databricks_metastore](metastore.md) and can contain a list of tables.

## Example Usage

-> **Note** In Terraform configuration, it is recommended to define objects in alphabetical order of their `name` arguments, so that you get consistent and readable diff. Whenever objects are added or removed, or `name` is renamed, you'll observe a change in the majority of tasks. It's related to the fact that the current version of the provider treats `object` blocks as an ordered list. Alternatively, `object` block could have been an unordered set, though end-users would see the entire block replaced upon a change in single property of the task.

Creating a Delta Sharing share and add some existing tables to it

```hcl
data "databricks_tables" "things" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

resource "databricks_share" "some" {
  name = "my_share"
  dynamic "object" {
    for_each = data.databricks_tables.things.ids
    content {
      name             = object.value
      data_object_type = "TABLE"
    }
  }
}
```

Creating a Delta Sharing share and add a schema to it(including all current and future tables).

```hcl
resource "databricks_share" "schema_share" {
  name = "schema_share"
  object {
    name                        = "catalog_name.schema_name"
    data_object_type            = "SCHEMA"
    history_data_sharing_status = "ENABLED"
  }
}
```

Creating a Delta Sharing share and share a table with partitions spec and history

```hcl
resource "databricks_share" "some" {
  name = "my_share"
  object {
    name                        = "my_catalog.my_schema.my_table"
    data_object_type            = "TABLE"
    history_data_sharing_status = "ENABLED"
    partition {
      value {
        name  = "year"
        op    = "EQUAL"
        value = "2009"
      }
      value {
        name  = "month"
        op    = "EQUAL"
        value = "12"
      }
    }
    partition {
      value {
        name  = "year"
        op    = "EQUAL"
        value = "2010"
      }
    }
  }
}
```

## Argument Reference

The following arguments are required:

* `name` (Required) - Name of share. Change forces creation of a new resource.
* `owner` (Optional) -  User name/group name/sp application_id of the share owner.

### object Configuration Block

* `name` (Required) - Full name of the object, e.g. `catalog.schema.name` for a table.
* `data_object_type` (Required) - Type of the data object, currently `TABLE`, `SCHEMA`, `VOLUME` are supported.
* `comment` (Optional) -  Description about the object.
* `shared_as` (Optional) - A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the `shared_as` name. The `shared_as` name must be unique within a Share. Change forces creation of a new resource.
* `cdf_enabled` (Optional) - Whether to enable Change Data Feed (cdf) on the shared object. When this field is set, field `history_data_sharing_status` can not be set.
* `start_version` (Optional) -  The start version associated with the object for cdf. This allows data providers to control the lowest object version that is accessible by clients.
* `history_data_sharing_status` (Optional) - Whether to enable history sharing, one of: `ENABLED`, `DISABLED`. When a table has history sharing enabled, recipients can query table data by version, starting from the current table version. If not specified, clients can only query starting from the version of the object at the time it was added to the share. *NOTE*: The start_version should be less than or equal the current version of the object. When this field is set, field `cdf_enabled` can not be set.

To share only part of a table when you add the table to a share, you can provide partition specifications. This is specified by a number of `partition` blocks. Each entry in `partition` block takes a list of `value` blocks. The field is documented below.

#### value Configuration Block

* `name` - The name of the partition column.
* `op` - The operator to apply for the value, one of: `EQUAL`, `LIKE`
* `recipient_property_key` (Optional) - The key of a Delta Sharing recipient's property. For example `databricks-account-id`. When this field is set, field `value` can not be set.
* `value` (Optional) - The value of the partition column. When this value is not set, it means null value. When this field is set, field `recipient_property_key` can not be set.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the ID of the share, the same as `name`.
* `created_at` - Time when the share was created.
* `created_by` - The principal that created the share.
* `status` - Status of the object, one of: `ACTIVE`, `PERMISSION_DENIED`.

## Related Resources

The following resources are often used in the same context:

* [databricks_recipient](recipient.md) to create Delta Sharing recipients.
* [databricks_grants](grants.md) to manage Delta Sharing permissions.
* [databricks_shares](../data-sources/shares.md) to read existing Delta Sharing shares.
