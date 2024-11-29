---
subcategory: "Deployment"
---
# databricks_mws_private_access_settings Resource

Allows you to create a Private Access Setting resource that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) or [GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html)

It is strongly recommended that customers read the [Enable AWS Private Link](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) [Enable GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) documentation before trying to leverage this resource.

## Example Usage

## Databricks on AWS usage

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`

```hcl
resource "databricks_mws_private_access_settings" "pas" {
  provider                     = databricks.mws
  account_id                   = var.databricks_account_id
  private_access_settings_name = "Private Access Settings for ${local.prefix}"
  region                       = var.region
  public_access_enabled        = true
}

```

The `databricks_mws_private_access_settings.pas.private_access_settings_id` can then be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource:

```hcl
resource "databricks_mws_workspaces" "this" {
  provider                   = databricks.mws
  aws_region                 = var.region
  workspace_name             = local.prefix
  credentials_id             = databricks_mws_credentials.this.credentials_id
  storage_configuration_id   = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "ENTERPRISE"
  depends_on                 = [databricks_mws_networks.this]
}
```

## Databricks on GCP usage

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.gcp.databricks.com"` and use `provider = databricks.mws`

```hcl
resource "databricks_mws_workspaces" "this" {
  provider       = databricks.mws
  workspace_name = "gcp-workspace"
  location       = var.subnet_region
  cloud_resource_container {
    gcp {
      project_id = var.google_project
    }
  }
  gke_config {
    connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
    master_ip_range   = "10.3.0.0/28"
  }
  network_id                 = databricks_mws_networks.this.network_id
  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  pricing_tier               = "PREMIUM"
  depends_on                 = [databricks_mws_networks.this]
}
```

## Argument Reference

The following arguments are available:

* `private_access_settings_name` - Name of Private Access Settings in Databricks Account
* `public_access_enabled` (Boolean, Optional, `false` by default on AWS, `true` by default on GCP) - If `true`, the [databricks_mws_workspaces](mws_workspaces.md) can be accessed over the [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) as well as over the public network. In such a case, you could also configure an [databricks_ip_access_list](ip_access_list.md) for the workspace, to restrict the source networks that could be used to access it over the public network. If `false`, the workspace can be accessed only over VPC endpoints, and not over the public network. Once explicitly set, this field becomes mandatory.
* `region` - Region of AWS VPC or the Google Cloud VPC network
* `private_access_level` - (Optional) The private access level controls which VPC endpoints can connect to the UI or API of any workspace that attaches this private access settings object. `ACCOUNT` level access _(default)_ lets only [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) that are registered in your Databricks account connect to your [databricks_mws_workspaces](mws_workspaces.md). `ENDPOINT` level access lets only specified [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) connect to your workspace. Please see the `allowed_vpc_endpoint_ids` documentation for more details.
* `allowed_vpc_endpoint_ids` - (Optional) An array of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md#vpc_endpoint_id) `vpc_endpoint_id` (not `id`). Only used when `private_access_level` is set to `ENDPOINT`. This is an allow list of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) that in your account that can connect to your [databricks_mws_workspaces](mws_workspaces.md) over AWS PrivateLink. If hybrid access to your workspace is enabled by setting `public_access_enabled` to true, then this control only works for PrivateLink connections. To control how your workspace is accessed via public internet, see the article for [databricks_ip_access_list](ip_access_list.md).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the ID of the Private Access Settings in form of `account_id/private_access_settings_id`.
* `private_access_settings_id` - Canonical unique identifier of Private Access Settings in Databricks Account
* `status` - (AWS only) Status of Private Access Settings

## Import

This resource can be imported by Databricks account ID and private access settings ID.

```sh
terraform import databricks_mws_private_access_settings.this '<account_id>/<private_access_settings_id>'
```

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with Private Link](../guides/aws-private-link-workspace.md) guide.
* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [Provisioning Databricks workspaces on GCP with Private Service Connect](../guides/gcp-private-service-connect-workspace.md) guide.
* [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) to register [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) resources with Databricks such that they can be used as part of a [databricks_mws_networks](mws_networks.md) configuration.
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_workspaces](mws_workspaces.md) to set up [AWS and GCP workspaces](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
