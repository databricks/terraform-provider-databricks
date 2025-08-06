---
subcategory: "Clean Rooms"
---
# databricks_clean_room_auto_approval_rule Resource


## Example Usage


## Arguments
The following arguments are supported:
* `author_collaborator_alias` (string, optional)
* `author_scope` (string, optional) - . Possible values are: `ANY_AUTHOR`
* `clean_room_name` (string, optional) - The name of the clean room this auto-approval rule belongs to
* `runner_collaborator_alias` (string, optional)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `created_at` (integer) - Timestamp of when the rule was created, in epoch milliseconds
* `rule_id` (string) - A generated UUID identifying the rule
* `rule_owner_collaborator_alias` (string) - The owner of the rule to whom the rule applies

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = rule_id
  to = databricks_clean_room_auto_approval_rule.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_clean_room_auto_approval_rule rule_id
```