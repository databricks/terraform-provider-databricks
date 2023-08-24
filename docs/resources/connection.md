---
subcategory: "Unity Catalog"
---
# databricks_connection (Resource)

Lakehouse Federation is the query federation platform for Databricks. Databricks uses Unity Catalog to manage query federation. To make a dataset available for read-only querying using Lakehouse Federation, you create the following:

- A connection, a securable object in Unity Catalog that specifies a path and credentials for accessing an external database system.
- A foreign [catalog](catalog.md)

This resource manages connections in Unity Catalog

## Example Usage

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

## Argument Reference

The following arguments are supported:

- `name` - Name of the Connection.
- `connection_type` - Connection type. `MYSQL` `POSTGRESQL` `SNOWFLAKE` `REDSHIFT` `SQLDW` `SQLSERVER` or `DATABRICKS` are supported. [Up-to-date list of connection type supported](https://docs.databricks.com/query-federation/index.html#supported-data-sources)
- `options` - The key value of options required by the connection, e.g. `host`, `port`, `user` and `password`.
- `owner` - (Optional) Name of the connection owner.
- `properties` -  (Optional) Free-form connection properties.
- `comment` - (Optional) Free-form text.

## Import

This resource can be imported by `name`

```bash
terraform import databricks_connection.this <connection_name>
```
