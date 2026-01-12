---
subcategory: "Settings"
---
# databricks_account_network_policy Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

Network policies control which network destinations can be accessed from the Databricks environment. 

Each Databricks account includes a default policy named `default-policy`. This policy is:

- Associated with any workspace lacking an explicit network policy assignment
- Automatically associated with each newly created workspace
- Reserved and cannot be deleted, but can be updated to customize the default network access rules for your account

The `default-policy` provides a baseline security configuration that ensures all workspaces have network access controls in place.

-> **Note** This resource can only be used with an account-level provider!

## Example Usage
```hcl
resource "databricks_account_network_policy" "example_network_policy" {
  network_policy_id    = "example-network-policy"
  egress = {
    network_access = {
      restriction_mode = "RESTRICTED_ACCESS"
      allowed_internet_destinations = [
        {
            destination = "example.com"
            internet_destination_type = "DNS_NAME"
        }
      ],
      allowed_storage_destinations = [
        {
          bucket_name = "example-aws-cloud-storage"
          region = "us-west-1"
          storage_destination_type = "AWS_S3"
        }
      ]
      policy_enforcement = {
        enforcement_mode = "ENFORCED"
      }
    }
  }
}
```

## Arguments
The following arguments are supported:
* `account_id` (string, optional) - The associated account ID for this Network Policy object
* `egress` (NetworkPolicyEgress, optional) - The network policies applying for egress traffic
* `network_policy_id` (string, optional) - The unique identifier for the network policy

### EgressNetworkPolicyNetworkAccessPolicy
* `restriction_mode` (string, required) - The restriction mode that controls how serverless workloads can access the internet. Possible values are: `FULL_ACCESS`, `RESTRICTED_ACCESS`
* `allowed_internet_destinations` (list of EgressNetworkPolicyNetworkAccessPolicyInternetDestination, optional) - List of internet destinations that serverless workloads are allowed to access when in RESTRICTED_ACCESS mode
* `allowed_storage_destinations` (list of EgressNetworkPolicyNetworkAccessPolicyStorageDestination, optional) - List of storage destinations that serverless workloads are allowed to access when in RESTRICTED_ACCESS mode
* `policy_enforcement` (EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement, optional) - Optional. When policy_enforcement is not provided, we default to ENFORCE_MODE_ALL_SERVICES

### EgressNetworkPolicyNetworkAccessPolicyInternetDestination
* `destination` (string, optional) - The internet destination to which access will be allowed. Format dependent on the destination type
* `internet_destination_type` (string, optional) - The type of internet destination. Currently only DNS_NAME is supported. Possible values are: `DNS_NAME`

### EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
* `dry_run_mode_product_filter` (list of string, optional) - When empty, it means dry run for all products.
  When non-empty, it means dry run for specific products and for the other products, they will run in enforced mode
* `enforcement_mode` (string, optional) - The mode of policy enforcement. ENFORCED blocks traffic that violates policy,
  while DRY_RUN only logs violations without blocking. When not specified,
  defaults to ENFORCED. Possible values are: `DRY_RUN`, `ENFORCED`

### EgressNetworkPolicyNetworkAccessPolicyStorageDestination
* `azure_storage_account` (string, optional) - The Azure storage account name
* `azure_storage_service` (string, optional) - The Azure storage service type (blob, dfs, etc.)
* `bucket_name` (string, optional)
* `region` (string, optional)
* `storage_destination_type` (string, optional) - The type of storage destination. Possible values are: `AWS_S3`, `AZURE_STORAGE`, `GOOGLE_CLOUD_STORAGE`

### NetworkPolicyEgress
* `network_access` (EgressNetworkPolicyNetworkAccessPolicy, optional) - The access policy enforced for egress traffic to the internet



## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "network_policy_id"
  to = databricks_account_network_policy.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_network_policy.this "network_policy_id"
```