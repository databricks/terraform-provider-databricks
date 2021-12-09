locals {
  prefix     = "dltp${random_string.naming.result}"
  cidr_block = data.external.env.result.TEST_CIDR
  region     = data.external.env.result.TEST_REGION
  account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
  tags       = azurerm_resource_group.this.tags
}

data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "this" {
  name     = "${local.prefix}-terraform-it"
  location = "West Europe"
  tags = {
    Owner = lookup(data.external.me.result, "name")
  }
}

resource "azurerm_user_assigned_identity" "this" {
  name                = "${local.prefix}-msi"
  resource_group_name = azurerm_resource_group.this.name
  location            = azurerm_resource_group.this.location
  tags                = azurerm_resource_group.this.tags
}

resource "azurerm_role_assignment" "msi_resource_group" {
  scope                = azurerm_resource_group.this.id
  role_definition_name = "Contributor"
  principal_id         = azurerm_user_assigned_identity.this.principal_id
}

resource "azurerm_role_assignment" "msi_blob" {
  scope                = azurerm_resource_group.this.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azurerm_user_assigned_identity.this.principal_id
}

resource "azurerm_databricks_workspace" "this" {
  managed_resource_group_name = "${local.prefix}-workspace-rg"
  name                        = "${local.prefix}-workspace"
  sku                         = "premium"
  resource_group_name         = azurerm_resource_group.this.name
  location                    = azurerm_resource_group.this.location
  tags                        = azurerm_resource_group.this.tags
}

resource "azurerm_storage_account" "this" {
  name                     = "${local.prefix}storage"
  resource_group_name      = azurerm_resource_group.this.name
  location                 = azurerm_resource_group.this.location
  tags                     = azurerm_resource_group.this.tags
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
}
