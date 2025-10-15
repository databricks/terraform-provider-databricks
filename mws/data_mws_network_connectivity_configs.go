package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfigs() common.Resource {
	type mwsNetworkConnectivityConfiguration struct {
		common.Namespace
		Names []string `json:"names" tf:"computed,optional"`
	}

	type mwsNetworkConnectivityConfigurationParams struct {
		common.Namespace
		Names  []string `json:"names" tf:"computed,optional"`
		Region string   `json:"region" tf:"optional"`
	}

	return common.AccountDataWithParams(func(ctx context.Context, data mwsNetworkConnectivityConfigurationParams, a *databricks.AccountClient) (*mwsNetworkConnectivityConfiguration, error) {
		list, err := a.NetworkConnectivity.ListNetworkConnectivityConfigurationsAll(ctx, settings.ListNetworkConnectivityConfigurationsRequest{})
		if err != nil {
			return nil, err
		}

		names := []string{}

		if data.Region != "" {
			for _, ncc := range list {
				if data.Region == ncc.Region {
					names = append(names, ncc.Name)
				}
			}
			return &mwsNetworkConnectivityConfiguration{Names: names}, nil
		}

		for _, ncc := range list {
			names = append(names, ncc.Name)
		}
		return &mwsNetworkConnectivityConfiguration{Names: names}, nil
	})
}
