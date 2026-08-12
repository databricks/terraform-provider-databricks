---
subcategory: "Machine Learning"
---
# databricks_feature_engineering_kafka_configs Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of results to return
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `kafka_configs`. It is a list of resources, each with the following attributes:
* `auth_config` (AuthConfig) - Authentication configuration for connection to topics
* `backfill_source` (BackfillSource) - A user-provided and managed source for backfilling data. Historical data is used when creating a training set from streaming features linked to this Kafka config.
  In the future, a separate table will be maintained by Databricks for forward filling data.
  The schema for this source must match exactly that of the key and value schemas specified for this Kafka config
* `bootstrap_servers` (string) - A comma-separated list of host/port pairs pointing to Kafka cluster
* `extra_options` (object) - Catch-all for miscellaneous options. Keys should be source options or Kafka consumer options (kafka.*)
* `ingestion_config` (IngestionConfig) - Configuration for ingesting Kafka data into a Databricks-managed
  Delta table
* `key_schema` (SchemaConfig) - Schema configuration for extracting message keys from topics. At least one of key_schema and value_schema must be provided
* `name` (string) - Name that uniquely identifies this Kafka config within the metastore. This will be the identifier used from the Feature object to reference these configs for a feature.
  Can be distinct from topic name
* `subscription_mode` (SubscriptionMode) - Options to configure which Kafka topics to pull data from
* `value_schema` (SchemaConfig) - Schema configuration for extracting message values from topics. At least one of key_schema and value_schema must be provided

### AuthConfig
* `mtls_config` (MtlsConfig) - Mutual-TLS authentication. See MtlsConfig
* `uc_service_credential_name` (string) - Name of the Unity Catalog service credential. This value will be set under the option databricks.serviceCredential

### BackfillSource
* `delta_table_name` (string) - The full three-part name (catalog, schema, name) of the Delta table containing the historical data to backfill
* `delta_table_source` (DeltaTableSource, deprecated) - Deprecated: Use delta_table_name instead. Kept for backwards compatibility.
  The Delta table source containing the historical data to backfill.
  Only the delta table name is used for backfill, other fields are ignored

### DeltaTableSource
* `dataframe_schema` (string) - Schema of the resulting dataframe after transformations, in Spark StructType JSON format (from df.schema.json()).
  Required if transformation_sql is specified.
  Example: {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
* `entity_columns` (list of string, deprecated) - Deprecated: Use Feature.entity instead. Kept for backwards compatibility.
  The entity columns of the Delta table
* `filter_condition` (string) - Single WHERE clause to filter delta table before applying transformations. Will be row-wise evaluated, so should only include conditionals and projections
* `full_name` (string) - The full three-part (catalog, schema, table) name of the Delta table
* `timeseries_column` (string, deprecated) - Deprecated: Use Feature.timeseries_column instead. Kept for backwards compatibility.
  The timeseries column of the Delta table
* `transformation_sql` (string) - A single SQL SELECT expression applied after filter_condition.
  Should contains all the columns needed (eg. "SELECT *, col_a + col_b AS col_c FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a + col_b AS col_c")
  If transformation_sql is not provided, all columns of the delta table are present in the DataSource dataframe

### IngestionConfig
* `backfill_job_id` (integer) - The ID of the Databricks Job that performs the historical backfill of the ingestion Delta table
* `backfill_source` (BackfillSource) - A user-provided source for backfilling data. Historical data is used when creating a training set from streaming features linked to this Stream.
  The backfill data stored in this location will be copied into the ingestion table for offline querying and training.
  The schema for this source must match exactly that of the key and payload schemas specified for this Stream
* `deduplication_columns` (list of string) - Column paths used to identify duplicate rows during ingestion; only one row per
  distinct combination of these values is kept. Use dot notation for nested fields
  (e.g. `value.user_id`). Empty list means every column is compared
* `ingestion_destination` (IngestionDestination) - Destination for the Databricks-managed Delta table that holds an offline copy of the streaming data for querying and training.
  This table contains both 1) forward-filled data from the Stream and 2) backfilled data from the BackfillSource (if provided).
  This table is created and managed by Databricks and is deleted when the Stream is deleted
* `ingestion_job_id` (integer) - The ID of the Databricks Job that performs the forward-fill ingestion
* `ingestion_pipeline_id` (string) - The ID of the SDP pipeline that continuously copies new events from the streaming source
  into the ingestion Delta table

### IngestionDestination
* `delta_table_name` (string) - The full three-part name (catalog, schema, name) of the Delta table to be created for ingestion

### MtlsConfig
* `disable_hostname_verification` (boolean) - Set to true only when the broker certificate's SAN intentionally does not match
  the connection endpoint — for example when reaching the cluster through a
  PrivateLink endpoint whose DNS name is not in the broker certificate. Skipping
  the hostname check removes a defense against man-in-the-middle attacks; do not
  enable casually. mTLS client authentication is unaffected by this option.
  
  See the Apache Kafka SSL security guide for background on this check:
  https://kafka.apache.org/42/security/encryption-and-authentication-using-ssl/#host-name-verification
* `key_password_ref` (SecretScopeReference) - Secret-scope reference for the private key password. Often the same value as the
  keystore password (keytool's default), but provided as a separate field because
  Apache Kafka requires it as a distinct option (kafka.ssl.key.password)
* `keystore_location` (string) - Unity Catalog volume path to the JKS keystore file containing the client certificate
  and private key. e.g. "/Volumes/<catalog>/<schema>/<volume>/client.jks". The
  materialization compute must have read permission on this volume
* `keystore_password_ref` (SecretScopeReference) - Secret-scope reference for the JKS keystore password
* `truststore_location` (string) - Unity Catalog volume path to the JKS truststore file containing the CA certificate(s)
  trusted to verify the Kafka broker's server certificate.
  e.g. "/Volumes/<catalog>/<schema>/<volume>/truststore.jks"
* `truststore_password_ref` (SecretScopeReference) - Secret-scope reference for the JKS truststore password

### ProtoSchemaSpec
* `message_name` (string) - The fully-qualified name of the message within schema_text that describes the Kafka payload
  (e.g. "Event" or "com.example.Event" if schema_text declares a package). Identifies which
  message is used to decode each Kafka record — a .proto file may declare multiple messages
  but only one represents the payload. Must not be empty
* `schema_text` (string) - The raw .proto file text (proto2 and proto3 syntax supported, see
  https://protobuf.dev/programming-guides/proto3/ and https://protobuf.dev/programming-guides/proto2/)

### SchemaConfig
* `avro_schema` (string) - Avro schema in JSON format (https://avro.apache.org/docs/current/specification/)
* `json_schema` (string) - Schema of the JSON object in standard IETF JSON schema format (https://json-schema.org/)
* `proto_schema` (ProtoSchemaSpec) - Protocol Buffer schema with its payload message name

### SecretScopeReference
* `key` (string) - The key within the scope
* `scope` (string) - The Databricks secret scope name

### SubscriptionMode
* `assign` (string) - A JSON string that contains the specific topic-partitions to consume from.
  For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th and 1st partitions will be consumed from
* `subscribe` (string) - A comma-separated list of Kafka topics to read from. For example, 'topicA,topicB,topicC'
* `subscribe_pattern` (string) - A regular expression matching topics to subscribe to. For example, 'topic.*' will subscribe to all topics starting with 'topic'