---
subcategory: "Agent Bricks"
---
# databricks_supervisor_agent Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/supervisoragents)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the Supervisor Agent.
  Format: supervisor-agents/{supervisor_agent_id}
  The server assigns the name when the agent is created
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - The time at which the Supervisor Agent was created, as an RFC 3339 timestamp
* `creator` (string) - The email address of the user who created the Supervisor Agent
* `description` (string) - A summary of what the agent does. It is shown to users of the agent
* `display_name` (string) - The display name of the Supervisor Agent. It must be unique among the agents in the workspace.
  Creating an agent with a display name that is already in use fails with ALREADY_EXISTS
* `endpoint_name` (string) - The name of the Model Serving endpoint that serves the Supervisor Agent. Query this endpoint to
  chat with the agent
* `experiment_id` (string) - The ID of the MLflow experiment that is created together with the Supervisor Agent
* `id` (string, deprecated) - Deprecated: Use supervisor_agent_id instead. This field is no longer populated
* `instructions` (string) - Natural-language guidelines for how the supervisor should respond and how it should delegate
  requests to its tools and subagents
* `name` (string) - The resource name of the Supervisor Agent.
  Format: supervisor-agents/{supervisor_agent_id}
  The server assigns the name when the agent is created
* `supervisor_agent_id` (string) - The universally unique identifier (UUID) of the Supervisor Agent. It is the last segment of `name`