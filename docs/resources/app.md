---
subcategory: "Apps"
---
# databricks_app Resource

[Databricks Apps](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html) run directly on a customer's Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on. This resource creates the application but does not handle app deployment, which should be handled separately as part of your CI/CD pipeline.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_app" "this" {
  name        = "my-custom-app"
  description = "My app"

  resources = [{
    name = "sql-warehouse"
    sql_warehouse = {
      id         = "e9ca293f79a74b5c"
      permission = "CAN_MANAGE"
    }
    },
    {
      name = "serving-endpoint"
      serving_endpoint = {
        name       = "databricks-meta-llama-3-1-70b-instruct"
        permission = "CAN_MANAGE"
      }
    },
    {
      name = "job"
      job = {
        id         = "1234"
        permission = "CAN_MANAGE"
      }
  }]
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) The name of the app. The name must contain only lowercase alphanumeric characters and hyphens. It must be unique within the workspace.
* `description` - (Optional) The description of the app.
* `budget_policy_id` - (Optional) The Budget Policy ID set for this resource.
* `usage_policy_id` - (Optional) The Usage Policy ID set for this resource.
* `resources` - (Optional) A list of resources that the app have access to.
* `user_api_scopes` - (Optional) A list of api scopes granted to the user access token.  See [REST API docs](https://docs.databricks.com/api/workspace/api/scopes) for full list of supported scopes.
* `compute_size` - (Optional) A string specifying compute size for the App. Possible values are `MEDIUM`, `LARGE`.
* `git_repository` - (Optional) Git repository configuration for app deployments (see [below](#git_repository-configuration-attribute)). When specified, deployments can reference code from this repository by providing only the git reference (branch, tag, or commit).
* `telemetry_export_destinations` - (Optional) A list of destinations to which the app's telemetry (logs, metrics, traces) is exported (see [below](#telemetry_export_destinations-configuration-attribute)).

### resources Configuration Attribute

This attribute describes a resource used by the app.

* `name` - (Required) The name of the resource.
* `description` - (Optional) The description of the resource.

Exactly one of the following attributes must be provided:

* `secret` attribute
  * `scope` - Scope of the secret to grant permission on.
  * `key` - Key of the secret to grant permission on.
  * `permission` - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: `READ`, `WRITE`, `MANAGE`.
* `sql_warehouse` attribute
  * `id` - Id of the SQL warehouse to grant permission on.
  * `permission` - Permission to grant on the SQL warehouse. Supported permissions are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`.
* `serving_endpoint` attribute
  * `name` - Name of the serving endpoint to grant permission on.
  * `permission` - Permission to grant on the serving endpoint. Supported permissions are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`.
* `job` attribute
  * `id` - Id of the job to grant permission on.
  * `permission` - Permissions to grant on the Job. Supported permissions are: `CAN_MANAGE`, `IS_OWNER`, `CAN_MANAGE_RUN`, `CAN_VIEW`.
* `uc_securable` attribute (see the [API docs](https://docs.databricks.com/api/workspace/apps/create#resources-uc_securable) for full list of supported UC objects)
  * `securable_type` - The type of UC securable. Supported values are `CONNECTION`, `FUNCTION`, `TABLE`, `VOLUME`.
  * `securable_full_name` - The full name of UC securable, i.e. `my-catalog.my-schema.my-volume`.
  * `permission` - Permission to grant on UC securable. Supported values depend on `securable_type`: `READ_VOLUME` and `WRITE_VOLUME` for `VOLUME`, `SELECT` and `MODIFY` for `TABLE`, `EXECUTE` for `FUNCTION`, `USE_CONNECTION` for `CONNECTION`.
* `database` attribute
  * `database_name` - The name of database.
  * `instance_name` - The name of database instance.
  * `permission` - Permission to grant on database. Supported permissions are: `CAN_CONNECT_AND_CREATE`.
* `postgres` attribute
  * `branch` - The resource path of the Lakebase Autoscaling branch to grant permission on (e.g. `projects/proj-abc123/branches/branch-xyz789`).
  * `database` - (Optional) The resource path of a specific database within the branch to grant permission on (e.g. `projects/proj-abc123/branches/branch-xyz789/databases/db-456`). If omitted, permission applies to the branch.
  * `permission` - Permission to grant on the Lakebase Autoscaling branch or database. Supported permissions are: `CAN_CONNECT_AND_CREATE`.
* `genie_space` attribute
  * `name` - The name of Genie Space.
  * `permission` - Permission to grant on Genie Space. Supported permissions are `CAN_MANAGE`, `CAN_EDIT`, `CAN_RUN`, `CAN_VIEW`.
  * `space_id` - The unique ID of Genie Space.
* `app` attribute - reference to another Databricks App.
  * `name` - The name of the app to grant permission on.
  * `permission` - Permission to grant on the app. Supported permissions are: `CAN_USE`.
* `experiment` attribute
  * `experiment_id` - The ID of the MLflow experiment to grant permission on.
  * `permission` - Permission to grant on the experiment. Supported permissions are: `CAN_READ`, `CAN_EDIT`, `CAN_MANAGE`.

### git_repository Configuration Attribute

This block configures a Git repository associated with the app, allowing deployments to reference code from this repository:

* `url` - (Required) URL of the Git repository.
* `provider` - (Required) Git provider. Case insensitive. Supported values: `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`, `awsCodeCommit`.

### telemetry_export_destinations Configuration Attribute

Each item describes a single destination to which the app's telemetry (logs, metrics, traces) is exported. Exactly one of the following nested blocks must be provided per destination:

* `unity_catalog` attribute - export telemetry to Unity Catalog tables (must already exist and be writable by the app's service principal).
  * `logs_table` - (Required) Full name of the Unity Catalog table for OpenTelemetry logs.
  * `metrics_table` - (Required) Full name of the Unity Catalog table for OpenTelemetry metrics.
  * `traces_table` - (Required) Full name of the Unity Catalog table for OpenTelemetry traces (spans).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the app.
* `compute_status` attribute
  * `state` - State of the app compute.
  * `message` - Compute status message
* `app_status` attribute
  * `state` - State of the application.
  * `message` - Application status message
* `url` - The URL of the app once it is deployed.
* `create_time` - The creation time of the app.
* `creator` - The email of the user that created the app.
* `update_time` - The update time of the app.
* `updater` - The email of the user that last updated the app.
* `service_principal_client_id` - client_id (application_id) of the app service principal
* `service_principal_id` - id of the app service principal
* `service_principal_name` - name of the app service principal
* `default_source_code_path` - The default workspace file system path of the source code from which app deployment are created. This field tracks the workspace source code path of the last active deployment.
* `effective_budget_policy_id` - The effective budget policy ID.
* `effective_usage_policy_id` - The effective usage policy ID.
* `effective_user_api_scopes` - A list of effective api scopes granted to the user access token.
* `oauth2_app_client_id` - The OAuth2 client ID of the app's integration, set when the app uses user authorization.
* `oauth2_app_integration_id` - The unique ID of the OAuth2 integration associated with the app.
* `thumbnail_url` - The URL of the thumbnail image for the app.
* `active_deployment` attribute - the active deployment of the app. A deployment is considered active when it has been deployed to the app compute.
  * `deployment_id` - The unique ID of the deployment.
  * `source_code_path` - The workspace file system path of the source code used to create the deployment.
  * `mode` - The deployment mode (`AUTO_SYNC` or `SNAPSHOT`).
  * `create_time` - The creation time of the deployment.
  * `creator` - The email of the user that created the deployment.
  * `update_time` - The update time of the deployment.
  * `status` attribute
    * `state` - The state of the deployment.
    * `message` - The status message of the deployment.
  * `deployment_artifacts` attribute
    * `source_code_path` - The snapshotted workspace file system path of the source code loaded by the deployed app.
* `pending_deployment` attribute - the pending deployment of the app. A deployment is considered pending when it is being prepared for deployment to the app compute. Schema is identical to `active_deployment`.

## Import

This resource can be imported by name:

```hcl
import {
  to = databricks_app.this
  id = "<app_name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_app.this <app_name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code.
