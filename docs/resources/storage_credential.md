---
subcategory: "Unity Catalog"
---
# databricks_storage_credential Resource

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
    privileges = ["CREATE_TABLE"]
  }
}
```

For Azure

```hcl
data "azurerm_resource_group" "this" {
  name = "example-rg"
}

resource "azurerm_databricks_access_connector" "example" {
  name                = "databrickstest"
  resource_group_name = azurerm_resource_group.this.name
  location            = azurerm_resource_group.this.location
  identity {
    type = "SystemAssigned"
  }
  tags = {
    Environment = "Production"
  }
}

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
    privileges = ["CREATE_TABLE"]
  }
}
```

## Argument Reference

The following arguments are required:

- `name` - Name of Storage Credentials, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
- `owner` - (Optional) Username/groupname/sp application_id of the storage credential owner.


`aws_iam_role` optional configuration block for credential details for AWS:

- `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`

`azure_managed_identity` optional configuration block for using managed identity as credential details for Azure (recommended over service principal):

- `access_connector_id` - The Resource ID of the Azure Databricks Access Connector resource, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name`

`azure_service_principal` optional configuration block to use service principal as credential details for Azure:

- `directory_id` - The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application
- `application_id` - The application ID of the application registration within the referenced AAD tenant
- `client_secret` - The client secret generated for the above app ID in AAD. **This field is redacted on output**

## Import

This resource can be imported by name:

```bash
terraform import databricks_storage_credential.this <name>
```
