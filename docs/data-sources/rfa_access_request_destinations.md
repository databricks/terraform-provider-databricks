---
subcategory: "Unity Catalog"
---
# databricks_rfa_access_request_destinations Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/rfa)

This data source can be used to get the Request for Access (RFA) access request destinations for a specific securable object.


## Example Usage
Referring to RFA access request destinations by securable type and full name:

```hcl
data "databricks_rfa_access_request_destinations" "customer_data_schema" {
  securable_type = "SCHEMA"
  full_name = "main.customer_data"
}
```


## Arguments
The following arguments are supported:
* `full_name` (string, required) - The full name of the securable, repeated from **securable**
* `securable_type` (string, required) - The type of the securable, repeated from **securable** as a plain string
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `are_any_destinations_hidden` (boolean) - Indicates whether any destinations are hidden from the caller due to a lack of permissions.
  This value is true if the caller does not have permission to see all destinations
* `destination_source_securable` (Securable) - The source securable from which the destinations are inherited. Either the same value as securable (if destination
  is set directly on the securable) or the nearest parent securable with destinations set
* `destinations` (list of NotificationDestination) - The access request destinations for the securable
* `full_name` (string) - The full name of the securable, repeated from **securable**
* `securable` (Securable) - The securable for which the access request destinations are being modified or read
* `securable_type` (string) - The type of the securable, repeated from **securable** as a plain string
* `url_configuration` (UrlConfiguration) - The access request URLs shared by every workspace in the metastore. Only set when the securable is a metastore.
  These URLs are used when neither the requested securable nor its parents have an explicit destination or a
  default owner email destination, and they take precedence over metastore destinations.
  
  A **dynamic_url** and default owner email destinations cannot both be on: saving one turns the other off.
  Saving a **dynamic_url** while they are on requires a metastore admin.
  
  Setting this field removes the metastore destinations in this workspace. The first save also removes the
  metastore URL destinations of other workspaces. Once set, the configuration can be cleared but not removed

### NotificationDestination
* `destination_id` (string) - The identifier for the destination. This is the email address for EMAIL destinations, the URL for URL destinations,
  or the unique Databricks notification destination ID for all other external destinations
* `destination_type` (string) - The type of the destination. Possible values are: `EMAIL`, `GENERIC_WEBHOOK`, `MICROSOFT_TEAMS`, `SLACK`, `URL`
* `special_destination` (string) - This field is used to denote whether the destination is the email of the owner of the securable object.
  The special destination cannot be assigned to a securable and only represents the default destination of the securable.
  The securable types that support default special destinations are: "catalog", "external_location", "connection", "credential", and "metastore".
  The **destination_type** of a **special_destination** is always EMAIL. Possible values are: `SPECIAL_DESTINATION_CATALOG_OWNER`, `SPECIAL_DESTINATION_CONNECTION_OWNER`, `SPECIAL_DESTINATION_CREDENTIAL_OWNER`, `SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER`, `SPECIAL_DESTINATION_FUNCTION_OWNER`, `SPECIAL_DESTINATION_METASTORE_OWNER`, `SPECIAL_DESTINATION_REGISTERED_MODEL_OWNER`, `SPECIAL_DESTINATION_SCHEMA_OWNER`, `SPECIAL_DESTINATION_TABLE_OWNER`, `SPECIAL_DESTINATION_VOLUME_OWNER`

### Securable
* `full_name` (string) - Required. The full name of the catalog/schema/table.
  Optional if resource_name is present
* `provider_share` (string) - Optional. The name of the Share object that contains the securable when the securable is
  getting shared in D2D Delta Sharing
* `type` (string) - Required. The type of securable (catalog/schema/table).
  Optional if resource_name is present. Possible values are: `AGENT_SERVICE`, `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FEATURE`, `FUNCTION`, `MCP_SERVICE`, `METASTORE`, `MODEL`, `MODEL_PROVIDER_SERVICE`, `MODEL_SERVICE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `SKILL`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`

### UrlConfiguration
* `dynamic_url` (string) - The URL template for access requests on catalogs and the securables inside them, such as schemas, tables,
  volumes, functions, and registered models. Placeholders are replaced with percent-encoded values from the
  requested securable: `{catalog}` with the catalog name (`main`), `{catalog.schema}` with the catalog and
  schema names (`main.sales`), `{catalog.schema.object}` with the full name (`main.sales.orders`),
  `{securable_type}` with the securable type (`TABLE`), and `{securable_id}` with the securable ID.
  Name placeholders leave out levels the securable does not have, so `{catalog.schema.object}` is `main.sales`
  for a schema. When the deepest name placeholder stops above the requested securable, such as
  `{catalog.schema}` for a table, `{securable_type}` and `{securable_id}` describe that parent.
  If the filled-in URL is longer than 4,096 characters, the **redirect_url** is used instead.
  
  Must start with `http://` or `https://`, contain only printable ASCII characters without spaces or backslashes,
  and be at most 2,048 characters. Braces are only allowed around a placeholder
* `redirect_url` (string) - The URL for access requests on the metastore and on securables outside catalogs, such as external locations,
  connections, and credentials. Also used for catalogs and the securables inside them when **dynamic_url** is
  not set or cannot produce a valid URL.
  Must start with `http://` or `https://`, contain only printable ASCII characters without spaces or backslashes,
  and be at most 2,048 characters
* `syntax_version` (integer) - The version of the placeholder syntax used in **dynamic_url**. The only supported version is `1`