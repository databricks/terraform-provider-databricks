provider "azurerm" {
  version = "~> 2.14"
  features {}
}

provider "random" {
  version = "~> 2.2"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

data "azurerm_client_config" "current" {
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

variable "owner" {
  type = string 
  default = ""
}

locals {
  // dltp - databricks labs terraform provider
  prefix = "dltp${random_string.naming.result}"
  tags = {
    Environment = "Testing"
    Owner       = coalesce(var.owner, lookup(data.external.me.result, "name"))
    Epoch       = random_string.naming.result
  }
}

resource "azurerm_resource_group" "example" {
  name     = "${local.prefix}-rg"
  location = "westeurope"
  tags     = local.tags
}

resource "azurerm_databricks_workspace" "example" {
  name                        = "${local.prefix}-workspace"
  resource_group_name         = azurerm_resource_group.example.name
  location                    = azurerm_resource_group.example.location
  sku                         = "premium"
  managed_resource_group_name = "${local.prefix}-workspace-rg"
  tags                        = local.tags
}

// ADLSv1 resource
resource "azurerm_data_lake_store" "gen1" {
  name                = "${local.prefix}gen1"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
}

# Create container in storage acc and container for use by blob mount tests
resource "azurerm_storage_account" "blobaccount" {
  name                     = "${local.prefix}blob"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
  tags                     = local.tags
}

resource "azurerm_storage_container" "blobexample" {
  name                  = "dev"
  storage_account_name  = azurerm_storage_account.blobaccount.name
  container_access_type = "private"
}

output "cloud_env" {
  value = "azure"
}

output "test_node_type" {
  // or i3.xlarge for AWS
  value = "Standard_D3_v2"
}

output "test_prefix" {
  value = local.prefix
}

output "arm_client_id" {
  value = data.azurerm_client_config.current.client_id
}

output "arm_subscription_id" {
  value = data.azurerm_client_config.current.subscription_id
}

output "arm_tenant_id" {
  value = data.azurerm_client_config.current.tenant_id
}

output "azure_region" {
  value = azurerm_resource_group.example.location
}

output "test_resource_group" {
  value = azurerm_resource_group.example.name
}

output "test_gen1_name" {
  value = azurerm_data_lake_store.gen1.name
}

output "test_storage_account_key" {
  value     = azurerm_storage_account.blobaccount.primary_access_key
  sensitive = true
}

output "test_storage_account_name" {
  value = azurerm_storage_account.blobaccount.name
}

output "databricks_azure_workspace_resource_id" {
  // The ID of the Databricks Workspace in the Azure management plane.
  value = azurerm_databricks_workspace.example.id
}

output "workspace_url" {
  // The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
  // this is not named as DATABRICKS_HOST, because it affect authentication
  value = "https://${azurerm_databricks_workspace.example.workspace_url}/"
}