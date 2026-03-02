package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceEntitlements() common.Resource {
	type entity struct {
		GroupId string `json:"group_id,omitempty" tf:"force_new"`
		UserId  string `json:"user_id,omitempty" tf:"force_new"`
		SpnId   string `json:"service_principal_id,omitempty" tf:"force_new"`
	}
	entitlementSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(m)
			alof := []string{"group_id", "user_id", "service_principal_id"}
			for _, field := range alof {
				m[field].AtLeastOneOf = alof
			}
			return m
		})
	addEntitlementsToSchema(entitlementSchema)
	common.AddNamespaceInSchema(entitlementSchema)
	common.NamespaceCustomizeSchemaMap(entitlementSchema)
	return common.Resource{
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.HostType() != config.WorkspaceHost {
				return fmt.Errorf("entitlements can only be managed with a provider configured at the workspace-level")
			}
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = patchEntitlements(ctx, d, newClient, "replace")
			if err != nil {
				return err
			}
			d.SetId(getId(d))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements: %s", d.Id())
			}
			switch strings.ToLower(split[0]) {
			case "group":
				group, err := NewGroupsAPI(ctx, newClient).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				d.Set("group_id", split[1])
				group.Entitlements.generateEmpty(d)
				return group.Entitlements.readIntoData(d)
			case "user":
				user, err := NewUsersAPI(ctx, newClient).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				d.Set("user_id", split[1])
				user.Entitlements.generateEmpty(d)
				return user.Entitlements.readIntoData(d)
			case "spn":
				spn, err := NewServicePrincipalsAPI(ctx, newClient).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				d.Set("service_principal_id", split[1])
				spn.Entitlements.generateEmpty(d)
				return spn.Entitlements.readIntoData(d)
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return patchEntitlements(ctx, d, newClient, "replace")
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return patchEntitlements(ctx, d, newClient, "remove")
		},
		Schema: entitlementSchema,
	}
}

func getId(d *schema.ResourceData) string {
	groupId := d.Get("group_id").(string)
	userId := d.Get("user_id").(string)
	spnId := d.Get("service_principal_id").(string)
	if groupId != "" {
		return "group/" + groupId
	}
	if userId != "" {
		return "user/" + userId
	}
	if spnId != "" {
		return "spn/" + spnId
	}
	return ""
}

func patchEntitlements(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient, op string) error {
	groupId := d.Get("group_id").(string)
	userId := d.Get("user_id").(string)
	spnId := d.Get("service_principal_id").(string)
	noEntitlementMessage := "invalidPath No such attribute with the name : entitlements in the current resource"
	entitlements := readEntitlementsFromData(d)
	if len(entitlements) == 1 && entitlements[0].Value == "" && op == "remove" {
		// No updates are needed, so return early
		return nil
	}
	request := PatchRequestComplexValue([]patchOperation{
		{
			op,
			"entitlements",
			entitlements,
		},
	})
	if groupId != "" {
		groupsAPI := NewGroupsAPI(ctx, c)
		err := groupsAPI.UpdateEntitlements(groupId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
	}
	if userId != "" {
		usersAPI := NewUsersAPI(ctx, c)
		err := usersAPI.UpdateEntitlements(userId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
	}
	if spnId != "" {
		spnAPI := NewServicePrincipalsAPI(ctx, c)
		err := spnAPI.UpdateEntitlements(spnId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
	}
	return nil
}
