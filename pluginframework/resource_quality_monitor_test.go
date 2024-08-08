package pluginframework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func getProtoV6ProviderFactory() map[string]func() (tfprotov6.ProviderServer, error) {
	return nil
}

func TestQualityMonitorCreate(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: getProtoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				ResourceName: "databricks_quality_monitor_pluginframework",
			},
		},
	})
}
