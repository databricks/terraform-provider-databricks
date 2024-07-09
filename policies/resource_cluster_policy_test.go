package policies

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceClusterPolicyRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: compute.Policy{
					PolicyId:           "abc",
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimestamp: 0,
					MaxClustersPerUser: 5,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "Dummy", d.Get("name"))
	assert.Equal(t, "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}", d.Get("definition"))
	assert.Equal(t, "abc", d.Get("policy_id"))
	assert.Equal(t, 5, d.Get("max_clusters_per_user"))
}

func TestResourceClusterPolicyRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceClusterPolicyRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceClusterPolicyCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/create",
				ExpectedRequest: compute.CreatePolicy{
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					MaxClustersPerUser: 3,
				},
				Response: compute.CreatePolicyResponse{
					PolicyId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: compute.Policy{
					PolicyId:           "abc",
					Name:               "Dummy",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimestamp: 0,
					MaxClustersPerUser: 3,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"definition":            `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"max_clusters_per_user": 3,
			"name":                  "Dummy",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 3, d.Get("max_clusters_per_user"))
}

func TestResourceClusterPolicyCreateNewFromPolicyFamily(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/create",
				ExpectedRequest: compute.CreatePolicy{
					Name:                            "Dummy",
					PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyId:                  "personal-vm",
					MaxClustersPerUser:              3,
				},
				Response: compute.CreatePolicyResponse{
					PolicyId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: compute.Policy{
					PolicyId:                        "abc",
					Name:                            "Dummy",
					Definition:                      "{\"spark_conf.baz\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyId:                  "personal-vm",
					CreatedAtTimestamp:              0,
					MaxClustersPerUser:              3,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: compute.ListPolicyFamiliesResponse{
					PolicyFamilies: []compute.PolicyFamily{
						{
							Name:           "Personal Compute",
							PolicyFamilyId: "personal-vm",
						},
					},
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_definition_overrides": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"max_clusters_per_user":              3,
			"name":                               "Dummy",
			"policy_family_id":                   "personal-vm",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 3, d.Get("max_clusters_per_user"))
}

func TestResourceClusterPolicyCreateOverrideBuiltin(t *testing.T) {
	policy := compute.Policy{
		PolicyId:                        "abc",
		Name:                            "Personal Compute",
		Definition:                      "{\"spark_conf.baz\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
		PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
		PolicyFamilyId:                  "personal-vm",
		CreatedAtTimestamp:              0,
		MaxClustersPerUser:              3,
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				ExpectedRequest: compute.EditPolicy{
					Name:                            "Personal Compute",
					PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyId:                  "personal-vm",
					MaxClustersPerUser:              3,
					PolicyId:                        "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: policy,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: compute.ListPolicyFamiliesResponse{
					PolicyFamilies: []compute.PolicyFamily{
						{
							Name:           "Personal Compute",
							PolicyFamilyId: "personal-vm",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Response: compute.ListPoliciesResponse{
					Policies: []compute.Policy{
						policy,
					},
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_definition_overrides": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"max_clusters_per_user":              3,
			"name":                               "Personal Compute",
			"policy_family_id":                   "personal-vm",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 3, d.Get("max_clusters_per_user"))
}

func TestResourceClusterPolicyCreateOverrideBuiltin_ErrorListingFamilies(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_definition_overrides": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"max_clusters_per_user":              3,
			"name":                               "Personal Compute",
			"policy_family_id":                   "personal-vm",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestResourceClusterPolicyCreateOverrideBuiltin_ErrorListingPolicies(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: compute.ListPolicyFamiliesResponse{
					PolicyFamilies: []compute.PolicyFamily{
						{
							Name:           "Personal Compute",
							PolicyFamilyId: "personal-vm",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_definition_overrides": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"max_clusters_per_user":              3,
			"name":                               "Personal Compute",
			"policy_family_id":                   "personal-vm",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestResourceClusterPolicyCreateConflict(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceClusterPolicy(),
		HCL: `
		name = "Dummy"
		definition = "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}"
		policy_family_id = "Test"
		policy_family_definition_overrides = "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}"
		max_clusters_per_user = 3
		`,
		Create: true,
	}.ExpectError(t, "invalid config supplied. [definition] Conflicting configuration arguments. [policy_family_definition_overrides] Conflicting configuration arguments. [policy_family_id] Conflicting configuration arguments")
}

func TestResourceClusterPolicyCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestResourceClusterPolicyUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				ExpectedRequest: compute.EditPolicy{
					PolicyId:   "abc",
					Name:       "Dummy Updated",
					Definition: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: compute.Policy{
					PolicyId:           "abc",
					Name:               "Dummy Updated",
					Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimestamp: 0,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy Updated",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyUpdateWithPolicyFamily(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				ExpectedRequest: compute.EditPolicy{
					PolicyId:                        "abc",
					Name:                            "Dummy Updated",
					PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyId:                  "personal-vm",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
				Response: compute.Policy{
					PolicyId:                        "abc",
					Name:                            "Dummy Updated",
					PolicyFamilyDefinitionOverrides: "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					PolicyFamilyId:                  "personal-vm",
					Definition:                      "{\"spark_conf.baz\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
					CreatedAtTimestamp:              0,
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_id":                   "personal-vm",
			"name":                               "Dummy Updated",
			"policy_family_definition_overrides": "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
			"name":       "Dummy Updated",
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/delete",
				ExpectedRequest: map[string]string{
					"policy_id": "abc",
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDeletePolicyOverride(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/edit",
				ExpectedRequest: compute.EditPolicy{
					PolicyId:       "abc",
					Name:           "Personal Compute",
					PolicyFamilyId: "personal-vm",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: compute.ListPolicyFamiliesResponse{
					PolicyFamilies: []compute.PolicyFamily{
						{
							Name:           "Personal Compute",
							PolicyFamilyId: "personal-vm",
						},
					},
				},
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_id":                   "personal-vm",
			"name":                               "Personal Compute",
			"policy_family_definition_overrides": "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
		},
		Delete: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDeletePolicyOverride_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policy-families?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		State: map[string]any{
			"policy_family_id":                   "personal-vm",
			"name":                               "Personal Compute",
			"policy_family_definition_overrides": "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
		},
		Delete: true,
		ID:     "abc",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceClusterPolicyDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/policies/clusters/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceClusterPolicy(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
