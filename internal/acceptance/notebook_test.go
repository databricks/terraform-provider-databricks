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
