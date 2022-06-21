---
page_title: "Troubleshooting Guide"
---

* [How to troubleshoot your problem](#how-to-troubleshoot-your-problem)
* [Typical problems](#typical-problems)
   * [Data resources and Authentication is not configured errors](#data-resources-and-authentication-is-not-configured-errors)
   * [Multiple Provider Configurations](#multiple-provider-configurations)
   * [Error while installing: registry does not have a provider](#error-while-installing-registry-does-not-have-a-provider)
   * [Error: Failed to install provider](#error-failed-to-install-provider)
   * [Error: Failed to query available provider packages](#error-failed-to-query-available-provider-packages)
   * [Error: Deployment name cannot be used until a deployment name prefix is defined](#error-deployment-name-cannot-be-used-until-a-deployment-name-prefix-is-defined)
   * [Azure KeyVault cannot yet be configured for Service Principal authorization](#azure-keyvault-cannot-yet-be-configured-for-service-principal-authorization)


# How to troubleshoot your problem

If you have problems with code that uses Databricks Terraform provider, follow these steps to solve them:

* Check symptoms and solutions in the [Typical problems](#typical-problems) section below.
* Upgrade provider to the latest version. The bug might have already been fixed.
* In case of authentication problems, see the [Data resources and Authentication is not configured errors](#data-resources-and-authentication-is-not-configured-errors) below.
* Collect debug information using following command:

```sh
TF_LOG=DEBUG DATABRICKS_DEBUG_TRUNCATE_BYTES=250000 terraform apply 2>&1 > tf-debug.log
```

* Open a [new GitHub issue](https://github.com/databricks/terraform-provider-databricks/issues/new/choose) providing all information described in the issue template - debug logs, your Terraform code, Terraform & plugin versions, etc.


# Typical problems

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
      source  = "databricks/databricks"
      version = "1.0.0"
    }
  }
}
```

... and copy the file in every module in your codebase. Our recommendation is to skip the `version` field for `versions.tf` file on module level, and keep it only on the environment level.

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

## Error: Failed to install provider

Running the `terraform init` command, you may see `Failed to install provider` error if you didn't check-in [`.terraform.lock.hcl`](https://www.terraform.io/language/files/dependency-lock#lock-file-location) to the source code version control:

```sh
Error: Failed to install provider

Error while installing databricks/databricks: v1.0.0: checksum list has no SHA-256 hash for "https://github.com/databricks/terraform-provider-databricks/releases/download/v1.0.0/terraform-provider-databricks_1.0.0_darwin_amd64.zip"
```

You can fix it by following three simple steps: 

* Replace `databrickslabs/databricks` with `databricks/databricks` in all your `.tf` files with the `python3 -c "$(curl -Ls https://dbricks.co/updtfns)"` command. 
* Run the `terraform state replace-provider databrickslabs/databricks databricks/databricks` command and approve the changes. See [Terraform CLI](https://www.terraform.io/cli/commands/state/replace-provider) docs for more information.
* Run `terraform init` to verify everything working.

The terraform apply command should work as expected now.

Alternatively, you can find the hashes of the last 30 provider versions in [`.terraform.lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl).

## Error: Failed to query available provider packages

see the same steps as in [Error: Failed to install provider](#error-failed-to-install-provider).


## Error: Deployment name cannot be used until a deployment name prefix is defined

You can get this error during provisioning of the Databricks workspace.  It arises when you're trying to set `deployment_name` by no deployment prefix was set on the Databricks side (you can't set it yourself).  The problem could be solved one of the following methods:

1. Contact your Databricks representative, like Solutions Architect, Customer Success Engineer, Account Executive, or Partner Solutions Architect to set a deployment prefix for your account.

1. Comment out the `deployment_name` parameter to create workspace with default URL: `dbc-XXXXXX.cloud.databricks.com`.


## Azure KeyVault cannot yet be configured for Service Principal authorization

This is a well known limitation of the Azure Databricks - currently you cannot create Azure Key Vault-based secret scope because OBO flow is not supported yet for service principals on Azure Active Directory side.  Use [azure-cli authentication](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authenticating-with-azure-cli) with user principal to create AKV-based secret scope. 
