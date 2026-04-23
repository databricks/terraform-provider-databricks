---
subcategory: "Agent Bricks"
---
# databricks_supervisor_agent Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the SupervisorAgent.
  Format: supervisor-agents/{supervisor_agent_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - Creation timestamp
* `creator` (string) - The creator of the Supervisor Agent
* `description` (string) - Description of what this agent can do (user-facing)
* `display_name` (string) - The display name of the Supervisor Agent, unique at workspace level
* `endpoint_name` (string) - The name of the supervisor agent's serving endpoint
* `experiment_id` (string) - The MLflow experiment ID
* `id` (string, deprecated) - Deprecated: Use supervisor_agent_id instead
* `instructions` (string) - Optional natural-language instructions for the supervisor agent
* `name` (string) - The resource name of the SupervisorAgent.
  Format: supervisor-agents/{supervisor_agent_id}
* `supervisor_agent_id` (string) - The universally unique identifier (UUID) of the Supervisor Agent