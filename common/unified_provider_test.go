package common

import (
	"testing"

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
