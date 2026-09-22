---
subcategory: "Postgres"
---
# databricks_postgres_cdf_config Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source retrieves a single Postgres CDF (Change Data Feed) configuration.


## Example Usage
### Retrieve CDF Configuration by Name

```hcl
data "databricks_postgres_cdf_config" "this" {
  name = "projects/my-project/branches/main/databases/app_db/cdf-configs/public"
}

output "cdf_target_catalog" {
  value = data.databricks_postgres_cdf_config.this.catalog
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The full resource name of the CdfConfig.
  Format: projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `catalog` (string) - The Unity Catalog catalog that replicated tables are written into.
  Set at creation; the CdfConfig is immutable
* `cdf_config_id` (string) - The user-specified id; equals the final segment of `name`. Defaults to the
  Postgres schema name for configs without an explicit id
* `create_time` (string) - When the CdfConfig was created
* `name` (string) - Output only. The full resource name of the CdfConfig.
  Format: projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
* `postgres_schema` (string) - The Postgres schema this CdfConfig replicates from. Unique within the
  parent database. Set at creation; the CdfConfig is immutable
* `schema` (string) - The Unity Catalog schema that replicated tables are written into.
  Set at creation; the CdfConfig is immutable