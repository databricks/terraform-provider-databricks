package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
	"github.com/databricks/databricks-sdk-go/service/supervisoragents"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListKnowledgeAssistants(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.1/knowledge-assistants?",
			ReuseRequest: true,
			Response: knowledgeassistants.ListKnowledgeAssistantsResponse{
				KnowledgeAssistants: []knowledgeassistants.KnowledgeAssistant{
					{
						Name:        "knowledge-assistants/abc-123",
						DisplayName: "Sales Assistant",
					},
					{
						Name:        "knowledge-assistants/def-456",
						DisplayName: "Support Assistant",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("agentbricks")
		err := resourcesMap["databricks_knowledge_assistant"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_knowledge_assistant[<unknown>] (id: knowledge-assistants/abc-123)"])
		assert.True(t, ic.testEmits["databricks_knowledge_assistant[<unknown>] (id: knowledge-assistants/def-456)"])
	})
}

func TestListKnowledgeAssistantsWithMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/knowledge-assistants?",
			Response: knowledgeassistants.ListKnowledgeAssistantsResponse{
				KnowledgeAssistants: []knowledgeassistants.KnowledgeAssistant{
					{
						Name:        "knowledge-assistants/abc-123",
						DisplayName: "Sales Assistant",
					},
					{
						Name:        "knowledge-assistants/def-456",
						DisplayName: "Other",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("agentbricks")
		ic.match = "Sales"
		err := resourcesMap["databricks_knowledge_assistant"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_knowledge_assistant[<unknown>] (id: knowledge-assistants/abc-123)"])
	})
}

func TestListSupervisorAgents(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.1/supervisor-agents?",
			ReuseRequest: true,
			Response: supervisoragents.ListSupervisorAgentsResponse{
				SupervisorAgents: []supervisoragents.SupervisorAgent{
					{
						Name:        "supervisor-agents/agent-1",
						DisplayName: "Primary Agent",
					},
					{
						Name:        "supervisor-agents/agent-2",
						DisplayName: "Secondary Agent",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("agentbricks")
		err := resourcesMap["databricks_supervisor_agent"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_supervisor_agent[<unknown>] (id: supervisor-agents/agent-1)"])
		assert.True(t, ic.testEmits["databricks_supervisor_agent[<unknown>] (id: supervisor-agents/agent-2)"])
	})
}

func TestListSupervisorAgentsWithMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/supervisor-agents?",
			Response: supervisoragents.ListSupervisorAgentsResponse{
				SupervisorAgents: []supervisoragents.SupervisorAgent{
					{
						Name:        "supervisor-agents/agent-1",
						DisplayName: "Primary Agent",
					},
					{
						Name:        "supervisor-agents/agent-2",
						DisplayName: "Other",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("agentbricks")
		ic.match = "Primary"
		err := resourcesMap["databricks_supervisor_agent"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_supervisor_agent[<unknown>] (id: supervisor-agents/agent-1)"])
	})
}

func TestSupervisorAgentToolNameUnified(t *testing.T) {
	ic := importContextForTest()
	ir := resourcesMap["databricks_supervisor_agent_tool"]

	wrapperWithToolId := &mockResourceDataWrapper{
		data: map[string]any{
			"id":      "supervisor-agents/abc/tools/xyz",
			"tool_id": "my_tool",
		},
	}
	assert.Equal(t, "my_tool_supervisor-agents/abc/tools/xyz", ir.NameUnified(ic, wrapperWithToolId))

	wrapperWithoutToolId := &mockResourceDataWrapper{
		data: map[string]any{
			"id": "supervisor-agents/abc/tools/xyz",
		},
	}
	assert.Equal(t, "supervisor-agents/abc/tools/xyz", ir.NameUnified(ic, wrapperWithoutToolId))
}
