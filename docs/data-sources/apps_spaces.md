---
subcategory: "Apps"
---
# databricks_apps_spaces Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `spaces`. It is a list of resources, each with the following attributes:
* `create_time` (string) - The creation time of the app space. Formatted timestamp in ISO 6801
* `creator` (string) - The email of the user that created the app space
* `description` (string) - The description of the app space
* `effective_usage_policy_id` (string) - The effective usage policy ID used by apps in the space
* `effective_user_api_scopes` (list of string) - The effective api scopes granted to the user access token
* `id` (string) - The unique identifier of the app space
* `name` (string) - The name of the app space. The name must contain only lowercase alphanumeric characters and hyphens.
  It must be unique within the workspace
* `oauth2_app_client_id` (string) - The OAuth2 app client ID for the app space
* `oauth2_app_integration_id` (string) - The OAuth2 app integration ID for the app space
* `resources` (list of AppResource) - Resources for the app space. Resources configured at the space level are available to all apps in the space
* `service_principal_client_id` (string) - The service principal client ID for the app space
* `service_principal_id` (integer) - The service principal ID for the app space
* `service_principal_name` (string) - The service principal name for the app space
* `status` (SpaceStatus) - The status of the app space
* `update_time` (string) - The update time of the app space. Formatted timestamp in ISO 6801
* `updater` (string) - The email of the user that last updated the app space
* `usage_policy_id` (string) - The usage policy ID for managing cost at the space level
* `user_api_scopes` (list of string) - OAuth scopes for apps in the space

### AppResource
* `database` (AppResourceDatabase)
* `description` (string) - Description of the App Resource
* `experiment` (AppResourceExperiment)
* `genie_space` (AppResourceGenieSpace)
* `job` (AppResourceJob)
* `name` (string) - Name of the App Resource
* `secret` (AppResourceSecret)
* `serving_endpoint` (AppResourceServingEndpoint)
* `sql_warehouse` (AppResourceSqlWarehouse)
* `uc_securable` (AppResourceUcSecurable)

### AppResourceDatabase
* `database_name` (string)
* `instance_name` (string)
* `permission` (string) - Possible values are: `CAN_CONNECT_AND_CREATE`

### AppResourceExperiment
* `experiment_id` (string)
* `permission` (string) - Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`

### AppResourceGenieSpace
* `name` (string)
* `permission` (string) - Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`
* `space_id` (string)

### AppResourceJob
* `id` (string) - Id of the job to grant permission on
* `permission` (string) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppResourceSecret
* `key` (string) - Key of the secret to grant permission on
* `permission` (string) - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: "READ", "WRITE", "MANAGE". Possible values are: `MANAGE`, `READ`, `WRITE`
* `scope` (string) - Scope of the secret to grant permission on

### AppResourceServingEndpoint
* `name` (string) - Name of the serving endpoint to grant permission on
* `permission` (string) - Permission to grant on the serving endpoint. Supported permissions are: "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`

### AppResourceSqlWarehouse
* `id` (string) - Id of the SQL warehouse to grant permission on
* `permission` (string) - Permission to grant on the SQL warehouse. Supported permissions are: "CAN_MANAGE", "CAN_USE", "IS_OWNER". Possible values are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`

### AppResourceUcSecurable
* `permission` (string) - Possible values are: `EXECUTE`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`
* `securable_full_name` (string)
* `securable_type` (string) - Possible values are: `CONNECTION`, `FUNCTION`, `TABLE`, `VOLUME`

### SpaceStatus
* `message` (string) - Message providing context about the current state
* `state` (string) - The state of the app space. Possible values are: `SPACE_ACTIVE`, `SPACE_CREATING`, `SPACE_DELETED`, `SPACE_DELETING`, `SPACE_ERROR`, `SPACE_UPDATING`