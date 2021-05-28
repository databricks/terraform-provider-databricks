---
subcategory: "Security"
---
# databricks_secret_scope Resource

Sometimes accessing data requires that you authenticate to external data sources through JDBC. Instead of directly entering your credentials into a notebook, use Databricks secrets to store your credentials and reference them in notebooks and jobs. Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

## Example Usage

```hcl
resource "databricks_secret_scope" "this" {
  name = "terraform-demo-scope"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
* `initial_manage_principal` - (Optional) The principal with the only possible value `users` that is initially granted `MANAGE` permission to the created scope.  If it's omitted, then the [databricks_secret_acl](secret_acl.md) with `MANAGE` permission applied to the scope is assigned to the API request issuer's user identity (see [documentation](https://docs.databricks.com/dev-tools/api/latest/secrets.html#create-secret-scope)). This part of the state cannot be imported.

## keyvault_metadata

On Azure it's possible to create and manage secrets in Azure Key Vault and have use Azure Databricks secret redaction & access control functionality for reading them. There has to be a single Key Vault per single secret scope. To define AKV access policies, you must use [azurerm_key_vault_access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_access_policy) instead of [access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault#access_policy) blocks on `azurerm_key_vault`, otherwise Terraform will remove access policies needed to access the Key Vault and secret scope won't be in a usable state anymore.

-> **Note** Currently, it's only possible to create Azure Key Vault scopes with Azure CLI authentication and not with Service Principal. That means, `az login --service-principal --username $ARM_CLIENT_ID --password $ARM_CLIENT_SECRET --tenant $ARM_TENANT_ID` won't work as well. This is the limitation from underlying cloud resources.

```hcl
data "azurerm_client_config" "current" {
}

resource "azurerm_key_vault" "this" {
  name                     = "${var.prefix}-kv"
  location                 = azurerm_resource_group.example.location
  resource_group_name      = azurerm_resource_group.example.name
  tenant_id                = data.azurerm_client_config.current.tenant_id
  soft_delete_enabled      = false
  purge_protection_enabled = false
  sku_name                 = "standard"
  tags                     = var.tags
}

resource "azurerm_key_vault_access_policy" "this" {
  key_vault_id       = azurerm_key_vault.this.id
  tenant_id          = data.azurerm_client_config.current.tenant_id
  object_id          = data.azurerm_client_config.current.object_id
  secret_permissions = ["delete", "get", "list", "set"]
}

resource "databricks_secret_scope" "kv" {
  name = "keyvault-managed"

  keyvault_metadata {
    resource_id = azurerm_key_vault.this.id
    dns_name = azurerm_key_vault.this.vault_uri
  }
}
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the secret scope object.
* `backend_type` - Either `DATABRICKS` or `AZURE_KEYVAULT`

## Import

The secret resource scope can be imported using the scope name. `initial_manage_principal` state won't be imported, because the underlying API doesn't include it in the response.

```bash
$ terraform import databricks_secret_scope.object <scopeName>
```
