---
subcategory: "Unity Catalog"
---
# databricks_external_metadata Resource
To enrich lineage with workloads that are run outside of Databricks (for example, first mile ETL or last mile BI),
Unity Catalog is introducing the external metadata object. UC lets you add external metadata objects to augment the data lineage it captures automatically, giving you an end-to-end lineage view in UC. 
This is useful when you want to capture where data came from (for example, Salesforce or MySQL) before it was ingested into UC or where data is being consumed outside UC (for example, Tableau or PowerBI).

-> **Note** This resource can only be used with an workspace-level provider!

## Example Usage
```hcl
resource "databricks_external_metadata" "this" {
  name = "security_events_stream"
  system_type = "KAFKA"
  entity_type = "Topic"
  url = "https://kafka.com/12345"
  description = "A stream of security related events in the critical services."
  columns = [
    "type",
    "message",
    "details",
    "date",
    "time"
  ]
  properties = {
    topic: "prod.security.events.raw",
    compression.enabled: "true",
    compression.format: "zstd"
  }
}
```

## Arguments
The following arguments are supported:
* `entity_type` (string, required) - Type of entity within the external system
* `name` (string, required) - Name of the external metadata object
* `system_type` (string, required) - Type of external system. Possible values are: `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `TABLEAU`, `TERADATA`, `WORKDAY`
* `columns` (list of string, optional) - List of columns associated with the external metadata object
* `description` (string, optional) - User-provided free-form text description
* `owner` (string, optional) - Owner of the external metadata object
* `properties` (object, optional) - A map of key-value properties attached to the external metadata object
* `url` (string, optional) - URL associated with the external metadata object
* `workspace_id` (string, optional) - Workspace ID of the resource

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time at which this external metadata object was created
* `created_by` (string) - Username of external metadata object creator
* `id` (string) - Unique identifier of the external metadata object
* `metastore_id` (string) - Unique identifier of parent metastore
* `update_time` (string) - Time at which this external metadata object was last modified
* `updated_by` (string) - Username of user who last modified external metadata object

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_external_metadata.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_external_metadata "name"
```