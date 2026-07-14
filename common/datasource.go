package common

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Deprecated: migrate to WorkspaceData
func DataResource(sc any, read func(context.Context, any, *DatabricksClient) error) Resource {
	// TODO: migrate to go1.18 and get schema from second function argument?..
	s := StructToSchema(sc, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	AddNamespaceInSchema(s)
	NamespaceCustomizeSchemaMap(s)
	return Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *DatabricksClient) (err error) {
			newClient, err := m.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			ptr := reflect.New(reflect.ValueOf(sc).Type())
			DataToReflectValue(d, s, ptr.Elem())
			err = read(ctx, ptr.Interface(), newClient)
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
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (*databricks.WorkspaceClient, error) {
			return client.WorkspaceClientUnifiedProvider(ctx, d)
		},
		func(ctx context.Context, s T, t *T, wc *databricks.WorkspaceClient) error {
			return read(ctx, t, wc)
		}, false, NamespaceCustomizeSchemaMap)
}

// WorkspaceDataWithParams defines a data source that can be used to read data from the workspace API.
// It differs from WorkspaceData in that it separates the definition of the computed fields (the resource type)
// from the definition of the user-supplied parameters.
//
// The first type parameter is the type representing parameters that a user may provide to the data source. These
// are the attributes that the user can specify in the data source configuration, but are not part of the resource
// type. If there are no extra attributes, this should be `struct{}`. If there are any fields with the same JSON
// name as fields in the resource type, these fields will override the values from the resource type.
//
// The second type parameter is the type of the resource. This can be a type directly from the SDK, or a custom
// type defined in the provider that embeds the SDK type.
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
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (*databricks.WorkspaceClient, error) {
			return client.WorkspaceClientUnifiedProvider(ctx, d)
		},
		func(ctx context.Context, o P, s *T, w *databricks.WorkspaceClient) error {
			res, err := read(ctx, o, w)
			if err != nil {
				return err
			}
			*s = *res
			return nil
		}, true, NamespaceCustomizeSchemaMap)
}

// WorkspaceDataWithCustomizeFunc defines a data source that can be used to read data from the workspace API.
// It differs from WorkspaceData in that it allows the schema to be customized further using a
// customizeSchemaFunc function.
//
// The additional argument is a function that will be called to customize the schema of the data source.
func WorkspaceDataWithCustomizeFunc[T any](
	read func(context.Context, *T, *databricks.WorkspaceClient) error,
	customizeSchemaFunc func(map[string]*schema.Schema) map[string]*schema.Schema) Resource {
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (*databricks.WorkspaceClient, error) {
			return client.WorkspaceClientUnifiedProvider(ctx, d)
		},
		func(ctx context.Context, s struct{}, t *T, wc *databricks.WorkspaceClient) error {
			return read(ctx, t, wc)
		}, false, customizeSchemaFunc)
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
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (*databricks.AccountClient, error) {
			return client.AccountClient()
		},
		func(ctx context.Context, s struct{}, t *T, ac *databricks.AccountClient) error {
			return read(ctx, t, ac)
		}, false, NoCustomize)
}

// AccountDataWithParams defines a data source that can be used to read data from the account API.
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
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (*databricks.AccountClient, error) {
			return client.AccountClient()
		},
		func(ctx context.Context, o P, s *T, a *databricks.AccountClient) error {
			res, err := read(ctx, o, a)
			if err != nil {
				return err
			}
			*s = *res
			return nil
		}, true, NoCustomize)
}

// genericDatabricksData is generic and common way to define both account and workspace data and calls their respective clients.
//
// If hasOther is true, all of the fields of SdkType will be marked as computed in the final schema, and the fields
// from OtherFields will be overlaid on top of the schema generated by SdkType. Otherwise, the schema generated by
// SdkType will be used directly.
func genericDatabricksData[T, P, C any](
	getClient func(*DatabricksClient, context.Context, *schema.ResourceData) (C, error),
	read func(context.Context, P, *T, C) error,
	hasOther bool,
	customizeSchemaFunc func(map[string]*schema.Schema) map[string]*schema.Schema) Resource {
	var dummy T
	var other P
	otherFields := StructToSchema(other, nil)

	s := StructToSchema(dummy, nil)
	// For WorkspaceData and AccountData, a single data type is used to represent all of the fields of
	// the resource, so its configuration is correct. For the *WithParams methods, the SdkType parameter
	// is copied directly from the resource definition, which means that all fields from that type are
	// computed and optional, and the fields from OtherFields are overlaid on top of the schema generated
	// by SdkType.
	if hasOther {
		for k := range s {
			s[k].Computed = true
			s[k].Required = false
			s[k].Optional = true
		}
		for k, v := range otherFields {
			s[k] = v
		}
	}
	// `id` attribute must be marked as computed, otherwise it's not set!
	if v, ok := s["id"]; ok {
		v.Computed = true
		v.Required = false
	}
	// allow c
	s = customizeSchemaFunc(s)

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
			c, err := getClient(client, ctx, d)
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

// NoClientData is a generic way to define data resources in Terraform provider that doesn't require any client.
// usage is similar to AccountData and WorkspaceData, but the read function doesn't take a client.
func NoClientData[T any](read func(context.Context, *T) error) Resource {
	return genericDatabricksData(
		func(client *DatabricksClient, ctx context.Context, d *schema.ResourceData) (any, error) {
			return nil, nil
		},
		func(ctx context.Context, s struct{}, t *T, ac any) error {
			return read(ctx, t)
		}, false, NoCustomize)
}
