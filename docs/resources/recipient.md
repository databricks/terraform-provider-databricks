---
subcategory: "Unity Catalog"
---
# databricks_share Resource

Within a metastore, Unity Catalog provides the ability to create a recipient to attach delta shares to..

A `databricks_recipient` is contained within [databricks_metastore](metastore.md) and can contain a list of shares.

## Example Usage

```hcl
resource "databricks_recipient" "db2open" {
  name = "db2open-recipient"
  comment = "made by terraform"
  authentication_type = "TOKEN"
  sharing_code = "my code "
  ip_access_list {
    allowed_ip_addresses = ["0.0.0.0/0"]
  }
}

resource "databricks_recipient" "db2db" {
  name = "sri-terraform-test-db2db-recipient"
  comment = "made by terraform"
  authentication_type = "DATABRICKS"
  data_recipient_global_metastore_id = "<cloud>:<region>:<guid>"
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of share. Change forces creation of a new resource.
* `comment` - (Optional) Description about the recipient.
* `sharing_code` - (Optional) The one-time sharing code provided by the data recipient.
* `authentication_type` - (Optional) The delta sharing authentication type. Valid values are `TOKEN` and `DATABRICKS`.
* `data_recipient_global_metastore_id` - Required when authentication_type is DATABRICKS.
* `ip_access_list` - (Optional) The one-time sharing code provided by the data recipient.

### Ip Access List Argument
Only one `ip_access_list` blocks is allowed in a recipient. It conflicts with authentication type DATABRICKS.

```hcl
  ip_access_list {
    allowed_ip_addresses = ["0.0.0.0/0"]
  }
```

Arguments for the `ip_access_list` block are:

Exactly one of the below arguments is required:
* `allowed_ip_addresses` - Allowed IP Addresses in CIDR notation. Limit of 100.

## Attribute Reference:

* `tokens` - (Optional) List of Recipient Tokens.

## Related Resources

The following resources are used in the same context:

* [databricks_table](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schema](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalog](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.