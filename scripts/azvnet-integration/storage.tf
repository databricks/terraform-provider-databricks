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

resource "azurerm_storage_container" "sample" {
  name                  = "sample"
  storage_account_name  = azurerm_storage_account.blobaccount.name
  container_access_type = "private"
}

output "test_storage_v2_account_name" {
  value = azurerm_storage_account.blobaccount.name
}

output "test_storage_v2_account_key" {
  value     = azurerm_storage_account.blobaccount.primary_access_key
  sensitive = true
}
