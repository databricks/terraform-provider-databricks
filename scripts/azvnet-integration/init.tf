terraform {
  required_providers {
    azurerm =  "~> 2.33"
  }
}

provider "azurerm" {
  features {}
}

provider "random" {
  version = "~> 2.2"
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

variable "cidr" {
  default = "10.4.0.0/16"
}

locals {
  // dltp - databricks labs terraform provider
  prefix   = "dltp${random_string.naming.result}"
  location = "eastus"
  cidr     = var.cidr
  // tags that are propagated down to all resources
  tags = {
    Environment = "Testing"
    Owner       = lookup(data.external.me.result, "name")
    Epoch       = random_string.naming.result
  }
}

resource "azurerm_resource_group" "example" {
  name     = "${local.prefix}-rg"
  location = local.location
  tags     = local.tags
}

output "test_resource_group" {
  value = azurerm_resource_group.example.name
}
