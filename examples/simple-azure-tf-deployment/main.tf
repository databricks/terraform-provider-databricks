variable "client_id" {
  type = string
}
variable "client_secret" {
  type = string
}
variable "tenant_id" {
  type = string
}
variable "subscription_id" {
  type = string
}

provider "azurerm" {
  client_id = var.client_id
  client_secret = var.client_secret
  tenant_id = var.tenant_id
  subscription_id = var.subscription_id
  features {}
}

resource "azurerm_resource_group" "db" {
  name     = "db-labs-resources"
  location = "West Europe"
}

resource "azurerm_databricks_workspace" "demo_test_workspace" {
  location = "centralus"
  name = "terraform-test-ws-6"
  resource_group_name = azurerm_resource_group.db.name
  sku = "premium"
}



provider "databricks" {
  azure_auth = {
    managed_resource_group = azurerm_databricks_workspace.demo_test_workspace.managed_resource_group_name
    azure_region = azurerm_databricks_workspace.demo_test_workspace.location
    workspace_name = azurerm_databricks_workspace.demo_test_workspace.name
    resource_group = azurerm_databricks_workspace.demo_test_workspace.resource_group_name
    client_id = var.client_id
    client_secret = var.client_secret
    tenant_id = var.tenant_id
    subscription_id = var.subscription_id
  }
}

resource "databricks_secret_scope" "my-scope" {
  name = "terraform-demo-scope2"
}

resource "databricks_secret" "test_secret" {
  key = "demo-test-secret-1"
  string_value = "hello world 123"
  scope = databricks_secret_scope.my-scope.name
}