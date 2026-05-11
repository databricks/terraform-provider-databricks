package acceptance

import "testing"

// TestMwsAccWorkspaceID_UnknownFromMwsWorkspaces creates a workspace via
// databricks_mws_workspaces and, in the SAME plan, creates a databricks_job
// inside it via provider_config.workspace_id = databricks_mws_workspaces.this.workspace_id.
// The workspace_id reference is unknown at plan time and only becomes known
// once the mws_workspaces resource is applied.
//
// Without the NewValueKnown short-circuit in NamespaceValidateWorkspaceID,
// plan-time validation falls back to c.Config.WorkspaceID — which for an
// account-level provider is empty (or the CLI's "none" sentinel) — and fails
// the entire plan before any resource is created. The expectation here is
// that plan defers validation, apply creates the workspace first, the job
// then resolves the now-known workspace_id and is created inside the new
// workspace, and destroy cleanly tears both down.
//
// The workspace template mirrors TestMwsAccWorkspaces (mws_workspaces_test.go)
// so the same deco-fixture env vars cover both tests.
func TestMwsAccWorkspaceID_UnknownFromMwsWorkspaces(t *testing.T) {
	AccountLevel(t, Step{
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

			custom_tags = {
				"randomkey" = "randomvalue"
			}

			token {
				comment = "Test {var.RANDOM}"
			}
		}
		resource "databricks_job" "j" {
			name = "tf-unknown-wsid-{var.RANDOM}"

			task {
				task_key = "noop"
				new_cluster {
					num_workers   = 1
					spark_version = "15.4.x-scala2.12"
					node_type_id  = "m5.large"
				}
				notebook_task {
					# Path doesn't need to exist for job creation; the task
					# would only fail if the job were run, which the test
					# doesn't do.
					notebook_path = "/Workspace/non-existent-notebook"
				}
			}

			provider_config {
				workspace_id = databricks_mws_workspaces.this.workspace_id
			}
		}`,
	})
}
