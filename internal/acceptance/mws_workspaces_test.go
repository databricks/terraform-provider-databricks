package acceptance

import (
	"testing"
)

func TestMwsAccWorkspaces(t *testing.T) {
	accountLevel(t, step{
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
	})
}

func TestMwsAccGcpWorkspaces(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
		}`,
	})
}

func TestMwsAccGcpByovpcWorkspaces(t *testing.T) {
	t.Skip()
	// FIXME: flaky with `Secondary IP range (pods, svc) is already in use by another GKE cluster`
	accountLevel(t, step{
		Template: `
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
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}

			network_id = databricks_mws_networks.this.network_id
			
			gke_config {
				connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
				master_ip_range = "10.3.0.0/28"
			}
		}`,
	})
}

func TestMwsAccGcpPscWorkspaces(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "{env.TEST_PREFIX}-network-{var.RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.VPC_NETWORK_ID}"
				subnet_id = "{env.SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
				pod_ip_range_name = "{env.POD_IP_RANGE_NAME}"
				service_ip_range_name = "{env.SVC_IP_RANGE_NAME}"
			}
		}

		resource "databricks_mws_private_access_settings" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			private_access_settings_name = "tf-pas-{var.RANDOM}"
			region = "{env.GOOGLE_REGION}"
			public_access_enabled = true
			private_access_level = "ACCOUNT"
		}
		
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
            
            private_access_settings_id = databricks_mws_private_access_settings.this.private_access_settings_id
			network_id = databricks_mws_networks.this.network_id
			
			gke_config {
				connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
				master_ip_range = "10.3.0.0/28"
			}
		}`,
	})
}
