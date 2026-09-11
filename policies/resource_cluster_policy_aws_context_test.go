package policies

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

const awsContextPolicyDefinition = `{"worker_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-worker"},"driver_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-driver"}}`

func TestResourceClusterPolicyAWSContextCreate(t *testing.T) {
	for _, mode := range []string{"definition", "family", "builtin"} {
		t.Run(mode, func(t *testing.T) {
			policy := compute.Policy{PolicyId: "abc", Name: "AWS Context ID", Definition: awsContextPolicyDefinition}
			state := map[string]any{"name": policy.Name, "definition": awsContextPolicyDefinition}
			request := compute.CreatePolicy{Name: policy.Name, Definition: awsContextPolicyDefinition}
			if mode != "definition" {
				delete(state, "definition")
				state["policy_family_id"] = "personal-vm"
				state["policy_family_definition_overrides"] = awsContextPolicyDefinition
				policy.PolicyFamilyId = "personal-vm"
				policy.PolicyFamilyDefinitionOverrides = awsContextPolicyDefinition
				request.Definition = ""
				request.PolicyFamilyId = "personal-vm"
				request.PolicyFamilyDefinitionOverrides = awsContextPolicyDefinition
			}
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					policies := w.GetMockClusterPoliciesAPI().EXPECT()
					if mode != "definition" {
						familyName := "Personal Compute"
						if mode == "builtin" {
							familyName = policy.Name
						}
						w.GetMockPolicyFamiliesAPI().EXPECT().ListAll(mock.Anything, compute.ListPolicyFamiliesRequest{}).
							Return([]compute.PolicyFamily{{PolicyFamilyId: "personal-vm", Name: familyName}}, nil).Once()
					}
					if mode == "builtin" {
						policies.GetByName(mock.Anything, policy.Name).
							Return(&compute.Policy{PolicyId: policy.PolicyId, Name: policy.Name}, nil).Once()
						policies.Edit(mock.Anything, compute.EditPolicy{
							PolicyId:                        policy.PolicyId,
							Name:                            policy.Name,
							PolicyFamilyId:                  policy.PolicyFamilyId,
							PolicyFamilyDefinitionOverrides: awsContextPolicyDefinition,
						}).Return(nil).Once()
					} else {
						policies.Create(mock.Anything, request).
							Return(&compute.CreatePolicyResponse{PolicyId: policy.PolicyId}, nil).Once()
					}
					policies.GetByPolicyId(mock.Anything, policy.PolicyId).Return(&policy, nil).Once()
				},
				Resource: ResourceClusterPolicy(),
				State:    state,
				Create:   true,
			}.ApplyAndExpectData(t, map[string]any{
				"id":                                 policy.PolicyId,
				"policy_id":                          policy.PolicyId,
				"definition":                         awsContextPolicyDefinition,
				"policy_family_id":                   policy.PolicyFamilyId,
				"policy_family_definition_overrides": policy.PolicyFamilyDefinitionOverrides,
			})
		})
	}
}

func TestResourceClusterPolicyAWSContextUpdate(t *testing.T) {
	for _, field := range []string{"definition", "policy_family_definition_overrides"} {
		for _, tc := range []struct {
			name       string
			definition string
		}{
			{"change", `{"worker_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-updated"},"driver_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-driver"}}`},
			{"remove_worker", `{"driver_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-driver"}}`},
			{"remove_driver", `{"worker_node_type_flexibility.aws_context_id":{"type":"fixed","value":"context-worker"}}`},
		} {
			t.Run(field+"/"+tc.name, func(t *testing.T) {
				state := map[string]any{"name": "AWS Context ID", field: tc.definition}
				oldState := map[string]string{"id": "abc", "name": "AWS Context ID", field: awsContextPolicyDefinition}
				request := compute.EditPolicy{PolicyId: "abc", Name: "AWS Context ID", Definition: tc.definition}
				policy := compute.Policy{PolicyId: "abc", Name: request.Name, Definition: tc.definition}
				if field == "policy_family_definition_overrides" {
					state["policy_family_id"] = "personal-vm"
					oldState["policy_family_id"] = "personal-vm"
					oldState["definition"] = awsContextPolicyDefinition
					request.Definition = ""
					request.PolicyFamilyId = "personal-vm"
					request.PolicyFamilyDefinitionOverrides = tc.definition
					policy.PolicyFamilyId = "personal-vm"
					policy.PolicyFamilyDefinitionOverrides = tc.definition
				}
				qa.ResourceFixture{
					MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
						policies := w.GetMockClusterPoliciesAPI().EXPECT()
						policies.Edit(mock.Anything, request).Return(nil).Once()
						policies.GetByPolicyId(mock.Anything, policy.PolicyId).Return(&policy, nil).Once()
					},
					Resource:      ResourceClusterPolicy(),
					State:         state,
					InstanceState: oldState,
					Update:        true,
					ID:            policy.PolicyId,
				}.ApplyAndExpectData(t, map[string]any{
					"id":                                 policy.PolicyId,
					"definition":                         tc.definition,
					"policy_family_id":                   policy.PolicyFamilyId,
					"policy_family_definition_overrides": policy.PolicyFamilyDefinitionOverrides,
				})
			})
		}
	}
}

func TestResourceClusterPolicyAWSContextRead(t *testing.T) {
	for _, familyID := range []string{"", "personal-vm"} {
		t.Run("family="+familyID, func(t *testing.T) {
			policy := compute.Policy{
				PolicyId:       "abc",
				Name:           "AWS Context ID",
				Definition:     awsContextPolicyDefinition,
				PolicyFamilyId: familyID,
			}
			if familyID != "" {
				policy.PolicyFamilyDefinitionOverrides = awsContextPolicyDefinition
			}
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockClusterPoliciesAPI().EXPECT().GetByPolicyId(mock.Anything, policy.PolicyId).
						Return(&policy, nil).Once()
				},
				Resource: ResourceClusterPolicy(),
				Read:     true,
				New:      true,
				ID:       policy.PolicyId,
			}.ApplyAndExpectData(t, map[string]any{
				"id":                                 policy.PolicyId,
				"policy_id":                          policy.PolicyId,
				"name":                               policy.Name,
				"definition":                         awsContextPolicyDefinition,
				"policy_family_id":                   policy.PolicyFamilyId,
				"policy_family_definition_overrides": policy.PolicyFamilyDefinitionOverrides,
			})
		})
	}
}
