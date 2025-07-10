---
subcategory: "Database Instances"
---
# databricks_database_instances Data Source
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
* `creation_time` (string) - The timestamp when the instance was created
* `creator` (string) - The email of the creator of the instance
* `effective_stopped` (boolean) - xref AIP-129. `stopped` is owned by the client, while `effective_stopped` is owned by the server.
  `stopped` will only be set in Create/Update response messages if and only if the user provides the field via the request.
  `effective_stopped` on the other hand will always bet set in all response messages (Create/Update/Get/List)
* `name` (string) - The name of the instance. This is the unique identifier for the instance
* `pg_version` (string) - The version of Postgres running on the instance
* `read_write_dns` (string) - The DNS endpoint to connect to the instance for read+write access
* `state` (string) - The current state of the instance. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
* `stopped` (boolean) - Whether the instance is stopped
* `uid` (string) - An immutable UUID identifier for the instance
