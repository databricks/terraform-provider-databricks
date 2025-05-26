---
subcategory: "Unity Catalog"
---
# databricks_database_instance Resource
Database Instances are managed Postgres instances, composed of a primary Postgres compute instance and 0 or more read replica instances. 

### Use Cases

Database Instances can be accessed for reading and writing structured data via SQL. 

Database Instances can be mapped to Catalogs and Tables within a Unity Catalog Metastore to enable read access via REST API and Feature Store.

### Permissions

Database Instances are compute resources created at the workspace level. As such, management operations are controlled by high-level ACLs.

* CAN_USE: Users with this permission can use the database instance for creating synced tables.
* CAN_MANAGE: Users with this permission can update and delete Database Instances.

### Sizing

Database Instances can be created with the following compute sizes in terms of Capacity Units (CUs):

* 1 (small)
* 2 (medium)
* 4 (large)

**Note:** Each workspace is limited to 5 Database Instances in total.


## Example Usage
```hcl
resource "databricks_database_instance" "this" {
  name = "my-database-instance"
  capacity = "CU_2"
  stopped = false
}
```


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