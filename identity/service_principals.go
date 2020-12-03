package identity

import (
	"fmt"
	"net/http"

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
	}
	for _, entitlement := range entitlements {
		createRequest.Entitlements = append(createRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
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
	return servicePrincipal, err
}

func (a ServicePrincipalsAPI) read(servicePrincipalID string) (ScimServicePrincipal, error) {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.readByPath(servicePrincipalPath)
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
		Groups        []GroupsListItem       `json:"groups,omitempty"`
	}{}
	scimServicePrincipalUpdateRequest.Schemas = []URN{ServicePrincipalSchema}
	scimServicePrincipalUpdateRequest.ApplicationId = applicationId
	scimServicePrincipalUpdateRequest.DisplayName = displayName
	scimServicePrincipalUpdateRequest.Entitlements = []EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimServicePrincipalUpdateRequest.Entitlements = append(scimServicePrincipalUpdateRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
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
