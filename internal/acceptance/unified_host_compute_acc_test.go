package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ==========================================
// databricks_cluster (clusters/clusters_api.go)
//
// Uses the databricks_node_type / databricks_spark_version data sources so the
// config is cloud-agnostic (they resolve the cloud-correct node type at apply
// time). Those data sources are workspace-routed too, so they also carry the
// provider_config block on a unified host.
// ==========================================

func createClusterWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	clusterName := "tf-" + RandomName() + "-cluster"
	step := Step{
		Template: `
		data "databricks_spark_version" "latest" {
			` + pcBlock(workspaceID) + `
		}
		data "databricks_node_type" "smallest" {
			local_disk = true
			` + pcBlock(workspaceID) + `
		}
		resource "databricks_cluster" "this" {
			cluster_name            = "` + clusterName + `"
			spark_version           = data.databricks_spark_version.latest.id
			node_type_id            = data.databricks_node_type.smallest.id
			num_workers             = 0
			autotermination_minutes = 10
			spark_conf = {
				"spark.databricks.cluster.profile" = "singleNode"
				"spark.master"                     = "local[*]"
			}
			custom_tags = {
				"ResourceClass" = "SingleNode"
			}
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_cluster.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			res, err := w.Clusters.GetByClusterId(ctx, id)
			if err != nil {
				return err
			}
			if res.ClusterName != clusterName {
				return fmt.Errorf("expected cluster name %q, got %q (cluster may be in the wrong workspace)", clusterName, res.ClusterName)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateCluster(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createClusterWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

// ==========================================
// databricks_instance_pool (pools/resource_instance_pool.go)
//
// No prior real (apply) acceptance test existed for this resource, so both a
// normal-workspace and a unified-host variant are added.
// ==========================================

func createInstancePoolWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	poolName := "tf-" + RandomName() + "-pool"
	step := Step{
		Template: `
		data "databricks_node_type" "smallest" {
			local_disk = true
			` + pcBlock(workspaceID) + `
		}
		resource "databricks_instance_pool" "this" {
			instance_pool_name                    = "` + poolName + `"
			node_type_id                          = data.databricks_node_type.smallest.id
			min_idle_instances                    = 0
			idle_instance_autotermination_minutes = 10
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_instance_pool.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			res, err := w.InstancePools.GetByInstancePoolId(ctx, id)
			if err != nil {
				return err
			}
			if res.InstancePoolName != poolName {
				return fmt.Errorf("expected instance pool name %q, got %q (pool may be in the wrong workspace)", poolName, res.InstancePoolName)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateInstancePool(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createInstancePoolWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestAccInstancePool_WorkspaceLevel(t *testing.T) {
	LoadWorkspaceEnv(t)
	createInstancePoolWithProviderConfig(t, currentWorkspaceID(t), nil)
}

// ==========================================
// databricks_instance_profile (aws/resource_instance_profile.go)
//
// AWS-only. Needs a real, registerable IAM instance-profile ARN
// (TEST_EC2_INSTANCE_PROFILE). The terraform ID is the ARN itself.
// ==========================================

func createInstanceProfileWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	if !IsAws(t) {
		Skipf(t)("databricks_instance_profile is AWS-only")
	}
	arn := GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	step := Step{
		Template: `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "` + arn + `"
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_instance_profile.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			profiles, err := w.InstanceProfiles.ListAll(ctx)
			if err != nil {
				return err
			}
			for _, p := range profiles {
				if p.InstanceProfileArn == id {
					return nil
				}
			}
			return fmt.Errorf("instance profile %q not found in workspace %s", id, workspaceID)
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateInstanceProfile(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createInstanceProfileWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}
