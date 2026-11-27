package compute

import (
	"context"
	"fmt"
	"strings"
)

// NodeTypeRequest is a wrapper for local filtering of node types
type NodeTypeRequest struct {
	Id                    string `json:"id,omitempty"`
	MinMemoryGB           int32  `json:"min_memory_gb,omitempty"`
	GBPerCore             int32  `json:"gb_per_core,omitempty"`
	MinCores              int32  `json:"min_cores,omitempty"`
	MinGPUs               int32  `json:"min_gpus,omitempty"`
	LocalDisk             bool   `json:"local_disk,omitempty"`
	LocalDiskMinSize      int32  `json:"local_disk_min_size,omitempty"`
	Category              string `json:"category,omitempty"`
	PhotonWorkerCapable   bool   `json:"photon_worker_capable,omitempty"`
	PhotonDriverCapable   bool   `json:"photon_driver_capable,omitempty"`
	Graviton              bool   `json:"graviton,omitempty"`
	IsIOCacheEnabled      bool   `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool   `json:"support_port_forwarding,omitempty"`
	Fleet                 bool   `json:"fleet,omitempty"`
}

// sort NodeTypes within this struct
func (ntl *ListNodeTypesResponse) sort() {
	sortByChain(ntl.NodeTypes, func(i int) sortCmp {
		var localDisks, localDiskSizeGB, localNVMeDisk, localNVMeDiskSizeGB int32
		if ntl.NodeTypes[i].NodeInstanceType != nil {
			localDisks = int32(ntl.NodeTypes[i].NodeInstanceType.LocalDisks)
			localNVMeDisk = int32(ntl.NodeTypes[i].NodeInstanceType.LocalNvmeDisks)
			localDiskSizeGB = int32(ntl.NodeTypes[i].NodeInstanceType.LocalDiskSizeGb)
			localNVMeDiskSizeGB = int32(ntl.NodeTypes[i].NodeInstanceType.LocalNvmeDiskSizeGb)
		}
		return sortChain{
			boolAsc(ntl.NodeTypes[i].IsDeprecated),
			intAsc(ntl.NodeTypes[i].NumCores),
			intAsc(ntl.NodeTypes[i].MemoryMb),
			intAsc(localDisks),
			intAsc(localDiskSizeGB),
			intAsc(localNVMeDisk),
			intAsc(localNVMeDiskSizeGB),
			intAsc(ntl.NodeTypes[i].NumGpus),
			strAsc(ntl.NodeTypes[i].InstanceTypeId),
		}
	})
}

func (nt NodeType) shouldBeSkipped() bool {
	if nt.NodeInfo == nil {
		return false
	}
	for _, st := range nt.NodeInfo.Status {
		switch st {
		case CloudProviderNodeStatusNotAvailableInRegion, CloudProviderNodeStatusNotEnabledOnSubscription:
			return true
		}
	}
	return false
}

func (ntl *ListNodeTypesResponse) Smallest(r NodeTypeRequest) (string, error) {
	// error is explicitly ingored here, because Azure returns
	// apparently too big of a JSON for Go to parse
	if len(ntl.NodeTypes) == 0 {
		return "", fmt.Errorf("cannot determine smallest node type with empty response")
	}
	ntl.sort()
	for _, nt := range ntl.NodeTypes {
		if nt.shouldBeSkipped() {
			continue
		}
		gbs := int32(nt.MemoryMb / 1024)
		if r.Fleet != strings.Contains(nt.NodeTypeId, "-fleet.") {
			continue
		}
		if r.MinMemoryGB > 0 && gbs < r.MinMemoryGB {
			continue
		}
		if r.GBPerCore > 0 && (gbs/int32(nt.NumCores)) < r.GBPerCore {
			continue
		}
		if r.MinCores > 0 && int32(nt.NumCores) < r.MinCores {
			continue
		}
		if (r.MinGPUs > 0 && int32(nt.NumGpus) < r.MinGPUs) || (r.MinGPUs == 0 && nt.NumGpus > 0) {
			continue
		}
		if (r.LocalDisk || r.LocalDiskMinSize > 0) && nt.NodeInstanceType != nil &&
			(nt.NodeInstanceType.LocalDisks < 1 &&
				nt.NodeInstanceType.LocalNvmeDisks < 1) {
			continue
		}
		if r.LocalDiskMinSize > 0 && nt.NodeInstanceType != nil &&
			(int32(nt.NodeInstanceType.LocalDiskSizeGb)+int32(nt.NodeInstanceType.LocalNvmeDiskSizeGb)) < r.LocalDiskMinSize {
			continue
		}
		if r.Category != "" && !strings.EqualFold(nt.Category, r.Category) {
			continue
		}
		if r.IsIOCacheEnabled && !nt.IsIoCacheEnabled {
			continue
		}
		if r.SupportPortForwarding && !nt.SupportPortForwarding {
			continue
		}
		if r.PhotonDriverCapable && !nt.PhotonDriverCapable {
			continue
		}
		if r.PhotonWorkerCapable && !nt.PhotonWorkerCapable {
			continue
		}
		if nt.IsGraviton != r.Graviton {
			continue
		}
		return nt.NodeTypeId, nil
	}
	return "", fmt.Errorf("cannot determine smallest node type")
}

func (a *ClustersAPI) SelectNodeType(ctx context.Context, r NodeTypeRequest) (string, error) {
	nodeTypes, err := a.ListNodeTypes(ctx)
	if err != nil {
		return "", err
	}
	return nodeTypes.Smallest(r)
}
