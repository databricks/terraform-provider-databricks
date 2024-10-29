---
subcategory: "Apps"
---
# databricks_app Resource

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

Apps run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on.

## Example Usage

```hcl
resource "databricks_app" "this" {
  name             = "my-custom-app"
  description      = "My app"
  source_code_path = "/Workspace/user@test.com/my_custom_app"
  mode             = "SNAPSHOT"
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) The name of the app. The name must contain only lowercase alphanumeric characters and hyphens. It must be unique within the workspace.
* `description` - The description of the app.

One or more `resource` block with the following arguments:

* `name` - Name of the App Resource.
* `description` - Description of the App Resource.

* One or more of the following resource blocks
  * `secret` block
      * `scope` - Scope of the secret to grant permission on.
      * `key` - Key of the secret to grant permission on.
      * `permission` - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: `READ`, `WRITE`, `MANAGE`.
  * `sql_warehouse` block
      * `id` - Id of the SQL warehouse to grant permission on.
      * `permission` - Permission to grant on the SQL warehouse. Supported permissions are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`.
  * `serving_endpoint` block
      * `name` - Name of the serving endpoint to grant permission on.
      * `permission` - Permission to grant on the serving endpoint. Supported permissions are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`.
  * `job` block
      * `id` - Id of the job to grant permission on.
      * `permission` - Permissions to grant on the Job. Supported permissions are: `CAN_MANAGE`, `IS_OWNER`, `CAN_MANAGE_RUN`, `CAN_VIEW`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `compute_status` block
  * `state` - State of the app compute.
  * `message` - Compute status message
* `app_status` block
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

```bash
terraform import databricks_app.this <app_name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
