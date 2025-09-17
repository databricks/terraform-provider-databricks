---
subcategory: "Apps"
---
# databricks_apps_settings_custom_templates Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch the list of Custom Templates within the workspace.
The list can then be accessed via the data object's `templates` field.


## Example Usage
Getting a list of all Custom Templates:

```hcl
data "databricks_apps_settings_custom_templates" "all" {}
output "all_custom_templates" {
  value = data.databricks_apps_settings_custom_templates.all.templates
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Upper bound for items returned


## Attributes
This data source exports a single attribute, `templates`. It is a list of resources, each with the following attributes:
* `creator` (string)
* `description` (string) - The description of the template
* `git_provider` (string) - The Git provider of the template
* `git_repo` (string) - The Git repository URL that the template resides in
* `manifest` (AppManifest) - The manifest of the template. It defines fields and default values when installing the template
* `name` (string) - The name of the template. It must contain only alphanumeric characters, hyphens, underscores, and whitespaces.
  It must be unique within the workspace
* `path` (string) - The path to the template within the Git repository

### AppManifest
* `description` (string) - Description of the app defined by manifest author / publisher
* `name` (string) - Name of the app defined by manifest author / publisher
* `resource_specs` (list of AppManifestAppResourceSpec)
* `version` (integer) - The manifest schema version, for now only 1 is allowed

### AppManifestAppResourceJobSpec
* `permission` (string) - Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE", "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`

### AppManifestAppResourceSecretSpec
* `permission` (string) - Permission to grant on the secret scope. For secrets, only one permission is allowed. Permission must be one of: "READ", "WRITE", "MANAGE". Possible values are: `MANAGE`, `READ`, `WRITE`

### AppManifestAppResourceServingEndpointSpec
* `permission` (string) - Permission to grant on the serving endpoint. Supported permissions are: "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW". Possible values are: `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`

### AppManifestAppResourceSpec
* `description` (string) - Description of the App Resource
* `job_spec` (AppManifestAppResourceJobSpec)
* `name` (string) - Name of the App Resource
* `secret_spec` (AppManifestAppResourceSecretSpec)
* `serving_endpoint_spec` (AppManifestAppResourceServingEndpointSpec)
* `sql_warehouse_spec` (AppManifestAppResourceSqlWarehouseSpec)
* `uc_securable_spec` (AppManifestAppResourceUcSecurableSpec)

### AppManifestAppResourceSqlWarehouseSpec
* `permission` (string) - Permission to grant on the SQL warehouse. Supported permissions are: "CAN_MANAGE", "CAN_USE", "IS_OWNER". Possible values are: `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`

### AppManifestAppResourceUcSecurableSpec
* `permission` (string) - . Possible values are: `MANAGE`, `READ_VOLUME`, `WRITE_VOLUME`
* `securable_type` (string) - . Possible values are: `VOLUME`