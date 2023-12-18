package acceptance

import (
	"testing"
)

func TestUcAccCatalog(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}`,
	})
}

func TestUcAccCatalogIsolated(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "ISOLATED"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "OPEN"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	})
}

func TestUcAccCatalogUpdate(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			owner = "account users"
		}`,
	}, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform - updated comment"
			properties     = {
				purpose = "testing"
			}
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
		}`,
	})
}
