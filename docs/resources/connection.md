---
subcategory: "Unity Catalog"
---
# databricks_connection (Resource)

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

- `name` - Name of the Connection.
- `connection_type` - Connection type. `MYSQL`, `POSTGRESQL`, `SNOWFLAKE`, `REDSHIFT` `SQLDW`, `SQLSERVER`, `DATABRICKS`, `SALESFORCE`, `BIGQUERY`, `WORKDAY_RAAS`, `HIVE_METASTORE`, `GA4_RAW_DATA`, `SERVICENOW`, `SALESFORCE_DATA_CLOUD`, `GLUE`, `ORACLE`, `TERADATA`, `HTTP` or `POWER_BI` are supported. Up-to-date list of connection type supported is in the [documentation](https://docs.databricks.com/query-federation/index.html#supported-data-sources). Change forces creation of a new resource.
- `options` - The key value of options required by the connection, e.g. `host`, `port`, `user`, `password`, `authorization_endpoint`, `client_id`, `client_secret` or `GoogleServiceAccountKeyJson`. Please consult the [documentation](https://docs.databricks.com/query-federation/index.html#supported-data-sources) for the required option.
- `owner` - (Optional) Name of the connection owner.
- `properties` -  (Optional) Free-form connection properties. Change forces creation of a new resource.
- `comment` - (Optional) Free-form text. Change forces creation of a new resource.
- `read_only` - (Optional) Indicates whether the connection is read-only. Change forces creation of a new resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this connection in form of `<metastore_id>|<name>`.
- `connection_id` - Unique ID of the connection.
- `created_at` - Time at which this connection was created, in epoch milliseconds.
- `created_by` -  Username of connection creator.
- `credential_type` - The type of credential for this connection.
- `full_name` - Full name of connection.
- `metastore_id` - Unique ID of the UC metastore for this connection.
- `provisioning_info` - Object with the status of an asynchronously provisioned resource.
- `updated_at` - Time at which connection this was last modified, in epoch milliseconds.
- `updated_by` - Username of user who last modified the connection.
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
