---
subcategory: "Unity Catalog"
---
# databricks_database_instance Resource


## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the instance. This is the unique identifier for the instance
* `admin_password` (string, optional) - Password for admin user to create. If not provided, no user will be created
* `admin_rolename` (string, optional) - Name of the admin role for the instance. If not provided, defaults to 'databricks_admin'
* `capacity` (string, optional) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4"
* `stopped` (boolean, optional) - Whether the instance is stopped

## Attributes
In addition to the above arguments, the following attributes are exported:
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `pg_version` (string) - The version of Postgres running on the instance
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `state` (string) - The current state of the instance. Possible values are: AVAILABLE, DELETING, FAILING_OVER, STARTING, STOPPED, UPDATING
* `uid` (string) - An immutable UUID identifier for the instance

## Import
As of terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_database_instance.this
}
```

If you are using an older version of terraform, you can import the resource using cli as follows:
```sh
$ terraform import databricks_database_instance name
```