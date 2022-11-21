package scim

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewServicePrincipalsAPI creates ServicePrincipalsAPI instance from provider meta
func NewServicePrincipalsAPI(ctx context.Context, m any) ServicePrincipalsAPI {
	return ServicePrincipalsAPI{m.(*common.DatabricksClient), ctx}
}

// ServicePrincipalsAPI exposes the scim servicePrincipal API
type ServicePrincipalsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// CreateR ..
func (a ServicePrincipalsAPI) Create(rsp User) (sp User, err error) {
	if rsp.Schemas == nil {
		rsp.Schemas = []URN{ServicePrincipalSchema}
	}
	err = a.client.Scim(a.context, "POST", "/preview/scim/v2/ServicePrincipals", rsp, &sp)
	return sp, err
}

func (a ServicePrincipalsAPI) Read(servicePrincipalID string) (sp User, err error) {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	err = a.client.Scim(a.context, "GET", servicePrincipalPath, nil, &sp)
	return
}

func (a ServicePrincipalsAPI) filter(filter string) (u []User, err error) {
	var sps UserList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	err = a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/ServicePrincipals", req, &sps)
	if err != nil {
		return
	}
	u = sps.Resources
	return
}

// Patch updates resource-friendly entity
func (a ServicePrincipalsAPI) Patch(servicePrincipalID string, r patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch, fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID), r, nil)
}

// Update replaces resource-friendly-entity
func (a ServicePrincipalsAPI) Update(servicePrincipalID string, updateRequest User) error {
	servicePrincipal, err := a.Read(servicePrincipalID)
	if err != nil {
		return err
	}
	if updateRequest.Schemas == nil {
		updateRequest.Schemas = []URN{ServicePrincipalSchema}
	}
	updateRequest.Groups = servicePrincipal.Groups
	return a.client.Scim(a.context, "PUT",
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID),
		updateRequest, nil)
}

func (a ServicePrincipalsAPI) UpdateEntitlements(servicePrincipalID string, entitlements patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch,
		fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID), entitlements, nil)
}

// Delete will delete the servicePrincipal given the servicePrincipal id
func (a ServicePrincipalsAPI) Delete(servicePrincipalID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.client.Scim(a.context, "DELETE", servicePrincipalPath, nil, nil)
}

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() *schema.Resource {
	type entity struct {
		ApplicationID string `json:"application_id,omitempty" tf:"computed,force_new"`
		DisplayName   string `json:"display_name,omitempty" tf:"computed,force_new"`
		Active        bool   `json:"active,omitempty"`
		ExternalID    string `json:"external_id,omitempty" tf:"suppress_diff"`
	}
	servicePrincipalSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(&m)
			m["active"].Default = true
			m["force"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			return m
		})
	spFromData := func(d *schema.ResourceData) User {
		var u entity
		common.DataToStructPointer(d, servicePrincipalSchema, &u)
		return User{
			ApplicationID: u.ApplicationID,
			DisplayName:   u.DisplayName,
			Active:        u.Active,
			Entitlements:  readEntitlementsFromData(d),
			ExternalID:    u.ExternalID,
		}
	}
	return common.Resource{
		Schema: servicePrincipalSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sp := spFromData(d)
			spAPI := NewServicePrincipalsAPI(ctx, c)
			servicePrincipal, err := spAPI.Create(sp)
			if err != nil {
				return createForceOverridesManuallyAddedServicePrincipal(err, d, spAPI, sp)
			}
			d.SetId(servicePrincipal.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sp, err := NewServicePrincipalsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(sp, servicePrincipalSchema, d)
			if err != nil {
				return err
			}
			return sp.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var applicationId string
			if c.IsAzure() {
				applicationId = d.Get("application_id").(string)
			}
			return NewServicePrincipalsAPI(ctx, c).Update(d.Id(), User{
				DisplayName:   d.Get("display_name").(string),
				Active:        d.Get("active").(bool),
				Entitlements:  readEntitlementsFromData(d),
				ExternalID:    d.Get("external_id").(string),
				ApplicationID: applicationId,
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewServicePrincipalsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}

func createForceOverridesManuallyAddedServicePrincipal(err error, d *schema.ResourceData, spAPI ServicePrincipalsAPI, u User) error {
	forceCreate := d.Get("force").(bool)
	if !forceCreate {
		return err
	}
	// corner-case for overriding manually provisioned service principals
	force := fmt.Sprintf("Service principal with application ID %s already exists.", u.ApplicationID)
	if err.Error() != force {
		return err
	}
	spList, err := spAPI.filter(fmt.Sprintf("applicationId eq '%s'", strings.ReplaceAll(u.ApplicationID, "'", "")))
	if err != nil {
		return err
	}
	if len(spList) == 0 {
		return fmt.Errorf("cannot find SP with ID %s for force import", u.ApplicationID)
	}
	sp := spList[0]
	d.SetId(sp.ID)
	return spAPI.Update(d.Id(), u)
}
