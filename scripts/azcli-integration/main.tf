provider "azurerm" {
  version = "~> 2.14"
  features {}
}

data "azurerm_client_config" "current" {
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

module "this" {
  source = "../az-common"
  owner  = lookup(data.external.me.result, "name")
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

output "test_gen1_name" {
  value = module.this.test_gen1_name
}

output "test_storage_account_key" {
  value     = module.this.test_storage_account_key
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