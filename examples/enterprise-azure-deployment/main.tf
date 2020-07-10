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
    Owner = "jason.robey@databricks.com"
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
  display_name  = var.breakglass_user
  default_roles = []
  roles = []
  set_admin = "true"
}

resource "databricks_ipaccesslists" "my-home-net"{
  preview_ipacl_enabled = "true"
  ip_acls = [
    {
      label = "my_net"
      type = "WHITELIST"
      ip_addresses = [
        "10.0.10.0/24",
      ]
    }
  ]
}
