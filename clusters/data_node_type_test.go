package clusters

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestNodeType(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/list-node-types",
				Response: NodeTypeList{
					[]NodeType{
						{
							NodeTypeID:     "vcpu-worker",
							InstanceTypeID: "vcpu-worker",
							MemoryMB:       0,
							NumCores:       0,
						},
						{
							NodeTypeID:     "Random_05",
							InstanceTypeID: "Random_05",
							MemoryMB:       1024,
							NumCores:       32,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGB: 100,
							},
						},
						{
							NodeTypeID:     "Standard_L80s_v2",
							InstanceTypeID: "Standard_L80s_v2",
							MemoryMB:       655360,
							NumCores:       80,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      2,
								InstanceTypeID:  "Standard_L80s_v2",
								LocalDiskSizeGB: 160,
								LocalNVMeDisks:  1,
							},
						},
						{
							NodeTypeID:     "Random_01",
							InstanceTypeID: "Random_01",
							MemoryMB:       8192,
							NumCores:       8,
							NodeInstanceType: &NodeInstanceType{
								InstanceTypeID: "_",
							},
						},
						{
							NodeTypeID:     "Random_02",
							InstanceTypeID: "Random_02",
							MemoryMB:       8192,
							NumCores:       8,
							NumGPUs:        2,
							NodeInstanceType: &NodeInstanceType{
								InstanceTypeID: "_",
							},
						},
						{
							NodeTypeID:     "Random_03",
							InstanceTypeID: "Random_03",
							MemoryMB:       8192,
							NumCores:       8,
							NumGPUs:        1,
							NodeInstanceType: &NodeInstanceType{
								InstanceTypeID:      "_",
								LocalNVMeDisks:      15,
								LocalNVMeDiskSizeGB: 235,
							},
						},
						{
							NodeTypeID:     "Random_04",
							InstanceTypeID: "Random_04",
							MemoryMB:       32000,
							NumCores:       32,
							IsDeprecated:   true,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      2,
								LocalDiskSizeGB: 20,
							},
						},
						{
							NodeTypeID:     "Standard_F4s",
							InstanceTypeID: "Standard_F4s",
							MemoryMB:       8192,
							NumCores:       4,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      1,
								LocalDiskSizeGB: 16,
								LocalNVMeDisks:  0,
							},
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]interface{}{
			"local_disk":    true,
			"min_memory_gb": 8,
			"min_cores":     8,
			"min_gpus":      1,
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
				Resource:     "/api/2.0/clusters/list-node-types",
				Response: NodeTypeList{
					[]NodeType{
						{
							NodeTypeID:     "Random_05",
							InstanceTypeID: "Random_05",
							MemoryMB:       1024,
							NumCores:       32,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGB: 100,
							},
						},
						{
							NodeTypeID:     "Random_01",
							InstanceTypeID: "Random_01",
							MemoryMB:       8192,
							NumCores:       8,
							NodeInstanceType: &NodeInstanceType{
								InstanceTypeID: "_",
							},
							Category: "Memory Optimized",
						},
						{
							NodeTypeID:     "Random_02",
							InstanceTypeID: "Random_02",
							MemoryMB:       8192,
							NumCores:       8,
							NumGPUs:        2,
							Category:       "Storage Optimized",
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]interface{}{
			"category": "Storage optimized",
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", d.Id())
}

func TestNodeTypeVCPU(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/list-node-types",
				Response: NodeTypeList{
					[]NodeType{
						{
							NodeTypeID:     "Random_05",
							InstanceTypeID: "Random_05",
							MemoryMB:       1024,
							NumCores:       32,
							NodeInstanceType: &NodeInstanceType{
								LocalDisks:      3,
								LocalDiskSizeGB: 100,
							},
						},
						{
							NodeTypeID:     "vcpu-worker",
							InstanceTypeID: "vcpu-worker",
							MemoryMB:       0,
							NumCores:       0,
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceNodeType(),
		NonWritable: true,
		State: map[string]interface{}{
			"vcpu": true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "vcpu-worker", d.Id())
}

func TestSmallestNodeTypeClouds(t *testing.T) {
	assert.Equal(t, "Standard_D3_v2", ClustersAPI{
		client: &common.DatabricksClient{
			Host: "foo.azuredatabricks.net",
		},
	}.defaultSmallestNodeType())

	assert.Equal(t, "n1-standard-4", ClustersAPI{
		client: &common.DatabricksClient{
			Host: "foo.gcp.databricks.com",
		},
	}.defaultSmallestNodeType())
}
