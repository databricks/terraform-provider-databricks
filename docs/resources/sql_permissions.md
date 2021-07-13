---
subcategory: "Security"
---
# databricks_sql_permissions Resource

This resource manages data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html). In order to enable Table Access control, you have to login to the workspace as administrator, go to `Admin Console`, pick `Access Control` tab, click on `Enable` button in `Table Access Control` section, and click `Confirm`. The security guarantees of table access control **will only be effective if cluster access control is also turned on**. Please make sure that no users can create clusters in your workspace and all [databricks_cluster](cluster.md) have approximately the following configuration:

```hcl
resource "databricks_cluster" "cluster_with_table_access_control" {
  // ...

  spark_conf = {
    "spark.databricks.acl.dfAclsEnabled": "true",
    "spark.databricks.repl.allowedLanguages": "python,sql",
  }

}  
```

It could be combined with creation of High-Concurrency and Single-Node clusters - in this case it should have corresponding `custom_tags` and `spark.databricks.cluster.profile` in Spark configuration as described in [documentation for `databricks_cluster` resource](cluster.md).

The created cluster could be referred by providing its ID as `cluster_id` property.


```hcl
resource "databricks_sql_permissions" "foo_table" {
  cluster_id = databricks_cluster.cluster_name.id
  ...
}
```

## Example Usage

The following resource definition will enforce access control on a table by executing the following SQL queries on a special auto-terminating cluster it would create for this operation:

* ```SHOW GRANT ON TABLE `default`.`foo` ```
* ```REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM ... every group and user that has access to it ...```
* ```GRANT MODIFY, SELECT ON TABLE `default`.`foo` TO `serge@example.com` ```
* ```GRANT SELECT ON TABLE `default`.`foo` TO `special group` ```

```hcl
resource "databricks_sql_permissions" "foo_table" {
    table = "foo"

    privilege_assignments {
        principal = "serge@example.com"
        privileges = ["SELECT", "MODIFY"]
    }

    privilege_assignments {
        principal = "special group"
        privileges = ["SELECT"]
    }
}
```

## Argument Reference

The following arguments are available to specify the data object you need to enforce access controls on. You must specify only one of those arguments (except for `table` and `view`), otherwise resource creation will fail.

* `database` - Name of the database. Has default value of `default`.
* `table` - Name of the table. Can be combined with `database`. 
* `view` - Name of the view. Can be combined with `database`. 
* `catalog` - (Boolean) If this access control for the entire catalog. Defaults to `false`.
* `any_file` - (Boolean) If this access control for reading any file. Defaults to `false`.
* `anonymous_function` - (Boolean) If this access control for using anonymous function. Defaults to `false`.

### `privilege_assignments` blocks

You must specify one or many `privilege_assignments` configuration blocks to declare `privileges` to a `principal`, which corresponds to `display_name` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name). Terraform would ensure that only those principals and privileges defined in the resource are applied for the data object and would remove anything else. It would not remove any transitive privileges. `DENY` statements are intentionally not supported. Every `privilege_assignments` has the following required arguments:

* `principal` - `display_name` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name).
* `privileges` - set of available privilege names in upper case.

[Available](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html) privilege names are:

* `SELECT` - gives read access to an object.
* `CREATE` - gives the ability to create an object (for example, a table in a database).
* `MODIFY` - gives the ability to add, delete, and modify data to or from an object.
* `USAGE` - do not give any abilities, but is an additional requirement to perform any action on a database object.
* `READ_METADATA` - gives the ability to view an object and its metadata.
* `CREATE_NAMED_FUNCTION` - gives the ability to create a named UDF in an existing catalog or database.
* `MODIFY_CLASSPATH` - gives the ability to add files to the Spark class path.

-> Even though the value `ALL PRIVILEGES` is mentioned in Table ACL documentation, it's not recommended to use it from terraform, as it may result in unnecessary state updates.

## Import

The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:

* `table/default.foo` - table `foo` in a `default` database. Database is always mandatory.
* `view/bar.foo` - view `foo` in `bar` database.
* `database/bar` - `bar` database.
* `catalog/` - entire catalog. `/` suffix is mandatory.
* `any file/` - direct access to any file. `/` suffix is mandatory.
* `anonymous function/` - anonymous function. `/` suffix is mandatory.

```bash
$ terraform import databricks_sql_permissions.foo /<object-type>/<object-name>
```
