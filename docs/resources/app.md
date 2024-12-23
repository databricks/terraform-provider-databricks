---
subcategory: "Apps"
---
# databricks_app Resource

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

[Databricks Apps](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html) run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on. This resource creates the application but does not handle app deployment, which should be handled separately as part of your CI/CD pipeline.

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
* `resources` - (Optional) A list of resources that the app have access to.

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

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

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

## Import

This resource can be imported by name:

```hcl
import {
  to = databricks_app.this
  id = "<app_name>"
}
```

or using the `terraform` CLI:

```bash
terraform import databricks_app.this <app_name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_model_serving](model_serving.md) to serve this model on a Databricks serving endpoint.
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code.
