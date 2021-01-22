terraform {
  required_providers {
    azurerm =  "~> 2.33"
  }
}

provider "azurerm" {
  features {}
}

// get any env var to tf
data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

module "this" {
  source = "../modules/az-common"
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

resource "azurerm_storage_container" "abfss" {
  name                  = "${local.prefix}-abfss"
  storage_account_name  = module.this.test_storage_v2_account
  container_access_type = "private"
}

output "test_storage_v2_abfss" {
  value = azurerm_storage_container.abfss.name
}

resource "azurerm_storage_blob" "example" {
  name                   = "main.tf"
  storage_account_name   = module.this.test_storage_v2_account
  storage_container_name = azurerm_storage_container.abfss.name
  type                   = "Block"
  source                 = "${path.module}/main.tf"
}

output "cloud_env" {
  value = "azure"
}

output "test_node_type" {
  value = "Standard_D3_v2"
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
}

output "test_storage_v2_wasbs_sas" {
  value = module.this.test_storage_v2_wasbs_sas
  sensitive = true
}

output "test_data_lake_store_name" {
  value = module.this.test_data_lake_store_name
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

output "databricks_host" {
  value = module.this.databricks_host
}
