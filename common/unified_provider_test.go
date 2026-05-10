package common

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkspaceIDValidateFunc(t *testing.T) {
	validateFunc := workspaceIDValidateFunc()

	testCases := []struct {
		name        string
		input       interface{}
		expectError bool
	}{
		{
			name:        "valid numeric workspace ID",
			input:       "123456789",
			expectError: false,
		},
		{
			name:        "valid single digit workspace ID",
			input:       "1",
			expectError: false,
		},
		{
			name:        "valid large workspace ID",
			input:       "999999999999999",
			expectError: false,
		},
		{
			name:        "invalid empty string",
			input:       "",
			expectError: true,
		},
		{
			name:        "invalid non-numeric string",
			input:       "abc123",
			expectError: true,
		},
		{
			name:        "invalid string with spaces",
			input:       "123 456",
			expectError: true,
		},
		{
			name:        "invalid string with special characters",
			input:       "123-456",
			expectError: true,
		},
		{
			name:        "invalid string with leading zero",
			input:       "0123",
			expectError: true, // Leading zeros are not allowed
		},
		{
			name:        "invalid single zero",
			input:       "0",
			expectError: true, // Zero is not a valid workspace ID
		},
		{
			name:        "invalid negative number",
			input:       "-123",
			expectError: true,
		},
		{
			name:        "invalid decimal number",
			input:       "123.456",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, errors := validateFunc(tc.input, "workspace_id")
			if tc.expectError {
				assert.NotEmpty(t, errors, "Expected validation errors but got none")
			} else {
				assert.Empty(t, errors, "Expected no validation errors but got: %v", errors)
			}
		})
	}
}

func TestNamespaceCustomizeSchema(t *testing.T) {
	// Create a test schema with provider_config structure
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}

	// Apply the customization
	NamespaceCustomizeSchema(&CustomizableSchema{
		Schema: &schema.Schema{
			Elem: &schema.Resource{
				Schema: testSchema,
			},
		},
	})

	// Verify that ValidateFunc was set
	providerConfig := testSchema["provider_config"]
	require.NotNil(t, providerConfig)
	elem, ok := providerConfig.Elem.(*schema.Resource)
	require.True(t, ok)
	workspaceID := elem.Schema["workspace_id"]
	require.NotNil(t, workspaceID)
	assert.NotNil(t, workspaceID.ValidateFunc, "ValidateFunc should be set on workspace_id")
}

func TestNamespaceCustomizeSchemaMap(t *testing.T) {
	t.Run("with valid provider_config", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"provider_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workspace_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"other_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}

		result := NamespaceCustomizeSchemaMap(testSchema)

		// Verify that ValidateFunc was set
		providerConfig := result["provider_config"]
		require.NotNil(t, providerConfig)
		elem, ok := providerConfig.Elem.(*schema.Resource)
		require.True(t, ok)
		workspaceID := elem.Schema["workspace_id"]
		require.NotNil(t, workspaceID)
		assert.NotNil(t, workspaceID.ValidateFunc, "ValidateFunc should be set on workspace_id")
	})

	t.Run("panic without provider_config", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"other_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}

		assert.Panics(t, func() {
			NamespaceCustomizeSchemaMap(testSchema)
		}, "NamespaceCustomizeSchemaMap should panic when provider_config is not present")
	})
}

func TestAddNamespaceInSchema(t *testing.T) {
	t.Run("adds provider_config to schema with existing fields", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}

		result := AddNamespaceInSchema(testSchema)

		// Verify provider_config was added
		require.NotNil(t, result)
		providerConfig, exists := result["provider_config"]
		assert.True(t, exists, "provider_config should exist in schema")
		require.NotNil(t, providerConfig)

		// Verify workspace_id field
		elem, ok := providerConfig.Elem.(*schema.Resource)
		require.True(t, ok, "Elem should be *schema.Resource")
		workspaceID, exists := elem.Schema["workspace_id"]
		assert.True(t, exists, "workspace_id should exist")
		require.NotNil(t, workspaceID)
		assert.Equal(t, schema.TypeString, workspaceID.Type)
		assert.True(t, workspaceID.Optional)
		assert.True(t, workspaceID.Computed)

		// Verify existing fields are preserved
		assert.Len(t, result, 3, "Should have 3 fields: name, description, provider_config")
		assert.NotNil(t, result["name"])
		assert.NotNil(t, result["description"])
		assert.NotNil(t, result["provider_config"])
	})

	t.Run("panics when provider_config already exists", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"provider_config": {
				Type:     schema.TypeList,
				Optional: true,
			},
		}

		assert.Panics(t, func() {
			AddNamespaceInSchema(testSchema)
		}, "Should panic when provider_config already exists in schema")
	})
}

func TestWorkspaceClientUnifiedProvider(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	ctx := context.Background()
	mockWorkspaceClient := &databricks.WorkspaceClient{}

	testCases := []struct {
		name          string
		resourceData  map[string]interface{}
		client        *DatabricksClient
		expectError   bool
		errorContains string
		description   string
	}{
		{
			name: "workspace_id not set - calls with empty string",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     0,
			},
			expectError: false,
			description: "When provider_config is not set, should use cached workspace client",
		},
		{
			name: "workspace_id set to valid value",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "123456",
					},
				},
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     123456,
			},
			expectError: false,
			description: "When workspace_id matches cached ID, should return workspace client",
		},
		{
			name: "workspace_id set to empty string",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "",
					},
				},
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     0,
			},
			expectError: false,
			description: "When workspace_id is explicitly empty, should use cached workspace client",
		},
		{
			name: "workspace_id with different numeric value",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "789012",
					},
				},
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     1234,
			},
			expectError:   true,
			errorContains: "failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1234 but got 789012 in provider_config",
			description:   "Should handle different workspace IDs correctly",
		},
		{
			name: "account level provider without workspace_id - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:      "https://accounts.cloud.databricks.com",
						AccountID: "test-account-id",
						Token:     "test-token",
					},
				},
			},
			expectError:   true,
			errorContains: "managing workspace-level resources requires a workspace_id, but none was found in the resource's provider_config block or the provider's workspace_id attribute",
			description:   "Account-level provider requires workspace_id to be set",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

			// Call WorkspaceClientUnifiedProvider
			result, err := tc.client.WorkspaceClientUnifiedProvider(ctx, d)

			// Verify results
			if tc.expectError {
				assert.Error(t, err, tc.description)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
			} else {
				assert.NoError(t, err, tc.description)
				assert.NotNil(t, result)
				assert.Equal(t, mockWorkspaceClient, result)
			}
		})
	}
}

func TestDatabricksClientForUnifiedProvider(t *testing.T) {
	cachedWorkspaceHost := "https://workspace.test.databricks.com"
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	ctx := context.Background()

	testCases := []struct {
		name             string
		resourceData     map[string]interface{}
		client           *DatabricksClient
		expectError      bool
		errorContains    string
		expectSameClient bool
		description      string
	}{
		{
			name: "workspace_id not set - returns current client",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
			},
			expectError:      false,
			expectSameClient: true,
			description:      "When provider_config is not set, should return current client",
		},
		{
			name: "workspace_id set to empty string - returns current client",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "",
					},
				},
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
			},
			expectError:      false,
			expectSameClient: true,
			description:      "When workspace_id is explicitly empty, should return current client",
		},
		{
			name: "workspace_id set to valid value - client from cache",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "123456",
					},
				},
			},
			client: func() *DatabricksClient {
				c := &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:  "https://test.cloud.databricks.com",
							Token: "test-token",
						},
					},
				}
				// Setup cached DatabricksClient
				mockDatabricksClient := &client.DatabricksClient{
					Config: &config.Config{
						Host: cachedWorkspaceHost,
					},
				}
				c.cachedDatabricksClients = map[int64]*client.DatabricksClient{
					123456: mockDatabricksClient,
				}
				return c
			}(),
			expectError:      false,
			expectSameClient: false,
			description:      "When workspace_id is set and client is cached, should return cached client",
		},
		{
			name: "workspace_id set to valid value - creates new client",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "789012",
					},
				},
			},
			client: func() *DatabricksClient {
				c := &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:  "https://test.cloud.databricks.com",
							Token: "test-token",
						},
					},
				}
				c.Config = c.Config.WithTesting()
				// Setup workspace client for the call
				mockWorkspaceClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  cachedWorkspaceHost,
						Token: "test-token",
					},
				}
				mockWorkspaceClient.Config = mockWorkspaceClient.Config.WithTesting()
				c.SetWorkspaceClient(mockWorkspaceClient)
				c.cachedWorkspaceID = 789012
				return c
			}(),
			expectError:      false,
			expectSameClient: false,
			description:      "When workspace_id is set and client is not cached, should create new client",
		},
		{
			name: "invalid workspace_id - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "invalid",
					},
				},
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:  "https://test.cloud.databricks.com",
						Token: "test-token",
					},
				},
			},
			expectError:   true,
			errorContains: "failed to parse workspace_id",
			description:   "When workspace_id is invalid, should return error",
		},
		{
			name: "account level provider without workspace_id - returns current client",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:      "https://accounts.cloud.databricks.com",
						AccountID: "test-account-id",
						Token:     "test-token",
					},
				},
			},
			expectError:      false,
			expectSameClient: true,
			description:      "Account-level provider without workspace_id returns current client for AccountOrWorkspaceRequest routing",
		},
		{
			name: "workspace_id mismatch - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "100",
					},
				},
			},
			client: func() *DatabricksClient {
				c := &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:  "https://test.cloud.databricks.com",
							Token: "test-token",
						},
					},
				}
				c.Config = c.Config.WithTesting()
				mockWorkspaceClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  cachedWorkspaceHost,
						Token: "test-token",
					},
				}
				mockWorkspaceClient.Config = mockWorkspaceClient.Config.WithTesting()
				c.SetWorkspaceClient(mockWorkspaceClient)
				c.cachedWorkspaceID = 200
				return c
			}(),
			expectError:   true,
			errorContains: "workspace_id mismatch",
			description:   "Should return error when workspace_id doesn't match configured workspace",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

			// Call DatabricksClientForUnifiedProvider
			result, err := tc.client.DatabricksClientForUnifiedProvider(ctx, d)

			// Verify results
			if tc.expectError {
				assert.Error(t, err, tc.description)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err, tc.description)
				assert.NotNil(t, result)
				assert.NotNil(t, result.DatabricksClient)

				if tc.expectSameClient {
					// Verify we got the same client back
					assert.Equal(t, tc.client, result, "Should return the same client instance")
				} else {
					// verify the host is the same as the one we get from the cached client
					assert.Equal(t, cachedWorkspaceHost, result.DatabricksClient.Config.Host)
				}
			}
		})
	}
}

// newTestResourceForCustomizeDiff creates a Resource with provider_config schema
// and NamespaceCustomizeDiff for testing the customize diff logic.
func newTestResourceForCustomizeDiff() *schema.Resource {
	r := Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		CustomizeDiff: NamespaceCustomizeDiff,
		Read: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			return nil
		},
	}
	AddNamespaceInSchema(r.Schema)
	NamespaceCustomizeSchemaMap(r.Schema)
	return r.ToResource()
}

// diffCustomizeDiff runs Diff on the resource with the given state and config,
// returning the diff and any error from CustomizeDiff.
//
// Uses NewResourceConfigShimmed with a cty value so that GetRawConfigAt works
// for tests of NamespaceValidateWorkspaceID (which inspects raw config to
// distinguish "user typed it" from "preserved from state"). When the test
// resource has the api field (dual resources), it is included in the cty
// schema; otherwise only name + provider_config are populated.
func diffCustomizeDiff(
	t *testing.T,
	resource *schema.Resource,
	instanceState map[string]string,
	rawConfig map[string]interface{},
	c *DatabricksClient,
) (*terraform.InstanceDiff, error) {
	t.Helper()
	ctx := context.Background()
	ctx = context.WithValue(ctx, ResourceName, "test_resource")

	ctyVal := testResourceConfigToCtyValue(resource, rawConfig)
	block := schema.InternalMap(resource.Schema).CoreConfigSchema()
	rc := terraform.NewResourceConfigShimmed(ctyVal, block)

	// Always pass an InstanceState with RawConfig set so that
	// ResourceDiff.GetRawConfigAt — used by NamespaceValidateWorkspaceID
	// to detect explicit user input — sees the cty value. Mirrors the gRPC
	// PlanResourceChange path where priorState.RawConfig = configVal.
	is := &terraform.InstanceState{RawConfig: ctyVal}
	if instanceState != nil {
		is.Attributes = instanceState
	}
	return resource.Diff(ctx, is, rc, c)
}

// testResourceConfigToCtyValue converts a Go map representing a test config to
// a cty.Value matching the resource's schema. It handles the provider_config
// TypeList block and the optional api field used by dual resources.
func testResourceConfigToCtyValue(resource *schema.Resource, rawConfig map[string]any) cty.Value {
	block := schema.InternalMap(resource.Schema).CoreConfigSchema()
	implType := block.ImpliedType()
	vals := map[string]cty.Value{}
	for name, attrType := range implType.AttributeTypes() {
		raw, ok := rawConfig[name]
		if !ok {
			vals[name] = cty.NullVal(attrType)
			continue
		}
		vals[name] = goValueToCty(raw, attrType)
	}
	return cty.ObjectVal(vals)
}

func goValueToCty(v any, t cty.Type) cty.Value {
	if v == nil {
		return cty.NullVal(t)
	}
	switch {
	case t.Equals(cty.String):
		return cty.StringVal(v.(string))
	case t.IsListType():
		items := v.([]any)
		if len(items) == 0 {
			return cty.ListValEmpty(t.ElementType())
		}
		elements := make([]cty.Value, 0, len(items))
		for _, item := range items {
			elements = append(elements, goValueToCty(item, t.ElementType()))
		}
		return cty.ListVal(elements)
	case t.IsObjectType():
		objMap := v.(map[string]any)
		objVals := map[string]cty.Value{}
		for fieldName, fieldType := range t.AttributeTypes() {
			if fv, ok := objMap[fieldName]; ok {
				objVals[fieldName] = goValueToCty(fv, fieldType)
			} else {
				objVals[fieldName] = cty.NullVal(fieldType)
			}
		}
		return cty.ObjectVal(objVals)
	default:
		return cty.NullVal(t)
	}
}

func TestNamespaceCustomizeDiff_MatchingWorkspaceID(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://test.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
		cachedWorkspaceClient: mockWS,
		cachedWorkspaceID:     123456,
	}
	diff, err := diffCustomizeDiff(t, resource, map[string]string{
		"name":                           "test",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "123456",
	}, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "123456",
			},
		},
	}, c)
	assert.NoError(t, err)
	assert.Nil(t, diff)
}

func TestNamespaceCustomizeDiff_MismatchedWorkspaceID(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://test.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
		cachedWorkspaceClient: mockWS,
		cachedWorkspaceID:     999999,
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "123",
			},
		},
	}, c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "workspace_id mismatch")
	assert.Contains(t, err.Error(), "please check the workspace_id provided in provider_config")
}

// TestNamespaceValidateWorkspaceID_DefersWhenConfigAbsent asserts the validator
// returns nil when the user did not write provider_config in HCL, even when
// the provider-level Config.WorkspaceID is set to a value that would mismatch
// the cached workspace ID. This is the v1.114-regression-fixing path: validators
// must not fall back to Config.WorkspaceID at plan time.
func TestNamespaceValidateWorkspaceID_DefersWhenConfigAbsent(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://test.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:        "https://test.cloud.databricks.com",
				Token:       "test-token",
				WorkspaceID: "123", // mismatching auto-resolved value
			},
		},
		cachedWorkspaceClient: mockWS,
		cachedWorkspaceID:     999999,
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		// no provider_config in raw config
	}, c)
	assert.NoError(t, err, "validator must defer when provider_config is not written by user")
}

// TestNamespaceValidateWorkspaceID_DefersWhenInnerEmpty asserts the validator
// returns nil when provider_config is written but workspace_id is empty —
// matches the schema-level "user wrote a placeholder block" case. The empty
// string is still a known value, but the validator treats it like an absent
// value because there is nothing concrete to validate.
//
// Note: cross-resource references (workspace_id = databricks_mws_workspaces.this.workspace_id)
// produce an *unknown* cty value at plan time. terraform-plugin-sdk's
// NewResourceConfigShimmed cannot synthesize cty.UnknownVal directly, but the
// underlying GetRawConfigAt code path is exercised by the IsKnown() check in
// workspaceIDFromRawDiffConfig, which has its own unit coverage.
func TestNamespaceValidateWorkspaceID_DefersWhenInnerEmpty(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://test.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:        "https://test.cloud.databricks.com",
				Token:       "test-token",
				WorkspaceID: "123", // mismatching auto-resolved value
			},
		},
		cachedWorkspaceClient: mockWS,
		cachedWorkspaceID:     999999,
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "",
			},
		},
	}, c)
	assert.NoError(t, err, "validator must defer when inner workspace_id is empty")
}

func TestNamespaceCustomizeDiff_AccountLevelProvider_ValidWorkspace(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://workspace.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "test-account-id",
				Token:     "test-token",
			},
		},
	}
	// Pre-cache the workspace client so WorkspaceClientForWorkspace returns it
	c.SetWorkspaceClientForWorkspace(123, mockWS)

	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "123",
			},
		},
	}, c)
	assert.NoError(t, err)
}

// unifiedHostConfig returns a config whose resolved host type is UnifiedHost.
// It installs a HostMetadataResolver that returns UnifiedHost and forces
// EnsureResolved to run so resolvedHostType is populated before HostType() is
// consulted.
func unifiedHostConfig(t *testing.T, host string) *config.Config {
	cfg := &config.Config{
		Host:  host,
		Token: "test-token",
		HostMetadataResolver: func(ctx context.Context, _ string) (*config.HostMetadata, error) {
			return &config.HostMetadata{HostType: config.UnifiedHost}, nil
		},
	}
	require.NoError(t, cfg.EnsureResolved())
	return cfg
}

func TestNamespaceCustomizeDiff_UnifiedHost_ValidWorkspace(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://workspace.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unified.cloud.databricks.com"),
		},
	}
	c.SetWorkspaceClientForWorkspace(456, mockWS)

	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "456",
			},
		},
	}, c)
	assert.NoError(t, err)
}

func TestNamespaceCustomizeDiff_UnifiedHost_DirectFallback(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unified.cloud.databricks.com"),
		},
	}
	// No cached workspace client — WorkspaceClientForWorkspace falls back to
	// tryWorkspaceClientDirect which succeeds for unified hosts (routes via
	// X-Databricks-Org-Id header). Actual workspace validation happens at apply time.
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "999",
			},
		},
	}, c)
	assert.NoError(t, err)
}

// newDualResourceForCustomizeDiff builds a test resource with the `api` field
// (mirroring dual resources that call AddApiField) wired to CustomizeDiffDualResources.
func newDualResourceForCustomizeDiff() *schema.Resource {
	r := Resource{
		IsDual: true,
		Schema: AddApiField(map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		}),
		CustomizeDiff: CustomizeDiffDualResources,
		Read: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			return nil
		},
	}
	AddNamespaceInSchema(r.Schema)
	NamespaceCustomizeSchemaMap(r.Schema)
	return r.ToResource()
}

func TestValidateApiLevelForUnifiedHost_ErrorsWhenApiMissing(t *testing.T) {
	resource := newDualResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unifiedhost.databricks.com"),
		},
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
	}, c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "please set api to account or workspace")
}

func TestValidateApiLevelForUnifiedHost_PassesWhenApiSet(t *testing.T) {
	resource := newDualResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unifiedhost.databricks.com"),
		},
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"api":  "account",
	}, c)
	assert.NoError(t, err)
}

func TestValidateApiLevelForUnifiedHostFromData_ErrorsWhenApiMissing(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unifiedhost.databricks.com"),
		},
	}
	d := schema.TestResourceDataRaw(t, AddApiField(map[string]*schema.Schema{}), map[string]interface{}{})
	err := ValidateApiLevelForUnifiedHostFromData(d, c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "please set api to account or workspace")
}

func TestValidateApiLevelForUnifiedHostFromData_PassesWhenApiSet(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: unifiedHostConfig(t, "https://unifiedhost.databricks.com"),
		},
	}
	d := schema.TestResourceDataRaw(t, AddApiField(map[string]*schema.Schema{}), map[string]interface{}{
		"api": "workspace",
	})
	assert.NoError(t, ValidateApiLevelForUnifiedHostFromData(d, c))
}

func TestValidateApiLevelForUnifiedHost_IgnoresNonUnifiedHost(t *testing.T) {
	resource := newDualResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
	}
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
	}, c)
	assert.NoError(t, err)
}

func TestNamespaceCustomizeDiff_ForceNewOnChange(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://test.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
		cachedWorkspaceClient: mockWS,
		cachedWorkspaceID:     789,
	}
	diff, err := diffCustomizeDiff(t, resource, map[string]string{
		"name":                           "test",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "456",
	}, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "789",
			},
		},
	}, c)
	assert.NoError(t, err)
	require.NotNil(t, diff)
	wsAttr, ok := diff.Attributes[workspaceIDSchemaKey]
	require.True(t, ok, "workspace_id should be in diff attributes")
	assert.True(t, wsAttr.RequiresNew, "changing workspace_id should require new resource")
}

func TestWorkspaceClientUnifiedProviderWithWorkspaceID(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	ctx := context.Background()
	mockWorkspaceClient := &databricks.WorkspaceClient{}

	testCases := []struct {
		name          string
		resourceData  map[string]interface{}
		client        *DatabricksClient
		expectError   bool
		errorContains string
		expectedHost  string // Host of the workspace client we expect to be returned
		description   string
	}{
		{
			name: "account-level with workspace_id and no provider_config - uses default",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: func() *DatabricksClient {
				c := &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:        "https://accounts.cloud.databricks.com",
							AccountID:   "test-account-id",
							Token:       "test-token",
							WorkspaceID: "123456",
						},
					},
				}
				c.Config = c.Config.WithTesting()
				// Create mock workspace client for DEFAULT workspace
				defaultWsClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  "https://default-workspace.test.databricks.com",
						Token: "test-token",
					},
				}
				defaultWsClient.Config = defaultWsClient.Config.WithTesting()
				// Create mock workspace client for OVERRIDE workspace
				overrideWsClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  "https://override-workspace.test.databricks.com",
						Token: "test-token",
					},
				}
				overrideWsClient.Config = overrideWsClient.Config.WithTesting()
				// Cache BOTH workspace clients
				c.cachedWorkspaceClients = map[int64]*databricks.WorkspaceClient{
					123456: defaultWsClient,  // workspace_id
					789012: overrideWsClient, // potential override
				}
				return c
			}(),
			expectError:  false,
			expectedHost: "https://default-workspace.test.databricks.com",
			description:  "Should use workspace_id from provider when provider_config is not set",
		},
		{
			name: "account-level with workspace_id and provider_config override - uses override",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "789012",
					},
				},
			},
			client: func() *DatabricksClient {
				c := &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:        "https://accounts.cloud.databricks.com",
							AccountID:   "test-account-id",
							Token:       "test-token",
							WorkspaceID: "123456",
						},
					},
				}
				c.Config = c.Config.WithTesting()
				// Create mock workspace client for DEFAULT workspace
				defaultWsClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  "https://default-workspace.test.databricks.com",
						Token: "test-token",
					},
				}
				defaultWsClient.Config = defaultWsClient.Config.WithTesting()
				// Create mock workspace client for OVERRIDE workspace
				overrideWsClient := &databricks.WorkspaceClient{
					Config: &config.Config{
						Host:  "https://override-workspace.test.databricks.com",
						Token: "test-token",
					},
				}
				overrideWsClient.Config = overrideWsClient.Config.WithTesting()
				// Cache BOTH workspace clients
				c.cachedWorkspaceClients = map[int64]*databricks.WorkspaceClient{
					123456: defaultWsClient,  // workspace_id - should NOT be used
					789012: overrideWsClient, // provider_config override - SHOULD be used
				}
				return c
			}(),
			expectError:  false,
			expectedHost: "https://override-workspace.test.databricks.com",
			description:  "Should use workspace_id from provider_config over provider workspace_id",
		},
		{
			name: "account-level without workspace_id and no provider_config - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:      "https://accounts.cloud.databricks.com",
						AccountID: "test-account-id",
						Token:     "test-token",
					},
				},
			},
			expectError:   true,
			errorContains: "managing workspace-level resources requires a workspace_id",
			description:   "Should return error when neither workspace_id nor provider_config.workspace_id is set",
		},
		{
			name: "workspace-level with matching workspace_id - validates and uses workspace client",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:        "https://workspace.test.databricks.com",
						Token:       "test-token",
						WorkspaceID: "123456",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     123456,
			},
			expectError: false,
			description: "Workspace-level provider falls back to Config.WorkspaceID; validation passes when it matches the cached workspace ID",
		},
		{
			name: "workspace-level with mismatching workspace_id - returns mismatch error",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			client: &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:        "https://workspace.test.databricks.com",
						Token:       "test-token",
						WorkspaceID: "999999",
					},
				},
				cachedWorkspaceClient: mockWorkspaceClient,
				cachedWorkspaceID:     123456,
			},
			expectError:   true,
			errorContains: "workspace_id mismatch",
			description:   "Workspace-level provider with Config.WorkspaceID that doesn't match the cached workspace ID errors at apply",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

			// Call WorkspaceClientUnifiedProvider
			result, err := tc.client.WorkspaceClientUnifiedProvider(ctx, d)

			// Verify results
			if tc.expectError {
				assert.Error(t, err, tc.description)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
			} else {
				assert.NoError(t, err, tc.description)
				assert.NotNil(t, result)
				// Verify the correct workspace client was returned by checking its Host
				if tc.expectedHost != "" {
					assert.Equal(t, tc.expectedHost, result.Config.Host,
						"Expected workspace client with host %s but got %s", tc.expectedHost, result.Config.Host)
				}
			}
		})
	}
}

func TestConfigWorkspaceID(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
	}

	// Initially, WorkspaceID should be empty
	assert.Empty(t, c.Config.WorkspaceID)

	// Set the workspace ID
	c.Config.WorkspaceID = "123456"

	// Verify it was set correctly
	assert.Equal(t, "123456", c.Config.WorkspaceID)

	// Set it to a different value
	c.Config.WorkspaceID = "789012"

	// Verify it was updated
	assert.Equal(t, "789012", c.Config.WorkspaceID)
}

// testCustomizeDiffForceNew is a helper that tests whether NamespaceCustomizeDiff
// triggers ForceNew for a given workspace_id transition.
// It builds a schema.Resource with provider_config, wires NamespaceCustomizeDiff as
// the CustomizeDiff, and runs the full diff pipeline to check ForceNew flags.
// It creates a proper cty-valued ResourceConfig so that GetRawConfigAt works in
// CustomizeDiff (terraform.NewResourceConfigRaw does not set cty values).
func testCustomizeDiffForceNew(t *testing.T, instanceState map[string]string, newConfig map[string]any, c *DatabricksClient) (bool, error) {
	t.Helper()

	testSchema := map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Required: true},
	}
	AddNamespaceInSchema(testSchema)

	r := &schema.Resource{
		Schema: testSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			dc, ok := m.(*DatabricksClient)
			if !ok {
				return fmt.Errorf("expected *DatabricksClient, got %T", m)
			}
			return namespaceForceNew(ctx, d, dc, true)
		},
	}

	// Build a cty.Value from the config map to ensure GetRawConfigAt works
	// in CustomizeDiff. NewResourceConfigRaw doesn't set cty values.
	ctyVal := configMapToCtyValue(newConfig)
	block := schema.InternalMap(testSchema).CoreConfigSchema()
	rc := terraform.NewResourceConfigShimmed(ctyVal, block)

	is := &terraform.InstanceState{
		ID:         "test-resource-id",
		Attributes: instanceState,
		// Set RawConfig to the new config's cty value, matching the real gRPC
		// PlanResourceChange path (grpc_provider.go:1081) where
		// priorState.RawConfig = configVal. schemaMapWithIdentity.Diff() copies
		// s.RawConfig to result.RawConfig, which ResourceDiff.GetRawConfig reads.
		RawConfig: ctyVal,
	}

	diff, err := r.Diff(context.Background(), is, rc, c)
	if err != nil {
		return false, err
	}
	if diff == nil {
		return false, nil
	}

	for _, v := range diff.Attributes {
		if v.RequiresNew {
			return true, nil
		}
	}
	return false, nil
}

// configMapToCtyValue converts a Go map (as used in test configs) to a cty.Value
// suitable for use with terraform.NewResourceConfigShimmed. It handles the
// provider_config TypeList structure used in unified provider tests.
func configMapToCtyValue(config map[string]any) cty.Value {
	providerConfigType := cty.Object(map[string]cty.Type{
		"workspace_id": cty.String,
	})

	vals := map[string]cty.Value{}
	if name, ok := config["name"]; ok {
		vals["name"] = cty.StringVal(name.(string))
	} else {
		vals["name"] = cty.NullVal(cty.String)
	}

	if pc, ok := config["provider_config"]; ok {
		pcList := pc.([]any)
		if len(pcList) > 0 {
			pcMap := pcList[0].(map[string]any)
			vals["provider_config"] = cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"workspace_id": cty.StringVal(pcMap["workspace_id"].(string)),
				}),
			})
		} else {
			vals["provider_config"] = cty.ListValEmpty(providerConfigType)
		}
	} else {
		vals["provider_config"] = cty.NullVal(cty.List(providerConfigType))
	}

	return cty.ObjectVal(vals)
}

func TestNamespaceCustomizeDiff_ForceNew(t *testing.T) {
	makeClient := func(defaultWSID string, cachedWSID int64) *DatabricksClient {
		cfg := &config.Config{
			Host:        "https://test.cloud.databricks.com",
			Token:       "test-token",
			WorkspaceID: defaultWSID,
		}
		return &DatabricksClient{
			DatabricksClient: &client.DatabricksClient{
				Config: cfg.WithTesting(),
			},
			cachedWorkspaceID: cachedWSID,
		}
	}

	tests := []struct {
		name           string
		instanceState  map[string]string
		newConfig      map[string]any
		defaultWSID    string
		cachedWSID     int64
		expectForceNew bool
		expectError    bool
		errorContains  string
	}{
		{
			name: "workspace_id changes A to B - ForceNew",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
				"provider_config": []any{
					map[string]any{"workspace_id": "200"},
				},
			},
			expectForceNew: true,
		},
		{
			name: "workspace_id same A to A - no ForceNew",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
				"provider_config": []any{
					map[string]any{"workspace_id": "100"},
				},
			},
			expectForceNew: false,
		},
		{
			name: "workspace_id added empty to A - no ForceNew",
			instanceState: map[string]string{
				"name": "test",
			},
			newConfig: map[string]any{
				"name": "test",
				"provider_config": []any{
					map[string]any{"workspace_id": "100"},
				},
			},
			expectForceNew: false,
		},
		{
			// Workspace host with old state value but no provider_config, no default,
			// and no cached ID: lazy resolution attempts CurrentWorkspaceID which fails
			// silently, then falls through to the "missing workspace_id" error.
			name: "workspace_id removed A to empty no default workspace host - error",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			expectError:   true,
			errorContains: "managing workspace-level resources requires a workspace_id",
		},
		{
			name: "workspace_id removed A to empty default=A - no ForceNew same effective",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			defaultWSID:    "100",
			expectForceNew: false,
		},
		{
			name: "workspace_id removed A to empty default=B - ForceNew effective changes",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			defaultWSID:    "200",
			expectForceNew: true,
		},
		{
			name: "workspace_id added empty to A default=A - no ForceNew same effective",
			instanceState: map[string]string{
				"name": "test",
			},
			newConfig: map[string]any{
				"name": "test",
				"provider_config": []any{
					map[string]any{"workspace_id": "100"},
				},
			},
			defaultWSID:    "100",
			expectForceNew: false,
		},
		{
			name: "workspace_id added empty to A default=B - no ForceNew without state",
			instanceState: map[string]string{
				"name": "test",
			},
			newConfig: map[string]any{
				"name": "test",
				"provider_config": []any{
					map[string]any{"workspace_id": "100"},
				},
			},
			defaultWSID:    "200",
			expectForceNew: false,
		},
		{
			name: "both empty no default - no ForceNew",
			instanceState: map[string]string{
				"name": "test",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			expectForceNew: false,
		},
		{
			name: "both empty with default - no ForceNew",
			instanceState: map[string]string{
				"name": "test",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			defaultWSID:    "100",
			expectForceNew: false,
		},
		// Workspace host change scenarios: cachedWorkspaceID reflects the
		// workspace ID resolved from the current provider host during init.
		// When the user changes the provider host (e.g. ws-A → ws-B),
		// cachedWorkspaceID changes, triggering ForceNew.
		{
			name: "host change - cachedWSID differs from state - ForceNew",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			cachedWSID:     200,
			expectForceNew: true,
		},
		{
			name: "host unchanged - cachedWSID matches state - no ForceNew",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			cachedWSID:     100,
			expectForceNew: false,
		},
		{
			// Same as "workspace_id removed" above but explicitly sets cachedWSID=0.
			// Lazy resolution fails silently, falls through to "missing workspace_id" error.
			name: "no provider_config no default cachedWSID=0 workspace host - error",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			cachedWSID:    0,
			expectError:   true,
			errorContains: "managing workspace-level resources requires a workspace_id",
		},
		{
			name: "default takes precedence over cachedWSID - ForceNew from default",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			defaultWSID:    "300",
			cachedWSID:     200,
			expectForceNew: true,
		},
		{
			name: "default matches state - cachedWSID ignored - no ForceNew",
			instanceState: map[string]string{
				"name":                           "test",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "100",
			},
			newConfig: map[string]any{
				"name": "test",
			},
			defaultWSID:    "100",
			cachedWSID:     200,
			expectForceNew: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := makeClient(tc.defaultWSID, tc.cachedWSID)
			forceNew, err := testCustomizeDiffForceNew(t, tc.instanceState, tc.newConfig, c)
			if tc.expectError {
				require.Error(t, err)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expectForceNew, forceNew,
				"ForceNew mismatch: expected %v, got %v", tc.expectForceNew, forceNew)
		})
	}
}

func TestValidateWorkspaceID(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		accountID   string
		workspaceID string
		expectError bool
	}{
		{
			name:        "workspace provider with workspace_id - error",
			host:        "https://workspace.cloud.databricks.com",
			workspaceID: "123456",
			expectError: true,
		},
		{
			name:        "account provider with workspace_id - success",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "test-account-id",
			workspaceID: "123456",
			expectError: false,
		},
		{
			name:        "workspace provider without workspace_id - success",
			host:        "https://workspace.cloud.databricks.com",
			workspaceID: "",
			expectError: false,
		},
		{
			name:        "account provider without workspace_id - success",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "test-account-id",
			workspaceID: "",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:        tc.host,
						AccountID:   tc.accountID,
						Token:       "test-token",
						WorkspaceID: tc.workspaceID,
					},
				},
			}
			// Replicate the validation logic from provider init
			hasError := c.Config.WorkspaceID != "" && c.HostTypeForTerraform() == config.WorkspaceHost
			assert.Equal(t, tc.expectError, hasError)
		})
	}
}

func TestPopulateProviderConfigInState(t *testing.T) {
	makeResourceData := func(t *testing.T, existingWSID string) *schema.ResourceData {
		t.Helper()
		testSchema := map[string]*schema.Schema{
			"name": {Type: schema.TypeString, Required: true},
		}
		AddNamespaceInSchema(testSchema)
		r := schema.Resource{Schema: testSchema}
		d := r.TestResourceData()
		d.SetId("test-resource")
		if existingWSID != "" {
			d.Set("provider_config", []map[string]any{{"workspace_id": existingWSID}})
		}
		return d
	}

	tests := []struct {
		name              string
		existingWSID      string // workspace_id already in state (simulates prior state)
		providerWSID      string // provider-level workspace_id (Config.WorkspaceID)
		cachedWorkspaceID int64
		host              string // defaults to workspace host if empty
		expectedWSID      string
	}{
		// --- First time (no state) scenarios: resolve from provider ---
		{
			name:         "first time - from provider workspace_id",
			existingWSID: "",
			providerWSID: "9876543210",
			expectedWSID: "9876543210",
		},
		{
			name:              "first time - from cachedWorkspaceID (workspace host)",
			existingWSID:      "",
			cachedWorkspaceID: 1234567890,
			expectedWSID:      "1234567890",
		},
		{
			name:              "first time - provider workspace_id takes precedence over cached",
			existingWSID:      "",
			providerWSID:      "1111111111",
			cachedWorkspaceID: 2222222222,
			expectedWSID:      "1111111111",
		},
		// Note: "account host with no workspace_id" is not tested here because
		// it cannot happen in the Terraform lifecycle — NamespaceValidateWorkspaceID
		// rejects account-level providers without workspace_id during plan, and
		// dual resources at account level are guarded by the api field early return.
		// --- Subsequent reads (has state) scenarios: preserve state ---
		{
			name:         "subsequent read - preserves state, ignores different provider workspace_id",
			existingWSID: "111",
			providerWSID: "222",
			expectedWSID: "111",
		},
		{
			name:              "subsequent read - preserves state, ignores different cached",
			existingWSID:      "111",
			cachedWorkspaceID: 222,
			expectedWSID:      "111",
		},
		{
			name:              "subsequent read - preserves state even when all sources differ",
			existingWSID:      "111",
			providerWSID:      "222",
			cachedWorkspaceID: 333,
			expectedWSID:      "111",
		},
		{
			name:         "subsequent read - preserves state with no other sources",
			existingWSID: "111",
			expectedWSID: "111",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := makeResourceData(t, tc.existingWSID)
			host := tc.host
			if host == "" {
				host = "https://my-workspace.databricks.com"
			}
			c := &DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host:        host,
						WorkspaceID: tc.providerWSID,
					},
				},
				cachedWorkspaceID: tc.cachedWorkspaceID,
			}

			err := populateProviderConfigInState(context.Background(), d, c)
			require.NoError(t, err)

			wsID := d.Get("provider_config.0.workspace_id")
			assert.Equal(t, tc.expectedWSID, wsID)
		})
	}
}

func TestGetDatabricksClientForUnifiedProvider_CopiesCommandFactory(t *testing.T) {
	// Setup: parent client with commandFactory set via WithCommandMock
	parentClient := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
	}
	parentClient.WithCommandMock(func(commandStr string) CommandResults {
		return CommandResults{ResultType: "text", Data: "mock"}
	})

	// Verify parent's CommandExecutor works
	assert.NotPanics(t, func() {
		parentClient.CommandExecutor(context.Background())
	}, "parent client should have a working CommandExecutor")

	// Pre-cache an inner client for workspace ID 123456
	innerClient := &client.DatabricksClient{
		Config: &config.Config{
			Host:  "https://workspace-123456.cloud.databricks.com",
			Token: "test-token",
		},
	}
	parentClient.cachedDatabricksClients = map[int64]*client.DatabricksClient{
		123456: innerClient,
	}

	// Call getDatabricksClientForUnifiedProvider — returned client must have
	// commandFactory copied from the parent so CommandExecutor works.
	result, err := parentClient.getDatabricksClientForUnifiedProvider(
		context.Background(), "123456",
	)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.NotPanics(t, func() {
		result.CommandExecutor(context.Background())
	}, "returned client should have commandFactory copied from parent")

	// Also verify via the public entry point DatabricksClientForUnifiedProvider
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, map[string]any{
		"name": "test",
		"provider_config": []any{
			map[string]any{
				"workspace_id": "123456",
			},
		},
	})
	result2, err := parentClient.DatabricksClientForUnifiedProvider(
		context.Background(), d,
	)
	require.NoError(t, err)
	require.NotNil(t, result2)
	assert.NotPanics(t, func() {
		result2.CommandExecutor(context.Background())
	}, "client from DatabricksClientForUnifiedProvider should have commandFactory")
}
