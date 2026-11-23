---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_service_principal_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads a single service principal by its resource name, `service-principals/{service_principal_id}`, using a workspace-scoped provider.


## Example Usage
Example usage:

Look up a service principal by its internal ID, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_service_principal_v2" "this" {
  name = "service-principals/123456789"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the service principal, in one of these formats:
  - Account APIs: `accounts/{account_id}/service-principals/{service_principal_id}`
  - Workspace APIs: `service-principals/{service_principal_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
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