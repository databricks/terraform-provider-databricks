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
* `filter` (string, optional) - Filters service principals by `application_id` or `external_id`, for example
  `application_id eq "<application-id>"`. Put the value in quotes
* `page_size` (integer, optional) - The maximum number of service principals to return, up to 1000. The default is 1000. A page
  can contain fewer service principals
* `parent` (string, optional) - The resource name of the account to list service principals in, in the format
  `accounts/{account_id}`. Defaults to the account that contains this workspace. Any other
  account is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `service_principals`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The ID of the account that contains the service principal
* `account_sp_status` (string) - Whether the service principal is active in the account. Possible values are: `ACTIVE`, `INACTIVE`
* `application_id` (string) - The service principal's application ID, a UUID. If you don't set it when you create the
  service principal, Databricks generates one. It must be unique in the account and can't be
  changed
* `display_name` (string) - The service principal's display name
* `external_id` (string) - The ID of the service principal in your identity provider (IdP)
* `name` (string) - The resource name of the service principal, in one of these formats:
  - Account APIs: `accounts/{account_id}/service-principals/{service_principal_id}`
  - Workspace APIs: `service-principals/{service_principal_id}`
* `parent` (string) - The resource name of the account that contains the service principal, in the format
  `accounts/{account_id}`. Workspace APIs return the same format
* `service_principal_id` (string) - The Databricks ID of the service principal