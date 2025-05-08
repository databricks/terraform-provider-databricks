---
subcategory: "Unity Catalog"
---
# databricks_database_instance Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the instance. This is the unique identifier for the instance

## Attributes
The following attributes are exported:
* `admin_password` (string) - Password for admin user to create. If not provided, no user will be created
* `admin_rolename` (string) - Name of the admin role for the instance. If not provided, defaults to 'databricks_admin'
* `capacity` (string) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4"
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `enable_readable_secondaries` (boolean) - Whether to enable secondaries to serve read-only traffic. Defaults to false
* `name` (string) - The name of the instance. This is the unique identifier for the instance
* `node_count` (integer) - The number of nodes in the instance, composed of 1 primary and 0 or more secondaries. Defaults to
  1 primary and 0 secondaries
* `pg_version` (string) - The version of Postgres running on the instance
* `read_only_dns` (string) - The DNS endpoint to connect to the instance for read only access. This is only available if
  enable_readable_secondaries is true
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `state` (string) - The current state of the instance. Possible values are: AVAILABLE, DELETING, FAILING_OVER, STARTING, STOPPED, UPDATING
* `stopped` (boolean) - Whether the instance is stopped
* `uid` (string) - An immutable UUID identifier for the instance