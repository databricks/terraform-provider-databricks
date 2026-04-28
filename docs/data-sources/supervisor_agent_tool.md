---
subcategory: "Agent Bricks"
---
# databricks_supervisor_agent_tool Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Full resource name:
  supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `app` (App)
* `uc_connection` (UcConnection)
* `description` (string) - Description of what this tool does (user-facing)
* `genie_space` (GenieSpace)
* `id` (string, deprecated) - Deprecated: Use tool_id instead
* `knowledge_assistant` (KnowledgeAssistant)
* `name` (string) - Full resource name:
  supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
* `tool_id` (string) - User specified id of the Tool
* `tool_type` (string) - Tool type. Must be one of: "genie_space", "knowledge_assistant", "uc_function", "connection", "app", "volume", "lakeview_dashboard", "serving_endpoint", "uc_table", "vector_search_index"
* `uc_function` (UcFunction)
* `volume` (Volume)

### App
* `name` (string) - App name

### UcConnection
* `name` (string)

### GenieSpace
* `id` (string) - The ID of the genie space

### KnowledgeAssistant
* `knowledge_assistant_id` (string) - The ID of the knowledge assistant
* `serving_endpoint_name` (string, deprecated) - Deprecated: use knowledge_assistant_id instead

### UcFunction
* `name` (string) - Full uc function name

### Volume
* `name` (string) - Full uc volume name