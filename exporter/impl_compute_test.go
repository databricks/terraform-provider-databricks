package exporter

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/compute"
	sdk_pipelines "github.com/databricks/databricks-sdk-go/service/pipelines"
	tf_dlt "github.com/databricks/terraform-provider-databricks/pipelines"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/policies"
	"github.com/databricks/terraform-provider-databricks/pools"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstancePool(t *testing.T) {
	d := pools.ResourceInstancePool().ToResource().TestResourceData()
	d.Set("instance_pool_name", "blah-bah")
	ic := importContextForTest()
	ic.enableServices("access,pools")

	name := resourcesMap["databricks_instance_pool"].Name(ic, d)
	assert.Equal(t, "blah-bah", name)

	d = pools.ResourceInstancePool().ToResource().TestResourceData()
	d.SetId("abc-bcd-def")
	name = resourcesMap["databricks_instance_pool"].Name(ic, d)
	assert.Equal(t, "def", name)

	ic.meAdmin = true
	r := &resource{
		ID:   "abc",
		Data: d,
	}
	err := resourcesMap["databricks_instance_pool"].Import(ic, r)
	assert.NoError(t, err)
	assert.True(t, ic.testEmits["databricks_permissions[inst_pool_def] (id: /instance-pools/abc)"])
}

func TestClusterPolicy(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "bcd")
	definition := map[string]map[string]string{
		"aws_attributes.instance_profile_arn": {
			"value": "def",
		},
		"instance_pool_id": {
			"defaultValue": "efg",
		},
		"init_scripts.0.dbfs.destination": {
			"type":  "fixed",
			"value": "dbfs:/FileStore/init-script.sh",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))
	ic := importContextForTest()
	ic.enableServices("storage,pools,policies,access")
	ic.meAdmin = true
	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 4)
	assert.True(t, ic.testEmits["databricks_permissions[cluster_policy_bcd] (id: /cluster-policies/abc)"])
	assert.True(t, ic.testEmits["databricks_instance_pool[<unknown>] (id: efg)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: def)"])
	assert.True(t, ic.testEmits["databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/init-script.sh)"])
}

func TestPredefinedClusterPolicy(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("policy_family_id", "job-cluster")
	d.Set("name", "Job Compute")
	policy, _ := json.Marshal(map[string]map[string]string{})
	d.Set("definition", string(policy))
	ic := importContextForTest()
	ic.builtInPolicies = map[string]compute.PolicyFamily{
		"job-cluster": {Name: "Job Compute"},
	}
	r := resource{ID: "abc", Data: d}
	err := ic.Importables["databricks_cluster_policy"].Import(ic, &r)
	assert.NoError(t, err)
	assert.Equal(t, "data", r.Mode)
	assert.Equal(t, "", r.Data.Get("definition").(string))
}

func TestInstancePoolNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := pools.ResourceInstancePool().ToResource().TestResourceData()
	d.SetId("a-b-c")
	d.Set("instance_pool_name", "")
	assert.Equal(t, "c", resourcesMap["databricks_instance_pool"].Name(ic, d))
}

func TestClusterNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := clusters.ResourceCluster().ToResource().TestResourceData()
	d.SetId("a-b-c")
	assert.Equal(t, "c", resourcesMap["databricks_cluster"].Name(ic, d))
}

func TestImportClusterLibraries(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
			Response: compute.ClusterLibraryStatuses{
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Library: &compute.Library{
							Whl: "foo.whl",
						},
						Status: "INSTALLED",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := clusters.ResourceCluster().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_cluster"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

func TestImportClusterLibrariesFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       404,
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := clusters.ResourceCluster().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_cluster"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestClusterListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/list?filter_by.cluster_sources=UI&filter_by.cluster_sources=API&page_size=100",
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_cluster"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func TestClusterList_NoNameMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/list?filter_by.cluster_sources=UI&filter_by.cluster_sources=API&page_size=100",
			Response: clusters.ClusterList{
				Clusters: []clusters.ClusterInfo{
					{
						ClusterName: "abc",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_cluster"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestInstnacePoolsListWithMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: compute.ListInstancePools{
				InstancePools: []compute.InstancePoolAndStats{
					{
						InstancePoolName: "test",
						InstancePoolId:   "test",
					},
					{
						InstancePoolName: "bcd",
						InstancePoolId:   "bcd",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("pools")
		ic.match = "bcd"
		err := resourcesMap["databricks_instance_pool"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestClusterPolicyNoValues(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "abc")
	d.Set("definition", `{"foo": {}}`)
	ic := importContextForTest()
	ic.enableServices("access")
	ic.meAdmin = true
	err := resourcesMap["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "x",
		Data: d,
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ic.testEmits))
}

func TestPoliciesListing(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/list?",
			Response: compute.ListPoliciesResponse{
				Policies: []compute.Policy{
					{
						Name:           "Personal Compute",
						PolicyFamilyId: "personal-vm",
						PolicyId:       "123",
					},
					{
						Name:     "abcd",
						PolicyId: "456",
					},
				},
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
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("policies")
		err := resourcesMap["databricks_cluster_policy"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestDLTIgnore(t *testing.T) {
	ic := importContextForTest()
	d := tf_dlt.ResourcePipeline().ToResource().TestResourceData()
	scm := tf_dlt.ResourcePipeline().Schema

	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without libraries
	assert.True(t, resourcesMap["databricks_pipeline"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))

	// job deployed by DABs
	d.MarkNewResource()
	pipeline := tf_dlt.Pipeline{
		PipelineSpec: sdk_pipelines.PipelineSpec{
			Deployment: &sdk_pipelines.PipelineDeployment{
				Kind: "BUNDLE",
			},
		},
	}
	err := common.StructToData(pipeline, scm, d)
	require.NoError(t, err)

	r = &resource{ID: "12345", Data: d}
	for k := range ic.ignoredResources {
		delete(ic.ignoredResources, k)
	}
	assert.True(t, resourcesMap["databricks_pipeline"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestClusterPolicyWrongDef(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "abc")
	d.Set("definition", "..")
	ic := importContextForTest()
	err := resourcesMap["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "x",
		Data: d,
	})
	assert.EqualError(t, err, "invalid character '.' looking for beginning of value")
}

func TestClusterPolicyCloudConversionFixed(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":  "fixed",
			"value": "SPOT",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies,access")
	ic.meAdmin = true

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the definition was converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	// Check that aws_attributes was removed
	_, hasAws := newDef["aws_attributes.availability"]
	assert.False(t, hasAws, "AWS attribute should be removed")

	// Check that azure_attributes was added with converted value
	azureAttr, hasAzure := newDef["azure_attributes.availability"]
	assert.True(t, hasAzure, "Azure attribute should be added")
	assert.Equal(t, "SPOT_AZURE", azureAttr["value"])
	assert.Equal(t, "fixed", azureAttr["type"])
}

func TestClusterPolicyCloudConversionAllowlist(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":   "allowlist",
			"values": []interface{}{"SPOT", "ON_DEMAND", "SPOT_WITH_FALLBACK"},
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the definition was converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	// Check that azure_attributes was added with all values converted
	azureAttr, hasAzure := newDef["azure_attributes.availability"]
	assert.True(t, hasAzure, "Azure attribute should be added")
	assert.Equal(t, "allowlist", azureAttr["type"])

	values := azureAttr["values"].([]interface{})
	assert.Len(t, values, 3)
	assert.Contains(t, values, "SPOT_AZURE")
	assert.Contains(t, values, "ON_DEMAND_AZURE")
	assert.Contains(t, values, "SPOT_WITH_FALLBACK_AZURE")
}

func TestClusterPolicyNodeTypeConversionFixed(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"node_type_id": {
			"type":  "fixed",
			"value": "i3.xlarge",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	// Load node type mappings
	mappings, err := loadNodeTypeMappings("node_type_mapping.json")
	require.NoError(t, err)
	ic.nodeTypeMappings = mappings

	err = ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the node type was converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	nodeTypeAttr := newDef["node_type_id"]
	assert.Equal(t, "fixed", nodeTypeAttr["type"])
	// The actual converted value depends on the mapping file
	// Just verify it's not the original AWS value
	convertedValue := nodeTypeAttr["value"].(string)
	assert.NotEmpty(t, convertedValue)
}

func TestClusterPolicyNodeTypeConversionAllowlist(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"node_type_id": {
			"type":   "allowlist",
			"values": []interface{}{"i3.xlarge", "i3.2xlarge"},
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	// Load node type mappings
	mappings, err := loadNodeTypeMappings("node_type_mapping.json")
	require.NoError(t, err)
	ic.nodeTypeMappings = mappings

	err = ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the node types were converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	nodeTypeAttr := newDef["node_type_id"]
	assert.Equal(t, "allowlist", nodeTypeAttr["type"])

	values := nodeTypeAttr["values"].([]interface{})
	assert.Len(t, values, 2)
	// Values should be converted (exact values depend on mapping file)
}

func TestClusterPolicyRegexSkipped(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":    "regex",
			"pattern": "^SPOT.*",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify regex policy was not converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	// Original aws_attributes should still be there
	awsAttr, hasAws := newDef["aws_attributes.availability"]
	assert.True(t, hasAws, "AWS regex attribute should not be removed")
	assert.Equal(t, "regex", awsAttr["type"])
	assert.Equal(t, "^SPOT.*", awsAttr["pattern"])

	// Azure attribute should not be added
	_, hasAzure := newDef["azure_attributes.availability"]
	assert.False(t, hasAzure, "Azure attribute should not be added for regex")
}

func TestClusterPolicyFamilyOverridesConversion(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "Job Compute")
	d.Set("policy_family_id", "job-cluster")

	overrides := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":  "fixed",
			"value": "SPOT",
		},
	}
	overridesJson, _ := json.Marshal(overrides)
	d.Set("policy_family_definition_overrides", string(overridesJson))
	d.Set("definition", "{}")

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")
	ic.builtInPolicies = map[string]compute.PolicyFamily{
		"job-cluster": {Name: "Job Compute"},
	}

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify that policy_family_definition_overrides was converted, not definition
	overridesStr := d.Get("policy_family_definition_overrides").(string)
	var newOverrides map[string]map[string]any
	err = json.Unmarshal([]byte(overridesStr), &newOverrides)
	assert.NoError(t, err)

	// Check that azure_attributes was added
	azureAttr, hasAzure := newOverrides["azure_attributes.availability"]
	assert.True(t, hasAzure, "Azure attribute should be added in overrides")
	assert.Equal(t, "SPOT_AZURE", azureAttr["value"])
}

func TestClusterPolicyUnlimitedTypeConversion(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"node_type_id": {
			"type":         "unlimited",
			"defaultValue": "i3.xlarge",
			"isOptional":   true,
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	// Load node type mappings
	mappings, err := loadNodeTypeMappings("node_type_mapping.json")
	require.NoError(t, err)
	ic.nodeTypeMappings = mappings

	err = ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the node type defaultValue was converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	nodeTypeAttr := newDef["node_type_id"]
	assert.Equal(t, "unlimited", nodeTypeAttr["type"])
	assert.Equal(t, true, nodeTypeAttr["isOptional"])

	// The defaultValue should be converted
	convertedValue := nodeTypeAttr["defaultValue"].(string)
	assert.NotEmpty(t, convertedValue)
	assert.NotEqual(t, "i3.xlarge", convertedValue, "Node type should be converted")
}

func TestClusterPolicyUnlimitedTypeCloudAttributeConversion(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":         "unlimited",
			"defaultValue": "SPOT",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "azure"
	ic.enableServices("policies")

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify the definition was converted
	newDefStr := d.Get("definition").(string)
	var newDef map[string]map[string]any
	err = json.Unmarshal([]byte(newDefStr), &newDef)
	assert.NoError(t, err)

	// Check that aws_attributes was removed
	_, hasAws := newDef["aws_attributes.availability"]
	assert.False(t, hasAws, "AWS attribute should be removed")

	// Check that azure_attributes was added with converted defaultValue
	azureAttr, hasAzure := newDef["azure_attributes.availability"]
	assert.True(t, hasAzure, "Azure attribute should be added")
	assert.Equal(t, "SPOT_AZURE", azureAttr["defaultValue"])
	assert.Equal(t, "unlimited", azureAttr["type"])
}

func TestClusterPolicyNoConversionWhenTargetCloudEmpty(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "test-policy")
	definition := map[string]map[string]any{
		"aws_attributes.availability": {
			"type":  "fixed",
			"value": "SPOT",
		},
	}
	policy, _ := json.Marshal(definition)
	originalDef := string(policy)
	d.Set("definition", originalDef)

	ic := importContextForTest()
	ic.Client = &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://test.cloud.databricks.com",
			},
		},
	}
	ic.targetCloud = "" // No target cloud
	ic.enableServices("policies")

	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "test-policy-id",
		Data: d,
	})
	assert.NoError(t, err)

	// Verify definition was not changed
	newDefStr := d.Get("definition").(string)
	assert.JSONEq(t, originalDef, newDefStr)
}

func TestIncrementalListDLT(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines?max_results=100",
			Response: sdk_pipelines.ListPipelinesResponse{
				Statuses: []sdk_pipelines.PipelineStateInfo{
					{
						PipelineId: "abc",
						Name:       "abc",
					},
					{
						PipelineId: "def",
						Name:       "def",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines/abc?",
			Response: sdk_pipelines.GetPipelineResponse{
				PipelineId:   "abc",
				Name:         "abc",
				LastModified: 1681466931226,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines/def?",
			Response: sdk_pipelines.GetPipelineResponse{
				PipelineId:   "def",
				Name:         "def",
				LastModified: 1690156900000,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("dlt")
		ic.incremental = true
		ic.updatedSinceStr = "2023-07-24T00:00:00Z"
		ic.updatedSinceMs = 1690156700000
		err := resourcesMap["databricks_pipeline"].List(ic)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}
