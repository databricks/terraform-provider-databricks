# TODO: this file is in progress - azsp will test for azure service principal and azcli will test just for cli (reduced set)
module "this" {
    source = "../azure-integration"
}

output "cloud_env" {
  value = "azure"
}

output "test_node_type" {
  // or i3.xlarge for AWS
  value = "Standard_D3_v2"
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

output "test_resource_group" {
  value = azurerm_resource_group.example.name
}

output "test_gen2_adal_name" {
  value = azurerm_storage_account.adlsaccount.name
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

output "azure_region" {
  value = azurerm_storage_account.adlsaccount.location
}

output "databricks_azure_workspace_resource_id" {
  value = module.this.databricks_azure_workspace_resource_id
}

output "workspace_url" {
  value = module.this.workspace_url
}