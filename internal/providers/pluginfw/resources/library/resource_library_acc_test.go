package library_test

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var commonClusterConfig = `resource "data "databricks_spark_version" "latest" {
}
resource "databricks_cluster" "this" {
	cluster_name = "test-library-{var.RANDOM}"
	spark_version = data.databricks_spark_version.latest.id
	instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
	autotermination_minutes = 10
	num_workers = 0
	spark_conf = {
		"spark.databricks.cluster.profile" = "singleNode"
		"spark.master" = "local[*]"
	}
	custom_tags = {
		"ResourceClass" = "SingleNode"
	}
}

`

func TestAccLibraryCreation(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: commonClusterConfig + `resource "databricks_library" "new_library" {
			cluster_id = databricks_cluster.this.id
			pypi {
				repo = "https://pypi.org/dummy"
				package = "databricks-sdk"
			}
		}
		`,
	})
}

func TestAccLibraryUpdate(t *testing.T) {
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
					cluster_id = databricks_cluster.this.id
					pypi {
						repo = "https://pypi.org/simple"
						package = "databricks-sdk"
					}
				}
				`,
		},
		acceptance.Step{
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					package = "networkx"
				}
			}
			`,
		},
	)
}

// Testing the transition from sdkv2 to plugin framework.
func TestAccLibraryUpdateTransitionFromSdkV2(t *testing.T) {
	sdkV2FallbackFactory := map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			return providers.GetProviderServer(context.Background(), providers.WithSdkV2FallbackOptions([]pluginfw.SdkV2FallbackOption{pluginfw.WithSdkV2ResourceFallbacks([]string{"databricks_library"})}))
		},
	}
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			ProtoV6ProviderFactories: sdkV2FallbackFactory,
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
					cluster_id = databricks_cluster.this.id
					pypi {
						repo = "https://pypi.org/simple"
						package = "databricks-sdk"
					}
				}
				`,
		},
		acceptance.Step{
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					package = "networkx"
				}
			}
			`,
		},
	)
}

// Testing the transition from plugin framework to sdkv2.
func TestAccLibraryUpdateTransitionFromPluginFw(t *testing.T) {
	sdkV2FallbackFactory := map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			return providers.GetProviderServer(context.Background(), providers.WithSdkV2FallbackOptions([]pluginfw.SdkV2FallbackOption{pluginfw.WithSdkV2ResourceFallbacks([]string{"databricks_library"})}))
		},
	}
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
					cluster_id = databricks_cluster.this.id
					pypi {
						repo = "https://pypi.org/simple"
						package = "databricks-sdk"
					}
				}
				`,
		},
		acceptance.Step{
			ProtoV6ProviderFactories: sdkV2FallbackFactory,
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					package = "networkx"
				}
			}
			`,
		},
	)
}
