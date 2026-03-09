---
subcategory: "Machine Learning"
---
# databricks_online_stores Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return. Defaults to 100 if not specified
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `online_stores`. It is a list of resources, each with the following attributes:
* `capacity` (string) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `name` (string) - The name of the online store. This is the unique identifier for the online store
* `read_replica_count` (integer) - The number of read replicas for the online store. Defaults to 0
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`
* `usage_policy_id` (string) - The usage policy applied to the online store to track billing