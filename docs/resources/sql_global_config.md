---
subcategory: "Databricks SQL"
---
# databricks_sql_global_config Resource

This resource configures the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace. *Please note that changing parameters of this resource will restart all running [databricks_sql_endpoint](sql_endpoint.md).*  To use this resource you need to be an administrator.

## Example usage

### AWS example

```hcl
resource "databricks_sql_global_config" "this" {
  security_policy      = "DATA_ACCESS_CONTROL"
  instance_profile_arn = "arn:...."
  data_access_config = {
    "spark.sql.session.timeZone" : "UTC"
  }
}
```

### Azure example

For Azure you should use the `data_access_config` to provide the service principal configuration. You can use the Databricks SQL Admin Console UI to help you generate the right configuration values.

```hcl
resource "databricks_sql_global_config" "this" {
  security_policy = "DATA_ACCESS_CONTROL"
  data_access_config = {
    "spark.hadoop.fs.azure.account.auth.type" : "OAuth",
    "spark.hadoop.fs.azure.account.oauth.provider.type" : "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
    "spark.hadoop.fs.azure.account.oauth2.client.id" : "${var.tenant_id}",
    "spark.hadoop.fs.azure.account.oauth2.client.secret" : "{{secrets/${local.secret_scope}/${local.secret_key}}}",
    "spark.hadoop.fs.azure.account.oauth2.client.endpoint" : "https://login.microsoftonline.com/${var.tenant_id}/oauth2/token"
  }
  sql_config_params = {
    "ANSI_MODE" : "true"
  }
}
```


## Argument Reference

The following arguments are supported (see [documentation](https://docs.databricks.com/sql/api/sql-endpoints.html#global-edit) for more details):

* `security_policy` (Optional, String) - The policy for controlling access to datasets. Default value: `DATA_ACCESS_CONTROL`, consult documentation for list of possible values
* `data_access_config` (Optional, Map) - Data access configuration for [databricks_sql_endpoint](sql_endpoint.md), such as configuration for an external Hive metastore, Hadoop Filesystem configuration, etc.  Please note that the list of supported configuration properties is limited, so refer to the [documentation](https://docs.databricks.com/sql/admin/data-access-configuration.html#supported-properties) for a full list.  Apply will fail if you're specifying not permitted configuration.
* `enable_serverless_compute` (optional, Boolean) - Allows the possibility to create Serverless SQL warehouses. Default value: false.
* `instance_profile_arn` (Optional, String) - [databricks_instance_profile](instance_profile.md) used to access storage from [databricks_sql_endpoint](sql_endpoint.md). Please note that this parameter is only for AWS, and will generate an error if used on other clouds. 
* `sql_config_params` (Optional, Map) - SQL Configuration Parameters let you override the default behavior for all sessions with all endpoints.

## Import

You can import a `databricks_sql_global_config` resource with command like the following (you need to use `global` as ID):

```bash
$ terraform import databricks_sql_global_config.this global
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
