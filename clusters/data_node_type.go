package clusters

import (
	"context"
	"sort"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NodeTypeRequest is a wrapper for local filtering of node types
type NodeTypeRequest struct {
	MinMemoryGB           int32  `json:"min_memory_gb,omitempty"`
	GBPerCore             int32  `json:"gb_per_core,omitempty"`
	MinCores              int32  `json:"min_cores,omitempty"`
	MinGPUs               int32  `json:"min_gpus,omitempty"`
	LocalDisk             bool   `json:"local_disk,omitempty"`
	Category              string `json:"category,omitempty"`
	PhotonWorkerCapable   bool   `json:"photon_worker_capable,omitempty"`
	PhotonDriverCapable   bool   `json:"photon_driver_capable,omitempty"`
	Graviton              bool   `json:"graviton,omitempty"`
	IsIOCacheEnabled      bool   `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool   `json:"support_port_forwarding,omitempty"`
}

// NodeTypeList contains a list of node types
type NodeTypeList struct {
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

// Sort NodeTypes within this struct
func (l *NodeTypeList) Sort() {
	sort.Slice(l.NodeTypes, func(i, j int) bool {
		if l.NodeTypes[i].IsDeprecated != l.NodeTypes[j].IsDeprecated {
			return !l.NodeTypes[i].IsDeprecated
		}
		if l.NodeTypes[i].NodeInstanceType != nil &&
			l.NodeTypes[j].NodeInstanceType != nil {
			if l.NodeTypes[i].NodeInstanceType.LocalDisks !=
				l.NodeTypes[j].NodeInstanceType.LocalDisks {
				return l.NodeTypes[i].NodeInstanceType.LocalDisks <
					l.NodeTypes[j].NodeInstanceType.LocalDisks
			}
			if l.NodeTypes[i].NodeInstanceType.LocalDiskSizeGB !=
				l.NodeTypes[j].NodeInstanceType.LocalDiskSizeGB {
				return l.NodeTypes[i].NodeInstanceType.LocalDiskSizeGB <
					l.NodeTypes[j].NodeInstanceType.LocalDiskSizeGB
			}
		}
		if l.NodeTypes[i].MemoryMB != l.NodeTypes[j].MemoryMB {
			return l.NodeTypes[i].MemoryMB < l.NodeTypes[j].MemoryMB
		}
		if l.NodeTypes[i].NumCores != l.NodeTypes[j].NumCores {
			return l.NodeTypes[i].NumCores < l.NodeTypes[j].NumCores
		}
		if l.NodeTypes[i].NumGPUs != l.NodeTypes[j].NumGPUs {
			return l.NodeTypes[i].NumGPUs < l.NodeTypes[j].NumGPUs
		}
		return l.NodeTypes[i].InstanceTypeID < l.NodeTypes[j].InstanceTypeID
	})
}

// ClusterCloudProviderNodeInfo encapsulates the existing quota available from the cloud service provider.
type ClusterCloudProviderNodeInfo struct {
	Status             []string `json:"status,omitempty"`
	AvailableCoreQuota float32  `json:"available_core_quota,omitempty"`
	TotalCoreQuota     float32  `json:"total_core_quota,omitempty"`
}

// NodeInstanceType encapsulates information about a specific node type
type NodeInstanceType struct {
	InstanceTypeID      string `json:"instance_type_id,omitempty"`
	LocalDisks          int32  `json:"local_disks,omitempty"`
	LocalDiskSizeGB     int32  `json:"local_disk_size_gb,omitempty"`
	LocalNVMeDisks      int32  `json:"local_nvme_disks,omitempty"`
	LocalNVMeDiskSizeGB int32  `json:"local_nvme_disk_size_gb,omitempty"`
}

// NodeType encapsulates information about a given node when using the list-node-types api
type NodeType struct {
	NodeTypeID            string                        `json:"node_type_id,omitempty"`
	MemoryMB              int32                         `json:"memory_mb,omitempty"`
	NumCores              float32                       `json:"num_cores,omitempty"`
	NumGPUs               int32                         `json:"num_gpus,omitempty"`
	SupportEBSVolumes     bool                          `json:"support_ebs_volumes,omitempty"`
	IsIOCacheEnabled      bool                          `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool                          `json:"support_port_forwarding,omitempty"`
	Description           string                        `json:"description,omitempty"`
	Category              string                        `json:"category,omitempty"`
	InstanceTypeID        string                        `json:"instance_type_id,omitempty"`
	IsDeprecated          bool                          `json:"is_deprecated,omitempty"`
	IsHidden              bool                          `json:"is_hidden,omitempty"`
	SupportClusterTags    bool                          `json:"support_cluster_tags,omitempty"`
	DisplayOrder          int32                         `json:"display_order,omitempty"`
	NodeInfo              *ClusterCloudProviderNodeInfo `json:"node_info,omitempty"`
	NodeInstanceType      *NodeInstanceType             `json:"node_instance_type,omitempty"`
	PhotonWorkerCapable   bool                          `json:"photon_worker_capable,omitempty"`
	PhotonDriverCapable   bool                          `json:"photon_driver_capable,omitempty"`
	Graviton              bool                          `json:"is_graviton,omitempty"`
}

func (a ClustersAPI) defaultSmallestNodeType() string {
	if a.client.IsAzure() {
		return "Standard_D3_v2"
	} else if a.client.IsGcp() {
		return "n1-standard-4"
	}
	return "i3.xlarge"
}

// ListNodeTypes returns a sorted list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() (l NodeTypeList, err error) {
	err = a.client.Get(a.context, "/clusters/list-node-types", nil, &l)
	return
}

// GetSmallestNodeType returns smallest (or default) node type id given the criteria
func (a ClustersAPI) GetSmallestNodeType(r NodeTypeRequest) string {
	list, _ := a.ListNodeTypes()
	// error is explicitly ingored here, because Azure returns
	// apparently too big of a JSON for Go to parse
	if len(list.NodeTypes) == 0 {
		return a.defaultSmallestNodeType()
	}
	list.Sort()
	for _, nt := range list.NodeTypes {
		gbs := (nt.MemoryMB / 1024)
		if r.MinMemoryGB > 0 && gbs < r.MinMemoryGB {
			continue
		}
		if r.GBPerCore > 0 && (gbs/int32(nt.NumCores)) < r.GBPerCore {
			continue
		}
		if r.MinCores > 0 && int32(nt.NumCores) < r.MinCores {
			continue
		}
		if r.MinGPUs > 0 && nt.NumGPUs < r.MinGPUs {
			continue
		}
		if r.LocalDisk && nt.NodeInstanceType != nil &&
			(nt.NodeInstanceType.LocalDisks < 1 &&
				nt.NodeInstanceType.LocalNVMeDisks < 1) {
			continue
		}
		if r.Category != "" && !strings.EqualFold(nt.Category, r.Category) {
			continue
		}
		if r.IsIOCacheEnabled && nt.IsIOCacheEnabled != r.IsIOCacheEnabled {
			continue
		}
		if r.SupportPortForwarding && nt.SupportPortForwarding != r.SupportPortForwarding {
			continue
		}
		if r.PhotonDriverCapable && nt.PhotonDriverCapable != r.PhotonDriverCapable {
			continue
		}
		if r.PhotonWorkerCapable && nt.PhotonWorkerCapable != r.PhotonWorkerCapable {
			continue
		}
		if r.Graviton && nt.Graviton != r.Graviton {
			continue
		}
		return nt.NodeTypeID
	}
	return a.defaultSmallestNodeType()
}

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() *schema.Resource {
	s := common.StructToSchema(NodeTypeRequest{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData,
			m interface{}) diag.Diagnostics {
			var this NodeTypeRequest
			common.DataToStructPointer(d, s, &this)
			clustersAPI := NewClustersAPI(ctx, m)
			d.SetId(clustersAPI.GetSmallestNodeType(this))
			return nil
		},
	}
}
