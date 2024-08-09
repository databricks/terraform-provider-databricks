---
subcategory: "Unity Catalog"
---
# databricks_storage_credential Resource

-> **Note** This resource could be used with account or workspace-level provider.

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:

- `databricks_storage_credential` represents authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal/managed identity for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location.

## Example Usage

For AWS

```hcl
resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_EXTERNAL_TABLE"]
  }
}
```

For Azure

```hcl
resource "databricks_storage_credential" "external_mi" {
  name = "mi_credential"
  azure_managed_identity {
    access_connector_id = azurerm_databricks_access_connector.example.id
  }
  comment = "Managed identity credential managed by TF"
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_EXTERNAL_TABLE"]
  }
}
```

For GCP

```hcl
resource "databricks_storage_credential" "external" {
  name = "the-creds"
  databricks_gcp_service_account {}
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_EXTERNAL_TABLE"]
  }
}
```

## Argument Reference

The following arguments are required:

- `name` - Name of Storage Credentials, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
- `metastore_id` - (Required for account-level) Unique identifier of the parent Metastore. If set for workspace-level, it must match the ID of the metastore assigned to the worspace. When changing the metastore assigned to a workspace, this field becomes required.
- `owner` - (Optional) Username/groupname/sp application_id of the storage credential owner.
- `read_only` - (Optional) Indicates whether the storage credential is only usable for read operations.
- `skip_validation` - (Optional) Suppress validation errors if any & force save the storage credential.
- `force_destroy` - (Optional) Delete storage credential regardless of its dependencies.
- `force_update` - (Optional) Update storage credential regardless of its dependents.
<<<<<<< HEAD
- `isolation_mode` - (Optional) Whether the storage credential is accessible from all workspaces or a specific set of workspaces. Can be `ISOLATED` or `OPEN`. Setting the credential to `ISOLATED` will automatically allow access from the current workspace.
=======
- `isolation_mode` - (Optional) Whether the storage credential is accessible from all workspaces or a specific set of workspaces. Can be `ISOLATION_MODE_ISOLATED` or `ISOLATION_MODE_OPEN`. Setting the credential to `ISOLATION_MODE_ISOLATED` will automatically allow access from the current workspace.
>>>>>>> 1a309c8195c9779dadd9a337e1dbd3496815833a

`aws_iam_role` optional configuration block for credential details for AWS:

- `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`

`azure_managed_identity` optional configuration block for using managed identity as credential details for Azure (recommended over service principal):

- `access_connector_id` - The Resource ID of the Azure Databricks Access Connector resource, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name`.

- `managed_identity_id` - (Optional) The Resource ID of the Azure User Assigned Managed Identity associated with Azure Databricks Access Connector, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-managed-identity-name`.

`databricks_gcp_service_account` optional configuration block for creating a Databricks-managed GCP Service Account:

- `email` (output only) - The email of the GCP service account created, to be granted access to relevant buckets.

`azure_service_principal` optional configuration block to use service principal as credential details for Azure (Legacy):

- `directory_id` - The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application
- `application_id` - The application ID of the application registration within the referenced AAD tenant
- `client_secret` - The client secret generated for the above app ID in AAD. **This field is redacted on output**

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this storage credential - same as the `name`.
- `storage_credential_id` - Unique ID of storage credential.

## Import

This resource can be imported by name:

```bash
terraform import databricks_storage_credential.this <name>
```
