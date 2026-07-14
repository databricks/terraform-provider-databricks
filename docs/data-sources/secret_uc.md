---
subcategory: "Unity Catalog"
---
# databricks_secret_uc Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/secretsuc)

The Secret data source allows you to read a single secret in Unity Catalog by its three-level fully qualified name (`catalog_name.schema_name.secret_name`).

This returns the secret's metadata. The secret value is only returned to principals with the `READ_SECRET` privilege.

### Permissions
- The calling principal must have the appropriate privileges to read the secret in the target schema.


## Example Usage
### Basic Example
This example reads a secret in Unity Catalog by its fully qualified name:

```hcl
data "databricks_secret_uc" "example" {
  full_name = "my_catalog.my_schema.my_secret"
}
```


## Arguments
The following arguments are supported:
* `full_name` (string, required) - The three-level (fully qualified) name of the secret, in the form of **catalog_name.schema_name.secret_name**
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `catalog_name` (string) - The name of the catalog where the schema and the secret reside
* `comment` (string) - User-provided free-form text description of the secret
* `create_time` (string) - The time at which this secret was created
* `created_by` (string) - The principal that created the secret
* `effective_owner` (string) - The effective owner of the secret, which may differ from the directly-set **owner** due to
  inheritance
* `effective_value` (string) - The secret value. Only populated in responses when you have the **READ_SECRET**
  privilege and **include_value** is set to true in the request. The maximum size is 60 KiB
* `expire_time` (string) - User-provided expiration time of the secret. This field indicates when the secret should no
  longer be used and may be displayed as a warning in the UI. It is purely informational and
  does not trigger any automatic actions or affect the secret's lifecycle
* `full_name` (string) - The three-level (fully qualified) name of the secret, in the form of **catalog_name.schema_name.secret_name**
* `metastore_id` (string) - Unique identifier of the metastore hosting the secret
* `name` (string) - The name of the secret, relative to its parent schema
* `owner` (string) - The owner of the secret. Defaults to the creating principal on creation. Can be updated to
  transfer ownership of the secret to another principal
* `schema_name` (string) - The name of the schema where the secret resides
* `update_time` (string) - The time at which this secret was last updated
* `updated_by` (string) - The principal that last updated the secret
* `value` (string) - The secret value to store. This field is input-only and is not returned in responses — use
  the **effective_value** field (via GetSecret with **include_value** set to true) to read the
  secret value. The maximum size is 60 KiB (pre-encryption). Accepted content includes
  passwords, tokens, keys, and other sensitive credential data