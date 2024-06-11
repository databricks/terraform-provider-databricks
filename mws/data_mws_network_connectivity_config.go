package mws

import (
	"context"
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfig() common.Resource {
	return common.AccountData(func(ctx context.Context, data *struct {
		Id                          string                   `json:"id" tf:"computed,optional"`
		AccountId                   string                   `json:"account_id,omitempty" tf:"computed"`
		CreationTime                int64                    `json:"creation_time,omitempty" tf:"computed"`
		EgressConfig                settings.NccEgressConfig `json:"egress_config,omitempty" tf:"computed"`
		Name                        string                   `json:"name,omitempty" tf:"computed"`
		NetworkConnectivityConfigId string                   `json:"network_connectivity_config_id"`
		Region                      string                   `json:"region,omitempty" tf:"computed"`
		UpdatedTime                 int64                    `json:"updated_time,omitempty" tf:"computed"`
	}, a *databricks.AccountClient) error {
		ncc, err := a.NetworkConnectivity.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(ctx, data.NetworkConnectivityConfigId)
		if err != nil {
			return err
		}
		data.Id = ncc.NetworkConnectivityConfigId
		data.AccountId = ncc.AccountId
		data.CreationTime = ncc.CreationTime
		data.EgressConfig = settings.NccEgressConfig{
			DefaultRules: ncc.EgressConfig.DefaultRules,
			TargetRules:  ncc.EgressConfig.TargetRules,
		}
		data.Name = ncc.Name
		data.Region = ncc.Region
		data.UpdatedTime = ncc.UpdatedTime
		return nil
	})
}
