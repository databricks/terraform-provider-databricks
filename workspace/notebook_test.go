package workspace_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccNotebookResourceDbcUpdate(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update2.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	})
}

func TestAccNotebookResourceJupiterUpdate(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update2.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	})
}
