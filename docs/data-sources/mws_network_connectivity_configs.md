---
subcategory: "Deployment"
---
# databricks_mws_network_connectivity_configs Data Source

-> **Note** This data source can only be used with an account-level provider!

Lists all [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) in Databricks Account.

## Example Usage

List all network connectivity configurations in Databricks Account

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_network_connectivity_configs" "this" {}

output "all" {
  value = data.databricks_mws_network_connectivity_configs.this
}
```

List network connectivity configurations from a specific region in Databricks Account

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_network_connectivity_configs" "this" {
    region = "us-east-1"
}

output "filtered" {
  value = data.databricks_mws_network_connectivity_configs.this
}
```

## Argument Reference

* `region` - (Optional) Filter network connectivity configurations by region.

## Attribute Reference

This data source exports the following attributes:

* `names` - List of names of [databricks_mws_network_connectivity_config](./databricks_mws_network_connectivity_config.md)

* [databricks_mws_network_connectivity_config](./mws_network_connectivity_config.md) to get information about a single network connectivity configuration.
* [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) to manage network connectivity configuration.
