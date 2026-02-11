package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
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
		assert.True(t, workspaceID.Required)

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
			errorContains: "workspace_id is not set, please set the workspace_id in the provider_config",
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
			description:      "Account-level provider without workspace_id should return current client",
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
	var is *terraform.InstanceState
	if instanceState != nil {
		is = &terraform.InstanceState{Attributes: instanceState}
	}
	rc := terraform.NewResourceConfigRaw(rawConfig)
	return resource.Diff(ctx, is, rc, c)
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

func TestNamespaceCustomizeDiff_AccountLevelProvider_ForceNewOnChange(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS123 := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://ws-123.cloud.databricks.com",
			Token: "test-token",
		},
	}
	mockWS456 := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://ws-456.cloud.databricks.com",
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
	c.SetWorkspaceClientForWorkspace(123, mockWS123)
	c.SetWorkspaceClientForWorkspace(456, mockWS456)

	diff, err := diffCustomizeDiff(t, resource, map[string]string{
		"name":                           "test",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "123",
	}, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "456",
			},
		},
	}, c)
	assert.NoError(t, err)
	require.NotNil(t, diff)
	wsAttr, ok := diff.Attributes[workspaceIDSchemaKey]
	require.True(t, ok, "workspace_id should be in diff attributes")
	assert.True(t, wsAttr.RequiresNew, "changing workspace_id should require new resource")
}

func TestNamespaceCustomizeDiff_AccountLevelProvider_InvalidWorkspace(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "test-account-id",
				Token:     "test-token",
			},
		},
	}
	// No cached workspace client — WorkspaceClientForWorkspace will fail
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "999",
			},
		},
	}, c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get workspace client with workspace_id 999")
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
			Config: &config.Config{
				Host:                       "https://unified.cloud.databricks.com",
				Token:                      "test-token",
				Experimental_IsUnifiedHost: true,
			},
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

func TestNamespaceCustomizeDiff_UnifiedHost_ForceNewOnChange(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	mockWS123 := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://ws-123.cloud.databricks.com",
			Token: "test-token",
		},
	}
	mockWS456 := &databricks.WorkspaceClient{
		Config: &config.Config{
			Host:  "https://ws-456.cloud.databricks.com",
			Token: "test-token",
		},
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:                       "https://unified.cloud.databricks.com",
				Token:                      "test-token",
				Experimental_IsUnifiedHost: true,
			},
		},
	}
	c.SetWorkspaceClientForWorkspace(123, mockWS123)
	c.SetWorkspaceClientForWorkspace(456, mockWS456)

	diff, err := diffCustomizeDiff(t, resource, map[string]string{
		"name":                           "test",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "123",
	}, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "456",
			},
		},
	}, c)
	assert.NoError(t, err)
	require.NotNil(t, diff)
	wsAttr, ok := diff.Attributes[workspaceIDSchemaKey]
	require.True(t, ok, "workspace_id should be in diff attributes")
	assert.True(t, wsAttr.RequiresNew, "changing workspace_id should require new resource")
}

func TestNamespaceCustomizeDiff_UnifiedHost_InvalidWorkspace(t *testing.T) {
	resource := newTestResourceForCustomizeDiff()
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:                       "https://unified.cloud.databricks.com",
				Token:                      "test-token",
				Experimental_IsUnifiedHost: true,
			},
		},
	}
	// No cached workspace client — WorkspaceClientForWorkspace will fail
	_, err := diffCustomizeDiff(t, resource, nil, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "999",
			},
		},
	}, c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get workspace client with workspace_id 999")
}

func TestNamespaceCustomizeDiff_ForceNew(t *testing.T) {
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
		"provider_config.0.workspace_id": "789",
	}, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "789",
			},
		},
	}, c)
	assert.NoError(t, err)
	// No change in workspace_id, so no ForceNew
	if diff != nil {
		for _, v := range diff.Attributes {
			assert.False(t, v.RequiresNew, "should not require new when workspace_id is unchanged")
		}
	}
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
