package pluginframework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestQualityMonitorCreate(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Steps: []resource.TestStep{},
	})
}
