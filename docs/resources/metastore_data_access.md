---
subcategory: "Unity Catalog"
---
# databricks_metastore_data_access (Resource)

Each [databricks_metastore](docs/resources/metastore.md) requires an IAM role that will be assumed by Unity Catalog to access data. `databricks_metastore_data_access` defines this

## Example Usage

For AWS

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = "uc admins"
  region        = "us-east-1"
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  metastore_id = databricks_metastore.this.id
  name         = aws_iam_role.metastore_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.metastore_data_access.arn
  }
  is_default = true
}
```

For Azure using managed identity as credential (recommended)

```hcl
resource "databricks_metastore" "this" {
  name = "primary"
  storage_root = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_container.unity_catalog.name,
  azurerm_storage_account.unity_catalog.name)
  owner         = "uc admins"
  region        = "eastus"
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  metastore_id = databricks_metastore.this.id
  name         = "mi_dac"
  azure_managed_identity {
    access_connector_id = var.access_connector_id
  }
  is_default = true
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of Data Access Configuration, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
* `metastore_id` - Unique identifier of the parent Metastore
* `owner` - (Optional) Username/groupname/sp application_id of the data access configuration owner.
* `force_destroy` - (Optional) Delete the data access configuration regardless of its dependencies.

`aws_iam_role` optional configuration block for credential details for AWS:

* `role_arn` - The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`

`azure_managed_identity` optional configuration block for using managed identity as credential details for Azure (Recommended):

* `access_connector_id` - The Resource ID of the Azure Databricks Access Connector resource, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name`.
* `managed_identity_id` - (Optional) The Resource ID of the Azure User Assigned Managed Identity associated with Azure Databricks Access Connector, of the form `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-managed-identity-name`.

`databricks_gcp_service_account` optional configuration block for creating a Databricks-managed GCP Service Account:

* `email` (output only) - The email of the GCP service account created, to be granted access to relevant buckets.

`azure_service_principal` optional configuration block for credential details for Azure (Legacy):

* `directory_id` - The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application
* `application_id` - The application ID of the application registration within the referenced AAD tenant
* `client_secret` - The client secret generated for the above app ID in AAD. **This field is redacted on output**

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this data access configuration in form of `<metastore_id>|<name>`.

## Import

This resource can be imported by combination of metastore id and the data access name.

```bash
terraform import databricks_metastore_data_access.this '<metastore_id>|<name>'
```
