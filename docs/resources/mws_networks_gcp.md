---
subcategory: "GCP"
---
# databricks_mws_networks Resource

-> **Note** Initialize provider with `alias = "mws"`, `host  = "https://accounts.gcp.databricks.com"` and use `provider = databricks.mws` for all `databricks_mws_*` resources.

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

Use this resource to [configure VPC](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) & subnet for new workspaces within GCP. It is essential to understand that this will require you to configure your provider separately for the multiple workspaces resources.

* Databricks must have access to a subnet in the same region as the workspace, of which IP range will be used to allocate your workspace’s GKE cluster nodes.
* The subnet must have a netmask between /29 and /9.
* Databricks must have access to 2 secondary IP ranges, one between /21 to /9 for workspace’s GKE cluster pods, and one between /27 to /16 for workspace’s GKE cluster services.
* Subnet must have outbound access to the public network using a [gcp_compute_router_nat](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_router_nat) or other similar customer-managed appliance infrastructure.

Please follow this [complete runnable example](../guides/gcp-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.gcp.databricks.com/"`) and another for the workspace you've created with `databricks_mws_workspaces` resource. If you want both creations of workspaces & clusters within the same Terraform module (essentially the same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module to create workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

resource "google_compute_network" "dbx_private_vpc" {
  project                 = var.google_project
  name                    = "tf-network-${random_string.suffix.result}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-dbx-${random_string.suffix.result}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.dbx_private_vpc.id
  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = "10.1.0.0/16"
  }
  secondary_ip_range {
    range_name    = "svc"
    ip_cidr_range = "10.2.0.0/20"
  }
  private_ip_google_access = true
}

resource "google_compute_router" "router" {
  name    = "my-router-${random_string.suffix.result}"
  region  = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  network = google_compute_network.dbx_private_vpc.id
}

resource "google_compute_router_nat" "nat" {
  name                               = "my-router-nat-${random_string.suffix.result}"
  router                             = google_compute_router.router.name
  region                             = google_compute_router.router.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}

resource "databricks_mws_networks" "this" {
  account_id   = var.databricks_account_id
  network_name = "test-demo-${random_string.suffix.result}"
  gcp_network_info {
    network_project_id    = var.google_project
    vpc_id                = google_compute_network.dbx_private_vpc.name
    subnet_id             = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
    subnet_region         = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
    pod_ip_range_name     = "pods"
    service_ip_range_name = "svc"
  }
}
```

## Argument Reference

The following arguments are available:

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `network_name` - name under which this network is registered
* `gcp_network_info` - a block consists of Google Cloud specific information for this network, for example the VPC ID, subnet ID, and secondary IP ranges. It has the following fields:
  * `network_project_id` - The Google Cloud project ID of the VPC network.
  * `vpc_id` - The ID of the VPC associated with this network. VPC IDs can be used in multiple network configurations.
  * `subnet_id` - The ID of the subnet associated with this network.
  * `subnet_region` - The Google Cloud region of the workspace data plane. For example, `us-east4`.
  * `pod_ip_range_name` - The name of the secondary IP range for pods. A Databricks-managed GKE cluster uses this IP range for its pods. This secondary IP range can only be used by one workspace.
  * `service_ip_range_name` - The name of the secondary IP range for services. A Databricks-managed GKE cluster uses this IP range for its services. This secondary IP range can only be used by one workspace.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws networks.
* `network_id` - (String) id of network to be used for [databricks_mws_workspace](mws_workspaces_gcp.md) resource.
* `vpc_status` - (String) VPC attachment status
* `workspace_id` - (Integer) id of associated workspace

## Import

-> **Note** Importing this resource is not currently supported.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on GCP](../guides/gcp-workspace.md) guide.
* [databricks_mws_workspaces](mws_workspaces_gcp.md) to set up [workspaces on GCP](https://docs.gcp.databricks.com/administration-guide/account-settings-gcp/workspaces.html).
