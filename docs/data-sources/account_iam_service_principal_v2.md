---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_service_principal_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single account service principal by its resource name, `accounts/{account_id}/service-principals/{service_principal_id}`.


## Example Usage
Example usage:

Look up an account service principal by its internal ID.

```hcl
data "databricks_account_iam_service_principal_v2" "this" {
  name = "accounts/00000000-0000-0000-0000-000000000000/service-principals/123456789"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the service principal. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/service-principals/{service_principal_id}
  - Workspace-scoped: service-principals/{service_principal_id}

## Attributes
The following attributes are exported:
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