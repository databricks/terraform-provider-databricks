provider "azurerm" {
  client_id         = var.client_id
  client_secret     = var.client_secret
  tenant_id         = var.tenant_id
  subscription_id   = var.subscription_id
  features {}
}

resource "azurerm_databricks_workspace" "enterprise-demo-1" {
  location                      = var.region
  name                          = var.dbws_name
  resource_group_name           = var.dbws_rg_name
  managed_resource_group_name   = var.managed_rg_name
  sku                           = "premium"
  custom_parameters {
    no_public_ip = "true"
    public_subnet_name = var.subnet_public
    private_subnet_name = var.subnet_private
    virtual_network_id = var.vnet_id
  }
  tags = {
    Owner = "thestuff@nonexistent.net"
  }
}

provider "databricks" {
  azure_auth = {
    managed_resource_group  = azurerm_databricks_workspace.enterprise-demo-1.managed_resource_group_name
    azure_region            = azurerm_databricks_workspace.enterprise-demo-1.location
    workspace_name          = azurerm_databricks_workspace.enterprise-demo-1.name
    resource_group          = azurerm_databricks_workspace.enterprise-demo-1.resource_group_name
    client_id               = var.client_id
    client_secret           = var.client_secret
    tenant_id               = var.tenant_id
    subscription_id         = var.subscription_id
  }
}

resource "databricks_scim_user" "break-glass-user" {
  user_name     = var.breakglass_user
  display_name  = "User"
  default_roles = []
  roles = []
  set_admin = "true"
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}

resource "databricks_scim_user" "user_jason" {
  user_name     = "jason@robeydespain.net"
  display_name  = "Robey"
  default_roles = []
  roles = []
  set_admin = "false"
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}

resource "databricks_cluster_policy" "billing-tags" {
  name = "RequiredBillingTags"
  definition = "{\"custom_tags.my_dept\": {\"type\": \"fixed\", \"value\": \"RD Enterprises\"}, \"dbus_per_hour\": {\"type\": \"range\",\"maxValue\": 35}, \"autotermination_minutes\": {\"type\": \"range\",\"minValue\": 15,\"maxValue\": 180,\"defaultValue\": 60}}"
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}

resource "databricks_scim_group" "ws_users_2" {
  display_name = "Workspace Users 2"
  members = [databricks_scim_user.user_jason.id]
  entitlements = ["allow-cluster-create","allow-instance-pool-create"]
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}

resource "databricks_workspace_conf" "features" {
  enable_ip_access_lists = "true"
}

resource "databricks_ip_access_list" "naughty_list" {
  label = "lumps_of_coal"
  list_type = "BLACKLIST"
  ip_addresses = [
    "10.0.10.25","10.0.10.0/24"
  ]
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}

resource "databricks_ip_access_list" "nice-list" {
  label = "gifts"
  list_type = "WHITELIST"
  ip_addresses = [
    "10.0.100.0/24",
    "74.36.5.42"
  ]
  depends_on = [azurerm_databricks_workspace.enterprise-demo-1]
}