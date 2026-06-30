---
subcategory: "Postgres"
---
# databricks_postgres_data_api Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

The Data API is a [PostgREST](https://docs.postgrest.org/)-compatible REST interface for a Lakebase Postgres database, served by a Databricks-managed implementation. See the [Lakebase Data API docs](https://docs.databricks.com/aws/en/oltp/projects/data-api) for an overview. Creating this resource enables the Data API on the parent database; deleting it disables it.

There is at most one Data API per database. The Data API is identified by the parent database — you do not provide a separate ID.

### Hierarchy Context

Data APIs exist within the Lakebase Autoscaling resource hierarchy:
- A **Data API** belongs to a **database** within a **branch** within a **project**
- A database may have at most one Data API enabled
- Removing the `databricks_postgres_data_api` resource (e.g. via `terraform destroy`) only disables the Data API; the underlying database is left intact

### Configuration

`spec` holds the desired Data API configuration. All fields are optional and have reasonable defaults.

The Data API is wire-compatible with PostgREST and most spec fields are named after the corresponding PostgREST configuration parameter. Where a setting maps to PostgREST, the upstream documentation link is provided and semantics match it unless explicitly noted otherwise. Settings marked **Currently not honored** are accepted, validated, and round-tripped through the API but are ignored by the runtime — they are reserved so configuration set today does not need to be migrated when the behavior is implemented.

#### Database settings

- `db_aggregates_enabled` (`bool`, default `true`) — enables aggregate functions (`count`, `sum`, `avg`, …) in Data API responses. Maps to PostgREST [`db-aggregates-enabled`](https://docs.postgrest.org/en/latest/references/configuration.html#db-aggregates-enabled).
- `db_schemas` (`list<string>`, default `["public"]`) — Postgres schemas exposed through the Data API. Maps to PostgREST [`db-schemas`](https://docs.postgrest.org/en/latest/references/configuration.html#db-schemas). Each schema must already exist in the database (application-managed).
- `db_extra_search_path` (`list<string>`) — additional schemas added to the Postgres `search_path` for resolving unqualified identifiers. Maps to PostgREST [`db-extra-search-path`](https://docs.postgrest.org/en/latest/references/configuration.html#db-extra-search-path).
- `db_max_rows` (`int32`, must be positive) — caps the rows returned by a single query. Maps to PostgREST [`db-max-rows`](https://docs.postgrest.org/en/latest/references/configuration.html#db-max-rows). When unset, no limit is applied (PostgREST default).

#### JWT / authentication settings

- `jwt_role_claim_key` (`string`, default `".sub"`) — JSON path to the JWT claim that selects the Postgres role for the request. Maps to PostgREST [`jwt-role-claim-key`](https://docs.postgrest.org/en/latest/references/configuration.html#jwt-role-claim-key). **Currently not honored** — the runtime hardcodes the role-claim path to `.sub` (the Databricks user email).
- `jwt_cache_max_lifetime` (`Duration`) — maximum lifetime for cached validated JWTs; `0s` disables caching, larger values amortize signature verification across requests. Maps to PostgREST [`jwt-cache-max-lifetime`](https://docs.postgrest.org/en/latest/references/configuration.html#jwt-cache-max-lifetime). **Currently not honored** — the runtime uses fixed caching behavior.

#### Server settings

- `server_cors_allowed_origins` (`list<string>`) — allowed CORS origins, or `["*"]` to allow any origin. Maps to PostgREST [`server-cors-allowed-origins`](https://docs.postgrest.org/en/latest/references/configuration.html#server-cors-allowed-origins).
- `server_timing_enabled` (`bool`) — emits the `Server-Timing` HTTP header on responses. Maps to PostgREST [`server-timing-enabled`](https://docs.postgrest.org/en/latest/references/configuration.html#server-timing-enabled).

#### OpenAPI settings

- `openapi_mode` (enum) — controls how the Data API exposes the auto-generated OpenAPI document. Maps to PostgREST [`openapi-mode`](https://docs.postgrest.org/en/latest/references/configuration.html#openapi-mode). **Only `OPEN_API_MODE_IGNORE_PRIVILEGES` and `OPEN_API_MODE_DISABLED` are supported** today; PostgREST's `follow-privileges` mode is not implemented yet.

`status` is read-only and reflects the actual Data API configuration applied to the database, plus:
- `url` — the HTTPS endpoint at which the Data API is served.
- `available_schemas` — schemas in the database, useful when configuring `db_schemas`.

### Use Cases

- **Public read API**: expose a `public` schema to anonymous clients with `db_max_rows` capped to a safe value.
- **OpenAPI catalog**: set `openapi_mode` so consumers can discover endpoints automatically.


## Example Usage
### Enable the Data API on a Database

Enable the Data API on a Lakebase database with default settings. The `parent` is the database resource name; the resulting Data API URL is exposed via `status.url`.

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "main" {
  branch_id = "main"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_role" "app_owner" {
  role_id = "app-owner"
  parent  = databricks_postgres_branch.main.name
  spec = {
    postgres_role = "app_owner"
  }
}

resource "databricks_postgres_database" "app" {
  database_id = "app"
  parent      = databricks_postgres_branch.main.name
  spec = {
    postgres_database = "app"
    role              = databricks_postgres_role.app_owner.name
  }
}

resource "databricks_postgres_data_api" "app" {
  parent = databricks_postgres_database.app.name
}

output "data_api_url" {
  value = databricks_postgres_data_api.app.status.url
}
```

> [!NOTE]
> **`db_schemas` is application-managed, not infrastructure.**
>
> Enabling the Data API only stands up the PostgREST machinery (`pgrst` schema, `authenticator` role, `pre_config` function); the schemas exposed via `db_schemas` must already exist in the database. They are application data, expected to be created by the application managing the database (SQL, sqitch, Flyway, etc.). Terraform should not own them.

### Custom PostgREST Settings

Tighten response size, restrict CORS, and configure the OpenAPI mode.

```hcl
resource "databricks_postgres_data_api" "app" {
  parent = databricks_postgres_database.app.name
  spec = {
    db_schemas  = ["public"]
    db_max_rows = 1000

    server_cors_allowed_origins = ["https://app.example.com"]
    server_timing_enabled       = true
    openapi_mode                = "OPEN_API_MODE_IGNORE_PRIVILEGES"
  }
}
```

### Caveat: Clearing a Previously-Set Field

Removing a field from `spec` (or setting it to `null`) does **not** clear the value on the server with the typed schema today; the previous value remains in effect. To clear a field, taint the resource and re-apply, or use the API directly.

### Refreshing the PostgREST schema cache

PostgREST caches the database's schemas, roles, and functions in memory and only re-reads them on a refresh. Create/Update normally trigger this refresh automatically, but if it fails (for example a transient network blip), the database is configured but PostgREST can't see the new schemas.

To re-trigger the refresh without changing any settings, call Update with no fields set in `spec` — the server treats this as a refresh-only request, re-reads the cache, and returns the unchanged state. Via the API directly:

```sh
curl -X PATCH "$WORKSPACE/api/2.0/postgres/projects/{p}/branches/{b}/databases/{d}/data-api?update_mask=spec" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"spec": {}}'
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The database containing this Data API configuration.
  Format: projects/{project_id}/branches/{branch_id}/databases/{database_id}
* `spec` (DataApiDataApiSpec, optional) - The desired Data API configuration
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DataApiDataApiSpec
* `db_aggregates_enabled` (boolean, optional) - Enable aggregate functions (count, sum, avg, etc.) in Data API responses.
  Default: true
* `db_extra_search_path` (list of string, optional) - Additional schemas to include in the PostgreSQL search path.
  Each entry must be a valid PostgreSQL schema name
* `db_max_rows` (integer, optional) - Maximum number of rows returned in a single Data API response.
  Must be a positive integer
* `db_schemas` (list of string, optional) - Database schemas exposed through the Data API.
  Each entry must be a valid PostgreSQL schema name (1-63 chars, [a-zA-Z_][a-zA-Z0-9_$]*).
  Maximum 100 entries. Default: ["public"]
* `jwt_cache_max_lifetime` (string, optional) - Maximum lifetime for cached JWT tokens. Zero duration disables caching
* `jwt_role_claim_key` (string, optional) - JSON path to the role claim in JWT tokens (e.g., ".sub").
  Default: ".sub"
* `openapi_mode` (string, optional) - OpenAPI documentation mode for the Data API endpoint. Possible values are: `OPEN_API_MODE_DISABLED`, `OPEN_API_MODE_IGNORE_PRIVILEGES`
* `server_cors_allowed_origins` (list of string, optional) - Allowed origins for CORS requests.
  Each entry should be a valid origin URL, or use "*" to allow all origins
* `server_timing_enabled` (boolean, optional) - Enable the Server-Timing header in Data API responses

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the Data API was first enabled
* `name` (string) - Resource name: projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
* `status` (DataApiDataApiStatus) - The observed Data API state (read-only)
* `update_time` (string) - A timestamp indicating when the Data API configuration was last updated

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

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_data_api.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_data_api.this "name"
```