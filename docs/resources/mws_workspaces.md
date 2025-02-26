---
subcategory: "Deployment"
---
# databricks_mws_workspaces resource

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`. We require all `databricks_mws_*` resources to be created within its own dedicated terraform module of your environment. Usually this module creates VPC and IAM roles as well. Code that creates workspaces and code that [manages workspaces](../guides/workspace-management.md) must be in separate terraform modules to avoid common confusion between `provider = databricks.mws` and `provider = databricks.created_workspace`. This is why we specify `databricks_host` and `databricks_token` outputs, that have to be used in the latter modules:

```hcl
provider "databricks" {
  host  = module.ai.databricks_host
  token = module.ai.databricks_token
}
```

This resource allows you to set up [workspaces on AWS](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1) or [workspaces on GCP](https://docs.gcp.databricks.com/administration-guide/account-settings-gcp/workspaces.html). Please follow this complete runnable example on [AWS](../guides/aws-workspace.md) or [GCP](../guides/gcp-workspace.md) with new VPC and new workspace setup.

-> On Azure you need to use [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace) resource to create Azure Databricks workspaces.

## Example Usage

### Creating a Databricks on AWS workspace

![Simplest multiworkspace](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/simplest-multiworkspace.png)

To get workspace running, you have to configure a couple of things:

* [databricks_mws_credentials](mws_credentials.md) - You can share a credentials (cross-account IAM role) configuration ID with multiple workspaces. It is not required to create a new one for each workspace.
* [databricks_mws_storage_configurations](mws_storage_configurations.md) - You can share a root S3 bucket with multiple workspaces in a single account. You do not have to create new ones for each workspace. If you share a root S3 bucket for multiple workspaces in an account, data on the root S3 bucket is partitioned into separate directories by workspace.
* [databricks_mws_networks](mws_networks.md) - (optional, but recommended) You can share one [customer-managed VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) with multiple workspaces in a single account. However, Databricks recommends using unique subnets and security groups for each workspace. If you plan to share one VPC with multiple workspaces, be sure to size your VPC and subnets accordingly. Because a Databricks [databricks_mws_networks](mws_networks.md) encapsulates this information, you cannot reuse it across workspaces.
* [databricks_mws_customer_managed_keys](mws_customer_managed_keys.md) - You can share a customer-managed key across workspaces.

```hcl
variable "databricks_account_id" {
  description = "Account ID that can be found in the dropdown under the email address in the upper-right corner of https://accounts.cloud.databricks.com/"
}

provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

// register cross-account ARN
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.databricks_account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = var.crossaccount_arn
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  storage_configuration_name = "${var.prefix}-storage"
  bucket_name                = var.root_bucket
}

// register VPC
resource "databricks_mws_networks" "this" {
  provider           = databricks.mws
  account_id         = var.databricks_account_id
  network_name       = "${var.prefix}-network"
  vpc_id             = var.vpc_id
  subnet_ids         = var.subnets_private
  security_group_ids = [var.security_group]
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider       = databricks.mws
  account_id     = var.databricks_account_id
  workspace_name = var.prefix
  aws_region     = var.region

  credentials_id           = databricks_mws_credentials.this.credentials_id
  storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id               = databricks_mws_networks.this.network_id

  token {}
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

### Creating a Databricks on AWS workspace with Databricks-Managed VPC

![VPCs](https://docs.databricks.com/_images/customer-managed-vpc.png)

By default, Databricks creates a VPC in your AWS account for each workspace. Databricks uses it for running clusters in the workspace. Optionally, you can use your VPC for the workspace, using the feature customer-managed VPC. Databricks recommends that you provide your VPC with [databricks_mws_networks](mws_networks.md) so that you can configure it according to your organization’s enterprise cloud standards while still conforming to Databricks requirements. You cannot migrate an existing workspace to your VPC. Please see the difference described through IAM policy actions [on this page](https://docs.databricks.com/administration-guide/account-api/iam-role.html).

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix = "dltp${random_string.naming.result}"
}

data "databricks_aws_assume_role_policy" "this" {
  external_id = var.databricks_account_id
}

resource "aws_iam_role" "cross_account_role" {
  name               = "${local.prefix}-crossaccount"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  tags               = var.tags
}

data "databricks_aws_crossaccount_policy" "this" {
}

resource "aws_iam_role_policy" "this" {
  name   = "${local.prefix}-policy"
  role   = aws_iam_role.cross_account_role.id
  policy = data.databricks_aws_crossaccount_policy.this.json
}

resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.databricks_account_id
  credentials_name = "${local.prefix}-creds"
  role_arn         = aws_iam_role.cross_account_role.arn
}

resource "aws_s3_bucket" "root_storage_bucket" {
  bucket        = "${local.prefix}-rootbucket"
  acl           = "private"
  force_destroy = true
  tags          = var.tags
}

resource "aws_s3_bucket_versioning" "root_versioning" {
  bucket = aws_s3_bucket.root_storage_bucket.id
  versioning_configuration {
    status = "Disabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "root_storage_bucket" {
  bucket = aws_s3_bucket.root_storage_bucket.bucket

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "root_storage_bucket" {
  bucket                  = aws_s3_bucket.root_storage_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
  depends_on              = [aws_s3_bucket.root_storage_bucket]
}

data "databricks_aws_bucket_policy" "this" {
  bucket = aws_s3_bucket.root_storage_bucket.bucket
}

resource "aws_s3_bucket_policy" "root_bucket_policy" {
  bucket     = aws_s3_bucket.root_storage_bucket.id
  policy     = data.databricks_aws_bucket_policy.this.json
  depends_on = [aws_s3_bucket_public_access_block.root_storage_bucket]
}

resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  storage_configuration_name = "${local.prefix}-storage"
  bucket_name                = aws_s3_bucket.root_storage_bucket.bucket
}

resource "databricks_mws_workspaces" "this" {
  provider       = databricks.mws
  account_id     = var.databricks_account_id
  workspace_name = local.prefix
  aws_region     = "us-east-1"

  credentials_id           = databricks_mws_credentials.this.credentials_id
  storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id

  token {}

  # Optional Custom Tags
  custom_tags = {

    "SoldToCode" = "1234"

  }
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

In order to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) please ensure that you have read and understood the [Enable Private Link](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html) documentation and then customise the example above with the relevant examples from [mws_vpc_endpoint](mws_vpc_endpoint.md), [mws_private_access_settings](mws_private_access_settings.md) and [mws_networks](mws_networks.md).

### Creating a Databricks on GCP workspace

To get workspace running, you have to configure a network object:

* [databricks_mws_networks](mws_networks.md) - (optional, but recommended) You can share one [customer-managed VPC](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/customer-managed-vpc.html) with multiple workspaces in a single account. You do not have to create a new VPC for each workspace. However, you cannot reuse subnets with other resources, including other workspaces or non-Databricks resources. If you plan to share one VPC with multiple workspaces, be sure to size your VPC and subnets accordingly. Because a Databricks [databricks_mws_networks](mws_networks.md) encapsulates this information, you cannot reuse it across workspaces.

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
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
  token {}
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

In order to create a [Databricks Workspace that leverages GCP Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) please ensure that you have read and understood the [Enable Private Service Connect](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/private-service-connect.html) documentation and then customise the example above with the relevant examples from [mws_vpc_endpoint](mws_vpc_endpoint.md), [mws_private_access_settings](mws_private_access_settings.md) and [mws_networks](mws_networks.md).

#### Creating a Databricks on GCP workspace with Databricks-Managed VPC

![VPCs](https://docs.databricks.com/_images/customer-managed-vpc.png)

By default, Databricks creates a VPC in your GCP project for each workspace. Databricks uses it for running clusters in the workspace. Optionally, you can use your VPC for the workspace, using the feature customer-managed VPC. Databricks recommends that you provide your VPC with [databricks_mws_networks](mws_networks.md) so that you can configure it according to your organization’s enterprise cloud standards while still conforming to Databricks requirements. You cannot migrate an existing workspace to your VPC.

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
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

  token {}
}

output "databricks_token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}
```

## Argument Reference

-> All workspaces would be verified to get into runnable state or deleted upon failure. You can only update `credentials_id`, `network_id`, and `storage_customer_managed_key_id`, `private_access_settings_id` on a running workspace.

The following arguments are available:

* `account_id` - Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/).
* `deployment_name` - (Optional) part of URL as in `https://<prefix>-<deployment-name>.cloud.databricks.com`. Deployment name cannot be used until a deployment name prefix is defined. Please contact your Databricks representative. Once a new deployment prefix is added/updated, it only will affect the new workspaces created.
* `workspace_name` - name of the workspace, will appear on UI.
* `network_id` - (Optional) `network_id` from [networks](mws_networks.md).
* `aws_region` - (AWS only) region of VPC.
* `storage_configuration_id` - (AWS only)`storage_configuration_id` from [storage configuration](mws_storage_configurations.md).
* `managed_services_customer_managed_key_id` - (Optional) `customer_managed_key_id` from [customer managed keys](mws_customer_managed_keys.md) with `use_cases` set to `MANAGED_SERVICES`. This is used to encrypt the workspace's notebook and secret data in the control plane.
* `storage_customer_managed_key_id` - (Optional) `customer_managed_key_id` from [customer managed keys](mws_customer_managed_keys.md) with `use_cases` set to `STORAGE`. This is used to encrypt the DBFS Storage & Cluster Volumes.
* `location` - (GCP only) region of the subnet.
* `cloud_resource_container` - (GCP only) A block that specifies GCP workspace configurations, consisting of following blocks:
  * `gcp` - A block that consists of the following field:
    * `project_id` - The Google Cloud project ID, which the workspace uses to instantiate cloud resources for your workspace.
* `gke_config` (Deprecated) - (GCP only) This field is deprecated and will be removed in a future release. A block that specifies GKE configuration for the Databricks workspace:
  * `connectivity_type`: Specifies the network connectivity types for the GKE nodes and the GKE master network. Possible values are: `PRIVATE_NODE_PUBLIC_MASTER`, `PUBLIC_NODE_PUBLIC_MASTER`.
  * `master_ip_range`: The IP range from which to allocate GKE cluster master resources. This field will be ignored if GKE private cluster is not enabled. It must be exactly as big as `/28`.
* `private_access_settings_id` - (Optional) Canonical unique identifier of [databricks_mws_private_access_settings](mws_private_access_settings.md) in Databricks Account.
* `custom_tags` - (Optional / AWS only) - The custom tags key-value pairing that is attached to this workspace. These tags will be applied to clusters automatically in addition to any `default_tags` or `custom_tags` on a cluster level. Please note it can take up to an hour for custom_tags to be set due to scheduling on Control Plane. After custom tags are applied, they can be modified however they can never be completely removed.
* `pricing_tier` - (Optional) - The pricing tier of the workspace.

### token block

You can specify a `token` block in the body of the workspace resource, so that Terraform manages the refresh of the PAT token for the deployment user. The other option is to create [databricks_obo_token](obo_token.md), though it requires Premium or Enterprise plan enabled as well as more complex setup. Token block exposes `token_value`, that holds sensitive PAT token and optionally it can accept two arguments:

-> Tokens managed by `token {}` block are recreated when expired.

* `comment` - (Optional) Comment, that will appear in "User Settings / Access Tokens" page on Workspace UI. By default it's "Terraform PAT".
* `lifetime_seconds` - (Optional) Token expiry lifetime. By default its 2592000 (30 days).

### Updating workspaces

On AWS, the following arguments could be modified after the workspace is running:

* `network_id` - Modifying [networks on running workspaces](mws_networks.md#modifying-networks-on-running-workspaces-aws-only) would require three separate `terraform apply` steps.
* `credentials_id`
* `storage_customer_managed_key_id`
* `private_access_settings_id`
* `custom_tags`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - (String) Canonical unique identifier for the workspace, of the format `<account-id>/<workspace-id>`
* `workspace_id` - (String) workspace id
* `workspace_status_message` - (String) updates on workspace status
* `workspace_status` - (String) workspace status
* `creation_time` - (Integer) time when workspace was created
* `workspace_url` - (String) URL of the workspace
* `custom_tags` - (Map) Custom Tags (if present) added to workspace
* `gcp_workspace_sa` - (String, GCP only) identifier of a service account created for the workspace in form of `db-<workspace-id>@prod-gcp-<region>.iam.gserviceaccount.com`

## Timeouts

The `timeouts` block allows you to specify `create`, `read` and `update` timeouts. It usually takes 5-7 minutes to provision Databricks workspace and another couple of minutes for your local DNS caches to resolve. Please launch `TF_LOG=DEBUG terraform apply` whenever you observe timeout issues.

```hcl
timeouts {
  create = "30m"
  read   = "10m"
  update = "20m"
}
```

You can reset local DNS caches before provisioning new workspaces with one of the following commands:

* Linux - `sudo /etc/init.d/nscd restart`
* Mac OS Sierra, X El Capitan, X Mavericks, X Mountain Lion, or X Lion - `sudo killall -HUP mDNSResponder`
* Mac OS X Yosemite - `sudo discoveryutil udnsflushcaches`
* Mac OS X Snow Leopard - `sudo dscacheutil -flushcache`
* Mac OS X Leopard and below - `sudo lookupd -flushcache`

## Import

This resource can be imported by Databricks account ID and workspace ID.

```sh
terraform import databricks_mws_networks.this '<account_id>/<workspace_id>'
```

~> Not all fields of `databricks_mws_workspaces` can be updated without causing the workspace to be recreated.
   If the configuration for these immutable fields does not match the existing workspace, the workspace will
   be deleted and recreated in the next `terraform apply`. After importing, verify that the configuration
   matches the existing resource by running `terraform plan`. The only fields that can be updated are
   `credentials_id`, `network_id`, `storage_customer_managed_key_id`, `private_access_settings_id`,
   `managed_services_customer_managed_key_id`, and `custom_tags`.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with Private Link](../guides/aws-private-link-workspace.md) guide.
* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide.
* [Provisioning Databricks on GCP](../guides/gcp-workspace.md) guide.
* [Provisioning Databricks workspaces on GCP with Private Service Connect](../guides/gcp-private-service-connect-workspace.md) guide.
* [databricks_mws_credentials](mws_credentials.md) to configure the cross-account role for creation of new workspaces within AWS.
* [databricks_mws_customer_managed_keys](mws_customer_managed_keys.md) to configure KMS keys for new workspaces within AWS.
* [databricks_mws_log_delivery](mws_log_delivery.md) to configure delivery of [billable usage logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html) and [audit logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_storage_configurations](mws_storage_configurations.md) to configure root bucket new workspaces within AWS.
* [databricks_mws_private_access_settings](mws_private_access_settings.md) to create a [Private Access Setting](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html#step-5-create-a-private-access-settings-configuration-using-the-databricks-account-api) that can be used as part of a [databricks_mws_workspaces](mws_workspaces.md) resource to create a [Databricks Workspace that leverages AWS PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
