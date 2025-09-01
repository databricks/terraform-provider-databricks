---
subcategory: "Clean Rooms"
---
# databricks_clean_room_auto_approval_rule Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `rule_id` (string, required) - A generated UUID identifying the rule
* `workspace_id` (string, optional) - Workspace ID of the resource

## Attributes
The following attributes are exported:
* `author_collaborator_alias` (string) - Collaborator alias of the author covered by the rule.
  Only one of `author_collaborator_alias` and `author_scope` can be set
* `author_scope` (string) - Scope of authors covered by the rule.
  Only one of `author_collaborator_alias` and `author_scope` can be set. Possible values are: `ANY_AUTHOR`
* `clean_room_name` (string) - The name of the clean room this auto-approval rule belongs to
* `created_at` (integer) - Timestamp of when the rule was created, in epoch milliseconds
* `rule_id` (string) - A generated UUID identifying the rule
* `rule_owner_collaborator_alias` (string) - The owner of the rule to whom the rule applies
* `runner_collaborator_alias` (string) - Collaborator alias of the runner covered by the rule