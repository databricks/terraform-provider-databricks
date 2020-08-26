---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: |-
  Terraform provider databricks.
---

# Databricks Provider

The Databricks provider is what is used to interact with the Databricks resources. This needs to be configured so that terraform can provision resources in your Databricks workspace on your behalf. 

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

No configuration options given to your provider will look up configured credentials in `~/.databrickscfg` file. 
It is created by `databricks configure --token` command. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication 
for docs. Config file credentials will only be used when `host`/`token` or `azure_auth` options are not provided. 
This is recommended way to use Databricks Terraform provider, in case you're using the same approach with 
[AWS Shared Credentials File](https://www.terraform.io/docs/providers/aws/index.html#shared-credentials-file) 
or [Azure CLI authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/azure_cli).

``` hcl
provider "databricks" {
}
```

One can specify non-standard location of configuration file through `config_file` parameter or `DATABRICKS_CONFIG_FILE` environment variable:

``` hcl
provider "databricks" {
  config_file = "/opt/databricks/cli-config"
}
```

One can specify a [CLI connection profile](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) through `profile` parameter or `DATABRICKS_CONFIG_PROFILE` environment variable:

``` hcl
provider "databricks" {
  profile = "ML_WORKSPACE"
}
```

### Authenticating with hostname and token

One can use `host` and `token` parameters to supply credentials to workspace. In case environment variables are preferred, `DATABRICKS_HOST` and `DATABRICKS_TOKEN` could be used instead. This is second most recommended way of configuring this provider.

``` hcl
provider "databricks" {
  host  = "http://abc-cdef-ghi.cloud.databricks.com"
  token = "dapitokenhere"
}
```

### Authenticating with hostname, username and password

!> **Warning** This approach is currently recommended only for provisioning AWS workspaces and should be avoided for regular use.

One can use `basic_auth` parameter to supply username and password credentials to workspace. `DATABRICKS_USERNAME` and `DATABRICKS_PASSWORD` environment variables could be used instead.

``` hcl
provider "databricks" {
  host = "http://abc-cdef-ghi.cloud.databricks.com"
  username = var.user
  password = var.password
}
```

## Argument Reference

The following arguments are supported by the db provider block:

* `host` - (optional) This is the host of the Databricks workspace. This is will be a url that you use to login to your workspace. 
Alternatively you can provide this value as an environment variable `DATABRICKS_HOST`.
* `token` - (optional) This is the api token to authenticate into the workspace. Alternatively you can provide this value as an 
environment variable `DATABRICKS_TOKEN`. 
* `username` - (optional) This is the username of the user that can log into the workspace. Alternatively you can provide this value as an environment variable `DATABRICKS_USERNAME`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `password` - (optional) This is the password of the user that can log into the workspace. Alternatively you can provide this value as an environment variable `DATABRICKS_PASSWORD`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `config_file` - (optional) Location of the Databricks CLI credentials file, that is created, by `databricks configure --token` command. By default, it is located in ~/.databrickscfg. Check [databricks cli documentation](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) for more details. Config file credentials will only be used when host/token/basic_auth/azure_auth are not provided. Alternatively you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to `~/.databrickscfg`. 
* `profile` - (optional) Connection profile specified within ~/.databrickscfg. Please check [connection profiles section](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) for more details. This field defaults to 
`DEFAULT`.

## Special configurations for Azure

!> **Warning** Please note that the azure service principal authentication currently uses a generated Databricks PAT token and not a AAD token for the authentication. This is due to the Databricks AAD feature not yet supporting AAD tokens for secret scopes. This will be refactored in a transparent manner when that support is enabled. The only field to be impacted is `pat_token_duration_seconds` which will be deprecated and after AAD support is fully supported. 

In order to work with Azure Databricks workspace, provider has to know it's `id` (or construct it from `azure_subscription_id`, `azure_workspace_name` and `azure_workspace_name`). Provider works with [Azure CLI authentication](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest) to facilitate local development workflows, though for automated scenarios a service principal auth is necessary (and specification of `azure_client_id`, `azure_client_secret` and `azure_tenant_id` parameters).

### Authenticating with Azure Service Principal

-> **Note** **Azure Service Principal Authentication** will only work on Azure Databricks where as the API Token authentication will work on both **Azure** and **AWS**. Internally `azure_auth` will generate a session-based PAT token.

```hcl
provider "azurerm" {
  client_id         = var.client_id
  client_secret     = var.client_secret
  tenant_id         = var.tenant_id
  subscription_id   = var.subscription_id
}

resource "azurerm_databricks_workspace" "demo_test_workspace" {
  location                      = "centralus"
  name                          = "my-workspace-name"
  resource_group_name           = var.resource_group
  sku                           = "premium"
}

provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.demo_test_workspace.id
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

-> **Note** **Azure Service Principal Authentication** will only work on Azure Databricks where as the API Token authentication will work on both **Azure** and **AWS**. Internally `azure_auth` will generate a session-based PAT token.

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_databricks_workspace" "demo_test_workspace" {
  location                      = "centralus"
  name                          = "my-workspace-name"
  resource_group_name           = var.resource_group
  sku                           = "premium"
}

provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.demo_test_workspace.id
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```

* `azure_workspace_resource_id` - (optional) `id` attribute of [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace) resource. Combination of subscription id, resource group name and workspace name. 
* `azure_workspace_name` - (optional) This is the name of your Azure Databricks Workspace. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_WORKSPACE_NAME`. Not needed with `azure_workspace_resource_id` is set.
* `azure_resource_group` - (optional) This is the resource group in which your Azure Databricks Workspace resides in. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_RESOURCE_GROUP`. Not needed with `azure_workspace_resource_id` is set.
* `azure_subscription_id` - (optional) This is the Azure Subscription id in which your Azure Databricks Workspace resides in. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID`. Not needed with `azure_workspace_resource_id` is set.
* `azure_client_secret` - (optional) This is the Azure Enterprise Application (Service principal) client secret. This service principal requires contributor access to your Azure Databricks deployment. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`.
* `azure_client_id` - (optional) This is the Azure Enterprise Application (Service principal) client id. This service principal requires contributor access to your Azure Databricks deployment. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`.
* `azure_tenant_id` - (optional) This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) 
resides in. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`.
* `pat_token_duration_seconds` - The current implementation of the azure auth via sp requires the provider to create a temporary personal access token within Databricks. The current AAD implementation does not cover all the APIs for Authentication. This field determines the duration in which that temporary PAT token is alive for. It is measured in seconds and will default to `3600` seconds. 

Where there are multiple environment variable options, the `DATABRICKS_AZURE_*` environment variables takes precedence and the `ARM_*` environment variables provide a way to share authentication configuration when using the `databricks` provider alongside the `azurerm` provider.

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

For example, with the following zero-argument configuration ...

``` hcl 
provider "databricks" {}
```

1. Provider will check all of the supported environment variables and set values of relevant arguments.
2. In case of any conflicting arguments are present, plan will end with error.
3. Will check for presence of `host` + `token` pair, continue trying otherwise.
4. Will check for `host` + `username` + `password` presence, continue trying otherwise.
5. Will check for Azure workspace ID, `azure_client_secret` + `azure_client_id` + `azure_tenant_id` presence, continue trying otherwise.
6. Will check for Azure workspace ID presence and if `az cli` is authenticated, continue trying otherwise.
7. Will check for `~/.databrickscfg` file in the home directory, will fail otherwise.
8. Will check for `profile` presence and try picking from that file, will fail otherwise.
9. Will check for `host` and `token` or `username`+`password` combination, will fail if nothing of these exist.

## Difference between apply and plan phases

Some arguments for the provider configuration might not yet be available when you run `terraform plan` because they depend on other resources that won't be available until the `terraform apply` phase.
That is why the provider doesn't use the provided arguments until you either refresh an existing resource (can be implicitly with `terraform plan`) or let the provider execute a CRUD operation.

Let's take a look at the example below:

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_databricks_workspace" "demo_test_workspace" {
  location                      = "centralus"
  name                          = "my-workspace-name"
  resource_group_name           = var.resource_group
  sku                           = "premium"
}

provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.demo_test_workspace.id
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```

Assume that we don't have any state yet, so the Databricks workspace still has to be created during our first call to `terraform apply`.
Therefore, Terraform doesn't know the actual value of the Databricks workspace resource ID. The example might use Azure CLI authentication or Azure Service Principal authentication.
The provider won't validate that you're authenticated in either way until the workspace is created and Terraform wants to create the `databricks_scim_user` resource.

Next time, when you run `terraform plan` and the provider has to refresh the `databricks_scim_user`, Terraform will already know the workspace ID and your authentication is validated during the implicit refresh of `terraform plan`.