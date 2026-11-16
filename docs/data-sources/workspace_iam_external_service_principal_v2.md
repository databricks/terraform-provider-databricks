---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_external_service_principal_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Retrieves an external service principal — a service principal that can be synced from your identity provider (IdP) — with the given external ID from the customer's IdP, scoped to a workspace. If the service principal does not exist in the account, it will be created. If the customer is not onboarded onto Automatic Identity Management (AIM), this returns an error.

The `name` uses the format `external-service-principals/{external_service_principal_id}`, where `external_service_principal_id` is the service principal's object ID in the IdP (for example, a Microsoft Entra ID object ID).

~> **Note** Reading this data source has a side effect: it resolves the service principal against the IdP and provisions the service principal in the account if it does not already exist. Provisioning happens at the account level; it does not assign the service principal to the workspace.

-> **Note** This data source can be used with an account-level or workspace-level provider. With an account-level provider, a `workspace_id` is required — set it in the `provider_config` block (or via the provider's `workspace_id` attribute). With a workspace-level provider, `workspace_id` is optional and defaults to the provider's workspace.


## Example Usage
Referring to an external service principal within a workspace, using an account-level provider with the target workspace selected via `provider_config`:

```hcl
data "databricks_workspace_iam_external_service_principal_v2" "example" {
  name = "external-service-principals/11111111-2222-3333-4444-555555555555"
  provider_config = {
    workspace_id = "1234567890123456"
  }
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the external service principal, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/external-service-principals/{external_service_principal_id}`
  - Workspace APIs: `external-service-principals/{external_service_principal_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the service principal
* `account_sp_status` (string) - Whether the service principal is active in the account. Possible values are: `ACTIVE`, `INACTIVE`
* `application_id` (string) - The service principal's application ID, from your IdP
* `display_name` (string) - The service principal's display name, from your IdP
* `external_service_principal_id` (string) - The ID of the service principal in your IdP
* `internal_id` (string) - The Databricks ID of the service principal. Use it as the `service_principal_id` in the
  service principal APIs
* `name` (string) - The resource name of the external service principal, in one of these formats:
  - Account APIs:
  `accounts/{account_id}/external-service-principals/{external_service_principal_id}`
  - Workspace APIs: `external-service-principals/{external_service_principal_id}`