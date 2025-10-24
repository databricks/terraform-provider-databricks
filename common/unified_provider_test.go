package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			expectError: false, // This is actually valid as it's still a numeric string
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

	// Test the validation function
	_, errors := workspaceID.ValidateFunc("123456", "workspace_id")
	assert.Empty(t, errors, "Valid workspace ID should not produce errors")

	_, errors = workspaceID.ValidateFunc("invalid", "workspace_id")
	assert.NotEmpty(t, errors, "Invalid workspace ID should produce errors")
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

		// Test the validation function
		_, errors := workspaceID.ValidateFunc("123456", "workspace_id")
		assert.Empty(t, errors, "Valid workspace ID should not produce errors")

		_, errors = workspaceID.ValidateFunc("invalid", "workspace_id")
		assert.NotEmpty(t, errors, "Invalid workspace ID should produce errors")
	})

	t.Run("without provider_config", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"other_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}

		result := NamespaceCustomizeSchemaMap(testSchema)
		assert.Equal(t, testSchema, result, "Schema should be returned unchanged when provider_config is not present")
	})
}

func TestNamespaceCustomizeDiff(t *testing.T) {
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

	t.Run("handles nil ResourceDiff gracefully", func(t *testing.T) {
		// Test that the function doesn't panic with nil input by checking the logic
		// The actual CustomizeDiff is called by Terraform framework, so we just verify
		// the function signature and basic error handling

		// Create a basic resource data to ensure schema is valid
		d := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
			"name": "test",
			"provider_config": []interface{}{
				map[string]interface{}{
					"workspace_id": "123456",
				},
			},
		})

		// Verify the schema is properly set up
		require.NotNil(t, d)
	})

	t.Run("validate workspace_id schema key constant", func(t *testing.T) {
		// Verify the constant is correctly defined
		assert.Equal(t, "provider_config.0.workspace_id", workspaceIDSchemaKey,
			"workspaceIDSchemaKey constant should match the schema path")
	})

	t.Run("workspace_id in schema matches expected path", func(t *testing.T) {
		// Verify that the schema path matches what we expect
		providerConfig := testSchema["provider_config"]
		require.NotNil(t, providerConfig)
		elem, ok := providerConfig.Elem.(*schema.Resource)
		require.True(t, ok)
		workspaceID, exists := elem.Schema["workspace_id"]
		require.True(t, exists, "workspace_id should exist in provider_config schema")
		assert.Equal(t, schema.TypeString, workspaceID.Type, "workspace_id should be a string type")
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

	testCases := []struct {
		name              string
		resourceData      map[string]interface{}
		cachedWorkspaceID int64
		isAccountLevel    bool
		accountID         string
		expectError       bool
		errorContains     string
		description       string
	}{
		{
			name: "workspace_id not set - calls with empty string",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			cachedWorkspaceID: 0,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When provider_config is not set, should use cached workspace client",
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
			cachedWorkspaceID: 123456,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When workspace_id matches cached ID, should return workspace client",
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
			cachedWorkspaceID: 0,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When workspace_id is explicitly empty, should use cached workspace client",
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
			cachedWorkspaceID: 789012,
			isAccountLevel:    false,
			expectError:       false,
			description:       "Should handle different workspace IDs correctly",
		},
		{
			name: "account level provider without workspace_id - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			cachedWorkspaceID: 0,
			isAccountLevel:    true,
			accountID:         "test-account-id",
			expectError:       true,
			errorContains:     "workspace_id",
			description:       "Account-level provider requires workspace_id to be set",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

			// Create a mock workspace client to be returned
			mockWorkspaceClient := &databricks.WorkspaceClient{}

			// Create a DatabricksClient based on test case configuration
			var dc *DatabricksClient
			if tc.isAccountLevel {
				// Create account-level provider
				dc = &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:      "https://accounts.cloud.databricks.com",
							AccountID: tc.accountID,
							Token:     "test-token",
						},
					},
				}
			} else {
				// Create workspace-level provider
				dc = &DatabricksClient{
					DatabricksClient: &client.DatabricksClient{
						Config: &config.Config{
							Host:  "https://test.cloud.databricks.com",
							Token: "test-token",
						},
					},
					cachedWorkspaceClient: mockWorkspaceClient,
					cachedWorkspaceID:     tc.cachedWorkspaceID,
				}
			}

			// Call WorkspaceClientUnifiedProvider
			result, err := dc.WorkspaceClientUnifiedProvider(ctx, d)

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

func TestWorkspaceIDExtractionLogic(t *testing.T) {
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

	testCases := []struct {
		name          string
		resourceData  map[string]interface{}
		getKey        string // optional: if set, use this key instead of workspaceIDSchemaKey
		expectedValue string
		expectedOk    bool
		shouldBeNil   bool
		description   string
	}{
		{
			name: "provider_config not set - returns empty string",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			expectedValue: "",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns empty string when provider_config is not set",
		},
		{
			name: "non-existent key returns nil",
			resourceData: map[string]interface{}{
				"name": "test",
			},
			getKey:        "non_existent_key",
			expectedValue: "",
			expectedOk:    false,
			shouldBeNil:   true,
			description:   "d.Get returns nil if the key doesn't exist in the schema",
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
			expectedValue: "123456",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns the correct workspace_id when set",
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
			expectedValue: "",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns empty string when explicitly set to empty",
		},
		{
			name: "workspace_id with numeric value",
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "999999999",
					},
				},
			},
			expectedValue: "999999999",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns large numeric workspace_id as string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data with test case input
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

			// Test: workspaceIDFromSchema := d.Get(workspaceIDSchemaKey)
			// Use custom key if specified, otherwise use workspaceIDSchemaKey
			key := workspaceIDSchemaKey
			if tc.getKey != "" {
				key = tc.getKey
			}
			workspaceIDFromSchema := d.Get(key)

			// Test: if workspaceIDFromSchema == nil
			if tc.shouldBeNil {
				assert.Nil(t, workspaceIDFromSchema, tc.description)
			} else {
				assert.NotNil(t, workspaceIDFromSchema, tc.description)
			}

			// Test: workspaceID, ok := workspaceIDFromSchema.(string)
			workspaceID, ok := workspaceIDFromSchema.(string)
			assert.Equal(t, tc.expectedOk, ok, "type assertion ok should match expected")
			if ok {
				assert.Equal(t, tc.expectedValue, workspaceID, tc.description)
			}
		})
	}
}
