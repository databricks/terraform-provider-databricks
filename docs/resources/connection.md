---
subcategory: "Unity Catalog"
---
# databricks_connection (Resource)

[API Documentation](https://docs.databricks.com/api/workspace/connections)

-> This resource can only be used with a workspace-level provider!

Lakehouse Federation is the query federation platform for Databricks. Databricks uses Unity Catalog to manage query federation. To make a dataset available for read-only querying using Lakehouse Federation, you create the following:

- A connection, a securable object in Unity Catalog that specifies a path and credentials for accessing an external database system.
- A foreign [catalog](catalog.md)

This resource manages connections in Unity Catalog. Please note that OAuth U2M is not supported as it requires user interaction for authentication.

## Example Usage

Create a connection to a MySQL database

```hcl
resource "databricks_connection" "mysql" {
  name            = "mysql_connection"
  connection_type = "MYSQL"
  comment         = "this is a connection to mysql db"
  options = {
    host     = "test.mysql.database.azure.com"
    port     = "3306"
    user     = "user"
    password = "password"
  }
  properties = {
    purpose = "testing"
  }
}
```

Create a connection to a BigQuery database

```hcl
resource "databricks_connection" "bigquery" {
  name            = "bq_connection"
  connection_type = "BIGQUERY"
  comment         = "this is a connection to BQ"
  options = {
    GoogleServiceAccountKeyJson = jsonencode({
      "type" : "service_account",
      "project_id" : "PROJECT_ID",
      "private_key_id" : "KEY_ID",
      "private_key" : "-----BEGIN PRIVATE KEY-----\nPRIVATE_KEY\n-----END PRIVATE KEY-----\n",
      "client_email" : "SERVICE_ACCOUNT_EMAIL",
      "client_id" : "CLIENT_ID",
      "auth_uri" : "https://accounts.google.com/o/oauth2/auth",
      "token_uri" : "https://accounts.google.com/o/oauth2/token",
      "auth_provider_x509_cert_url" : "https://www.googleapis.com/oauth2/v1/certs",
      "client_x509_cert_url" : "https://www.googleapis.com/robot/v1/metadata/x509/SERVICE_ACCOUNT_EMAIL",
      "universe_domain" : "googleapis.com"
    })
  }
  properties = {
    purpose = "testing"
  }
}
```

Create a connection to builtin Hive Metastore

```hcl
resource "databricks_connection" "hms" {
  name            = "hms-builtin"
  connection_type = "HIVE_METASTORE"
  comment         = "This is a connection to builtin HMS"
  options = {
    builtin = "true"
  }
}
```

Create a HTTP connection with bearer token

```hcl
resource "databricks_connection" "http_bearer" {
  name            = "http_bearer"
  connection_type = "HTTP"
  comment         = "This is a connection to a HTTP service"
  options = {
    host         = "https://example.com"
    port         = "8433"
    base_path    = "/api/"
    bearer_token = "bearer_token"
  }
}
```

Create a HTTP connection with OAuth M2M

```hcl
resource "databricks_connection" "http_oauth" {
  name            = "http_oauth"
  connection_type = "HTTP"
  comment         = "This is a connection to a HTTP service"
  options = {
    host           = "https://example.com"
    port           = "8433"
    base_path      = "/api/"
    client_id      = "client_id"
    client_secret  = "client_secret"
    oauth_scope    = "channels:read channels:history chat:write"
    token_endpoint = "https://authorization-server.com/oauth/token"
  }
}
```

Create a PowerBI connection with OAuth M2M

```hcl
resource "databricks_connection" "pbi" {
  name            = "test-pbi"
  connection_type = "POWER_BI"
  options = {
    authorization_endpoint = "https://login.microsoftonline.com/{tenant}/oauth2/v2.0/authorize"
    client_id              = "client_id"
    client_secret          = "client_secret"
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of the connection.
- `connection_type` - (Required) The type of connection. Possible values are: `BIGQUERY`, `CONFLUENCE`, `DATABRICKS`, `GA4_RAW_DATA`, `GITHUB`, `GLUE`, `HIVE_METASTORE`, `HTTP`, `HUBSPOT`, `META_MARKETING`, `MYSQL`, `ORACLE`, `OUTLOOK`, `POSTGRESQL`, `POWER_BI`, `REDSHIFT`, `SALESFORCE`, `SALESFORCE_DATA_CLOUD`, `SERVICENOW`, `SMARTSHEET`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`, `TERADATA`, `WORKDAY_RAAS`, or `ZENDESK`. For an up-to-date list of connection types and required options, see the [documentation](https://docs.databricks.com/query-federation/index.html#supported-data-sources). Change forces creation of a new resource.
- `options` - (Required) A map of key-value properties attached to the securable. The required keys depend on the connection type, e.g. `host`, `port`, `user`, `password`, `authorization_endpoint`, `client_id`, `client_secret`, or `GoogleServiceAccountKeyJson`. Please consult the [documentation](https://docs.databricks.com/query-federation/index.html#supported-data-sources) for the required options. This field is sensitive.
- `comment` - (Optional) User-provided free-form text description. Change forces creation of a new resource.
- `environment_settings` - (Optional) Connection environment settings. This block consists of the following fields:
  - `environment_version` - Environment version.
  - `java_dependencies` - List of Java dependencies.
- `owner` - (Optional) Username of current owner of the connection.
- `properties` - (Optional) A map of key-value properties attached to the securable. Change forces creation of a new resource.
- `read_only` - (Optional) If the connection is read only. Change forces creation of a new resource.
- `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  - `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this connection in form of `<metastore_id>|<name>`.
- `connection_id` - Unique identifier of the Connection.
- `created_at` - Time at which this connection was created, in epoch milliseconds.
- `created_by` - Username of connection creator.
- `credential_type` - The type of credential.
- `full_name` - Full name of connection.
- `metastore_id` - Unique identifier of parent metastore.
- `provisioning_info` - Status of an asynchronously provisioned resource. This block consists of the following fields:
  - `state` - The provisioning state of the resource. Possible values are: `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, or `UPDATING`.
- `securable_type` - Securable type.
- `updated_at` - Time at which this connection was updated, in epoch milliseconds.
- `updated_by` - Username of user who last modified connection.
- `url` - URL of the remote data source, extracted from options.

## Import

This resource can be imported by `id`:

```hcl
import {
  to = databricks_connection.this
  id = "<metastore_id>|<name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_connection.this "<metastore_id>|<name>"
```
