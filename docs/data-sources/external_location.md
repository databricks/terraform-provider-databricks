---
subcategory: "Unity Catalog"
---
# databricks_external_location Data Source

Retrieves details about a [databricks_external_location](../resources/external_location.md) that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Getting details of an existing external location in the metastore

```hcl
data "databricks_external_location" "this" {
  name = "this"
}

output "created_by" {
  value = data.databricks_external_location.this.external_location_info[0].created_by
}
```

## Argument Reference

* `name` - (Required) The name of the external location

## Attribute Reference

This data source exports the following attributes:

* `id` - external location ID - same as name.
* `external_location_info` - array of objects with information about external location:
  * `url` - Path URL in cloud storage, of the form: `s3://[bucket-host]/[bucket-dir]` (AWS), `abfss://[user]@[host]/[path]` (Azure), `gs://[bucket-host]/[bucket-dir]` (GCP).
  * `credential_name` - Name of the [databricks_storage_credential](storage_credential.md) to use with this external location.
  * `credential_id` - Unique ID of storage credential.
  * `metastore_id` - Unique identifier of the parent Metastore.
  * `owner` - Username/groupname/sp application_id of the external location owner.
  * `comment` - User-supplied comment.
  * `read_only` - Indicates whether the external location is read-only.
  * `created_at` - Time at which this catalog was created, in epoch milliseconds.
  * `created_by` - Username of catalog creator.
  * `updated_at` - Time at which this catalog was last modified, in epoch milliseconds.
  * `updated_by` - Username of user who last modified catalog.
  * `access_point` - The ARN of the s3 access point to use with the external location (AWS).
  * `encryption_details` - A block describing encryption options that apply to clients connecting to cloud storage. Consisting of the following attributes:
    * sse_encryption_details - a block describing server-Side Encryption properties for clients communicating with AWS S3. Consists of the following attributes:
      * `algorithm` - Encryption algorithm value. Sets the value of the `x-amz-server-side-encryption` header in S3 request.
      * `aws_kms_key_arn` - ARN of the SSE-KMS key used with the S3 location, when `algorithm = "SSE-KMS"`. 

## Related Resources

The following resources are used in the same context:

* [databricks_external_locations](./external_locations.md) to get names of all external locations
* [databricks_external_location](../resources/external_location.md) to manage external locations within Unity Catalog.
