---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: |-
  Terraform provider databricks.
---

# Databricks Provider

Use the Databricks Terraform provider to interact with almost all of [Databricks](http://databricks.com/) resources. 

## Example Usage

```hcl
provider "databricks" {
  host = "https://abc-defg-024.cloud.databricks.com/"
  token = "<your PAT token>"
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20

  autoscale {
    min_workers = 1
    max_workers = 50
  }
}
```

## Authentication

!> **Warning** Please be aware that hard coding any credentials in plain text is not something that is recommended. We strongly recommend using Terraform backend that supports encryption. Please use [environment variables](#environment-variables), `~/.databrickscfg` file, encrypted `.tfvars` files or secret store of your choice (Hashicorp [Vault](https://www.vaultproject.io/), AWS [Secrets Manager](https://aws.amazon.com/secrets-manager/), AWS [Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html), Azure [Key Vault](https://azure.microsoft.com/en-us/services/key-vault/))


There are currently three supported methods [to authenticate into](https://docs.databricks.com/dev-tools/api/latest/authentication.html) the Databricks platform to create resources:

* [PAT Tokens](https://docs.databricks.com/dev-tools/api/latest/authentication.html)
* Username and password pair
* Azure Active Directory Tokens via [Azure CLI](#authenticating-with-azure-cli) or [Service Principals](#authenticating-with-azure-service-principal)

### Authenticating with Databricks CLI credentials

No configuration options given to your provider will look up configured credentials in `~/.databrickscfg` file. It is created by the `databricks configure --token` command. Check [this page](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) 
for more details. The provider uses config file credentials only when `host`/`token` or `azure_auth` options are not specified. 
It is the recommended way to use Databricks Terraform provider, in case you're already using the same approach with 
[AWS Shared Credentials File](https://www.terraform.io/docs/providers/aws/index.html#shared-credentials-file) 
or [Azure CLI authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/azure_cli).

``` hcl
provider "databricks" {
}
```

You can specify non-standard location of configuration file through `config_file` parameter or `DATABRICKS_CONFIG_FILE` environment variable:

``` hcl
provider "databricks" {
  config_file = "/opt/databricks/cli-config"
}
```

You can specify a [CLI connection profile](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) through `profile` parameter or `DATABRICKS_CONFIG_PROFILE` environment variable:

``` hcl
provider "databricks" {
  profile = "ML_WORKSPACE"
}
```

### Authenticating with hostname and token

You can use `host` and `token` parameters to supply credentials to the workspace. When environment variables are preferred, then you can specify `DATABRICKS_HOST` and `DATABRICKS_TOKEN` instead. Environment variables are the second most recommended way of configuring this provider.

``` hcl
provider "databricks" {
  host  = "http://abc-cdef-ghi.cloud.databricks.com"
  token = "dapitokenhere"
}
```

### Authenticating with hostname, username, and password

!> **Warning** This approach is currently recommended only for provisioning AWS workspaces and should be avoided for regular use.

You can use the `basic_auth` parameter to supply username and password credentials to the workspace. You can use `DATABRICKS_USERNAME` and `DATABRICKS_PASSWORD` environment variables instead.

``` hcl
provider "databricks" {
  host = "http://abc-cdef-ghi.cloud.databricks.com"
  username = var.user
  password = var.password
}
```

## Argument Reference

The provider block supports the following arguments:

* `host` - (optional) This is the host of the Databricks workspace. It is a URL that you use to login to your workspace. 
Alternatively, you can provide this value as an environment variable `DATABRICKS_HOST`.
* `token` - (optional) This is the API token to authenticate into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_TOKEN`. 
* `username` - (optional) This is the username of the user that can log into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_USERNAME`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `password` - (optional) This is the user's password that can log into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_PASSWORD`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `config_file` - (optional) Location of the Databricks CLI credentials file created by `databricks configure --token` command (~/.databrickscfg by default). Check [Databricks CLI documentation](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) for more details. The provider uses configuration file credentials when you don't specify host/token/basic_auth/azure attributes. Alternatively, you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to `~/.databrickscfg`. 
* `profile` - (optional) Connection profile specified within ~/.databrickscfg. Please check [connection profiles section](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) for more details. This field defaults to 
`DEFAULT`.

## Special configurations for Azure

To work with Azure Databricks workspace, the provider must know its `id` (or construct it from `azure_subscription_id`, `azure_workspace_name` and `azure_workspace_name`). The provider works with [Azure CLI authentication](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest) to facilitate local development workflows, though for automated scenarios a service principal auth is necessary (and specification of `azure_client_id`, `azure_client_secret` and `azure_tenant_id` parameters).

### Authenticating with Azure Service Principal

!> **Warning** Please note that the azure service principal authentication currently uses a generated Databricks PAT token and not an AAD token for the authentication. Azure Databricks does not yet support AAD tokens for [secret scopes](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/secrets#--create-secret-scope). Databricks Labs team will refactor it transparently once that support is available. The only impacted field is `pat_token_duration_seconds`, which will be deprecated and fully supported after AAD support. 

```hcl
provider "azurerm" {
  client_id         = var.client_id
  client_secret     = var.client_secret
  tenant_id         = var.tenant_id
  subscription_id   = var.subscription_id
}

resource "azurerm_databricks_workspace" "this" {
  location                      = "centralus"
  name                          = "my-workspace-name"
  resource_group_name           = var.resource_group
  sku                           = "premium"
}

provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.this.id
  azure_client_id             = var.client_id
  azure_client_secret         = var.client_secret
  azure_tenant_id             = var.tenant_id
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```

### Authenticating with Azure CLI

It's possible to use _experimental_ [Azure CLI](https://docs.microsoft.com/cli/azure/) authentication, where the provider would rely on access token cached by `az login` command so that local development scenarios are possible. Technically, the provider will call `az account get-access-token` each time before an access token is about to expire. It is [verified to work](https://github.com/databrickslabs/terraform-provider-databricks/pull/282) with all API. It could be turned off by setting `azure_use_pat_for_cli` to `true` on provider configuration.

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_databricks_workspace" "this" {
  location                      = "centralus"
  name                          = "my-workspace-name"
  resource_group_name           = var.resource_group
  sku                           = "premium"
}

provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.this.id
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```

* `azure_workspace_resource_id` - (optional) `id` attribute of [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace) resource. Combination of subscription id, resource group name, and workspace name. 
* `azure_workspace_name` - (optional) This is the name of your Azure Databricks Workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_AZURE_WORKSPACE_NAME`. Not needed with `azure_workspace_resource_id` is set.
* `azure_resource_group` - (optional) This is the resource group in which your Azure Databricks Workspace resides. Alternatively, you can provide this value as an environment variable `DATABRICKS_AZURE_RESOURCE_GROUP`. Not needed with `azure_workspace_resource_id` is set.
* `azure_subscription_id` - (optional) This is the Azure Subscription id in which your Azure Databricks Workspace resides. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID`. Not needed with `azure_workspace_resource_id` is set.
* `azure_client_secret` - (optional) This is the Azure Enterprise Application (Service principal) client secret. This service principal requires contributor access to your Azure Databricks deployment. Alternatively, you can provide this value as an environment variable `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`.
* `azure_client_id` - (optional) This is the Azure Enterprise Application (Service principal) client id. This service principal requires contributor access to your Azure Databricks deployment. Alternatively, you can provide this value as an environment variable `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`.
* `azure_tenant_id` - (optional) This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) 
resides. Alternatively, you can provide this value as an environment variable `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`.
* `pat_token_duration_seconds` - The current implementation of the azure auth via sp requires the provider to create a temporary personal access token within Databricks. The current AAD implementation does not cover all the APIs for Authentication. This field determines the duration in which that temporary PAT token is alive. It is measured in seconds and will default to `3600` seconds. 

There are multiple environment variable options, the `DATABRICKS_AZURE_*` environment variables take precedence, and the `ARM_*` environment variables provide a way to share authentication configuration using the `databricks` provider alongside the `azurerm` provider.

## Environment variables

The following configuration attributes can be passed via environment variables:

| Argument | Environment variable |
| --: | --- |
| `host` | `DATABRICKS_HOST` |
| `token` | `DATABRICKS_TOKEN` |
| `username` | `DATABRICKS_USERNAME` |
| `password` | `DATABRICKS_PASSWORD` |
| `config_file` | `DATABRICKS_CONFIG_FILE` |
| `profile` | `DATABRICKS_CONFIG_PROFILE` |
| `azure_workspace_resource_id` | `DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID` |
| `azure_workspace_name` | `DATABRICKS_AZURE_WORKSPACE_NAME` |
| `azure_resource_group` | `DATABRICKS_AZURE_RESOURCE_GROUP` |
| `azure_subscription_id` | `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID` |
| `azure_client_secret` | `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET` |
| `azure_client_id` | `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID` |
| `azure_tenant_id` | `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID` |

## Empty provider block

For example, with the following zero-argument configuration:

``` hcl 
provider "databricks" {}
```

1. Provider will check all of the supported environment variables and set values of relevant arguments.
2. In case any conflicting arguments are present, the plan will end with an error.
3. Will check for the presence of `host` + `token` pair, continue trying otherwise.
4. Will check for `host` + `username` + `password` presence, continue trying otherwise.
5. Will check for Azure workspace ID, `azure_client_secret` + `azure_client_id` + `azure_tenant_id` presence, continue trying otherwise.
6. Will check for Azure workspace ID presence, and if `AZ CLI` returns an access token, continue trying otherwise.
7. Will check for the `~/.databrickscfg` file in the home directory, will fail otherwise.
8. Will check for `profile` presence and try picking from that file will fail otherwise.
9. Will check for `host` and `token` or `username`+`password` combination, will fail if nothing of these exist.

## Project Support

**Important:** Projects in the `databrickslabs` GitHub account, including the Databricks Terraform Provider, are not formally supported by Databricks. They are maintained by Databricks Field teams and provided as-is. There is no service level agreement (SLA). Databricks makes no guarantees of any kind. If you discover an issue with the provider, please file a GitHub Issue on the repo, and it will be reviewed by project maintainers as time permits.
