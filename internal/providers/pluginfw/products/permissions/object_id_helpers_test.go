package permissions

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestObjectIdentifiers_ToFieldValuesMap(t *testing.T) {
	tests := []struct {
		name     string
		objIds   ObjectIdentifiers
		expected map[string]string
	}{
		{
			name: "cluster_id set",
			objIds: ObjectIdentifiers{
				ClusterId: types.StringValue("cluster-123"),
			},
			expected: map[string]string{
				"cluster_id": "cluster-123",
				// All other fields should be empty strings
				"cluster_policy_id":         "",
				"instance_pool_id":          "",
				"job_id":                    "",
				"pipeline_id":               "",
				"notebook_id":               "",
				"notebook_path":             "",
				"directory_id":              "",
				"directory_path":            "",
				"workspace_file_id":         "",
				"workspace_file_path":       "",
				"registered_model_id":       "",
				"experiment_id":             "",
				"sql_dashboard_id":          "",
				"sql_endpoint_id":           "",
				"sql_query_id":              "",
				"sql_alert_id":              "",
				"dashboard_id":              "",
				"repo_id":                   "",
				"repo_path":                 "",
				"authorization":             "",
				"serving_endpoint_id":       "",
				"vector_search_endpoint_id": "",
				"app_name":                  "",
				"database_instance_name":    "",
				"alert_v2_id":               "",
			},
		},
		{
			name: "multiple fields set",
			objIds: ObjectIdentifiers{
				JobId:         types.StringValue("job-456"),
				Authorization: types.StringValue("tokens"),
				NotebookPath:  types.StringValue("/Users/test/notebook"),
			},
			expected: map[string]string{
				"cluster_id":                "",
				"cluster_policy_id":         "",
				"instance_pool_id":          "",
				"job_id":                    "job-456",
				"pipeline_id":               "",
				"notebook_id":               "",
				"notebook_path":             "/Users/test/notebook",
				"directory_id":              "",
				"directory_path":            "",
				"workspace_file_id":         "",
				"workspace_file_path":       "",
				"registered_model_id":       "",
				"experiment_id":             "",
				"sql_dashboard_id":          "",
				"sql_endpoint_id":           "",
				"sql_query_id":              "",
				"sql_alert_id":              "",
				"dashboard_id":              "",
				"repo_id":                   "",
				"repo_path":                 "",
				"authorization":             "tokens",
				"serving_endpoint_id":       "",
				"vector_search_endpoint_id": "",
				"app_name":                  "",
				"database_instance_name":    "",
				"alert_v2_id":               "",
			},
		},
		{
			name:   "all fields empty",
			objIds: ObjectIdentifiers{},
			expected: map[string]string{
				"cluster_id":                "",
				"cluster_policy_id":         "",
				"instance_pool_id":          "",
				"job_id":                    "",
				"pipeline_id":               "",
				"notebook_id":               "",
				"notebook_path":             "",
				"directory_id":              "",
				"directory_path":            "",
				"workspace_file_id":         "",
				"workspace_file_path":       "",
				"registered_model_id":       "",
				"experiment_id":             "",
				"sql_dashboard_id":          "",
				"sql_endpoint_id":           "",
				"sql_query_id":              "",
				"sql_alert_id":              "",
				"dashboard_id":              "",
				"repo_id":                   "",
				"repo_path":                 "",
				"authorization":             "",
				"serving_endpoint_id":       "",
				"vector_search_endpoint_id": "",
				"app_name":                  "",
				"database_instance_name":    "",
				"alert_v2_id":               "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.objIds.ToFieldValuesMap()

			// Check that all expected keys are present
			assert.Len(t, result, len(tt.expected), "Map should have correct number of entries")

			// Check specific values we care about
			for key, expectedValue := range tt.expected {
				assert.Contains(t, result, key, "Map should contain key: %s", key)
				assert.Equal(t, expectedValue, result[key], "Value for key %s should match", key)
			}
		})
	}
}

func TestGetObjectIdentifierFields(t *testing.T) {
	// Verify that GetObjectIdentifierFields returns all expected fields
	expectedFields := []string{
		"cluster_id",
		"cluster_policy_id",
		"instance_pool_id",
		"job_id",
		"pipeline_id",
		"notebook_id",
		"notebook_path",
		"directory_id",
		"directory_path",
		"workspace_file_id",
		"workspace_file_path",
		"registered_model_id",
		"experiment_id",
		"sql_dashboard_id",
		"sql_endpoint_id",
		"sql_query_id",
		"sql_alert_id",
		"dashboard_id",
		"repo_id",
		"repo_path",
		"authorization",
		"serving_endpoint_id",
		"vector_search_endpoint_id",
		"app_name",
		"database_instance_name",
		"alert_v2_id",
	}

	fields := GetObjectIdentifierFields()
	assert.Len(t, fields, len(expectedFields), "Should have all object identifier fields")

	// Check that all expected fields are present
	fieldMap := make(map[string]bool)
	for _, field := range fields {
		fieldMap[field] = true
	}

	for _, expected := range expectedFields {
		assert.True(t, fieldMap[expected], "GetObjectIdentifierFields() should contain: %s", expected)
	}
}
