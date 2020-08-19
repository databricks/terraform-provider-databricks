provider "azurerm" {
  version = "~> 2.14"
  features {}
}

// get any env var to tf
data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

module "this" {
  source = "../az-common"
  owner  = data.external.env.result.OWNER
}

data "azurerm_client_config" "current" {
}

locals {
  prefix = module.this.test_prefix
  tags = {
    Environment = "Testing"
    Owner       = data.external.env.result.OWNER
    Epoch       = module.this.test_prefix
  }
}

# Create container in storage acc and container for use by adls gen2 mount tests
resource "azurerm_storage_account" "adlsaccount" {
  name                     = "${local.prefix}datalake"
  resource_group_name      = module.this.test_resource_group
  location                 = module.this.azure_region
  account_tier             = "Standard"
  account_replication_type = "GRS"
  account_kind             = "StorageV2"
  is_hns_enabled           = true
  tags                     = local.tags
}

# SP has to have Microsoft.Authorization/roleAssignments/write
# resource "azurerm_role_assignment" "datalake" {
#   scope = azurerm_storage_account.adlsaccount.id
#   role_definition_name = "Storage Blob Data Contributor"
#   principal_id         = data.azurerm_client_config.current.object_id
# }

resource "azurerm_storage_container" "adlsexample" {
  name                  = "dev"
  storage_account_name  = azurerm_storage_account.adlsaccount.name
  container_access_type = "private"
}

output "cloud_env" {
  value = "azure"
}

output "test_node_type" {
  value = "Standard_D3_v2"
}

output "test_resource_group" {
  value = module.this.test_resource_group
}

output "azure_region" {
  value = module.this.azure_region
}


output "test_gen2_adal_name" {
  value = azurerm_storage_account.adlsaccount.name
}

output "test_gen1_name" {
  value = module.this.test_gen1_name
}

output "test_storage_account_key" {
  value = module.this.test_storage_account_key
  sensitive = true
}

output "test_storage_account_name" {
  value = module.this.test_storage_account_name
}

output "databricks_azure_workspace_resource_id" {
  value = module.this.databricks_azure_workspace_resource_id
}

output "workspace_url" {
  value = module.this.workspace_url
}