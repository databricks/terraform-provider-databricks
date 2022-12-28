---
subcategory: "GCP"
---
# databricks_mws_workspaces resource

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

-> **Note** Initialize provider with `alias = "mws"`, `host  = "https://accounts.gcp.databricks.com"` and use `provider = databricks.mws` for all `databricks_mws_*` resources. We require all `databricks_mws_*` resources to be created within its own dedicated terraform module of your environment, which also creates the VPC as well. Code that creates workspaces and code that [manages workspaces](../guides/workspace-management.md) must be in separate terraform modules to avoid common confusion between `provider = databricks.mws` and `provider = databricks.created_workspace`. This is why we specify `databricks_host` and `databricks_token` outputs, that have to be used in the latter modules:

```hcl
provider "databricks" {
  host  = module.ai.databricks_host
  token = module.ai.databricks_token
}
```

This resource allows you to set up [workspaces on GCP](https://docs.gcp.databricks.com/administration-guide/account-settings-gcp/workspaces.html). Please follow this [complete runnable example](../guides/gcp-workspace.md) with new VPC and new workspace setup.

## Example Usage

* [databricks_mws_networks](mws_networks_gcp.md) - (optional, but recommended) You can share one [customer-managed VPC](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) with multiple workspaces in a single account. You do not have to create a new VPC for each workspace. However, you cannot reuse subnets with other resources, including other workspaces or non-Databricks resources. If you plan to share one VPC with multiple workspaces, be sure to size your VPC and subnets accordingly. Because a Databricks [databricks_mws_networks](mws_networks_gcp.md) encapsulates this information, you cannot reuse it across workspaces.

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}
variable "databricks_google_service_account" {}
variable "google_project" {}

provider "databricks" {
  alias = "mws"
  host  = "https://accounts.gcp.databricks.com"
}


// register VPC
resource "databricks_mws_networks" "this" {
  account_id   = var.databricks_account_id
  network_name = "${var.prefix}-network"
  gcp_network_info {
    network_project_id    = var.google_project
    vpc_id                = var.vpc_id
    subnet_id             = var.subnet_id
    subnet_region         = var.subnet_region
    pod_ip_range_name     = "pods"
    service_ip_range_name = "svc"
  }
}

// create workspace in given VPC
resource "databricks_mws_workspaces" "this" {
  account_id     = var.databricks_account_id
  workspace_name = var.prefix
  location       = var.subnet_region
  cloud_resource_container {
    gcp {
      project_id = var.google_project
    }
  }

  network_id = databricks_mws_networks.this.network_id
  gke_config {
    connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
    master_ip_range   = "10.3.0.0/28"
  }
  
  token {}
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

## Workspace with Databricks-Managed VPC

![VPCs](https://docs.databricks.com/_images/customer-managed-vpc.png)

By default, Databricks creates a VPC in your GCP project for each workspace. Databricks uses it for running clusters in the workspace. Optionally, you can use your VPC for the workspace, using the feature customer-managed VPC. Databricks recommends that you provide your VPC with [databricks_mws_networks](mws_networks_gcp.md) so that you can configure it according to your organization’s enterprise cloud standards while still conforming to Databricks requirements. You cannot migrate an existing workspace to your VPC.

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

data "google_client_openid_userinfo" "me" {
}
 
data "google_client_config" "current" {
}

resource "databricks_mws_workspaces" "this" {
 provider       = databricks.accounts
 account_id     = var.databricks_account_id
 workspace_name = var.prefix
 location       = data.google_client_config.current.region
 
 cloud_resource_container {
   gcp {
     project_id = data.google_client_config.current.project
   }
 }

 gke_config {
    connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
    master_ip_range   = "10.3.0.0/28"
 } 

 token {}
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

## Argument Reference

-> **Note** All workspaces would be verified to get into runnable state or deleted upon failure.

The following arguments are available and cannot be changed after workspace is created:

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/).
* `workspace_name` - name of the workspace, will appear on UI
* `location` - GCP region of the subnet
* `network_id` - (Optional) `network_id` from [networks](mws_networks.md).
* `cloud_resource_container` - A block that specifies GCP workspace configurations, consisting of following blocks:
  * `gcp` - A block that consists of the following field:
    * `project_id` - The Google Cloud project ID, which the workspace uses to instantiate cloud resources for your workspace.
* `gke_config` - A block that specifies GKE configuration for the Databricks workspace:
  * `connectivity_type`: Specifies the network connectivity types for the GKE nodes and the GKE master network. Possible values are: `PRIVATE_NODE_PUBLIC_MASTER`, `PUBLIC_NODE_PUBLIC_MASTER`
  * `master_ip_range`: The IP range from which to allocate GKE cluster master resources. This field will be ignored if GKE private cluster is not enabled. It must be exactly as big as `/28`.

## token block

You can specify a `token` block in the body of the workspace resource, so that Terraform manages the refresh of the PAT token for the deployment user. The other option is to create [databricks_obo_token](obo_token.md), though it requires Premium or Enterprise plan enabled as well as more complex setup. Token block exposes `token_value`, that holds sensitive PAT token and optionally it can accept two arguments:

-> **Note** Tokens managed by `token {}` block are recreated when expired.

* `comment` - (Optional) Comment, that will appear in "User Settings / Access Tokens" page on Workspace UI. By default it's "Terraform PAT".
* `lifetime_seconds` - (Optional) Token expiry lifetime. By default its 2592000 (30 days).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the workspace.
* `workspace_status_message` - (String) updates on workspace status
* `workspace_status` - (String) workspace status
* `creation_time` - (Integer) time when workspace was created
* `workspace_url` - (String) URL of the workspace

## Import

-> **Note** Importing this resource is not currently supported.

## Timeouts

The `timeouts` block allows you to specify `create`, `read` and `update` timeouts. It usually takes 5-7 minutes to provision Databricks E2 Workspace and another couple of minutes for your local DNS caches to resolve. Please launch `TF_LOG=DEBUG terraform apply` whenever you observe timeout issues.

```hcl
timeouts {
  create = "30m"
  read   = "10m"
  update = "20m
}
```

You can reset local DNS caches before provisioning new workspaces with one of the following commands:

* Linux - `sudo /etc/init.d/nscd restart`
* Mac OS Sierra, X El Capitan, X Mavericks, X Mountain Lion, or X Lion - `sudo killall -HUP mDNSResponder`
* Mac OS X Yosemite - `sudo discoveryutil udnsflushcaches`
* Mac OS X Snow Leopard - `sudo dscacheutil -flushcache`
* Mac OS X Leopard and below - `sudo lookupd -flushcache`

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on GCP](../guides/gcp-workspace.md) guide.
* [databricks_mws_networks](mws_networks_gcp.md) to [configure VPC](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) & subnet for new workspaces within GCP.
