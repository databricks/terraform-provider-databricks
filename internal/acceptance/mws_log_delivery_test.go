package acceptance

import (
	"testing"
)

func TestMwsAccLogDelivery(t *testing.T) {
	accountLevel(t, LegacyStep{
		Template: `resource "databricks_mws_credentials" "ld" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "tf-acceptance-logdelivery-{var.RANDOM}"
			role_arn         = "{env.TEST_LOGDELIVERY_ARN}"
		}
	
		resource "databricks_mws_storage_configurations" "ld" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "tf-acceptance-logdelivery-{var.RANDOM}"
			bucket_name                = "{env.TEST_LOGDELIVERY_BUCKET}"
		}
	
		resource "databricks_mws_log_delivery" "usage_logs" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_id = databricks_mws_credentials.ld.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.ld.storage_configuration_id
			delivery_path_prefix = "tf-{var.RANDOM}/billable-usage"
			config_name = "Usage {var.RANDOM}"
			log_type = "BILLABLE_USAGE"
			output_format = "CSV"
		}
		
		resource "databricks_mws_log_delivery" "audit_logs" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_id = databricks_mws_credentials.ld.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.ld.storage_configuration_id
			delivery_path_prefix = "tf-{var.RANDOM}/audit-logs"
			config_name = "Audit {var.RANDOM}"
			log_type = "AUDIT_LOGS"
			output_format = "JSON"
		}`,
	})
}
