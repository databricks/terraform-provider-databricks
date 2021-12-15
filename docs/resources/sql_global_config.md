---
subcategory: "Databricks SQL"
---
# databricks_sql_global_config Resource

This resource configures the security policy, [databricks_instance_profile](instance_profile.md), and data access properties for all [databricks_sql_endpoint](sql_endpoint.md) of workspace. *Please note that changing parameters of this resources will restart all running [databricks_sql_endpoint](sql_endpoint.md).*  To use this resource you need to be an administrator.

## Example usage

```hcl
resource "databricks_sql_global_config" "this" {
  security_policy = "DATA_ACCESS_CONTROL"
  instance_profile_arn = "arn:...."
  data_access_config = {
    "spark.sql.session.timeZone": "UTC"
  }
}
```

## Argument Reference

The following arguments are supported (see [documentation](https://docs.databricks.com/sql/api/sql-endpoints.html#global-edit) for more details):

* `security_policy` (Optional, String) - The policy for controlling access to datasets. Default value: `DATA_ACCESS_CONTROL`, consult documentation for list of possible values
* `data_access_config` (Optional, Map) - data access configuration for [databricks_sql_endpoint](sql_endpoint.md), such as configuration for an external Hive metastore, Hadoop Filesystem configuration, etc.  Please note that the list of supported configuration properties is limited, so refer to the [documentation](https://docs.databricks.com/sql/admin/data-access-configuration.html#supported-properties) for a full list.  Apply will fail if you're specifying not permitted configuration.
* `instance_profile_arn` (Optional, String) - [databricks_instance_profile](instance_profile.md) used to access storage from [databricks_sql_endpoint](sql_endpoint.md). Please note that this parameter is only for AWS, and will generate an error if used on other clouds.

## Import

You can import a `databricks_sql_global_config` resource with command like the following (you need to use `global` as ID):

```bash
$ terraform import databricks_sql_global_config.this global
```
