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
