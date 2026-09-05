---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_workspace_identity_detail_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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

Import the resource by the principal's internal Databricks ID:

```sh
terraform import databricks_workspace_iam_workspace_identity_detail_v2.this <principal_id>
```


## Arguments
The following arguments are supported:
* `workspace_identity_status` (string, optional) - The activity status of an identity in a Databricks workspace. Possible values are: `ACTIVE`, `INACTIVE`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `assignment_type` (string) - The type of assignment the principal has to the workspace (direct or indirect). Possible values are: `DIRECT`, `INDIRECT`
* `principal_id` (integer) - The internal ID of the principal (user/sp/group) in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "principal_id"
  to = databricks_workspace_iam_workspace_identity_detail_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_workspace_identity_detail_v2.this "principal_id"
```