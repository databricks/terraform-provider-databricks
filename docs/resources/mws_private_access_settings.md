---
subcategory: "AWS"
---
# databricks_mws_private_access_settings Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

## Example Usage

```hcl
resource "databricks_mws_private_access_settings" "pas" {
  account_id                   = var.databricks_account_id
  private_access_settings_name = "Private Access Settings for ${aws_vpc.main.id}"
  region                       = local.region
}
```

## Argument Reference

The following arguments are available:

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `private_access_settings_name` - Name of Private Access Settings in Databricks Account
* `public_access_enabled` (Boolean, Optional, `false` by default) - If `true`, the [databricks_mws_workspaces](mws_workspaces.md) can be accessed over the [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) as well as over the public network. In such a case, you could also configure an [databricks_ip_access_list](ip_access_list.md) for the workspace, to restrict the source networks that could be used to access it over the public network. If `false` (default), the workspace can be accessed only over VPC endpoints, and not over the public network.
* `region` - Region of AWS VPC

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `private_access_settings_id` - Canonical unique identifier of Private Access Settings in Databricks Account
* `status` - Status of Private Access Settings
