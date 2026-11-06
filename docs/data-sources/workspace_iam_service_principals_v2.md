---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_service_principals_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the service principals visible to a workspace, using a workspace-scoped provider.


## Example Usage
Example usage:

List all service principals visible to the workspace, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_service_principals_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `filter` (string, optional) - Optional. Allows filtering service principals by application id or external id
* `page_size` (integer, optional) - The maximum number of SPs to return. The service may return fewer than this value.
  If not provided, defaults to 1000, which is also the maximum allowed. Requests for more than the maximum are clamped to 1000
* `parent` (string, optional) - The resource name of the account to list service principals in, `accounts/{account_id}`. Optional:
  defaults to the account of the calling workspace, and any other account is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `service_principals`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The parent account ID for the service principal in Databricks
* `account_sp_status` (string) - The activity status of a service principal in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `application_id` (string) - Application ID of the service principal. Set at creation time and cannot be changed
  afterwards; when omitted, the server generates one
* `display_name` (string) - Display name of the service principal
* `external_id` (string) - ExternalId of the service principal in the customer's IdP
* `name` (string) - The resource name of the service principal. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/service-principals/{service_principal_id}
  - Workspace-scoped: service-principals/{service_principal_id}
* `parent` (string) - The resource name of the account that contains the service principal,
  `accounts/{account_id}`. Both the account-scoped and workspace-scoped APIs return it
* `service_principal_id` (string) - Internal service principal ID of the service principal in Databricks