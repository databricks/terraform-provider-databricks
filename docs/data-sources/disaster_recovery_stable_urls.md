---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_stable_urls Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/disasterrecovery)

This data source can be used to fetch the list of stable URLs in the account.

-> **Note** This data source can only be used with an account-level provider!


## Example Usage
Getting a list of all stable URLs in the account:

```hcl
data "databricks_disaster_recovery_stable_urls" "all" {
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The parent resource. Format: accounts/{account_id}
* `page_size` (integer, optional) - Maximum number of stable URLs to return per page:
  - when set to a value greater than 0, the page length is the minimum of this value
  and a server configured value;
  - when set to 0 or unset, the page length is set to a server configured value
  (recommended);
  - when set to a value less than 0, an invalid parameter error is returned


## Attributes
This data source exports a single attribute, `stable_urls`. It is a list of resources, each with the following attributes:
* `failover_group_name` (string) - Fully qualified resource name of the FailoverGroup this stable URL is
  currently linked to, in the format
  `accounts/{account_id}/failover-groups/{failover_group_id}`. Empty when
  the stable URL is not attached to any failover group
* `initial_workspace_id` (string) - The workspace this stable URL is initially bound to. Used only in Create
  requests to associate the stable URL with a workspace. Not returned in
  responses
* `name` (string) - Fully qualified resource name.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}
* `url` (string) - The stable URL endpoint. Generated on creation and
  immutable thereafter. For non-Private-Link workspaces this is
  `https://<spog_host>/?w=<connection_id>`. For Private-Link workspaces
  this is the per-connection hostname