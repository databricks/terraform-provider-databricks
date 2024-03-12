---
page_title: "Provisioning Azure Databricks with Private Link - Standard deployment."
---

# Deploying pre-requisite resources and enabling Private Link connections - Standard deployment

-> **Note**

- Refer to [adb-with-private-link-standard](https://github.com/databricks/terraform-databricks-examples/tree/main/modules/adb-with-private-link-standard), a Terraform module that contains code used to deploy an Azure Databricks workspace with Azure Private Link using the Standard deployment approach.
- Refer to the [Databricks Terraform Registry modules](https://registry.terraform.io/modules/databricks/examples/databricks/latest) for more Terraform modules and examples to deploy Azure Databricks resources.
- This guide assumes that connectivity from the on-premises user environment is already configured using ExpressRoute or a VPN gateway connection.

[Azure Private Link](https://learn.microsoft.com/en-us/azure/private-link/private-link-overview) support enables private connectivity between users and their Databricks workspaces and between clusters on the data plane and core services on the control plane within the Databricks workspace infrastructure.

You can use Terraform to deploy the underlying cloud resources and the private access settings resources automatically using a programmatic approach.

This guide covers a [standard deployment](https://learn.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/private-link-standard) to configure Azure Databricks with Private Link:
- Two separate VNets are used:
  - A transit VNet
  - A customer Data Plane VNet
- A private endpoint is used for back-end connectivity and deployed in the customer Data Plane VNet.
- A private endpoint is used for front-end connectivity and deployed in the transit VNet.
- A private endpoint is used for web authentication and deployed in the transit VNet.
- A dedicated Databricks workspace, called Web Auth workspace, is used for web authentication traffic. This workspace is configured with the sub-resource **browser_authentication** and deployed using subnets in the transit VNet.

-> **Note**  
- A separate Web Auth workspace is not mandatory but recommended.
- DNS mapping for SSO login callbacks to the Azure Databricks web application can be managed by the Web Auth workspace or another workspace associated with the **browser_authentication** private endpoint.

![Azure Databricks with Private Link - Standard deployment](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/azure-private-link-standard.png)

This guide uses the following variables:

- `cidr_transit`: The network range (CIDR) for the Azure transit VNet
- `cidr_dp`: The network range (CIDR) for the Azure Data Plane VNet
- `rg_transit`: The name of the existing resource group that will contain the Azure transit VNet and the private DNS zone for the Frontend and the Web auth private endpoint
- `rg_dp`: The name of the existing resource group that will contain the Azure Data Plane VNet and the private DNS zone for the Backend private endpoint
- `location`: The location for Azure resources

This guide is provided as-is, and you can use it as the basis for your custom Terraform module.

This guide takes you through the following high-level steps to set up a workspace with Azure Private Link:

- Initialize the required providers
- Configure Azure objects:
  - Deploy two Azure VNets with the following subnets:
   	- Public and private subnets for each Azure Databricks workspace in the Data Plane VNet
   	- Private Link subnet in the Data Plane VNet that will contain the Backend private endpoint
  - Private Link subnet in the Transit VNet that will contain the following private endpoints:
    - Frontend private endpoint
    - Web auth private endpoint
  - Configure the private DNS zone to add:
   	- DNS A record to map connection for workspace access
   	- DNS A record(s) for web_auth
- Workspace Creation

## Provider initialization

Initialize provider

```hcl
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">=3.43.0"
    }
  }
}

provider "azurerm" {
  features {}
}

provider "random" {
}
```

Define the required variables

```hcl
variable "cidr_transit" {
  type = string
}

variable "cidr_dp" {
  type = string
}

variable "rg_transit" {
  type = string
}

variable "rg_dp" {
  type = string
}

variable "location" {
  type = string
}

data "azurerm_client_config" "current" {
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix   = "adb-pl"
  dbfsname = join("", ["dbfs", "${random_string.naming.result}"])
  tags = {
    Environment = "Demo"
    Owner       = lookup(data.external.me.result, "name")
  }
}
```

## Summary

- In the Transit resource group:
  1. Create a Transit VNet
  2. Create a private DNS zone
  3. Create Web Auth Databricks workspace with the sub-resource **browser_authentication**
  4. Create a Frontend private endpoint with the sub-resource **databricks_ui_api**
- In the Data Plane resource group:
  1. Create a Data Plane VNet
  2. Create a private DNS zone
  3. Create a new Azure Databricks workspace
  4. Create a Backend private endpoint with the sub-resource **databricks_ui_api**

## Deploy Transit resources

1. Create a Transit VNet

```hcl
resource "azurerm_virtual_network" "transit_vnet" {
  name                = "${local.prefix}-transit-vnet"
  location            = var.location
  resource_group_name = var.rg_transit
  address_space       = [var.cidr_transit]
  tags                = local.tags
}

resource "azurerm_network_security_group" "transit_sg" {
  name                = "${local.prefix}-transit-nsg"
  location            = var.location
  resource_group_name = var.rg_transit
  tags                = local.tags
}

resource "azurerm_network_security_rule" "transit_aad" {
  name                        = "AllowAAD-transit"
  priority                    = 200
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureActiveDirectory"
  resource_group_name         = var.rg_transit
  network_security_group_name = azurerm_network_security_group.transit_sg.name
}

resource "azurerm_network_security_rule" "transit_azfrontdoor" {
  name                        = "AllowAzureFrontDoor-transit"
  priority                    = 201
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureFrontDoor.Frontend"
  resource_group_name         = var.rg_transit
  network_security_group_name = azurerm_network_security_group.transit_sg.name
}
```

2. Create a private DNS zone:

```hcl
resource "azurerm_private_dns_zone" "dns_auth_front" {
  name                = "privatelink.azuredatabricks.net"
  resource_group_name = var.rg_transit
}

resource "azurerm_private_dns_zone_virtual_network_link" "transitdnszonevnetlink" {
  name                  = "dpcpspokevnetconnection"
  resource_group_name   = var.rg_transit
  private_dns_zone_name = azurerm_private_dns_zone.dns_auth_front.name
  virtual_network_id    = azurerm_virtual_network.transit_vnet.id
}
```

3. Create Web Auth Databricks workspace with the sub-resource **browser_authentication**:

```hcl
resource "azurerm_subnet" "transit_public" {
  name                 = "${local.prefix}-transit-public"
  resource_group_name  = var.rg_transit
  virtual_network_name = azurerm_virtual_network.transit_vnet.name
  address_prefixes     = [cidrsubnet(local.cidr_transit, 6, 0)]

  delegation {
    name = "databricks"
    service_delegation {
      name = "Microsoft.Databricks/workspaces"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
        "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
      "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"]
    }
  }
}

resource "azurerm_subnet_network_security_group_association" "transit_public" {
  subnet_id                 = azurerm_subnet.transit_public.id
  network_security_group_id = azurerm_network_security_group.transit_sg.id
}

variable "transit_private_subnet_endpoints" {
  default = []
}

resource "azurerm_subnet" "transit_private" {
  name                 = "${local.prefix}-transit-private"
  resource_group_name  = var.rg_transit
  virtual_network_name = azurerm_virtual_network.transit_vnet.name
  address_prefixes     = [cidrsubnet(local.cidr_transit, 6, 1)]

  enforce_private_link_endpoint_network_policies = true
  enforce_private_link_service_network_policies  = true

  delegation {
    name = "databricks"
    service_delegation {
      name = "Microsoft.Databricks/workspaces"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
        "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
      "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"]
    }
  }

  service_endpoints = var.transit_private_subnet_endpoints
}

resource "azurerm_subnet_network_security_group_association" "transit_private" {
  subnet_id                 = azurerm_subnet.transit_private.id
  network_security_group_id = azurerm_network_security_group.transit_sg.id
}


resource "azurerm_subnet" "transit_plsubnet" {
  name                                           = "${local.prefix}-transit-privatelink"
  resource_group_name                            = var.rg_transit
  virtual_network_name                           = azurerm_virtual_network.transit_vnet.name
  address_prefixes                               = [cidrsubnet(local.cidr_transit, 6, 2)]
  enforce_private_link_endpoint_network_policies = true
}

resource "azurerm_private_endpoint" "transit_auth" {
  name                = "aadauthpvtendpoint-transit"
  location            = var.location
  resource_group_name = var.rg_transit
  subnet_id           = azurerm_subnet.transit_plsubnet.id

  private_service_connection {
    name                           = "ple-${var.prefix}-auth"
    private_connection_resource_id = azurerm_databricks_workspace.web_auth_workspace.id
    is_manual_connection           = false
    subresource_names              = ["browser_authentication"]
  }

  private_dns_zone_group {
    name                 = "private-dns-zone-auth"
    private_dns_zone_ids = [azurerm_private_dns_zone.dns_auth_front.id]
  }
}

resource "azurerm_databricks_workspace" "web_auth_workspace" {
  name                                  = "${local.prefix}-transit-workspace"
  resource_group_name                   = var.rg_transit
  location                              = var.location
  sku                                   = "premium"
  tags                                  = local.tags
  public_network_access_enabled         = false                    //use private endpoint
  network_security_group_rules_required = "NoAzureDatabricksRules" //use private endpoint
  customer_managed_key_enabled          = true
  custom_parameters {
    no_public_ip                                         = true
    virtual_network_id                                   = azurerm_virtual_network.transit_vnet.id
    private_subnet_name                                  = azurerm_subnet.transit_private.name
    public_subnet_name                                   = azurerm_subnet.transit_public.name
    public_subnet_network_security_group_association_id  = azurerm_subnet_network_security_group_association.transit_public.id
    private_subnet_network_security_group_association_id = azurerm_subnet_network_security_group_association.transit_private.id
    storage_account_name                                 = local.dbfsname
  }
  depends_on = [
    azurerm_subnet_network_security_group_association.transit_public,
    azurerm_subnet_network_security_group_association.transit_private
  ]
}
```

4. Create a Frontend private endpoint with the sub-resource **databricks_ui_api**:

```hcl
resource "azurerm_private_endpoint" "front_pe" {
  name                = "frontprivatendpoint"
  location            = var.location
  resource_group_name = var.rg_transit
  subnet_id           = azurerm_subnet.transit_plsubnet.id

  private_service_connection {
    name                           = "ple-${var.prefix}-uiapi"
    private_connection_resource_id = azurerm_databricks_workspace.app_workspace.id
    is_manual_connection           = false
    subresource_names              = ["databricks_ui_api"]
  }

  private_dns_zone_group {
    name                 = "private-dns-zone-uiapi"
    private_dns_zone_ids = [azurerm_private_dns_zone.dns_auth_front.id]
  }
}
```

## Deploy Data Plane resources

1. Create a Data Plane VNet

```hcl
resource "azurerm_virtual_network" "app_vnet" {
  name                = "${local.prefix}-app-vnet"
  location            = var.location
  resource_group_name = var.rg_dp
  address_space       = [local.cidr_dp]
  tags                = local.tags
}

resource "azurerm_network_security_group" "app_sg" {
  name                = "${local.prefix}-app-nsg"
  location            = var.location
  resource_group_name = var.rg_dp
  tags                = local.tags
}

resource "azurerm_network_security_rule" "app_aad" {
  name                        = "AllowAAD-app"
  priority                    = 200
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureActiveDirectory"
  resource_group_name         = var.rg_dp
  network_security_group_name = azurerm_network_security_group.app_sg.name
}

resource "azurerm_network_security_rule" "app_azfrontdoor" {
  name                        = "AllowAzureFrontDoor-app"
  priority                    = 201
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureFrontDoor.Frontend"
  resource_group_name         = var.rg_dp
  network_security_group_name = azurerm_network_security_group.app_sg.name
}
```

2. Create a private DNS zone:

```hcl
resource "azurerm_private_dns_zone" "dnsdpcp" {
  name                = "privatelink.azuredatabricks.net"
  resource_group_name = var.rg_dp
}

resource "azurerm_private_dns_zone_virtual_network_link" "uiapidnszonevnetlink" {
  name                  = "dpcpvnetconnection"
  resource_group_name   = var.rg_dp
  private_dns_zone_name = azurerm_private_dns_zone.dnsdpcp.name
  virtual_network_id    = azurerm_virtual_network.app_vnet.id
}
```

3. Create a new Azure Databricks workspace

```hcl
resource "azurerm_subnet" "app_public" {
  name                 = "${local.prefix}-app-public"
  resource_group_name  = var.rg_dp
  virtual_network_name = azurerm_virtual_network.app_vnet.name
  address_prefixes     = [cidrsubnet(local.cidr_dp, 6, 0)]

  delegation {
    name = "databricks"
    service_delegation {
      name = "Microsoft.Databricks/workspaces"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
        "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
      "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"]
    }
  }
}

resource "azurerm_subnet_network_security_group_association" "app_public" {
  subnet_id                 = azurerm_subnet.app_public.id
  network_security_group_id = azurerm_network_security_group.app_sg.id
}

variable "private_subnet_endpoints" {
  default = []
}

resource "azurerm_subnet" "app_private" {
  name                 = "${local.prefix}-app-private"
  resource_group_name  = var.rg_dp
  virtual_network_name = azurerm_virtual_network.app_vnet.name
  address_prefixes     = [cidrsubnet(local.cidr_dp, 6, 1)]

  enforce_private_link_endpoint_network_policies = true
  enforce_private_link_service_network_policies  = true

  delegation {
    name = "databricks"
    service_delegation {
      name = "Microsoft.Databricks/workspaces"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
        "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
      "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"]
    }
  }

  service_endpoints = var.private_subnet_endpoints
}

resource "azurerm_subnet_network_security_group_association" "app_private" {
  provider                  = azurerm.app
  subnet_id                 = azurerm_subnet.app_private.id
  network_security_group_id = azurerm_network_security_group.app_sg.id
}


resource "azurerm_subnet" "app_plsubnet" {
  provider                                       = azurerm.app
  name                                           = "${local.prefix}-app-privatelink"
  resource_group_name                            = var.rg_dp
  virtual_network_name                           = azurerm_virtual_network.app_vnet.name
  address_prefixes                               = [cidrsubnet(local.cidr_dp, 6, 2)]
  enforce_private_link_endpoint_network_policies = true
}

resource "azurerm_databricks_workspace" "app_workspace" {
  name                                  = "${local.prefix}-app-workspace"
  resource_group_name                   = var.rg_dp
  location                              = var.location
  sku                                   = "premium"
  tags                                  = local.tags
  public_network_access_enabled         = false                    //use private endpoint
  network_security_group_rules_required = "NoAzureDatabricksRules" //use private endpoint
  customer_managed_key_enabled          = true
  custom_parameters {
    no_public_ip                                         = true
    virtual_network_id                                   = azurerm_virtual_network.app_vnet.id
    private_subnet_name                                  = azurerm_subnet.app_private.name
    public_subnet_name                                   = azurerm_subnet.app_public.name
    public_subnet_network_security_group_association_id  = azurerm_subnet_network_security_group_association.app_public.id
    private_subnet_network_security_group_association_id = azurerm_subnet_network_security_group_association.app_private.id
    storage_account_name                                 = "dbfsapp723k4b3"
  }

  depends_on = [
    azurerm_subnet_network_security_group_association.app_public,
    azurerm_subnet_network_security_group_association.app_private
  ]
}
```

4. Create a Backend private endpoint with the sub-resource **databricks_ui_api**:

```hcl
resource "azurerm_private_endpoint" "app_dpcp" {
  name                = "dpcppvtendpoint"
  resource_group_name = var.rg_dp
  location            = var.location
  subnet_id           = azurerm_subnet.app_plsubnet.id

  private_service_connection {
    name                           = "ple-${var.prefix}-dpcp"
    private_connection_resource_id = azurerm_databricks_workspace.app_workspace.id
    is_manual_connection           = false
    subresource_names              = ["databricks_ui_api"]
  }

  private_dns_zone_group {
    name                 = "app-private-dns-zone-dpcp"
    private_dns_zone_ids = [azurerm_private_dns_zone.dnsdpcp.id]
  }
}
```

-> **Note**

- The public network access to the workspace is disabled. You can access the workspace only through private connectivity to the on-premises user environment. For testing purposes, you can deploy an Azure VM in the Transit VNet to test the frontend connectivity.
- If you wish to deploy a test VM in the Data Plane VNet, you should configure a peering connection between the two VNets
