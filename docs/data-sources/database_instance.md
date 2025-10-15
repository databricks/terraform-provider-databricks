---
subcategory: "Database Instances"
---
# databricks_database_instance Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single Database Instance.


## Example Usage
Referring to a Database Instance by name:

```hcl
data "databricks_database_instance" "this" {
  name = "my-database-instance"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the instance. This is the unique identifier for the instance
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

## Attributes
The following attributes are exported:
* `capacity` (string) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `child_instance_refs` (list of DatabaseInstanceRef) - The refs of the child instances. This is only available if the instance is
  parent instance
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `custom_tags` (list of CustomTag) - Custom tags associated with the instance. This field is only included on create and update responses
* `effective_capacity` (string, deprecated) - Deprecated. The sku of the instance; this field will always match the value of capacity
* `effective_custom_tags` (list of CustomTag) - The recorded custom tags associated with the instance
* `effective_enable_pg_native_login` (boolean) - Whether the instance has PG native password login enabled
* `effective_enable_readable_secondaries` (boolean) - Whether secondaries serving read-only traffic are enabled. Defaults to false
* `effective_node_count` (integer) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries
* `effective_retention_window_in_days` (integer) - The retention window for the instance. This is the time window in days
  for which the historical data is retained
* `effective_stopped` (boolean) - Whether the instance is stopped
* `effective_usage_policy_id` (string) - The policy that is applied to the instance
* `enable_pg_native_login` (boolean) - Whether to enable PG native password login on the instance. Defaults to false
* `enable_readable_secondaries` (boolean) - Whether to enable secondaries to serve read-only traffic. Defaults to false
* `name` (string) - The name of the instance. This is the unique identifier for the instance
* `node_count` (integer) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries. This field is input only, see effective_node_count for the output
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
* `stopped` (boolean) - Whether to stop the instance. An input only param, see effective_stopped for the output
* `uid` (string) - An immutable UUID identifier for the instance
* `usage_policy_id` (string) - The desired usage policy to associate with the instance

### CustomTag
* `key` (string) - The key of the custom tag
* `value` (string) - The value of the custom tag

### DatabaseInstanceRef
* `branch_time` (string) - Branch time of the ref database instance.
  For a parent ref instance, this is the point in time on the parent instance from which the
  instance was created.
  For a child ref instance, this is the point in time on the instance from which the child
  instance was created.
  Input: For specifying the point in time to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `effective_lsn` (string) - For a parent ref instance, this is the LSN on the parent instance from which the
  instance was created.
  For a child ref instance, this is the LSN on the instance from which the child instance
  was created
* `lsn` (string) - User-specified WAL LSN of the ref database instance.
  
  Input: For specifying the WAL LSN to create a child instance. Optional.
  Output: Only populated if provided as input to create a child instance
* `name` (string) - Name of the ref database instance
* `uid` (string) - Id of the ref database instance