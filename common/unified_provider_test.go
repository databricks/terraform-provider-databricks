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
<<<<<<< HEAD
			expectError: false, // This is actually valid as it's still a numeric string
=======
			expectError: true, // Leading zeros are not allowed
		},
		{
			name:        "invalid single zero",
			input:       "0",
			expectError: true, // Zero is not a valid workspace ID
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD

	// Test the validation function
	_, errors := workspaceID.ValidateFunc("123456", "workspace_id")
	assert.Empty(t, errors, "Valid workspace ID should not produce errors")

	_, errors = workspaceID.ValidateFunc("invalid", "workspace_id")
	assert.NotEmpty(t, errors, "Invalid workspace ID should produce errors")
=======
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD

		// Test the validation function
		_, errors := workspaceID.ValidateFunc("123456", "workspace_id")
		assert.Empty(t, errors, "Valid workspace ID should not produce errors")

		_, errors = workspaceID.ValidateFunc("invalid", "workspace_id")
		assert.NotEmpty(t, errors, "Invalid workspace ID should produce errors")
	})

	t.Run("without provider_config", func(t *testing.T) {
=======
	})

	t.Run("panic without provider_config", func(t *testing.T) {
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
		testSchema := map[string]*schema.Schema{
			"other_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}

<<<<<<< HEAD
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
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD

	testCases := []struct {
		name              string
		resourceData      map[string]interface{}
		cachedWorkspaceID int64
		isAccountLevel    bool
		accountID         string
		expectError       bool
		errorContains     string
		description       string
=======
	mockWorkspaceClient := &databricks.WorkspaceClient{}

	testCases := []struct {
		name          string
		resourceData  map[string]interface{}
		client        *DatabricksClient
		expectError   bool
		errorContains string
		description   string
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
	}{
		{
			name: "workspace_id not set - calls with empty string",
			resourceData: map[string]interface{}{
				"name": "test",
			},
<<<<<<< HEAD
			cachedWorkspaceID: 0,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When provider_config is not set, should use cached workspace client",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD
			cachedWorkspaceID: 123456,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When workspace_id matches cached ID, should return workspace client",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD
			cachedWorkspaceID: 0,
			isAccountLevel:    false,
			expectError:       false,
			description:       "When workspace_id is explicitly empty, should use cached workspace client",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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
<<<<<<< HEAD
			cachedWorkspaceID: 789012,
			isAccountLevel:    false,
			expectError:       false,
			description:       "Should handle different workspace IDs correctly",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
		},
		{
			name: "account level provider without workspace_id - returns error",
			resourceData: map[string]interface{}{
				"name": "test",
			},
<<<<<<< HEAD
			cachedWorkspaceID: 0,
			isAccountLevel:    true,
			accountID:         "test-account-id",
			expectError:       true,
			errorContains:     "workspace_id",
			description:       "Account-level provider requires workspace_id to be set",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create resource data
			d := schema.TestResourceDataRaw(t, testSchema, tc.resourceData)

<<<<<<< HEAD
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
=======
			// Call WorkspaceClientUnifiedProvider
			result, err := tc.client.WorkspaceClientUnifiedProvider(ctx, d)
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c

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

<<<<<<< HEAD
func TestWorkspaceIDExtractionLogic(t *testing.T) {
=======
func TestDatabricksClientForUnifiedProvider(t *testing.T) {
	cachedWorkspaceHost := "https://workspace.test.databricks.com"
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
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

<<<<<<< HEAD
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
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
						"workspace_id": "",
					},
				},
			},
<<<<<<< HEAD
			expectedValue: "",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns empty string when explicitly set to empty",
		},
		{
			name: "workspace_id with numeric value",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
			resourceData: map[string]interface{}{
				"name": "test",
				"provider_config": []interface{}{
					map[string]interface{}{
<<<<<<< HEAD
						"workspace_id": "999999999",
					},
				},
			},
			expectedValue: "999999999",
			expectedOk:    true,
			shouldBeNil:   false,
			description:   "d.Get returns large numeric workspace_id as string",
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
<<<<<<< HEAD
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
=======
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
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
			}
		})
	}
}
