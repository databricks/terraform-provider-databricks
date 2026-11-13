package terraformtest

import (
	"fmt"
	"slices"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
)

type sdkImplementation string

const (
	classicSdkImplementation sdkImplementation = "classic"
	sdkGoImplementation      sdkImplementation = "sdk_go"
)

// sdkImplementations selects, by resource type, whether a resource uses the provider's
// existing clients or the modular sdk-go clients. Unlisted resources use the existing
// clients: replay ignores USE_SDK_GO_RESOURCES.
type sdkImplementations map[string]sdkImplementation

func (s sdkImplementations) validate() error {
	for name, implementation := range s {
		if name == "" {
			return fmt.Errorf("sdk_implementations requires a resource type name")
		}
		switch implementation {
		case classicSdkImplementation, sdkGoImplementation:
		default:
			return fmt.Errorf("sdk_implementations[%q]: expected classic or sdk_go, got %q",
				name, implementation)
		}
		if !slices.Contains(sdkv2.SdkGoResourceNames(), name) {
			return fmt.Errorf("sdk_implementations[%q]: resource has no sdk_go implementation",
				name)
		}
	}
	return nil
}

func (s sdkImplementations) sdkGoResources() []string {
	names := []string{}
	for name, implementation := range s {
		if implementation == sdkGoImplementation {
			names = append(names, name)
		}
	}
	slices.Sort(names)
	return names
}

// ProviderFactoriesWithSdkGo is ProviderFactories with the given resource types using
// sdk-go, as sdk_implementations selects them in YAML. Native Go replays use it.
func ProviderFactoriesWithSdkGo(customize func(*config.Config) error,
	resources []string) map[string]func() (tfprotov6.ProviderServer, error) {
	selection := sdkImplementations{}
	for _, name := range resources {
		selection[name] = sdkGoImplementation
	}
	return providerFactories(customize, nil, selection)
}
