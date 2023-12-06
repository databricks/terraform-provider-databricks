package settings

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func retryOnEtagError[Req, Resp any](f func(req Req) (Resp, error), firstReq Req, updateReq func(req *Req, newEtag string)) (Resp, error) {
	req := firstReq
	// Retry once on etag error.
	res, err := f(req)
	if err == nil {
		return res, nil
	}
	etag, err := getEtagFromError(err)
	if err != nil {
		return res, err
	}
	updateReq(&req, etag)
	return f(req)
}

func isEtagVersionError(err error) bool {
	return errors.Is(err, databricks.ErrResourceConflict) || errors.Is(err, databricks.ErrNotFound)
}

func getEtagFromError(err error) (string, error) {
	if !isEtagVersionError(err) {
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

type genericSettingDefinition[T, U any] interface {
	// Returns the struct corresponding to the setting. The schema of the Terraform resource will be generated from this struct.
	SettingStruct() T

	// Read the setting from the server. The etag is provided as the third argument.
	Read(ctx context.Context, c U, etag string) (*T, error)

	// Update the setting to the value specified by t, and return the new etag.
	Update(ctx context.Context, c U, t T) (string, error)

	// Delete the setting with the given etag, and return the new etag.
	Delete(ctx context.Context, c U, etag string) (string, error)

	// Get the etag from the setting.
	GetETag(t *T) string

	// Update the etag in the setting.
	SetETag(t *T, newEtag string)
}

type workspaceSettingDefinition[T any] genericSettingDefinition[T, *databricks.WorkspaceClient]
type accountSettingDefinition[T any] genericSettingDefinition[T, *databricks.AccountClient]

// Instructions for adding a new setting:
//
//  1. Add a new setting name constant. The resulting Terraform resource name will be "databricks_<your settings_name>_setting".
//  2. Create a struct corresponding to your setting, and implement either the workspaceSettingDefinition or accountSettingDefinition interface.
//     Add an assertion to ensure that your type implements said interface. If the setting name is user-settable, it will be provided in the
//     third argument to the Update method. If not, you must set the SettingName field appropriately. You must also set AllowMissing: true
//     and the field mask to the field to update.
//  3. Add a new entry to the AllSettingsResources map below.
const (
	defaultNamespaceSettingName string = "default_namespace"
)

// Default Namespace Setting
type defaultNamespaceSetting struct{}

func (defaultNamespaceSetting) SettingStruct() settings.DefaultNamespaceSetting {
	return settings.DefaultNamespaceSetting{}
}
func (defaultNamespaceSetting) Read(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.DefaultNamespaceSetting, error) {
	return w.Settings.ReadDefaultWorkspaceNamespace(ctx, settings.ReadDefaultWorkspaceNamespaceRequest{
		Etag: etag,
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
func (defaultNamespaceSetting) Delete(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
	res, err := w.Settings.DeleteDefaultWorkspaceNamespace(ctx, settings.DeleteDefaultWorkspaceNamespaceRequest{
		Etag: etag,
	})
	return res.Etag, err
}
func (defaultNamespaceSetting) GetETag(t *settings.DefaultNamespaceSetting) string {
	return t.Etag
}
func (defaultNamespaceSetting) SetETag(t *settings.DefaultNamespaceSetting, newEtag string) {
	t.Etag = newEtag
}

var _ workspaceSettingDefinition[settings.DefaultNamespaceSetting] = defaultNamespaceSetting{}

func AllSettingsResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		defaultNamespaceSettingName: makeSettingResource[settings.DefaultNamespaceSetting, *databricks.WorkspaceClient](defaultNamespaceSetting{}),
	}
}

// Candidates for code generation: end

func makeSettingResource[T, U any](defn genericSettingDefinition[T, U]) *schema.Resource {
	resourceSchema := common.StructToSchema(defn.SettingStruct(),
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["etag"].Computed = true
			// Note: this may not always be computed, but it is for the default namespace setting. If other settings
			// are added for which setting_name is not computed, we'll need to expose this somehow as part of the setting
			// definition.
			s["setting_name"].Computed = true
			return s
		})

	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient, setting T) error {
		common.DataToStructPointer(d, resourceSchema, &setting)
		var res string
		switch defn := defn.(type) {
		case workspaceSettingDefinition[T]:
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			res, err = retryOnEtagError[T, string](func(setting T) (string, error) { return defn.Update(ctx, w, setting) }, setting, defn.SetETag)
			if err != nil {
				return err
			}
		case accountSettingDefinition[T]:
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			res, err = retryOnEtagError(func(setting T) (string, error) { return defn.Update(ctx, a, setting) }, setting, defn.SetETag)
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
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var setting T
			return createOrUpdate(ctx, d, c, setting)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var res *T
			switch defn := defn.(type) {
			case workspaceSettingDefinition[T]:
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				res, err = retryOnEtagError(func(etag string) (*T, error) { return defn.Read(ctx, w, etag) }, d.Id(), func(req *string, newEtag string) { *req = newEtag })
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				res, err = retryOnEtagError(func(etag string) (*T, error) { return defn.Read(ctx, a, etag) }, d.Id(), func(req *string, newEtag string) { *req = newEtag })
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
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var setting T
			defn.SetETag(&setting, d.Id())
			return createOrUpdate(ctx, d, c, setting)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var etag string
			switch defn := defn.(type) {
			case workspaceSettingDefinition[T]:
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				etag, err = retryOnEtagError(func(etag string) (string, error) { return defn.Delete(ctx, w, etag) }, d.Id(), func(req *string, newEtag string) { *req = newEtag })
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				etag, err = retryOnEtagError(func(etag string) (string, error) { return defn.Delete(ctx, a, etag) }, d.Id(), func(req *string, newEtag string) { *req = newEtag })
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
