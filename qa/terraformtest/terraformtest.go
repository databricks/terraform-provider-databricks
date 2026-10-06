package terraformtest

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// HTTPFixtures runs Terraform against the production muxed provider and a
// local HTTP server so tests cover provider dispatch without live credentials.
func HTTPFixtures(t *testing.T, fixtures []qa.HTTPFixture, steps ...resource.TestStep) {
	t.Helper()
	_, server, err := qa.StrictHTTPFixtureClient(t, fixtures)
	if err != nil {
		t.Fatal(err)
	}

	customizeConfig := func(cfg *config.Config) error {
		cfg.Host = server.URL
		cfg.Token = "..."
		cfg.AuthType = "pat"
		cfg.AzureEnvironment = "PUBLIC"
		return nil
	}
	providerFactories := map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			sdkV2Provider := sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(customizeConfig))
			pluginFrameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(
				pluginfw.WithConfigCustomizer(customizeConfig),
			)
			return providers.GetProviderServer(
				context.Background(),
				providers.WithSdkV2Provider(sdkV2Provider),
				providers.WithPluginFrameworkProvider(pluginFrameworkProvider),
			)
		},
	}
	for i := range steps {
		steps[i].ProtoV6ProviderFactories = providerFactories
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Steps:      steps,
	})
}
