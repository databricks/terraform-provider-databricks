---
subcategory: "Apps"
---
# databricks_apps_settings_custom_template Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Custom App Templates store the metadata of custom app code hosted in an external Git repository, enabling users to reuse boilerplate code when creating apps.

### Use Cases

Custom Templates can be used for standardizing app code and establishing best practices when creating apps.

### Permissions

Custom Templates are objects created at the workspace level. As such, management operations are performed by workspace admins.

The workspace admin does not need permissions on the Git repository to add Custom Templates into the workspace. However, the user installing from the custom template needs Read (or higher) permissions to the repository.

**Note:** Each workspace is limited to 20 Custom Templates in total.


## Example Usage
### Basic Example

This example creates a Custom Template in the workspace with the specified name.

```hcl
resource "databricks_apps_settings_custom_template" "this" {
  name         = "my-custom-template"
  description  = "A sample custom app template"
  git_repo     = "https://github.com/example/repo.git"
  path         = "path-to-template"
  git_provider = "github"
  manifest     = {
    version    = 1
    name       = "my-custom-app"
  }
}
```

### Example with API Scopes

This example creates a custom template that declares required user API scopes.

```hcl
resource "databricks_apps_settings_custom_template" "api_scopes_example" {
  name         = "my-api-template"
  description  = "A template that requests user API scopes"
  git_repo     = "https://github.com/example/my-app.git"
  path         = "templates/app"
  git_provider = "github"

  manifest = {
    version     = 1
    name        = "my-databricks-app"
    description = "This app requires the SQL API scope."
    user_api_scopes = [
      "sql"
    ]
  }
}
```

### Example with Resource Requirements

This example defines a template that requests specific workspace resources with permissions granted.

```hcl
resource "databricks_apps_settings_custom_template" "resources_example" {
  name         = "my-resource-template"
  description  = "Template that requires secret and SQL warehouse access"
  git_repo     = "https://github.com/example/resource-app.git"
  path         = "resource-template"
  git_provider = "github"
  manifest = {
    version     = 1
    name        = "resource-consuming-app"
    description = "This app requires access to a secret and SQL warehouse."
    resource_specs = [
      {
        name        = "my-secret"
        description = "A secret needed by the app"
        secret_spec = {
          permission = "READ"
        }
      },
      {
        name        = "warehouse"
        description = "Warehouse access"
        sql_warehouse_spec = {
          permission = "CAN_USE"
        }
      }
    ]
  }
}
```


## Arguments
The following arguments are supported:
* `git_provider` (string, required) - The Git provider of the template
* `git_repo` (string, required) - The Git repository URL that the template resides in
* `manifest` (AppManifest, required) - The manifest of the template. It defines fields and default values when installing the template
* `name` (string, required) - The name of the template. It must contain only alphanumeric characters, hyphens, underscores, and whitespaces.
  It must be unique within the workspace
* `path` (string, required) - The path to the template within the Git repository
* `description` (string, optional) - The description of the template

### AppManifest
* `name` (string, required) - Name of the app defined by manifest author / publisher
* `version` (integer, required) - The manifest schema version, for now only 1 is allowed
* `description` (string, optional) - Description of the app defined by manifest author / publisher
* `resource_specs` (list of AppManifestAppResourceSpec, optional)

### AppManifestAppResourceJobSpec
* `permission` (string, required) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppManifestAppResourceSecretSpec
* `permission` (string, required) - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: "READ", "WRITE", "MANAGE". Possible values are: `MANAGE`, `READ`, `WRITE`

### AppManifestAppResourceServingEndpointSpec
* `permission` (string, required) - Permission to grant on the serving endpoint. Supported permissions are: "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`

### AppManifestAppResourceSpec
* `name` (string, required) - Name of the App Resource
* `description` (string, optional) - Description of the App Resource
* `job_spec` (AppManifestAppResourceJobSpec, optional)
* `secret_spec` (AppManifestAppResourceSecretSpec, optional)
* `serving_endpoint_spec` (AppManifestAppResourceServingEndpointSpec, optional)
* `sql_warehouse_spec` (AppManifestAppResourceSqlWarehouseSpec, optional)
* `uc_securable_spec` (AppManifestAppResourceUcSecurableSpec, optional)

### AppManifestAppResourceSqlWarehouseSpec
* `permission` (string, required) - Permission to grant on the SQL warehouse. Supported permissions are: "CAN_MANAGE", "CAN_USE", "IS_OWNER". Possible values are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`

### AppManifestAppResourceUcSecurableSpec
* `permission` (string, required) - Possible values are: `MANAGE`, `READ_VOLUME`, `WRITE_VOLUME`
* `securable_type` (string, required) - Possible values are: `VOLUME`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `creator` (string)

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_apps_settings_custom_template.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_apps_settings_custom_template.this "name"
```