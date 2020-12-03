package identity

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewServicePrincipalsAPI creates ServicePrincipalsAPI instance from provider meta
func NewServicePrincipalsAPI(m interface{}) ServicePrincipalsAPI {
	return ServicePrincipalsAPI{client: m.(*common.DatabricksClient)}
}

// ServicePrincipalsAPI exposes the scim servicePrincipal API
type ServicePrincipalsAPI struct {
	client *common.DatabricksClient
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
	err = a.client.Scim(http.MethodPost, "/preview/scim/v2/ServicePrincipals", rsp.toRequest(), &sp)
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
	err = a.client.Scim(http.MethodGet, servicePrincipalPath, nil, &sp)
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
	return a.client.Scim(http.MethodPut,
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID),
		updateRequest, nil)
}

// PatchR updates resource-friendly entity
func (a ServicePrincipalsAPI) PatchR(servicePrincipalID string, r patchRequest) error {
	return a.client.Scim(http.MethodPatch, fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID), r, nil)
}

// Delete will delete the servicePrincipal given the servicePrincipal id
func (a ServicePrincipalsAPI) Delete(servicePrincipalID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.client.Scim(http.MethodDelete, servicePrincipalPath, nil, nil)
}

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() *schema.Resource {
	servicePrincipalSchema := internal.StructToSchema(ServicePrincipalEntity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["application_id"].ForceNew = true
		s["active"].Default = true
		return s
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		servicePrincipal, err := NewServicePrincipalsAPI(m).ReadR(d.Id())
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(servicePrincipal, servicePrincipalSchema, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:      servicePrincipalSchema,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru ServicePrincipalEntity
			err := internal.DataToStructPointer(d, servicePrincipalSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			servicePrincipal, err := NewServicePrincipalsAPI(m).CreateR(ru)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(servicePrincipal.ID)
			return readContext(ctx, d, m)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru ServicePrincipalEntity
			err := internal.DataToStructPointer(d, servicePrincipalSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewServicePrincipalsAPI(m).UpdateR(d.Id(), ru)
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewServicePrincipalsAPI(m).Delete(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
