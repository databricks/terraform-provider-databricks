---
subcategory: "Postgres"
---
# databricks_postgres_cdf_statuses Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source lists the replication statuses of all Postgres tables replicated under a CDF (Change Data Feed) configuration.


## Example Usage
### List All Replicated Table Statuses in a CDF Configuration

```hcl
data "databricks_postgres_cdf_statuses" "all" {
  parent = "projects/my-project/branches/main/databases/app_db/cdf-configs/public"
}

output "cdf_table_states" {
  value = {
    for status in data.databricks_postgres_cdf_statuses.all.cdf_statuses :
    status.postgres_table => status.state
  }
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The parent CdfConfig to list CdfStatuses for.
  Format: projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
* `page_size` (integer, optional) - Maximum number of CdfStatuses to return
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `cdf_statuses`. It is a list of resources, each with the following attributes:
* `committed_lsn` (string) - The high-watermark Log Sequence Number (LSN) committed to Delta Lake
* `create_time` (string) - When replication for this table was first established
* `last_sync_time` (string) - The last time changes for this table were written to Delta Lake
* `name` (string) - Output only. The full resource name of the CdfStatus.
  Format: projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}/cdf-statuses/{cdf_status}
  The {cdf_status} segment is the Postgres table name
* `postgres_table` (string) - The Postgres table being replicated
* `state` (string) - The current replication state of this table. Possible values are: `CDF_STATE_SKIPPED`, `CDF_STATE_SNAPSHOTTING`, `CDF_STATE_STREAMING`, `CDF_STATE_TERMINATED`
* `status_detail` (string) - Human-readable detail for the current state (e.g. the skip/error reason).
  Empty for healthy states
* `uc_table` (string) - The Unity Catalog table receiving replicated data