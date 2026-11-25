package catalog

import (
	"context"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
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
	v, err := vc.ReadByName(ctx, d.Id())
	if err != nil {
		return err
	}
	info := volumeInfoFromSdkGo(v)
	if err := common.StructToData(&info, s, d); err != nil {
		return err
	}
	volumePath := "/Volumes/" + strings.ReplaceAll(deref(v.FullName), ".", "/")
	return d.Set("volume_path", volumePath)
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

// sdkGoVolumeWriter sends the volume writes through sdk-go uc/volumes/v1 with the requests
// the existing client sends.
type sdkGoVolumeWriter struct {
	vc *common.SdkGoVolumes
}

func newSdkGoVolumeWriter(ctx context.Context,
	w *databricks.WorkspaceClient) (sdkGoVolumeWriter, error) {
	vc, err := common.SdkGoVolumesClient(ctx, w)
	if err != nil {
		return sdkGoVolumeWriter{}, err
	}
	return sdkGoVolumeWriter{vc}, nil
}

func (s sdkGoVolumeWriter) create(ctx context.Context,
	r catalog.CreateVolumeRequestContent) (string, error) {
	v, err := s.vc.Create(ctx, createVolumeRequestToSdkGo(r))
	if err != nil {
		return "", err
	}
	return deref(v.FullName), nil
}

func (s sdkGoVolumeWriter) update(ctx context.Context,
	r catalog.UpdateVolumeRequestContent) (string, error) {
	v, err := s.vc.Update(ctx, updateVolumeRequestToSdkGo(r))
	if err != nil {
		return "", err
	}
	return deref(v.FullName), nil
}

func (s sdkGoVolumeWriter) delete(ctx context.Context, fullName string) error {
	return s.vc.DeleteByName(ctx, fullName)
}

// createVolumeRequestToSdkGo sends the fields the existing client sends. Unlike it,
// sdk-go omits an empty volume_type, but the schema only accepts MANAGED and EXTERNAL.
func createVolumeRequestToSdkGo(
	r catalog.CreateVolumeRequestContent) sdkvolumes.CreateVolumeRequest {
	return sdkvolumes.CreateVolumeRequest{
		CatalogName:     &r.CatalogName,
		SchemaName:      &r.SchemaName,
		Name:            &r.Name,
		VolumeType:      sdkvolumes.VolumeType(r.VolumeType),
		StorageLocation: sentString(r.StorageLocation, "StorageLocation", r.ForceSendFields),
		Comment:         sentString(r.Comment, "Comment", r.ForceSendFields),
	}
}

// updateVolumeRequestToSdkGo sends the fields the existing client sends. The full name
// goes to full_name_arg, which sdk-go needs for the path.
func updateVolumeRequestToSdkGo(
	r catalog.UpdateVolumeRequestContent) sdkvolumes.UpdateVolumeRequest {
	return sdkvolumes.UpdateVolumeRequest{
		FullNameArg: &r.Name,
		Comment:     sentString(r.Comment, "Comment", r.ForceSendFields),
		NewName:     sentString(r.NewName, "NewName", r.ForceSendFields),
		Owner:       sentString(r.Owner, "Owner", r.ForceSendFields),
	}
}

// sentString is what the existing client sends for an omitempty string: nothing when it
// is empty, unless field is in forceSendFields.
func sentString(value, field string, forceSendFields []string) *string {
	if value == "" && !slices.Contains(forceSendFields, field) {
		return nil
	}
	return &value
}

func deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}
