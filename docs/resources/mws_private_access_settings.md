---
subcategory: "AWS"
---
# databricks_mws_private_access_settings Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Allows you to create a [Private Access Setting](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-5-create-a-private-access-settings-configuration-using-the-databricks-account-api) that can be used as part of a [databricks_mws_workspaces](mws_networks.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).

It is strongly recommended that customers read the [Enable Private Link](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) documentation before trying to leverage this resource.

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

## Example Usage

```hcl
resource "databricks_mws_private_access_settings" "pas" {
  provider                     = databricks.mws
  account_id                   = var.databricks_account_id
  private_access_settings_name = "Private Access Settings for ${local.prefix}"
  region                       = var.region
  public_access_enabled        = true
}

```

The `databricks_mws_private_access_settings.pas.private_access_settings_id` can then be used as part of a [databricks_mws_workspaces](databricks_mws_workspaces.md) resource:

```hcl
resource "databricks_mws_workspaces" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  aws_region                 = var.region
  workspace_name             = local.prefix
  deployment_name            = local.prefix
  credentials_id             = databricks_mws_credentials.this.credentials_id
  storage_configuration_id   = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "ENTERPRISE"
  depends_on                 = [databricks_mws_networks.this]
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
