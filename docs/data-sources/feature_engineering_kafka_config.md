---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_kafka_config Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Name that uniquely identifies this Kafka config within the metastore. This will be the identifier used from the Feature object to reference these configs for a feature.
  Can be distinct from topic name

## Attributes
The following attributes are exported:
* `auth_config` (AuthConfig) - Authentication configuration for connection to topics
* `backfill_source` (BackfillSource) - A user-provided and managed source for backfilling data. Historical data is used when creating a training set from streaming features linked to this Kafka config.
  In the future, a separate table will be maintained by Databricks for forward filling data.
  The schema for this source must match exactly that of the key and value schemas specified for this Kafka config
* `bootstrap_servers` (string) - A comma-separated list of host/port pairs pointing to Kafka cluster
* `extra_options` (object) - Catch-all for miscellaneous options. Keys should be source options or Kafka consumer options (kafka.*)
* `key_schema` (SchemaConfig) - Schema configuration for extracting message keys from topics. At least one of key_schema and value_schema must be provided
* `name` (string) - Name that uniquely identifies this Kafka config within the metastore. This will be the identifier used from the Feature object to reference these configs for a feature.
  Can be distinct from topic name
* `subscription_mode` (SubscriptionMode) - Options to configure which Kafka topics to pull data from
* `value_schema` (SchemaConfig) - Schema configuration for extracting message values from topics. At least one of key_schema and value_schema must be provided

### AuthConfig
* `uc_service_credential_name` (string) - Name of the Unity Catalog service credential. This value will be set under the option databricks.serviceCredential

### BackfillSource
* `delta_table_source` (DeltaTableSource) - The Delta table source containing the historic data to backfill.
  Only the delta table name is used for backfill, the entity columns and timeseries column are ignored as they are defined by the associated KafkaSource

### DeltaTableSource
* `entity_columns` (list of string) - The entity columns of the Delta table
* `full_name` (string) - The full three-part (catalog, schema, table) name of the Delta table
* `timeseries_column` (string) - The timeseries column of the Delta table

### SchemaConfig
* `json_schema` (string) - Schema of the JSON object in standard IETF JSON schema format (https://json-schema.org/)

### SubscriptionMode
* `assign` (string) - A JSON string that contains the specific topic-partitions to consume from.
  For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th and 1st partitions will be consumed from
* `subscribe` (string) - A comma-separated list of Kafka topics to read from. For example, 'topicA,topicB,topicC'
* `subscribe_pattern` (string) - A regular expression matching topics to subscribe to. For example, 'topic.*' will subscribe to all topics starting with 'topic'