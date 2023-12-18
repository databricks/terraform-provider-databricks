package acceptance

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func getStorageRoot(cloudEnv string) string {
	switch cloudEnv {
	case "ucacct":
		return "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
	case "azure-ucacct":
		return "abfss://{var.RANDOM}@{var.RANDOM}/"
	case "gcp-accounts":
		return "gs://{var.RANDOM}/metastore"
	default:
		return ""
	}
}

func getRegion(cloudEnv string) string {
	switch cloudEnv {
	case "ucacct":
		return "us-east-1"
	case "azure-ucacct":
		return "eastus"
	case "gcp-accounts":
		return "us-east1"
	default:
		return ""
	}
}

func TestUcAccRootlessMetastore(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	runMetastoreTest(t, map[string]any{
		"region": getRegion(os.Getenv("CLOUD_ENV")),
	})
}

func TestUcAccMetastore(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	runMetastoreTest(t, map[string]any{
		"storage_root": getStorageRoot(os.Getenv("CLOUD_ENV")),
		"region":       getRegion(os.Getenv("CLOUD_ENV")),
	})
}

func TestUcAccMetastoreDeltaSharing(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	runMetastoreTest(t, map[string]any{
		"storage_root":        getStorageRoot(os.Getenv("CLOUD_ENV")),
		"region":              getRegion(os.Getenv("CLOUD_ENV")),
		"delta_sharing_scope": "INTERNAL",
		"delta_sharing_recipient_token_lifetime_in_seconds": 3600,
		"delta_sharing_organization_name":                   "databricks-tf-provider-test",
	})
}

func TestUcAccMetastoreDeltaSharingInfiniteLifetime(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	runMetastoreTest(t, map[string]any{
		"storage_root":        getStorageRoot(os.Getenv("CLOUD_ENV")),
		"region":              getRegion(os.Getenv("CLOUD_ENV")),
		"delta_sharing_scope": "INTERNAL",
		"delta_sharing_recipient_token_lifetime_in_seconds": 0,
	})
}

func TestUcAccMetastoreWithOwnerUpdates(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	runMetastoreTestWithOwnerUpdates(t, map[string]any{
		"storage_root": getStorageRoot(os.Getenv("CLOUD_ENV")),
		"region":       getRegion(os.Getenv("CLOUD_ENV")),
	})
}

func getTemplateFromExtraAttributes(t *testing.T, extraAttributes map[string]any) string {
	params := make([]string, len(extraAttributes))
	for k, v := range extraAttributes {
		jsonValue, err := json.Marshal(v)
		require.NoError(t, err)
		params = append(params, k+" = "+string(jsonValue))
	}
	return strings.Join(params, "\n\t\t\t")
}

func runMetastoreTest(t *testing.T, extraAttributes map[string]any) {
	template := getTemplateFromExtraAttributes(t, extraAttributes)
	unityAccountLevel(t, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.RANDOM}"
			force_destroy = true
			%s
		}`, template),
	})
}

func runMetastoreTestWithOwnerUpdates(t *testing.T, extraAttributes map[string]any) {
	template := getTemplateFromExtraAttributes(t, extraAttributes)
	unityAccountLevel(t, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.STICKY_RANDOM}"
			force_destroy = true
			%s
		}`, template),
	}, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.STICKY_RANDOM}"
			force_destroy = true
			owner = "account users"
			%s
		}`, template),
	}, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.STICKY_RANDOM}"
			force_destroy = true
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
			%s
		}`, template),
	})
}
