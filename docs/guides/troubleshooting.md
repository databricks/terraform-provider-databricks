---
page_title: "Troubleshooting Guide"
---

# How to troubleshoot your problem

If you have problems with code that uses Databricks Terraform provider, follow these steps to solve them:

* Check symptoms and solutions in the [Typical problems](#typical-problems) section below.
* Upgrade provider to the latest version. The bug might have already been fixed.
* In case of authentication problems, see the [Data resources and Authentication is not configured errors](#data-resources-and-authentication-is-not-configured-errors) below.
* Collect debug information using following command:

```sh
TF_LOG=DEBUG DATABRICKS_DEBUG_TRUNCATE_BYTES=250000 terraform apply -no-color 2>&1 |tee tf-debug.log
```

* Open a [new GitHub issue](https://github.com/databricks/terraform-provider-databricks/issues/new/choose) providing all information described in the issue template - debug logs, your Terraform code, Terraform & plugin versions, etc.


# Typical problems

## Data resources and Authentication is not configured errors

*In Terraform 0.13 and later*, data resources have the same dependency resolution behavior [as defined for managed resources](https://www.terraform.io/docs/language/resources/behavior.html#resource-dependencies). Most data resources make an API call to a workspace. If a workspace doesn't exist yet, `default auth: cannot configure default credentials` error is raised. To work around this issue and guarantee a proper lazy authentication with data resources, you should add `depends_on = [azurerm_databricks_workspace.this]` or `depends_on = [databricks_mws_workspaces.this]` to the body. This issue doesn't occur if workspace is created *in one module* and resources [within the workspace](guides/workspace-management.md) are created *in another*. We do not recommend using Terraform 0.12 and earlier, if your usage involves data resources.

## Multiple Provider Configurations

The most common reason for technical difficulties might be related to missing `alias` attribute in `provider "databricks" {}` blocks or `provider` attribute in `resource "databricks_..." {}` blocks, when using multiple provider configurations. Please make sure to read [`alias`: Multiple Provider Configurations](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) documentation article. 



## Error while installing: registry does not have a provider

```sh
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
      version = "1.0.1"
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

Alternatively, you can find the hashes of the last 30 provider versions in [`.terraform.lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl). As a temporary measure, you can lock on a prior version by following the following steps:

* Copy [`versions-lock.hcl`](https://github.com/databrickslabs/terraform-provider-databricks/blob/v0.6.2/scripts/versions-lock.hcl) to the root folder of your terraform project.
* Rename to `terraform.lock.hcl`
* Run `terraform init` and verify the provider is installed.
* Be sure to commit the new `.terraform.lock.hcl` file to your source code repository.


## Error: Failed to query available provider packages

See the same steps as in [Error: Failed to install provider](#error-failed-to-install-provider).


## Error: Deployment name cannot be used until a deployment name prefix is defined

You can get this error during provisioning of the Databricks workspace.  It arises when you're trying to set `deployment_name` by no deployment prefix was set on the Databricks side (you can't set it yourself).  The problem could be solved one of the following methods:

1. Contact your Databricks representative, like Solutions Architect, Customer Success Engineer, Account Executive, or Partner Solutions Architect to set a deployment prefix for your account.

1. Comment out the `deployment_name` parameter to create workspace with default URL: `dbc-XXXXXX.cloud.databricks.com`.

## Error: 'strconv.ParseInt parsing "...." value out of range' or "Attribute must be a whole number, got N.NNNNe+XX"

This kind of errors happens when the 32-bit version of Databricks Terraform provider is used, usually on Microsoft Windows.  To fix the issue you need to switch to use of the 64-bit versions of Terraform and Databricks Terraform provider.

## Error: cannot create xxxx: HTTP method POST is not supported by this URL

This error may appear when creating Databricks users/groups/service principals on Databricks account level when no `account_id` is specified in the provider's configuration.  Make sure that `account_id` is specified & has a correct value.

## Error: oauth-m2m: oidc: parse .well-known: invalid character '<' looking for beginning of value.

This similar to previous item.  Make sure that `account_id` is specified in the provider configuration & it has a correct value.
