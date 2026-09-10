---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_direct_group_members_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Lists the direct members of an account group.


## Example Usage
Example usage:

List all direct members of a group.

```hcl
data "databricks_account_iam_direct_group_members_v2" "this" {
  group_id = 123456789
}
```


## Arguments
The following arguments are supported:
* `group_id` (integer, required) - Required. Internal ID of the group in Databricks whose direct members are being listed
* `page_size` (integer, optional) - The maximum number of members to return. The service may return fewer than this value.
  If not provided, defaults to 1000, which is also the maximum allowed. Requests for more than the maximum are clamped to 1000


## Attributes
This data source exports a single attribute, `direct_group_members`. It is a list of resources, each with the following attributes:
* `display_name` (string) - Display name of the principal
* `external_id` (string) - The external ID of the principal in Databricks
* `group_id` (integer) - The internal ID of the group this member belongs to
* `membership_source` (string) - The source of group membership (internal or from identity provider). Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `principal_id` (integer) - Internal ID of the principal in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`