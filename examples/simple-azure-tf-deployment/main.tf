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



provider "db" {
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

//

resource "db_scim_user" "my-user" {
  count = 2
  user_name = join("", ["demo-test-user", "+",count.index,"@databricks.com"])
  display_name = "demo Test User"
  entitlements = [
    "allow-cluster-create",
  ]
}

resource "db_scim_group" "my-group" {
  display_name = "demo Test Group"
  members = [db_scim_user.my-user[0].id]
}
resource "db_secret_scope" "my-scope" {
  name = "terraform-demo-scope2"
}

resource "db_secret" "test_secret" {
  key = "demo-test-secret-1"
  string_value = "hello world 123"
  scope = db_secret_scope.my-scope.name
}

resource "db_secret_acl" "my-acl" {
  principal = "USERS"
  permission = "READ"
  scope = db_secret_scope.my-scope.name
}

resource "db_instance_pool" "my-pool" {
  instance_pool_name = "demo-terraform-pool"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "Standard_DS3_v2"
  idle_instance_autotermination_minutes = 10
  disk_spec = {
    ebs_volume_type = "GENERAL_PURPOSE_SSD"
    disk_size = 80
    disk_count = 1
  }
  custom_tags = {
    "creator": "demo user"
    "testChange": "demo user"
  }
}

resource "db_token" "my-token" {
  lifetime_seconds = 6000
  comment = "Testing terraform v2"
}

output "db_user_ids" {
  value = db_scim_user.my-user[*].id
}

output "db_scope" {
  value = db_secret_scope.my-scope.name
}
