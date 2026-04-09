---
subcategory: "Settings"
---
# databricks_account_network_policy Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `ingress` (CustomerFacingIngressNetworkPolicy) - The network policies applying for ingress traffic
* `ingress_dry_run` (CustomerFacingIngressNetworkPolicy) - The ingress policy for dry run mode. Dry run will always run even if the request
  is allowed by the ingress policy. When this field is set, the policy will be evaluated
  and emit logs only without blocking requests
* `network_policy_id` (string) - The unique identifier for the network policy

### CustomerFacingIngressNetworkPolicy
* `public_access` (CustomerFacingIngressNetworkPolicyPublicAccess)

### CustomerFacingIngressNetworkPolicyAppsDestination
* `all_destinations` (boolean) - Must be set to true

### CustomerFacingIngressNetworkPolicyAuthentication
* `identities` (list of CustomerFacingIngressNetworkPolicyAuthenticationIdentity) - Valid only when IdentityType is IDENTITY_TYPE_SELECTED_IDENTITIES
* `identity_type` (string) - Possible values are: `IDENTITY_TYPE_ALL_SERVICE_PRINCIPALS`, `IDENTITY_TYPE_ALL_USERS`, `IDENTITY_TYPE_SELECTED_IDENTITIES`

### CustomerFacingIngressNetworkPolicyAuthenticationIdentity
* `principal_id` (integer)
* `principal_type` (string) - Possible values are: `PRINCIPAL_TYPE_SERVICE_PRINCIPAL`, `PRINCIPAL_TYPE_USER`

### CustomerFacingIngressNetworkPolicyIpRanges
* `ip_ranges` (list of string) - We only support IPv4 and IPv4 CIDR notation for now

### CustomerFacingIngressNetworkPolicyLakebaseDestination
* `all_destinations` (boolean) - Must be set to true

### CustomerFacingIngressNetworkPolicyPublicAccess
* `allow_rules` (list of CustomerFacingIngressNetworkPolicyPublicIngressRule)
* `deny_rules` (list of CustomerFacingIngressNetworkPolicyPublicIngressRule)
* `restriction_mode` (string) - Possible values are: `FULL_ACCESS`, `RESTRICTED_ACCESS`

### CustomerFacingIngressNetworkPolicyPublicIngressRule
* `authentication` (CustomerFacingIngressNetworkPolicyAuthentication)
* `destination` (CustomerFacingIngressNetworkPolicyRequestDestination)
* `label` (string) - User-provided name for this ingress rule. Helps identify which rule
  caused a request to be denied or dry-run denied
* `origin` (CustomerFacingIngressNetworkPolicyPublicRequestOrigin)

### CustomerFacingIngressNetworkPolicyPublicRequestOrigin
* `all_ip_ranges` (boolean) - Matches all IPv4 and IPv6 ranges (both public and private)
* `excluded_ip_ranges` (CustomerFacingIngressNetworkPolicyIpRanges) - Excluded means: all public IP ranges except this one
* `included_ip_ranges` (CustomerFacingIngressNetworkPolicyIpRanges) - Will not allow IP ranges with private IPs

### CustomerFacingIngressNetworkPolicyRequestDestination
* `all_destinations` (boolean) - When true, match all destinations, no other destination fields can be set.
  When not set or false, at least one specific destination must be provided
* `apps` (CustomerFacingIngressNetworkPolicyAppsDestination)
* `lakebase` (CustomerFacingIngressNetworkPolicyLakebaseDestination)
* `workspace_api` (CustomerFacingIngressNetworkPolicyWorkspaceApiDestination)
* `workspace_ui` (CustomerFacingIngressNetworkPolicyWorkspaceUiDestination)

### CustomerFacingIngressNetworkPolicyWorkspaceApiDestination
* `scopes` (list of string)

### CustomerFacingIngressNetworkPolicyWorkspaceUiDestination
* `all_destinations` (boolean) - Must be set to true

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