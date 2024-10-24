---
page_title: "Troubleshooting Guide"
---

# How to troubleshoot your problem

If you have problems with code that uses Databricks Terraform provider, follow these steps to solve them:

* Check symptoms and solutions in the [Typical problems](#typical-problems) section below.
* Upgrade the provider to the latest version. The bug might have already been fixed.
* In case of authentication problems, see the [Data resources and Authentication is not configured errors](#data-resources-and-authentication-is-not-configured-errors) below.
* Collect debug information using the following command:

```sh
TF_LOG=DEBUG DATABRICKS_DEBUG_TRUNCATE_BYTES=250000 terraform apply -no-color 2>&1 |tee tf-debug.log
```

* Open a [new GitHub issue](https://github.com/databricks/terraform-provider-databricks/issues/new/choose) providing all information described in the issue template - debug logs, your Terraform code, Terraform & plugin versions, etc.

## Plugin Framework Migration Problems
The following resources and data sources have been migrated from sdkv2 to plugin framework。 If you encounter any problem with those, you can fallback to sdkv2 by setting the `USE_SDK_V2_RESOURCES` and `USE_SDK_V2_DATA_SOURCES` environment variables.

Example: `export USE_SDK_V2_RESOURCES="databricks_library,databricks_quality_monitor"`

### Resources migrated
  - databricks_quality_monitor
  - databricks_library
### Data sources migrated
  - databricks_volumes


## Typical problems

### Data resources and Authentication is not configured errors

*In Terraform 0.13 and later*, data resources have the same dependency resolution behavior [as defined for managed resources](https://www.terraform.io/docs/language/resources/behavior.html#resource-dependencies). Most data resources make an API call to a workspace. If a workspace doesn't exist yet, `default auth: cannot configure default credentials` error is raised. To work around this issue and guarantee a proper lazy authentication with data resources, you should add `depends_on = [azurerm_databricks_workspace.this]` or `depends_on = [databricks_mws_workspaces.this]` to the body. This issue doesn't occur if a workspace is created *in one module* and resources [within the workspace](guides/workspace-management.md) are created *in another*. We do not recommend using Terraform 0.12 and earlier if your usage involves data resources.

### Multiple Provider Configurations

The most common reason for technical difficulties might be related to missing `alias` attribute in `provider "databricks" {}` blocks or `provider` attribute in `resource "databricks_..." {}` blocks when using multiple provider configurations. Please make sure to read [`alias`: Multiple Provider Configurations](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) documentation article.

### Error while installing: registry does not have a provider

```sh
Error while installing hashicorp/databricks: provider registry
registry.terraform.io does not have a provider named
registry.terraform.io/hashicorp/databricks
```

If you notice the below error, it might be because [required_providers](https://www.terraform.io/docs/language/providers/requirements.html#requiring-providers) block is not defined in *every module*that uses Databricks Terraform Provider. Create `versions.tf` file with the following contents:

```hcl
# versions.tf
terraform {
  required_providers {
    databricks = {
      source  = "databricks/databricks"
      version = "1.0.1"
    }
  }
}
```

... and copy the file in every module in your codebase. We recommend skipping the `version` field for `versions.tf` file on the module level and keeping it only on the environment level.

```
├── environments
│   ├── sandbox
│   │   ├── README.md
│   │   ├── main.tf
│   │   └── versions.tf
│   └── production
│    ├── README.md
│    ├── main.tf
│    └── versions.tf
└── modules
 ├── first-module
 │   ├── ...
 │   └── versions.tf
 └── second-module
     ├── ...
     └── versions.tf
```

### Error: Failed to install provider

Running the `terraform init` command, you may see `Failed to install provider` error if you didn't check in [`.terraform.lock.hcl`](https://www.terraform.io/language/files/dependency-lock#lock-file-location) to the source code version control:

```sh
Error: Failed to install provider

Error while installing databricks/databricks: v1.0.0: checksum list has no SHA-256 hash for "https://github.com/databricks/terraform-provider-databricks/releases/download/v1.0.0/terraform-provider-databricks_1.0.0_darwin_amd64.zip"
```

You can fix it by following three simple steps:

* Replace `databrickslabs/databricks` with `databricks/databricks` in all your `.tf` files with the `python3 -c "$(curl -Ls https://dbricks.co/updtfns)"` command.
* Run the `terraform state replace-provider databrickslabs/databricks databricks/databricks` command and approve the changes. See [Terraform CLI](https://www.terraform.io/cli/commands/state/replace-provider) docs for more information.
* Run `terraform init` to verify everything is working.

The terraform apply command should work as expected now.

Alternatively, you can find the hashes of the last 30 provider versions in [`.terraform.lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl). As a temporary measure, you can lock on a prior version by following the following steps:

* Copy [`versions-lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl) to the root folder of your terraform project.
* Rename to `terraform.lock.hcl`
* Run `terraform init` and verify the provider is installed.
* Commit the new `.terraform.lock.hcl` file to your source code repository.

### Error: Failed to query available provider packages

See the same steps as in [Error: Failed to install provider](#error-failed-to-install-provider).

### Error: Deployment name cannot be used until a deployment name prefix is defined

You can get this error during provisioning of the Databricks workspace.  It arises when you're trying to set `deployment_name` with no deployment prefix on the Databricks side (you can't set it yourself).  The problem could be solved by one of the following methods:

1. Contact your Databricks representatives, like Solutions Architect, Customer Success Engineer, Account Executive, or Partner Solutions Architect, to set a deployment prefix for your account.

1. Comment out the `deployment_name` parameter to create a workspace with the default URL: `dbc-XXXXXX.cloud.databricks.com`.

### Error: 'strconv.ParseInt parsing "...." value out of range' or "Attribute must be a whole number, got N.NNNNe+XX"

This kind of error happens when the 32-bit version of Databricks Terraform provider is used, usually on Microsoft Windows.  To fix the issue, you need to switch to use of the 64-bit versions of Terraform and Databricks Terraform provider.

### Error: cannot create xxxx: HTTP method POST is not supported by this URL

This error may appear when creating Databricks users/groups/service principals on Databricks account level when no `account_id` is specified in the provider's configuration.  Make sure that `account_id` is set and has a correct value.

### Error: oauth-m2m: oidc: parse .well-known: invalid character '<' looking for beginning of value

This problem is similar to the previous item.  Ensure that `account_id` is specified in the provider configuration and it has a correct value.

### Error: cannot create ...: invalid character '<' looking for beginning of value

This error may appear when creating workspace-level objects, but the provider is configured to account-level.

### Error: cannot ...: unexpected error handling request: invalid character '<' for beginning of value

If you see the following HTTP request when running Terraform in the debug mode:

```
GET /login.html?error=private-link-validation-error:NNNNNNNNNN
```

then it means that you're trying to access a workspace that uses private link with private access set to disabled, but you're trying to reach it via public endpoint.  Make sure that domain names resolution is configured correctly to resolve workspace URL to a private endpoint.  Also, this may happen when you’re accessing the internet via an HTTP proxy, so all traffic from Terraform is forwarded to the HTTP proxy, and routed via the public internet.

### Error: ....: Unauthorized access to Org: NNNNNNNNNN

There are a few possible reasons for this error:

* You’re trying to access a Databricks workspace with a private link enabled and public network access set to disabled.  Typically this happens when a computer from which you’re running terraform apply or terraform plan doesn’t have domain name resolution configured correctly, and Terraform is reaching the workspace via a public IP address. Also, this may happen when you’re accessing the internet via an HTTP proxy, so all traffic from Terraform is forwarded to the proxy, and routed via the public internet.
* You have a Databricks workspace with IP Access Lists enabled and you’re trying to access from a computer that isn’t in the list of approved IP addresses.

### Error: Provider registry.terraform.io/databricks/databricks v... does not have a package available for your current platform, windows_386

This error happens when the 32-bit version of Databricks Terraform provider is used, usually on Microsoft Windows.  To fix the issue, you need to switch to the 64-bit versions of Terraform and Databricks Terraform provider.

### Permanent configuration drifts with `databricks_grants` or `databricks_permissions`

For both resources, each single resource instance should manage all the grants/permissions for a given object. If multiple instances are set up against an object, they will keep overwriting one another, leading to permanent configuration drifts.

To prevent that, you need to have only one resource instance per object, and inside that resource instance, use [Dynamic Blocks](https://developer.hashicorp.com/terraform/language/expressions/dynamic-blocks) to specify the variable number of nested grant blocks.

For example

```hcl
locals {
  groups = ["group1", "group2"]
}

resource "databricks_grants" "catalog_grants" {
  catalog = databricks_catalog.catalog_raw.name

  dynamic "grant" {
    for_each = local.groups
    content {
      principal  = grant.value
      privileges = ["ALL_PRIVILEGES"]
    }
  }
}
```

See `databricks_grant` for managing grants for a single principal.

### Error updating UC catalog resources after a metastore_id change

After changing the metastore assigned to a workspace, some resources may fail to update with the following error:

```
metastore_id must be empty or equal to the metastore id assigned to the workspace: <metastore_id>. 
If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set
```

To solve this error, the new Metastore ID must be set in the field `metastore_id` of the failing resources.

### More than one authorization method configured error

If you notice the below error:

```sh
Error: validate: more than one authorization method configured
```

Ensure that you only have one authorization method set. All available authorization methods are documented [here](https://registry.terraform.io/providers/databricks/databricks/latest/docs#auth_type).

If you want to enforce a specific authorization method, you can set the `auth_type` attribute in the provider block:

```hcl
provider "databricks" {
  # other configurations
  auth_type = "pat"
}
```

The above would enforce the use of PAT authorization.


### oauth-m2m: oidc: databricks OAuth is not supported for this host.

There could be different reasons for this error:

- Old version of Terraform provider is used - there were [problems reported](https://github.com/databricks/terraform-provider-databricks/issues/3023) with versions lower than 1.35.0, so try to upgrade your provider.
- You use multiple provider instances in your code and don't specify that specific instance in your resource or data source - in this case, Terraform will try to use "default instance" that could be initialized from the `DEFAULT` profile in your `~/.databrickscfg` file if you have it.  Check that correct provider instance is used everywhere.
- If you're using authentication data from `~/.databrickscfg` file, check that profile you're using has correct data - URL, client ID and secret.

### Provider "registry.terraform.io/databricks/databricks" planned an invalid value for ...: planned value ... for a non-computed attribute.

Starting with version v1.51.0, the Terraform provider for Databricks supports `terraform` versions 1.1.5 and later. Older versions of `terraform`, such as v0.15.5, are known to erroneously generate this error. Check the version of `terraform` that you're using by running `terraform version` and upgrade it if necessary.
