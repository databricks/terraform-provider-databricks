---
subcategory: "Delta Sharing"
---
# databricks_recipient Resource

-> This resource can only be used with a workspace-level provider!

In Delta Sharing, a recipient is an entity that receives shares from a provider. In Unity Catalog, a share is a securable object that represents an organization and associates it with a credential or secure sharing identifier that allows that organization to access one or more shares.

As a data provider (sharer), you can define multiple recipients for any given Unity Catalog metastore, but if you want to share data from multiple metastores with a particular user or group of users, you must define the recipient separately for each metastore. A recipient can have access to multiple shares.

A `databricks_recipient` is contained within [databricks_metastore](metastore.md) and can have permissions to `SELECT` from a list of shares.

## Example Usage

### Databricks Sharing with non databricks recipient

Setting `authentication_type` type to `TOKEN` creates a temporary url to download a credentials file. This is used to
authenticate to the sharing server to access data. This is for when the recipient is not using Databricks.

```hcl
resource "random_password" "db2opensharecode" {
  length  = 16
  special = true
}

data "databricks_current_user" "current" {}

resource "databricks_recipient" "db2open" {
  name                = "${data.databricks_current_user.current.alphanumeric}-recipient"
  comment             = "made by terraform"
  authentication_type = "TOKEN"
  sharing_code        = random_password.db2opensharecode.result
  ip_access_list {
    allowed_ip_addresses = [] // .. fill in allowed IPv4 addresses (CIDR notation allowed)
  }
}
```

### Databricks to Databricks Sharing

Setting `authentication_type` type to `DATABRICKS` allows you to automatically create a provider for a recipient who
is using Databricks. To do this they would need to provide the global metastore id that you will be sharing with. The
global metastore id follows the format: `<cloud>:<region>:<guid>`

```hcl
data "databricks_current_user" "current" {}

resource "databricks_metastore" "recipient_metastore" {
  name = "recipient"
  storage_root = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_container.unity_catalog.name,
  azurerm_storage_account.unity_catalog.name)
  delta_sharing_scope                               = "INTERNAL"
  delta_sharing_recipient_token_lifetime_in_seconds = "60000000"
  force_destroy                                     = true
}

resource "databricks_recipient" "db2db" {
  name                               = "${data.databricks_current_user.current.alphanumeric}-recipient"
  comment                            = "made by terraform"
  authentication_type                = "DATABRICKS"
  data_recipient_global_metastore_id = databricks_metastore.recipient_metastore.global_metastore_id
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of recipient. Change forces creation of a new resource.
* `comment` - (Optional) Description about the recipient.
* `sharing_code` - (Optional) The one-time sharing code provided by the data recipient.
* `owner` - (Optional) Username/groupname/sp application_id of the recipient owner.
* `authentication_type` - (Optional) The delta sharing authentication type. Valid values are `TOKEN` and `DATABRICKS`.
* `data_recipient_global_metastore_id` - Required when `authentication_type` is `DATABRICKS`.
* `ip_access_list` - (Optional) Recipient IP access list.
* `properties_kvpairs` - (Optional) Recipient properties - object consisting of following fields:
  * `properties` (Required) a map of string key-value pairs with recipient's properties.  Properties with name starting with `databricks.` are reserved.

### Ip Access List Argument

Only one `ip_access_list` block is allowed in a recipient. It conflicts with authentication type `DATABRICKS`.

```hcl
ip_access_list {
  allowed_ip_addresses = ["0.0.0.0/0"]
}
```

Arguments for the `ip_access_list` block are:

Exactly one of the below arguments is required:

* `allowed_ip_addresses` - Allowed IP Addresses in CIDR notation. Limit of 100.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the ID of the recipient - the same as the `name`.
* `tokens` - List of Recipient Tokens. This field is only present when the authentication_type is TOKEN. Each list element is an object with following attributes:
  * `id` - Unique ID of the recipient token.
  * `created_at` - Time at which this recipient Token was created, in epoch milliseconds.
  * `created_by` - Username of recipient token creator.
  * `activation_url` - Full activation URL to retrieve the access token. It will be empty if the token is already retrieved.
  * `expiration_time` - Expiration timestamp of the token in epoch milliseconds.
  * `updated_at` - Time at which this recipient Token was updated, in epoch milliseconds.
  * `updated_by` - Username of recipient Token updater.
* `created_at` - Time at which this recipient was created, in epoch milliseconds.
* `created_by` - Username of recipient creator.
* `updated_at` - Time at which this recipient was updated, in epoch milliseconds.
* `updated_by` - Username of recipient Token updater.
* `metastore_id` - Unique identifier of recipient's Unity Catalog metastore. This field is only present when the authentication_type is `DATABRICKS`.
* `cloud` - Cloud vendor of the recipient's Unity Catalog Metstore. This field is only present when the authentication_type is `DATABRICKS`.
* `region` - Cloud region of the recipient's Unity Catalog Metstore. This field is only present when the authentication_type is `DATABRICKS`.

## Import

The recipient resource can be imported using the name of the recipient.

```bash
terraform import databricks_recipient.this <recipient_name>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_share](share.md) to create Delta Sharing shares.
* [databricks_grants](grants.md) to manage Delta Sharing permissions.
* [databricks_shares](../data-sources/shares.md) to read existing Delta Sharing shares.
