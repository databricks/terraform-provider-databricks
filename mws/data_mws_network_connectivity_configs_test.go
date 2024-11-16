package mws

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func getTestNccs() []settings.NetworkConnectivityConfiguration {
	return []settings.NetworkConnectivityConfiguration{
		{
			AccountId:    "abc",
			CreationTime: 0,
			EgressConfig: &settings.NccEgressConfig{
				DefaultRules: nil,
				TargetRules:  nil,
			},
			Name:                        "def",
			NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			Region:                      "us-east-1",
			UpdatedTime:                 0,
		},
		{
			AccountId:    "abc",
			CreationTime: 0,
			EgressConfig: &settings.NccEgressConfig{
				DefaultRules: nil,
				TargetRules:  nil,
			},
			Name:                        "ghi",
			NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			Region:                      "us-west-1",
			UpdatedTime:                 0,
		},
	}
}

func TestDataSourceMwsNetworkConnectivityConfigs_All(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockNetworkConnectivityAPI().EXPECT()
			api.ListNetworkConnectivityConfigurationsAll(mock.Anything, settings.ListNetworkConnectivityConfigurationsRequest{}).Return(
				[]settings.NetworkConnectivityConfiguration{
					{
						AccountId:    "abc",
						CreationTime: 0,
						EgressConfig: &settings.NccEgressConfig{
							DefaultRules: nil,
							TargetRules:  nil,
						},
						Name:                        "def",
						NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Region:                      "us-east-1",
						UpdatedTime:                 0,
					},
					{
						AccountId:    "abc",
						CreationTime: 0,
						EgressConfig: &settings.NccEgressConfig{
							DefaultRules: nil,
							TargetRules:  nil,
						},
						Name:                        "ghi",
						NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Region:                      "us-west-1",
						UpdatedTime:                 0,
					},
				}, nil,
			)
		},
		AccountID:   "abc",
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceMwsNetworkConnectivityConfigs(),
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"names": []interface{}{"def", "ghi"},
	})
}

func TestDataSourceMwsNetworkConnectivityConfigs_Filter(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockNetworkConnectivityAPI().EXPECT()
			api.ListNetworkConnectivityConfigurationsAll(mock.Anything, settings.ListNetworkConnectivityConfigurationsRequest{}).Return(
				[]settings.NetworkConnectivityConfiguration{
					{
						AccountId:    "abc",
						CreationTime: 0,
						EgressConfig: &settings.NccEgressConfig{
							DefaultRules: nil,
							TargetRules:  nil,
						},
						Name:                        "def",
						NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Region:                      "us-east-1",
						UpdatedTime:                 0,
					},
					{
						AccountId:    "abc",
						CreationTime: 0,
						EgressConfig: &settings.NccEgressConfig{
							DefaultRules: nil,
							TargetRules:  nil,
						},
						Name:                        "def-3",
						NetworkConnectivityConfigId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
						Region:                      "us-west-1",
						UpdatedTime:                 0,
					},
				}, nil,
			)
		},
		AccountID:   "abc",
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceMwsNetworkConnectivityConfigs(),
		ID:          "_",
		HCL: fmt.Sprintf(`
        region = "%s"
        `, getTestNccs()[0].Region),
	}.ApplyAndExpectData(t, map[string]any{
		"names": []interface{}{"def"},
	})
}

func TestDataSourceMwsNetworkConnectivityConfigs_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		AccountID:   "abc",
		Resource:    DataSourceMwsNetworkConnectivityConfigs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
        region = "us-east-1"
        `,
	}.ExpectError(t, "i'm a teapot")
}
