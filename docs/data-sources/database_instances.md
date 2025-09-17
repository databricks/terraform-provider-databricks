---
subcategory: "Database Instances"
---
# databricks_database_instances Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch the list of Database Instances within the workspace.
The list can then be accessed via the data object's `database_instances` field.


## Example Usage
Getting a list of all Database Instances:

```hcl
data "databricks_database_instances" "all" {
}
output "all_database_instances" {
  value = data.databricks_database_instances.all.database_instances
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned


## Attributes
This data source exports a single attribute, `database_instances`. It is a list of resources, each with the following attributes:
* `capacity` (string) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `child_instance_refs` (list of DatabaseInstanceRef) - The refs of the child instances. This is only available if the instance is
  parent instance
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `effective_enable_pg_native_login` (boolean) - xref AIP-129. `enable_pg_native_login` is owned by the client, while `effective_enable_pg_native_login` is owned by the server.
  `enable_pg_native_login` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_enable_pg_native_login` on the other hand will always bet set in all response messages (Create/Update/Get/List)
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
* `enable_pg_native_login` (boolean) - Whether the instance has PG native password login enabled. Defaults to true
* `enable_readable_secondaries` (boolean) - Whether to enable secondaries to serve read-only traffic. Defaults to false
* `name` (string) - The name of the instance. This is the unique identifier for the instance
* `node_count` (integer) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries
* `parent_instance_ref` (DatabaseInstanceRef) - The ref of the parent instance. This is only available if the instance is
  child instance.
  Input: For specifying the parent instance to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `pg_version` (string) - The version of Postgres running on the instance
* `read_only_dns` (string) - The DNS endpoint to connect to the instance for read only access. This is only available if
  enable_readable_secondaries is true
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `retention_window_in_days` (integer) - The retention window for the instance. This is the time window in days
  for which the historical data is retained. The default value is 7 days.
  Valid values are 2 to 35 days
* `state` (string) - The current state of the instance. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
* `stopped` (boolean) - Whether the instance is stopped
* `uid` (string) - An immutable UUID identifier for the instance

### DatabaseInstanceRef
* `branch_time` (string) - Branch time of the ref database instance.
  For a parent ref instance, this is the point in time on the parent instance from which the
  instance was created.
  For a child ref instance, this is the point in time on the instance from which the child
  instance was created.
  Input: For specifying the point in time to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `effective_lsn` (string) - xref AIP-129. `lsn` is owned by the client, while `effective_lsn` is owned by the server.
  `lsn` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_lsn` on the other hand will always bet set in all response messages (Create/Update/Get/List).
  For a parent ref instance, this is the LSN on the parent instance from which the
  instance was created.
  For a child ref instance, this is the LSN on the instance from which the child instance
  was created
* `lsn` (string) - User-specified WAL LSN of the ref database instance.
  
  Input: For specifying the WAL LSN to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `name` (string) - Name of the ref database instance
* `uid` (string) - Id of the ref database instance