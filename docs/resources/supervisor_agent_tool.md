---
subcategory: "Agent Bricks"
---
# databricks_supervisor_agent_tool Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `description` (string, required) - Description of what this tool does (user-facing)
* `parent` (string, required) - Parent resource where this tool will be created.
  Format: supervisor-agents/{supervisor_agent_id}
* `tool_id` (string, required) - User specified id of the Tool
* `tool_type` (string, required) - Tool type. Must be one of: "genie_space", "knowledge_assistant", "uc_function", "uc_connection", "app", "volume", "lakeview_dashboard", "serving_endpoint", "uc_table", "vector_search_index"
* `app` (App, optional)
* `genie_space` (GenieSpace, optional)
* `knowledge_assistant` (KnowledgeAssistant, optional)
* `uc_connection` (UcConnection, optional)
* `uc_function` (UcFunction, optional)
* `volume` (Volume, optional)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### App
* `name` (string, required) - App name

### GenieSpace
* `id` (string, required) - The ID of the genie space

### KnowledgeAssistant
* `knowledge_assistant_id` (string, required) - The ID of the knowledge assistant
* `serving_endpoint_name` (string, optional, deprecated) - Deprecated: use knowledge_assistant_id instead

### UcConnection
* `name` (string, required)

### UcFunction
* `name` (string, required) - Full uc function name

### Volume
* `name` (string, required) - Full uc volume name

## Attributes
In addition to the above arguments, the following attributes are exported:
* `id` (string, deprecated) - Deprecated: Use tool_id instead
* `name` (string) - Full resource name:
  supervisor-agents/{supervisor_agent_id}/tools/{tool_id}

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_supervisor_agent_tool.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_supervisor_agent_tool.this "name"
```