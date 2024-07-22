package mws

import (
	"context"
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsNetworkConnectivityConfig() common.Resource {
	type MwsNetworkConnectivity struct {
		settings.NetworkConnectivityConfiguration
	}

	type MwsNetworkConnectivityParams struct {
		Name   string `json:"name" tf:"computed,optional"`
		Region string `json:"region" tf:"computed,optional"`
	}

	return common.AccountDataWithParams(func(ctx context.Context, data MwsNetworkConnectivityParams, a *databricks.AccountClient) (*MwsNetworkConnectivity, error) {
		nccList := a.NetworkConnectivity.ListNetworkConnectivityConfigurations(ctx, settings.ListNetworkConnectivityConfigurationsRequest{})
		for nccList.HasNext(ctx) {
			ncc, err := nccList.Next(ctx)

			if err != nil {
				return nil, err
			}
			if data.Region == ncc.Region || data.Name == ncc.Name {
				return &MwsNetworkConnectivity{NetworkConnectivityConfiguration: ncc}, nil
			}
		}
		return nil, nil
	})
}
