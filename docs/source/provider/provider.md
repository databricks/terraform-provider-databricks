# Databricks Provider

The Databricks provider is what is used to interact with the Databricks resources. This needs to be configured so that 
terraform can provision resources in your Databricks workspace on your behalf.  

## Example Usage

.. code-block:: tf
    
    provider "databricks" {
      host = "http://databricks.domain.com"
      token = "dapitokenhere"
    }
    
    resource "databricks_scim_user" "my-user" {
      user_name = "test-user@databricks.com"
      display_name = "Test User"
    }

.. WARNING:: Please be aware that hard coding credentials is not something that is recommended.
             It may be best if you store the credentials environment variables or use tfvars file.

## Authentication

There are currently two supported methods to authenticate into the Databricks platform to create resources.

* **API Token** 
* **Azure Service Principal Authentication** 

.. Note:: **Azure Service Principal Authentication** will only work on Azure Databricks where as the API Token
          authentication will work on both **Azure** and **AWS**

### API Token

Databricks hostname for the workspace and api token can be provided here. This configuration is very similar to the 
Databricks CLI

.. code-block:: tf

    provider "databricks" {
      host = "http://databricks.domain.com"
      token = "dapitokenhere"
    }

.. WARNING:: Please be aware that hard coding credentials is not something that is recommended.
             It may be best if you store the credentials environment variables or use tfvars file.


### Azure Service Principal Auth

.. code-block:: tf

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

### Environment variables

The following variables can be passed via environment variables:

* `host` → `HOST`
* `token` → `TOKEN`
* `subscription_id` → `ARM_SUBSCRIPTION_ID`
* `client_secret` → `ARM_CLIENT_SECRET`
* `client_id` → `ARM_CLIENT_ID`
* `tenant_id` → `ARM_TENANT_ID`

For example you can have the following provider definition:

.. code-block:: tf

    provider "databricks" {}

Then run the following code and the following environment variables will be injected into the provider.    

.. code-block:: sh

    $ export HOST="http://databricks.domain.com"
    $ export TOKEN="dapitokenhere"
    $ terraform plan

## Provider Argument Reference

The following arguments are supported by the db provider block:

.. _provider-host:
* :ref:`host <provider-host>` - This is the host of the Databricks workspace. This is will be a url that you use to login to your workspace. 
Alternatively you can provide this value as an environment variable `HOST`.

.. _provider-token:
* :ref:`token <provider-token>` - This is the api token to authenticate into the workspace. Alternatively you can provide this value as an 
environment variable `TOKEN`. 

.. _provider-azure-auth:
* :ref:`azure_auth <provider-azure-auth>` - This is the authentication required to authenticate to the Databricks via an azure service principal 
that has access to the workspace. This is optional as you can use the api token based auth. The `azure_auth` block 
contains the following arguments:
    
    * `managed_resource_group` - This is the managed workgroup id when the Databricks workspace is provisioned
    
    * `azure_region` - This is the azure region in which your workspace is deployed.
    
    * `workspace_name` - This is the name of your Azure Databricks Workspace.
    
    * `resource_group` - This is the resource group in which your Azure Databricks Workspace resides in.
    
    * `subscription_id` - This is the Azure Subscription id in which your Azure Databricks Workspace resides in. 
    Alternatively you can provide this value as an environment variable `ARM_SUBSCRIPTION_ID`.
                                                                                  
    * `client_secret` - This is the Azure Enterprise Application (Service principal) client secret. This service 
    principal requires contributor access to your Azure Databricks deployment. Alternatively you can provide this 
    value as an environment variable `ARM_CLIENT_SECRET`.  
    
    * `client_id` - This is the Azure Enterprise Application (Service principal) client id. This service principal 
    requires contributor access to your Azure Databricks deployment. Alternatively you can provide this value as an 
    environment variable `ARM_CLIENT_ID`.  
    
    * `tenant_id` - This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal) 
    resides in. Alternatively you can provide this value as an environment variable `ARM_TENANT_ID`.  
