package settings

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewDefaultNamespaceSettingsAPI creates a DefaultNamespaceSettingsAPI instance
func NewDefaultNamespaceSettingsAPI(ctx context.Context, m any) DefaultNamespaceSettingsAPI {
	client := m.(common.DatabricksAPI)
	return DefaultNamespaceSettingsAPI{client, ctx}
}

// DefaultNamespaceSettingsAPI exposes the Default Namespace Settings API
type DefaultNamespaceSettingsAPI struct {
	client  common.DatabricksAPI
	context context.Context
}

func (a DefaultNamespaceSettingsAPI) isEtagVersionError(err error) bool {
	var aerr *apierr.APIError
	if !errors.As(err, &aerr) {
		return false
	}
	return aerr.StatusCode == http.StatusNotFound || (aerr.StatusCode == http.StatusConflict && aerr.ErrorCode == "RESOURCE_CONFLICT")
}

func (a DefaultNamespaceSettingsAPI) getEtagFromError(err error) (string, error) {
	if !a.isEtagVersionError(err) {
		return "", err
	}
	errorInfos := apierr.GetErrorInfo(err)
	if len(errorInfos) > 0 {
		metadata := errorInfos[0].Metadata
		if etag, ok := metadata["etag"]; ok {
			return etag, nil
		}
	}
	return "", fmt.Errorf("error fetching the default workspace namespace settings: %w", err)
}

func (a DefaultNamespaceSettingsAPI) DeleteWithRetry(etag string) (string, error) {
	etag, err := a.executeDelete(etag)
	if err == nil {
		return etag, nil
	}
	etag, err = a.getEtagFromError(err)
	if err != nil {
		return "", err
	}
	return a.executeDelete(etag)
}

func (a DefaultNamespaceSettingsAPI) executeDelete(etag string) (string, error) {
	var response settings.DefaultNamespaceSetting
	err := a.client.DeleteWithResponse(a.context, "/settings/types/default_namespace_ws/names/default", map[string]string{
		"etag": etag,
	}, &response)
	if err != nil {
		return "", err
	}
	return response.Etag, nil
}

func (a DefaultNamespaceSettingsAPI) UpdateWithRetry(request settings.UpdateDefaultWorkspaceNamespaceRequest) (string, error) {
	etag, err := a.executeUpdate(request)
	if err == nil {
		return etag, nil
	}
	etag, err = a.getEtagFromError(err)
	if err != nil {
		return "", err
	}
	request.Setting.Etag = etag
	return a.executeUpdate(request)
}

func (a DefaultNamespaceSettingsAPI) executeUpdate(request settings.UpdateDefaultWorkspaceNamespaceRequest) (string, error) {
	var response settings.DefaultNamespaceSetting
	err := a.client.PatchWithResponse(a.context, "/settings/types/default_namespace_ws/names/default", request, &response)
	if err != nil {
		return "", err
	}
	return response.Etag, nil
}

func (a DefaultNamespaceSettingsAPI) Read(etag string) (settings.DefaultNamespaceSetting, error) {
	var setting settings.DefaultNamespaceSetting
	err := a.client.Get(a.context, "/settings/types/default_namespace_ws/names/default", map[string]string{
		"etag": etag,
	}, &setting)
	return setting, err
}

var resourceSchema = common.StructToSchema(settings.DefaultNamespaceSetting{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["etag"].Computed = true
		s["setting_name"].Computed = true

		return s
	})

func ResourceDefaultNamespaceSetting() *schema.Resource {
	return common.Resource{
		Schema: resourceSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			var setting settings.DefaultNamespaceSetting
			common.DataToStructPointer(d, resourceSchema, &setting)
			setting.SettingName = "default"
			request := settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				Setting:      &setting,
				FieldMask:    "namespace.value",
			}
			settingAPI := NewDefaultNamespaceSettingsAPI(ctx, c)
			etag, err := settingAPI.UpdateWithRetry(request)
			if err != nil {
				return err
			}
			d.SetId(etag)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			settingAPI := NewDefaultNamespaceSettingsAPI(ctx, c)
			res, err := settingAPI.Read(d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(res, resourceSchema, d)
			if err != nil {
				return err
			}
			// Update the etag. The server will accept any etag and respond
			// with a response which is at least as recent as the etag.
			// Updating, while not always necessary, ensures that the
			// server responds with an updated response.
			d.SetId(res.Etag)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			var setting settings.DefaultNamespaceSetting
			common.DataToStructPointer(d, resourceSchema, &setting)
			setting.SettingName = "default"
			setting.Etag = d.Id()
			request := settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				Setting:      &setting,
				FieldMask:    "namespace.value",
			}
			settingAPI := NewDefaultNamespaceSettingsAPI(ctx, c)
			etag, err := settingAPI.UpdateWithRetry(request)
			if err != nil {
				return err
			}
			// Update the etag. The server will accept any etag and respond
			// with a response which is at least as recent as the etag.
			// Updating, while not always necessary, ensures that the
			// server responds with an updated response.
			d.SetId(etag)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			etag, err := NewDefaultNamespaceSettingsAPI(ctx, c).DeleteWithRetry(d.Id())
			if err != nil {
				return err
			}
			d.SetId(etag)
			return nil
		},
	}.ToResource()
}
