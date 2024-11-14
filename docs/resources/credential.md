---
subcategory: "Unity Catalog"
---
# databricks_credential Resource

-> This resource can only be used with a workspace-level provider.

-> This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

A credential represents an authentication and authorization mechanism for accessing services on your cloud tenant. Each credential is subject to Unity Catalog access-control policies that control which users and groups can access the credential.

To create credentials, you must be a Databricks account admin or have the `CREATE SERVICE CREDENTIAL` privilege. The user who creates the credential can delegate ownership to another user or group to manage permissions on it

On AWS, the IAM role for a credential requires a trust policy. See [documentation](https://docs.databricks.com/en/connect/unity-catalog/cloud-services/service-credentials.html#step-1-create-an-iam-role) for more details. The data source [databricks_aws_unity_catalog_assume_role_policy](../data-sources/aws_unity_catalog_assume_role_policy.md) can be used to create the necessary AWS Unity Catalog assume role policy.

## Example Usage

For AWS

```hcl
resource "databricks_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  purpose = "SERVICE"
  comment = "Managed by TF"
}

resource "databricks_grants" "external_creds" {
  credential = databricks_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["ACCESS"]
  }
}
```

For Azure

```hcl
resource "databricks_credential" "external_mi" {
  name = "mi_credential"
  azure_managed_identity {
    access_connector_id = azurerm_databricks_access_connector.example.id
  }
  purpose = "SERVICE"
  comment = "Managed identity credential managed by TF"
}

resource "databricks_grants" "external_creds" {
  credential = databricks_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["ACCESS"]
  }
}
```

## Argument Reference

The following arguments are required:

- `name` - Name of Credentials, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
- `purpose` - Indicates the purpose of the credential. Can be `SERVICE`.
- `owner` - (Optional) Username/groupname/sp application_id of the credential owner.
- `read_only` - (Optional) Indicates whether the credential is only usable for read operations.
- `skip_validation` - (Optional) Suppress validation errors if any & force save the credential.
- `force_destroy` - (Optional) Delete credential regardless of its dependencies.
- `force_update` - (Optional) Update credential regardless of its dependents.
- `isolation_mode` - (Optional) Whether the credential is accessible from all workspaces or a specific set of workspaces. Can be `ISOLATION_MODE_ISOLATED` or `ISOLATION_MODE_OPEN`. Setting the credential to `ISOLATION_MODE_ISOLATED` will automatically allow access from the current workspace.

`aws_iam_role` optional configuration block for credential details for AWS:

- `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`

`azure_managed_identity` optional configuration block for using managed identity as credential details for Azure (recommended over service principal):

- `access_connector_id` - The Resource ID of the Azure Databricks Access Connector resource, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name`.

- `managed_identity_id` - (Optional) The Resource ID of the Azure User Assigned Managed Identity associated with Azure Databricks Access Connector, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-managed-identity-name`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this credential - same as the `name`.
- `credential_id` - Unique ID of the credential.

## Import

This resource can be imported by name:

```bash
terraform import databricks_credential.this <name>
```
