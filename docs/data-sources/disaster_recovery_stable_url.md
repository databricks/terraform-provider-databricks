---
subcategory: "Disaster Recovery"
---
# databricks_disaster_recovery_stable_url Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/disasterrecovery)

This data source can be used to get a single stable URL by its fully qualified resource name.

-> **Note** This data source can only be used with an account-level provider!


## Example Usage
Referring to a stable URL by its resource name:

```hcl
data "databricks_disaster_recovery_stable_url" "this" {
  name = "accounts/${var.account_id}/stable-urls/accounting-stable-url"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Fully qualified resource name.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}

## Attributes
The following attributes are exported:
* `effective_workspace_id` (string) - The workspace this stable URL currently routes to. Set to
  `initial_workspace_id` at creation, advanced to the failover group's primary
  while attached (including across a failover), and preserved when the stable
  URL is detached from its failover group. Read this to see where an unattached
  stable URL points: after a failover followed by a detach it reflects the
  post-failover primary, not `initial_workspace_id`
* `failover_group_name` (string) - Fully qualified resource name of the FailoverGroup this stable URL is
  currently linked to, in the format
  `accounts/{account_id}/failover-groups/{failover_group_id}`. Empty when
  the stable URL is not attached to any failover group
* `initial_workspace_id` (string) - The workspace this stable URL is initially bound to. Used only in Create
  requests to associate the stable URL with a workspace. Not returned in
  responses
* `name` (string) - Fully qualified resource name.
  Format: accounts/{account_id}/stable-urls/{stable_url_id}
* `stable_workspace_id` (string) - The stable workspace ID for this stable URL. Generated on creation and
  immutable thereafter; identifies the URL across failovers and is the same
  value embedded in the `url` (as the `w=` query parameter for SPOG URLs,
  or in the `conn-<id>` hostname for Private-Link URLs)
* `url` (string) - The stable URL endpoint. Generated on creation and
  immutable thereafter. For non-Private-Link workspaces this is
  `https://<spog_host>/?w=<connection_id>`. For Private-Link workspaces
  this is the per-connection hostname