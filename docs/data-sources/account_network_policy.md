---
subcategory: "Settings"
---
# databricks_account_network_policy Data Source
This data source can be used to get a single network policy.

-> **Note** This data source can only be used with an account-level provider!

## Example Usage
Referring to a network policy by id:

```hcl
data "databricks_account_network_policy" "this" {
  policy_id = "test"
}
```

## Arguments
The following arguments are supported:
* `network_policy_id` (string, required) - The unique identifier for the network policy

## Attributes
The following attributes are exported:
* `account_id` (string) - The associated account ID for this Network Policy object
* `egress` (NetworkPolicyEgress) - The network policies applying for egress traffic
* `network_policy_id` (string) - The unique identifier for the network policy

### EgressNetworkPolicyNetworkAccessPolicy
* `allowed_internet_destinations` (list of EgressNetworkPolicyNetworkAccessPolicyInternetDestination) - List of internet destinations that serverless workloads are allowed to access when in RESTRICTED_ACCESS mode
* `allowed_storage_destinations` (list of EgressNetworkPolicyNetworkAccessPolicyStorageDestination) - List of storage destinations that serverless workloads are allowed to access when in RESTRICTED_ACCESS mode
* `policy_enforcement` (EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) - Optional. When policy_enforcement is not provided, we default to ENFORCE_MODE_ALL_SERVICES
* `restriction_mode` (string) - The restriction mode that controls how serverless workloads can access the internet. Possible values are: `FULL_ACCESS`, `RESTRICTED_ACCESS`

### EgressNetworkPolicyNetworkAccessPolicyInternetDestination
* `destination` (string) - The internet destination to which access will be allowed. Format dependent on the destination type
* `internet_destination_type` (string) - The type of internet destination. Currently only DNS_NAME is supported. Possible values are: `DNS_NAME`

### EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
* `dry_run_mode_product_filter` (list of string) - When empty, it means dry run for all products.
  When non-empty, it means dry run for specific products and for the other products, they will run in enforced mode
* `enforcement_mode` (string) - The mode of policy enforcement. ENFORCED blocks traffic that violates policy,
  while DRY_RUN only logs violations without blocking. When not specified,
  defaults to ENFORCED. Possible values are: `DRY_RUN`, `ENFORCED`

### EgressNetworkPolicyNetworkAccessPolicyStorageDestination
* `azure_storage_account` (string) - The Azure storage account name
* `azure_storage_service` (string) - The Azure storage service type (blob, dfs, etc.)
* `bucket_name` (string)
* `region` (string)
* `storage_destination_type` (string) - The type of storage destination. Possible values are: `AWS_S3`, `AZURE_STORAGE`, `GOOGLE_CLOUD_STORAGE`

### NetworkPolicyEgress
* `network_access` (EgressNetworkPolicyNetworkAccessPolicy) - The access policy enforced for egress traffic to the internet