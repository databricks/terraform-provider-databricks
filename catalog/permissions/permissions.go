package permissions

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API
type UnityCatalogPermissionsAPI struct {
	client  *databricks.WorkspaceClient
	context context.Context
}

func NewUnityCatalogPermissionsAPI(ctx context.Context, client *databricks.WorkspaceClient) UnityCatalogPermissionsAPI {
	return UnityCatalogPermissionsAPI{client, ctx}
}

func (a UnityCatalogPermissionsAPI) GetPermissions(securable catalog.SecurableType, name string) (list *catalog.GetPermissionsResponse, err error) {
	if securable.String() == "share" {
		sharePermissions, err := a.client.Shares.SharePermissions(a.context, sharing.SharePermissionsRequest{
			Name: name,
		})
		if err != nil {
			return nil, err
		}
		list = &catalog.GetPermissionsResponse{
			PrivilegeAssignments: make([]catalog.PrivilegeAssignment, len(sharePermissions.PrivilegeAssignments)),
		}
		for i, pa := range sharePermissions.PrivilegeAssignments {
			list.PrivilegeAssignments[i] = catalog.PrivilegeAssignment{
				Principal:  pa.Principal,
				Privileges: toCatalogPrivileges(pa.Privileges),
			}
		}
		return list, nil
	}
	list, err = a.client.Grants.GetBySecurableTypeAndFullName(a.context, securable.String(), name)
	return
}

func (a UnityCatalogPermissionsAPI) UpdatePermissions(securable catalog.SecurableType, name string, diff []catalog.PermissionsChange) error {
	if securable.String() == "share" {
		var shareDiff []sharing.PermissionsChange
		for _, c := range diff {
			shareDiff = append(shareDiff, sharing.PermissionsChange{
				Add:       toSharingPrivileges(c.Add),
				Remove:    toSharingPrivileges(c.Remove),
				Principal: c.Principal,
			})
		}
		_, err := a.client.Shares.UpdatePermissions(a.context, sharing.UpdateSharePermissions{
			Changes: shareDiff,
			Name:    name,
		})
		return err
	}
	_, err := a.client.Grants.Update(a.context, catalog.UpdatePermissions{
		Changes:       diff,
		SecurableType: securable.String(),
		FullName:      name,
	})
	return err
}

func (a UnityCatalogPermissionsAPI) WaitForUpdate(timeout time.Duration, securable catalog.SecurableType, name string, desired []catalog.PrivilegeAssignment, diff func([]catalog.PrivilegeAssignment, []catalog.PrivilegeAssignment) []catalog.PermissionsChange) error {
	return retry.RetryContext(a.context, timeout, func() *retry.RetryError {
		current, err := a.GetPermissions(securable, name)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		log.Printf("[DEBUG] Permissions for %s-%s are: %v", securable.String(), name, current.PrivilegeAssignments)
		if diff(current.PrivilegeAssignments, desired) == nil {
			return nil
		}
		return retry.RetryableError(
			fmt.Errorf("permissions for %s-%s are %v, but have to be %v", securable.String(), name, current.PrivilegeAssignments, desired),
		)
	})
}

// Terraform Schema
type UnityCatalogPrivilegeAssignment struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
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

// securableResourceNamePrefixes maps a grants field to the resource-name prefix
// its databricks_ai_gateway_* resource carries on the `name` attribute (e.g.
// "model-provider-services/{catalog}.{schema}.{name}"). The grants/permissions
// API addresses these securables by their bare full name
// ({catalog}.{schema}.{name}), but users naturally wire
// databricks_ai_gateway_*.<x>.name into databricks_grant(s), so the prefix must
// be stripped before it reaches the API path and the resource ID.
var securableResourceNamePrefixes = map[string]string{
	"mcp_service":            "mcp-services/",
	"model_provider_service": "model-provider-services/",
	"model_service":          "model-services/",
}

// normalizeSecurableName strips the AI Gateway resource-name prefix so a caller
// can pass either the bare full name or the databricks_ai_gateway_*.name
// attribute. A bare name, or any non-AI-Gateway securable, is returned unchanged.
func normalizeSecurableName(field, name string) string {
	if prefix, ok := securableResourceNamePrefixes[field]; ok {
		return strings.TrimPrefix(name, prefix)
	}
	return name
}

// SuppressResourceNamePrefixDiff returns a DiffSuppressFunc that treats an AI
// Gateway securable value as equal whether or not it carries the resource-name
// prefix. Without it, referencing databricks_ai_gateway_*.<x>.name (prefixed) in
// config perpetually diffs against the bare full name that Read stores back,
// forcing replacement on every plan. Returns nil for securables without a
// prefix, so callers can attach it unconditionally.
func SuppressResourceNamePrefixDiff(field string) schema.SchemaDiffSuppressFunc {
	if _, ok := securableResourceNamePrefixes[field]; !ok {
		return nil
	}
	return func(_, old, new string, _ *schema.ResourceData) bool {
		return normalizeSecurableName(field, old) == normalizeSecurableName(field, new)
	}
}

func (sm SecurableMapping) KeyValue(d attributeGetter) (string, string) {
	for field := range sm {
		v := d.Get(field).(string)
		if v == "" {
			continue
		}
		return field, normalizeSecurableName(field, v)
	}
	log.Printf("[WARN] Unexpected resource or permissions. Please proceed at your own risk.")
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
	"catalog":                catalog.SecurableType("catalog"),
	"credential":             catalog.SecurableType("credential"),
	"foreign_connection":     catalog.SecurableType("connection"),
	"external_location":      catalog.SecurableType("external_location"),
	"function":               catalog.SecurableType("function"),
	"mcp_service":            catalog.SecurableType("mcp_service"),
	"metastore":              catalog.SecurableType("metastore"),
	"model":                  catalog.SecurableType("function"),
	"model_provider_service": catalog.SecurableType("model_provider_service"),
	"model_service":          catalog.SecurableType("model_service"),
	"pipeline":               catalog.SecurableType("pipeline"),
	"recipient":              catalog.SecurableType("recipient"),
	"schema":                 catalog.SecurableType("schema"),
	"secret":                 catalog.SecurableType("secret"),
	"share":                  catalog.SecurableType("share"),
	"storage_credential":     catalog.SecurableType("storage_credential"),
	"table":                  catalog.SecurableType("table"),
	"volume":                 catalog.SecurableType("volume"),
}

// Unity Catalog accepts privileges with spaces, but will automatically convert them to underscores
func NormalizePrivilege(privilege string) string {
	return strings.ToUpper(strings.Replace(privilege, " ", "_", -1))
}

// Utils for Slice and Set
func SliceToSet(in []catalog.Privilege) *schema.Set {
	var out []any
	for _, v := range in {
		out = append(out, NormalizePrivilege(v.String()))
	}
	return schema.NewSet(schema.HashString, out)
}

func SetToSlice(set *schema.Set) (ss []catalog.Privilege) {
	for _, v := range set.List() {
		ss = append(ss, catalog.Privilege(NormalizePrivilege(v.(string))))
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

func toCatalogPrivileges(in []sharing.Privilege) []catalog.Privilege {
	var out = make([]catalog.Privilege, len(in))
	for i, p := range in {
		out[i] = catalog.Privilege(p)
	}
	return out
}

func toSharingPrivileges(in []catalog.Privilege) []string {
	var out = make([]string, len(in))
	for i, p := range in {
		out[i] = p.String()
	}
	return out
}
