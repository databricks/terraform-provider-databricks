package exporter

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseInstanceName(t *testing.T) {
	ic := importContextForTest()
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}, map[string]any{
		"id":   "test-instance",
		"name": "test-instance",
	})
	d.SetId("test-instance")

	// Create wrapper for the resource data
	wrapper := &SDKv2ResourceData{
		data:   d,
		schema: &schema.Resource{Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}}},
	}

	name := resourcesMap["databricks_database_instance"].NameUnified(ic, wrapper)
	assert.Equal(t, "test-instance", name)
}

func TestDatabaseInstanceImport(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("lakebase,access")
	ic.meAdmin = true
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}, map[string]any{
		"id":   "test-instance",
		"name": "test-instance",
	})
	d.SetId("test-instance")
	r := &resource{
		ID:   "test-instance",
		Name: "test-instance",
		Data: d,
	}
	err := resourcesMap["databricks_database_instance"].Import(ic, r)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 1)
	assert.True(t, ic.testEmits["databricks_permissions[database_instance_test-instance] (id: /database-instances/test-instance)"])
}

func TestDatabaseInstanceIgnore(t *testing.T) {
	ic := importContextForTest()

	// Test with empty name - should be ignored
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]any{})
	d.SetId("test-instance")
	r := &resource{
		ID:   "test-instance",
		Data: d,
	}
	ignore := resourcesMap["databricks_database_instance"].Ignore(ic, r)
	assert.True(t, ignore)

	// Test with valid name - should not be ignored
	d2 := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]any{
		"name": "test-instance",
	})
	d2.SetId("test-instance")
	r2 := &resource{
		ID:   "test-instance",
		Data: d2,
	}
	ignore2 := resourcesMap["databricks_database_instance"].Ignore(ic, r2)
	assert.False(t, ignore2)
}

func TestDatabaseInstanceExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances?",
			Response: database.ListDatabaseInstancesResponse{
				DatabaseInstances: []database.DatabaseInstance{
					{
						Name:                      "prod-instance",
						Capacity:                  "CU_2",
						State:                     "AVAILABLE",
						NodeCount:                 2,
						EnableReadableSecondaries: true,
						UsagePolicyId:             "policy-123",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/database/instances/prod-instance?",
			Response: database.DatabaseInstance{
				Name:                               "prod-instance",
				Capacity:                           "CU_2",
				EffectiveCapacity:                  "CU_2",
				State:                              "AVAILABLE",
				NodeCount:                          2,
				EffectiveNodeCount:                 2,
				EnableReadableSecondaries:          true,
				EffectiveEnableReadableSecondaries: true,
				UsagePolicyId:                      "policy-123",
				EffectiveUsagePolicyId:             "policy-123",
				EffectiveCustomTags: []database.CustomTag{
					{
						Key:   "Environment",
						Value: "Production",
					},
					{
						Key:   "Team",
						Value: "DataPlatform",
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/permissions/database-instances/prod-instance?",
			ReuseRequest: true,
			Response: entity.PermissionsEntity{
				ObjectType:        "database-instances",
				AccessControlList: []iam.AccessControlRequest{},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("lakebase")
		ic.enableServices("lakebase")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the database instance was generated in the Terraform code
		content, err := os.ReadFile(tmpDir + "/lakebase.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with expected fields
		assert.Contains(t, contentStr, `resource "databricks_database_instance" "prod_instance"`)
		assert.Contains(t, contentStr, `name = "prod-instance"`)
		assert.Contains(t, contentStr, `capacity = "CU_2"`)
		// These simple-type fields are automatically exported from their effective_* counterparts
		assert.Contains(t, contentStr, `node_count = 2`)
		assert.Contains(t, contentStr, `enable_readable_secondaries = true`)
		assert.Contains(t, contentStr, `usage_policy_id = "policy-123"`)
		// Note: Complex types like custom_tags require deeper Plugin Framework integration
		// The conversion logic exists but wrapper.Set() validation is blocking it
		// This is documented in COMPLEX_TYPES_HANDLING.md as a known limitation
	})
}
