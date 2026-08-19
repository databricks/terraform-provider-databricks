---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_service_principal_v2 Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Manages a service principal using a workspace-scoped provider. Creating the resource provisions the service principal; deleting it removes the service principal. `application_id` is immutable — changing it replaces the resource.

When Automatic Identity Management (AIM) is enabled for the account, this resource manages local service principals only: a service principal cannot be created with `external_id` set. Service principals that already have an `external_id` can still be read through this resource, but only their `external_id` may be updated — their other fields are sourced from the identity provider.


## Example Usage
Example usage:

Creates a service principal using a workspace-scoped provider. When AIM is enabled, the service principal cannot be created with `external_id` set.

```hcl
resource "databricks_workspace_iam_service_principal_v2" "this" {
  display_name      = "ci-runner"
  account_sp_status = "ACTIVE"
}
```


## Arguments
The following arguments are supported:
* `account_sp_status` (string, required) - The activity status of a service principal in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `display_name` (string, required) - Display name of the service principal
* `external_id` (string, optional) - ExternalId of the service principal in the customer's IdP
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The parent account ID for the service principal in Databricks
* `application_id` (string) - Application ID of the service principal. Set at creation time and cannot be changed
  afterwards; when omitted, the server generates one
* `service_principal_id` (string) - Internal service principal ID of the service principal in Databricks

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "service_principal_id"
  to = databricks_workspace_iam_service_principal_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_service_principal_v2.this "service_principal_id"
```