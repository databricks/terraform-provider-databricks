package clusters

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestNodeType(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response: compute.ListNodeTypesResponse{
					NodeTypes: []compute.NodeType{
						{
							NodeTypeId:     "m-fleet.xlarge",
							InstanceTypeId: "m-fleet.xlarge",
							MemoryMb:       16384,
							NumCores:       4,
						},
						{
							NodeTypeId:     "Random_05",
							InstanceTypeId: "Random_05",
							MemoryMb:       1024,
							NumCores:       32,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGb: 100,
							},
						},
						{
							NodeTypeId:     "Standard_L80s_v2",
							InstanceTypeId: "Standard_L80s_v2",
							MemoryMb:       655360,
							NumCores:       80,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      2,
								InstanceTypeId:  "Standard_L80s_v2",
								LocalDiskSizeGb: 160,
								LocalNvmeDisks:  1,
							},
						},
						{
							NodeTypeId:     "Random_01",
							InstanceTypeId: "Random_01",
							MemoryMb:       8192,
							NumCores:       8,
							NodeInstanceType: &compute.NodeInstanceType{
								InstanceTypeId: "_",
							},
						},
						{
							NodeTypeId:     "Random_02",
							InstanceTypeId: "Random_02",
							MemoryMb:       8192,
							NumCores:       8,
							NumGpus:        2,
							NodeInstanceType: &compute.NodeInstanceType{
								InstanceTypeId: "_",
							},
						},
						{
							NodeTypeId:     "Random_03",
							InstanceTypeId: "Random_03",
							MemoryMb:       8192,
							NumCores:       8,
							NumGpus:        1,
							NodeInstanceType: &compute.NodeInstanceType{
								InstanceTypeId:      "_",
								LocalNvmeDisks:      15,
								LocalNvmeDiskSizeGb: 235,
							},
						},
						{
							NodeTypeId:     "Random_04",
							InstanceTypeId: "Random_04",
							MemoryMb:       32000,
							NumCores:       32,
							IsDeprecated:   true,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      2,
								LocalDiskSizeGb: 20,
							},
						},
						{
							NodeTypeId:     "Standard_F4s",
							InstanceTypeId: "Standard_F4s",
							MemoryMb:       8192,
							NumCores:       4,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      1,
								LocalDiskSizeGb: 16,
								LocalNvmeDisks:  0,
							},
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]any{
			"local_disk_min_size": 200,
			"min_memory_gb":       8,
			"min_cores":           8,
			"min_gpus":            1,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "Random_03", d.Id())
}

func TestNodeTypeCategory(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response: compute.ListNodeTypesResponse{
					NodeTypes: []compute.NodeType{
						{
							NodeTypeId:     "Random_05",
							InstanceTypeId: "Random_05",
							MemoryMb:       1024,
							NumCores:       32,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGb: 100,
							},
						},
						{
							NodeTypeId:     "Random_01",
							InstanceTypeId: "Random_01",
							MemoryMb:       8192,
							NumCores:       8,
							NodeInstanceType: &compute.NodeInstanceType{
								InstanceTypeId: "_",
							},
							Category: "Memory Optimized",
						},
						{
							NodeTypeId:     "Random_02_GPU",
							InstanceTypeId: "Random_02_GPU",
							MemoryMb:       8192,
							NumCores:       8,
							NumGpus:        2,
							Category:       "Storage Optimized",
						},
						{
							NodeTypeId:     "Random_02",
							InstanceTypeId: "Random_02",
							MemoryMb:       8192,
							NumCores:       8,
							Category:       "Storage Optimized",
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]any{
			"category": "Storage optimized",
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", d.Id())
}

// func TestSmallestNodeTypeClouds(t *testing.T) {
// 	assert.Equal(t, "Standard_D3_v2", ClustersAPI{
// 		client: &common.DatabricksClient{
// 			DatabricksClient: &client.DatabricksClient{
// 				Config: &config.Config{
// 					Host: "foo.azuredatabricks.net",
// 				},
// 			},
// 		},
// 	}.defaultSmallestNodeType())

// 	assert.Equal(t, "n1-standard-4", ClustersAPI{
// 		client: &common.DatabricksClient{
// 			DatabricksClient: &client.DatabricksClient{
// 				Config: &config.Config{
// 					Host: "foo.gcp.databricks.com",
// 				},
// 			},
// 		},
// 	}.defaultSmallestNodeType())

// 	assert.Equal(t, "i3.xlarge", ClustersAPI{
// 		client: &common.DatabricksClient{
// 			DatabricksClient: &client.DatabricksClient{
// 				Config: &config.Config{
// 					Host: "foo.cloud.databricks.com",
// 				},
// 			},
// 		},
// 	}.defaultSmallestNodeType())
// }

func TestNodeTypeCategoryNotAvailable(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response: compute.ListNodeTypesResponse{
					NodeTypes: []compute.NodeType{
						{
							NodeTypeId:     "Random_05",
							InstanceTypeId: "Random_05",
							MemoryMb:       1024,
							NumCores:       32,
							NodeInstanceType: &compute.NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGb: 100,
							},
						},
						{
							NodeTypeId:     "Random_01",
							InstanceTypeId: "Random_01",
							MemoryMb:       8192,
							NumCores:       8,
							NodeInstanceType: &compute.NodeInstanceType{
								InstanceTypeId: "_",
							},
							Category: "Memory Optimized",
						},
						{
							NodeTypeId:     "Random_02",
							InstanceTypeId: "Random_02",
							MemoryMb:       8192,
							NumCores:       8,
							NumGpus:        2,
							Category:       "Storage Optimized",
							NodeInfo: &compute.CloudProviderNodeInfo{
								Status: []compute.CloudProviderNodeStatus{
									compute.CloudProviderNodeStatusNotAvailableInRegion,
									compute.CloudProviderNodeStatusNotEnabledOnSubscription,
								},
							},
						},
						{
							NodeTypeId:     "Random_03",
							InstanceTypeId: "Random_03",
							MemoryMb:       8192,
							NumCores:       8,
							Category:       "Storage Optimized",
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]any{
			"category": "Storage optimized",
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "Random_03", d.Id())
}

func TestNodeTypeFleet(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response: compute.ListNodeTypesResponse{
					NodeTypes: []compute.NodeType{
						{
							NodeTypeId:     "Random_05",
							InstanceTypeId: "Random_05",
							MemoryMb:       1024,
							NumCores:       4,
						},
						{
							NodeTypeId:     "m-fleet.xlarge",
							InstanceTypeId: "m-fleet.xlarge",
							MemoryMb:       16384,
							NumCores:       4,
						},
						{
							NodeTypeId:     "m-fleet.2xlarge",
							InstanceTypeId: "m-fleet.2xlarge",
							MemoryMb:       32768,
							NumCores:       8,
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]any{
			"fleet":     true,
			"min_cores": 8,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "m-fleet.2xlarge", d.Id())
}

func TestNodeTypeEmptyList(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response:     compute.ListNodeTypesResponse{},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		Azure:       true,
		State:       map[string]any{},
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "Standard_D3_v2", d.Id())
}

func TestNodeTypeFleetEmptyList(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/list-node-types",
				Response:     compute.ListNodeTypesResponse{},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]any{
			"fleet": true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "md-fleet.xlarge", d.Id())
}
