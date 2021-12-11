---
subcategory: "Compute"
---
# databricks_clusters Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves a list of [databricks_cluster](../resources/cluster.md#cluster_id) ids, that were created by Terraform or manually, with or without [databricks_cluster_policy](../resources/cluster_policy.md).

## Example Usage

Retrieve all clusters on this workspace on AWS or GCP:

```hcl
data "databricks_clusters" "all" {
    depends_on = [databricks_mws_workspaces.this]
}
```

Retrieve all clusters with "Shared" in their cluster name on this Azure Databricks workspace:

```hcl
data "databricks_clusters" "all_shared" {
    depends_on = [azurerm_databricks_workspace.this]
    cluster_name_contains = "shared"
}
```

## Argument Reference

* `cluster_name_contains` - (Optional) Only return [databricks_cluster](../resources/cluster.md#cluster_id) ids that match the given name string.

## Attribute Reference

This data source exports the following attributes:

* `ids` - list of [databricks_cluster](../resources/cluster.md#cluster_id) ids
