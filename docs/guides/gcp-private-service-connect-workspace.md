---
page_title: "Provisioning Databricks on Google Cloud with Private Service Connect"
---

# Provisioning Databricks workspaces on GCP with Private Service Connect

-> **Note** Refer to the [Databricks Terraform Registry modules](https://registry.terraform.io/modules/databricks/examples/databricks/latest) for Terraform modules and examples to deploy Azure Databricks resources.

Secure a workspace with private connectivity and mitigate data exfiltration risks by [enabling Google Private Service Connect (PSC) on the workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html). This guide assumes that you are already familiar with Hashicorp Terraform and provisioned some of the Google Compute Cloud infrastructure with it.

## Creating a GCP service account for Databricks Provisioning and Authenticate with Databricks account API

To work with Databricks in GCP in an automated way, please create a service account and manually add it in the [Accounts Console](https://accounts.gcp.databricks.com/users) as an account admin. Databricks account-level APIs can only be called by account owners and account admins, and can only be authenticated using Google-issued OIDC tokens. The simplest way to do this would be via [Google Cloud CLI](https://cloud.google.com/sdk/gcloud). For details, please refer to [Provisioning Databricks workspaces on GCP](gcp_workspace.md).

## Creating a VPC network

The very first step is VPC creation with the necessary resources. Please consult [main documentation page](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. A GCP VPC is registered as [databricks_mws_networks](../resources/mws_networks.md) resource.

To enable [back-end Private Service Connect (data plane to control plane)](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html#two-private-service-connect-options), configure the network with the two back-end VPC endpoints:

- Back-end VPC endpoint for [Secure cluster connectivity](https://docs.gcp.databricks.com/security/secure-cluster-connectivity.html) relay
- Back-end VPC endpoint for REST APIs

-> Note: If you want to implement the front-end VPC endpoint as well for the connections from users to to the Databricks web application, REST API, and Databricks Connect API over a Virtual Private Cloud (VPC) endpoint, use the transit (bastion) VPC. Once the front-end endpoint is created, use the databricks_mws_private_access_settings resource to control which VPC endpoints can connect to the UI or API of any workspace that attaches this private access settings object.

```hcl
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

resource "databricks_mws_vpc_endpoint" "backend_rest_vpce" {
  account_id        = var.databricks_account_id
  vpc_endpoint_name = "vpce-backend-rest-${random_string.suffix.result}"
  gcp_vpc_endpoint_info {
    project_id        = var.google_project
    psc_endpoint_name = var.backend_rest_psce
    endpoint_region   = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  }
}

resource "databricks_mws_vpc_endpoint" "relay_vpce" {
  account_id        = var.databricks_account_id
  vpc_endpoint_name = "vpce-relay-${random_string.suffix.result}"
  gcp_vpc_endpoint_info {
    project_id        = var.google_project
    psc_endpoint_name = var.relay_psce
    endpoint_region   = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  }
}

resource "databricks_mws_networks" "this" {
  provider     = databricks.accounts
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
  vpc_endpoints {
    dataplane_relay = [databricks_mws_vpc_endpoint.relay_vpce.vpc_endpoint_id]
    rest_api        = [databricks_mws_vpc_endpoint.backend_rest_vpce.vpc_endpoint_id]
  }
}
```

## Creating a Databricks Workspace

Once [the VPC](#creating-a-vpc-network) is set up, you can create Databricks workspace through [databricks_mws_workspaces](../resources/mws_workspaces.md) resource.

For a workspace to support any of the Private Service Connect connectivity scenarios, the workspace must be created with an attached [databricks_mws_private_access_settings](../resources/mws_private_access_settings.md) resource.

Code that creates workspaces and code that [manages workspaces](workspace-management.md) must be in separate terraform modules to avoid common confusion between `provider = databricks.accounts` and `provider = databricks.created_workspace`. This is why we specify `databricks_host` and `databricks_token` outputs, which have to be used in the latter modules.

-> **Note** If you experience technical difficulties with rolling out resources in this example, please make sure that [environment variables](../index.md#environment-variables) don't [conflict with other](../index.md#empty-provider-block) provider block attributes. When in doubt, please run `TF_LOG=DEBUG terraform apply` to enable [debug mode](https://www.terraform.io/docs/internals/debugging.html) through the [`TF_LOG`](https://www.terraform.io/docs/cli/config/environment-variables.html#tf_log) environment variable. Look specifically for `Explicit and implicit attributes` lines, which should indicate authentication attributes used. The other common reason for technical difficulties might be related to missing `alias` attribute in `provider "databricks" {}` blocks or `provider` attribute in `resource "databricks_..." {}` blocks. Please make sure to read [`alias`: Multiple Provider Configurations](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) documentation article.

```hcl
resource "databricks_mws_private_access_settings" "pas" {
  account_id                   = var.databricks_account_id
  private_access_settings_name = "pas-${random_string.suffix.result}"
  region                       = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  public_access_enabled        = true
  private_access_level         = "ACCOUNT"
}

resource "databricks_mws_workspaces" "this" {
  provider       = databricks.accounts
  account_id     = var.databricks_account_id
  workspace_name = "tf-demo-test-${random_string.suffix.result}"
  location       = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  cloud_resource_container {
    gcp {
      project_id = var.google_project
    }
  }

  private_access_settings_id = databricks_mws_private_access_settings.pas.private_access_settings_id
  network_id                 = databricks_mws_networks.this.network_id
  gke_config {
    connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
    master_ip_range   = "10.3.0.0/28"
  }

  token {
    comment = "Terraform"
  }

  # this makes sure that the NAT is created for outbound traffic before creating the workspace
  depends_on = [google_compute_router_nat.nat]
}

output "databricks_host" {
  value = databricks_mws_workspaces.this.workspace_url
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

### Data resources and Authentication is not configured errors

*In Terraform 0.13 and later*, data resources have the same dependency resolution behavior [as defined for managed resources](https://www.terraform.io/docs/language/resources/behavior.html#resource-dependencies). Most data resources make an API call to a workspace. If a workspace doesn't exist yet, `default auth: cannot configure default credentials` error is raised. To work around this issue and guarantee proper lazy authentication with data resources, you should add `depends_on = [databricks_mws_workspaces.this]` to the body. This issue doesn't occur if workspace is created *in one module* and resources [within the workspace](workspace-management.md) are created *in another*. We do not recommend using Terraform 0.12 and earlier if your usage involves data resources.

```hcl
data "databricks_current_user" "me" {
  depends_on = [databricks_mws_workspaces.this]
}
```
