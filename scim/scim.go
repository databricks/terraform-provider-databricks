package scim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// URN is a custom type for the SCIM spec for the schema
type URN string

// Possible schema URNs for the Databricks SCIM api
const (
	UserSchema             URN = "urn:ietf:params:scim:schemas:core:2.0:User"
	ServicePrincipalSchema URN = "urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal"
	WorkspaceUserSchema    URN = "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"
	PatchOp                URN = "urn:ietf:params:scim:api:messages:2.0:PatchOp"
	GroupSchema            URN = "urn:ietf:params:scim:schemas:core:2.0:Group"
)

// Generalisation of most common complex values from SCIM protocol
// Details at https://datatracker.ietf.org/doc/html/rfc7643#section-2.3.8
type ComplexValue struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`

	// https://tools.ietf.org/html/rfc7643#page-64
	Type string `json:"type,omitempty"`
}

type ComplexValues []ComplexValue

func (cv ComplexValues) HasValue(value string) bool {
	for _, v := range cv {
		if v.Value == value {
			return true
		}
	}
	return false
}

var entitlementMapping = map[string]string{
	"allow-cluster-create":       "allow_cluster_create",
	"allow-instance-pool-create": "allow_instance_pool_create",
	"databricks-sql-access":      "databricks_sql_access",
	"workspace-access":           "workspace_access",
	"workspace-consume":          "workspace_consume",
}

// order is important for tests
var possibleEntitlements = []string{
	"allow-cluster-create",
	"allow-instance-pool-create",
	"databricks-sql-access",
	"workspace-access",
	"workspace-consume",
}

type entitlements []ComplexValue

func (e entitlements) generateEmpty(d *schema.ResourceData) error {
	for _, entitlement := range possibleEntitlements {
		d.Set(entitlementMapping[entitlement], false)
	}
	return nil
}

func (e entitlements) readIntoData(d *schema.ResourceData) error {
	for _, ent := range e {
		field_name := entitlementMapping[ent.Value]
		if err := d.Set(field_name, true); err != nil {
			return err
		}
	}
	return nil
}

func readEntitlementsFromData(d *schema.ResourceData) entitlements {
	var e entitlements
	for _, entitlement := range possibleEntitlements {
		field_name := entitlementMapping[entitlement]
		if d.Get(field_name).(bool) {
			e = append(e, ComplexValue{
				Value: entitlement,
			})
		}
	}
	// if there is no nil value
	if e == nil {
		e = append(e, ComplexValue{
			Value: "",
		})
	}
	return e
}

func addEntitlementsToSchema(s map[string]*schema.Schema) {
	for _, entitlement := range possibleEntitlements {
		field_name := entitlementMapping[entitlement]
		s[field_name] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		}
	}
}

// ResourceMeta is a struct that contains the meta information about the SCIM group
type ResourceMeta struct {
	// ResourceType is the type of the resource: "Group" or "WorkspaceGroup"
	ResourceType string `json:"resourceType,omitempty"`
}

// Group contains information about the SCIM group
type Group struct {
	ID           string         `json:"id,omitempty"`
	Schemas      []URN          `json:"schemas,omitempty"`
	DisplayName  string         `json:"displayName,omitempty"`
	Members      []ComplexValue `json:"members,omitempty"`
	Groups       []ComplexValue `json:"groups,omitempty"`
	Roles        []ComplexValue `json:"roles,omitempty"`
	Entitlements entitlements   `json:"entitlements,omitempty"`
	ExternalID   string         `json:"externalId,omitempty"`
	Meta         *ResourceMeta  `json:"meta,omitempty" tf:"computed"`
}

// GroupList contains a list of groups fetched from a list api call from SCIM api
type GroupList struct {
	TotalResults int32   `json:"totalResults,omitempty"`
	StartIndex   int32   `json:"startIndex,omitempty"`
	ItemsPerPage int32   `json:"itemsPerPage,omitempty"`
	Schemas      []URN   `json:"schemas,omitempty"`
	Resources    []Group `json:"resources,omitempty"`
}

type email struct {
	Type    any    `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Primary any    `json:"primary,omitempty"`
}

// User is a struct that contains all the information about a SCIM user
type User struct {
	ID            string            `json:"id,omitempty"`
	Emails        []email           `json:"emails,omitempty"`
	DisplayName   string            `json:"displayName,omitempty" tf:"alias:display_name"`
	Active        bool              `json:"active"`
	Schemas       []URN             `json:"schemas,omitempty"`
	UserName      string            `json:"userName,omitempty" tf:"alias:user_name"`
	ApplicationID string            `json:"applicationId,omitempty" tf:"alias:application_id"`
	Groups        []ComplexValue    `json:"groups,omitempty"`
	Name          map[string]string `json:"name,omitempty"`
	Roles         []ComplexValue    `json:"roles,omitempty"`
	Entitlements  entitlements      `json:"entitlements,omitempty"`
	ExternalID    string            `json:"externalId,omitempty"`
}

// UserList contains a list of Users fetched from a list api call from SCIM api
type UserList struct {
	TotalResults int32  `json:"totalResults,omitempty"`
	StartIndex   int32  `json:"startIndex,omitempty"`
	ItemsPerPage int32  `json:"itemsPerPage,omitempty"`
	Schemas      []URN  `json:"schemas,omitempty"`
	Resources    []User `json:"resources,omitempty"`
}

type patchOperation struct {
	Op    string `json:"op,omitempty"`
	Path  string `json:"path,omitempty"`
	Value any    `json:"value,omitempty"`
}

type patchRequest struct {
	Schemas    []URN            `json:"schemas,omitempty"`
	Operations []patchOperation `json:"Operations,omitempty"`
}

func PatchRequest(op, path string) patchRequest {
	o := patchOperation{
		Op:   op,
		Path: path,
	}
	return PatchRequestComplexValue([]patchOperation{o})
}

func PatchRequestWithValue(op, path, value string) patchRequest {
	o := patchOperation{
		Op:    op,
		Path:  path,
		Value: []ComplexValue{{Value: value}},
	}
	return PatchRequestComplexValue([]patchOperation{o})
}

func PatchRequestComplexValue(operations []patchOperation) patchRequest {
	return patchRequest{
		Schemas:    []URN{PatchOp},
		Operations: operations,
	}
}
