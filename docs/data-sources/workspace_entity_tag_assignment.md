---
subcategory: "Tags"
---
# databricks_workspace_entity_tag_assignment Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source allows you to get information about a tag assignment for a specific workspace scoped entity using the entity type, entity id, and tag key.

## Example Usage
```hcl
data "databricks_workspace_entity_tag_assignment" "dashboard_tag" {
  entity_type = "dashboards"
  entity_id   = "2807324866692453"
  tag_key     = "sensitivity_level"
}

data "databricks_workspace_entity_tag_assignment" "geniespace_tag" {
  entity_type = "geniespaces"
  entity_id   = "2807324866692453"
  tag_key     = "sensitivity_level"
}
```


## Arguments
The following arguments are supported:
* `entity_id` (string, required) - The identifier of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of entity to which the tag is assigned. Allowed values are dashboards, geniespaces
* `tag_key` (string, required) - The key of the tag. The characters , . : / - = and leading/trailing spaces are not allowed

## Attributes
The following attributes are exported:
* `entity_id` (string) - The identifier of the entity to which the tag is assigned
* `entity_type` (string) - The type of entity to which the tag is assigned. Allowed values are dashboards, geniespaces
* `tag_key` (string) - The key of the tag. The characters , . : / - = and leading/trailing spaces are not allowed
* `tag_value` (string) - The value of the tag