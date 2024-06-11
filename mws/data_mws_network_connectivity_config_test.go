package mws

import (
	"fmt"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceMwsNetworkConnectivityConfig(t *testing.T) {
	id := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	var ncc = settings.NetworkConnectivityConfiguration{
		AccountId:    "abc",
		CreationTime: 0,
		EgressConfig: &settings.NccEgressConfig{
			DefaultRules: nil,
			TargetRules:  nil,
		},
		Name:                        "def",
		NetworkConnectivityConfigId: "",
		Region:                      "us-east-1",
		UpdatedTime:                 0,
	}

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/accounts/%s/network-connectivity-configs/%s?", ncc.AccountId, id),
				Response: ncc,
			},
		},
		AccountID:   "abc",
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceMwsNetworkConnectivityConfig(),
		ID:          id,
		HCL: fmt.Sprintf(`
		network_connectivity_config_id = "%s"
		`, id),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"account_id":                     "abc",
		"creation_time":                  0,
		"egress_config":                  []interface{}{map[string]interface{}{"default_rules": []interface{}{}, "target_rules": []interface{}{}}},
		"id":                             id,
		"name":                           "def",
		"network_connectivity_config_id": id,
		"region":                         "us-east-1",
		"updated_time":                   0,
	})
}

func TestDataSourceMwsNetworkConnectivityConfig_AccountID(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Resource:    DataSourceMwsNetworkConnectivityConfig(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "invalid Databricks Account configuration")
}

func TestDataSourceMwsNetworkConnectivityConfig_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		AccountID:   "abc",
		Resource:    DataSourceMwsNetworkConnectivityConfig(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
