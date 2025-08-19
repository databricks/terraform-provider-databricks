package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type groupResource struct {
	entitlements
	DisplayName    string `json:"display_name"`
	ExternalID     string `json:"external_id,omitempty" tf:"force_new,suppress_diff"`
	URL            string `json:"url,omitempty" tf:"computed"`
	AclPrincipalID string `json:"acl_principal_id,omitempty" tf:"computed"`
	Force          bool   `json:"force,omitempty"`
}

// ResourceGroup manages user groups
func ResourceGroup() common.Resource {
	groupSchema := common.StructToSchema(groupResource{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			// https://github.com/databricks/terraform-provider-databricks/issues/1089
			m["display_name"].ValidateDiagFunc = validation.ToDiagFunc(
				validation.StringNotInSlice([]string{"users", "admins"}, false))
			return customizeEntitlementsSchema(m)
		})
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var groupResource groupResource
			common.DataToStructPointer(d, groupSchema, &groupResource)
			g := Group{
				DisplayName:  groupResource.DisplayName,
				Entitlements: groupResource.toComplexValueList(),
				ExternalID:   groupResource.ExternalID,
			}
			groupsAPI := NewGroupsAPI(ctx, c)
			group, err := groupsAPI.Create(g)
			if err != nil {
				return createForceOverridesManuallyAddedGroup(err, d, groupsAPI, g)
			}
			d.SetId(group.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(d.Id(), "displayName,externalId,entitlements")
			if err != nil {
				return err
			}
			groupResource := groupResource{
				entitlements:   newEntitlements(ctx, group.Entitlements),
				DisplayName:    group.DisplayName,
				ExternalID:     group.ExternalID,
				AclPrincipalID: fmt.Sprintf("groups/%s", group.DisplayName),
			}
			if c.Config.IsAccountClient() {
				groupResource.URL = c.FormatURL("users/groups/", d.Id(), "/information")
			} else {
				groupResource.URL = c.FormatURL("#setting/accounts/groups/", d.Id())
			}
			return common.StructToData(groupResource, groupSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var groupResource groupResource
			common.DataToStructPointer(d, groupSchema, &groupResource)
			return NewGroupsAPI(ctx, c).UpdateNameAndEntitlements(d.Id(), groupResource.DisplayName, groupResource.ExternalID, groupResource.entitlements.toComplexValueList())
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Delete(d.Id())
		},
		Schema: groupSchema,
	}
}

func createForceOverridesManuallyAddedGroup(err error, d *schema.ResourceData, groupsAPI GroupsAPI, g Group) error {
	forceCreate := d.Get("force").(bool)
	if !forceCreate {
		return err
	}
	// corner-case for overriding manually provisioned groups
	groupName := strings.ReplaceAll(g.DisplayName, "'", "")
	force := fmt.Sprintf("Group with name %s already exists.", groupName)
	if err.Error() != force {
		return err
	}
	group, err := groupsAPI.ReadByDisplayName(groupName, "")
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return groupsAPI.UpdateNameAndEntitlements(d.Id(), g.DisplayName, g.ExternalID, g.Entitlements)
}
