---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_external_service_principal_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Retrieves an external service principal — a service principal that can be synced from your identity provider (IdP) — with the given external ID from the customer's IdP. If the service principal does not exist in the account, it will be created. If the customer is not onboarded onto Automatic Identity Management (AIM), this returns an error.

The `name` uses the format `accounts/{account_id}/external-service-principals/{external_service_principal_id}`, where `external_service_principal_id` is the service principal's object ID in the IdP (for example, a Microsoft Entra ID object ID).

~> **Note** Reading this data source has a side effect: it resolves the service principal against the IdP and provisions the service principal in the account if it does not already exist.

-> **Note** This data source can only be used with an account-level provider.


## Example Usage
Referring to an external service principal by its resource name:

```hcl
data "databricks_account_iam_external_service_principal_v2" "example" {
  name = "accounts/00000000-0000-0000-0000-000000000000/external-service-principals/11111111-2222-3333-4444-555555555555"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the external service principal. The format depends on
  the API that returned it:
  - Account-scoped: accounts/{account_id}/external-service-principals/{external_service_principal_id}
  - Workspace-scoped: external-service-principals/{external_service_principal_id}

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID, from Databricks
* `account_sp_status` (string) - The activity status of the service principal in the Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `application_id` (string) - Application ID of the service principal, from the customer's IdP
* `display_name` (string) - Display name of the service principal, from the customer's IdP
* `external_service_principal_id` (string) - The external ID of the service principal in the customer's IdP
* `internal_id` (string) - Internal servicePrincipalId of the service principal in Databricks
* `name` (string) - The resource name of the external service principal. The format depends on
  the API that returned it:
  - Account-scoped: accounts/{account_id}/external-service-principals/{external_service_principal_id}
  - Workspace-scoped: external-service-principals/{external_service_principal_id}