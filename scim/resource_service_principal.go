package scim

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"golang.org/x/exp/slices"

	"github.com/databricks/terraform-provider-databricks/workspace"
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

func (a ServicePrincipalsAPI) Read(servicePrincipalID string, attributes string) (sp User, err error) {
	attrs := ""
	if attributes != "" {
		attrs = "?attributes=" + attributes
	}
	servicePrincipalPath := fmt.Sprintf("/preview/scim/v2/ServicePrincipals/%v%s", servicePrincipalID, attrs)
	err = a.client.Scim(a.context, "GET", servicePrincipalPath, nil, &sp)
	return
}

func (a ServicePrincipalsAPI) Filter(filter string, excludeRoles bool) (u []User, err error) {
	var sps UserList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	// We exclude roles to reduce load on the scim service
	if excludeRoles {
		req["excludedAttributes"] = "roles"
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
	servicePrincipal, err := a.Read(servicePrincipalID, "groups,roles")
	if err != nil {
		return err
	}
	if updateRequest.Schemas == nil {
		updateRequest.Schemas = []URN{ServicePrincipalSchema}
	}
	updateRequest.Groups = servicePrincipal.Groups
	updateRequest.Roles = servicePrincipal.Roles
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

type servicePrincipalResource struct {
	entitlements
	ApplicationID         string `json:"application_id,omitempty" tf:"computed,force_new"`
	DisplayName           string `json:"display_name,omitempty" tf:"computed"`
	Active                bool   `json:"active,omitempty"`
	ExternalID            string `json:"external_id,omitempty" tf:"suppress_diff"`
	Force                 bool   `json:"force,omitempty"`
	Home                  string `json:"home,omitempty" tf:"computed"`
	Repos                 string `json:"repos,omitempty" tf:"computed"`
	ForceDeleteRepos      bool   `json:"force_delete_repos,omitempty"`
	ForceDeleteHomeDir    bool   `json:"force_delete_home_dir,omitempty"`
	DisableAsUserDeletion bool   `json:"disable_as_user_deletion,omitempty"`
	AclPrincipalID        string `json:"acl_principal_id,omitempty" tf:"computed"`
}

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() common.Resource {
	servicePrincipalSchema := common.StructToSchema(servicePrincipalResource{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["active"].Default = true
			m["application_id"].AtLeastOneOf = []string{"application_id", "display_name"}
			m["display_name"].AtLeastOneOf = []string{"application_id", "display_name"}
			return customizeEntitlementsSchema(m)
		})
	return common.Resource{
		Schema: servicePrincipalSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var u servicePrincipalResource
			common.DataToStructPointer(d, servicePrincipalSchema, &u)
			sp := User{
				ApplicationID: u.ApplicationID,
				DisplayName:   u.DisplayName,
				Active:        u.Active,
				Entitlements:  u.entitlements.toComplexValueList(),
				ExternalID:    u.ExternalID,
			}
			spAPI := NewServicePrincipalsAPI(ctx, c)
			servicePrincipal, err := spAPI.Create(sp)
			if err != nil {
				if !u.Force {
					return err
				}
				return createForceOverridesManuallyAddedServicePrincipal(err, d, spAPI, sp)
			}
			d.SetId(servicePrincipal.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sp, err := NewServicePrincipalsAPI(ctx, c).Read(d.Id(), userAttributes)
			if err != nil {
				return err
			}
			spResource := servicePrincipalResource{
				entitlements:   newEntitlements(ctx, sp.Entitlements),
				ApplicationID:  sp.ApplicationID,
				DisplayName:    sp.DisplayName,
				Active:         sp.Active,
				ExternalID:     sp.ExternalID,
				Home:           getUserHomeDir(sp.ApplicationID),
				Repos:          getUserReposDir(sp.ApplicationID),
				AclPrincipalID: fmt.Sprintf("servicePrincipals/%s", sp.ApplicationID),
			}
			return common.StructToData(spResource, servicePrincipalSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var u servicePrincipalResource
			common.DataToStructPointer(d, servicePrincipalSchema, &u)
			sp := User{
				DisplayName:  u.DisplayName,
				Active:       u.Active,
				Entitlements: u.entitlements.toComplexValueList(),
				ExternalID:   u.ExternalID,
			}
			if c.IsAzure() {
				sp.ApplicationID = u.ApplicationID
			}
			return NewServicePrincipalsAPI(ctx, c).Update(d.Id(), sp)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var u servicePrincipalResource
			common.DataToStructPointer(d, servicePrincipalSchema, &u)
			spAPI := NewServicePrincipalsAPI(ctx, c)
			var err error = nil
			isAccount := c.Config.IsAccountClient() && c.Config.AccountID != ""
			isForceDeleteRepos := u.ForceDeleteRepos
			isForceDeleteHomeDir := u.ForceDeleteHomeDir
			// Determine if disable or delete
			var isDisable bool
			if isDisableP, exists := d.GetOkExists("disable_as_user_deletion"); exists {
				isDisable = isDisableP.(bool)
			} else {
				// Default is true for Account SCIM, false otherwise
				isDisable = isAccount
			}
			// Validate input
			if !isAccount && isDisable && isForceDeleteRepos {
				return fmt.Errorf("force_delete_repos: cannot force delete if disable_as_user_deletion is set")
			}
			if !isAccount && isDisable && isForceDeleteHomeDir {
				return fmt.Errorf("force_delete_home_dir: cannot force delete if disable_as_user_deletion is set")
			}
			// Disable or delete
			if isDisable {
				r := PatchRequestWithValue("replace", "active", "false")
				err = spAPI.Patch(d.Id(), r)
			} else {
				err = spAPI.Delete(d.Id())
			}
			if err != nil {
				return err
			}
			// Handle force delete flags
			if !isAccount && !isDisable {
				if isForceDeleteRepos {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(getUserReposDir(u.ApplicationID), true)
					if err != nil && !apierr.IsMissing(err) {
						return fmt.Errorf("force_delete_repos: %s", err.Error())
					}
				}
				if isForceDeleteHomeDir {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(getUserHomeDir(u.ApplicationID), true)
					if err != nil && !apierr.IsMissing(err) {
						return fmt.Errorf("force_delete_home_dir: %s", err.Error())
					}
				}
			}
			return nil
		},
	}
}

func createForceOverridesManuallyAddedServicePrincipal(err error, d *schema.ResourceData, spAPI ServicePrincipalsAPI, u User) error {
	// corner-case for overriding manually provisioned service principals
	knownErrs := []string{
		fmt.Sprintf("Service principal with application ID %s already exists.", u.ApplicationID),
		fmt.Sprintf("User with email %s already exists in this account", u.ApplicationID),
	}
	if !slices.Contains(knownErrs, err.Error()) {
		return err
	}
	spList, err := spAPI.Filter(fmt.Sprintf(`applicationId eq "%s"`, strings.ReplaceAll(u.ApplicationID, "'", "")), true)
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
