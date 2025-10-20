# Unified Terraform Provider
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

## Introduction
Manage workspace level resource through acccount level provider!

## Usability
Specify `provider_config` at the resource level which would contain the workspace_id that the resource will belong to. This can be used as either block or attribute depending on the internal implementations of the resource. For details, please see the documentation of the specific resource or data source.

### Block
```hcl
resource "workspace_level_resource" "this" {
    provider_config {
        workspace_id = "12345"
    }
    ...
}
```

#### Attribute
```hcl
resource "workspace_level_resource" "this" {
    provider_config = {
        workspace_id = "12345"
    }
    ...
}
```

## Migrating to Unified Provider
If you are managing both workspace and account level resources, you would have a 2 provider setup. For example:
```hcl
```

To migrate to unified provider, you can remove the workspace level provider and specify the workspace_id which can be through the workspace resource or environment variable. Example

```hcl
```

## Limitations
There are some limitations to this feature and we plan to address them in near future.
1. Databricks and Azure CLI aren't supported currently
2. Some resources don't support unified provider. Please see the documentation for relevant resource or datasource if they have the `provider_config` attribute or block.

## Issues
This feature is in Public Beta. In case you encounter an issue. Please report it to us on https://github.com/databricks/terraform-provider-databricks/issues and `Unified Provider` label.
