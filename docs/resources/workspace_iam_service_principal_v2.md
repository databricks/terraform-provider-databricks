---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_service_principal_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `account_sp_status` (string, required) - Whether the service principal is active in the account. Possible values are: `ACTIVE`, `INACTIVE`
* `display_name` (string, required) - The service principal's display name
* `external_id` (string, optional) - The ID of the service principal in your identity provider (IdP)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The ID of the account that contains the service principal
* `application_id` (string) - The service principal's application ID, a UUID. If you don't set it when you create the
  service principal, Databricks generates one. It must be unique in the account and can't be
  changed
* `name` (string) - The resource name of the service principal, in one of these formats:
  - Account APIs: `accounts/{account_id}/service-principals/{service_principal_id}`
  - Workspace APIs: `service-principals/{service_principal_id}`
* `parent` (string) - The resource name of the account that contains the service principal, in the format
  `accounts/{account_id}`. Workspace APIs return the same format
* `service_principal_id` (string) - The Databricks ID of the service principal

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_workspace_iam_service_principal_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_service_principal_v2.this "name"
```