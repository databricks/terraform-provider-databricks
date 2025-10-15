package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfig() common.Resource {
	type mwsNetworkConnectivityConfiguration struct {
		common.Namespace
		settings.NetworkConnectivityConfiguration
	}

	type mwsNetworkConnectivityConfigurationParams struct {
		common.Namespace
		Name string `json:"name"`
	}

	return common.AccountDataWithParams(func(ctx context.Context, data mwsNetworkConnectivityConfigurationParams, a *databricks.AccountClient) (*mwsNetworkConnectivityConfiguration, error) {
		list, err := a.NetworkConnectivity.ListNetworkConnectivityConfigurationsAll(ctx, settings.ListNetworkConnectivityConfigurationsRequest{})
		if err != nil {
			return nil, err
		}

		for _, ncc := range list {
			if data.Name == ncc.Name {
				return &mwsNetworkConnectivityConfiguration{NetworkConnectivityConfiguration: ncc}, nil
			}
		}
		return nil, nil
	})
}
