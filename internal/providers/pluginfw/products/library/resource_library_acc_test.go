package library_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var commonClusterConfig = `data "databricks_spark_version" "latest" {
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

func TestAccLibraryReinstalledIfClusterDeleted(t *testing.T) {
	var clusterId string
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					repo = "https://pypi.org/dummy"
					package = "databricks-sdk"
				}
		    }`,
			Check: func(s *terraform.State) error {
				clusterId = s.RootModule().Resources["databricks_cluster.this"].Primary.ID
				return nil
			},
		},
		// If the cluster is deleted before apply, it should be recreated and the library reinstalled on the new cluster.
		acceptance.Step{
			PreConfig: func() {
				// Delete the created cluster
				w := databricks.Must(databricks.NewWorkspaceClient())
				w.Clusters.PermanentDeleteByClusterId(context.Background(), clusterId)
				// Wait for the cluster to be completely deleted
				errClusterExists := errors.New("cluster still exists")
				retries.New[struct{}](retries.OnErrors(errClusterExists)).Wait(context.Background(), func(ctx context.Context) error {
					_, err := w.Clusters.GetByClusterId(context.Background(), clusterId)
					if err != nil && apierr.IsMissing(err) {
						return nil
					}
					if err != nil {
						return err
					}
					return errClusterExists
				})
			},
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					repo = "https://pypi.org/dummy"
					package = "databricks-sdk"
				}
		    }`,
		})
}

func TestAccLibraryInstallIfClusterTerminated(t *testing.T) {
	var clusterId string
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: commonClusterConfig,
			Check: func(s *terraform.State) error {
				clusterId = s.RootModule().Resources["databricks_cluster.this"].Primary.ID
				return nil
			},
		},
		// If the cluster is Terminated before apply, it should be restarted before installing library.
		acceptance.Step{
			PreConfig: func() {
				// Delete the created cluster
				w := databricks.Must(databricks.NewWorkspaceClient())
				getter, err := w.Clusters.Delete(context.Background(), compute.DeleteCluster{
					ClusterId: clusterId,
				})
				if err != nil {
					t.Fatalf("Error deleting cluster: %s", err)
				}
				_, err = getter.GetWithTimeout(60 * time.Minute)
				if err != nil {
					t.Fatalf("Error waiting for cluster to be deleted: %s", err)
				}
			},
			Template: commonClusterConfig + `resource "databricks_library" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					repo = "https://pypi.org/dummy"
					package = "databricks-sdk"
				}
		    }`,
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

var sdkV2FallbackFactory = map[string]func() (tfprotov6.ProviderServer, error){
	"databricks": func() (tfprotov6.ProviderServer, error) {
		return providers.GetProviderServer(context.Background(), providers.WithSdkV2FallbackOptions(pluginfw.WithSdkV2ResourceFallbacks("databricks_library")))
	},
}

// Testing the transition from sdkv2 to plugin framework.
func TestAccLibraryUpdateTransitionFromSdkV2(t *testing.T) {
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
