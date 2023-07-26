package acceptance

import (
	"os"
	"testing"
)

func TestUcAccMetastore(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "ucacct":
		unityAccountLevel(t, step{
			Template: `resource "databricks_metastore" "this" {
				name = "{var.RANDOM}"
				storage_root = "s3://{var.RANDOM}/metastore"
				region = "us-east-1"
				force_destroy = true
			}`,
		})
	case "azure-ucacct":
		unityAccountLevel(t, step{
			Template: `resource "databricks_metastore" "this" {
				name = "{var.RANDOM}"
				storage_root = "abfss://{var.RANDOM}@{var.RANDOM}.dfs.core.windows.net/"
				region = "eastus"
				force_destroy = true
			}`,
		})
	case "gcp-accounts":
		unityAccountLevel(t, step{
			Template: `resource "databricks_metastore" "this" {
				name = "{var.RANDOM}"
				storage_root = "gs://{var.RANDOM}/metastore"
				region = "us-east1"
				force_destroy = true
			}`,
		})
	default:
		t.Skipf("not available on %s", cloudEnv)
	}
}
