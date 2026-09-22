---
subcategory: "Postgres"
---
# databricks_postgres_cdf_configs Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source lists all Postgres CDF (Change Data Feed) configurations under a database.


## Example Usage
### List All CDF Configurations in a Database

```hcl
data "databricks_postgres_cdf_configs" "all" {
  parent = "projects/my-project/branches/main/databases/app_db"
}

output "cdf_config_names" {
  value = [for config in data.databricks_postgres_cdf_configs.all.cdf_configs : config.name]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The parent database to list CdfConfigs for.
  Format: projects/{project}/branches/{branch}/databases/{database}
* `page_size` (integer, optional) - Maximum number of CdfConfigs to return
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `cdf_configs`. It is a list of resources, each with the following attributes:
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