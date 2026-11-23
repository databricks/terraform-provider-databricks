---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_service_principals_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the service principals in an account.


## Example Usage
Example usage:

List all account service principals.

```hcl
data "databricks_account_iam_service_principals_v2" "this" {
  parent = "accounts/00000000-0000-0000-0000-000000000000"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the account to list service principals in, in the format
  `accounts/{account_id}`
* `filter` (string, optional) - Filters service principals by `application_id` or `external_id`, for example
  `application_id eq "<application-id>"`. Put the value in quotes
* `page_size` (integer, optional) - The maximum number of service principals to return, up to 1000. The default is 1000. A page
  can contain fewer service principals


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