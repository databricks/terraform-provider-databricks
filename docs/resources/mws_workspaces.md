# databricks_mws_workspaces Resource

-> **Note** This resource has evolving API, which may change in future versions of provider.

This resource to configure new workspaces within AWS.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api. 

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

// register cross-account ARN
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = var.crossaccount_arn
}

// register root bucket
resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.account_id
  storage_configuration_name = "${var.prefix}-storage"
  bucket_name                = var.root_bucket
}

// register VPC
resource "databricks_mws_networks" "this" {
  provider     = databricks.mws
  account_id   = var.account_id
  network_name = "${var.prefix}-network"
  vpc_id       = var.vpc_id
  subnet_ids = [var.subnet_public, var.subnet_private]
  security_group_ids = [var.security_group]
}

// create workspace in given VPC with DBFS on root bucket
resource "databricks_mws_workspaces" "this" {
  provider        = databricks.mws
  account_id      = var.account_id
  workspace_name  = var.prefix
  deployment_name = var.prefix
  aws_region      = var.region

  credentials_id            = databricks_mws_credentials.this.credentials_id
  storage_configuration_id  = databricks_mws_storage_configurations.this.storage_configuration_id
  network_id                = databricks_mws_networks.this.network_id
  verify_workspace_runnning = true
}

provider "databricks" {
  // in normal scenario you won't have to give providers aliases
  alias = "created_workspace" 
  
  host  = databricks_mws_workspaces.this.workspace_url
}

// create PAT token to provision entities within workspace
resource "databricks_token" "pat" {
  provider = databricks.created_workspace
  comment  = "Terraform Provisioning"
  // 1 day token
  lifetime_seconds = 86400
}
```

## Argument Reference

The following arguments are required:

* `network_id` - (Optional) (String) `network_id` from [networks](mws_networks.md)
* `account_id` - (Required) (String) master account id (also used for `sts:ExternaId` of `sts:AssumeRole`)
* `credentials_id` - (Required) (String) `credentials_id` from [credentials](mws_credentials.md)
* `customer_managed_key_id` - (Optional) (String) `customer_managed_key_id` from [customer managed keys](mws_customer_managed_keys.md)
* `deployment_name` - (Required) (String) part of URL: `https://<deployment-name>.cloud.databricks.com`
* `workspace_name` - (Required) (String) name of the workspace, will appear on UI
* `aws_region` - (Required) (String) AWS region of VPC
* `storage_configuration_id` - (Required) (String) `storage_configuration_id` from [storage configuration](mws_storage_configurations.md)
* `verify_workspace_runnning` - (Required) (Bool) wait until workspace is running.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the workspace.
* `workspace_status_message` - (String) updates on workspace status
* `workspace_status` - (String) workspace status
* `creation_time` - (Integer) time when workspace was created
* `workspace_url` - (String) URL of the workspace
* `workspace_id` - (Integer) same as `id`
