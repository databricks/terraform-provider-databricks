---
subcategory: "Database Instances"
---
# databricks_database_instance Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

Lakebase Database Instances are managed Postgres instances, composed of a primary Postgres compute instance and 0 or more read replica instances.

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

**Note:** Each workspace is limited to 10 Database Instances in total.


## Example Usage
### Basic Example

This example creates a simple Database Instance with the specified name and capacity.

```hcl
resource "databricks_database_instance" "this" {
  name = "my-database-instance"
  capacity = "CU_2"
}
```

### Example with Readable Secondaries

This example creates a Database Instance with readable secondaries (and HA) enabled.

```hcl
resource "databricks_database_instance" "this" {
  name = "my-database-instance"
  capacity = "CU_2"
  node_count = 2
  enable_readable_secondaries = true
}
```

### Example Child Instance Created From Parent

This example creates a child Database Instance from a specified parent Database Instance at the current point in time.

```hcl
resource "databricks_database_instance" "child" {
  name = "my-database-instance"
  capacity = "CU_2"
  parent_instance_ref = {
    name = "my-parent-instance"
  }
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the instance. This is the unique identifier for the instance
* `capacity` (string, optional) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `enable_pg_native_login` (boolean, optional) - Whether to enable PG native password login on the instance. Defaults to false
* `enable_readable_secondaries` (boolean, optional) - Whether to enable secondaries to serve read-only traffic. Defaults to false
* `node_count` (integer, optional) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries. This field is input only, see effective_node_count for the output
* `parent_instance_ref` (DatabaseInstanceRef, optional) - The ref of the parent instance. This is only available if the instance is
  child instance.
  Input: For specifying the parent instance to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `retention_window_in_days` (integer, optional) - The retention window for the instance. This is the time window in days
  for which the historical data is retained. The default value is 7 days.
  Valid values are 2 to 35 days
* `stopped` (boolean, optional) - Whether to stop the instance. An input only param, see effective_stopped for the output
* `purge_on_delete` (boolean, optional) - Purge the resource on delete

### DatabaseInstanceRef
* `branch_time` (string, optional) - Branch time of the ref database instance.
  For a parent ref instance, this is the point in time on the parent instance from which the
  instance was created.
  For a child ref instance, this is the point in time on the instance from which the child
  instance was created.
  Input: For specifying the point in time to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `lsn` (string, optional) - User-specified WAL LSN of the ref database instance.
  
  Input: For specifying the WAL LSN to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `name` (string, optional) - Name of the ref database instance

## Attributes
In addition to the above arguments, the following attributes are exported:
* `child_instance_refs` (list of DatabaseInstanceRef) - The refs of the child instances. This is only available if the instance is
  parent instance
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `effective_capacity` (string, deprecated) - Deprecated. The sku of the instance; this field will always match the value of capacity
* `effective_enable_pg_native_login` (boolean) - Whether the instance has PG native password login enabled
* `effective_enable_readable_secondaries` (boolean) - Whether secondaries serving read-only traffic are enabled. Defaults to false
* `effective_node_count` (integer) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries
* `effective_retention_window_in_days` (integer) - The retention window for the instance. This is the time window in days
  for which the historical data is retained
* `effective_stopped` (boolean) - Whether the instance is stopped
* `pg_version` (string) - The version of Postgres running on the instance
* `read_only_dns` (string) - The DNS endpoint to connect to the instance for read only access. This is only available if
  enable_readable_secondaries is true
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `state` (string) - The current state of the instance. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
* `uid` (string) - An immutable UUID identifier for the instance

### DatabaseInstanceRef
* `effective_lsn` (string) - For a parent ref instance, this is the LSN on the parent instance from which the
  instance was created.
  For a child ref instance, this is the LSN on the instance from which the child instance
  was created
* `uid` (string) - Id of the ref database instance

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_database_instance.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_database_instance "name"
```