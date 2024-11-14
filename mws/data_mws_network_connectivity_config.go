package mws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfig() common.Resource {
	type mwsNetworkConnectivityConfiguration struct {
		settings.NetworkConnectivityConfiguration
	}

	type mwsNetworkConnectivityConfigurationParams struct {
		Name string `json:"name" tf:"computed,optional"`
	}

	return common.AccountDataWithParams(func(ctx context.Context, data mwsNetworkConnectivityConfigurationParams, a *databricks.AccountClient) (*mwsNetworkConnectivityConfiguration, error) {
		if data.Name == "" {
			return nil, fmt.Errorf("`name` should be provided")
		}

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
