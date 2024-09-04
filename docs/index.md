---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: Terraform provider for the Databricks Lakehouse platform
---

# Databricks Provider

Use the Databricks Terraform provider to interact with almost all of [Databricks](http://databricks.com/) resources. If you're new to Databricks, please follow guide to create a workspace on [Azure](guides/azure-workspace.md), [AWS](guides/aws-workspace.md) or [GCP](guides/gcp-workspace.md) and then this [workspace management](guides/workspace-management.md) tutorial.  Take advantage of [Terraform Modules](https://www.terraform.io/docs/modules/index.html) to make your code simpler and reuse [existing modules for Databricks resources](https://registry.terraform.io/search/modules?namespace=databricks).   Changelog is available [on GitHub](https://github.com/databricks/terraform-provider-databricks/blob/main/CHANGELOG.md).

![Resources](https://raw.githubusercontent.com/databricks/terraform-provider-databricks/main/docs/resources.png)

Compute resources

* Deploy [databricks_cluster](resources/cluster.md) on selected [databricks_node_type](data-sources/node_type.md)
* Schedule automated [databricks_job](resources/job.md)
* Control cost and data access with [databricks_cluster_policy](resources/cluster_policy.md)
* Speedup job & cluster startup with [databricks_instance_pool](resources/instance_pool.md)
* Customize clusters with [databricks_global_init_script](resources/global_init_script.md)
* Manage few [databricks_notebook](resources/notebook.md), and even [list them](data-sources/notebook_paths.md)
* Manage [databricks_repo](resources/repo.md)

Storage

* Manage JAR, Wheel & Egg libraries through [databricks_dbfs_file](resources/dbfs_file.md)
* List entries on DBFS with [databricks_dbfs_file_paths](data-sources/dbfs_file_paths.md) data source
* Get contents of small files with [databricks_dbfs_file](data-sources/dbfs_file.md) data source
* Mount storage with [databricks_mount](resources/mount.md) resource

Security

* Organize [databricks_user](resources/user.md) into [databricks_group](resources/group.md) through [databricks_group_member](resources/group_member.md), also reading [metadata](data-sources/group.md)
* Create [databricks_service_principal](resources/service_principal.md) with [databricks_obo_token](resources/obo_token.md) to enable even more restricted access control.
* Create [databricks_service_principal](resources/service_principal.md) with [databricks_service_principal_secret](resources/service_principal_secret.md) to authenticate with the service principal OAuth tokens (Only for AWS deployments)
* Manage data access with [databricks_instance_profile](resources/instance_profile.md), which can be assigned through [databricks_group_instance_profile](resources/group_instance_profile.md) and [databricks_user_instance_profile](resources/user_instance_profile.md)
* Control which networks can access workspace with [databricks_ip_access_list](resources/ip_access_list.md)
* Generically manage [databricks_permissions](resources/permissions.md)
* Manage data object access control lists with [databricks_sql_permissions](resources/sql_permissions.md)
* Keep sensitive elements like passwords in [databricks_secret](resources/secret.md), grouped into [databricks_secret_scope](resources/secret_scope.md) and controlled by [databricks_secret_acl](resources/secret_acl.md)

[Databricks workspace on AWS](../docs/guides/aws-workspace.md)

* Create [workspaces](resources/mws_workspaces.md) in your [VPC](resources/mws_networks.md) with [DBFS](resources/mws_storage_configurations.md) using [cross-account IAM roles](resources/mws_credentials.md), having your notebooks encrypted with [CMK](resources/mws_customer_managed_keys.md).
* Use predefined AWS IAM Policy Templates: [databricks_aws_assume_role_policy](data-sources/aws_assume_role_policy.md), [databricks_aws_crossaccount_policy](data-sources/aws_crossaccount_policy.md), [databricks_aws_bucket_policy](data-sources/aws_bucket_policy.md)
* Configure billing and audit [databricks_mws_log_delivery](resources/mws_log_delivery.md)

Databricks SQL

* Create [databricks_sql_endpoint](resources/sql_endpoint.md) controlled by [databricks_permissions](resources/permissions.md).
* Manage [queries](resources/sql_query.md) and their [visualizations](resources/sql_visualization.md).
* Manage [dashboards](resources/sql_dashboard.md) and their [widgets](resources/sql_widget.md).
* Provide [global configuration for all SQL warehouses](resources/sql_global_config.md)

Machine Learning

* Create [models in Unity Catalog](resources/registered_model.md).
* Create [MLflow experiments](resources/mlflow_experiment.md).
* Create [models in the workspace model registry](resources/mlflow_model.md).
* Create [model serving endpoints](resources/model_serving.md).

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

  task {
    task_key = "task1"

    notebook_task {
      notebook_path = databricks_notebook.this.path
    }

    new_cluster {
      num_workers   = 1
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }
  }
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Switching from `databrickslabs` to `databricks` namespace

To make Databricks Terraform Provider generally available, we've moved it from [https://github.com/databrickslabs](https://github.com/databrickslabs) to [https://github.com/databricks](https://github.com/databricks). We've worked closely with the Terraform Registry team at Hashicorp to ensure a smooth migration. Existing terraform deployments continue to work as expected without any action from your side. We ask you to replace `databrickslabs/databricks` with `databricks/databricks` in all your `.tf` files.

You should have [`.terraform.lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl) file in your state directory that is checked into source control. terraform init will give you the following warning.

```text
Warning: Additional provider information from registry 

The remote registry returned warnings for registry.terraform.io/databrickslabs/databricks:
- For users on Terraform 0.13 or greater, this provider has moved to databricks/databricks. Please update your source in required_providers.
```

After you replace `databrickslabs/databricks` with `databricks/databricks` in the `required_providers` block, the warning will disappear. Do a global "search and replace" in `*.tf` files. Alternatively you can run `python3 -c "$(curl -Ls https://dbricks.co/updtfns)"` from the command-line, that would do all the boring work for you.

If you didn't check-in [`.terraform.lock.hcl`](https://www.terraform.io/language/files/dependency-lock#lock-file-location) to the source code version control, you may see `Failed to install provider` error. Please follow the simple steps described in the [troubleshooting guide](guides/troubleshooting.md).

## Troubleshooting

In case of the problems using Databricks Terraform provider follow the steps outlined in the [troubleshooting guide](guides/troubleshooting.md).

## Authentication

!> **Warning** Please be aware that hard coding any credentials in plain text is not something that is recommended. We strongly recommend using a Terraform backend that supports encryption. Please use [environment variables](#environment-variables), `~/.databrickscfg` file, encrypted `.tfvars` files or secret store of your choice (Hashicorp [Vault](https://www.vaultproject.io/), AWS [Secrets Manager](https://aws.amazon.com/secrets-manager/), AWS [Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html), Azure [Key Vault](https://azure.microsoft.com/en-us/services/key-vault/))

There are currently a number of supported methods to [authenticate](https://docs.databricks.com/dev-tools/api/latest/authentication.html) into the Databricks platform to create resources:

* [PAT Tokens](#authenticating-with-hostname-and-token)
* AWS, Azure and GCP via [Databricks-managed Service Principals](#authenticating-with-databricks-managed-service-principal)
* GCP via [Google Cloud CLI](#special-configurations-for-gcp)
* Azure Active Directory Tokens via [Azure CLI](#authenticating-with-azure-cli), [Azure-managed Service Principals](#authenticating-with-azure-managed-service-principal), or [Managed Service Identities](#authenticating-with-azure-msi)

### Authenticating with Databricks CLI credentials

If no configuration option is given, the provider will look up configured credentials in `~/.databrickscfg` file. It is created by the `databricks configure --token` command. Check [this page](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication)
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
  host  = "https://abc-cdef-ghi.cloud.databricks.com"
  token = "dapitokenhere"
}
```

### Authenticating with Databricks-managed Service Principal

You can use the `client_id` + `client_secret` attributes to authenticate with a Databricks-managed service principal at both the account and workspace levels in all supported clouds. The `client_id` is the `application_id` of the [Service Principal](resources/service_principal.md) and `client_secret` is its secret. You can generate the secret from Databricks Accounts Console (see [instruction](https://docs.databricks.com/dev-tools/authentication-oauth.html#step-2-create-an-oauth-secret-for-a-service-principal)) or by using the Terraform resource [databricks_service_principal_secret](resources/service_principal_secret.md).

``` hcl
provider "databricks" {
  host          = "https://abc-cdef-ghi.cloud.databricks.com"
  client_id     = var.client_id
  client_secret = var.client_secret
}
```

To create resources at both the account and workspace levels, you can create two providers as shown below

``` hcl
provider "databricks" {
  alias         = "accounts"
  host          = "https://accounts.cloud.databricks.com"
  client_id     = var.client_id
  client_secret = var.client_secret
  account_id    = "00000000-0000-0000-0000-000000000000"
}

provider "databricks" {
  alias         = "workspace"
  host          = var.workspace_host
  client_id     = var.client_id
  client_secret = var.client_secret
}
```

Next, you can specify the corresponding provider when creating the resource. For example, you can use the workspace provider to create a workspace group

``` hcl
resource "databricks_group" "cluster_admin" {
  provider                   = databricks.workspace
  display_name               = "cluster_admin"
  allow_cluster_create       = true
  allow_instance_pool_create = false
}
```

* `client_id` - The `application_id` of the [Service Principal](resources/service_principal.md). Alternatively, you can provide this value as an environment variable `DATABRICKS_CLIENT_ID`.
* `client_secret` - Secret of the service principal. Alternatively, you can provide this value as an environment variable `DATABRICKS_CLIENT_SECRET`.

## Argument Reference

-> **Note** If you experience technical difficulties with rolling out resources in this example, please make sure that [environment variables](#environment-variables) don't [conflict with other](#empty-provider-block) provider block attributes. When in doubt, please run `TF_LOG=DEBUG terraform apply` to enable [debug mode](https://www.terraform.io/docs/internals/debugging.html) through the [`TF_LOG`](https://www.terraform.io/docs/cli/config/environment-variables.html#tf_log) environment variable. Look specifically for `Explicit and implicit attributes` lines, that should indicate authentication attributes used.

The provider block supports the following arguments:

* `host` - (optional) This is the host of the Databricks workspace. It is a URL that you use to login to your workspace.
Alternatively, you can provide this value as an environment variable `DATABRICKS_HOST`.
* `token` - (optional) This is the API token to authenticate into the workspace. Alternatively, you can provide this value as an environment variable `DATABRICKS_TOKEN`.
* `config_file` - (optional) Location of the Databricks CLI credentials file created by `databricks configure --token` command (~/.databrickscfg by default). Check [Databricks CLI documentation](https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication) for more details. The provider uses configuration file credentials when you don't specify host/token/azure attributes. Alternatively, you can provide this value as an environment variable `DATABRICKS_CONFIG_FILE`. This field defaults to `~/.databrickscfg`.
* `profile` - (optional) Connection profile specified within ~/.databrickscfg. Please check [connection profiles section](https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles) for more details. This field defaults to
`DEFAULT`.
* `account_id` - (optional for workspace-level operations, but required for account-level) Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/). Alternatively, you can provide this value as an environment variable `DATABRICKS_ACCOUNT_ID`. Only has effect when `host = "https://accounts.cloud.databricks.com/"`, and is currently used to provision account admins via [databricks_user](resources/user.md). In the future releases of the provider this property will also be used specify account for `databricks_mws_*` resources as well.
* `auth_type` - (optional) enforce specific auth type to be used in very rare cases, where a single Terraform state manages Databricks workspaces on more than one cloud and `more than one authorization method configured` error is a false positive. Valid values are `pat`, `basic`, `oauth-m2m`, `azure-client-secret`, `azure-msi`, `azure-cli`, `google-credentials`, and `google-id`.

## Special configurations for Azure

The below Azure authentication options are supported at both the account and workspace levels. The provider works with [Azure CLI authentication](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest) to facilitate local development workflows, though for automated scenarios, managed identity or service principal auth is recommended (and specification of `azure_use_msi`, `azure_client_id`, `azure_client_secret` and `azure_tenant_id` parameters).

### Authenticating with Azure MSI

Since v0.3.8, it's possible to leverage [Azure Managed Service Identity](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/managed_service_identity) authentication, which is using the same environment variables as `azurerm` provider. Both `SystemAssigned` and `UserAssigned` identities work, as long as they have `Contributor` role on subscription level and created the workspace resource, or directly added to workspace through [databricks_service_principal](resources/service_principal.md).

```hcl
provider "databricks" {
  host                        = data.azurerm_databricks_workspace.this.workspace_url
  azure_workspace_resource_id = azurerm_databricks_workspace.this.id

  # ARM_USE_MSI environment variable is recommended
  azure_use_msi = true
}
```

### Authenticating with Azure CLI

It's possible to use [Azure CLI](https://docs.microsoft.com/cli/azure/) authentication, where the provider would rely on access token cached by `az login` command so that local development scenarios are possible. Technically, the provider will call `az account get-access-token` each time before an access token is about to expire.

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_databricks_workspace" "this" {
  location            = "centralus"
  name                = "my-workspace-name"
  resource_group_name = var.resource_group
  sku                 = "premium"
}

provider "databricks" {
  host = azurerm_databricks_workspace.this.workspace_url
}

resource "databricks_user" "my-user" {
  user_name    = "test-user@databricks.com"
  display_name = "Test User"
}
```

### Authenticating with Azure-managed Service Principal

```hcl
provider "azurerm" {
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
  subscription_id = var.subscription_id
}

resource "azurerm_databricks_workspace" "this" {
  location            = "centralus"
  name                = "my-workspace-name"
  resource_group_name = var.resource_group
  sku                 = "premium"
}

provider "databricks" {
  host                        = azurerm_databricks_workspace.this.workspace_url
  azure_workspace_resource_id = azurerm_databricks_workspace.this.id
  azure_client_id             = var.client_id
  azure_client_secret         = var.client_secret
  azure_tenant_id             = var.tenant_id
}

resource "databricks_user" "my-user" {
  user_name = "test-user@databricks.com"
}
```

* `azure_workspace_resource_id` - (optional) `id` attribute of [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace) resource. Combination of subscription id, resource group name, and workspace name. Required with `auzre_use_msi` or `azure_client_secret`.
* `azure_client_secret` - (optional) This is the Azure Enterprise Application (Service principal) client secret. This service principal requires contributor access to your Azure Databricks deployment. Alternatively, you can provide this value as an environment variable `ARM_CLIENT_SECRET`.
* `azure_client_id` - (optional) This is the Azure Enterprise Application (Service principal) client id. This service principal requires contributor access to your Azure Databricks deployment. Alternatively, you can provide this value as an environment variable `ARM_CLIENT_ID`.
* `azure_tenant_id` - (optional) This is the Azure Active Directory Tenant id in which the Enterprise Application (Service Principal)
resides. Alternatively, you can provide this value as an environment variable `ARM_TENANT_ID`.
* `azure_environment` - (optional) This is the Azure Environment which defaults to the `public` cloud. Other options are `german`, `china` and `usgovernment`. Alternatively, you can provide this value as an environment variable `ARM_ENVIRONMENT`.
* `azure_use_msi` - (optional) Use [Azure Managed Service Identity](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/managed_service_identity) authentication. Alternatively, you can provide this value as an environment variable `ARM_USE_MSI`.

There are `ARM_*` environment variables provide a way to share authentication configuration using the `databricks` provider alongside the [`azurerm` provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest).

When a workspace is created using a service principal account, that service principal account is automatically added to the workspace as a member of the admins group. To add a new service principal account to an existing workspace, create a [databricks_service_principal](resources/service_principal.md).

## Special configurations for GCP

The provider works with [Google Cloud CLI authentication](https://cloud.google.com/sdk/docs/authorizing) to facilitate local development workflows. For automated scenarios, a service principal auth is necessary using `google_service_account` parameter with [impersonation](https://cloud.google.com/docs/authentication#service-accounts) and Application Default Credentials. Alternatively, you could provide the service account key directly by passing it to `google_credentials` parameter (or `GOOGLE_CREDENTIALS` environment variable)

## Special configuration for Unity Catalog

Except for metastore, metastore assignment and storage credential objects, Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future.

If you are configuring a new Databricks account for the first time, please create at least one workspace with an identity (user or service principal) that you intend to use for Unity Catalog rollout. You can then configure the provider using that identity and workspace to provision the required Unity Catalog resources.

## Special considerations for Unity Catalog Resources

When performing a single Terraform apply to update both the owner and other fields for Unity Catalog resources, the process first updates the owner, followed by the other fields using the new owner's permissions. If your principal is not the owner (specifically, the newly updated owner), you will not have the authority to modify those fields. In cases where you wish to change the owner to another individual and also update other fields, we recommend initially updating the fields using your principal, which should have owner permissions, and then updating the owner in a separate step.

## Miscellaneous configuration parameters

!> **Warning** Combination of `debug_headers` and `debug_truncate_bytes` results in dumping of sensitive information to logs. Use it for troubleshooting purposes only.

This section covers configuration parameters not related to authentication. They could be used when debugging problems, or do an additional tuning of provider's behaviour:

* `http_timeout_seconds` - the amount of time Terraform waits for a response from Databricks REST API. Default is *60*.
* `rate_limit` - defines maximum number of requests per second made to Databricks REST API by Terraform. Default is *15*.
* `debug_truncate_bytes` - Applicable only when `TF_LOG=DEBUG` is set. Truncate JSON fields in HTTP requests and responses above this limit. Default is *96*.
* `debug_headers` - Applicable only when `TF_LOG=DEBUG` is set. Debug HTTP headers of requests made by the provider. Default is *false*. We recommend turning this flag on only under exceptional circumstances, when troubleshooting authentication issues. Turning this flag on will log first `debug_truncate_bytes` of any HTTP header value in cleartext.
* `skip_verify` - skips SSL certificate verification for HTTP calls. *Use at your own risk.* Default is *false* (don't skip verification).

## Environment variables

The following configuration attributes can be passed via environment variables:

|                      Argument | Environment variable              |
| ----------------------------: | --------------------------------- |
|                   `auth_type` | `DATABRICKS_AUTH_TYPE`            |
|                        `host` | `DATABRICKS_HOST`                 |
|                       `token` | `DATABRICKS_TOKEN`                |
|                  `account_id` | `DATABRICKS_ACCOUNT_ID`           |
|                 `config_file` | `DATABRICKS_CONFIG_FILE`          |
|                     `profile` | `DATABRICKS_CONFIG_PROFILE`       |
|                   `client_id` | `DATABRICKS_CLIENT_ID`            |
|               `client_secret` | `DATABRICKS_CLIENT_SECRET`        |
|         `azure_client_secret` | `ARM_CLIENT_SECRET`               |
|             `azure_client_id` | `ARM_CLIENT_ID`                   |
|             `azure_tenant_id` | `ARM_TENANT_ID`                   |
| `azure_workspace_resource_id` | `DATABRICKS_AZURE_RESOURCE_ID`    |
|               `azure_use_msi` | `ARM_USE_MSI`                     |
|           `azure_environment` | `ARM_ENVIRONMENT`                 |
|          `google_credentials` | `GOOGLE_CREDENTIALS`              |
|      `google_service_account` | `GOOGLE_SERVICE_ACCOUNT`          |
|        `debug_truncate_bytes` | `DATABRICKS_DEBUG_TRUNCATE_BYTES` |
|               `debug_headers` | `DATABRICKS_DEBUG_HEADERS`        |
|               `rate_limit`    | `DATABRICKS_RATE_LIMIT`           |

## Empty provider block

For example, with the following zero-argument configuration:

```hcl
provider "databricks" {}
```

1. Provider will check all the supported environment variables and set values of relevant arguments.
2. In case any conflicting arguments are present, the plan will end with an error.
3. Will check for the presence of `host` + `token` pair, continue trying otherwise.
4. Will check for Azure workspace ID, `azure_client_secret` + `azure_client_id` + `azure_tenant_id` presence, continue trying otherwise.
5. Will check for availability of Azure MSI, if enabled via `azure_use_msi`, continue trying otherwise.
6. Will check for Azure workspace ID presence, and if `AZ CLI` returns an access token, continue trying otherwise.
7. Will check for the `~/.databrickscfg` file in the home directory, will fail otherwise.
8. Will check for `profile` presence and try picking from that file will fail otherwise.

Please check [Default Authentication Flow](https://github.com/databricks/databricks-sdk-go#default-authentication-flow) from [Databricks SDK for Go](https://docs.databricks.com/dev-tools/sdk-go.html) in case you need more details.
