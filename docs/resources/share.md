---
subcategory: "Unity Catalog"
---
# databricks_share Resource

Within a metastore, Unity Catalog provides the ability to create a share, which is a named object that contains a collection of tables in a metastore that you want to share as a group. A share can contain tables from only a single metastore. You can add or remove tables from a share at any time.

A `databricks_share` is contained within [databricks_metastore](metastore.md) and can contain a list of shares.

## Example Usage

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

## Argument Reference

The following arguments are required:

* `name` (Required) - Name of share. Change forces creation of a new resource.

### object Configuration Block

* `name` (Required) - Full name of the object, e.g. `catalog.schema.name` for a table.
* `data_object_type` (Required) - Type of the object, currently only `TABLE` is allowed.
* `comment` (Optional) -  Description about the object.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `created_at` - Time when the share was created.
* `created_by` - The principal that created the share.

## Related Resources

The following resources are often used in the same context:

* [databricks_recipient](recipient.md) to create Delta Sharing recipients.
* [databricks_grants](grants.md) to manage Delta Sharing permissions.
* [databricks_shares](../data-sources/shares.md) to read existing Delta Sharing shares.
