---
page_title: "Provisioning Azure Databricks Workspace"
subcategory: "Guides"
---

# Provisioning Azure Databricks

The following sample configuration assumes you have authorized with `az login` on your local machine and have `Contributor` rights to your subscription.

## Simple setup

```hcl
terraform {
  required_providers {
    azurerm =  "~> 2.33"
    random = "~> 2.2"
  }
}

provider "azurerm" {
  features {}
}

variable "region" {
  type = string
  default = "westeurope"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

data "azurerm_client_config" "current" {
}

data "external" "me" {
  program = ["az", "account", "show", "--query", "user"]
}

locals {
  prefix = "databricksdemo${random_string.naming.result}"
  tags = {
    Environment = "Demo"
    Owner       = lookup(data.external.me.result, "name")
  }
}

resource "azurerm_resource_group" "this" {
  name     = "${local.prefix}-rg"
  location = var.region
  tags     = local.tags
}

resource "azurerm_databricks_workspace" "this" {
  name                        = "${local.prefix}-workspace"
  resource_group_name         = azurerm_resource_group.this.name
  location                    = azurerm_resource_group.this.location
  sku                         = "premium"
  managed_resource_group_name = "${local.prefix}-workspace-rg"
  tags                        = local.tags
}

output "databricks_host" {
  value = "https://${azurerm_databricks_workspace.this.workspace_url}/"
}
```

## Provider configuration

In [the next step](workspace-management.md), please use the [special configurations for Azure](../index.md#special-configurations-for-azure):

```hcl
provider "databricks" {
    azure_workspace_resource_id = azurerm_databricks_workspace.this.id
}
```