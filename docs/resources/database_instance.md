---
subcategory: "Database Instances"
---
# databricks_database_instance Resource
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
* `enable_readable_secondaries` (boolean, optional) - Whether to enable secondaries to serve read-only traffic. Defaults to false
* `node_count` (integer, optional) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries
* `parent_instance_ref` (DatabaseInstanceRef, optional) - The ref of the parent instance. This is only available if the instance is
  child instance.
  Input: For specifying the parent instance to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `retention_window_in_days` (integer, optional) - The retention window for the instance. This is the time window in days
  for which the historical data is retained. The default value is 7 days.
  Valid values are 2 to 35 days
* `stopped` (boolean, optional) - Whether the instance is stopped
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
* `effective_enable_readable_secondaries` (boolean) - xref AIP-129. `enable_readable_secondaries` is owned by the client, while `effective_enable_readable_secondaries` is owned by the server.
  `enable_readable_secondaries` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_enable_readable_secondaries` on the other hand will always bet set in all response messages (Create/Update/Get/List)
* `effective_node_count` (integer) - xref AIP-129. `node_count` is owned by the client, while `effective_node_count` is owned by the server.
  `node_count` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_node_count` on the other hand will always bet set in all response messages (Create/Update/Get/List)
* `effective_retention_window_in_days` (integer) - xref AIP-129. `retention_window_in_days` is owned by the client, while `effective_retention_window_in_days` is owned by the server.
  `retention_window_in_days` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_retention_window_in_days` on the other hand will always bet set in all response messages (Create/Update/Get/List)
* `effective_stopped` (boolean) - xref AIP-129. `stopped` is owned by the client, while `effective_stopped` is owned by the server.
  `stopped` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_stopped` on the other hand will always bet set in all response messages (Create/Update/Get/List)
* `pg_version` (string) - The version of Postgres running on the instance
* `read_only_dns` (string) - The DNS endpoint to connect to the instance for read only access. This is only available if
  enable_readable_secondaries is true
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `state` (string) - The current state of the instance. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
* `uid` (string) - An immutable UUID identifier for the instance

### DatabaseInstanceRef
* `effective_lsn` (string) - xref AIP-129. `lsn` is owned by the client, while `effective_lsn` is owned by the server.
  `lsn` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_lsn` on the other hand will always bet set in all response messages (Create/Update/Get/List).
  For a parent ref instance, this is the LSN on the parent instance from which the
  instance was created.
  For a child ref instance, this is the LSN on the instance from which the child instance
  was created
* `uid` (string) - Id of the ref database instance

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_database_instance.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_database_instance name
```