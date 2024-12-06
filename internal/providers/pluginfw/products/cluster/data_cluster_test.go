package cluster

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
)

func TestNoClusterError(t *testing.T) {
	clusterName := "test-cluster-name"
	clusters := []compute.ClusterDetails{}
	actualDiagnostics := validateClustersList(context.Background(), clusters, clusterName)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no cluster with name '%s'", clusterName), "")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}

func TestMultipleClustersError(t *testing.T) {
	clusterName := "test-cluster-name"
	clusters := []compute.ClusterDetails{
		{
			ClusterName: "test-cluster-name",
			ClusterId:   "123",
		},
		{
			ClusterName: "test-cluster-name",
			ClusterId:   "456",
		},
	}
	actualDiagnostics := validateClustersList(context.Background(), clusters, clusterName)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is more than one cluster with name '%s'", clusterName), "The IDs of those clusters are: 123, 456. When specifying a cluster name, the name must be unique. Alternatively, specify the cluster by ID using the cluster_id attribute.")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}
