---
subcategory: "Apps"
---
# databricks_apps_space Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `description` (string, optional) - The description of the app space
* `resources` (list of AppResource, optional) - Resources for the app space. Resources configured at the space level are available to all apps in the space
* `usage_policy_id` (string, optional) - The usage policy ID for managing cost at the space level
* `user_api_scopes` (list of string, optional) - OAuth scopes for apps in the space
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AppResource
* `name` (string, required) - Name of the App Resource
* `database` (AppResourceDatabase, optional)
* `description` (string, optional) - Description of the App Resource
* `experiment` (AppResourceExperiment, optional)
* `genie_space` (AppResourceGenieSpace, optional)
* `job` (AppResourceJob, optional)
* `secret` (AppResourceSecret, optional)
* `serving_endpoint` (AppResourceServingEndpoint, optional)
* `sql_warehouse` (AppResourceSqlWarehouse, optional)
* `uc_securable` (AppResourceUcSecurable, optional)

### AppResourceDatabase
* `database_name` (string, required)
* `instance_name` (string, required)
* `permission` (string, required) - Possible values are: `CAN_CONNECT_AND_CREATE`

### AppResourceExperiment
* `experiment_id` (string, required)
* `permission` (string, required) - Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`

### AppResourceGenieSpace
* `name` (string, required)
* `permission` (string, required) - Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`
* `space_id` (string, required)

### AppResourceJob
* `id` (string, required) - Id of the job to grant permission on
* `permission` (string, required) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppResourceSecret
* `key` (string, required) - Key of the secret to grant permission on
* `permission` (string, required) - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: "READ", "WRITE", "MANAGE". Possible values are: `MANAGE`, `READ`, `WRITE`
* `scope` (string, required) - Scope of the secret to grant permission on

### AppResourceServingEndpoint
* `name` (string, required) - Name of the serving endpoint to grant permission on
* `permission` (string, required) - Permission to grant on the serving endpoint. Supported permissions are: "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`

### AppResourceSqlWarehouse
* `id` (string, required) - Id of the SQL warehouse to grant permission on
* `permission` (string, required) - Permission to grant on the SQL warehouse. Supported permissions are: "CAN_MANAGE", "CAN_USE", "IS_OWNER". Possible values are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`

### AppResourceUcSecurable
* `permission` (string, required) - Possible values are: `EXECUTE`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`
* `securable_full_name` (string, required)
* `securable_type` (string, required) - Possible values are: `CONNECTION`, `FUNCTION`, `TABLE`, `VOLUME`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - The creation time of the app space. Formatted timestamp in ISO 6801
* `creator` (string) - The email of the user that created the app space
* `effective_usage_policy_id` (string) - The effective usage policy ID used by apps in the space
* `effective_user_api_scopes` (list of string) - The effective api scopes granted to the user access token
* `id` (string) - The unique identifier of the app space
* `name` (string) - The name of the app space. The name must contain only lowercase alphanumeric characters and hyphens.
  It must be unique within the workspace
* `oauth2_app_client_id` (string) - The OAuth2 app client ID for the app space
* `oauth2_app_integration_id` (string) - The OAuth2 app integration ID for the app space
* `service_principal_client_id` (string) - The service principal client ID for the app space
* `service_principal_id` (integer) - The service principal ID for the app space
* `service_principal_name` (string) - The service principal name for the app space
* `status` (SpaceStatus) - The status of the app space
* `update_time` (string) - The update time of the app space. Formatted timestamp in ISO 6801
* `updater` (string) - The email of the user that last updated the app space

### SpaceStatus
* `message` (string) - Message providing context about the current state
* `state` (string) - The state of the app space. Possible values are: `SPACE_ACTIVE`, `SPACE_CREATING`, `SPACE_DELETED`, `SPACE_DELETING`, `SPACE_ERROR`, `SPACE_UPDATING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_apps_space.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_apps_space.this "name"
```