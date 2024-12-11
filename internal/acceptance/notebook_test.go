package acceptance

import (
	"testing"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccNotebookResourceDbcUpdate(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/acc-test-update1.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	}, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/acc-test-update2.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	})
}

func TestAccNotebookResourceJupiterUpdate(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/acc-test-update1.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	}, Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/acc-test-update2.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	})
}
