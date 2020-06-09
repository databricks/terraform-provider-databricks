+++
title = "mws_networks"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_mws_networks`

This resource to configure the vpc for the multiple workspaces api if the BYOVPC option is chosen.

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
````
## Argument Reference

The following arguments are supported:

#### - `account_id`:
> **(Required)** Databricks multi-workspace master account ID.

#### - `network_name`:
> **(Required)** The human-readable name of the network configuration.
                 
#### - `vpc_id`:
> **(Required)** The ID of the VPC associated with this network. VPC IDs can be used in multiple network configurations.

#### - `subnet_ids`:
> **(Required)** IDs of at least 2 subnets associated with this network. 
>Subnet IDs cannot be used in multiple network configurations.

#### - `security_group_ids`:
> **(Required)** IDs of 1 to 5 security groups associated with this network. 
>Security groups IDs cannot be used in multiple network configurations.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id of the resource which follows the format accountId/networksId.

#### - `vpc_status`:
> Enum: "UNATTACHED" "VALID" "BROKEN" "WARNED"
> The status of this network configuration object in terms of its use in a workspace:
> 
> UNATTACHED — Unattached.
> VALID — Valid.
> BROKEN — Broken.
> WARNED — Warned.

#### - `error_messages`:
> Array of error messages about the network configuration.
> Contains the following objects: 
> error_type: The AWS resource associated with this error: credentials, VPC, subnet, security group, or network ACL.
> error_message: Details of the error.

#### - `workspace_id`:
> Workspace ID associated with this network configuration. Can be empty.

#### - `creation_time`:
> Time in epoch milliseconds when the network was created.

#### - `network_id`:
> The Databricks network configuration ID.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
