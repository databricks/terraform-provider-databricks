package exporter

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestGetSourceCloud(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		expected string
	}{
		{
			name:     "AWS workspace",
			host:     "https://test.cloud.databricks.com",
			expected: "aws",
		},
		{
			name:     "Azure workspace",
			host:     "https://adb-123.azuredatabricks.net",
			expected: "azure",
		},
		{
			name:     "GCP workspace",
			host:     "https://test.gcp.databricks.com",
			expected: "gcp",
		},
		{
			name:     "AWS accounts",
			host:     "https://accounts.cloud.databricks.com",
			expected: "aws",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &common.DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host: tt.host,
					},
				},
			}
			ic := &importContext{Client: client}
			result := ic.getSourceCloud()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertAvailability(t *testing.T) {
	tests := []struct {
		name         string
		availability string
		sourceCloud  string
		targetCloud  string
		expected     string
	}{
		{
			name:         "AWS SPOT to Azure",
			availability: "SPOT",
			sourceCloud:  "aws",
			targetCloud:  "azure",
			expected:     "SPOT_AZURE",
		},
		{
			name:         "AWS ON_DEMAND to GCP",
			availability: "ON_DEMAND",
			sourceCloud:  "aws",
			targetCloud:  "gcp",
			expected:     "ON_DEMAND_GCP",
		},
		{
			name:         "Azure SPOT_WITH_FALLBACK_AZURE to AWS",
			availability: "SPOT_WITH_FALLBACK_AZURE",
			sourceCloud:  "azure",
			targetCloud:  "aws",
			expected:     "SPOT_WITH_FALLBACK",
		},
		{
			name:         "GCP PREEMPTIBLE_GCP to Azure",
			availability: "PREEMPTIBLE_GCP",
			sourceCloud:  "gcp",
			targetCloud:  "azure",
			expected:     "SPOT_AZURE",
		},
		{
			name:         "Same cloud no change",
			availability: "SPOT",
			sourceCloud:  "aws",
			targetCloud:  "aws",
			expected:     "SPOT",
		},
		{
			name:         "Unknown availability unchanged",
			availability: "UNKNOWN",
			sourceCloud:  "aws",
			targetCloud:  "azure",
			expected:     "UNKNOWN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertAvailability(tt.availability, tt.sourceCloud, tt.targetCloud)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsCompatibleAttribute(t *testing.T) {
	tests := []struct {
		name        string
		attrName    string
		sourceCloud string
		targetCloud string
		attrValue   interface{}
		expected    bool
	}{
		{
			name:        "availability is always compatible",
			attrName:    "availability",
			sourceCloud: "aws",
			targetCloud: "azure",
			attrValue:   "SPOT",
			expected:    true,
		},
		{
			name:        "first_on_demand is always compatible",
			attrName:    "first_on_demand",
			sourceCloud: "aws",
			targetCloud: "gcp",
			attrValue:   1,
			expected:    true,
		},
		{
			name:        "zone_id auto is compatible between AWS and GCP",
			attrName:    "zone_id",
			sourceCloud: "aws",
			targetCloud: "gcp",
			attrValue:   "auto",
			expected:    true,
		},
		{
			name:        "zone_id AUTO is compatible (case insensitive)",
			attrName:    "zone_id",
			sourceCloud: "aws",
			targetCloud: "gcp",
			attrValue:   "AUTO",
			expected:    true,
		},
		{
			name:        "zone_id specific value not compatible",
			attrName:    "zone_id",
			sourceCloud: "aws",
			targetCloud: "gcp",
			attrValue:   "us-west-2a",
			expected:    false,
		},
		{
			name:        "ebs_volume_count compatible with local_ssd_count when GENERAL_PURPOSE_SSD",
			attrName:    "ebs_volume_count",
			sourceCloud: "aws",
			targetCloud: "gcp",
			attrValue:   2,
			expected:    true,
		},
		{
			name:        "instance_profile_arn not compatible",
			attrName:    "instance_profile_arn",
			sourceCloud: "aws",
			targetCloud: "azure",
			attrValue:   "arn:aws:iam::123:instance-profile/test",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For ebs_volume_count test, we need to pass additional context
			additionalData := map[string]interface{}{}
			if tt.attrName == "ebs_volume_count" {
				additionalData["ebs_volume_type"] = "GENERAL_PURPOSE_SSD"
			}
			result := isCompatibleAttribute(tt.attrName, tt.attrValue, tt.sourceCloud, tt.targetCloud, additionalData)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertCloudAttributesBlock(t *testing.T) {
	tests := []struct {
		name         string
		sourceCloud  string
		targetCloud  string
		sourceAttrs  map[string]interface{}
		expectedType string
		expectedLen  int
	}{
		{
			name:        "AWS to Azure conversion",
			sourceCloud: "aws",
			targetCloud: "azure",
			sourceAttrs: map[string]interface{}{
				"availability":         "SPOT",
				"first_on_demand":      1,
				"zone_id":              "us-west-2a",                             // not compatible
				"instance_profile_arn": "arn:aws:iam::123:instance-profile/test", // not compatible
			},
			expectedType: "azure_attributes",
			expectedLen:  2, // only availability and first_on_demand
		},
		{
			name:        "AWS to GCP with auto zone",
			sourceCloud: "aws",
			targetCloud: "gcp",
			sourceAttrs: map[string]interface{}{
				"availability":     "ON_DEMAND",
				"first_on_demand":  1,
				"zone_id":          "auto",
				"ebs_volume_count": 2,
				"ebs_volume_type":  "GENERAL_PURPOSE_SSD",
			},
			expectedType: "gcp_attributes",
			expectedLen:  4, // availability, first_on_demand, zone_id, local_ssd_count
		},
		{
			name:        "Same cloud returns nil",
			sourceCloud: "aws",
			targetCloud: "aws",
			sourceAttrs: map[string]interface{}{
				"availability": "SPOT",
			},
			expectedType: "",
			expectedLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertCloudAttributesBlock(tt.sourceAttrs, tt.sourceCloud, tt.targetCloud)
			if tt.expectedType == "" {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedLen, len(result))

				// Verify specific conversions
				if tt.name == "AWS to Azure conversion" {
					assert.Equal(t, "SPOT_AZURE", result["availability"])
				}
				if tt.name == "AWS to GCP with auto zone" {
					assert.Equal(t, "ON_DEMAND_GCP", result["availability"])
					assert.Contains(t, result, "local_ssd_count")
					assert.NotContains(t, result, "ebs_volume_count")
				}
			}
		})
	}
}

func TestConvertTopLevelNodeTypes(t *testing.T) {
	tests := []struct {
		name        string
		nodeType    string
		driverType  string
		sourceCloud string
		targetCloud string
		expectNode  string
		expectDrv   string
	}{
		{
			name:        "AWS to Azure conversion",
			nodeType:    "i3.xlarge",
			driverType:  "i3.2xlarge",
			sourceCloud: "aws",
			targetCloud: "azure",
			expectNode:  "Standard_F4s",
			expectDrv:   "Standard_F8s",
		},
		{
			name:        "Same cloud unchanged",
			nodeType:    "i3.xlarge",
			driverType:  "i3.2xlarge",
			sourceCloud: "aws",
			targetCloud: "aws",
			expectNode:  "i3.xlarge",
			expectDrv:   "i3.2xlarge",
		},
		{
			name:        "Unmapped type unchanged",
			nodeType:    "unknown-type",
			driverType:  "i3.xlarge",
			sourceCloud: "aws",
			targetCloud: "azure",
			expectNode:  "unknown-type",
			expectDrv:   "Standard_F4s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
				},
			}

			client := &common.DatabricksClient{
				DatabricksClient: &client.DatabricksClient{
					Config: &config.Config{
						Host: "https://test.cloud.databricks.com",
					},
				},
			}
			ic := &importContext{
				Client:           client,
				targetCloud:      tt.targetCloud,
				nodeTypeMappings: mappings,
			}

			// Create a mock wrapper
			data := map[string]interface{}{
				"node_type_id":        tt.nodeType,
				"driver_node_type_id": tt.driverType,
			}
			wrapper := &mockWrapper{data: data}

			converted := ic.convertTopLevelNodeTypes(wrapper, tt.sourceCloud, tt.targetCloud)

			if tt.sourceCloud == tt.targetCloud {
				assert.False(t, converted, "Should not convert for same cloud")
			} else if tt.expectNode != tt.nodeType || tt.expectDrv != tt.driverType {
				assert.True(t, converted, "Should convert when types change")
			}

			assert.Equal(t, tt.expectNode, data["node_type_id"])
			assert.Equal(t, tt.expectDrv, data["driver_node_type_id"])
		})
	}
}

// mockWrapper is a simple mock implementation for testing
type mockWrapper struct {
	data map[string]interface{}
	id   string
}

func (m *mockWrapper) GetOk(key string) (interface{}, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *mockWrapper) Get(key string) interface{} {
	return m.data[key]
}

func (m *mockWrapper) Set(key string, value interface{}) error {
	m.data[key] = value
	return nil
}

func (m *mockWrapper) SetId(id string) {
	m.id = id
}

func (m *mockWrapper) Id() string {
	if m.id != "" {
		return m.id
	}
	return "test-resource"
}

func (m *mockWrapper) Schema() map[string]*SchemaWrapper {
	return nil
}

func (m *mockWrapper) GetSchema() SchemaWrapper {
	return nil
}

func (m *mockWrapper) IsPluginFramework() bool {
	return false
}

func (m *mockWrapper) GetTypedStruct(ctx context.Context, target interface{}) error {
	return fmt.Errorf("not implemented for mock")
}

func TestConvertCloudAttributesInWrapper(t *testing.T) {
	tests := []struct {
		name                string
		sourceAttrs         map[string]interface{}
		sourceCloud         string
		targetCloud         string
		expectConverted     bool
		expectSourceCleared bool
		validateTarget      func(t *testing.T, result map[string]interface{})
	}{
		{
			name: "AWS to Azure with compatible attrs",
			sourceAttrs: map[string]interface{}{
				"availability":    "SPOT",
				"first_on_demand": 1,
				"zone_id":         "auto",
			},
			sourceCloud:         "aws",
			targetCloud:         "azure",
			expectConverted:     true,
			expectSourceCleared: true,
			validateTarget: func(t *testing.T, result map[string]interface{}) {
				assert.Equal(t, "SPOT_AZURE", result["availability"])
				assert.Equal(t, 1, result["first_on_demand"])
				// zone_id not compatible between AWS and Azure
				assert.NotContains(t, result, "zone_id")
			},
		},
		{
			name: "AWS to Azure with no compatible attrs",
			sourceAttrs: map[string]interface{}{
				"instance_profile_arn": "arn:aws:iam::123:instance-profile/test",
				"zone_id":              "us-west-2a",
			},
			sourceCloud:         "aws",
			targetCloud:         "azure",
			expectConverted:     false,
			expectSourceCleared: true, // Clear even if nothing converts
			validateTarget: func(t *testing.T, result map[string]interface{}) {
				assert.Nil(t, result, "Should be nil when no compatible attributes")
			},
		},
		{
			name: "AWS to GCP with ebs_volume conversion",
			sourceAttrs: map[string]interface{}{
				"availability":     "ON_DEMAND",
				"first_on_demand":  1,
				"zone_id":          "auto",
				"ebs_volume_count": 2,
				"ebs_volume_type":  "GENERAL_PURPOSE_SSD",
			},
			sourceCloud:         "aws",
			targetCloud:         "gcp",
			expectConverted:     true,
			expectSourceCleared: true,
			validateTarget: func(t *testing.T, result map[string]interface{}) {
				assert.Equal(t, "ON_DEMAND_GCP", result["availability"])
				assert.Equal(t, "auto", result["zone_id"])
				assert.Equal(t, 2, result["local_ssd_count"])
				assert.NotContains(t, result, "ebs_volume_count")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertCloudAttributesBlock(tt.sourceAttrs, tt.sourceCloud, tt.targetCloud)

			if tt.expectConverted {
				assert.NotNil(t, result)
			}

			if tt.validateTarget != nil {
				tt.validateTarget(t, result)
			}
		})
	}
}

func TestConvertInstancePoolDiskSpec(t *testing.T) {
	tests := []struct {
		name         string
		sourceCloud  string
		targetCloud  string
		diskSpec     []interface{}
		expectChange bool
		validate     func(t *testing.T, wrapper *mockWrapper)
	}{
		{
			name:        "AWS to Azure - GENERAL_PURPOSE_SSD to PREMIUM_LRS",
			sourceCloud: "aws",
			targetCloud: "azure",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 2,
					"disk_size":  100,
					"disk_type": []interface{}{
						map[string]interface{}{
							"ebs_volume_type": "GENERAL_PURPOSE_SSD",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				assert.Len(t, diskSpec, 1)
				diskSpecMap := diskSpec[0].(map[string]interface{})

				diskType := diskSpecMap["disk_type"].([]interface{})
				assert.Len(t, diskType, 1)
				diskTypeMap := diskType[0].(map[string]interface{})

				assert.Equal(t, "PREMIUM_LRS", diskTypeMap["azure_disk_volume_type"])
				assert.NotContains(t, diskTypeMap, "ebs_volume_type")
			},
		},
		{
			name:        "AWS to Azure - THROUGHPUT_OPTIMIZED_HDD to STANDARD_LRS",
			sourceCloud: "aws",
			targetCloud: "azure",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 1,
					"disk_type": []interface{}{
						map[string]interface{}{
							"ebs_volume_type": "THROUGHPUT_OPTIMIZED_HDD",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})
				diskType := diskSpecMap["disk_type"].([]interface{})
				diskTypeMap := diskType[0].(map[string]interface{})

				assert.Equal(t, "STANDARD_LRS", diskTypeMap["azure_disk_volume_type"])
				assert.NotContains(t, diskTypeMap, "ebs_volume_type")
			},
		},
		{
			name:        "Azure to AWS - PREMIUM_LRS to GENERAL_PURPOSE_SSD",
			sourceCloud: "azure",
			targetCloud: "aws",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 2,
					"disk_type": []interface{}{
						map[string]interface{}{
							"azure_disk_volume_type": "PREMIUM_LRS",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})
				diskType := diskSpecMap["disk_type"].([]interface{})
				diskTypeMap := diskType[0].(map[string]interface{})

				assert.Equal(t, "GENERAL_PURPOSE_SSD", diskTypeMap["ebs_volume_type"])
				assert.NotContains(t, diskTypeMap, "azure_disk_volume_type")
			},
		},
		{
			name:        "Azure to AWS - STANDARD_LRS to THROUGHPUT_OPTIMIZED_HDD",
			sourceCloud: "azure",
			targetCloud: "aws",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 1,
					"disk_type": []interface{}{
						map[string]interface{}{
							"azure_disk_volume_type": "STANDARD_LRS",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})
				diskType := diskSpecMap["disk_type"].([]interface{})
				diskTypeMap := diskType[0].(map[string]interface{})

				assert.Equal(t, "THROUGHPUT_OPTIMIZED_HDD", diskTypeMap["ebs_volume_type"])
				assert.NotContains(t, diskTypeMap, "azure_disk_volume_type")
			},
		},
		{
			name:        "AWS to GCP - disk_type removed",
			sourceCloud: "aws",
			targetCloud: "gcp",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 2,
					"disk_type": []interface{}{
						map[string]interface{}{
							"ebs_volume_type": "GENERAL_PURPOSE_SSD",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})

				// disk_type should be removed for GCP
				assert.NotContains(t, diskSpecMap, "disk_type")
				// disk_count should remain
				assert.Equal(t, 2, diskSpecMap["disk_count"])
			},
		},
		{
			name:        "Azure to GCP - disk_type removed",
			sourceCloud: "azure",
			targetCloud: "gcp",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 1,
					"disk_type": []interface{}{
						map[string]interface{}{
							"azure_disk_volume_type": "PREMIUM_LRS",
						},
					},
				},
			},
			expectChange: true,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})

				assert.NotContains(t, diskSpecMap, "disk_type")
				assert.Equal(t, 1, diskSpecMap["disk_count"])
			},
		},
		{
			name:        "GCP to AWS - no disk_type to convert",
			sourceCloud: "gcp",
			targetCloud: "aws",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 2,
					// No disk_type on GCP
				},
			},
			expectChange: false,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})

				// Should remain unchanged
				assert.Equal(t, 2, diskSpecMap["disk_count"])
				assert.NotContains(t, diskSpecMap, "disk_type")
			},
		},
		{
			name:        "Same cloud - no conversion",
			sourceCloud: "aws",
			targetCloud: "aws",
			diskSpec: []interface{}{
				map[string]interface{}{
					"disk_count": 2,
					"disk_type": []interface{}{
						map[string]interface{}{
							"ebs_volume_type": "GENERAL_PURPOSE_SSD",
						},
					},
				},
			},
			expectChange: false,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				diskSpec := wrapper.data["disk_spec"].([]interface{})
				diskSpecMap := diskSpec[0].(map[string]interface{})
				diskType := diskSpecMap["disk_type"].([]interface{})
				diskTypeMap := diskType[0].(map[string]interface{})

				// Should remain unchanged
				assert.Equal(t, "GENERAL_PURPOSE_SSD", diskTypeMap["ebs_volume_type"])
			},
		},
		{
			name:         "No disk_spec - no conversion",
			sourceCloud:  "aws",
			targetCloud:  "azure",
			diskSpec:     nil,
			expectChange: false,
			validate: func(t *testing.T, wrapper *mockWrapper) {
				assert.Nil(t, wrapper.data["disk_spec"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := &mockWrapper{
				data: make(map[string]interface{}),
				id:   "test-pool-123",
			}

			if tt.diskSpec != nil {
				wrapper.data["disk_spec"] = tt.diskSpec
			}

			ic := &importContext{}
			result := ic.convertInstancePoolDiskSpec(wrapper, tt.sourceCloud, tt.targetCloud)

			assert.Equal(t, tt.expectChange, result)

			if tt.validate != nil {
				tt.validate(t, wrapper)
			}
		})
	}
}
