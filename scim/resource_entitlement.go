package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type entitlementsResource struct {
	entitlements
	GroupId string `json:"group_id,omitempty" tf:"force_new"`
	UserId  string `json:"user_id,omitempty" tf:"force_new"`
	SpnId   string `json:"service_principal_id,omitempty" tf:"force_new"`
}

// ResourceGroup manages user groups
func ResourceEntitlements() common.Resource {
	entitlementSchema := common.StructToSchema(entitlementsResource{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			alof := []string{"group_id", "user_id", "service_principal_id"}
			for _, field := range alof {
				m[field].AtLeastOneOf = alof
			}
			return m
		})
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.IsAccountClient() {
				return fmt.Errorf("entitlements can only be managed with a provider configured at the workspace-level")
			}
			var e entitlementsResource
			common.DataToStructPointer(d, entitlementSchema, &e)
			err := patchEntitlements(ctx, e, c, "replace")
			if err != nil {
				return err
			}
			d.SetId(getId(e))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements: %s", d.Id())
			}
			var entitlements entitlementsResource
			switch strings.ToLower(split[0]) {
			case "group":
				group, err := NewGroupsAPI(ctx, c).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				entitlements.GroupId = split[1]
				entitlements.entitlements = newEntitlements(ctx, group.Entitlements)
			case "user":
				user, err := NewUsersAPI(ctx, c).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				entitlements.UserId = split[1]
				entitlements.entitlements = newEntitlements(ctx, user.Entitlements)
			case "spn":
				spn, err := NewServicePrincipalsAPI(ctx, c).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				entitlements.SpnId = split[1]
				entitlements.entitlements = newEntitlements(ctx, spn.Entitlements)
			}
			return common.StructToData(entitlements, entitlementSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e entitlementsResource
			common.DataToStructPointer(d, entitlementSchema, &e)
			return patchEntitlements(ctx, e, c, "replace")
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e entitlementsResource
			common.DataToStructPointer(d, entitlementSchema, &e)
			return patchEntitlements(ctx, e, c, "remove")
		},
		Schema: entitlementSchema,
	}
}

func getId(e entitlementsResource) string {
	groupId := e.GroupId
	userId := e.UserId
	spnId := e.SpnId
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

func patchEntitlements(ctx context.Context, e entitlementsResource, c *common.DatabricksClient, op string) error {
	groupId := e.GroupId
	userId := e.UserId
	spnId := e.SpnId
	noEntitlementMessage := "invalidPath No such attribute with the name : entitlements in the current resource"
	entitlements := e.toComplexValueList()
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
