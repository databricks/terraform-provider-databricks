package settings

import (
	"context"
	"fmt"
	"regexp"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewSettingsAPI creates SettingsAPI instance from provider meta
func NewSettingsAPI(ctx context.Context, m any) SettingsAPI {
	client := m.(*common.DatabricksClient)
	return SettingsAPI{client, ctx}
}

// SettingsAPI exposes the Settings API
type SettingsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a SettingsAPI) getCurrentEtag(etag string) (string, error) {
	res, err := a.Read(etag)
	if err != nil {
		// If settings is not set-up, we will get an error with the etag in it.
		// This will happen, for instance, during the first ever call to GET.
		return a.getEtagFromError(err)
	}
	return res.Etag, nil
}

func (a SettingsAPI) getEtagFromError(err error) (string, error) {
	errorInfos := apierr.GetErrorInfo(err)
	if len(errorInfos) > 0 {
		metadata := errorInfos[0].Metadata
		if etag, ok := metadata["etag"]; ok {
			return etag, nil
		}
	}
	return "", fmt.Errorf("error fetching the default workspace namespace settings: %w", err)
}

func (a SettingsAPI) Create(request settings.UpdateDefaultWorkspaceNamespaceRequest) (string, error) {
	return a.Update("", request)
}

func (a SettingsAPI) Read(etag string) (settings.DefaultNamespaceSetting, error) {
	var setting settings.DefaultNamespaceSetting
	err := a.client.Get(a.context, "/settings/types/default_namespace_ws/names/default", map[string]string{
		"etag": etag,
	}, &setting)
	return setting, err
}

func (a SettingsAPI) Delete(etag string) (string, error) {
	etag, err := a.getCurrentEtag(etag)
	if err != nil {
		return "", err
	}
	var response settings.DeleteDefaultWorkspaceNamespaceResponse
	err = a.client.DeleteWithResponse(a.context, "/settings/types/default_namespace_ws/names/default", map[string]string{
		"etag": etag,
	}, &response)
	if err != nil {
		return "", err
	}
	return response.Etag, nil
}

func (a SettingsAPI) Update(etag string, request settings.UpdateDefaultWorkspaceNamespaceRequest) (string, error) {
	etag, err := a.getCurrentEtag(etag)
	if err != nil {
		return "", err
	}
	request.Setting.Etag = etag
	var response settings.DefaultNamespaceSetting
	err = a.client.PatchWithResponse(a.context, "/settings/types/default_namespace_ws/names/default", request, &response)
	if err != nil {
		return "", err
	}
	return response.Etag, nil
}

var resourceSchema = common.StructToSchema(settings.DefaultNamespaceSetting{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		delete(s, "etag")
		regex, _ := regexp.Compile("default")
		s["setting_name"].ValidateFunc = validation.StringMatch(regex, `only "default" setting name is currently supported`)

		return s
	})

func ResourceNamespaceSettings() *schema.Resource {
	return common.Resource{
		Schema: resourceSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var setting settings.DefaultNamespaceSetting
			common.DataToStructPointer(d, resourceSchema, &setting)
			request := settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				Setting:      &setting,
				FieldMask:    "namespace.value",
			}
			settingAPI := NewSettingsAPI(ctx, c)
			etag, err := settingAPI.Create(request)
			if err != nil {
				return err
			}
			d.SetId(etag)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			settingAPI := NewSettingsAPI(ctx, c)
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
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var setting settings.DefaultNamespaceSetting
			common.DataToStructPointer(d, resourceSchema, &setting)
			request := settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				Setting:      &setting,
				FieldMask:    "namespace.value",
			}
			settingAPI := NewSettingsAPI(ctx, c)
			etag, err := settingAPI.Update(d.Id(), request)
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
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			etag, err := NewSettingsAPI(ctx, c).Delete(d.Id())
			if err != nil {
				return err
			}
			d.SetId(etag)
			return nil
		},
	}.ToResource()
}
