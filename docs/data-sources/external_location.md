---
subcategory: "Unity Catalog"
---
# databricks_external_location Data Source

-> **Note** This data source could be only used with workspace-level provider!

Retrieves details about a [databricks_external_location](../resources/external_location.md) that were created by Terraform or manually.

## Example Usage

Getting details of an existing external location in the metastore

```hcl
data "databricks_external_location" "this" {
  name = "this"
}

output "created_by" {
  value     = data.databricks_external_location.this.created_by
  sensitive = false
}
```

## Argument Reference

* `name` - (Required) The name of the storage credential

## Attribute Reference

* `url` - Path URL in cloud storage, of the form: `s3://[bucket-host]/[bucket-dir]` (AWS), `abfss://[user]@[host]/[path]` (Azure), `gs://[bucket-host]/[bucket-dir]` (GCP).
* `credential_name` - Name of the [databricks_storage_credential](storage_credential.md) to use with this external location.
* `owner` - Username/groupname/sp application_id of the external location owner.
* `comment` - User-supplied comment.
* `read_only` - Indicates whether the external location is read-only.
* `access_point` - The ARN of the s3 access point to use with the external location (AWS).
* `encryption_details` - The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).

## Related Resources

The following resources are used in the same context:

* [databricks_external_locations](./external_locations.md) to get names of all external locations
* [databricks_external_location](../resources/external_location.md) to manage external locations within Unity Catalog.
