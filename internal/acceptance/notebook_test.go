package acceptance

import (
	"testing"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, LegacyStep{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}
