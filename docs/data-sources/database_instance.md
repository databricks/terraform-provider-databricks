---
subcategory: "Database Instances"
---
# databricks_database_instance Data Source
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

## Attributes
The following attributes are exported:
* `capacity` (string) - The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `creation_time` ([time.Time](../../README.md#well-known-types)) - The timestamp when the instance was created
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