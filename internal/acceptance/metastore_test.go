package acceptance

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func getStorageRoot() string {
	if isAws() {
		return "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
	} else if isAzure() {
		return "abfss://{var.RANDOM}@{var.RANDOM}/"
	} else if isGcp() {
		return "gs://{var.RANDOM}/metastore"
	}
	return ""
}

func getRegion() string {
	if isAws() {
		return "us-east-1"
	} else if isAzure() {
		return "eastus"
	} else if isGcp() {
		return "us-east1"
	}
	return ""
}

func TestUcAccRootlessMetastore(t *testing.T) {
	loadUcacctEnv(t)
	runMetastoreTest(t, map[string]any{
		"region": getRegion(),
	})
}

func TestUcAccMetastore(t *testing.T) {
	loadUcacctEnv(t)
	runMetastoreTest(t, map[string]any{
		"storage_root": getStorageRoot(),
		"region":       getRegion(),
	})
}

func TestUcAccMetastoreDeltaSharing(t *testing.T) {
	loadUcacctEnv(t)
	runMetastoreTest(t, map[string]any{
		"storage_root":        getStorageRoot(),
		"region":              getRegion(),
		"delta_sharing_scope": "INTERNAL",
		"delta_sharing_recipient_token_lifetime_in_seconds": 3600,
		"delta_sharing_organization_name":                   "databricks-tf-provider-test",
	})
}

func TestUcAccMetastoreDeltaSharingInfiniteLifetime(t *testing.T) {
	loadUcacctEnv(t)
	runMetastoreTest(t, map[string]any{
		"storage_root":        getStorageRoot(),
		"region":              getRegion(),
		"delta_sharing_scope": "INTERNAL",
		"delta_sharing_recipient_token_lifetime_in_seconds": 0,
	})
}

func TestUcAccMetastoreWithOwnerUpdates(t *testing.T) {
	loadUcacctEnv(t)
	runMetastoreTestWithOwnerUpdates(t, map[string]any{
		"storage_root": getStorageRoot(),
		"region":       getRegion(),
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
			owner = "account users"
			%s
		}`, template),
	}, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.STICKY_RANDOM}"
			force_destroy = true
			owner = "{env.TEST_DATA_ENG_GROUP}"
			%s
		}`, template),
	}, step{
		Template: fmt.Sprintf(`resource "databricks_metastore" "this" {
			name = "{var.STICKY_RANDOM}-updated"
			force_destroy = true
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
			%s
		}`, template),
	})
}
