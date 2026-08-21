---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_service_principal_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads a single service principal by its internal ID, using a workspace-scoped provider.


## Example Usage
Example usage:

Look up a service principal by its internal ID, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_service_principal_v2" "this" {
  service_principal_id = "123456789"
}
```


## Arguments
The following arguments are supported:
* `service_principal_id` (string, required) - Internal service principal ID of the service principal in Databricks
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID for the service principal in Databricks
* `account_sp_status` (string) - The activity status of a service principal in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `application_id` (string) - Application ID of the service principal. Set at creation time and cannot be changed
  afterwards; when omitted, the server generates one
* `display_name` (string) - Display name of the service principal
* `external_id` (string) - ExternalId of the service principal in the customer's IdP
* `service_principal_id` (string) - Internal service principal ID of the service principal in Databricks