---
subcategory: "Security"
---
# databricks_secret_scope Resource

Sometimes accessing data requires that you authenticate to external data sources through JDBC. Instead of directly entering your credentials into a notebook, use Databricks secrets to store your credentials and reference them in notebooks and jobs. Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

-> This resource can only be used with a workspace-level provider!

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

### keyvault_metadata

On Azure, it is possible to create Azure Databricks secret scopes backed by Azure Key Vault. Secrets are stored in Azure Key Vault and can be accessed through the Azure Databricks secrets utilities, making use of Azure Databricks access control and secret redaction. A secret scope may be configured with at most one Key Vault.

-> **Warning** To create a secret scope from Azure Key Vault, you must use one of the [Azure-specific authentication methods](../index.md#special-configurations-for-azure). Secret scopes backed by Azure Key Vault cannot be created using personal access tokens (PAT).

To define AKV access policies, you must use [azurerm_key_vault_access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_access_policy) instead of [access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault#access_policy) blocks on `azurerm_key_vault`, otherwise Terraform will remove access policies needed to access the Key Vault and the secret scope won't be in a usable state anymore.

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
  secret_permissions = ["Delete", "Get", "List", "Set"]
}

resource "databricks_secret_scope" "kv" {
  name = "keyvault-managed"

  keyvault_metadata {
    resource_id = azurerm_key_vault.this.id
    dns_name    = azurerm_key_vault.this.vault_uri
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
terraform import databricks_secret_scope.object <scopeName>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_git_folder](git_folder.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_acl](secret_acl.md) to manage access to [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
