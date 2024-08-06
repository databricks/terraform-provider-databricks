package pluginframework_acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"context"
)

var commonPartQualityMonitoring = `resource "databricks_catalog" "sandbox" {
	name         = "sandbox{var.STICKY_RANDOM}"
	comment      = "this catalog is managed by terraform"
	properties = {
		purpose = "testing"
	}
	force_destroy = true
}

resource "databricks_schema" "things" {
	catalog_name = databricks_catalog.sandbox.id
	name         = "things{var.STICKY_RANDOM}"
	comment      = "this database is managed by terraform"
	properties = {
		kind = "various"
	}
}

resource "databricks_sql_table" "myInferenceTable" {
	catalog_name = databricks_catalog.sandbox.id
	schema_name = databricks_schema.things.name
	name = "bar{var.STICKY_RANDOM}_inference"
	table_type = "MANAGED"
	data_source_format = "DELTA"
	
	column {
		name = "model_id"
		type = "int"
	}
	column {
		name = "timestamp"
		type = "int"
	}
	column {
		name = "prediction"
		type = "int"
	}
}

`

func TestUcAccUpdateQualityMonitorPluginFramework(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_quality_monitor resource is not available on GCP")
	}
	resource.Test(t, resource.TestCase{
		// ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
		// 	"databricks": func() (tfprotov6.ProviderServer, error) {
		// 		ctx := context.Background()

		// 		sdkPluginProvider := provider.DatabricksProvider()

		// 		upgradedSdkServer, err := tf5to6server.UpgradeServer(
		// 			context.Background(),
		// 			sdkPluginProvider.GRPCProvider,
		// 		)

		// 		if err != nil {
		// 			return nil, err
		// 		}

		// 		providers := []func() tfprotov6.ProviderServer{
		// 			providerserver.NewProtocol6(provider.GetDatabricksProviderPluginFramework()), // Example terraform-plugin-framework provider
		// 			func() tfprotov6.ProviderServer {
		// 				return upgradedSdkServer
		// 			},
		// 		}

		// 		muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

		// 		if err != nil {
		// 			return nil, err
		// 		}

		// 		return muxServer.ProviderServer(), nil
		// 	},
		// },
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
					"databricks": func() (tfprotov6.ProviderServer, error) {
						ctx := context.Background()

						sdkPluginProvider := provider.DatabricksProvider()

						upgradedSdkServer, err := tf5to6server.UpgradeServer(
							context.Background(),
							sdkPluginProvider.GRPCProvider,
						)

						if err != nil {
							return nil, err
						}

						providers := []func() tfprotov6.ProviderServer{
							providerserver.NewProtocol6(provider.GetDatabricksProviderPluginFramework()), // Example terraform-plugin-framework provider
							func() tfprotov6.ProviderServer {
								return upgradedSdkServer
							},
						}

						muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

						if err != nil {
							return nil, err
						}

						return muxServer.ProviderServer(), nil
					},
				},
				Config: commonPartQualityMonitoring + `
				resource "databricks_lakehouse_monitor_pluginframework" "testMonitorInference" {
					table_name = databricks_sql_table.myInferenceTable.id
					assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myInferenceTable.name}"
					output_schema_name = databricks_schema.things.id
					inference_log = {
					  granularities = ["1 hour"]
					  timestamp_col = "timestamp"
					  prediction_col = "prediction"
					  model_id_col = "model_id"
					  problem_type = "PROBLEM_TYPE_REGRESSION"
					} 
				}
				`,
			},
		},
	})
}
