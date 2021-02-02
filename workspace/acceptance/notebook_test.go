package acceptance

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_notebook" "this" {
				source = "{var.CWD}/testdata/tf-test-python.py"
				path = "/Shared/provider-test/xx_{var.RANDOM}"
			}`,
		},
		{
			Template: `resource "databricks_notebook" "this" {
				source = "{var.CWD}/testdata/tf-test-python.py"
				path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
			}`,
		},
	})
}
