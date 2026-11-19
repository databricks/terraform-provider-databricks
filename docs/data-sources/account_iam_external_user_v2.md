---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_external_user_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Retrieves an external user — a user that can be synced from your identity provider (IdP) — with the given external ID from the customer's IdP. If the user does not exist in the account, it will be created. If the customer is not onboarded onto Automatic Identity Management (AIM), this returns an error.

The `name` uses the format `accounts/{account_id}/external-users/{external_user_id}`, where `external_user_id` is the user's object ID in the IdP (for example, a Microsoft Entra ID object ID).

~> **Note** Reading this data source has a side effect: it resolves the user against the IdP and provisions the user in the account if it does not already exist.

-> **Note** This data source can only be used with an account-level provider.


## Example Usage
Referring to an external user by its resource name:

```hcl
data "databricks_account_iam_external_user_v2" "example" {
  name = "accounts/00000000-0000-0000-0000-000000000000/external-users/11111111-2222-3333-4444-555555555555"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the external user, in one of these formats:
  - Account APIs: `accounts/{account_id}/external-users/{external_user_id}`
  - Workspace APIs: `external-users/{external_user_id}`

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the user
* `account_user_status` (string) - Whether the user is active in the account. Possible values are: `ACTIVE`, `INACTIVE`
* `display_name` (string) - The user's given and family names from your IdP, separated by a space
* `external_user_id` (string) - The ID of the user in your IdP
* `full_name` (FullName) - The user's name, from your IdP
* `internal_id` (string) - The Databricks ID of the user. Use it as the `user_id` in the user APIs
* `name` (string) - The resource name of the external user, in one of these formats:
  - Account APIs: `accounts/{account_id}/external-users/{external_user_id}`
  - Workspace APIs: `external-users/{external_user_id}`
* `username` (string) - The user's username, which is their email address

### FullName
* `family_name` (string) - The user's family (last) name
* `given_name` (string) - The user's given (first) name