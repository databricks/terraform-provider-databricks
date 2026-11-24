---
subcategory: "Unity Catalog"
---
# databricks_external_metadata Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/externalmetadata)

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
* `entity_type` (string, required) - Type of the object within the external system, such as a table, a report, or a dashboard. The value is
  free-form text of at most 255 characters
* `name` (string, required) - Name of the external metadata object. The name is unique within the metastore, must follow the
  [Unity Catalog naming requirements](https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements),
  and is stored in lowercase. It cannot be changed after creation, and external lineage relationships
  refer to the object by this name
* `system_type` (string, required) - Type of the external system that holds the object. Use `OTHER` for a system that is not in the list. Possible values are: `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `STREAM_NATIVE`, `TABLEAU`, `TERADATA`, `WORKDAY`
* `columns` (list of string, optional) - Names of the columns of the object, used for column-level mappings in external lineage relationships.
  Column names must be unique within the list, and the combined size of all column names is limited
* `description` (string, optional) - User-provided free-form text description of the object. The description must not exceed 5000 characters
* `owner` (string, optional) - Owner of the external metadata object. Defaults to the creator. To change the owner, send an update whose
  update mask includes `owner`
* `properties` (object, optional) - Key-value properties attached to the object. Keys and values must be non-empty, and the combined length
  of all keys and values must not exceed 5000 characters
* `url` (string, optional) - URL of the object in the external system. When set, lineage graph viewers can follow it to the external
  asset. The URL must not exceed 2048 characters
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time at which this external metadata object was created
* `created_by` (string) - Username of external metadata object creator
* `id` (string) - Unique identifier of the external metadata object. The ID is a UUID that the system assigns on creation
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
terraform import databricks_external_metadata.this "name"
```