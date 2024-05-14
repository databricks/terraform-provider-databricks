package policies

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataSourceClusterPolicy(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Response: compute.ListPoliciesResponse{
					Policies: []compute.Policy{
						{
							PolicyId:                        "abc",
							Name:                            "policy",
							Definition:                      `{"abc":"123"}`,
							Description:                     "A description",
							PolicyFamilyId:                  "def",
							PolicyFamilyDefinitionOverrides: `{"def":"456"}`,
							IsDefault:                       true,
							MaxClustersPerUser:              42,
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceClusterPolicy(),
		ID:          ".",
		HCL:         `name = "policy"`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                                 "abc",
		"definition":                         `{"abc":"123"}`,
		"description":                        "A description",
		"policy_family_id":                   "def",
		"policy_family_definition_overrides": `{"def":"456"}`,
		"is_default":                         true,
		"max_clusters_per_user":              42,
	})
}

func TestDataSourceClusterPolicyError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Status:   404,
				Response: apierr.APIError{
					Message: "searching_error",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceClusterPolicy(),
		ID:          ".",
		HCL:         `name = "policy"`,
	}.ExpectError(t, "searching_error")
}

func TestDataSourceClusterPolicyNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Response: compute.ListPoliciesResponse{},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceClusterPolicy(),
		ID:          ".",
		HCL:         `name = "policy"`,
	}.ExpectError(t, "Policy named 'policy' does not exist")
}

func TestDataSourceClusterPolicyStateUpgrader(t *testing.T) {
	state, err := removeZeroMaxClustersPerUser(context.Background(),
		map[string]any{
			"max_clusters_per_user": 0,
		}, nil)
	assert.NoError(t, err)
	_, ok := state["max_clusters_per_user"]
	assert.False(t, ok)

	state, err = removeZeroMaxClustersPerUser(context.Background(),
		map[string]any{
			"max_clusters_per_user": 1,
		}, nil)
	assert.NoError(t, err)
	_, ok = state["max_clusters_per_user"]
	assert.True(t, ok)
}
