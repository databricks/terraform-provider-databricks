---
page_title: "Provisioning Databricks workspaces on GCP"
---

# Provisioning Databricks workspaces on GCP

You can provision multiple Databricks workspaces with Terraform.

## Creating a GCP service account for Databricks Provisioning

This guide assumes that you are already familiar with Hashicorp Terraform and provisioned some of the Google Compute Cloud infrastructure with it. To work with Databricks in GCP in an automated way, please create a service account and manually add it in the [Accounts Console](https://accounts.gcp.databricks.com/users) as an account admin. You can use the following Terraform configuration to create a Service Account for Databricks Provisioning, which can be impersonated by a list of principals defined in delegate_from variable. Service Account would be automatically assigned to the newly created Databricks Workspace Creator custom role

```hcl
variable "prefix" {}
 
variable "project" {
  type    = string
  default = "<my-project-id>"
}
 
provider "google" {
  project = var.project
}
 
variable "delegate_from" {
 description = "Allow either user:user.name@example.com, group:deployers@example.com or serviceAccount:sa1@project.iam.gserviceaccount.com to impersonate created service account"
 type        = list(string)
}
 
resource "google_service_account" "sa2" {
 account_id   = "${var.prefix}-sa2"
 display_name = "Service Account for Databricks Provisioning"
}
 
output "service_account" {
 value       = google_service_account.sa2.email
 description = "Add this email as a user in the Databricks account console"
}
 
data "google_iam_policy" "this" {
 binding {
   role    = "roles/iam.serviceAccountTokenCreator"
   members = var.delegate_from
 }
}
 
resource "google_service_account_iam_policy" "impersonatable" {
 service_account_id = google_service_account.sa2.name
 policy_data        = data.google_iam_policy.this.policy_data
}
 
resource "google_project_iam_custom_role" "workspace_creator" {
 role_id = "${var.prefix}_workspace_creator"
 title   = "Databricks Workspace Creator"
 permissions = [
   "iam.serviceAccounts.getIamPolicy",
   "iam.serviceAccounts.setIamPolicy",
   "iam.roles.create",
   "iam.roles.delete",
   "iam.roles.get",
   "iam.roles.update",
   "resourcemanager.projects.get",
   "resourcemanager.projects.getIamPolicy",
   "resourcemanager.projects.setIamPolicy",
   "serviceusage.services.get",
   "serviceusage.services.list",
   "serviceusage.services.enable"
 ]
}
 
data "google_client_config" "current" {}
 
output "custom_role_url" {
 value = "https://console.cloud.google.com/iam-admin/roles/details/projects%3C${data.google_client_config.current.project}%3Croles%3C${google_project_iam_custom_role.workspace_creator.role_id}"
}
 
resource "google_project_iam_member" "sa2_can_create_workspaces" {
 role   = google_project_iam_custom_role.workspace_creator.id
 member = "serviceAccount:${google_service_account.sa2.email}"
}
```

After you’ve added Service Account to Databricks Accounts Console, please copy its name into `databricks_google_service_account` variable. If you prefer environment variables - `DATABRICKS_GOOGLE_SERVICE_ACCOUNT` is the one you’ll use instead. Please also copy Account ID into `databricks_account_id` variable.

## Authenticate with Databricks account API

Databricks account-level APIs can only be called by account owners and account admins, and can only be authenticated using Google-issued OIDC tokens. The simplest way to do this would be via [Google Cloud CLI](https://cloud.google.com/sdk/gcloud). The `gcloud` command is available after installing the SDK. Then run the following commands

* `gcloud auth application-default login` to authorise your user with Google Cloud Platform.
* `terraform init` to load Google and Databricks Terraform providers.
* `terraform apply` to apply the configuration changes. Terraform will use your credential to impersonate the service account specified in `databricks_google_service_account` to call the Databricks account-level API.

Alternatively, if you cannot use impersonation and [Application Default Credentials](https://cloud.google.com/docs/authentication/production) as configured by `gcloud`, consider using the service account key directly by passing it to `google_credentials` parameter (or `GOOGLE_CREDENTIALS` environment variable) to avoid using `gcloud`, impersonation, and ADC altogether. The content of this parameter must be either the path to `.json` file or the full JSON content of the Google service account key.

## Provider initialization

```hcl
variable "databricks_account_id" {}
variable "databricks_google_service_account" {}
variable "google_project" {}
variable "google_region" {}
variable "google_zone" {}


terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
    google = {
      source  = "hashicorp/google"
      version = "4.47.0"
    }
  }
}

provider "google" {
  project = var.google_project
  region  = var.google_region
  zone    = var.google_zone
}

// initialize provider in "accounts" mode to provision new workspace

provider "databricks" {
  alias                  = "accounts"  
  host                   = "https://accounts.gcp.databricks.com"
  google_service_account = var.databricks_google_service_account
  account_id             = var.databricks_account_id
}

data "google_client_openid_userinfo" "me" {
}

data "google_client_config" "current" {
}

resource "random_string" "suffix" {
  special = false
  upper   = false
  length  = 6
}
```

## Creating a VPC

The very first step is VPC creation with necessary resources. Please consult [main documentation page](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) for **the most complete and up-to-date details on networking**. A GCP VPC is registered as [databricks_mws_networks](../resources/mws_networks_gcp.md) resource.

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
}
```

## Creating a Databricks Workspace

Once [the VPC](#Creating a VPC) is set up, you can create Databricks workspace through [databricks_mws_workspaces](../resources/mws_workspaces_gcp.md) resource.

Code that creates workspaces and code that [manages workspaces](workspace-management.md) must be in separate terraform modules to avoid common confusion between `provider = databricks.accounts` and `provider = databricks.created_workspace`. This is why we specify `databricks_host` and `databricks_token` outputs, which have to be used in the latter modules.

-> **Note** If you experience technical difficulties with rolling out resources in this example, please make sure that [environment variables](../index.md#environment-variables) don't [conflict with other](../index.md#empty-provider-block) provider block attributes. When in doubt, please run `TF_LOG=DEBUG terraform apply` to enable [debug mode](https://www.terraform.io/docs/internals/debugging.html) through the [`TF_LOG`](https://www.terraform.io/docs/cli/config/environment-variables.html#tf_log) environment variable. Look specifically for `Explicit and implicit attributes` lines, that should indicate authentication attributes used. The other common reason for technical difficulties might be related to missing `alias` attribute in `provider "databricks" {}` blocks or `provider` attribute in `resource "databricks_..." {}` blocks. Please make sure to read [`alias`: Multiple Provider Configurations](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) documentation article.

```hcl
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

  network_id = databricks_mws_networks.this.network_id
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

*In Terraform 0.13 and later*, data resources have the same dependency resolution behavior [as defined for managed resources](https://www.terraform.io/docs/language/resources/behavior.html#resource-dependencies). Most data resources make an API call to a workspace. If a workspace doesn't exist yet, `authentication is not configured for provider` error is raised. To work around this issue and guarantee a proper lazy authentication with data resources, you should add `depends_on = [databricks_mws_workspaces.this]` to the body. This issue doesn't occur if workspace is created *in one module* and resources [within the workspace](workspace-management.md) are created *in another*. We do not recommend using Terraform 0.12 and earlier, if your usage involves data resources.

```hcl
data "databricks_current_user" "me" {
  depends_on = [databricks_mws_workspaces.this]
}
```

## Provider configuration

In [the next step](workspace-management.md), please use the following configuration for the provider:

```hcl
provider "databricks" {
  host  = module.dbx_gcp.workspace_url
  token = module.dbx_gcp.token_value
}
```

We assume that you have a terraform module in your project that creats a workspace (using [Databricks Workspace](#creating-a-databricks-workspace) section) and you named it as `dbx_gcp` while calling it in the **main.tf** file of your terraform project. And `workspace_url` and `token_value` are the output attributes of that module. This provider configuration will allow you to use the generated token during workspace creation to authenticate to the created workspace.
