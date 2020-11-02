package identity

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
)

// NewServicePrincipalsAPI creates ServicePrincipalsAPI instance from provider meta
func NewServicePrincipalsAPI(m interface{}) ServicePrincipalsAPI {
	return ServicePrincipalsAPI{C: m.(*common.DatabricksClient)}
}

// ServicePrincipalsAPI exposes the scim servicePrincipal API
type ServicePrincipalsAPI struct {
	C *common.DatabricksClient
}

// ServicePrincipalEntity entity from which resource schema is made
type ServicePrincipalEntity struct {
	ApplicationId           string `json:"application_id"`
	DisplayName             string `json:"display_name,omitempty" tf:"computed"`
	Active                  bool   `json:"active,omitempty"`
	AllowClusterCreate      bool   `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool   `json:"allow_instance_pool_create,omitempty"`
}

func (sp ServicePrincipalEntity) toRequest() ScimServicePrincipal {
	entitlements := []EntitlementsListItem{}
	if sp.AllowClusterCreate {
		entitlements = append(entitlements, EntitlementsListItem{
			Value: Entitlement("allow-cluster-create"),
		})
	}
	if sp.AllowInstancePoolCreate {
		entitlements = append(entitlements, EntitlementsListItem{
			Value: Entitlement("allow-instance-pool-create"),
		})
	}
	return ScimServicePrincipal{
		Schemas:       []URN{ServicePrincipalSchema},
		ApplicationId: sp.ApplicationId,
		Active:        sp.Active,
		DisplayName:   sp.DisplayName,
		Entitlements:  entitlements,
	}
}

// CreateR ..
func (a ServicePrincipalsAPI) CreateR(rsp ServicePrincipalEntity) (servicePrincipal ScimServicePrincipal, err error) {
	err = a.C.Scim(http.MethodPost, "/preview/scim/v2/ServicePrincipals", rsp.toRequest(), &servicePrincipal)
	return servicePrincipal, err
}

// Create given a applicationId, displayname, entitlements, and roles will create a scim servicePrincipal via SCIM api
func (a ServicePrincipalsAPI) Create(applicationId string, displayName string, entitlements []string, roles []string) (ScimServicePrincipal, error) {
	var servicePrincipal ScimServicePrincipal
	createRequest := ScimServicePrincipal{
		Schemas:       []URN{ServicePrincipalSchema},
		ApplicationId: applicationId,
		DisplayName:   displayName,
		Entitlements:  []EntitlementsListItem{},
		Roles:         []RoleListItem{},
	}
	for _, entitlement := range entitlements {
		createRequest.Entitlements = append(createRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	for _, role := range roles {
		createRequest.Roles = append(createRequest.Roles, RoleListItem{Value: role})
	}
	err := a.C.Scim(http.MethodPost, "/preview/scim/v2/ServicePrincipals", createRequest, &servicePrincipal)
	return servicePrincipal, err
}

// ReadR reads resource-friendly entity
func (a ServicePrincipalsAPI) ReadR(servicePrincipalID string) (rsp ServicePrincipalEntity, err error) {
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return
	}
	rsp.ApplicationId = servicePrincipal.ApplicationId
	rsp.DisplayName = servicePrincipal.DisplayName
	rsp.Active = servicePrincipal.Active
	for _, ent := range servicePrincipal.Entitlements {
		switch ent.Value {
		case AllowClusterCreateEntitlement:
			rsp.AllowClusterCreate = true
		case AllowInstancePoolCreateEntitlement:
			rsp.AllowInstancePoolCreate = true
		}
	}
	return
}

// Read returns the servicePrincipal object and all the attributes of a scim servicePrincipal
func (a ServicePrincipalsAPI) Read(servicePrincipalID string) (ScimServicePrincipal, error) {
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return servicePrincipal, err
	}

	//get groups
	var groups []ScimGroup
	for _, group := range servicePrincipal.Groups {
		group, err := GroupsAPI{a.C}.Read(group.Value)
		if err != nil {
			return servicePrincipal, err
		}
		groups = append(groups, group)
	}
	inherited, unInherited := a.getInheritedAndNonInheritedRoles(servicePrincipal, groups)
	servicePrincipal.InheritedRoles = inherited
	servicePrincipal.UnInheritedRoles = unInherited
	return servicePrincipal, err
}

func (a ServicePrincipalsAPI) read(servicePrincipalID string) (ScimServicePrincipal, error) {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.readByPath(servicePrincipalPath)
}

// Me gets servicePrincipal information about caller
func (a ServicePrincipalsAPI) Me() (ScimServicePrincipal, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a ServicePrincipalsAPI) readByPath(servicePrincipalPath string) (servicePrincipal ScimServicePrincipal, err error) {
	err = a.C.Scim(http.MethodGet, servicePrincipalPath, nil, &servicePrincipal)
	return
}

// UpdateR replaces resource-friendly-entity
func (a ServicePrincipalsAPI) UpdateR(servicePrincipalID string, rsp ServicePrincipalEntity) error {
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return err
	}
	updateRequest := rsp.toRequest()
	updateRequest.Groups = servicePrincipal.Groups
	updateRequest.Roles = servicePrincipal.Roles
	return a.C.Scim(http.MethodPut,
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID),
		updateRequest, nil)
}

// PatchR updates resource-friendly entity
func (a ServicePrincipalsAPI) PatchR(servicePrincipalID string, r patchRequest) error {
	return a.C.Scim(http.MethodPatch, fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID), r, nil)
}

// Update will update the servicePrincipal given the servicePrincipal id, application id, display name, entitlements and roles
func (a ServicePrincipalsAPI) Update(servicePrincipalID string, applicationId string, displayName string, entitlements []string, roles []string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	scimServicePrincipalUpdateRequest := struct {
		Schemas       []URN                  `json:"schemas,omitempty"`
		ApplicationId string                 `json:"applicationId,omitempty"`
		Entitlements  []EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName   string                 `json:"displayName,omitempty"`
		Roles         []RoleListItem         `json:"roles,omitempty"`
		Groups        []GroupsListItem       `json:"groups,omitempty"`
	}{}
	scimServicePrincipalUpdateRequest.Schemas = []URN{ServicePrincipalSchema}
	scimServicePrincipalUpdateRequest.ApplicationId = applicationId
	scimServicePrincipalUpdateRequest.DisplayName = displayName
	scimServicePrincipalUpdateRequest.Entitlements = []EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimServicePrincipalUpdateRequest.Entitlements = append(scimServicePrincipalUpdateRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	scimServicePrincipalUpdateRequest.Roles = []RoleListItem{}
	for _, role := range roles {
		scimServicePrincipalUpdateRequest.Roles = append(scimServicePrincipalUpdateRequest.Roles, RoleListItem{Value: role})
	}
	//Get any existing groups that the servicePrincipal is part of
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return err
	}
	scimServicePrincipalUpdateRequest.Groups = servicePrincipal.Groups
	return a.C.Scim(http.MethodPut, servicePrincipalPath, scimServicePrincipalUpdateRequest, nil)
}

// Delete will delete the servicePrincipal given the servicePrincipal id
func (a ServicePrincipalsAPI) Delete(servicePrincipalID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.C.Scim(http.MethodDelete, servicePrincipalPath, nil, nil)
}

// SetServicePrincipalAsAdmin will add the servicePrincipal to a admin group given the admin group id and servicePrincipal id
func (a ServicePrincipalsAPI) SetServicePrincipalAsAdmin(servicePrincipalID string, adminGroupID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	var addOperations UserPatchOperations
	servicePrincipalPatchRequest := UserPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []UserPatchOperations{},
	}
	addOperations = UserPatchOperations{
		Op: "add",
		Value: &GroupsValue{
			Groups: []ValueListItem{{Value: adminGroupID}},
		},
	}
	servicePrincipalPatchRequest.Operations = append(servicePrincipalPatchRequest.Operations, addOperations)
	return a.C.Scim(http.MethodPatch, servicePrincipalPath, servicePrincipalPatchRequest, nil)
}

// VerifyServicePrincipalAsAdmin will verify the servicePrincipal belongs to the admin group given the admin group id and servicePrincipal id
func (a ServicePrincipalsAPI) VerifyServicePrincipalAsAdmin(servicePrincipalID string, adminGroupID string) (bool, error) {
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return false, err
	}
	for _, group := range servicePrincipal.Groups {
		if group.Value == adminGroupID {
			return true, nil
		}
	}
	return false, nil
}

// RemoveServicePrincipalAsAdmin will remove the servicePrincipal from the admin group given the admin group id and servicePrincipal id
func (a ServicePrincipalsAPI) RemoveServicePrincipalAsAdmin(servicePrincipalID string, adminGroupID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	var removeOperations UserPatchOperations
	servicePrincipalPatchRequest := UserPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []UserPatchOperations{},
	}
	path := fmt.Sprintf("groups[value eq \"%s\"]", adminGroupID)
	removeOperations = UserPatchOperations{
		Op:   "remove",
		Path: path,
	}
	servicePrincipalPatchRequest.Operations = append(servicePrincipalPatchRequest.Operations, removeOperations)
	return a.C.Scim(http.MethodPatch, servicePrincipalPath, servicePrincipalPatchRequest, nil)
}

// GetOrCreateDefaultMetaServicePrincipal ...
func (a ServicePrincipalsAPI) GetOrCreateDefaultMetaServicePrincipal(metaServicePrincipalDisplayName string, metaServicePrincipalName string, deleteAfterCreate bool) (servicePrincipal ScimServicePrincipal, err error) {
	var servicePrincipals ServicePrincipalList
	err = a.C.Scim(http.MethodGet, "/preview/scim/v2/ServicePrincipals", map[string]string{
		"filter": "displayName+eq+" + metaServicePrincipalDisplayName,
	}, &servicePrincipals)
	if err != nil {
		return servicePrincipal, err
	}
	resources := servicePrincipals.Resources
	if len(resources) == 1 {
		return resources[0], err
	} else if len(resources) > 1 {
		return ScimServicePrincipal{}, errors.New("more than one meta servicePrincipal")
	}

	log.Printf("Meta ServicePrincipal not found will create a new meta servicePrincipal with name: %s\n", metaServicePrincipalDisplayName)

	newCreatedServicePrincipal, err := a.Create(metaServicePrincipalName, metaServicePrincipalDisplayName, nil, nil)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			time.Sleep(time.Second * 1)
			return a.GetOrCreateDefaultMetaServicePrincipal(metaServicePrincipalDisplayName, metaServicePrincipalName, deleteAfterCreate)
		}
		return servicePrincipal, err
	}
	if deleteAfterCreate {
		defer func() {
			deferErr := a.Delete(newCreatedServicePrincipal.ID)
			err = deferErr
		}()
	}
	return newCreatedServicePrincipal, err
}

func (a ServicePrincipalsAPI) getInheritedAndNonInheritedRoles(servicePrincipal ScimServicePrincipal, groups []ScimGroup) (inherited []RoleListItem, unInherited []RoleListItem) {
	allRoles := servicePrincipal.Roles
	var inheritedRoles []RoleListItem
	inheritedRolesKeys := []string{}
	inheritedRolesMap := map[string]RoleListItem{}
	for _, group := range groups {
		inheritedRoles = append(inheritedRoles, group.Roles...)
	}
	for _, role := range inheritedRoles {
		inheritedRolesKeys = append(inheritedRolesKeys, role.Value)
		inheritedRolesMap[role.Value] = role
	}
	sort.Strings(inheritedRolesKeys)
	for _, roleKey := range inheritedRolesKeys {
		inherited = append(inherited, inheritedRolesMap[roleKey])
	}
	for _, role := range allRoles {
		if _, ok := inheritedRolesMap[role.Value]; !ok {
			unInherited = append(unInherited, role)
		}
	}
	return inherited, unInherited
}
