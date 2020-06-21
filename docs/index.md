---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: |-
  Terraform provider databricks.
---

# Authentication

!> **Warning** Please be aware that hard coding any credentials is not something that is recommended. It may be best if 
you store the credentials environment variables, `~/.databrickscfg` file or use tfvars file.

The Databricks provider is what is used to interact with the Databricks resources. This needs to be configured so that 
terraform can provision resources in your Databricks workspace on your behalf. There are currently three supported methods [to authenticate into](https://docs.databricks.com/dev-tools/api/latest/authentication.html) the Databricks platform to create resources:

* [PAT Tokens](https://docs.databricks.com/dev-tools/api/latest/authentication.html)
* Username+Password pair
* Azure Active Directory Tokens

## Authenticating with Databricks CLI credentials

No configuration options given to your provider will look up configured credentials in `~/.databrickscfg` file. It is created by `databricks configure --token` command. Check  https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for docs. Config file credetials will only be used when `host`/`token` or `azure_auth` options are not provided. This is recommended way to use Databricks Terraform provider, in case you're using the same approach with [AWS Shared Credentials File](https://www.terraform.io/docs/providers/aws/index.html#shared-credentials-file) or [Azure CLI authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/azure_cli).

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

## Authenticating with hostname and token

One can use `host` and `token` parameters to supply credentials to workspace. In case environment variables are preferred, `DATABRICKS_HOST` and `DATABRICKS_TOKEN` could be used instead. This is second most recommended way of configuring this provider.

``` hcl
provider "databricks" {
  host  = "http://abc-cdef-ghi.cloud.databricks.com"
  token = "dapitokenhere"
}
```

## Authenticating with hostname, username and password

!> **Warning** This approach is currently recommended only for provisioning AWS workspaces and should be avoided for regular use.

One can use `basic_auth` parameter to supply username and password credentials to workspace. `DATABRICKS_USERNAME` and `DATABRICKS_PASSWORD` environment variables could be used instead.

``` hcl
provider "databricks" {
  host = "http://abc-cdef-ghi.cloud.databricks.com"
  basic_auth {
    username = var.user
    password = var.password
  }
}
```

## Authenticating with Azure Service Principal

-> **Note** **Azure Service Principal Authentication** will only work on Azure Databricks where as the API Token authentication will work on both **Azure** and **AWS**. Internally `azure_auth` will generate a session-based PAT token.

``` hcl
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
  managed_resource_group_name   = var.managed_resource_group_name
  sku                           = "premium"
}

provider "databricks" {
  azure_auth {
    managed_resource_group  = azurerm_databricks_workspace.demo_test_workspace.managed_resource_group_name
    azure_region            = azurerm_databricks_workspace.demo_test_workspace.location
    workspace_name          = azurerm_databricks_workspace.demo_test_workspace.name
    resource_group          = azurerm_databricks_workspace.demo_test_workspace.resource_group_name
    client_id               = var.client_id
    client_secret           = var.client_secret
    tenant_id               = var.tenant_id
    subscription_id         = var.subscription_id
  }
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```

# Argument Reference

The following arguments are supported by the db provider block:

* `host` - (optional) This is the host of the Databricks workspace. This is will be a url that you use to login to your workspace. 
Alternatively you can provide this value as an environment variable `DATABRICKS_HOST`.

* `token` - (optional) This is the api token to authenticate into the workspace. Alternatively you can provide this value as an 
environment variable `DATABRICKS_TOKEN`. 

* `basic_auth` - (optional) This is a basic_auth block ([documented below](#basic_auth-configuration-block)) to authenticate to the Databricks via basic auth through a user 
that has access to the workspace. This is optional as you can use the api token based auth.

* `config_file` - (optional) Location of the Databricks CLI credentials file, that is created, by `databricks configure --token` command. 
By default, it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication 
for docs. Config file credentials will only be used when host/token/basic_auth/azure_auth are not provided. 
Alternatively you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to 
`~/.databrickscfg`. 

* `profile` - (optional) Connection profile specified within ~/.databrickscfg. Please check 
https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation. This field defaults to 
`DEFAULT`.

* `azure_auth` - (optional) This is a azure_auth block ([documented below]((#azure_auth-configuration-block))) required to authenticate to the Databricks via an azure service 
principal that has access to the workspace. This is optional as you can use the api token based auth. 

## basic_auth Configuration Block

Example:

```hcl
basic_auth = {
    username = "user"
    password = "mypass-123"
}
```

The basic_auth block contains the following arguments:

* `username` - (required) This is the username of the user that can log into the workspace. 
Alternatively you can provide this value as an environment variable `DATABRICKS_USERNAME`.

* `password` - (required) This is the password of the user that can log into the workspace.
Alternatively you can provide this value as an environment variable `DATABRICKS_PASSWORD`.


## azure_auth Configuration Block

Example:

```hcl
azure_auth = {
    azure_region = "centralus"
    managed_resource_group = "my-databricks-managed-rg"
    workspace_name = "test-managed-workspace"
    resource_group = "1-test-rg"
    client_id = var.client_id
    client_secret = var.client_secret
    tenant_id = var.tenant_id
    subscription_id = var.subscription_id
}
```

This is the authentication required to authenticate to the Databricks via an azure service 
principal that has access to the workspace. This is optional as you can use the api token based auth. 
The azure_auth block contains the following arguments:

* `managed_resource_group` - (required) This is the managed workgroup id when the Databricks workspace is provisioned. 
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP`.

* `azure_region` - (required) This is the azure region in which your workspace is deployed.
Alternatively you can provide this value as an environment variable `AZURE_REGION`.

* `workspace_name` - (required) This is the name of your Azure Databricks Workspace.
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_WORKSPACE_NAME`.

* `resource_group` - (required) This is the resource group in which your Azure Databricks Workspace resides in.
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_RESOURCE_GROUP`.

* `subscription_id` - (required) This is the Azure Subscription id in which your Azure Databricks Workspace resides in. 
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID`.
                                                                              
* `client_secret` - (required) This is the Azure Enterprise Application (Service principal) client secret. This service 
principal requires contributor access to your Azure Databricks deployment. Alternatively you can provide this 
value as an environment variable `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`.

* `client_id` - (required) This is the Azure Enterprise Application (Service principal) client id. This service principal 
requires contributor access to your Azure Databricks deployment. Alternatively you can provide this value as an 
environment variable `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`.

* `tenant_id` - (required) This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) 
resides in. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`.

Where there are multiple environment variable options, the `DATABRICKS_AZURE_*` environment variables takes precedence 
and the `ARM_*` environment variables provide a way to share authentication configuration when using the `databricks-terraform` 
provider alongside the `azurerm` provider.

# Environment variables

The following variables can be passed via environment variables:

* `host` → `DATABRICKS_HOST`
* `token` → `DATABRICKS_TOKEN`
* `basic_auth.username` → `DATABRICKS_USERNAME`
* `basic_auth.password` → `DATABRICKS_PASSWORD`
* `config_file` → `DATABRICKS_CONFIG_FILE`
* `managed_resource_group` → `DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP`
* `azure_region` → `AZURE_REGION`
* `workspace_name` → `DATABRICKS_AZURE_WORKSPACE_NAME`
* `resource_group` → `DATABRICKS_AZURE_RESOURCE_GROUP`
* `subscription_id` → `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID`
* `client_secret` → `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`
* `client_id` → `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`
* `tenant_id` → `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`

For example you can have the following provider definition:

``` hcl 
provider "databricks" {}
```

Then run the following code and the following environment variables will be injected into the provider.    

``` bash
$ export HOST="http://databricks.domain.com"
$ export TOKEN="dapitokenhere"
$ terraform plan
```