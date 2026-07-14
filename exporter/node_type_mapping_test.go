package exporter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadNodeTypeMappings(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectError bool
		validate    func(t *testing.T, mappings *nodeTypeMappings)
	}{
		{
			name: "Valid mapping file",
			content: `{
				"version": "1.0",
				"mappings": [
					{
						"azure": "Standard_F4s",
						"aws": "i3.xlarge",
						"gcp": "n1-standard-4"
					},
					{
						"azure": "Standard_F8s",
						"aws": "i3.2xlarge",
						"gcp": "n1-standard-8"
					}
				]
			}`,
			expectError: false,
			validate: func(t *testing.T, mappings *nodeTypeMappings) {
				assert.Equal(t, "1.0", mappings.Version)
				assert.Len(t, mappings.Mappings, 2)
				assert.Equal(t, "Standard_F4s", mappings.Mappings[0].Azure)
				assert.Equal(t, "i3.xlarge", mappings.Mappings[0].AWS)
				assert.Equal(t, "n1-standard-4", mappings.Mappings[0].GCP)
			},
		},
		{
			name: "Empty file",
			content: `{
				"version": "1.0",
				"mappings": []
			}`,
			expectError: false,
			validate: func(t *testing.T, mappings *nodeTypeMappings) {
				assert.Equal(t, "1.0", mappings.Version)
				assert.Len(t, mappings.Mappings, 0)
			},
		},
		{
			name:        "Invalid JSON",
			content:     `{invalid json}`,
			expectError: true,
		},
		{
			name: "Missing version",
			content: `{
				"mappings": [
					{"azure": "Standard_F4s", "aws": "i3.xlarge"}
				]
			}`,
			expectError: false,
			validate: func(t *testing.T, mappings *nodeTypeMappings) {
				assert.Empty(t, mappings.Version)
				assert.Len(t, mappings.Mappings, 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "mappings.json")
			err := os.WriteFile(tmpFile, []byte(tt.content), 0644)
			require.NoError(t, err)

			// Load mappings
			mappings, err := loadNodeTypeMappings(tmpFile)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, mappings)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, mappings)
				if tt.validate != nil {
					tt.validate(t, mappings)
				}
			}
		})
	}
}

func TestLoadNodeTypeMappingsNonExistentFile(t *testing.T) {
	mappings, err := loadNodeTypeMappings("/nonexistent/file.json")
	assert.Error(t, err)
	assert.Nil(t, mappings)
}

func TestConvertNodeType(t *testing.T) {
	mappings := &nodeTypeMappings{
		Version: "1.0",
		Mappings: []nodeTypeMapping{
			{
				Azure: "Standard_F4s",
				AWS:   "i3.xlarge",
				GCP:   "n1-standard-4",
			},
			{
				Azure: "Standard_F8s",
				AWS:   "i3.2xlarge",
				GCP:   "n1-standard-8",
			},
			{
				Azure: "Standard_D4s_v3",
				AWS:   "m5.xlarge",
				// No GCP mapping
			},
		},
	}

	tests := []struct {
		name        string
		nodeType    string
		sourceCloud string
		targetCloud string
		expected    string
	}{
		{
			name:        "AWS to Azure",
			nodeType:    "i3.xlarge",
			sourceCloud: "aws",
			targetCloud: "azure",
			expected:    "Standard_F4s",
		},
		{
			name:        "Azure to GCP",
			nodeType:    "Standard_F4s",
			sourceCloud: "azure",
			targetCloud: "gcp",
			expected:    "n1-standard-4",
		},
		{
			name:        "GCP to AWS",
			nodeType:    "n1-standard-8",
			sourceCloud: "gcp",
			targetCloud: "aws",
			expected:    "i3.2xlarge",
		},
		{
			name:        "Same cloud unchanged",
			nodeType:    "i3.xlarge",
			sourceCloud: "aws",
			targetCloud: "aws",
			expected:    "i3.xlarge",
		},
		{
			name:        "Unmapped type unchanged",
			nodeType:    "some-unknown-type",
			sourceCloud: "aws",
			targetCloud: "azure",
			expected:    "some-unknown-type",
		},
		{
			name:        "Target cloud missing in mapping",
			nodeType:    "Standard_D4s_v3",
			sourceCloud: "azure",
			targetCloud: "gcp",
			expected:    "Standard_D4s_v3",
		},
		{
			name:        "Case sensitivity",
			nodeType:    "I3.XLARGE",
			sourceCloud: "aws",
			targetCloud: "azure",
			expected:    "I3.XLARGE", // No match due to case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertNodeType(tt.nodeType, tt.sourceCloud, tt.targetCloud, mappings)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertNodeTypeNilMappings(t *testing.T) {
	result := convertNodeType("i3.xlarge", "aws", "azure", nil)
	assert.Equal(t, "i3.xlarge", result, "Should return original when mappings is nil")
}
