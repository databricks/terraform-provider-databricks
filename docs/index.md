---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: |-
  Terraform provider databricks.
---

# Databricks Provider

Use the Databricks Terraform provider to interact with almost all of [Databricks](http://databricks.com/) resources. If you're new to Databricks, please follow guide to create a workspace on [Azure](guides/azure-workspace.md) or [AWS](guides/aws-workspace.md) and then this [workspace management](guides/workspace-management.md) tutorial. If you're migrating from version *0.2.x*, please follow [this guide](guides/migration-0.3.x.md). Changelog is available [on GitHub](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/CHANGELOG.md).

![Resources](https://github.com/databrickslabs/terraform-provider-databricks/raw/master/docs/resources.png)

Compute resources
* Deploy [databricks_cluster](resources/cluster.md) on selected [databricks_node_type](data-sources/node_type.md)
* Schedule automated [databricks_job](resources/job.md)
* Control cost and data access with [databricks_cluster_policy](resources/cluster_policy.md)
* Speedup job & cluster startup with [databricks_instance_pool](resources/instance_pool.md)
* Customize clusters with [databricks_global_init_script](resources/global_init_script.md)
* Manage few [databricks_notebook](resources/notebook.md), and even [list them](data-sources/notebook_paths.md)

Storage
* Manage JAR, Wheel & Egg libraries through [databricks_dbfs_file](resources/dbfs_file.md)
* List entries on DBFS with [databricks_dbfs_file_paths](data-sources/dbfs_file_paths.md) data source
* Get contents of small files with [databricks_dbfs_file](data-sources/dbfs_file.md) data source
* Mount your AWS storage using [databricks_aws_s3_mount](resources/aws_s3_mount.md)
* Mount your Azure storage using [databricks_azure_adls_gen1_mount](resources/azure_adls_gen1_mount.md), [databricks_azure_adls_gen2_mount](resources/azure_adls_gen2_mount.md), [databricks_azure_blob_mount](resources/azure_blob_mount.md)

Security
* Organize [databricks_user](resources/user.md) into [databricks_group](resources/group.md) through [databricks_group_member](resources/group_member.md), also reading [metadata](data-sources/group.md)
* Manage data access with [databricks_instance_profile](resources/instance_profile.md), which can be assigned through [databricks_group_instance_profile](resources/group_instance_profile.md) and [databricks_user_instance_profile](resources/user_instance_profile.md)
* Control which networks can access workspace with [databricks_ip_access_list](resources/ip_access_list.md)
* Generically manage [databricks_permissions](resources/permissions.md)
* Manage data object access control lists with [databricks_sql_permissions](resources/sql_permissions.md)
* Keep sensitive elements like passwords in [databricks_secret](resources/secret.md), grouped into [databricks_secret_scope](resources/secret_scope.md) and controlled by [databricks_secret_acl](resources/secret_acl.md)


[E2 Architecture](../docs/guides/aws-workspace.md)
* Create [workspaces](resources/mws_workspaces.md) in your [VPC](resources/mws_networks.md) with [DBFS](resources/mws_storage_configurations.md) using [cross-account IAM roles](resources/mws_credentials.md), having your notebooks encrypted with [CMK](resources/mws_customer_managed_keys.md).
* Use predefined AWS IAM Policy Templates: [databricks_aws_assume_role_policy](data-sources/aws_assume_role_policy.md), [databricks_aws_crossaccount_policy](data-sources/aws_crossaccount_policy.md), [databricks_aws_bucket_policy](data-sources/aws_bucket_policy.md)
* Configure billing and audit [databricks_mws_log_delivery](resources/mws_log_delivery.md)

SQL Analytics
* Create [databricks_sql_endpoint](resources/sql_endpoint.md) controlled by [databricks_permissions](resources/permissions.md).
* Manage [queries](resources/sql_query.md) and their [visualizations](resources/sql_visualization.md).
* Manage [dashboards](resources/sql_dashboard.md) and their [widgets](resources/sql_widget.md).

## Example Usage

```hcl
provider "databricks" {
}

data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/Terraform"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
    EOT
  )
}

resource "databricks_job" "this" {
  name = "Terraform Demo (${data.databricks_current_user.me.alphanumeric})"

  new_cluster {
    num_workers   = 1
    spark_version = data.databricks_spark_version.latest.id
    node_type_id  = data.databricks_node_type.smallest.id
  }

  notebook_task {
    notebook_path = databricks_notebook.this.path
  }

  email_notifications {}
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Authentication

!> **Warning** Please be aware that hard coding any credentials in plain text is not something that is recommended. We strongly recommend using a Terraform backend that supports encryption. Please use [environment variables](#environment-variables), `~/.databrickscfg` file, encrypted `.tfvars` files or secret store of your choice (Hashicorp [Vault](https://www.vaultproject.io/), AWS [Secrets Manager](https://aws.amazon.com/secrets-manager/), AWS [Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html), Azure [Key Vault](https://azure.microsoft.com/en-us/services/key-vault/))


There are currently three supported methods to [authenticate](https://docs.databricks.com/dev-tools/api/latest/authentication.html) into the Databricks platform to create resources:

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

You can use the `username` + `password` attributes to authenticate provider for E2 workspace setup. Respective `DATABRICKS_USERNAME` and `DATABRICKS_PASSWORD` environment variables are applicable as well.

``` hcl
provider "databricks" {
  host = "http://accounts.cloud.databricks.com"
  username = var.user
  password = var.password
}
```

## Argument Reference

-> **Note** If you experience technical difficulties with rolling out resources in this example, please make sure that [environment variables](#environment-variables) don't [conflict with other](#empty-provider-block) provider block attributes. When in doubt, please run `TF_LOG=DEBUG terraform apply` to enable [debug mode](https://www.terraform.io/docs/internals/debugging.html) through the [`TF_LOG`](https://www.terraform.io/docs/cli/config/environment-variables.html#tf_log) environment variable. Look specifically for `Explicit and implicit attributes` lines, that should indicate authentication attributes used.

The provider block supports the following arguments:

* `host` - (optional) This is the host of the Databricks workspace. It is a URL that you use to login to your workspace. 
Alternatively, you can provide this value as an environment variable `DATABRICKS_HOST`.
* `token` - (optional) This is the API token to authenticate into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_TOKEN`. 
* `username` - (optional) This is the username of the user that can log into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_USERNAME`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `password` - (optional) This is the user's password that can log into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_PASSWORD`. Recommended only for [creating workspaces in AWS](resources/mws_workspaces.md).
* `config_file` - (optional) Location of the Databricks CLI credentials file created by `databricks configure --token` command (~/.databrickscfg by default). Check [Databricks CLI documentation](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) for more details. The provider uses configuration file credentials when you don't specify host/token/username/password/azure attributes. Alternatively, you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to `~/.databrickscfg`. 
* `profile` - (optional) Connection profile specified within ~/.databrickscfg. Please check [connection profiles section](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) for more details. This field defaults to 
`DEFAULT`.

## Special configurations for Azure

To work with Azure Databricks workspace, the provider must know its `azure_workspace_resource_id` (or construct it from `azure_subscription_id`, `azure_resource_group` and `azure_workspace_name`). The provider works with [Azure CLI authentication](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest) to facilitate local development workflows, though for automated scenarios a service principal auth is necessary (and specification of `azure_client_id`, `azure_client_secret` and `azure_tenant_id` parameters).

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

resource "databricks_user" "my-user" {
  user_name = "test-user@databricks.com"
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

resource "databricks_user" "my-user" {
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
* `azure_environment` - (optional) This is the Azure Environment which defaults to the `public` cloud. Other options are `german`, `china` and `usgovernment`. Alternatively, you can provide this value as an environment variable `ARM_ENVIRONMENT`.
* `pat_token_duration_seconds` - The current implementation of the azure auth via sp requires the provider to create a temporary personal access token within Databricks. The current AAD implementation does not cover all the APIs for Authentication. This field determines the duration in which that temporary PAT token is alive. It is measured in seconds and will default to `3600` seconds. 

There are multiple environment variable options, the `DATABRICKS_AZURE_*` environment variables take precedence, and the `ARM_*` environment variables provide a way to share authentication configuration using the `databricks` provider alongside the `azurerm` provider.

## Miscellaneous configuration parameters

This section covers configuration parameters not related to authentication.  They could be used when debugging problems, or do an additional tuning of provider's behaviour:

* `rate_limit` - defines maximum number of requests per second made to Databricks REST API by Terraform. Default is *15*.
* `debug_truncate_bytes` - Applicable only when `TF_LOG=DEBUG` is set. Truncate JSON fields in HTTP requests and responses above this limit. Default is *96*.
* `debug_headers` - Applicable only when `TF_LOG=DEBUG` is set. Debug HTTP headers of requests made by the provider. Default is *false*. We recommend to turn this flag on only under exceptional circumstances, when troubleshooting authentication issues. Turning this flag on will log first `debug_truncate_bytes` of any HTTP header value in cleartext.
* `skip_verify` - skips SSL certificate verification for HTTP calls. *Use at your own risk.* Default is *false* (don't skip verification).


## Environment variables

The following configuration attributes can be passed via environment variables:

|                      Argument | Environment variable                                        |
| ----------------------------: | ----------------------------------------------------------- |
|                        `host` | `DATABRICKS_HOST`                                           |
|                       `token` | `DATABRICKS_TOKEN`                                          |
|                    `username` | `DATABRICKS_USERNAME`                                       |
|                    `password` | `DATABRICKS_PASSWORD`                                       |
|                 `config_file` | `DATABRICKS_CONFIG_FILE`                                    |
|                     `profile` | `DATABRICKS_CONFIG_PROFILE`                                 |
| `azure_workspace_resource_id` | `DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID`                    |
|        `azure_workspace_name` | `DATABRICKS_AZURE_WORKSPACE_NAME`                           |
|        `azure_resource_group` | `DATABRICKS_AZURE_RESOURCE_GROUP`                           |
|       `azure_subscription_id` | `DATABRICKS_AZURE_SUBSCRIPTION_ID` or `ARM_SUBSCRIPTION_ID` |
|         `azure_client_secret` | `DATABRICKS_AZURE_CLIENT_SECRET` or `ARM_CLIENT_SECRET`     |
|             `azure_client_id` | `DATABRICKS_AZURE_CLIENT_ID` or `ARM_CLIENT_ID`             |
|             `azure_tenant_id` | `DATABRICKS_AZURE_TENANT_ID` or `ARM_TENANT_ID`             |
|           `azure_environment` | `ARM_ENVIRONMENT`                                           |
|        `debug_truncate_bytes` | `DATABRICKS_DEBUG_TRUNCATE_BYTES`                           |
|               `debug_headers` | `DATABRICKS_DEBUG_HEADERS`                                  |
|               `rate_limit`    | `DATABRICKS_RATE_LIMIT`                                     |


## Empty provider block

For example, with the following zero-argument configuration:

``` hcl 
provider "databricks" {}
```

1. Provider will check all the supported environment variables and set values of relevant arguments.
2. In case any conflicting arguments are present, the plan will end with an error.
3. Will check for the presence of `host` + `token` pair, continue trying otherwise.
4. Will check for `host` + `username` + `password` presence, continue trying otherwise.
5. Will check for Azure workspace ID, `azure_client_secret` + `azure_client_id` + `azure_tenant_id` presence, continue trying otherwise.
6. Will check for Azure workspace ID presence, and if `AZ CLI` returns an access token, continue trying otherwise.
7. Will check for the `~/.databrickscfg` file in the home directory, will fail otherwise.
8. Will check for `profile` presence and try picking from that file will fail otherwise.
9. Will check for `host` and `token` or `username`+`password` combination, will fail if nothing of these exist.

## Data resources and Authentication is not configured errors

*In Terraform 0.13 and later*, data resources have the same dependency resolution behavior [as defined for managed resources](https://www.terraform.io/docs/language/resources/behavior.html#resource-dependencies). Most data resources make an API call to a workspace. If a workspace doesn't exist yet, `authentication is not configured for provider` error is raised. To work around this issue and guarantee a proper lazy authentication with data resources, you should add `depends_on = [azurerm_databricks_workspace.this]` or `depends_on = [databricks_mws_workspaces.this]` to the body. This issue doesn't occur if workspace is created *in one module* and resources [within the workspace](guides/workspace-management.md) are created *in another*. We do not recommend using Terraform 0.12 and earlier, if your usage involves data resources.

## Multiple Provider Configurations

 The most common reason for technical difficulties might be related to missing `alias` attribute in `provider "databricks" {}` blocks or `provider` attribute in `resource "databricks_..." {}` blocks, when using multiple provider configurations. Please make sure to read [`alias`: Multiple Provider Configurations](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) documentation article. 

## Error while installing: registry does not have a provider

```
Error while installing hashicorp/databricks: provider registry
registry.terraform.io does not have a provider named
registry.terraform.io/hashicorp/databricks
```

If you notice below error, it might be due to the fact that [required_providers](https://www.terraform.io/docs/language/providers/requirements.html#requiring-providers) block is not defined in *every module*, that uses Databricks Terraform Provider. Create `versions.tf` file with the following contents:

```hcl
# versions.tf
terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
      version = "0.3.4"
    }
  }
}
```

... and copy the file in every module in your codebase. Our recommendation is to skip `version` field for `versions.tf` file on module level, and keep it only on environment level.

```
├── environments
│   ├── sandbox
│   │   ├── README.md
│   │   ├── main.tf
│   │   └── versions.tf
│   └── production
│       ├── README.md
│       ├── main.tf
│       └── versions.tf
└── modules
    ├── first-module
    │   ├── ...
    │   └── versions.tf
    └── second-module
        ├── ...
        └── versions.tf
```


## Project Support

**Important:** Projects in the `databrickslabs` GitHub account, including the Databricks Terraform Provider, are not formally supported by Databricks. They are maintained by Databricks Field teams and provided as-is. There is no service level agreement (SLA). Databricks makes no guarantees of any kind. If you discover an issue with the provider, please file a GitHub Issue on the repo, and it will be reviewed by project maintainers as time permits.
