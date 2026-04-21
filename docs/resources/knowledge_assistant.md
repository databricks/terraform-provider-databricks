---
subcategory: "Agent Bricks"
---
# databricks_knowledge_assistant Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `description` (string, required) - Description of what this agent can do (user-facing).
  Required when creating a Knowledge Assistant.
  When updating a Knowledge Assistant, optional unless included in
  update_mask
* `display_name` (string, required) - The display name of the Knowledge Assistant, unique at workspace level.
  Required when creating a Knowledge Assistant.
  When updating a Knowledge Assistant, optional unless included in
  update_mask
* `instructions` (string, optional) - Additional global instructions on how the agent should generate answers.
  Optional on create and update.
  When updating a Knowledge Assistant, include this field in update_mask to
  modify it
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Creation timestamp
* `creator` (string) - The creator of the Knowledge Assistant
* `endpoint_name` (string) - The name of the knowledge assistant agent endpoint
* `error_info` (string) - Error details when the Knowledge Assistant is in FAILED state
* `experiment_id` (string) - The MLflow experiment ID
* `id` (string) - The universally unique identifier (UUID) of the Knowledge Assistant
* `name` (string) - The resource name of the Knowledge Assistant.
  Format: knowledge-assistants/{knowledge_assistant_id}
* `state` (string) - State of the Knowledge Assistant. Not returned in List responses. Possible values are: `ACTIVE`, `CREATING`, `FAILED`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_knowledge_assistant.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_knowledge_assistant.this "name"
```