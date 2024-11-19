package cluster

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestNoClusterError(t *testing.T) {
	clusterName := "test-cluster-name"
	clusters := []compute_tf.ClusterDetails{}
	actualDiagnostics := validateClustersList(context.Background(), clusters, clusterName)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no cluster with name '%s'", clusterName), "")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}

func TestMultipleClustersError(t *testing.T) {
	clusterName := "test-cluster-name"
	clusters := []compute_tf.ClusterDetails{
		{
			ClusterName: types.StringValue("test-cluster-name"),
			ClusterId:   types.StringValue("123"),
		},
		{
			ClusterName: types.StringValue("test-cluster-name"),
			ClusterId:   types.StringValue("456"),
		},
	}
	actualDiagnostics := validateClustersList(context.Background(), clusters, clusterName)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is more than one cluster with name '%s'", clusterName), "The IDs of those clusters are: 123, 456. When specifying a cluster name, the name must be unique. Alternatively, specify the cluster by ID using the cluster_id attribute.")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}
