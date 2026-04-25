---
subcategory: "Unity Catalog"
---
# databricks_secret_uc Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

The Secret resource allows you to manage secrets in Unity Catalog. Secrets provide a secure way to store and access sensitive information such as credentials, API keys, and tokens within Unity Catalog.

Secrets are scoped to a schema and identified by their three-level fully qualified name (`catalog_name.schema_name.secret_name`).

### Permissions
- The calling principal must have the appropriate privileges to create, read, update, or delete secrets in the target schema.


## Example Usage
### Basic Example
This example creates a secret in Unity Catalog:

```hcl
resource "databricks_secret_uc" "example" {
  catalog_name = "my_catalog"
  schema_name  = "my_schema"
  secret {
    name    = "my_secret"
    value   = "secret_value"
    comment = "My secret for external service authentication"
  }
}
```


## Arguments
The following arguments are supported:
* `catalog_name` (string, required) - The name of the catalog where the schema and the secret reside
* `name` (string, required) - The name of the secret, relative to its parent schema
* `schema_name` (string, required) - The name of the schema where the secret resides
* `value` (string, required) - The secret value to store. This field is input-only and is not returned in responses — use
  the **effective_value** field (via GetSecret with **include_value** set to true) to read the
  secret value. The maximum size is 60 KiB (pre-encryption). Accepted content includes
  passwords, tokens, keys, and other sensitive credential data
* `comment` (string, optional) - User-provided free-form text description of the secret
* `expire_time` (string, optional) - User-provided expiration time of the secret. This field indicates when the secret should no
  longer be used and may be displayed as a warning in the UI. It is purely informational and
  does not trigger any automatic actions or affect the secret's lifecycle
* `owner` (string, optional) - The owner of the secret. Defaults to the creating principal on creation. Can be updated to
  transfer ownership of the secret to another principal
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `browse_only` (boolean) - Indicates whether the principal is limited to retrieving metadata for the associated object
  through the **BROWSE** privilege when **include_browse** is enabled in the request
* `create_time` (string) - The time at which this secret was created
* `created_by` (string) - The principal that created the secret
* `effective_owner` (string) - The effective owner of the secret, which may differ from the directly-set **owner** due to
  inheritance
* `effective_value` (string) - The secret value. Only populated in responses when you have the **READ_SECRET**
  privilege and **include_value** is set to true in the request. The maximum size is 60 KiB
* `external_secret_id` (string)
* `full_name` (string) - The three-level (fully qualified) name of the secret, in the form of **catalog_name.schema_name.secret_name**
* `metastore_id` (string) - Unique identifier of the metastore hosting the secret
* `update_time` (string) - The time at which this secret was last updated
* `updated_by` (string) - The principal that last updated the secret

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "full_name"
  to = databricks_secret_uc.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_secret_uc.this "full_name"
```