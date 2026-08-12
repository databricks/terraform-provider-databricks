package exporter

import (
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
	"github.com/databricks/databricks-sdk-go/service/supervisoragents"
	knowledge_assistant_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/knowledge_assistant"
	knowledge_source_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/knowledge_assistant_knowledge_source"
	supervisor_agent_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/supervisor_agent"
	supervisor_agent_tool_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/supervisor_agent_tool"
)

func listKnowledgeAssistants(ic *importContext) error {
	assistants, err := ic.workspaceClient.KnowledgeAssistants.ListKnowledgeAssistantsAll(ic.Context,
		knowledgeassistants.ListKnowledgeAssistantsRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, assistant := range assistants {
		if !ic.MatchesName(assistant.DisplayName) {
			log.Printf("[INFO] Skipping knowledge assistant %s because it doesn't match %s",
				assistant.DisplayName, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_knowledge_assistant",
			ID:       assistant.Name,
		}, 0, fmt.Sprintf("knowledge assistant '%s'", assistant.DisplayName))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Knowledge Assistants", i)
	}
	return nil
}

func importKnowledgeAssistant(ic *importContext, r *resource) error {
	var assistant knowledgeassistants.KnowledgeAssistant
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		knowledge_assistant_resource.KnowledgeAssistant{}, &assistant); err != nil {
		return err
	}

	sources, err := ic.workspaceClient.KnowledgeAssistants.ListKnowledgeSourcesAll(ic.Context,
		knowledgeassistants.ListKnowledgeSourcesRequest{Parent: r.ID})
	if err != nil {
		log.Printf("[WARN] Failed to list knowledge sources for %s: %s", r.ID, err)
		return nil
	}
	for _, source := range sources {
		ic.Emit(&resource{
			Resource: "databricks_knowledge_assistant_knowledge_source",
			ID:       source.Name,
		})
	}
	return nil
}

func importKnowledgeSource(ic *importContext, r *resource) error {
	var source knowledgeassistants.KnowledgeSource
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		knowledge_source_resource.KnowledgeSource{}, &source); err != nil {
		return err
	}

	if source.FileTable != nil && source.FileTable.TableName != "" {
		ic.Emit(&resource{
			Resource: "databricks_sql_table",
			ID:       source.FileTable.TableName,
		})
	}
	if source.Index != nil && source.Index.IndexName != "" {
		ic.Emit(&resource{
			Resource: "databricks_vector_search_index",
			ID:       source.Index.IndexName,
		})
	}
	if source.Files != nil && source.Files.Path != "" {
		// Path format: /Volumes/<catalog>/<schema>/<volume>/<optional subdirs/files>
		parts := strings.Split(source.Files.Path, "/")
		if len(parts) >= 5 && parts[1] == "Volumes" {
			ic.Emit(&resource{
				Resource: "databricks_volume",
				ID:       strings.Join(parts[2:5], "."),
			})
		}
	}
	return nil
}

func listSupervisorAgents(ic *importContext) error {
	agents, err := ic.workspaceClient.SupervisorAgents.ListSupervisorAgentsAll(ic.Context,
		supervisoragents.ListSupervisorAgentsRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, agent := range agents {
		if !ic.MatchesName(agent.DisplayName) {
			log.Printf("[INFO] Skipping supervisor agent %s because it doesn't match %s",
				agent.DisplayName, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_supervisor_agent",
			ID:       agent.Name,
		}, 0, fmt.Sprintf("supervisor agent '%s'", agent.DisplayName))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Supervisor Agents", i)
	}
	return nil
}

func importSupervisorAgent(ic *importContext, r *resource) error {
	var agent supervisoragents.SupervisorAgent
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		supervisor_agent_resource.SupervisorAgent{}, &agent); err != nil {
		return err
	}

	tools, err := ic.workspaceClient.SupervisorAgents.ListToolsAll(ic.Context,
		supervisoragents.ListToolsRequest{Parent: r.ID})
	if err != nil {
		log.Printf("[WARN] Failed to list tools for %s: %s", r.ID, err)
		return nil
	}
	for _, tool := range tools {
		log.Printf("[INFO] Tool: %s", tool.Name)
		ic.Emit(&resource{
			Resource: "databricks_supervisor_agent_tool",
			ID:       tool.Name,
		})
	}
	return nil
}

// shouldOmitSupervisorAgentToolField skips deprecated API-only fields that should not appear in exported HCL.
func shouldOmitSupervisorAgentToolField(ic *importContext, pathString string, fieldSchema FieldSchema, wrapper ResourceDataWrapper, r *resource) bool {
	if pathString == "knowledge_assistant.serving_endpoint_name" {
		return true
	}
	return DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, r)
}

func importSupervisorAgentTool(ic *importContext, r *resource) error {
	var tool supervisoragents.Tool
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		supervisor_agent_tool_resource.Tool{}, &tool); err != nil {
		return err
	}

	if tool.App != nil && tool.App.Name != "" {
		ic.Emit(&resource{
			Resource: "databricks_app",
			ID:       tool.App.Name,
		})
	}
	if tool.UcConnection != nil && tool.UcConnection.Name != "" {
		ic.Emit(&resource{
			Resource: "databricks_connection",
			ID:       tool.UcConnection.Name,
		})
	}
	if tool.Volume != nil && tool.Volume.Name != "" {
		ic.Emit(&resource{
			Resource: "databricks_volume",
			ID:       tool.Volume.Name,
		})
	}
	if tool.KnowledgeAssistant != nil && tool.KnowledgeAssistant.KnowledgeAssistantId != "" {
		ic.Emit(&resource{
			Resource: "databricks_knowledge_assistant",
			ID:       "knowledge-assistants/" + tool.KnowledgeAssistant.KnowledgeAssistantId,
		})
	}
	return nil
}
