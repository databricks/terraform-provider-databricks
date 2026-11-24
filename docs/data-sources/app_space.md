---
subcategory: "Apps"
---
# databricks_app_space Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/apps)

This data source can be used to get a single app space.


## Example Usage
Referring to an app space by name:

```hcl
data "databricks_app_space" "this" {
  name = "my-space"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the app space. The name must contain only lowercase alphanumeric characters and hyphens.
  It must be unique within the workspace.
  It can't start or end with a hyphen, and it can't be changed after the space is created
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `assume_group_id` (string) - The group whose permissions users assume via Role Authorization for apps in this space. When
  set, user tokens assume the role of this group instead of doing regular obo token downscoping.
  Set only at space creation
* `create_time` (string) - The creation time of the app space. Formatted timestamp in ISO 8601
* `creator` (string) - The email of the user that created the app space
* `description` (string) - The description of the app space
* `effective_usage_policy_id` (string) - The effective usage policy ID used by apps in the space
* `effective_user_api_scopes` (list of string) - The effective api scopes granted to the user access token
* `id` (string) - The unique identifier of the app space
* `name` (string) - The name of the app space. The name must contain only lowercase alphanumeric characters and hyphens.
  It must be unique within the workspace.
  It can't start or end with a hyphen, and it can't be changed after the space is created
* `resources` (list of AppResource) - Resources for the app space. Resources configured at the space level are available to all apps in the space
* `service_principal_client_id` (string) - OAuth client ID of the app space's service principal
* `service_principal_id` (integer) - Numeric ID of the app space's service principal
* `service_principal_name` (string) - Display name of the app space's service principal
* `status` (SpaceStatus) - The status of the app space
* `update_time` (string) - The update time of the app space. Formatted timestamp in ISO 8601
* `updater` (string) - The email of the user that last updated the app space
* `usage_policy_id` (string) - The usage policy ID for managing cost at the space level
* `user_api_scopes` (list of string) - API scopes that apps in the space can request for user authorization. See
  https://docs.databricks.com/aws/en/dev-tools/databricks-apps/auth

### AppResource
* `app` (AppResourceApp) - Another app in the workspace that this app is granted access to
* `database` (AppResourceDatabase) - Lakebase Provisioned database the app is granted access to
* `description` (string) - Description of the App Resource
* `experiment` (AppResourceExperiment) - MLflow experiment the app is granted access to
* `genie_space` (AppResourceGenieSpace) - Genie space the app is granted access to
* `job` (AppResourceJob) - Lakeflow job the app is granted access to
* `name` (string) - Key that identifies the resource within the app. The app configuration (app.yaml) references
  the resource by this key
* `postgres` (AppResourcePostgres) - Lakebase Autoscaling Postgres database the app is granted access to
* `secret` (AppResourceSecret) - Secret the app is granted access to
* `serving_endpoint` (AppResourceServingEndpoint) - Model serving endpoint the app is granted access to
* `sql_warehouse` (AppResourceSqlWarehouse) - SQL warehouse the app is granted access to
* `uc_securable` (AppResourceUcSecurable) - Unity Catalog securable (volume, table, function, connection or model service) the app is
  granted a privilege on

### AppResourceApp
* `name` (string) - Name of the app to grant access to
* `permission` (string) - Permission to grant on that app. Possible values are: `CAN_USE`

### AppResourceDatabase
* `database_name` (string) - Name of the database within the instance
* `instance_name` (string) - Name of the Lakebase database instance
* `permission` (string) - Permission to grant the app on the database. Possible values are: `CAN_CONNECT_AND_CREATE`

### AppResourceExperiment
* `experiment_id` (string) - ID of the MLflow experiment
* `permission` (string) - Permission to grant the app on the experiment. Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`

### AppResourceGenieSpace
* `name` (string)
* `permission` (string) - Permission to grant the app on the Genie space. Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`
* `space_id` (string) - ID of the Genie space

### AppResourceJob
* `id` (string) - Id of the job to grant permission on
* `permission` (string) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppResourcePostgres
* `branch` (string) - Resource name of the Lakebase Autoscaling branch, in the form
  `projects/{project}/branches/{branch}`
* `database` (string) - Resource name of the database within the branch, in the form
  `projects/{project}/branches/{branch}/databases/{database}`
* `permission` (string) - Permission to grant the app on the database. Possible values are: `CAN_CONNECT_AND_CREATE`

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
* `permission` (string) - Privilege to grant the app on the securable. The privilege must be valid for the securable
  type, such as READ_VOLUME for a volume or SELECT for a table. Possible values are: `EXECUTE`, `MODIFY`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`
* `securable_full_name` (string) - Full name of the Unity Catalog securable. For volumes, tables and functions this is the
  three-level catalog.schema.object name
* `securable_kind` (string) - The securable kind from Unity Catalog.
  See https://docs.databricks.com/api/workspace/tables/get#securable_kind_manifest-securable_kind
* `securable_type` (string) - Kind of Unity Catalog securable that `securable_full_name` refers to. Possible values are: `CONNECTION`, `FUNCTION`, `MODEL_SERVICE`, `TABLE`, `VOLUME`

### SpaceStatus
* `message` (string) - Message providing context about the current state
* `state` (string) - The state of the app space. Possible values are: `SPACE_ACTIVE`, `SPACE_CREATING`, `SPACE_DELETED`, `SPACE_DELETING`, `SPACE_ERROR`, `SPACE_UPDATING`