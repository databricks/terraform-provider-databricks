---
subcategory: "Unity Catalog"
---
# databricks_external_metadata Data Source
This data source can be used to get a single external metadata object.

-> **Note** This resource can only be used with an workspace-level provider!

## Example Usage
Referring to an external metadata object by name:

```hcl
data "databricks_external_metadata" "this" {
  name = "security_events_stream"
}
```

## Arguments
The following arguments are supported:
* `name` (string, required) - Name of the external metadata object

## Attributes
The following attributes are exported:
* `columns` (list of string) - List of columns associated with the external metadata object
* `create_time` (string) - Time at which this external metadata object was created
* `created_by` (string) - Username of external metadata object creator
* `description` (string) - User-provided free-form text description
* `entity_type` (string) - Type of entity within the external system
* `id` (string) - Unique identifier of the external metadata object
* `metastore_id` (string) - Unique identifier of parent metastore
* `name` (string) - Name of the external metadata object
* `owner` (string) - Owner of the external metadata object
* `properties` (object) - A map of key-value properties attached to the external metadata object
* `system_type` (string) - Type of external system. Possible values are: `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `TABLEAU`, `TERADATA`, `WORKDAY`
* `update_time` (string) - Time at which this external metadata object was last modified
* `updated_by` (string) - Username of user who last modified external metadata object
* `url` (string) - URL associated with the external metadata object