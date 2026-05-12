package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceGroup manages user groups
func ResourceGroup() common.Resource {
	type entity struct {
		DisplayName string `json:"display_name"`
		ExternalID  string `json:"external_id,omitempty" tf:"force_new,suppress_diff"`
		URL         string `json:"url,omitempty" tf:"computed"`
	}
	groupSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(m)
			common.AddApiField(m)
			// https://github.com/databricks/terraform-provider-databricks/issues/1089
			m["display_name"].ValidateDiagFunc = validation.ToDiagFunc(
				validation.StringNotInSlice([]string{"users", "admins"}, false))
			m["force"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["acl_principal_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}
			return m
		})
	addEntitlementsToSchema(groupSchema)
	common.AddNamespaceInSchema(groupSchema)
	common.NamespaceCustomizeSchemaMap(groupSchema)
	return common.Resource{
		IsDual: true,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.CustomizeDiffDualResources(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			g := Group{
				DisplayName:  d.Get("display_name").(string),
				Entitlements: readEntitlementsFromData(d),
				ExternalID:   d.Get("external_id").(string),
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			group, err := groupsAPI.Create(g)
			if err != nil {
				return createForceOverridesManuallyAddedGroup(err, d, groupsAPI, g)
			}
			d.SetId(group.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			group, err := groupsAPI.Read(d.Id(), "displayName,externalId,entitlements")
			if err != nil {
				return err
			}
			d.Set("display_name", group.DisplayName)
			d.Set("external_id", group.ExternalID)
			d.Set("acl_principal_id", fmt.Sprintf("groups/%s", group.DisplayName))
			if common.IsAccountLevel(d, c) {
				d.Set("url", c.FormatURL("users/groups/", d.Id(), "/information"))
			} else {
				d.Set("url", c.FormatURL("#setting/accounts/groups/", d.Id()))
			}
			return group.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			groupName := d.Get("display_name").(string)
			return groupsAPI.UpdateNameAndEntitlements(d.Id(), groupName,
				d.Get("external_id").(string), readEntitlementsFromData(d))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			return groupsAPI.Delete(d.Id())
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
