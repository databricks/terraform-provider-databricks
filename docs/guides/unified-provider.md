# Unified Terraform Provider
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

## Introduction

The Unified Terraform Provider allows you to manage workspace-level resources and data sources through an account-level provider. This significantly simplifies Terraform configurations and resource management, as only one provider is needed per account.

**Note:** This feature is in [Public Beta](https://docs.databricks.com/aws/en/release-notes/release-types). If you experience any issues, please refer to the [Reporting Issues](#reporting-issues) section below.

## Usage

To manage a workspace-level resource through the account provider, specify the `provider_config` at the resource level with the `workspace_id` of the workspace the resource belongs to.

Depending on the internal implementation of the resource or data source, `provider_config` can be either a block or an attribute. For details, please refer to the documentation of the specific resource or data source.

### Block
```hcl
resource "workspace_level_resource" "this" {
    provider_config {
        workspace_id = "12345"
    }
    ...
}
```

### Attribute
```hcl
resource "workspace_level_resource" "this" {
    provider_config = {
        workspace_id = "12345"
    }
    ...
}
```

**Note:** This feature is being rolled out incrementally. Some resources do not yet support the unified provider. Please check the resource-specific documentation to see if the `provider_config` attribute or block is available.

## Migrating to Unified Provider

If you are currently managing both workspace-level and account-level resources and data sources, you likely have multiple provider configurations that you specify for each resource using aliases. For example:
```hcl
// Define an account provider with alias
provider "databricks" {
    alias         = "account"
    host          = var.account_host
    account_id    = var.account_id
    client_id     = var.account_client_id
    client_secret = var.account_client_secret
}

// Create a workspace
resource "databricks_mws_workspaces" "this" {
    provider = databricks.account
    ...
}
```

Once the workspace is created, create a workspace-level provider.
```hcl
// Define a workspace provider with alias
provider "databricks" {
    alias           = "workspace"
    host            = var.workspace_host
    client_id       = var.workspace_client_id
    client_secret   = var.workspace_client_secret
}

// Use the workspace provider for workspace-level resources
resource "databricks_workspace_level_resource" "this" {
    provider = databricks.workspace
    name = "resource name"
}
```

Migration to unified provider happens in 2 steps.
1. Add `provider_config` and `workspace_id` to the resource without removing the workspace-level provider. Then do terraform apply so these values are part of the state.
2. Remove the workspace-level provider.

you can remove the workspace-level provider and specify the `workspace_id` in the `provider_config` attribute or block instead. For example:

```hcl
// Define an account provider
provider "databricks" {
    host          = var.account_host
    account_id    = var.account_id
    client_id     = var.account_client_id
    client_secret = var.account_client_secret
}

// Create a workspace under the account
resource "databricks_mws_workspaces" "this" {
    account_id   = var.account_id
    aws_region   = "us-east-1"
    compute_mode = "SERVERLESS"
}

// Create a workspace-level resource using provider_config
resource "databricks_workspace_level_resource" "this" {
    provider_config {
        workspace_id = databricks_mws_workspaces.this.workspace_id
    }
    name = "resource name"
}
```

## FAQ
* Workspace for which the workspace_id is supplied to the resource through provider_config must belong to the account the provider is configured with.
* Migration: Doing the migration in 1 step will lead to issues. This happens because the state doesn't have workspace_id during the refresh.

## Limitations

There are some limitations to this feature that we plan to address in the near future:

1. Databricks CLI and Azure CLI authentication methods are not currently supported
2. Some resources do not yet support the unified provider. Please refer to the documentation for each resource or data source to check if they support the `provider_config` attribute or block.

## Reporting Issues

This feature is in Public Beta. If you encounter any issues, please report them on [GitHub Issues](https://github.com/databricks/terraform-provider-databricks/issues) with the `Unified Provider` label.
