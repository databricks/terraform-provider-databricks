package policies

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	clusterpolicies "github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceClusterPolicy(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Response: clusterpolicies.ListPoliciesResponse{
					Policies: []clusterpolicies.Policy{
						{
							PolicyId:   "abc",
							Name:       "policy",
							Definition: `{"abc":"123"}`,
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
		"id":         "abc",
		"definition": `{"abc":"123"}`,
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
				Response: clusterpolicies.ListPoliciesResponse{},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceClusterPolicy(),
		ID:          ".",
		HCL:         `name = "policy"`,
	}.ExpectError(t, "Policy named 'policy' does not exist")
}
