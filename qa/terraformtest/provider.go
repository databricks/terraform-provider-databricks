package terraformtest

import (
	"context"
	"fmt"
	"slices"

	"github.com/databricks/databricks-sdk-go/config"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
)

type providerImplementation string

const (
	sdkv2Implementation           providerImplementation = "sdkv2"
	pluginFrameworkImplementation providerImplementation = "plugin_framework"
)

// Overrides name resource types, not Terraform instance addresses. Unlisted
// resources and all data sources retain their production registration.
type providerImplementations map[string]providerImplementation

func (p providerImplementations) validate() error {
	for name, implementation := range p {
		if name == "" {
			return fmt.Errorf("provider_implementations requires a resource type name")
		}
		switch implementation {
		case sdkv2Implementation, pluginFrameworkImplementation:
		default:
			return fmt.Errorf("provider_implementations[%q]: expected sdkv2 or plugin_framework, got %q", name, implementation)
		}
	}
	return nil
}

// ProviderFactories builds the production mux with a caller-supplied client
// customizer. Native Go replays can wrap RunNative's isolated configuration to
// customize transports or authentication without duplicating provider setup.
func ProviderFactories(customize func(*config.Config) error) map[string]func() (tfprotov6.ProviderServer, error) {
	return providerFactories(customize, nil)
}

func providerFactories(customize func(*config.Config) error, implementations providerImplementations) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			var sdkResources, frameworkResources []string
			for name, implementation := range implementations {
				switch implementation {
				case sdkv2Implementation:
					sdkResources = append(sdkResources, name)
				case pluginFrameworkImplementation:
					frameworkResources = append(frameworkResources, name)
				}
			}
			slices.Sort(sdkResources)
			slices.Sort(frameworkResources)
			sdkProvider := sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(customize), sdkv2.WithSdkV2ResourceFallbacks(sdkResources), sdkv2.WithPluginFrameworkResources(frameworkResources))
			frameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(pluginfw.WithConfigCustomizer(customize), pluginfw.WithSdkV2ResourceFallbacks(sdkResources), pluginfw.WithPluginFrameworkResources(frameworkResources))
			if err := implementations.validateResources(sdkProvider.ResourcesMap, frameworkProvider.Resources(context.Background())); err != nil {
				return nil, err
			}
			return providers.GetProviderServer(context.Background(), providers.WithSdkV2Provider(sdkProvider), providers.WithPluginFrameworkProvider(frameworkProvider))
		},
	}
}

// Reject unavailable selections rather than silently exercising the other
// implementation. Metadata inspection requires no credentials or API calls.
func (p providerImplementations) validateResources(sdkResources map[string]*schema.Resource, frameworkResources []func() frameworkresource.Resource) error {
	if len(p) == 0 {
		return nil
	}
	frameworkNames := map[string]bool{}
	for _, factory := range frameworkResources {
		var metadata frameworkresource.MetadataResponse
		factory().Metadata(context.Background(), frameworkresource.MetadataRequest{ProviderTypeName: "databricks"}, &metadata)
		frameworkNames[metadata.TypeName] = true
	}
	for name, implementation := range p {
		if implementation == sdkv2Implementation && sdkResources[name] != nil || implementation == pluginFrameworkImplementation && frameworkNames[name] {
			continue
		}
		return fmt.Errorf("provider_implementations[%q]: %s implementation is not available", name, implementation)
	}
	return nil
}

// cleanupConfigSteps tracks which step supplies the retained configuration
// after each operation. HCL and provider selection must follow the same state:
// failed updates, later previews and discarded imports do not own that state.
func (d *testDefinition) cleanupConfigSteps() []int {
	retained := make([]int, len(d.Steps))
	current, hasApply := -1, false
	for i, s := range d.Steps {
		if s.ExpectError == nil && (s.Import == nil || *s.Import.Persist) && (!s.PlanOnly || !hasApply) {
			current, hasApply = i, !s.PlanOnly
		} else if s.Import != nil && !hasApply && d.InitialState != nil {
			current = i
		}
		retained[i] = current
	}
	return retained
}

// Install step-level factories only for journeys which select implementations;
// checkpoint/file staging preserves them because the case-level factory is nil.
func configureProviderImplementations(d *testDefinition, tc *resource.TestCase, customize func(*config.Config) error) {
	if !slices.ContainsFunc(d.Steps, func(s step) bool { return len(s.ProviderImplementations) > 0 }) {
		return
	}
	useStepProviderFactories(tc)
	for i, s := range d.Steps {
		tc.Steps[i].ProtoV6ProviderFactories = providerFactories(customize, s.ProviderImplementations)
	}
	var cleanup providerImplementations
	retained := d.cleanupConfigSteps()
	if index := retained[len(retained)-1]; index >= 0 {
		cleanup = d.Steps[index].ProviderImplementations
	}
	tc.Steps[len(d.Steps)].ProtoV6ProviderFactories = providerFactories(customize, cleanup)
}
