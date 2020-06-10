+++
title = "mws_workspaces"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_mws_workspaces`

This resource to configure the vpc for the multiple workspaces api if the BYOVPC option is chosen.

{{% notice warning %}}
This provider does not yet support the customer_managed_key resource yet so you will need to manually create that 
and provide the cmk object guid into the workspace api. You can see it on the argument reference below.  
{{% /notice %}}

{{% notice warning %}}
It is important to understand that this will require you to configure your provider separately for the 
multiple workspaces resources
{{% /notice %}}

{{% notice note %}}
This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth 
as that is the only authentication method available for multiple workspaces api.
{{% /notice %}}


## Example Usage

````hcl
provider "databricks" {
  host = "https://accounts.cloud.databricks.com"
  basic_auth {
    username = "username"
    password = "password"
  }
}
resource "databricks_mws_credentials" "my_mws_credentials" {
  account_id       = "my-mws-account-id"
  credentials_name = "my-cusotom-credentials"
  role_arn         = "arn:aws:iam::9999999999999:role/my-custom-cross-account-role"
}
resource "databricks_mws_storage_configurations" "my_mws_storage_configurations" {
  account_id = "my-mws-acct-id"
  storage_configuration_name = "storage-configuration-name"
  bucket_name         = "my-root-s3-bucket"
}
resource "databricks_mws_networks" "my_network" {
  account_id = "my-mws-acct-id"
  network_name = "my-custom-network-config-name"
  vpc_id = "my-aws-vpc-id"
  subnet_ids = [
    "my-first-subnet",
    "my-second-subnet",
  ]
  security_group_ids = [
    "my-security-group-1",
  ]
}
resource "databricks_mws_workspaces" "my_mws_workspace" {
  account_id = "my-mws-acct-id"
  workspace_name = "my-workspace-name"
  deployment_name = "my-deployment-urlname"
  aws_region = "my-aws-region"
  credentials_id = databricks_mws_credentials.my_mws_credentials.credentials_id
  storage_configuration_id = databricks_mws_storage_configurations.my_mws_storage_configurations.storage_configuration_id
  network_id = databricks_mws_networks.my_network.network_id
  verify_workspace_runnning = true
}
````
## Argument Reference

The following arguments are supported:

#### - `account_id`:
> **(Required)** Databricks multi-workspace master account ID.

#### - `workspace_name`:
> **(Required)** The workspace's human-readable name. It is used as part of the workspace URL.
                 
#### - `deployment_name`:
> **(Required)** The name of the deployment you want. The URL prefix of the workspace. 
>Append .cloud.databricks.com to get the full URL.

#### - `aws_region`:
> **(Required)** The AWS region of the workspace's Data Plane.

#### - `credentials_id`:
> **(Required)** ID of the workspace's credential configuration object

#### - `storage_configuration_id`:
> **(Required)** The ID of the workspace's storage configuration object.

#### - `verify_workspace_runnning`:
> **(Required)** Validates that the workspace is functioning post creation. Recommended to turn this on 
>to verify post apply that the workspace is in a running conidition

#### - `network_id`:
> **(Optional)** The ID of the workspace's network configuration object.

#### - `customer_managed_key_id`:
> **(Optional)** The ID of the workspace's notebook encryption key configuration object.

#### - `is_no_public_ip_enabled`:
> **(Optional)** Specifies whether secure cluster connectivity (sometimes called no public IP) is enabled on this workspace.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id of the resource which follows the format accountId/workspaceId.

#### - `workspace_status`:
> Enum: "NOT_PROVISIONED" "PROVISIONING" "RUNNING" "FAILED" "BANNED" "CANCELLING"
>The status of the workspace. For workspace creation, it is typically initially PROVISIONING. 
>Continue to check the status until the status is RUNNING. 
>For detailed instructions of creating a new workspace with this API including error handling see 
>Create a new workspace with the Multi-workspace API.

#### - `workspace_status_message`:
> Message describing the current workspace status.

#### - `creation_time`:
> Time in epoch milliseconds when the workspace was created.

#### - `workspace_id`:
> The Databricks workspace ID.

#### - `workspace_url`:
> The URL for the workspace.

#### - `network_error_messages`:
> Array of error messages about the network configuration.
> Contains the following objects: 
> error_type: The AWS resource associated with this error: credentials, VPC, subnet, security group, or network ACL.
> error_message: Details of the error.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
