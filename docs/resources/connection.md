---
subcategory: "Unity Catalog"
---
# databricks_connection (Resource)

-> This resource can only be used with a workspace-level provider!

Lakehouse Federation is the query federation platform for Databricks. Databricks uses Unity Catalog to manage query federation. To make a dataset available for read-only querying using Lakehouse Federation, you create the following:

- A connection, a securable object in Unity Catalog that specifies a path and credentials for accessing an external database system.
- A foreign [catalog](catalog.md)

This resource manages connections in Unity Catalog

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

## Argument Reference

The following arguments are supported:

- `name` - Name of the Connection.
- `connection_type` - Connection type. `BIGQUERY` `MYSQL` `POSTGRESQL` `SNOWFLAKE` `REDSHIFT` `SQLDW` `SQLSERVER`, `SALESFORCE` or `DATABRICKS` are supported. [Up-to-date list of connection type supported](https://docs.databricks.com/query-federation/index.html#supported-data-sources)
- `options` - The key value of options required by the connection, e.g. `host`, `port`, `user`, `password` or `GoogleServiceAccountKeyJson`. Please consult the [documentation](https://docs.databricks.com/query-federation/index.html#supported-data-sources) for the required option.
- `owner` - (Optional) Name of the connection owner.
- `properties` -  (Optional) Free-form connection properties.
- `comment` - (Optional) Free-form text.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this connection in form of `<metastore_id>|<name>`.

## Import

This resource can be imported by `id`:

```bash
terraform import databricks_connection.this '<metastore_id>|<name>'
```
