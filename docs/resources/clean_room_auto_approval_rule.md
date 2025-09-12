---
subcategory: "Clean Rooms"
---
# databricks_clean_room_auto_approval_rule Resource


## Example Usage


## Arguments
The following arguments are supported:
* `author_collaborator_alias` (string, optional) - Collaborator alias of the author covered by the rule.
  Only one of `author_collaborator_alias` and `author_scope` can be set
* `author_scope` (string, optional) - Scope of authors covered by the rule.
  Only one of `author_collaborator_alias` and `author_scope` can be set. Possible values are: `ANY_AUTHOR`
* `clean_room_name` (string, optional) - The name of the clean room this auto-approval rule belongs to
* `runner_collaborator_alias` (string, optional) - Collaborator alias of the runner covered by the rule

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