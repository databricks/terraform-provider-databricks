# Azure Databricks workspaces within custom VNet

## main.tf

* Creates Azure Databricks workspace in VNet
* exposes `databricks_azure_workspace_resource_id` and `workspace_url` variables

## init.tf

* Creates resource group within specified location (eastus by default)
* creates random local prefix variable for all resources

## storage.tf

* Creates StorageV2 account, `dev` container in it and exposes account name with access key

## vnet.tf

* Creates virtual network with supplied CIDR
* Creates signle network security group for public and private subnets of a vnet
* Delegates the use of databricks endpoints for use within those subnets