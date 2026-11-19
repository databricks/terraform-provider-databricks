---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_identity_detail_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Manages the workspace-level activity status of a principal already assigned to the workspace. This is a singleton adopted with `terraform import` (the principal is not created here); once imported, `workspace_identity_status` can be toggled between ACTIVE and INACTIVE.


## Example Usage
Example usage:

This is a singleton resource: the identity detail already exists for a principal
already assigned to the workspace, so `principal_id` is read-only and the resource
is adopted with `terraform import` rather than created. Once imported, Terraform
manages `workspace_identity_status` — the only updatable field. Edit it between
`ACTIVE` and `INACTIVE` and re-apply to update in place.

```hcl
resource "databricks_workspace_iam_workspace_identity_detail_v2" "this" {
  workspace_identity_status = "ACTIVE"
}
```

Import the resource by its name, which contains the principal's internal Databricks ID:

```sh
terraform import databricks_workspace_iam_workspace_identity_detail_v2.this principals/<principal_id>/workspace-identity-detail
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the parent
* `workspace_identity_status` (string, optional) - Whether the principal is active in the workspace. Groups are always `ACTIVE`. Possible values are: `ACTIVE`, `INACTIVE`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `assignment_type` (string) - Whether the principal is assigned to the workspace directly or has access through a group. Possible values are: `DIRECT`, `INDIRECT`
* `name` (string) - The resource name of the workspace identity detail, in the format
  `principals/{principal_id}/workspace-identity-detail`
* `principal_id` (integer) - The Databricks ID of the user, service principal, or group
* `principal_type` (string) - Whether the principal is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_workspace_iam_workspace_identity_detail_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_workspace_identity_detail_v2.this "name"
```