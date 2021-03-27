---
subcategory: "Security"
---
# databricks_table_acl Resource

This resource manages data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html). 

## Example Usage

The following resource definition will enforce access control on a table by executing the following SQL queries on a special auto-terminating cluster it would create for this operation:

* ```SHOW GRANT ON TABLE `default`.`foo` ```
* ```REVOKE ALL PRIVILEGES ON TABLE `default`.`foo` FROM ... every group and user that has access to it ...```
* ```GRANT READ, MODIFY, SELECT ON TABLE `default`.`foo` TO `serge@example.com` ```
* ```GRANT READ, MODIFY, SELECT ON TABLE `default`.`foo` TO `special group` ```
* ```DENY SELECT, READ ON TABLE `default`.`foo` TO `users` ```

```hcl
resource "databricks_table_acl" "foo_table" {
    table = "foo"

    grant {
        principal = "serge@example.com"
        privileges = ["SELECT", "READ", "MODIFY"]
    }

    grant {
        principal = "special group"
        privileges = ["SELECT", "READ", "MODIFY"]
    }

    deny {
        principal = "users"
        privileges = ["SELECT", "READ"]
    }
}
```

## Argument Reference

The following arguments are available to specify the data object you need to enfore access controls on. You must specify only one of those arguments (except for `table` and `view`), otherwise resource creation will fail.

* `database` - Name of the database. Has default value of `default`.
* `table` - Name of the table. Can be combined with `database`. 
* `view` - Name of the view. Can be combined with `database`. 
* `catalog` - (Boolean) If this access control for the entire catalog. Defaults to `false`.
* `any_file` - (Boolean) If this access control for reading any file. Defaults to `false`.
* `anonymous_function` - (Boolean) If this access control for using anonymous function. Defaults to `false`.

### `grant` and `deny` blocks

You must specify one or many `grant` and/or `deny` configuration blocks to declare `privileges` to a `principal`, which corresponds to `display_name` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name). Terraform would ensure that only those principals and priviliges defined in the resource are applied for the data object and would remove anything else. It would not remove any transitive privileges. Every `grant` or `deny` has the following required arguments:

* `principal` - `display_name` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name).
* `privileges` - set of available privilege names in upper case.

[Available](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html) piivilege names are:

* `SELECT` - gives read access to an object.
* `CREATE` - gives ability to create an object (for example, a table in a database).
* `MODIFY` - gives ability to add, delete, and modify data to or from an object.
* `USAGE` - does not give any abilities, but is an additional requirement to perform any action on a database object.
* `READ_METADATA` - gives ability to view an object and its metadata.
* `CREATE_NAMED_FUNCTION` - gives ability to create a named UDF in an existing catalog or database.
* `MODIFY_CLASSPATH` - gives ability to add files to the Spark class path.
* `ALL PRIVILEGES` - gives all privileges (is translated into all the above privileges).


## Import

The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:

* `table/default.foo` - table `foo` in a `default` database. Database is always mandatory.
* `view/bar.foo` - view `foo` in `bar` database.
* `database/bar` - `bar` database.
* `catalog/` - entire catalog. `/` suffix is mandatory.
* `any file/` - direct access to any file. `/` suffix is mandatory.
* `anonymous function/` - anonymous function. `/` suffix is mandatory.

```bash
$ terraform import databricks_table_acl.foo /<object-type>/<object-name>
```
