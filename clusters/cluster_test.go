package clusters_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/clusters"
)

func TestClusterResourceWorkloadType(t *testing.T) {
	tests := []struct {
		name     string
		clients  map[string]any
		expected map[string]any
	}{
		{
			name:     "clients default to enabled",
			clients:  map[string]any{},
			expected: map[string]any{"jobs": true, "notebooks": true},
		},
		{
			name:     "clients can be disabled",
			clients:  map[string]any{"jobs": false, "notebooks": false},
			expected: map[string]any{"jobs": false, "notebooks": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := clusters.ResourceCluster().ToResource()
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
				"workload_type": []any{
					map[string]any{"clients": []any{tt.clients}},
				},
			})

			workloadType := d.Get("workload_type").([]any)
			require.Len(t, workloadType, 1)
			clients := workloadType[0].(map[string]any)["clients"].([]any)
			require.Len(t, clients, 1)
			assert.Equal(t, tt.expected, clients[0])
		})
	}
}
