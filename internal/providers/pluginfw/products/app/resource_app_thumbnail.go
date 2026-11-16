package app

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// The Apps UI rejects larger thumbnails; the API itself has no size limit.
const maxThumbnailBytes = 250000

// readThumbnail reads a thumbnail image and rejects what the Apps API or UI would reject: anything
// but a PNG or JPEG (detected by the same signatures the API checks) and files over 250 KB.
func readThumbnail(thumbnailPath string) ([]byte, error) {
	content, err := common.ReadFileContent(thumbnailPath)
	if err != nil {
		return nil, err
	}
	if len(content) > maxThumbnailBytes {
		return nil, fmt.Errorf("%s is %d bytes, larger than the %d bytes (250 KB) an app thumbnail may have",
			thumbnailPath, len(content), maxThumbnailBytes)
	}
	if contentType := http.DetectContentType(content); contentType != "image/png" && contentType != "image/jpeg" {
		return nil, fmt.Errorf("%s must be a PNG or JPEG image, but its content is %s", thumbnailPath, contentType)
	}
	return content, nil
}

// planThumbnailMd5 plans the hash of the thumbnail file, so that a changed image shows as a diff
// although its path didn't change. The image can't be read back from the API.
func planThumbnailMd5(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var thumbnailPath types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("thumbnail_path"), &thumbnailPath)...)
	if resp.Diagnostics.HasError() || thumbnailPath.IsUnknown() {
		return
	}
	md5 := types.StringNull()
	if !thumbnailPath.IsNull() {
		content, err := readThumbnail(thumbnailPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("thumbnail_path"), "invalid app thumbnail", err.Error())
			return
		}
		md5 = types.StringValue(common.CalculateMd5Hash(content))
	}
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("thumbnail_md5"), md5)...)
}

// syncThumbnail uploads the thumbnail when its hash differs from priorMd5, or deletes the app's
// thumbnail when thumbnail_path was removed. It returns the hash to keep in state.
func syncThumbnail(ctx context.Context, a apps.AppsInterface, name string, thumbnailPath, priorMd5 types.String) (types.String, error) {
	if thumbnailPath.IsNull() {
		if priorMd5.IsNull() {
			return priorMd5, nil
		}
		err := a.DeleteAppThumbnail(ctx, apps.DeleteAppThumbnailRequest{Name: name})
		if err != nil && !apierr.IsMissing(err) {
			return priorMd5, fmt.Errorf("failed to delete app thumbnail: %w", err)
		}
		return types.StringNull(), nil
	}
	content, err := readThumbnail(thumbnailPath.ValueString())
	if err != nil {
		return priorMd5, err
	}
	md5 := common.CalculateMd5Hash(content)
	if md5 == priorMd5.ValueString() {
		return priorMd5, nil
	}
	_, err = a.UpdateAppThumbnail(ctx, apps.UpdateAppThumbnailRequest{
		Name:         name,
		AppThumbnail: &apps.AppThumbnail{Thumbnail: base64.StdEncoding.EncodeToString(content)},
	})
	if err != nil {
		return priorMd5, fmt.Errorf("failed to update app thumbnail: %w", err)
	}
	return types.StringValue(md5), nil
}

// onlyThumbnailChanged reports whether an update changes nothing but the thumbnail, so that a new
// image doesn't run a full app update. An unknown value the configuration leaves unset is a
// computed value the provider fills in, not a change.
func onlyThumbnailChanged(req resource.UpdateRequest) (bool, error) {
	var plan, state, config map[string]tftypes.Value
	if err := req.Plan.Raw.As(&plan); err != nil {
		return false, err
	}
	if err := req.State.Raw.As(&state); err != nil {
		return false, err
	}
	if err := req.Config.Raw.As(&config); err != nil {
		return false, err
	}
	for name, planned := range plan {
		if name == "thumbnail_path" || name == "thumbnail_md5" || planned.Equal(state[name]) {
			continue
		}
		if !planned.IsKnown() && config[name].IsNull() {
			continue
		}
		return false, nil
	}
	return true, nil
}
