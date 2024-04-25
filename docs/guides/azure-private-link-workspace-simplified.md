---
page_title: "Provisioning Azure Databricks with Private Link - Simple deployment."
---

# Deploying pre-requisite resources and enabling Private Link connections - Simple deployment

-> **Note** Refer to the [Databricks Terraform Registry modules](https://registry.terraform.io/modules/databricks/examples/databricks/latest) for Terraform modules and examples to deploy Azure Databricks resources.

-> **Note** This guide assumes that connectivity from the on-premises user environment is already configured using ExpressRoute or a VPN gateway connection.

[Azure Private Link](https://learn.microsoft.com/en-us/azure/private-link/private-link-overview) support enables private connectivity between users and their Databricks workspaces and between clusters on the data plane and core services on the control plane within the Databricks workspace infrastructure.

You can use Terraform to deploy the underlying cloud resources and the private access settings resources automatically using a programmatic approach.

This guide covers a [simple deployment](https://learn.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/private-link-simplified) to configure Azure Databricks with Private Link:

* No separate VNet separates user access from the VNet that you use for your compute resources in the Classic data plane
* A transit subnet in the data plane VNet is used for user access
* Only a single private endpoint is used for both front-end and back-end connectivity.
* A separate private endpoint is used for web authentication
* The same Databricks workspace is used for web authentication traffic. Databricks still strongly recommends creating a separate workspace called a private web auth workspace for each region to host the web auth private network settings.

![Azure Databricks with Private Link - Simple deployment](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/images/azure-private-link-simplified.png)

This guide uses the following variables:

* `cidr`: The CIDR for the Azure Vnet
* `rg_name`: The name of the existing resource group
* `location`: The location for Azure resources

This guide is provided as-is, and you can use it as the basis for your custom Terraform module.

This guide takes you through the following high-level steps to set up a workspace with Azure Private Link:

* Initialize the required providers
* Configure Azure objects:
  * Deploy an Azure Vnet with the following subnets:
   	* Public and private subnets for Azure Databricks workspace
  * Private Link subnet that will contain the following private endpoints:
    * Frontend / Backend private endpoint
    * Web_auth private endpoint
  * Configure the private DNS zone to add:
   	* DNS A record to map connection for workspace access
   	* DNS A record(s) for web_auth
* Workspace Creation

## Provider initialization

Initialize provider

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">=3.43.0"
    }
  }
}

provider "azurerm" {
  features {}
}
```

Define the required variables

```hcl
variable "cidr" {
  type = string
}

variable "rg_name" {
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

locals {
  prefix = "abd-pl"
  tags = {
    Environment = "Demo"
    Owner       = lookup(data.external.me.result, "name")
  }
}
```

## Configure network

### Deploy Azure VNet and subnets

Create a new Azure VNet, the required subnets, and associated security groups:

```hcl
resource "azurerm_virtual_network" "this" {
  name                = "${local.prefix}-vnet"
  location            = var.location
  resource_group_name = var.rg_name
  address_space       = [var.cidr]
  tags                = local.tags
}

resource "azurerm_network_security_group" "this" {
  name                = "${local.prefix}-nsg"
  location            = var.location
  resource_group_name = var.rg_name
  tags                = local.tags
}

resource "azurerm_network_security_rule" "aad" {
  name                        = "AllowAAD"
  priority                    = 200
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureActiveDirectory"
  resource_group_name         = var.rg_name
  network_security_group_name = azurerm_network_security_group.this.name
}

resource "azurerm_network_security_rule" "azfrontdoor" {
  name                        = "AllowAzureFrontDoor"
  priority                    = 201
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "443"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "AzureFrontDoor.Frontend"
  resource_group_name         = var.rg_name
  network_security_group_name = azurerm_network_security_group.this.name
}

resource "azurerm_subnet" "public" {
  name                 = "${local.prefix}-public"
  resource_group_name  = var.rg_name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = [cidrsubnet(var.cidr, 3, 0)]

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

resource "azurerm_subnet_network_security_group_association" "public" {
  subnet_id                 = azurerm_subnet.public.id
  network_security_group_id = azurerm_network_security_group.this.id
}

variable "private_subnet_endpoints" {
  default = []
}

resource "azurerm_subnet" "private" {
  name                 = "${local.prefix}-private"
  resource_group_name  = var.rg_name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = [cidrsubnet(var.cidr, 3, 1)]

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

resource "azurerm_subnet_network_security_group_association" "private" {
  subnet_id                 = azurerm_subnet.private.id
  network_security_group_id = azurerm_network_security_group.this.id
}


resource "azurerm_subnet" "plsubnet" {
  name                                           = "${local.prefix}-privatelink"
  resource_group_name                            = var.rg_name
  virtual_network_name                           = azurerm_virtual_network.this.name
  address_prefixes                               = [cidrsubnet(var.cidr, 3, 2)]
  enforce_private_link_endpoint_network_policies = true
}

```

### Deploy Azure private endpoints

#### Frontend / Backend private endpoint

Create a private endpoint with sub-resource **databricks_ui_api**:

```hcl
resource "azurerm_private_endpoint" "uiapi" {
  name                = "uiapipvtendpoint"
  location            = var.location
  resource_group_name = var.rg_name
  subnet_id           = azurerm_subnet.plsubnet.id

  private_service_connection {
    name                           = "ple-${var.workspace_prefix}-uiapi"
    private_connection_resource_id = azurerm_databricks_workspace.this.id
    is_manual_connection           = false
    subresource_names              = ["databricks_ui_api"]
  }

  private_dns_zone_group {
    name                 = "private-dns-zone-uiapi"
    private_dns_zone_ids = [azurerm_private_dns_zone.dnsuiapi.id]
  }
}

resource "azurerm_private_dns_zone" "dnsuiapi" {
  name                = "privatelink.azuredatabricks.net"
  resource_group_name = var.rg_name
}

resource "azurerm_private_dns_zone_virtual_network_link" "uiapidnszonevnetlink" {
  name                  = "uiapispokevnetconnection"
  resource_group_name   = var.rg_name
  private_dns_zone_name = azurerm_private_dns_zone.dnsuiapi.name
  virtual_network_id    = azurerm_virtual_network.this.id // connect to spoke vnet
}
```

#### Web auth private endpoint

Create a private endpoint with sub-resource **browser_authentication**:

```hcl
resource "azurerm_private_endpoint" "auth" {
  name                = "aadauthpvtendpoint"
  location            = var.location
  resource_group_name = var.rg_name
  subnet_id           = azurerm_subnet.plsubnet.id

  private_service_connection {
    name                           = "ple-${var.workspace_prefix}-auth"
    private_connection_resource_id = azurerm_databricks_workspace.this.id
    is_manual_connection           = false
    subresource_names              = ["browser_authentication"]
  }

  private_dns_zone_group {
    name                 = "private-dns-zone-auth"
    private_dns_zone_ids = [azurerm_private_dns_zone.dnsdpcp.id]
  }
}
```

## Configure workspace

Deploy an Azure Databricks workspace:

```hcl
resource "azurerm_databricks_workspace" "this" {
  name                                  = "${local.prefix}-workspace"
  resource_group_name                   = var.rg_name
  location                              = var.location
  sku                                   = "premium"
  tags                                  = local.tags
  public_network_access_enabled         = false
  network_security_group_rules_required = "NoAzureDatabricksRules"
  customer_managed_key_enabled          = true
  custom_parameters {
    no_public_ip                                         = true
    virtual_network_id                                   = azurerm_virtual_network.this.id
    private_subnet_name                                  = azurerm_subnet.private.name
    public_subnet_name                                   = azurerm_subnet.public.name
    public_subnet_network_security_group_association_id  = azurerm_subnet_network_security_group_association.public.id
    private_subnet_network_security_group_association_id = azurerm_subnet_network_security_group_association.private.id
    storage_account_name                                 = "dbfs"
  }

  depends_on = [
    azurerm_subnet_network_security_group_association.public,
    azurerm_subnet_network_security_group_association.private
  ]
}
```

-> **Note** The public network access to the workspace is disabled. You can access the workspace only through private connectivity to the on-premises user environment. For testing purposes, you can deploy an Azure VM in the same VNet to test the frontend connectivity.
