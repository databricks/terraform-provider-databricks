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
  source = "../modules/az-common"
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

output "test_data_lake_store_name" {
  value = module.this.test_data_lake_store_name
}

output "test_storage_v2_account" {
  value = module.this.test_storage_v2_account
}

output "test_storage_v2_key" {
  value = module.this.test_storage_v2_key
  sensitive = true
}

output "test_storage_v2_wasbs" {
  value = module.this.test_storage_v2_wasbs
  sensitive = true
}

output "test_key_vault_name" {
  value = module.this.test_key_vault_name
}

output "test_key_vault_resource_id" {
  value = module.this.test_key_vault_resource_id
}

output "test_key_vault_dns_name" {
  value = module.this.test_key_vault_dns_name
}

output "test_key_vault_secret" {
  value = module.this.test_key_vault_secret
  sensitive = true # :)
}

output "test_key_vault_secret_value" {
  value = module.this.test_key_vault_secret_value
}

output "databricks_azure_workspace_resource_id" {
  value = module.this.databricks_azure_workspace_resource_id
}

output "workspace_url" {
  // explicitly having this as different variable name,
  // so we test ensureWorkspaceInfo
  value = module.this.databricks_host
}