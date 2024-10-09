---
subcategory: "Unity Catalog"
---
# databricks_external_location Data Source

-> **Note** This data source can only be used with a workspace-level provider!

Retrieves details about a [databricks_external_location](../resources/external_location.md) that were created by Terraform or manually.

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
  * `encryption_details` - The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).

## Related Resources

The following resources are used in the same context:

* [databricks_external_locations](./external_locations.md) to get names of all external locations
* [databricks_external_location](../resources/external_location.md) to manage external locations within Unity Catalog.
