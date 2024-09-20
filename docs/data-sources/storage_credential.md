---
subcategory: "Unity Catalog"
---
# databricks_storage_credential Data Source

-> **Note** This data source can only be used with a workspace-level provider!

Retrieves details about a [databricks_storage_credential](../resources/storage_credential.md) that were created by Terraform or manually.

## Example Usage

Getting details of an existing storage credential in the metastore

```hcl
data "databricks_storage_credential" "this" {
  name = "this"
}

output "created_by" {
  value = data.databricks_storage_credential.this.storage_credential_info[0].created_by
}
```

## Argument Reference

* `name` - (Required) The name of the storage credential

## Attribute Reference

This data source exports the following attributes:

* `id` - Unique ID of storage credential.
* `storage_credential_info` - array of objects with information about storage credential.
  * `metastore_id` - Unique identifier of the parent Metastore.
  * `owner` - Username/groupname/sp application_id of the storage credential owner.
  * `read_only` - Indicates whether the storage credential is only usable for read operations.
  * `created_at` - Time at which this catalog was created, in epoch milliseconds.
  * `created_by` - Username of catalog creator.
  * `updated_at` - Time at which this catalog was last modified, in epoch milliseconds.
  * `updated_by` - Username of user who last modified catalog.
  * `aws_iam_role` credential details for AWS:
    * `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`
    * `external_id` (output only) - The external ID used in role assumption to prevent confused deputy problem.
    * `unity_catalog_iam_arn` (output only) - The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks. This is the identity that is going to assume the AWS IAM role.
  * `azure_managed_identity` managed identity credential details for Azure
    * `access_connector_id` - The Resource ID of the Azure Databricks Access Connector resource, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name`.
    * `managed_identity_id` - The Resource ID of the Azure User Assigned Managed Identity associated with Azure Databricks Access Connector, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-managed-identity-name`.
  * `azure_service_principal` service principal credential details for Azure:
    * `directory_id` - The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application
    * `application_id` - The application ID of the application registration within the referenced AAD tenant
  * `databricks_gcp_service_account` credential details for GCP:
    * `email` - The email of the GCP service account created, to be granted access to relevant buckets.

## Related Resources

The following resources are used in the same context:

* [databricks_storage_credentials](./storage_credentials.md) to get names of all credentials
* [databricks_storage_credential](../resources/storage_credential.md) to manage Storage Credentials within Unity Catalog.
