---
subcategory: "Tags"
---
# databricks_tag_policy Data Source
This data source can be used to get a single tag policy by its tag key.

-> **Note** This resource can only be used with an account-level provider!

## Example Usage
```hcl
```

## Arguments
The following arguments are supported:
* `tag_key` (string, required)
* `workspace_id` (string, optional) - Workspace ID of the resource

## Attributes
The following attributes are exported:
* `description` (string)
* `id` (string)
* `tag_key` (string)
* `values` (list of Value)

### Value
* `name` (string)