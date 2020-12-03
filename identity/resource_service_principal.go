package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewServicePrincipalsAPI creates ServicePrincipalsAPI instance from provider meta
func NewServicePrincipalsAPI(ctx context.Context, m interface{}) ServicePrincipalsAPI {
	return ServicePrincipalsAPI{m.(*common.DatabricksClient), ctx}
}

// ServicePrincipalsAPI exposes the scim servicePrincipal API
type ServicePrincipalsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// ServicePrincipalEntity entity from which resource schema is made
type ServicePrincipalEntity struct {
	ApplicationID           string `json:"application_id"`
	DisplayName             string `json:"display_name,omitempty" tf:"computed"`
	Active                  bool   `json:"active,omitempty"`
	AllowClusterCreate      bool   `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool   `json:"allow_instance_pool_create,omitempty"`
}

func (sp ServicePrincipalEntity) toRequest() ScimUser {
	entitlements := []entitlementsListItem{}
	if sp.AllowClusterCreate {
		entitlements = append(entitlements, entitlementsListItem{
			Value: Entitlement("allow-cluster-create"),
		})
	}
	if sp.AllowInstancePoolCreate {
		entitlements = append(entitlements, entitlementsListItem{
			Value: Entitlement("allow-instance-pool-create"),
		})
	}
	return ScimUser{
		Schemas:       []URN{ServicePrincipalSchema},
		ApplicationID: sp.ApplicationID,
		Active:        sp.Active,
		DisplayName:   sp.DisplayName,
		Entitlements:  entitlements,
	}
}

// CreateR ..
func (a ServicePrincipalsAPI) CreateR(rsp ServicePrincipalEntity) (sp ScimUser, err error) {
	err = a.client.Scim(a.context, "POST", "/preview/scim/v2/ServicePrincipals", rsp.toRequest(), &sp)
	return sp, err
}

// ReadR reads resource-friendly entity
func (a ServicePrincipalsAPI) ReadR(servicePrincipalID string) (rsp ServicePrincipalEntity, err error) {
	servicePrincipal, err := a.read(servicePrincipalID)
	if err != nil {
		return
	}
	rsp.ApplicationID = servicePrincipal.ApplicationID
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

func (a ServicePrincipalsAPI) read(servicePrincipalID string) (sp ScimUser, err error) {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	err = a.client.Scim(a.context, "GET", servicePrincipalPath, nil, &sp)
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
	return a.client.Scim(a.context, "PUT",
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID),
		updateRequest, nil)
}

// PatchR updates resource-friendly entity
func (a ServicePrincipalsAPI) PatchR(servicePrincipalID string, r patchRequest) error {
	return a.client.Scim(a.context, "PATCH",
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v",
			servicePrincipalID), r, nil)
}

// Delete will delete the servicePrincipal given the servicePrincipal id
func (a ServicePrincipalsAPI) Delete(servicePrincipalID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.client.Scim(a.context, "DELETE", servicePrincipalPath, nil, nil)
}

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() *schema.Resource {
	servicePrincipalSchema := internal.StructToSchema(ServicePrincipalEntity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["application_id"].ForceNew = true
		s["active"].Default = true
		return s
	})
	return util.CommonResource{
		Schema: servicePrincipalSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var sp ServicePrincipalEntity
			if err := internal.DataToStructPointer(d, servicePrincipalSchema, &sp); err != nil {
				return err
			}
			servicePrincipal, err := NewServicePrincipalsAPI(ctx, c).CreateR(sp)
			if err != nil {
				return err
			}
			d.SetId(servicePrincipal.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			servicePrincipal, err := NewServicePrincipalsAPI(ctx, c).ReadR(d.Id())
			if err != nil {
				return err
			}
			return internal.StructToData(servicePrincipal, servicePrincipalSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var sp ServicePrincipalEntity
			if err := internal.DataToStructPointer(d, servicePrincipalSchema, &sp); err != nil {
				return err
			}
			return NewServicePrincipalsAPI(ctx, c).UpdateR(d.Id(), sp)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewServicePrincipalsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
