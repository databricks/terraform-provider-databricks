---
subcategory: "Tags"
---
# databricks_workspace_entity_tag_assignment Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This resource allows you to create, update, list, and delete tag assignments for workspace scoped entities.

## Example Usage
resource "databricks_workspace_entity_tag_assignment" "dashboard_tag" {
  entity_type = "dashboards"
  entity_id   = "2807324866692453"
  tag_key     = "sensitivity_level"
  tag_value   = "high"
}

resource "databricks_workspace_entity_tag_assignment" "geniespace_tag" {
  entity_type = "geniespaces"
  entity_id   = "2807324866692453"
  tag_key     = "sensitivity_level"
  tag_value   = "high"
}


## Arguments
The following arguments are supported:
* `entity_id` (string, required) - The identifier of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of entity to which the tag is assigned. Allowed values are dashboards, geniespaces
* `tag_key` (string, required) - The key of the tag. The characters , . : / - = and leading/trailing spaces are not allowed
* `tag_value` (string, optional) - The value of the tag



## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "entity_type,entity_id,tag_key"
  to = databricks_workspace_entity_tag_assignment.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_entity_tag_assignment.this "entity_type,entity_id,tag_key"
```