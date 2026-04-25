---
subcategory: "Data Classification"
---
# databricks_data_classification_catalog_config Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch the Data Classification configuration for a Unity Catalog catalog.

To read the Data Classification configuration, you must have browse permissions on the catalog.

-> **Note** This data source can only be used with a workspace-level provider.


## Example Usage
To get the Data Classification configuration for a catalog:

```hcl
data "databricks_catalog" "prod" {
  name = "prod_catalog"
}

data "databricks_data_classification_catalog_config" "prod_config" {
  name = "catalogs/${data.databricks_catalog.prod.name}/config"
}

# Use the fetched configuration
output "auto_tag_configs" {
  value = data.databricks_data_classification_catalog_config.prod_config.auto_tag_configs
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name in the format: catalogs/{catalog_name}/config
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `auto_tag_configs` (list of AutoTaggingConfig) - List of auto-tagging configurations for this catalog.
  Empty list means no auto-tagging is enabled
* `included_schemas` (CatalogConfigSchemaNames) - Schemas to include in the scan. Empty list is not supported as it results in a no-op
  scan. If `included_schemas` is not set, all schemas are scanned
* `name` (string) - Resource name in the format: catalogs/{catalog_name}/config

### AutoTaggingConfig
* `auto_tagging_mode` (string) - Whether auto-tagging is enabled or disabled for this classification tag. Possible values are: `AUTO_TAGGING_DISABLED`, `AUTO_TAGGING_ENABLED`
* `classification_tag` (string) - The Classification Tag (e.g., "class.name", "class.location")

### CatalogConfigSchemaNames
* `names` (list of string)