---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_direct_group_member_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `name` (string, required) - The resource name of the group membership. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}
  - Workspace-scoped: groups/{group_id}/direct-members/{principal_id}

## Attributes
The following attributes are exported:
* `display_name` (string) - Display name of the principal
* `external_id` (string) - The external ID of the principal in the customer's IdP
* `group_id` (integer) - The internal ID of the group this member belongs to
* `membership_source` (string) - The source of group membership (internal or from identity provider). Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `name` (string) - The resource name of the group membership. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}
  - Workspace-scoped: groups/{group_id}/direct-members/{principal_id}
* `parent` (string) - The resource name of the group that contains the membership. The format depends on the API
  that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}
  - Workspace-scoped: groups/{group_id}
* `principal_id` (integer) - Internal ID of the principal in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`