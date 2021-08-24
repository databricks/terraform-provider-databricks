package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"

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

// CreateR ..
func (a ServicePrincipalsAPI) Create(rsp ScimUser) (sp ScimUser, err error) {
	if rsp.Schemas == nil {
		rsp.Schemas = []URN{ServicePrincipalSchema}
	}
	err = a.client.Scim(a.context, "POST", "/preview/scim/v2/ServicePrincipals", rsp, &sp)
	return sp, err
}

func (a ServicePrincipalsAPI) read(servicePrincipalID string) (sp ScimUser, err error) {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	err = a.client.Scim(a.context, "GET", servicePrincipalPath, nil, &sp)
	return
}

// Update replaces resource-friendly-entity
func (a ServicePrincipalsAPI) Update(servicePrincipalID string, updateRequest ScimUser) error {
	servicePrincipal, err := a.read(servicePrincipalID)
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

// Delete will delete the servicePrincipal given the servicePrincipal id
func (a ServicePrincipalsAPI) Delete(servicePrincipalID string) error {
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v", servicePrincipalID)
	return a.client.Scim(a.context, "DELETE", servicePrincipalPath, nil, nil)
}

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() *schema.Resource {
	type entity struct {
		ApplicationID string `json:"application_id,omitempty" tf:"computed,force_new"`
		DisplayName   string `json:"display_name,omitempty" tf:"computed"`
		Active        bool   `json:"active,omitempty"`
	}
	servicePrincipalSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(&m)
			m["active"].Default = true
			return m
		})
	spFromData := func(d *schema.ResourceData) (user ScimUser, err error) {
		var u entity
		if err = common.DataToStructPointer(d, servicePrincipalSchema, &u); err != nil {
			return
		}
		return ScimUser{
			ApplicationID: u.ApplicationID,
			DisplayName:   u.DisplayName,
			Active:        u.Active,
			Entitlements:  readEntitlementsFromData(d),
		}, nil
	}
	return common.Resource{
		Schema: servicePrincipalSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c interface{}) error {
			var sp entity
			if err := common.DiffToStructPointer(d, servicePrincipalSchema, &sp); err != nil {
				return err
			}
			client := c.(*common.DatabricksClient)
			if client.IsAzure() && sp.ApplicationID == "" {
				return fmt.Errorf("application_id is required for service principals in Azure Databricks")
			}
			if client.IsAws() && sp.DisplayName == "" {
				return fmt.Errorf("display_name is required for service principals in Databricks on AWS")
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sp, err := spFromData(d)
			if err != nil {
				return err
			}
			if c.IsAws() && sp.ApplicationID != "" {
				return fmt.Errorf("application_id is not allowed for service principals in Databricks on AWS")
			}
			servicePrincipal, err := NewServicePrincipalsAPI(ctx, c).Create(sp)
			if err != nil {
				return err
			}
			d.SetId(servicePrincipal.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sp, err := NewServicePrincipalsAPI(ctx, c).read(d.Id())
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
			sp, err := spFromData(d)
			if err != nil {
				return err
			}
			return NewServicePrincipalsAPI(ctx, c).Update(d.Id(), sp)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewServicePrincipalsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
