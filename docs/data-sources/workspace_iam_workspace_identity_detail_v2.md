---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_identity_detail_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads the workspace identity detail for a principal — its principal type, assignment type (direct or indirect), and workspace activity status. The `name` uses the format `principals/{principal_id}/workspace-identity-detail`.


## Example Usage
Example usage:

```hcl
data "databricks_workspace_iam_workspace_identity_detail_v2" "this" {
  name = "principals/987654321/workspace-identity-detail"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the workspace identity detail, in the format
  `principals/{principal_id}/workspace-identity-detail`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `assignment_type` (string) - Whether the principal is assigned to the workspace directly or has access through a group. Possible values are: `DIRECT`, `INDIRECT`
* `name` (string) - The resource name of the workspace identity detail, in the format
  `principals/{principal_id}/workspace-identity-detail`
* `principal_id` (integer) - The Databricks ID of the user, service principal, or group
* `principal_type` (string) - Whether the principal is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `workspace_identity_status` (string) - Whether the principal is active in the workspace. Groups are always `ACTIVE`. Possible values are: `ACTIVE`, `INACTIVE`