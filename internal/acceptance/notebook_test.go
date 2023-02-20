package acceptance

import (
	"testing"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}
