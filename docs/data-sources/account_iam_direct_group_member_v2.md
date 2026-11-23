---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_direct_group_member_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Reads a single direct group membership in an account by its resource name, `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`.


## Example Usage
Example usage:

Look up a single direct group member by group and principal ID.

```hcl
data "databricks_account_iam_direct_group_member_v2" "this" {
  name = "accounts/00000000-0000-0000-0000-000000000000/groups/123456789/direct-members/987654321"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the membership, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`
  - Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`

## Attributes
The following attributes are exported:
* `display_name` (string) - The member's display name
* `external_id` (string) - The ID of the member in your identity provider (IdP)
* `group_id` (integer) - The Databricks ID of the group
* `membership_source` (string) - Where the membership comes from. Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `name` (string) - The resource name of the membership, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`
  - Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`
* `parent` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `principal_id` (integer) - The Databricks ID of the member
* `principal_type` (string) - Whether the member is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`