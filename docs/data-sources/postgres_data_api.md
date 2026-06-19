---
subcategory: "Postgres"
---
# databricks_postgres_data_api Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

This data source retrieves the Data API configuration for a single Lakebase database, including the public Data API URL.


## Example Usage
### Retrieve the Data API for a Database

```hcl
data "databricks_postgres_data_api" "app" {
  name = "projects/my-project/branches/main/databases/app/data-api"
}

output "data_api_url" {
  value = data.databricks_postgres_data_api.app.status.url
}

output "available_schemas" {
  value = data.databricks_postgres_data_api.app.status.available_schemas
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name: projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - A timestamp indicating when the Data API was first enabled
* `name` (string) - Resource name: projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
* `parent` (string) - The database containing this Data API configuration.
  Format: projects/{project_id}/branches/{branch_id}/databases/{database_id}
* `spec` (DataApiDataApiSpec) - The desired Data API configuration
* `status` (DataApiDataApiStatus) - The observed Data API state (read-only)
* `update_time` (string) - A timestamp indicating when the Data API configuration was last updated

### DataApiDataApiSpec
* `db_aggregates_enabled` (boolean) - Enable aggregate functions (count, sum, avg, etc.) in Data API responses.
  Default: true
* `db_extra_search_path` (list of string) - Additional schemas to include in the PostgreSQL search path.
  Each entry must be a valid PostgreSQL schema name
* `db_max_rows` (integer) - Maximum number of rows returned in a single Data API response.
  Must be a positive integer
* `db_schemas` (list of string) - Database schemas exposed through the Data API.
  Each entry must be a valid PostgreSQL schema name (1-63 chars, [a-zA-Z_][a-zA-Z0-9_$]*).
  Maximum 100 entries. Default: ["public"]
* `jwt_cache_max_lifetime` (string) - Maximum lifetime for cached JWT tokens. Zero duration disables caching
* `jwt_role_claim_key` (string) - JSON path to the role claim in JWT tokens (e.g., ".sub").
  Default: ".sub"
* `openapi_mode` (string) - OpenAPI documentation mode for the Data API endpoint. Possible values are: `OPEN_API_MODE_DISABLED`, `OPEN_API_MODE_IGNORE_PRIVILEGES`
* `server_cors_allowed_origins` (list of string) - Allowed origins for CORS requests.
  Each entry should be a valid origin URL, or use "*" to allow all origins
* `server_timing_enabled` (boolean) - Enable the Server-Timing header in Data API responses

### DataApiDataApiStatus
* `available_schemas` (list of string) - Schemas available in the database (for reference when configuring db_schemas)
* `db_aggregates_enabled` (boolean) - Actual aggregate function setting read from the database
* `db_extra_search_path` (list of string) - Actual extra search path schemas read from the database
* `db_max_rows` (integer) - Actual max rows setting read from the database
* `db_schemas` (list of string) - Actual exposed schemas read from the database
* `jwt_cache_max_lifetime` (string) - Actual JWT cache max lifetime read from the database
* `jwt_role_claim_key` (string) - Actual JWT role claim key read from the database
* `openapi_mode` (string) - Actual OpenAPI mode read from the database. Possible values are: `OPEN_API_MODE_DISABLED`, `OPEN_API_MODE_IGNORE_PRIVILEGES`
* `server_cors_allowed_origins` (list of string) - Actual CORS allowed origins read from the database
* `server_timing_enabled` (boolean) - Actual Server-Timing header setting read from the database
* `url` (string) - Data API endpoint URL