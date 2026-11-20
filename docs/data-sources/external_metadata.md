---
subcategory: "Unity Catalog"
---
# databricks_external_metadata Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/externalmetadata)

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
* `name` (string, required) - Name of the external metadata object. The name is unique within the metastore, must follow the
  [Unity Catalog naming requirements](https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements),
  and is stored in lowercase. It cannot be changed after creation, and external lineage relationships
  refer to the object by this name
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `columns` (list of string) - Names of the columns of the object, used for column-level mappings in external lineage relationships.
  Column names must be unique within the list, and the combined size of all column names is limited
* `create_time` (string) - Time at which this external metadata object was created
* `created_by` (string) - Username of external metadata object creator
* `description` (string) - User-provided free-form text description of the object. The description must not exceed 5000 characters
* `entity_type` (string) - Type of the object within the external system, such as a table, a report, or a dashboard. The value is
  free-form text of at most 255 characters
* `id` (string) - Unique identifier of the external metadata object. The ID is a UUID that the system assigns on creation
* `metastore_id` (string) - Unique identifier of parent metastore
* `name` (string) - Name of the external metadata object. The name is unique within the metastore, must follow the
  [Unity Catalog naming requirements](https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements),
  and is stored in lowercase. It cannot be changed after creation, and external lineage relationships
  refer to the object by this name
* `owner` (string) - Owner of the external metadata object. Defaults to the creator. To change the owner, send an update whose
  update mask includes `owner`
* `properties` (object) - Key-value properties attached to the object. Keys and values must be non-empty, and the combined length
  of all keys and values must not exceed 5000 characters
* `system_type` (string) - Type of the external system that holds the object. Use `OTHER` for a system that is not in the list. Possible values are: `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `STREAM_NATIVE`, `TABLEAU`, `TERADATA`, `WORKDAY`
* `update_time` (string) - Time at which this external metadata object was last modified
* `updated_by` (string) - Username of user who last modified external metadata object
* `url` (string) - URL of the object in the external system. When set, lineage graph viewers can follow it to the external
  asset. The URL must not exceed 2048 characters