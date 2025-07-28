---
subcategory: "Unity Catalog"
---
# databricks_request_for_access Resource
Request for Access (RFA) access request destinations allow you to configure where notifications are sent when users request access to securable objects in Unity Catalog. This resource enables you to manage access request destinations for specific securable objects, such as tables, catalogs, or schemas.

When a user requests access to a securable object, notifications can be sent to various destinations including email addresses, Slack channels, or Microsoft Teams channels. This resource allows you to configure these destinations to ensure that the appropriate stakeholders are notified of access requests.


## Example Usage
```hcl
resource "databricks_rfa_access_request_destinations" "customer_data_table" {
  destinations = [
    {
      destination_id = "john.doe@databricks.com"
      destination_type = "EMAIL"
    },
    {
      destination_id = "https://www.databricks.com/"
      destination_type = "URL"
    },
    {
      destination_id = "456e7890-e89b-12d3-a456-426614174001"
      destination_type = "SLACK"
    },
    {
      destination_id = "789e0123-e89b-12d3-a456-426614174002"
      destination_type = "MICROSOFT_TEAMS"
    },
    {
      destination_id = "012e3456-e89b-12d3-a456-426614174003"
      destination_type = "GENERIC_WEBHOOK"
    }
  ] 
  securable = {
    type = "SCHEMA"
    full_name = "main.customer_data"
  }
  are_any_destinations_hidden = false
}


## Arguments
The following arguments are supported:
* `destinations` (list of NotificationDestination, required) - The access request destinations for the securable
* `securable` (Securable, required) - The securable for which the access request destinations are being retrieved

### NotificationDestination
* `destination_id` (string, optional) - The identifier for the destination. This is the email address for EMAIL destinations, the URL for URL destinations,
  or the unique Databricks notification destination ID for all other external destinations
* `destination_type` (string, optional) - The type of the destination. Possible values are: `EMAIL`, `GENERIC_WEBHOOK`, `MICROSOFT_TEAMS`, `SLACK`, `URL`
* `special_destination` (string, optional) - This field is used to denote whether the destination is the email of the owner of the securable object.
  The special destination cannot be assigned to a securable and only represents the default destination of the securable.
  The securable types that support default special destinations are: "catalog", "external_location", "connection", "credential", and "metastore".
  The **destination_type** of a **special_destination** is always EMAIL. Possible values are: `SPECIAL_DESTINATION_CATALOG_OWNER`, `SPECIAL_DESTINATION_CONNECTION_OWNER`, `SPECIAL_DESTINATION_CREDENTIAL_OWNER`, `SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER`, `SPECIAL_DESTINATION_METASTORE_OWNER`

### Securable
* `full_name` (string, optional) - Required. The full name of the catalog/schema/table.
  Optional if resource_name is present
* `provider_share` (string, optional) - Optional. The name of the Share object that contains the securable when the securable is
  getting shared in D2D Delta Sharing
* `type` (string, optional) - Required. The type of securable (catalog/schema/table).
  Optional if resource_name is present. Possible values are: `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`



## Attributes
In addition to the above arguments, the following attributes are exported:
* `are_any_destinations_hidden` (boolean) - Indicates whether any destinations are hidden from the caller due to a lack of permissions.
  This value is true if the caller does not have permission to see all destinations

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = 
  to = databricks_request_for_access.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_request_for_access 
```