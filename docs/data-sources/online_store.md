---
subcategory: "Machine Learning"
---
# databricks_online_store Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the online store. This is the unique identifier for the online store

## Attributes
The following attributes are exported:
* `capacity` (string) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `name` (string) - The name of the online store. This is the unique identifier for the online store
* `read_replica_count` (integer) - The number of read replicas for the online store. Defaults to 0
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`