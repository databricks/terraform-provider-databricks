package compute

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
)

func TestGetOrCreateRunningCluster_Aws_NoAwsAttributes(t *testing.T) {
	client, server := qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/list?",
			Response: ListClustersResponse{},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/list-node-types",
			Response: ListNodeTypesResponse{
				NodeTypes: []NodeType{
					{
						NodeTypeId:     "m5.large",
						InstanceTypeId: "m5.large",
						MemoryMb:       8192,
						NumCores:       2,
						NodeInstanceType: &NodeInstanceType{
							LocalDisks:      1,
							InstanceTypeId:  "m5.large",
							LocalDiskSizeGb: 16,
						},
					},
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/spark-versions",
			Response: GetSparkVersionsResponse{
				Versions: []SparkVersion{
					{
						Key:  "7.3.x-scala2.12",
						Name: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.1/clusters/create",
			ExpectedRequest: CreateCluster{
				AutoterminationMinutes: 10,
				ClusterName:            "mount",
				NodeTypeId:             "m5.large",
				NumWorkers:             1,
				SparkVersion:           "7.3.x-scala2.12",
			},
			Response: CreateClusterResponse{
				ClusterId: "bcd",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.1/clusters/get?cluster_id=bcd",
			Response: ClusterDetails{
				State: StateRunning,
			},
		},
	}.Client(t)
	defer server.Close()

	ctx := context.Background()
	api := NewClusters(client)
	clusterInfo, err := api.GetOrCreateRunningCluster(ctx, "mount")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if clusterInfo == nil {
		t.Fatal("expected non-nil cluster info")
	}
	if clusterInfo.AwsAttributes != nil {
		t.Fatalf("expected AwsAttributes to be nil, got %v", clusterInfo.AwsAttributes)
	}
}
