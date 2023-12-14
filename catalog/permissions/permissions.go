package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API
func NewUnityCatalogPermissionsAPI(ctx context.Context, m any) UnityCatalogPermissionsAPI {
	client, _ := m.(*common.DatabricksClient).WorkspaceClient()
	return UnityCatalogPermissionsAPI{client, ctx}
}

func (a UnityCatalogPermissionsAPI) GetPermissions(securable string, name string) (list *catalog.PermissionsList, err error) {
	if securable == "share" {
		list, err = a.client.Shares.SharePermissions(a.context, sharing.SharePermissionsRequest{name})
		return
	}
	list, err = a.client.Grants.GetBySecurableTypeAndFullName(a.context, Mappings.GetSecurableType(securable), name)
	return
}

func (a UnityCatalogPermissionsAPI) UpdatePermissions(securable, name string, diff []catalog.PermissionsChange) error {
	if securable == "share" {
		return a.client.Shares.UpdatePermissions(a.context, sharing.UpdateSharePermissions{
			Changes: diff,
			Name:    name,
		})
	}
	_, err := a.client.Grants.Update(a.context, catalog.UpdatePermissions{
		Changes:       diff,
		SecurableType: Mappings.GetSecurableType(securable),
		FullName:      name,
	})
	return err
}

// Terraform Schema
type UnityCatalogPrivilegeAssignment struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
}

type UnityCatalogPermissionsAPI struct {
	client  *databricks.WorkspaceClient
	context context.Context
}

// Permission Mappings

type SecurableMapping map[string]catalog.SecurableType

// reuse ResourceDiff and ResourceData
type attributeGetter interface {
	Get(key string) any
}

func (sm SecurableMapping) GetSecurableType(securable string) catalog.SecurableType {
	return sm[securable]
}

func (sm SecurableMapping) KeyValue(d attributeGetter) (string, string) {
	for field := range sm {
		v := d.Get(field).(string)
		if v == "" {
			continue
		}
		return field, v
	}
	return "unknown", "unknown"
}
func (sm SecurableMapping) Id(d *schema.ResourceData) string {
	securable, name := sm.KeyValue(d)
	return fmt.Sprintf("%s/%s", securable, name)
}

// Mappings
// See https://docs.databricks.com/api/workspace/grants/update for full list
// Omitting provider as a reserved keyword
var Mappings = SecurableMapping{
	"catalog":            catalog.SecurableType("catalog"),
	"foreign_connection": catalog.SecurableType("connection"),
	"external_location":  catalog.SecurableType("external_location"),
	"function":           catalog.SecurableType("function"),
	"metastore":          catalog.SecurableType("metastore"),
	"pipeline":           catalog.SecurableType("pipeline"),
	"recipient":          catalog.SecurableType("recipient"),
	"schema":             catalog.SecurableType("schema"),
	"share":              catalog.SecurableType("share"),
	"storage_credential": catalog.SecurableType("storage_credential"),
	"table":              catalog.SecurableType("table"),
	"volume":             catalog.SecurableType("volume"),
}

// Utils for Slice and Set
func SliceToSet(in []catalog.Privilege) *schema.Set {
	var out []any
	for _, v := range in {
		out = append(out, v.String())
	}
	return schema.NewSet(schema.HashString, out)
}

func SetToSlice(set *schema.Set) (ss []catalog.Privilege) {
	for _, v := range set.List() {
		ss = append(ss, catalog.Privilege(v.(string)))
	}
	return
}

func ToPrivilegeSlice(in []interface{}) (out []catalog.Privilege) {
	for _, v := range in {
		out = append(out, catalog.Privilege(v.(string)))
	}
	return
}

func SliceWithoutString(in []string, without string) (out []string) {
	for _, v := range in {
		if v == without {
			continue
		}
		out = append(out, v)
	}
	return
}
