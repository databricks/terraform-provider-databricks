package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfigs() common.Resource {
	type mwsNetworkConnectivityConfiguration struct {
		Names []string `json:"names" tf:"computed,optional"`
	}

	type mwsNetworkConnectivityConfigurationParams struct {
		Names  []string `json:"names" tf:"computed,optional"`
		Region string   `json:"region" tf:"optional"`
	}

	return common.AccountDataWithParams(func(ctx context.Context, data mwsNetworkConnectivityConfigurationParams, a *databricks.AccountClient) (*mwsNetworkConnectivityConfiguration, error) {
		list, err := a.NetworkConnectivity.ListNetworkConnectivityConfigurationsAll(ctx, settings.ListNetworkConnectivityConfigurationsRequest{})
		if err != nil {
			return nil, err
		}

		if data.Region != "" {
			filtered := []string{}
			for _, ncc := range list {
				if data.Region == ncc.Region {
					filtered = append(filtered, ncc.Name)
				}
			}
			return &mwsNetworkConnectivityConfiguration{Names: filtered}, nil
		}

		names := []string{}
		for _, ncc := range list {
			names = append(names, ncc.Name)
		}
		return &mwsNetworkConnectivityConfiguration{Names: names}, nil
	})
}
