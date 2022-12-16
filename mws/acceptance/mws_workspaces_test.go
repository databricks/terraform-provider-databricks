package acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccWorkspaces(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=MWS is set")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_mws_credentials" "this" {
				account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
				credentials_name = "credentials-ws-{var.RANDOM}"
				role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
			}
			resource "databricks_mws_customer_managed_keys" "this" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				aws_key_info {
					key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
					key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
				}
                use_cases = ["MANAGED_SERVICES"]
			}
			resource "databricks_mws_storage_configurations" "this" {
				account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
				storage_configuration_name = "storage-ws-{var.RANDOM}"
				bucket_name                = "{env.TEST_ROOT_BUCKET}"
			}
			resource "databricks_mws_networks" "this" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				network_name = "network-ws-{var.RANDOM}"
				vpc_id       = "{env.TEST_VPC_ID}"
				subnet_ids   = [
					"{env.TEST_SUBNET_PRIVATE}",
					"{env.TEST_SUBNET_PRIVATE2}",
				]
				security_group_ids = [
					"{env.TEST_SECURITY_GROUP}",
				]
			}
			resource "databricks_mws_workspaces" "this" {
				account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name  = "terra-{var.RANDOM}"
				aws_region      = "{env.AWS_REGION}"
		
				network_id = databricks_mws_networks.this.network_id
				credentials_id = databricks_mws_credentials.this.credentials_id
				storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
				managed_services_customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id

				token {
					comment = "Test {var.RANDOM}"
				}
			}`,
		},
	})
}

func TestGcpAccaWorkspaces(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "gcp-accounts" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=gcp-accounts is set")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_mws_workspaces" "this" {
				account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
				location        = "{env.GOOGLE_REGION}"
		
				cloud_resource_bucket {
					gcp {
						project_id = "{env.GOOGLE_PROJECT}"
					}
				}
			}`,
		},
	})
}

func TestGcpAccByovpcWorkspaces(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "gcp-accounts" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=gcp-accounts is set")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "google_compute_router" "this" {
			    name    = "{env.TEST_PREFIX}-router-{var.RANDOM}"
			    region  = "{env.GOOGLE_REGION}"
			    network = "{env.TEST_VPC_ID}"
			}
			
			resource "google_compute_router_nat" "nat" {
			    name                               = "{env.TEST_PREFIX}-nat-{var.RANDOM}"
			    router                             = google_compute_router.this.name
			    region                             = google_compute_router.this.region
			    nat_ip_allocate_option             = "AUTO_ONLY"
			    source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
			}

			resource "databricks_mws_networks" "this" {
				account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
				network_name = "{env.TEST_PREFIX}-network-{var.RANDOM}"
				gcp_network_info {
					network_project_id = "{env.GOOGLE_PROJECT}"
					vpc_id = "{env.TEST_VPC_ID}"
					subnet_id = "{env.TEST_SUBNET_ID}"
					subnet_region = "{env.GOOGLE_REGION}"
					pod_ip_range_name = "pods"
					service_ip_range_name = "svc"
		  		}
			}
			
			resource "databricks_mws_workspaces" "this" {
				account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
				location        = "{env.GOOGLE_REGION}"
		
				cloud_resource_bucket {
					gcp {
						project_id = "{env.GOOGLE_PROJECT}"
					}
				}
				network {
					network_id = databricks_mws_networks.this.network_id
					gcp_common_network_config {
						gke_connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
						gke_cluster_master_ip_range = "10.3.0.0/28"
					}
  				}
			}`,
		},
	})
}
