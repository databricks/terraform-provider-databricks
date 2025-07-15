---
subcategory: "Apps"
---
# databricks_apps Data Source

-> This data source can only be used with a workspace-level provider!

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

[Databricks Apps](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html) run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on. This resource creates the application but does not handle app deployment, which should be handled separately as part of your CI/CD pipeline.

This data source allows you to fetch information about all Databricks Apps within a workspace.

## Example Usage

```hcl
data "databricks_apps" "all_apps" {}
```

## Attribute Reference

The following attributes are exported:

* `apps` - A list of [databricks_app](../resources/app.md) resources.
  * `name` - The name of the app.
  * `description` - The description of the app.
  * `resources` - A list of resources that the app have access to.
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
  * `service_principal_id` - id of the app service principal
  * `service_principal_name` - name of the app service principal
  * `default_source_code_path` - The default workspace file system path of the source code from which app deployment are created. This field tracks the workspace source code path of the last active deployment.
  * `budget_policy_id` - The Budget Policy ID set for this resource.
  * `effective_budget_policy_id` - The effective budget policy ID.
  * `effective_user_api_scopes` - A list of effective api scopes granted to the user access token.

### resources Attribute

This attribute describes a resource used by the app.

* `name` - The name of the resource.
* `description` - The description of the resource.

Exactly one of the following attributes will be provided:

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
* `uc_securable` attribute
  * `securable_type` - the type of UC securable, i.e. `VOLUME`.
  * `securable_full_name` - the full name of UC securable, i.e. `my-catalog.my-schema.my-volume`.
  * `permission` - Permissions to grant on UC securable, i.e. `READ_VOLUME`, `WRITE_VOLUME`.
* `database` attribute
  * `database_name` - The name of database.
  * `instance_name` - The name of database instance.
  * `permission` - Permission to grant on database. Supported permissions are: `CAN_CONNECT_AND_CREATE`.

## Related Resources

The following resources are used in the same context:

* [databricks_app](../resources/app.md) to manage [Databricks Apps](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code.
