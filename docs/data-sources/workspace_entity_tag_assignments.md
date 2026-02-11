---
subcategory: "Tags"
---
# databricks_workspace_entity_tag_assignments Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source allows you to retrieve tag assignments that have been applied to a particular workspace scoped entity.

## Example Usage
```hcl
data "databricks_workspace_entity_tag_assignments" "app_tags" {
  entity_type = "apps"
  entity_id   = "2807324866692453"
}

data "databricks_workspace_entity_tag_assignments" "dashboard_tags" {
  entity_type = "dashboards"
  entity_id   = "2807324866692453"
}

data "databricks_workspace_entity_tag_assignments" "geniespace_tags" {
  entity_type = "geniespaces"
  entity_id   = "2807324866692453"
}
```


## Arguments
The following arguments are supported:
* `entity_id` (string, required) - The identifier of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of entity to which the tag is assigned. Allowed values are apps, dashboards, geniespaces
* `page_size` (integer, optional) - Optional. Maximum number of tag assignments to return in a single page
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `tag_assignments`. It is a list of resources, each with the following attributes:
* `entity_id` (string) - The identifier of the entity to which the tag is assigned
* `entity_type` (string) - The type of entity to which the tag is assigned. Allowed values are apps, dashboards, geniespaces
* `tag_key` (string) - The key of the tag. The characters , . : / - = and leading/trailing spaces are not allowed
* `tag_value` (string) - The value of the tag