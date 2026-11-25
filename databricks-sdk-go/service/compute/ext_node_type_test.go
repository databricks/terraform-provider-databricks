package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeType(t *testing.T) {
	lst := ListNodeTypesResponse{
		NodeTypes: []NodeType{
			{
				NodeTypeId:     "m-fleet.xlarge",
				InstanceTypeId: "m-fleet.xlarge",
				MemoryMb:       16384,
				NumCores:       4,
			},
			{
				NodeTypeId: "Random_05",
				MemoryMb:   1024,
				NumCores:   32,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId: "Standard_L80s_v2",
				MemoryMb:   655360,
				NumCores:   80,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      2,
					InstanceTypeId:  "Standard_L80s_v2",
					LocalDiskSizeGb: 160,
					LocalNvmeDisks:  1,
				},
			},
			{
				NodeTypeId: "Random_01",
				MemoryMb:   8192,
				NumCores:   8,
				NodeInstanceType: &NodeInstanceType{
					InstanceTypeId: "_",
				},
			},
			{
				NodeTypeId: "Random_02",
				MemoryMb:   8192,
				NumCores:   8,
				NumGpus:    2,
				NodeInstanceType: &NodeInstanceType{
					InstanceTypeId: "_",
				},
			},
			{
				NodeTypeId: "Random_03",
				MemoryMb:   8192,
				NumCores:   8,
				NumGpus:    1,
				NodeInstanceType: &NodeInstanceType{
					InstanceTypeId:      "_",
					LocalNvmeDisks:      15,
					LocalNvmeDiskSizeGb: 235,
				},
			},
			{
				NodeTypeId:   "Random_04",
				MemoryMb:     32000,
				NumCores:     32,
				IsDeprecated: true,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      2,
					LocalDiskSizeGb: 20,
				},
			},
			{
				NodeTypeId: "Standard_F4s",
				MemoryMb:   8192,
				NumCores:   4,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      1,
					LocalDiskSizeGb: 16,
					LocalNvmeDisks:  0,
				},
			},
		},
	}
	nt, err := lst.Smallest(NodeTypeRequest{LocalDiskMinSize: 200, MinMemoryGB: 8, MinCores: 8, MinGPUs: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Random_03", nt)
}

func TestNodeTypeCategory(t *testing.T) {
	lst := ListNodeTypesResponse{
		NodeTypes: []NodeType{
			{
				NodeTypeId:     "Random_05",
				InstanceTypeId: "Random_05",
				MemoryMb:       1024,
				NumCores:       32,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId:     "Random_01",
				InstanceTypeId: "Random_01",
				MemoryMb:       8192,
				NumCores:       8,
				NodeInstanceType: &NodeInstanceType{
					InstanceTypeId: "_",
				},
				Category: "Memory Optimized",
			},
			{
				NodeTypeId:     "Random_02",
				InstanceTypeId: "Random_02",
				MemoryMb:       8192,
				NumCores:       8,
				Category:       "Storage Optimized",
			},
		},
	}
	nt, err := lst.Smallest(NodeTypeRequest{Category: "Storage optimized"})
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", nt)
}

func TestNodeTypeCategoryNotAvailable(t *testing.T) {
	lst := ListNodeTypesResponse{
		NodeTypes: []NodeType{
			{
				NodeTypeId:     "Random_05",
				InstanceTypeId: "Random_05",
				MemoryMb:       1024,
				NumCores:       32,
				NodeInstanceType: &NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId:     "Random_01",
				InstanceTypeId: "Random_01",
				MemoryMb:       8192,
				NumCores:       8,
				NodeInstanceType: &NodeInstanceType{
					InstanceTypeId: "_",
				},
				NodeInfo: &CloudProviderNodeInfo{
					Status: []CloudProviderNodeStatus{
						CloudProviderNodeStatusNotAvailableInRegion,
						CloudProviderNodeStatusNotEnabledOnSubscription},
				},
				Category: "Storage Optimized",
			},
			{
				NodeTypeId:     "Random_02",
				InstanceTypeId: "Random_02",
				MemoryMb:       8192,
				NumCores:       8,
				Category:       "Storage Optimized",
			},
		},
	}
	nt, err := lst.Smallest(NodeTypeRequest{Category: "Storage optimized"})
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", nt)
}

func TestNodeTypeFleet(t *testing.T) {
	lst := ListNodeTypesResponse{
		NodeTypes: []NodeType{
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
	}
	nt, err := lst.Smallest(NodeTypeRequest{Fleet: true, MinCores: 8})
	assert.NoError(t, err)
	assert.Equal(t, "m-fleet.2xlarge", nt)
}

func TestNodeTypeEmptyList(t *testing.T) {
	lst := ListNodeTypesResponse{
		NodeTypes: []NodeType{},
	}
	_, err := lst.Smallest(NodeTypeRequest{Fleet: true})
	assert.ErrorContains(t, err, "cannot determine smallest node type with empty response")
}
