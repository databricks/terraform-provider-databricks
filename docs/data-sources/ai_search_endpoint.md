---
subcategory: "AI Search"
---
# databricks_ai_search_endpoint Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Name of the AI Search endpoint. Server-assigned full resource path
  (`workspaces/{workspace}/endpoints/{endpoint}`) on output. On create, the
  user-supplied short name is conveyed via `CreateEndpointRequest.endpoint_id`;
  the server composes the full `name` and returns it on the response
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `budget_policy_id` (string) - The user-selected budget policy id for the endpoint
* `create_time` (string) - Time the endpoint was created
* `creator` (string) - Creator of the endpoint
* `custom_tags` (list of CustomTag) - The custom tags assigned to the endpoint
* `effective_budget_policy_id` (string) - The budget policy id applied to the endpoint
* `endpoint_status` (EndpointStatus) - Current status of the endpoint
* `endpoint_type` (string) - Type of endpoint. Required on create and immutable thereafter. Possible values are: `STANDARD`, `STORAGE_OPTIMIZED`
* `id` (string) - Unique identifier of the endpoint
* `index_count` (integer) - Number of indexes on the endpoint
* `last_updated_user` (string) - User who last updated the endpoint
* `name` (string) - Name of the AI Search endpoint. Server-assigned full resource path
  (`workspaces/{workspace}/endpoints/{endpoint}`) on output. On create, the
  user-supplied short name is conveyed via `CreateEndpointRequest.endpoint_id`;
  the server composes the full `name` and returns it on the response
* `replica_count` (integer) - The client-supplied desired number of replicas for the endpoint, applied at
  create/update time. Mutually exclusive with `target_qps`
* `scaling_info` (EndpointScalingInfo) - Scaling information for the endpoint
* `target_qps` (integer) - Target QPS for the endpoint. Mutually exclusive with `replica_count`. Best-effort;
  the system does not guarantee this QPS will be achieved
* `throughput_info` (EndpointThroughputInfo) - Throughput information for the endpoint
* `update_time` (string) - Time the endpoint was last updated
* `usage_policy_id` (string) - The usage policy id applied to the endpoint

### CustomTag
* `key` (string) - Key field for an AI Search endpoint tag
* `value` (string) - [Optional] Value field for an AI Search endpoint tag

### EndpointScalingInfo
* `requested_target_qps` (integer) - The requested QPS target for the endpoint. Best-effort; the system does not
  guarantee this QPS will be achieved
* `state` (string) - The current state of the scaling change request. Possible values are: `SCALING_CHANGE_APPLIED`, `SCALING_CHANGE_IN_PROGRESS`, `SCALING_CHANGE_UNSPECIFIED`

### EndpointStatus
* `message` (string) - Human-readable detail about the endpoint's current state or the reason for a state transition
* `state` (string) - Current lifecycle state of the endpoint. See `State` for the meaning of each value. Possible values are: `DELETED`, `OFFLINE`, `ONLINE`, `PROVISIONING`, `RED_STATE`, `YELLOW_STATE`

### EndpointThroughputInfo
* `change_request_message` (string) - Additional information about the throughput change request
* `change_request_state` (string) - The state of the most recent throughput change request. Possible values are: `CHANGE_ADJUSTED`, `CHANGE_FAILED`, `CHANGE_IN_PROGRESS`, `CHANGE_REACHED_MAXIMUM`, `CHANGE_REACHED_MINIMUM`, `CHANGE_SUCCESS`
* `current_concurrency` (number) - The current concurrency (total CPU) allocated to the endpoint
* `current_concurrency_utilization_percentage` (number) - The current utilization of concurrency as a percentage (0-100)
* `current_num_replicas` (integer) - The current number of replicas allocated to the endpoint
* `maximum_concurrency_allowed` (number) - The maximum concurrency allowed for this endpoint
* `minimal_concurrency_allowed` (number) - The minimum concurrency allowed for this endpoint
* `requested_concurrency` (number) - The requested concurrency (total CPU) for the endpoint
* `requested_num_replicas` (integer) - The requested number of replicas for the endpoint