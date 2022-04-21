---
subcategory: "Unity Catalog"
---
# databricks_storage_credential Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:
- `databricks_storage_credential` represents authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
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
resource "databricks_storage_credential" "external" {
  name = azuread_application.ext_cred.display_name
  azure_service_principal {
    directory_id   = var.tenant_id
    application_id = azuread_application.ext_cred.application_id
    client_secret  = azuread_application_password.ext_cred.value
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

## Argument Reference

The following arguments are required:

* `name` - Name of Storage Credentials, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.

`aws_iam_role` optional configuration block for credential details for AWS:
* `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`

`azure_service_principal` optional configuration block for credential details for Azure:
* `directory_id` - The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application
* `application_id` - The application ID of the application registration within the referenced AAD tenant
* `client_secret` - The client secret generated for the above app ID in AAD. **This field is redacted on output**

## Import

This resource can be imported by name:

```bash
$ terraform import databricks_storage_credential.this <name>
```
