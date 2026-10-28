package notificationdestination_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestNotificationDestination(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
