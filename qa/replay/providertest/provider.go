// Package providertest selects provider implementations for native Go replay.
package providertest

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
)

// Factories explicitly selects one resource's implementation without changing
// process-wide rollout settings. An empty resource keeps the production mux.
// Both implementations receive the runner's isolated HTTP/auth configuration.
func Factories(customize func(*config.Config) error, resource string, framework bool) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			sdkOptions := []sdkv2.SdkV2ProviderOption{sdkv2.WithConfigCustomizer(customize)}
			frameworkOptions := []pluginfw.PluginFrameworkOption{pluginfw.WithConfigCustomizer(customize)}
			if resource != "" {
				if framework {
					sdkOptions = append(sdkOptions, sdkv2.WithPluginFrameworkResources([]string{resource}))
					frameworkOptions = append(frameworkOptions, pluginfw.WithPluginFrameworkResources([]string{resource}))
				} else {
					sdkOptions = append(sdkOptions, sdkv2.WithSdkV2ResourceFallbacks([]string{resource}))
					frameworkOptions = append(frameworkOptions, pluginfw.WithSdkV2ResourceFallbacks([]string{resource}))
				}
			}
			return providers.GetProviderServer(context.Background(),
				providers.WithSdkV2Provider(sdkv2.DatabricksProvider(sdkOptions...)),
				providers.WithPluginFrameworkProvider(pluginfw.GetDatabricksProviderPluginFramework(frameworkOptions...)))
		},
	}
}
