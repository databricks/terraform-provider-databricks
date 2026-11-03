package sql_alert_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestSqlAlert(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
