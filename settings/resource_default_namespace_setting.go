package settings

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewDefaultNamespaceSettingsAPI creates a DefaultNamespaceSettingsAPI instance
func NewDefaultNamespaceSettingsAPI(ctx context.Context, m any) DefaultNamespaceSettingsAPI {
	client := m.(*common.DatabricksClient)
	return DefaultNamespaceSettingsAPI{client, ctx}
}

// DefaultNamespaceSettingsAPI exposes the Default Namespace Settings API
type DefaultNamespaceSettingsAPI struct {
	client  *common.DatabricksClient
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

type genericSettingDefinition[T, U any] interface {
	SettingStruct() T
	Read(ctx context.Context, c U, id string) (*T, error)
	Update(ctx context.Context, c U, t T) (string, error)
	Delete(ctx context.Context, c U, id string) (string, error)
	GetETag(t *T) string
}

type workspaceSettingDefinition[T any] genericSettingDefinition[T, *databricks.WorkspaceClient]
type accountSettingDefinition[T any] genericSettingDefinition[T, *databricks.AccountClient]

// Candidates for code generation: begin
const (
	defaultNamespaceSettingName string = "databricks_default_namespace_setting"
)

// Default Namespace Setting
type defaultNamespaceSetting struct{}

func (defaultNamespaceSetting) SettingStruct() settings.DefaultNamespaceSetting {
	return settings.DefaultNamespaceSetting{}
}
func (defaultNamespaceSetting) Read(ctx context.Context, w *databricks.WorkspaceClient, id string) (*settings.DefaultNamespaceSetting, error) {
	return w.Settings.ReadDefaultWorkspaceNamespace(ctx, settings.ReadDefaultWorkspaceNamespaceRequest{
		Etag: id,
	})
}
func (defaultNamespaceSetting) Update(ctx context.Context, w *databricks.WorkspaceClient, t settings.DefaultNamespaceSetting) (string, error) {
	t.SettingName = "default"
	res, err := w.Settings.UpdateDefaultWorkspaceNamespace(ctx, settings.UpdateDefaultWorkspaceNamespaceRequest{
		AllowMissing: true,
		Setting:      &t,
		FieldMask:    "namespace.value",
	})
	return res.Etag, err
}
func (defaultNamespaceSetting) Delete(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error) {
	res, err := w.Settings.DeleteDefaultWorkspaceNamespace(ctx, settings.DeleteDefaultWorkspaceNamespaceRequest{
		Etag: id,
	})
	return res.Etag, err
}
func (defaultNamespaceSetting) GetETag(t *settings.DefaultNamespaceSetting) string {
	return t.Etag
}

func AllSettingsResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		defaultNamespaceSettingName: makeSettingResource[settings.DefaultNamespaceSetting, *databricks.WorkspaceClient](defaultNamespaceSetting{}),
	}
}

// Candidates for code generation: end

func makeSettingResource[T, U any](defn genericSettingDefinition[T, U]) *schema.Resource {
	resourceSchema := common.StructToSchema(defn.SettingStruct,
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["etag"].Computed = true
			s["setting_name"].Computed = true
			return s
		})

	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		var setting T
		common.DataToStructPointer(d, resourceSchema, &setting)
		var res string
		switch defn := defn.(type) {
		case workspaceSettingDefinition[T]:
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			res, err = defn.Update(ctx, w, setting)
			if err != nil {
				return err
			}
		case accountSettingDefinition[T]:
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			res, err = defn.Update(ctx, a, setting)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected setting type: %T", defn)
		}
		d.SetId(res)
		return nil
	}

	return common.Resource{
		Schema: resourceSchema,
		Create: createOrUpdate,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var res *T
			switch defn := defn.(type) {
			case workspaceSettingDefinition[T]:
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				res, err = defn.Read(ctx, w, d.Id())
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				res, err = defn.Read(ctx, a, d.Id())
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected setting type: %T", defn)
			}
			err := common.StructToData(res, resourceSchema, d)
			if err != nil {
				return err
			}
			// Update the etag. The server will accept any etag and respond
			// with a response which is at least as recent as the etag.
			// Updating, while not always necessary, ensures that the
			// server responds with an updated response.
			d.SetId(defn.GetETag(res))
			return nil
		},
		Update: createOrUpdate,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var etag string
			switch defn := defn.(type) {
			case workspaceSettingDefinition[T]:
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				etag, err = defn.Delete(ctx, w, d.Id())
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				etag, err = defn.Delete(ctx, a, d.Id())
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected setting type: %T", defn)
			}
			d.SetId(etag)
			return nil
		},
	}.ToResource()
}
