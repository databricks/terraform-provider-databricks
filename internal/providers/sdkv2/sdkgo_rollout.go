package sdkv2

// Resources listed in sdkGoResources can read through the modular
// github.com/databricks/sdk-go clients instead of the provider's existing clients.
// Users opt in per resource with the USE_SDK_GO_RESOURCES environment variable, for
// example USE_SDK_GO_RESOURCES="databricks_service_principal". It is separate from the
// plugin framework rollout variables.

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

const useSdkGoResourcesEnv = "USE_SDK_GO_RESOURCES"

// Resource types that have an sdk-go implementation. Keep this list sorted.
var sdkGoResources = []string{
	"databricks_service_principal",
}

// The selection is read once, so it stays fixed for the life of the provider process.
var sdkGoResourcesFromEnv = sync.OnceValues(func() ([]string, error) {
	return parseSdkGoResources(os.Getenv(useSdkGoResourcesEnv))
})

// SdkGoResourceNames returns the resource types that USE_SDK_GO_RESOURCES accepts.
func SdkGoResourceNames() []string {
	return slices.Clone(sdkGoResources)
}

// WithSdkGoResources selects the resources that use sdk-go instead of reading
// USE_SDK_GO_RESOURCES. This should only be used for testing.
func WithSdkGoResources(resources []string) SdkV2ProviderOption {
	return func(o *sdkV2ProviderOptions) {
		o.sdkGoResources = &resources
	}
}

// parseSdkGoResources splits a comma-separated list, ignoring spaces and empty entries.
func parseSdkGoResources(value string) ([]string, error) {
	names := strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || unicode.IsSpace(r)
	})
	return names, validateSdkGoResources(names)
}

func validateSdkGoResources(names []string) error {
	for _, name := range names {
		if !slices.Contains(sdkGoResources, name) {
			return fmt.Errorf("%s: resource %q has no sdk-go implementation; supported resources: %s",
				useSdkGoResourcesEnv, name, strings.Join(sdkGoResources, ", "))
		}
	}
	return nil
}

// selectedSdkGoResources returns the test selection if one is set, otherwise the
// environment selection.
func (o *sdkV2ProviderOptions) selectedSdkGoResources() ([]string, error) {
	if o.sdkGoResources != nil {
		return *o.sdkGoResources, validateSdkGoResources(*o.sdkGoResources)
	}
	return sdkGoResourcesFromEnv()
}

// enableSdkGoResources marks the selected resources' operations to use sdk-go.
func enableSdkGoResources(resources map[string]*schema.Resource, names []string) {
	for _, name := range names {
		if r, ok := resources[name]; ok {
			common.EnableSdkGo(r)
		}
	}
}
