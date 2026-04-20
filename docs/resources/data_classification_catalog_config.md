---
subcategory: "Data Classification"
---
# databricks_data_classification_catalog_config Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This resource allows you to manage the Data Classification configuration for Unity Catalog catalogs.

Data Classification automatically identifies and tags sensitive data (personally identifiable information, or PII) in Unity Catalog tables. Creating this resource enables Data Classification for the specified catalog, while deleting it disables Data Classification.

To manage Data Classification configuration, you must either:
1. be an owner of the catalog, OR
2. have **USE_CATALOG** and **MANAGE** permissions on the catalog

-> **Note** This resource can only be used with a workspace-level provider.


## Example Usage
```hcl
# Enable Data Classification for a set of schemas in a catalog
resource "databricks_data_classification_catalog_config" "example" {
  name = "catalogs/prod_catalog/config"

  included_schemas = {
    names = ["sales", "marketing", "customer_data"]
  }

  auto_tag_configs = [
    {
      classification_tag = "class.credit_card"
      auto_tagging_mode  = "AUTO_TAGGING_ENABLED"
    },
    {
      classification_tag = "class.email"
      auto_tagging_mode  = "AUTO_TAGGING_ENABLED"
    }
  ]
}

# Enable Data Classification for the entire catalog (all current and future schemas)
resource "databricks_data_classification_catalog_config" "all_schemas" {
  name = "catalogs/staging_catalog/config"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - Parent resource in the format: catalogs/{catalog_name}
* `auto_tag_configs` (list of AutoTaggingConfig, optional) - List of auto-tagging configurations for this catalog.
  Empty list means no auto-tagging is enabled
* `included_schemas` (CatalogConfigSchemaNames, optional) - Schemas to include in the scan. Empty list is not supported as it results in a no-op
  scan. If `included_schemas` is not set, all schemas are scanned
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AutoTaggingConfig
* `auto_tagging_mode` (string, required) - Whether auto-tagging is enabled or disabled for this classification tag. Possible values are: `AUTO_TAGGING_DISABLED`, `AUTO_TAGGING_ENABLED`
* `classification_tag` (string, required) - The Classification Tag (e.g., "class.name", "class.location")

### CatalogConfigSchemaNames
* `names` (list of string, required)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `name` (string) - Resource name in the format: catalogs/{catalog_name}/config

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_data_classification_catalog_config.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_data_classification_catalog_config.this "name"
```