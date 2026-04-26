---
subcategory: "Unity Catalog"
---
# databricks_secret_ucs Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `catalog_name` (string, optional) - The name of the catalog under which to list secrets. Both **catalog_name** and
  **schema_name** must be specified together
* `include_browse` (boolean, optional) - Whether to include secrets in the response for which you only have the **BROWSE** privilege,
  which limits access to metadata
* `page_size` (integer, optional) - Maximum number of secrets to return.
  
  - If not specified, at most 10000 secrets are returned.
  - If set to a value greater than 0, the page length is the minimum of this value and 10000.
  - If set to 0, the page length is set to 10000.
  - If set to a value less than 0, an invalid parameter error is returned
* `schema_name` (string, optional) - The name of the schema under which to list secrets. Both **catalog_name** and
  **schema_name** must be specified together
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `secrets`. It is a list of resources, each with the following attributes:
* `browse_only` (boolean) - Indicates whether the principal is limited to retrieving metadata for the associated object
  through the **BROWSE** privilege when **include_browse** is enabled in the request
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
* `external_secret_id` (string)
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