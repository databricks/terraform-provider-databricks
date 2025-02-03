package mws

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceMwsNetworkConnectivityConfig(t *testing.T) {
	ncc := settings.NetworkConnectivityConfiguration{
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
	}

	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockNetworkConnectivityAPI().EXPECT()
			api.ListNetworkConnectivityConfigurationsAll(mock.Anything, settings.ListNetworkConnectivityConfigurationsRequest{}).Return(
				[]settings.NetworkConnectivityConfiguration{ncc}, nil,
			)
		},
		AccountID:   "abc",
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceMwsNetworkConnectivityConfig(),
		ID:          "_",
		HCL: fmt.Sprintf(`
		    name = "%s"
		`, ncc.Name),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"account_id":                     "abc",
		"creation_time":                  0,
		"egress_config":                  []interface{}{map[string]interface{}{"default_rules": []interface{}{}, "target_rules": []interface{}{}}},
		"id":                             "_",
		"name":                           "def",
		"network_connectivity_config_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		"region":                         "us-east-1",
		"updated_time":                   0,
	})
}

func TestDataSourceMwsNetworkConnectivityConfig_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		AccountID:   "abc",
		Resource:    DataSourceMwsNetworkConnectivityConfig(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
        name = "def"
        `,
	}.ExpectError(t, "i'm a teapot")
}
