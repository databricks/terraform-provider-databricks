package catalog

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/sdk-go/options/call"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

// readVolumeWithSdkGo reads the volume through sdk-go uc/volumes/v1 into the same
// attributes as the existing path.
func readVolumeWithSdkGo(ctx context.Context, w *databricks.WorkspaceClient,
	s map[string]*schema.Schema, d *schema.ResourceData) error {
	vc, err := common.SdkGoVolumesClient(ctx, w)
	if err != nil {
		return err
	}
	fullName := d.Id()
	get := func(ctx context.Context, opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		req := sdkvolumes.GetVolumeRequest{FullNameArg: &fullName}
		return vc.GetVolume(ctx, req, opts...)
	}
	v, err := common.CallSdkGoVolumes(ctx, vc, get)
	if err != nil {
		return err
	}
	info := volumeInfoFromSdkGo(v)
	if err := common.StructToData(&info, s, d); err != nil {
		return err
	}
	return d.Set("volume_path", "/Volumes/"+strings.ReplaceAll(deref(v.FullName), ".", "/"))
}

// volumeInfoFromSdkGo keeps the fields that the existing path reads into state.
// StructToData then applies the same rules to them, such as keeping the prior value of
// an optional field the response leaves empty.
func volumeInfoFromSdkGo(v *sdkvolumes.VolumeInfo) VolumeInfo {
	return VolumeInfo{
		Name:            deref(v.Name),
		CatalogName:     deref(v.CatalogName),
		SchemaName:      deref(v.SchemaName),
		VolumeType:      catalog.VolumeType(v.VolumeType),
		StorageLocation: deref(v.StorageLocation),
		Owner:           deref(v.Owner),
		Comment:         deref(v.Comment),
	}
}

func deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}
