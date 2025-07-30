---
subcategory: "Unity Catalog"
---
# databricks_request_for_access Data Source
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

## Attributes
The following attributes are exported:
* `are_any_destinations_hidden` (boolean) - Indicates whether any destinations are hidden from the caller due to a lack of permissions.
  This value is true if the caller does not have permission to see all destinations
* `destinations` (list of NotificationDestination) - The access request destinations for the securable
* `securable` (Securable) - The securable for which the access request destinations are being retrieved

### NotificationDestination
* `destination_id` (string) - The identifier for the destination. This is the email address for EMAIL destinations, the URL for URL destinations,
  or the unique Databricks notification destination ID for all other external destinations
* `destination_type` (string) - The type of the destination. Possible values are: `EMAIL`, `GENERIC_WEBHOOK`, `MICROSOFT_TEAMS`, `SLACK`, `URL`
* `special_destination` (string) - This field is used to denote whether the destination is the email of the owner of the securable object.
  The special destination cannot be assigned to a securable and only represents the default destination of the securable.
  The securable types that support default special destinations are: "catalog", "external_location", "connection", "credential", and "metastore".
  The **destination_type** of a **special_destination** is always EMAIL. Possible values are: `SPECIAL_DESTINATION_CATALOG_OWNER`, `SPECIAL_DESTINATION_CONNECTION_OWNER`, `SPECIAL_DESTINATION_CREDENTIAL_OWNER`, `SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER`, `SPECIAL_DESTINATION_METASTORE_OWNER`

### Securable
* `full_name` (string) - Required. The full name of the catalog/schema/table.
  Optional if resource_name is present
* `provider_share` (string) - Optional. The name of the Share object that contains the securable when the securable is
  getting shared in D2D Delta Sharing
* `type` (string) - Required. The type of securable (catalog/schema/table).
  Optional if resource_name is present. Possible values are: `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`