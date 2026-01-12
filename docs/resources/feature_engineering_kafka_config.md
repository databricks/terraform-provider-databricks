---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_kafka_config Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `auth_config` (AuthConfig, required) - Authentication configuration for connection to topics
* `bootstrap_servers` (string, required) - A comma-separated list of host/port pairs pointing to Kafka cluster
* `subscription_mode` (SubscriptionMode, required) - Options to configure which Kafka topics to pull data from
* `extra_options` (object, optional) - Catch-all for miscellaneous options. Keys should be source options or Kafka consumer options (kafka.*)
* `key_schema` (SchemaConfig, optional) - Schema configuration for extracting message keys from topics. At least one of key_schema and value_schema must be provided
* `value_schema` (SchemaConfig, optional) - Schema configuration for extracting message values from topics. At least one of key_schema and value_schema must be provided

### AuthConfig
* `uc_service_credential_name` (string, optional) - Name of the Unity Catalog service credential. This value will be set under the option databricks.serviceCredential

### SchemaConfig
* `json_schema` (string, optional) - Schema of the JSON object in standard IETF JSON schema format (https://json-schema.org/)

### SubscriptionMode
* `assign` (string, optional) - A JSON string that contains the specific topic-partitions to consume from.
  For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th and 1st partitions will be consumed from
* `subscribe` (string, optional) - A comma-separated list of Kafka topics to read from. For example, 'topicA,topicB,topicC'
* `subscribe_pattern` (string, optional) - A regular expression matching topics to subscribe to. For example, 'topic.*' will subscribe to all topics starting with 'topic'

## Attributes
In addition to the above arguments, the following attributes are exported:
* `name` (string) - Name that uniquely identifies this Kafka config within the metastore. This will be the identifier used from the Feature object to reference these configs for a feature.
  Can be distinct from topic name

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_feature_engineering_kafka_config.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_feature_engineering_kafka_config.this "name"
```
