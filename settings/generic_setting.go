package settings

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	defaultSettingId = "global"
	etagAttrName     = "etag"
)

func retryOnEtagError[Req, Resp any](f func(req Req) (Resp, error), firstReq Req, updateReq func(req *Req, newEtag string), retriableErrors []error) (Resp, error) {
	req := firstReq
	// Retry once on etag error.
	res, err := f(req)
	if err == nil {
		return res, nil
	}
	if !isRetriableError(err, retriableErrors) {
		return res, err
	}
	etag, err := getEtagFromError(err)
	if err != nil {
		return res, err
	}
	updateReq(&req, etag)
	return f(req)
}

func isRetriableError(err error, retriableErrors []error) bool {
	for _, retriableError := range retriableErrors {
		if errors.Is(err, retriableError) {
			return true
		}
	}
	return false
}

func getEtagFromError(err error) (string, error) {
	errorInfos := apierr.GetErrorInfo(err)
	if len(errorInfos) > 0 {
		metadata := errorInfos[0].Metadata
		if etag, ok := metadata[etagAttrName]; ok {
			return etag, nil
		}
	}
	return "", fmt.Errorf("error fetching the settings: %w", err)
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

	// Generate resource ID from settings instance
	GetId(t *T) string

	// Schema customization function
	GetCustomizeSchemaFunc() func(map[string]*schema.Schema) map[string]*schema.Schema
}

func getEtag[T any](t T) string {
	rv := reflect.ValueOf(t)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv.FieldByName("Etag").String()
}

func setEtag[T any](t T, newEtag string) {
	rv := reflect.ValueOf(t)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rv.FieldByName("Etag").SetString(newEtag)
}

type workspaceSettingDefinition[T any] genericSettingDefinition[T, *databricks.WorkspaceClient]

// A workspace setting is a setting that is scoped to a workspace.
type workspaceSetting[T any] struct {
	// The struct corresponding to the setting. The schema of the Terraform resource will be generated from this struct.
	// This struct must have an Etag field of type string.
	settingStruct T

	// Read the setting from the server. The etag is provided as the third argument.
	readFunc func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*T, error)

	// Update the setting to the value specified by t, and return the new etag. If the setting name is user-settable,
	// it will be provided in the third argument. If not, you must set the SettingName field appropriately. You must
	// also set AllowMissing: true and the field mask to the field to update.
	updateFunc func(ctx context.Context, w *databricks.WorkspaceClient, setting T) (string, error)

	// Delete the setting with the given etag, and return the new etag.
	deleteFunc func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error)

	// Optional function to generate resource ID from the settings. If not provided, will use predefined value `global`
	generateIdFunc func(setting *T) string

	// Optional function to customize the schema. If not provided, will use the default customization
	customizeSchemaFunc func(map[string]*schema.Schema) map[string]*schema.Schema
}

func (w workspaceSetting[T]) SettingStruct() T {
	return w.settingStruct
}
func (w workspaceSetting[T]) Read(ctx context.Context, c *databricks.WorkspaceClient, etag string) (*T, error) {
	return w.readFunc(ctx, c, etag)
}
func (w workspaceSetting[T]) Update(ctx context.Context, c *databricks.WorkspaceClient, t T) (string, error) {
	return w.updateFunc(ctx, c, t)
}
func (w workspaceSetting[T]) Delete(ctx context.Context, c *databricks.WorkspaceClient, etag string) (string, error) {
	return w.deleteFunc(ctx, c, etag)
}
func (w workspaceSetting[T]) GetETag(t *T) string {
	return getEtag(t)
}

func (w workspaceSetting[T]) GetId(t *T) string {
	id := defaultSettingId
	if w.generateIdFunc != nil {
		id = w.generateIdFunc(t)
	}
	return id
}

func (w workspaceSetting[T]) GetCustomizeSchemaFunc() func(map[string]*schema.Schema) map[string]*schema.Schema {
	if w.customizeSchemaFunc != nil {
		return w.customizeSchemaFunc
	}
	return func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s[etagAttrName].Computed = true
		// Note: this may not always be computed, but it is for the default namespace setting. If other settings
		// are added for which setting_name is not computed, we'll need to expose this somehow as part of the setting
		// definition.
		s["setting_name"].Computed = true
		return s
	}
}

func (w workspaceSetting[T]) SetETag(t *T, newEtag string) {
	setEtag(t, newEtag)
}

var _ workspaceSettingDefinition[struct{}] = workspaceSetting[struct{}]{}

type accountSettingDefinition[T any] genericSettingDefinition[T, *databricks.AccountClient]

// An account setting is a setting that is scoped to an account.
type accountSetting[T any] struct {
	// The struct corresponding to the setting. The schema of the Terraform resource will be generated from this struct.
	// This struct must have an Etag field of type string.
	settingStruct T

	// Read the setting from the server. The etag is provided as the third argument.
	readFunc func(ctx context.Context, w *databricks.AccountClient, etag string) (*T, error)

	// Update the setting to the value specified by t, and return the new etag. If the setting name is user-settable,
	// it will be provided in the third argument. If not, you must set the SettingName field appropriately. You must
	// also set AllowMissing: true and the field mask to the field to update.
	updateFunc func(ctx context.Context, w *databricks.AccountClient, setting T) (string, error)

	// Delete the setting with the given etag, and return the new etag.
	deleteFunc func(ctx context.Context, w *databricks.AccountClient, etag string) (string, error)

	// Optional function to generate resource ID from the settings. If not provided, will use predefined value `global`
	generateIdFunc func(setting *T) string

	// Optional function to customize the schema. If not provided, will use the default customization
	customizeSchemaFunc func(map[string]*schema.Schema) map[string]*schema.Schema
}

func (w accountSetting[T]) SettingStruct() T {
	return w.settingStruct
}
func (w accountSetting[T]) Read(ctx context.Context, c *databricks.AccountClient, etag string) (*T, error) {
	return w.readFunc(ctx, c, etag)
}
func (w accountSetting[T]) Update(ctx context.Context, c *databricks.AccountClient, t T) (string, error) {
	return w.updateFunc(ctx, c, t)
}
func (w accountSetting[T]) Delete(ctx context.Context, c *databricks.AccountClient, etag string) (string, error) {
	return w.deleteFunc(ctx, c, etag)
}
func (w accountSetting[T]) GetETag(t *T) string {
	return getEtag(t)
}
func (w accountSetting[T]) SetETag(t *T, newEtag string) {
	setEtag(t, newEtag)
}

func (w accountSetting[T]) GetId(t *T) string {
	id := defaultSettingId
	if w.generateIdFunc != nil {
		id = w.generateIdFunc(t)
	}
	return id
}

func (w accountSetting[T]) GetCustomizeSchemaFunc() func(map[string]*schema.Schema) map[string]*schema.Schema {
	if w.customizeSchemaFunc != nil {
		return w.customizeSchemaFunc
	}
	return func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s[etagAttrName].Computed = true
		// Note: this may not always be computed, but it is for the default namespace setting. If other settings
		// are added for which setting_name is not computed, we'll need to expose this somehow as part of the setting
		// definition.
		s["setting_name"].Computed = true
		return s
	}
}

var _ accountSettingDefinition[struct{}] = accountSetting[struct{}]{}

type accountWorkspaceSettingDefinition[T any] genericSettingDefinition[T, *common.DatabricksClient]

// An account workspace setting is a setting that can be scoped to either an account or a workspace.
type accountWorkspaceSetting[T any] struct {
	// The struct corresponding to the setting. The schema of the Terraform resource will be generated from this struct.
	// This struct must have an Etag field of type string.
	settingStruct T

	// Read the setting from the server. The etag is provided as the third argument.
	readAccFunc func(ctx context.Context, acc *databricks.AccountClient, etag string) (*T, error)

	readWsFunc func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*T, error)

	// Update the setting to the value specified by t, and return the new etag. If the setting name is user-settable,
	// it will be provided in the third argument. If not, you must set the SettingName field appropriately. You must
	// also set AllowMissing: true and the field mask to the field to update.
	updateAccFunc func(ctx context.Context, w *databricks.AccountClient, setting T) (string, error)

	updateWsFunc func(ctx context.Context, w *databricks.WorkspaceClient, setting T) (string, error)

	// Delete the setting with the given etag, and return the new etag.
	deleteAccFunc func(ctx context.Context, acc *databricks.AccountClient, etag string) (string, error)

	deleteWsFunc func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error)

	// Optional function to generate resource ID from the settings. If not provided, will use predefined value `global`
	generateIdFunc func(setting *T) string

	// Optional function to customize the schema. If not provided, will use the default customization
	customizeSchemaFunc func(map[string]*schema.Schema) map[string]*schema.Schema
}

func (aw accountWorkspaceSetting[T]) SettingStruct() T {
	return aw.settingStruct
}
func (aw accountWorkspaceSetting[T]) Read(ctx context.Context, c *common.DatabricksClient, etag string) (*T, error) {
	if c.Config.IsAccountClient() {
		a, err := c.AccountClient()
		if err != nil {
			return nil, err
		}
		return aw.readAccFunc(ctx, a, etag)
	} else {
		ws, err := c.WorkspaceClient()
		if err != nil {
			return nil, err
		}
		return aw.readWsFunc(ctx, ws, etag)
	}
}
func (aw accountWorkspaceSetting[T]) Update(ctx context.Context, c *common.DatabricksClient, t T) (string, error) {
	if c.Config.IsAccountClient() {
		a, err := c.AccountClient()
		if err != nil {
			return "", err
		}
		return aw.updateAccFunc(ctx, a, t)
	} else {
		ws, err := c.WorkspaceClient()
		if err != nil {
			return "", err
		}
		return aw.updateWsFunc(ctx, ws, t)
	}
}
func (aw accountWorkspaceSetting[T]) Delete(ctx context.Context, c *common.DatabricksClient, etag string) (string, error) {
	if c.Config.IsAccountClient() {
		a, err := c.AccountClient()
		if err != nil {
			return "", err
		}
		return aw.deleteAccFunc(ctx, a, etag)
	} else {
		ws, err := c.WorkspaceClient()
		if err != nil {
			return "", err
		}
		return aw.deleteWsFunc(ctx, ws, etag)
	}
}
func (aw accountWorkspaceSetting[T]) GetETag(t *T) string {
	return getEtag(t)
}
func (aw accountWorkspaceSetting[T]) SetETag(t *T, newEtag string) {
	setEtag(t, newEtag)
}

func (aw accountWorkspaceSetting[T]) GetId(t *T) string {
	id := defaultSettingId
	if aw.generateIdFunc != nil {
		id = aw.generateIdFunc(t)
	}
	return id
}

func (aw accountWorkspaceSetting[T]) GetCustomizeSchemaFunc() func(map[string]*schema.Schema) map[string]*schema.Schema {
	if aw.customizeSchemaFunc != nil {
		return aw.customizeSchemaFunc
	}
	return func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s[etagAttrName].Computed = true
		// Note: this may not always be computed, but it is for the default namespace setting. If other settings
		// are added for which setting_name is not computed, we'll need to expose this somehow as part of the setting
		// definition.
		s["setting_name"].Computed = true
		return s
	}
}

var _ accountWorkspaceSettingDefinition[struct{}] = accountWorkspaceSetting[struct{}]{}

func makeSettingResource[T, U any](defn genericSettingDefinition[T, U]) common.Resource {
	resourceSchema := common.StructToSchema(defn.SettingStruct(),
		defn.GetCustomizeSchemaFunc())
	createOrUpdateRetriableErrors := []error{apierr.ErrNotFound, apierr.ErrResourceConflict}
	deleteRetriableErrors := []error{apierr.ErrResourceConflict}
	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient, setting T) error {
		common.DataToStructPointer(d, resourceSchema, &setting)
		var res string
		switch defn := defn.(type) {
		case workspaceSettingDefinition[T]:
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			res, err = retryOnEtagError[T, string](
				func(setting T) (string, error) {
					return defn.Update(ctx, w, setting)
				},
				setting,
				defn.SetETag,
				createOrUpdateRetriableErrors)
			if err != nil {
				return err
			}
		case accountSettingDefinition[T]:
			a, err := c.AccountClient()
			if err != nil {
				return err
			}
			res, err = retryOnEtagError(
				func(setting T) (string, error) {
					return defn.Update(ctx, a, setting)
				},
				setting,
				defn.SetETag,
				createOrUpdateRetriableErrors)
			if err != nil {
				return err
			}
		case accountWorkspaceSettingDefinition[T]:
			var err error
			res, err = retryOnEtagError(
				func(setting T) (string, error) {
					return defn.Update(ctx, c, setting)
				},
				setting,
				defn.SetETag,
				createOrUpdateRetriableErrors)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected setting type: %T", defn)
		}
		d.Set(etagAttrName, res)
		d.SetId(defn.GetId(&setting))
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
				res, err = defn.Read(ctx, w, d.Get(etagAttrName).(string))
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				res, err = defn.Read(ctx, a, d.Get(etagAttrName).(string))
				if err != nil {
					return err
				}
			case accountWorkspaceSettingDefinition[T]:
				var err error
				res, err = defn.Read(ctx, c, d.Get(etagAttrName).(string))
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
			d.Set(etagAttrName, defn.GetETag(res))
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var setting T
			defn.SetETag(&setting, d.Get(etagAttrName).(string))
			return createOrUpdate(ctx, d, c, setting)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var etag string
			updateETag := func(req *string, newEtag string) { *req = newEtag }
			switch defn := defn.(type) {
			case workspaceSettingDefinition[T]:
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				etag, err = retryOnEtagError(
					func(etag string) (string, error) {
						return defn.Delete(ctx, w, etag)
					},
					d.Get(etagAttrName).(string),
					updateETag,
					deleteRetriableErrors)
				if err != nil {
					return err
				}
			case accountSettingDefinition[T]:
				a, err := c.AccountClient()
				if err != nil {
					return err
				}
				etag, err = retryOnEtagError(
					func(etag string) (string, error) {
						return defn.Delete(ctx, a, etag)
					},
					d.Get(etagAttrName).(string),
					updateETag,
					deleteRetriableErrors)
				if err != nil {
					return err
				}
			case accountWorkspaceSettingDefinition[T]:
				var err error
				etag, err = retryOnEtagError(
					func(etag string) (string, error) {
						return defn.Delete(ctx, c, etag)
					},
					d.Get(etagAttrName).(string),
					updateETag,
					deleteRetriableErrors)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected setting type: %T", defn)
			}
			d.Set(etagAttrName, etag)
			return nil
		},
	}
}
