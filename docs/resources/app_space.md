---
subcategory: "Apps"
---
# databricks_app_space Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/apps)

App Spaces group related Databricks Apps and share configuration across them. Resources, user API scopes, and the usage policy configured on a space are available to every app in the space.

### Use Cases

App Spaces can be used to give a set of apps access to the same resources, such as SQL warehouses or secrets, without configuring each app individually.

### Permissions

Each app space has its own service principal. Resources configured on the space are granted to that service principal and are available to all apps that belong to the space.


## Example Usage
### Basic Example

This example creates an app space in the workspace with the specified name.

```hcl
resource "databricks_app_space" "this" {
  name        = "my-space"
  description = "My app space for team applications."
}
```

### Example with Resources

This example creates an app space that gives all of its apps access to a SQL warehouse and a secret.

```hcl
resource "databricks_app_space" "with_resources" {
  name        = "analytics-space"
  description = "Shared resources for analytics apps"
  resources = [
    {
      name = "warehouse"
      sql_warehouse = {
        id         = "e9ca293f79a74b5c"
        permission = "CAN_USE"
      }
    },
    {
      name = "api-key"
      secret = {
        scope      = "my-scope"
        key        = "my-key"
        permission = "READ"
      }
    }
  ]
  user_api_scopes = ["sql"]
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the app space. The name must contain only lowercase alphanumeric characters and hyphens.
  It must be unique within the workspace.
  It can't start or end with a hyphen, and it can't be changed after the space is created
* `assume_group_id` (string, optional) - The group whose permissions users assume via Role Authorization for apps in this space. When
  set, user tokens assume the role of this group instead of doing regular obo token downscoping.
  Set only at space creation
* `description` (string, optional) - The description of the app space
* `resources` (list of AppResource, optional) - Resources for the app space. Resources configured at the space level are available to all apps in the space
* `usage_policy_id` (string, optional) - The usage policy ID for managing cost at the space level
* `user_api_scopes` (list of string, optional) - API scopes that apps in the space can request for user authorization. See
  https://docs.databricks.com/aws/en/dev-tools/databricks-apps/auth
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AppResource
* `name` (string, required) - Key that identifies the resource within the app. The app configuration (app.yaml) references
  the resource by this key
* `app` (AppResourceApp, optional) - Another app in the workspace that this app is granted access to
* `database` (AppResourceDatabase, optional) - Lakebase Provisioned database the app is granted access to
* `description` (string, optional) - Description of the App Resource
* `experiment` (AppResourceExperiment, optional) - MLflow experiment the app is granted access to
* `genie_space` (AppResourceGenieSpace, optional) - Genie space the app is granted access to
* `job` (AppResourceJob, optional) - Lakeflow job the app is granted access to
* `postgres` (AppResourcePostgres, optional) - Lakebase Autoscaling Postgres database the app is granted access to
* `secret` (AppResourceSecret, optional) - Secret the app is granted access to
* `serving_endpoint` (AppResourceServingEndpoint, optional) - Model serving endpoint the app is granted access to
* `sql_warehouse` (AppResourceSqlWarehouse, optional) - SQL warehouse the app is granted access to
* `uc_securable` (AppResourceUcSecurable, optional) - Unity Catalog securable (volume, table, function, connection or model service) the app is
  granted a privilege on

### AppResourceApp
* `name` (string, optional) - Name of the app to grant access to
* `permission` (string, optional) - Permission to grant on that app. Possible values are: `CAN_USE`

### AppResourceDatabase
* `database_name` (string, required) - Name of the database within the instance
* `instance_name` (string, required) - Name of the Lakebase database instance
* `permission` (string, required) - Permission to grant the app on the database. Possible values are: `CAN_CONNECT_AND_CREATE`

### AppResourceExperiment
* `experiment_id` (string, required) - ID of the MLflow experiment
* `permission` (string, required) - Permission to grant the app on the experiment. Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`

### AppResourceGenieSpace
* `name` (string, required)
* `permission` (string, required) - Permission to grant the app on the Genie space. Possible values are: `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`
* `space_id` (string, required) - ID of the Genie space

### AppResourceJob
* `id` (string, required) - Id of the job to grant permission on
* `permission` (string, required) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppResourcePostgres
* `branch` (string, optional) - Resource name of the Lakebase Autoscaling branch, in the form
  `projects/{project}/branches/{branch}`
* `database` (string, optional) - Resource name of the database within the branch, in the form
  `projects/{project}/branches/{branch}/databases/{database}`
* `permission` (string, optional) - Permission to grant the app on the database. Possible values are: `CAN_CONNECT_AND_CREATE`

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
* `permission` (string, required) - Privilege to grant the app on the securable. The privilege must be valid for the securable
  type, such as READ_VOLUME for a volume or SELECT for a table. Possible values are: `EXECUTE`, `MODIFY`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`
* `securable_full_name` (string, required) - Full name of the Unity Catalog securable. For volumes, tables and functions this is the
  three-level catalog.schema.object name
* `securable_type` (string, required) - Kind of Unity Catalog securable that `securable_full_name` refers to. Possible values are: `CONNECTION`, `FUNCTION`, `MODEL_SERVICE`, `TABLE`, `VOLUME`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - The creation time of the app space. Formatted timestamp in ISO 8601
* `creator` (string) - The email of the user that created the app space
* `effective_usage_policy_id` (string) - The effective usage policy ID used by apps in the space
* `effective_user_api_scopes` (list of string) - The effective api scopes granted to the user access token
* `id` (string) - The unique identifier of the app space
* `service_principal_client_id` (string) - OAuth client ID of the app space's service principal
* `service_principal_id` (integer) - Numeric ID of the app space's service principal
* `service_principal_name` (string) - Display name of the app space's service principal
* `status` (SpaceStatus) - The status of the app space
* `update_time` (string) - The update time of the app space. Formatted timestamp in ISO 8601
* `updater` (string) - The email of the user that last updated the app space

### AppResourceUcSecurable
* `securable_kind` (string) - The securable kind from Unity Catalog.
  See https://docs.databricks.com/api/workspace/tables/get#securable_kind_manifest-securable_kind

### SpaceStatus
* `message` (string) - Message providing context about the current state
* `state` (string) - The state of the app space. Possible values are: `SPACE_ACTIVE`, `SPACE_CREATING`, `SPACE_DELETED`, `SPACE_DELETING`, `SPACE_ERROR`, `SPACE_UPDATING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_app_space.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_app_space.this "name"
```