---
subcategory: "Security"
---
# databricks_sql_permissions Resource

-> Please switch to [databricks_grants](grants.md) with Unity Catalog to manage data access, which provides a better and faster way for managing data security. `databricks_grants` resource *doesn't require a technical cluster to perform operations*. On workspaces with Unity Catalog enabled, you may run into errors such as `Error: cannot create sql permissions: cannot read current grants: For unity catalog, please specify the catalog name explicitly. E.g. SHOW GRANT ``your.address@email.com`` ON CATALOG main`. This happens if your `default_catalog_name` was set to a UC catalog instead of `hive_metastore`. The workaround is to re-assign the metastore again with the default catalog set to `hive_metastore`. See [databricks_metastore_assignment](metastore_assignment.md).

This resource manages data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html). In order to enable Table Access control, you have to login to the workspace as administrator, go to `Admin Console`, pick the `Access Control` tab, click on the `Enable` button in the `Table Access Control` section, and click `Confirm`. The security guarantees of table access control **will only be effective if cluster access control is also turned on**. Please make sure that no users can create clusters in your workspace and all [databricks_cluster](cluster.md) have approximately the following configuration:

```hcl
resource "databricks_cluster" "cluster_with_table_access_control" {
  // ...
  spark_conf = {
    "spark.databricks.acl.dfAclsEnabled" : "true",
    "spark.databricks.repl.allowedLanguages" : "python,sql",
  }

}
```

-> This resource can only be used with a workspace-level provider!

It is required to define all permissions for a securable in a single resource, otherwise Terraform cannot guarantee config drift prevention.

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
    principal  = "serge@example.com"
    privileges = ["SELECT", "MODIFY"]
  }


  privilege_assignments {
    principal  = "special group"
    privileges = ["SELECT"]
  }
}
```

## Argument Reference

* `cluster_id` - (Optional) Id of an existing [databricks_cluster](cluster.md), where the appropriate `GRANT`/`REVOKE` commands are executed. This cluster must have the appropriate data security mode (`USER_ISOLATION` or `LEGACY_TABLE_ACL` specified). If no `cluster_id` is specified, a TACL-enabled cluster with the name `terraform-table-acl` is automatically created.

```hcl
resource "databricks_sql_permissions" "foo_table" {
  cluster_id = databricks_cluster.cluster_name.id
  #...
}
```

The following arguments are available to specify the data object you need to enforce access controls on. You must specify only one of those arguments (except for `table` and `view`), otherwise resource creation will fail.

* `database` - Name of the database. Has a default value of `default`.
* `table` - Name of the table. Can be combined with the `database`.
* `view` - Name of the view. Can be combined with the `database`.
* `catalog` - (Boolean) If this access control for the entire catalog. Defaults to `false`.
* `any_file` - (Boolean) If this access control for reading/writing any file. Defaults to `false`.
* `anonymous_function` - (Boolean) If this access control for using an anonymous function. Defaults to `false`.

### `privilege_assignments` blocks

You must specify one or many `privilege_assignments` configuration blocks to declare `privileges` to a `principal`, which corresponds to `display_name` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name). Terraform would ensure that only those principals and privileges defined in the resource are applied for the data object and would remove anything else. It would not remove any transitive privileges. `DENY` statements are intentionally not supported. Every `privilege_assignments` has the following required arguments:

* `principal` - `display_name` for a [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name), `application_id` for a [databricks_service_principal](service_principal.md).
* `privileges` - set of available privilege names in upper case.

[Available](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html) privilege names are:

* `SELECT` - gives read access to an object.
* `CREATE` - gives the ability to create an object (for example, a table in a database).
* `MODIFY` - gives the ability to add, delete, and modify data to or from an object.
* `USAGE` - do not give any abilities, but is an additional requirement to perform any action on a database object.
* `READ_METADATA` - gives the ability to view an object and its metadata.
* `CREATE_NAMED_FUNCTION` - gives the ability to create a named UDF in an existing catalog or database.
* `MODIFY_CLASSPATH` - gives the ability to add files to the Spark classpath.

-> Even though the value `ALL PRIVILEGES` is mentioned in Table ACL documentation, it's not recommended to use it from Terraform, as it may result in unnecessary state updates.

## Import

The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:

* `table/default.foo` - table `foo` in a `default` database. The `database` is always mandatory.
* `view/bar.foo` - view `foo` in `bar` database.
* `database/bar` - `bar` database.
* `catalog/` - entire catalog. `/` suffix is mandatory.
* `any file/` - direct access to any file. `/` suffix is mandatory.
* `anonymous function/` - anonymous function. `/` suffix is mandatory.

```hcl
import {
  to = databricks_sql_permissions.foo
  id = "/<object-type>/<object-name>"
}
```

Alternatively, when using `terraform` version 1.5 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_permissions.foo /<object-type>/<object-name>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group](group.md) to manage [groups in Databricks Workspace](https://docs.databricks.com/administration-guide/users-groups/groups.html) or [Account Console](https://accounts.cloud.databricks.com/) (for AWS deployments).
* [databricks_grants](grants.md) to manage data access in Unity Catalog.
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_user](user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](group.md) within the workspace.
