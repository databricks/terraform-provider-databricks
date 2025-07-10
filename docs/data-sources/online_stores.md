---
subcategory: "Machine Learning"
---
# databricks_online_stores Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return. Defaults to 100 if not specified



## Attributes
This data source exports a single attribute, `online_stores`. It is a list of resources, each with the following attributes:
* `capacity` (string) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `name` (string) - The name of the online store. This is the unique identifier for the online store
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
