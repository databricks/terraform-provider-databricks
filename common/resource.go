package common

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type WorkspaceIdField int

const (
	// WorkspaceId is the defalt value, since all non-UC workspace resources belong to a workspace.
	WorkspaceId WorkspaceIdField = iota
	// Do not expose a workspace_id field. This is used for account-level-only resources.
	NoWorkspaceId
	// Use management_workspace_id to specify the workspace_id. This is used for UC resources,
	// where resources do not belong to a workspace but are managed using a specific workspace's
	// API.
	ManagementWorkspaceId
	// Used for deletion.
	OriginalWorkspaceId
)

func (w WorkspaceIdField) Field() string {
	switch w {
	case WorkspaceId:
		return "workspace_id"
	case NoWorkspaceId:
		return ""
	case ManagementWorkspaceId:
		return "management_workspace_id"
	case OriginalWorkspaceId:
		return "original_workspace_id"
	default:
		panic(fmt.Sprintf("unknown workspace_id field type: %d", w))
	}
}

func (w WorkspaceIdField) IsUserSpecified() bool {
	switch w {
	case WorkspaceId:
		return true
	case ManagementWorkspaceId:
		return true
	default:
		return false
	}
}

// Resource aims to simplify things like error & deleted entities handling
type Resource struct {
	Create             func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Read               func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Update             func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	Delete             func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error
	CustomizeDiff      func(ctx context.Context, d *schema.ResourceDiff) error
	StateUpgraders     []schema.StateUpgrader
	Schema             map[string]*schema.Schema
	SchemaVersion      int
	Timeouts           *schema.ResourceTimeout
	DeprecationMessage string
	Importer           *schema.ResourceImporter

	// WorkspaceIdField indicates whether a resource can only be managed using a
	// provider at the appropriate level, i.e. workspace-level resources that can only be
	// managed with a workspace-level provider, and if so, what the name is of the field
	// that should be used to specify the workspace_id.
	//
	// All account-level resources can only be managed with an account-level provider, so this
	// must be set to `true` for all account-level resources.
	//
	// When false, the resource will include a workspace_id field that can be set when using an
	// account-level provider to target a specific workspace in the account.
	WorkspaceIdField WorkspaceIdField
	// workspace_id = check that specified workspace_id matches current workspace ID when using workspace-level provider
	// What happens if your workspace client still doesn't work yet for existing resource?
	//   We could always block and ask to apply once cleanly, then future applies should work.
	// can only be skipped if workspace_id is known after apply (new resource case)
	// management_workspace_id = should never be in the diff.
}

func nicerError(ctx context.Context, err error, action string) error {
	name := ResourceName.GetOrUnknown(ctx)
	if name == "unknown" {
		return err
	}
	return fmt.Errorf("cannot %s %s: %w", action,
		strings.ReplaceAll(name, "_", " "), err)
}

func recoverable(cb func(
	ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error) func(
	ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
	return func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) (err error) {
		defer func() {
			// this is deliberate decision to convert a panic into error,
			// so that any unforeseen bug would we visible to end-user
			// as an error and not a provider crash, which is way less
			// of pleasant experience.
			if panic := recover(); panic != nil {
				err = fmt.Errorf("panic: %v", panic)
			}
		}()
		err = cb(ctx, d, c)
		return
	}
}

func (r Resource) verifyWorkspaceId(ctx context.Context, rd *schema.ResourceDiff, c *DatabricksClient) error {
	if !r.WorkspaceIdField.IsUserSpecified() {
		return nil
	}
	if !rd.NewValueKnown(r.WorkspaceIdField.Field()) {
		return nil
	}
	if c.Config.IsAccountClient() {
		return nil
	}
	configuredWorkspaceId := rd.Get(r.WorkspaceIdField.Field()).(int64)
	currentWorkspaceId, err := c.CurrentWorkspaceID(ctx)
	if err != nil {
		return err
	}
	if configuredWorkspaceId != currentWorkspaceId {
		return fmt.Errorf("configured %s (%d) does not match current workspace ID (%d)", r.WorkspaceIdField.Field(), configuredWorkspaceId, currentWorkspaceId)
	}
	return nil
}

func (r Resource) saferCustomizeDiff() schema.CustomizeDiffFunc {
	if !r.isResource() {
		return nil
	}
	return func(ctx context.Context, rd *schema.ResourceDiff, m any) (err error) {
		defer func() {
			// this is deliberate decision to convert a panic into error,
			// so that any unforeseen bug would we visible to end-user
			// as an error and not a provider crash, which is way less
			// of pleasant experience.
			if panic := recover(); panic != nil {
				err = nicerError(ctx, fmt.Errorf("panic: %v", panic),
					"customize diff for")
			}
		}()
		// TODO: check that the provided workspace ID matches that for the provider
		c := m.(*DatabricksClient)
		err = r.verifyWorkspaceId(ctx, rd, c)
		if err != nil {
			return err
		}
		// Note: you cannot clear workspace_id field because it is not computed.
		// We could make it computed, but I think that would be too implicit for users.

		// we don't propagate instance of SDK client to the diff function, because
		// authentication is not deterministic at this stage with the recent Terraform
		// versions. Diff customization must be limited to hermetic checks only anyway.
		if r.CustomizeDiff != nil {
			err = r.CustomizeDiff(ctx, rd)
			if err != nil {
				err = nicerError(ctx, err, "customize diff for")
				return
			}
		}
		return
	}
}

func (r Resource) getSchema() map[string]*schema.Schema {
	if r.Schema == nil {
		return make(map[string]*schema.Schema)
	}
	return r.Schema
}

func (r Resource) isResource() bool {
	return r.Create != nil || r.Update != nil || r.Delete != nil
}

// ToResource converts to Terraform resource definition
func (r Resource) ToResource() *schema.Resource {
	if r.WorkspaceIdField == OriginalWorkspaceId {
		panic("OriginalWorkspaceId is not a valid value for WorkspaceIdField")
	}
	getClient := func(_ context.Context, m any, d *schema.ResourceData) (c *DatabricksClient, err error) {
		return m.(*DatabricksClient), nil
	}
	getDeleteClient := getClient
	if r.WorkspaceIdField != NoWorkspaceId {
		getClient = func(ctx context.Context, m any, d *schema.ResourceData) (c *DatabricksClient, err error) {
			return m.(*DatabricksClient).InConfiguredWorkspace(ctx, d, r.WorkspaceIdField)
		}
		getDeleteClient = func(ctx context.Context, m any, d *schema.ResourceData) (c *DatabricksClient, err error) {
			return m.(*DatabricksClient).InConfiguredWorkspace(ctx, d, OriginalWorkspaceId)
		}
	}
	// Ignore missing for read for resources, but not for data sources.
	ignoreMissingForRead := r.isResource()
	generateReadFunc := func(ignoreMissing bool) func(ctx context.Context, d *schema.ResourceData,
		m any) diag.Diagnostics {
		return func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
			c, err := getClient(ctx, m, d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = recoverable(r.Read)(ctx, d, c)
			// TODO: https://github.com/databricks/terraform-provider-databricks/issues/2021
			if ignoreMissing && apierr.IsMissing(err) {
				log.Printf("[INFO] %s[id=%s] is removed on backend",
					ResourceName.GetOrUnknown(ctx), d.Id())
				d.SetId("")
				return nil
			}
			if err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		}
	}
	resource := &schema.Resource{
		Schema:             r.getSchema(),
		SchemaVersion:      r.SchemaVersion,
		StateUpgraders:     r.StateUpgraders,
		CustomizeDiff:      r.saferCustomizeDiff(),
		ReadContext:        generateReadFunc(ignoreMissingForRead),
		Importer:           r.Importer,
		Timeouts:           r.Timeouts,
		DeprecationMessage: r.DeprecationMessage,
	}
	if r.Create != nil {
		resource.CreateContext = func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			c, err := getClient(ctx, m, d)
			if err != nil {
				return diag.FromErr(err)
			}
			if err = recoverable(r.Create)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "create")
				return diag.FromErr(err)
			}
			if err = recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		}
	}
	if r.Delete != nil {
		resource.DeleteContext = func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			c, err := getDeleteClient(ctx, m, d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = recoverable(r.Delete)(ctx, d, c)
			if apierr.IsMissing(err) {
				// TODO: https://github.com/databricks/terraform-provider-databricks/issues/2021
				log.Printf("[INFO] %s[id=%s] is removed on backend",
					ResourceName.GetOrUnknown(ctx), d.Id())
				d.SetId("")
				return nil
			}
			if err != nil {
				err = nicerError(ctx, err, "delete")
				return diag.FromErr(err)
			}
			return nil
		}
	}
	if resource.Importer == nil {
		resource.Importer = &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData,
				m any) (data []*schema.ResourceData, e error) {
				d.MarkNewResource()
				diags := generateReadFunc(false)(ctx, d, m)
				var err error
				if diags.HasError() {
					err = diags[0].Validate()
				}
				return []*schema.ResourceData{d}, err
			},
		}
	}
	if r.WorkspaceIdField != NoWorkspaceId {
		resource.Schema = AddWorkspaceIdField(resource.Schema, r.WorkspaceIdField)
	}
	if r.Update != nil {
		resource.UpdateContext = func(ctx context.Context, d *schema.ResourceData,
			m any) diag.Diagnostics {
			c, err := getClient(ctx, m, d)
			if err != nil {
				return diag.FromErr(err)
			}
			if err := recoverable(r.Update)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "update")
				return diag.FromErr(err)
			}
			if err := recoverable(r.Read)(ctx, d, c); err != nil {
				err = nicerError(ctx, err, "read")
				return diag.FromErr(err)
			}
			return nil
		}
	} else {
		// set ForceNew to all attributes with CRD
		queue := []*schema.Resource{
			{Schema: r.Schema},
		}
		for {
			head := queue[0]
			queue = queue[1:]
			for _, v := range head.Schema {
				if v.Computed {
					continue
				}
				if nested, ok := v.Elem.(*schema.Resource); ok {
					queue = append(queue, nested)
				}
				v.ForceNew = true
			}
			if len(queue) == 0 {
				break
			}
		}
	}
	return resource
}

func MustCompileKeyRE(name string) *regexp.Regexp {
	regexFromName := strings.ReplaceAll(name, ".", "\\.")
	regexFromName = strings.ReplaceAll(regexFromName, ".0", ".\\d+")
	return regexp.MustCompile(regexFromName)
}

// Deprecated: migrate to WorkspaceData
func DataResource(sc any, read func(context.Context, any, *DatabricksClient) error) Resource {
	// TODO: migrate to go1.18 and get schema from second function argument?..
	s := StructToSchema(sc, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	return Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *DatabricksClient) (err error) {
			ptr := reflect.New(reflect.ValueOf(sc).Type())
			DataToReflectValue(d, s, ptr.Elem())
			err = read(ctx, ptr.Interface(), m)
			if err != nil {
				err = nicerError(ctx, err, "read data")
			}
			StructToData(ptr.Elem().Interface(), s, d)
			// check if the resource schema has the `id` attribute (marked with `json:"id"` in the provided structure).
			// and if yes, then use it as resource ID. If not, then use default value for resource ID (`_`)
			if _, ok := s["id"]; ok {
				d.SetId(d.Get("id").(string))
			} else {
				d.SetId("_")
			}
			return
		},
	}
}

// WorkspaceData is a generic way to define workspace data resources in Terraform provider.
//
// Example usage:
//
//	type catalogsData struct {
//		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
//	}
//	return common.WorkspaceData(func(ctx context.Context, data *catalogsData, w *databricks.WorkspaceClient) error {
//		catalogs, err := w.Catalogs.ListAll(ctx)
//		...
//	})
func WorkspaceData[T any](read func(context.Context, *T, *databricks.WorkspaceClient) error) Resource {
	return genericDatabricksData((*DatabricksClient).WorkspaceClient, func(ctx context.Context, s struct{}, t *T, wc *databricks.WorkspaceClient) error {
		return read(ctx, t, wc)
	}, false)
}

// WorkspaceDataWithParams defines a data source that can be used to read data from the workspace API.
// It differs from WorkspaceData in that it separates the definition of the computed fields (the resource type)
// from the definition of the user-supplied parameters.
//
// The first type parameter is the type of the resource. This can be a type directly from the SDK, or a custom
// type defined in the provider that embeds the SDK type.
//
// The second type parameter is the type representing parameters that a user may provide to the data source. These
// are the attributes that the user can specify in the data source configuration, but are not part of the resource
// type. If there are no extra attributes, this should be `struct{}`. If there are any fields with the same JSON
// name as fields in the resource type, these fields will override the values from the resource type.
//
// The single argument is a function that will be called to read the data from the workspace API, returning the
// value of the resource type. The function should return an error if the data cannot be read or the resource cannot
// be found.
//
// Example usage:
//
//	 type SqlWarehouse struct { ... }
//
//	 type SqlWarehouseDataParams struct {
//		     Id   string `json:"id" tf:"computed,optional"`
//		     Name string `json:"name" tf:"computed,optional"`
//	 }
//
//	 WorkspaceDataWithParams(
//	     func(ctx context.Context, data SqlWarehouseDataParams, w *databricks.WorkspaceClient) (*SqlWarehouse, error) {
//	         // User-provided attributes are present in the `data` parameter.
//	         // The resource should be returned.
//	         ...
//	     })
func WorkspaceDataWithParams[T, P any](read func(context.Context, P, *databricks.WorkspaceClient) (*T, error)) Resource {
	return genericDatabricksData((*DatabricksClient).WorkspaceClient, func(ctx context.Context, o P, s *T, w *databricks.WorkspaceClient) error {
		res, err := read(ctx, o, w)
		if err != nil {
			return err
		}
		*s = *res
		return nil
	}, true)
}

// AccountData is a generic way to define account data resources in Terraform provider.
//
// Example usage:
//
//	type metastoresData struct {
//		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
//	}
//	return common.AccountData(func(ctx context.Context, d *metastoresData, acc *databricks.AccountClient) error {
//		metastores, err := acc.Metastores.List(ctx)
//		...
//	})
func AccountData[T any](read func(context.Context, *T, *databricks.AccountClient) error) Resource {
	return genericDatabricksData((*DatabricksClient).AccountClient, func(ctx context.Context, s struct{}, t *T, ac *databricks.AccountClient) error {
		return read(ctx, t, ac)
	}, false)
}

// AccountDataWithParams defines a data source that can be used to read data from the workspace API.
// It differs from AccountData in that it allows extra attributes to be provided as a separate argument,
// so the original type used to define the resource can also be used to define the data source.
//
// The first type parameter is the type of the resource. This can be a type directly from the SDK, or a custom
// type defined in the provider that embeds the SDK type.
//
// The second type parameter is the type of the extra attributes that should be provided to the data source. These
// are the attributes that the user can specify in the data source configuration, but are not part of the resource
// type. If there are no extra attributes, this should be `struct{}`. If there are any fields with the same JSON
// name as fields in the resource type, these fields will override the values from the resource type.
//
// The single argument is a function that will be called to read the data from the workspace API, returning the
// requested resource. The function should return an error if the data cannot be read or the resource cannot be
// found.
//
// Example usage:
//
//	 type MwsWorkspace struct { ... }
//
//	 type MwsWorkspaceDataParams struct {
//		     Id   string `json:"id" tf:"computed,optional"`
//		     Name string `json:"name" tf:"computed,optional"`
//	 }
//
//	 AccountDataWithParams(
//	     func(ctx context.Context, data MwsWorkspaceDataParams, a *databricks.AccountClient) (*MwsWorkspace, error) {
//	         // User-provided attributes are present in the `data` parameter.
//	         // The resource should be populated in the `workspace` parameter.
//	         ...
//		  })
func AccountDataWithParams[T, P any](read func(context.Context, P, *databricks.AccountClient) (*T, error)) Resource {
	return genericDatabricksData((*DatabricksClient).AccountClient, func(ctx context.Context, o P, s *T, a *databricks.AccountClient) error {
		res, err := read(ctx, o, a)
		if err != nil {
			return err
		}
		*s = *res
		return nil
	}, true)
}

// genericDatabricksData is generic and common way to define both account and workspace data and calls their respective clients.
//
// If hasOther is true, all of the fields of SdkType will be marked as computed in the final schema, and the fields
// from OtherFields will be overlaid on top of the schema generated by SdkType. Otherwise, the schema generated by
// SdkType will be used directly.
func genericDatabricksData[T, P, C any](
	getClient func(*DatabricksClient) (C, error),
	read func(context.Context, P, *T, C) error,
	hasOther bool) Resource {
	var dummy T
	var other P
	otherFields := StructToSchema(other, NoCustomize)
	s := StructToSchema(dummy, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		// For WorkspaceData and AccountData, a single data type is used to represent all of the fields of
		// the resource, so its configuration is correct. For the *WithParams methods, the SdkType parameter
		// is copied directly from the resource definition, which means that all fields from that type are
		// computed and optional, and the fields from OtherFields are overlaid on top of the schema generated
		// by SdkType.
		if hasOther {
			for k := range m {
				m[k].Computed = true
				m[k].Required = false
				m[k].Optional = true
			}
			for k, v := range otherFields {
				m[k] = v
			}
		}
		// `id` attribute must be marked as computed, otherwise it's not set!
		if v, ok := m["id"]; ok {
			v.Computed = true
			v.Required = false
		}
		return m
	})
	return Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, client *DatabricksClient) (err error) {
			defer func() {
				// using recoverable() would cause more complex rewrapping of DataToStructPointer & StructToData
				if panic := recover(); panic != nil {
					err = fmt.Errorf("panic: %v", panic)
				}
			}()
			var dummy T
			var other P
			DataToStructPointer(d, s, &other)
			DataToStructPointer(d, s, &dummy)
			c, err := getClient(client)
			if err != nil {
				return nicerError(ctx, err, "get client")
			}
			err = read(ctx, other, &dummy, c)
			if err != nil {
				err = nicerError(ctx, err, "read data")
			}
			StructToData(&dummy, s, d)
			// check if the resource schema has the `id` attribute (marked with `json:"id"` in the provided structure).
			// and if yes, then use it as resource ID. If not, then use default value for resource ID (`_`)
			if _, ok := s["id"]; ok {
				d.SetId(d.Get("id").(string))
			} else {
				d.SetId("_")
			}
			return
		},
	}
}

func EqualFoldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if strings.EqualFold(old, new) {
		log.Printf("[INFO] Suppressing diff on %s", k)
		return true
	}
	return false
}

func NoCustomize(m map[string]*schema.Schema) map[string]*schema.Schema {
	return m
}

var NoAuth string = "default auth: cannot configure default credentials, " +
	"please check https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication " +
	"to configure credentials for your preferred authentication method"

func AddAccountIdField(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["account_id"] = &schema.Schema{
		Type:       schema.TypeString,
		Optional:   true,
		Deprecated: "Configuring `account_id` at the resource-level is deprecated; please specify it in the `provider {}` configuration block instead",
	}
	return s
}

func AddWorkspaceIdField(s map[string]*schema.Schema, f WorkspaceIdField) map[string]*schema.Schema {
	s[f.Field()] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		// Changing the workspace does require force new, but adding it for an existing resource results in no diff.
		Description: "The workspace id of the workspace, for example 1234567890. This can be retrieved from `databricks_mws_workspaces.<YOUR_WORKSPACE>.workspace_id`. This attribute is required when using an account-level provider.",
	}
	// Separate field to track the original workspace ID when using an account-level provider.
	// Needed when changing the workspace ID, as `workspace_id` in state will reflect
	// the new ID during apply.
	s[OriginalWorkspaceId.Field()] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: false,
		Required: false,
		Computed: true,
	}
	return s
}
