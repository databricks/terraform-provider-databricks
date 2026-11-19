package catalog

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

// forEachVolumeClient runs test once with the volume calls on the existing client and
// once on sdk-go uc/volumes/v1, as USE_SDK_GO_RESOURCES selects it.
func forEachVolumeClient(t *testing.T, test func(t *testing.T, sdkGo bool)) {
	t.Run("existing_path", func(t *testing.T) { test(t, false) })
	t.Run("sdk_go", func(t *testing.T) { test(t, true) })
}

// volumeResource is the provider's wrapped volume resource, with its operations on
// sdk-go if sdkGo is set.
func volumeResource(sdkGo bool) *schema.Resource {
	r := ResourceVolume().ToResource()
	if sdkGo {
		common.EnableSdkGo(r)
	}
	return r
}
