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
variable "token" {
  type = string
}
variable "host" {
  type = string
}
variable "resource_group" {
  type = string
}
variable "managed_resource_group_name" {
  type = string
}

provider "azurerm" {
  client_id = var.client_id
  client_secret = var.client_secret
  tenant_id = var.tenant_id
  subscription_id = var.subscription_id
}

resource "azurerm_databricks_workspace" "demo_test_workspace" {
  location = "centralus"
  name = "terraform-test-ws-6"
  resource_group_name = var.resource_group
  managed_resource_group_name = var.managed_resource_group_name
  sku = "premium"
}



provider "databricks" {
  azure_auth = {
    managed_resource_group = azurerm_databricks_workspace.demo_test_workspace.managed_resource_group_name
    azure_region = azurerm_databricks_workspace.demo_test_workspace.location
    workspace_name = azurerm_databricks_workspace.demo_test_workspace.name
    resource_group = azurerm_databricks_workspace.demo_test_workspace.resource_group_name
    workspace_url = azurerm_databricks_workspace.demo_test_workspace.workspace_url
    client_id = var.client_id
    client_secret = var.client_secret
    tenant_id = var.tenant_id
    subscription_id = var.subscription_id
  }
}

//

resource "databricks_scim_user" "my-user" {
  count = 2
  user_name = join("", ["demo-test-user", "+",count.index,"@databricks.com"])
  display_name = "demo Test User"
  entitlements = [
    "allow-cluster-create",
  ]
}

resource "databricks_scim_group" "my-group" {
  display_name = "demo Test Group"
  members = [databricks_scim_user.my-user[0].id]
}
resource "databricks_secret_scope" "my-scope" {
  name = "terraform-demo-scope2"
}

resource "databricks_secret" "test_secret" {
  key = "demo-test-secret-1"
  string_value = "hello world 123"
  scope = databricks_secret_scope.my-scope.name
}

resource "databricks_secret_acl" "my-acl" {
  principal = "USERS"
  permission = "READ"
  scope = databricks_secret_scope.my-scope.name
}

resource "databricks_instance_pool" "my-pool" {
  instance_pool_name = "demo-terraform-pool"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "Standard_DS3_v2"
  idle_instance_autotermination_minutes = 10
  disk_spec = {
    azure_disk_volume_type = "PREMIUM_LRS"
    disk_size = 80
    disk_count = 1
  }
  custom_tags = {
    "creator": "demo user"
    "testChange": "demo user"
  }
}

resource "databricks_token" "my-token" {
  lifetime_seconds = 6000
  comment = "Testing terraform v2"
}

resource "databricks_notebook" "my-dbc-base" {
  content = filebase64("${path.module}/demo-terraform.dbc")
  path = "/terraform-test-folder/folder1/folder2/terraformsamplecode"
  overwrite = false
  mkdirs = true
  format = "DBC"
}

data "databricks_notebook" "my-notebook" {
  path = "/terraform-test-folder/folder1/folder2/terraformsamplecode"
  format = "DBC"
}


output "databricks_user_ids" {
  value = databricks_scim_user.my-user[*].id
}

output "databricks_scope" {
  value = databricks_secret_scope.my-scope.name
}

output "notebook-content" {
  value = data.databricks_notebook.my-notebook.content
}