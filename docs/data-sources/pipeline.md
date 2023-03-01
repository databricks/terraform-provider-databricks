---
subcategory: "Other"
---
# databricks_pipeline Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

TODO: write me

## Example Usage

Doing X:

```hcl
data "databricks_pipeline" "all" {}

output "all_pipeline" {
  value = data.databricks_pipeline.all
}
```

## Attribute Reference

This data source exports the following attributes:

* `add_field_name` - write doc

## Related Resources

The following resources are used in the same context:

* TODO: write me