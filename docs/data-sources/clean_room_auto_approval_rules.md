---
subcategory: "Clean Rooms"
---
# databricks_clean_room_auto_approval_rules Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `clean_room_name` (string, required)
* `page_size` (integer, optional) - Maximum number of auto-approval rules to return. Defaults to 100



## Attributes
This data source exports a single attribute, `rules`. It is a list of resources, each with the following attributes:
* `author_collaborator_alias` (string)
* `author_scope` (string) - . Possible values are: `ANY_AUTHOR`
* `clean_room_name` (string) - The name of the clean room this auto-approval rule belongs to
* `created_at` (integer) - Timestamp of when the rule was created, in epoch milliseconds
* `rule_id` (string) - A generated UUID identifying the rule
* `rule_owner_collaborator_alias` (string) - The owner of the rule to whom the rule applies
* `runner_collaborator_alias` (string)