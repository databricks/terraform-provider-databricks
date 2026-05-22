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
* `backfill_source` (BackfillSource, optional) - A user-provided and managed source for backfilling data. Historical data is used when creating a training set from streaming features linked to this Kafka config.
  In the future, a separate table will be maintained by Databricks for forward filling data.
  The schema for this source must match exactly that of the key and value schemas specified for this Kafka config
* `extra_options` (object, optional) - Catch-all for miscellaneous options. Keys should be source options or Kafka consumer options (kafka.*)
* `key_schema` (SchemaConfig, optional) - Schema configuration for extracting message keys from topics. At least one of key_schema and value_schema must be provided
* `value_schema` (SchemaConfig, optional) - Schema configuration for extracting message values from topics. At least one of key_schema and value_schema must be provided
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### AuthConfig
* `mtls_config` (MtlsConfig, optional) - Mutual-TLS authentication. See MtlsConfig
* `uc_service_credential_name` (string, optional) - Name of the Unity Catalog service credential. This value will be set under the option databricks.serviceCredential

### BackfillSource
* `delta_table_name` (string, optional) - The full three-part name (catalog, schema, name) of the Delta table containing the historical data to backfill
* `delta_table_source` (DeltaTableSource, optional, deprecated) - Deprecated: Use delta_table_name instead. Kept for backwards compatibility.
  The Delta table source containing the historical data to backfill.
  Only the delta table name is used for backfill, other fields are ignored

### DeltaTableSource
* `full_name` (string, required) - The full three-part (catalog, schema, table) name of the Delta table
* `dataframe_schema` (string, optional) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `entity_columns` (list of string, optional, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity columns of the Delta table
* `filter_condition` (string, optional) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `timeseries_column` (string, optional, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column of the Delta table
* `transformation_sql` (string, optional) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### MtlsConfig
* `key_password_ref` (SecretScopeReference, required) - Secret-scope reference for the private key password. Often the same value as the
  keystore password (keytool's default), but provided as a separate field because
  Apache Kafka requires it as a distinct option (kafka.ssl.key.password)
* `keystore_location` (string, required) - Unity Catalog volume path to the JKS keystore file containing the client certificate
  and private key. e.g. "/Volumes/<catalog>/<schema>/<volume>/client.jks". The
  materialization compute must have read permission on this volume
* `keystore_password_ref` (SecretScopeReference, required) - Secret-scope reference for the JKS keystore password
* `truststore_location` (string, required) - Unity Catalog volume path to the JKS truststore file containing the CA certificate(s)
  trusted to verify the Kafka broker's server certificate.
  e.g. "/Volumes/<catalog>/<schema>/<volume>/truststore.jks"
* `truststore_password_ref` (SecretScopeReference, required) - Secret-scope reference for the JKS truststore password
* `disable_hostname_verification` (boolean, optional) - Set to true only when the broker certificate's SAN intentionally does not match
  the connection endpoint — for example when reaching the cluster through a
  PrivateLink endpoint whose DNS name is not in the broker certificate. Skipping
  the hostname check removes a defense against man-in-the-middle attacks; do not
  enable casually. mTLS client authentication is unaffected by this option.
  
  See the Apache Kafka SSL security guide for background on this check:
  https://kafka.apache.org/42/security/encryption-and-authentication-using-ssl/#host-name-verification

### SchemaConfig
* `json_schema` (string, optional) - Schema of the JSON object in standard IETF JSON schema format (https://json-schema.org/)

### SecretScopeReference
* `key` (string, required) - The key within the scope
* `scope` (string, required) - The Databricks secret scope name

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