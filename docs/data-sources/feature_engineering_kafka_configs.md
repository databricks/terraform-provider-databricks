---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_kafka_configs Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return


## Attributes
This data source exports a single attribute, `kafka_configs`. It is a list of resources, each with the following attributes:
* `auth_config` (AuthConfig) - Authentication configuration for connection to topics
* `bootstrap_servers` (string) - A comma-separated list of host/port pairs pointing to Kafka cluster
* `extra_options` (object) - Catch-all for miscellaneous options. Keys should be source options or Kafka consumer options (kafka.*)
* `key_schema` (SchemaConfig) - Schema configuration for extracting message keys from topics. At least one of key_schema and value_schema must be provided
* `name` (string) - Name that uniquely identifies this Kafka config within the metastore. This will be the identifier used from the Feature object to reference these configs for a feature.
  Can be distinct from topic name
* `subscription_mode` (SubscriptionMode) - Options to configure which Kafka topics to pull data from
* `value_schema` (SchemaConfig) - Schema configuration for extracting message values from topics. At least one of key_schema and value_schema must be provided

### AuthConfig
* `uc_service_credential_name` (string) - Name of the Unity Catalog service credential. This value will be set under the option databricks.serviceCredential

### SchemaConfig
* `json_schema` (string) - Schema of the JSON object in standard IETF JSON schema format (https://json-schema.org/)

### SubscriptionMode
* `assign` (string) - A JSON string that contains the specific topic-partitions to consume from.
  For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th and 1st partitions will be consumed from
* `subscribe` (string) - A comma-separated list of Kafka topics to read from. For example, 'topicA,topicB,topicC'
* `subscribe_pattern` (string) - A regular expression matching topics to subscribe to. For example, 'topic.*' will subscribe to all topics starting with 'topic'
