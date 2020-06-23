+++
title = "Provider"
date = 2020-04-20T23:34:03-04:00
weight = 10
chapter = false
pre = ""
+++

## Databricks Provider

The Databricks provider is what is used to interact with the Databricks resources. This needs to be configured so that 
terraform can provision resources in your Databricks workspace on your behalf.  

## Example Usages

{{< tabs groupId="providerConfigs" >}}
 {{% tab name="Token" %}}
``` hcl
provider "databricks" {
  host  = "http://databricks.domain.com"
  token = "dapitokenhere"
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```
 {{% /tab %}}
 {{% tab name="Basic Auth" %}}
``` hcl
provider "databricks" {
  host          = "http://databricks.domain.com"
  basic_auth {
    username = var.user
    password = var.password
  }
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```
{{% /tab %}}
 {{% tab name="Profile" %}}
``` hcl
provider "databricks" {
  config_file = "~/.databrickscfg"
  profile     = "DEFAULT"
}

resource "databricks_scim_user" "my-user" {
  user_name     = "test-user@databricks.com"
  display_name  = "Test User"
}
```
 {{% /tab %}}
 {{% tab name="Azure SP Auth" %}}
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
  azure_auth = {
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
{{% /tab %}}
{{< /tabs >}}


{{% notice warning %}}
Please be aware that hard coding credentials is not something that is recommended.
It may be best if you store the credentials environment variables or use tfvars file.
{{% /notice %}}

## Authentication

There are currently two supported methods to authenticate into the Databricks platform to create resources.

* **API Token** 
* **Azure Service Principal Authentication** 

{{% notice note %}}
**Azure Service Principal Authentication** will only work on Azure Databricks where as the API Token
authentication will work on both **Azure** and **AWS**
{{% /notice %}}

### API Token

Databricks hostname for the workspace and api token can be provided here. This configuration is very similar to the 
Databricks CLI

``` hcl
provider "databricks" {
  host = "http://databricks.domain.com"
  token = "dapitokenhere"
}
```
{{% notice warning %}}
Please be aware that hard coding credentials is not something that is recommended.
It may be best if you store the credentials environment variables or use tfvars file.
{{% /notice %}}


### Azure Service Principal Auth

``` hcl
provider "databricks" {
  azure_auth = {
    managed_resource_group = "${azurerm_databricks_workspace.sri_test_workspace.managed_resource_group_name}"
    azure_region = "${azurerm_databricks_workspace.sri_test_workspace.location}"
    workspace_name = "${azurerm_databricks_workspace.sri_test_workspace.name}"
    resource_group = "${azurerm_databricks_workspace.sri_test_workspace.resource_group_name}"
    client_id = "${var.client_id}"
    client_secret = "${var.client_secret}"
    tenant_id = "${var.tenant_id}"
    subscription_id = "${var.subscription_id}"
  }
}
```

### Environment variables

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

## Provider Argument Reference

The following arguments are supported by the db provider block:

#### `host`
> This is the host of the Databricks workspace. This is will be a url that you use to login to your workspace. 
Alternatively you can provide this value as an environment variable `DATABRICKS_HOST`.

#### `token`:
> This is the api token to authenticate into the workspace. Alternatively you can provide this value as an 
environment variable `DATABRICKS_TOKEN`. 

#### `basic_auth`:
> #### **Usage**
>```hcl
>basic_auth = {
>    username = "user"
>    password = "mypass-123"
>}
>```
> {{%chevron default=`This is the authentication required to authenticate to the Databricks via basic auth through a user 
 that has access to the workspace. This is optional as you can use the api token based auth. 
The basic_auth block contains the following arguments:` display="true" %}}

* `username` - This is the username of the user that can log into the workspace. 
Alternatively you can provide this value as an environment variable `DATABRICKS_USERNAME`.

* `password` - This is the password of the user that can log into the workspace.
Alternatively you can provide this value as an environment variable `DATABRICKS_PASSWORD`.
{{% /chevron %}}

#### `config_file`:
> Location of the Databricks CLI credentials file, that is created, by `databricks configure --token` command. 
>By default, it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication 
>for docs. Config file credentials will only be used when host/token/basic_auth/azure_auth are not provided. 
>Alternatively you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to 
>`~/.databrickscfg`. 

#### `profile`:
> Connection profile specified within ~/.databrickscfg. Please check 
>https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation. This field defaults to 
>`DEFAULT`.

#### `azure_auth`:
> #### **Usage**
>```hcl
>azure_auth = {
>    azure_region = "centralus"
>    managed_resource_group = "my-databricks-managed-rg"
>    workspace_name = "test-managed-workspace"
>    resource_group = "1-test-rg"
>    client_id = var.client_id
>    client_secret = var.client_secret
>    tenant_id = var.tenant_id
>    subscription_id = var.subscription_id
>}
>```
> {{%chevron default=`This is the authentication required to authenticate to the Databricks via an azure service 
principal that has access to the workspace. This is optional as you can use the api token based auth. 
The azure_auth block contains the following arguments:` display="true" %}}

* `managed_resource_group` - This is the managed workgroup id when the Databricks workspace is provisioned. 
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP`.

* `azure_region` - This is the azure region in which your workspace is deployed.
Alternatively you can provide this value as an environment variable `AZURE_REGION`.

* `workspace_name` - This is the name of your Azure Databricks Workspace.
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_WORKSPACE_NAME`.

* `resource_group` - This is the resource group in which your Azure Databricks Workspace resides in.
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_RESOURCE_GROUP`.

* `subscription_id` - This is the Azure Subscription id in which your Azure Databricks Workspace resides in. 
Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID`.
                                                                              
* `client_secret` - This is the Azure Enterprise Application (Service principal) client secret. This service 
principal requires contributor access to your Azure Databricks deployment. Alternatively you can provide this 
value as an environment variable `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`.

* `client_id` - This is the Azure Enterprise Application (Service principal) client id. This service principal 
requires contributor access to your Azure Databricks deployment. Alternatively you can provide this value as an 
environment variable `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`.

* `tenant_id` - This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) 
resides in. Alternatively you can provide this value as an environment variable `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`.

Where there are multiple environment variable options, the `DATABRICKS_AZURE_*` environment variables takes precedence and the `ARM_*` environment variables provide a way to share authentication configuration when using the `databricks-terraform` provider alongside the `azurerm` provider.
{{% /chevron %}}