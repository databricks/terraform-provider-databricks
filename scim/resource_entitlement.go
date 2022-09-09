package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceEntitlement() *schema.Resource {
	type entity struct {
		GroupId string `json:"group_id,omitempty" tf:"force_new"`
		UserId  string `json:"user_id,omitempty" tf:"force_new"`
		SpnId   string `json:"spn_id,omitempty" tf:"force_new"`
	}
	entitlementSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(&m)
			// https://github.com/databricks/terraform-provider-databricks/issues/1089
			alof := []string{"group_id", "user_id", "spn_id"}
			for _, field := range alof {
				m[field].AtLeastOneOf = alof
			}
			return m
		})
	addEntitlementsToSchema(&entitlementSchema)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return createEntitlement(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements: %s", d.Id())
			}
			switch strings.ToLower(split[0]) {
			case "group":
				group, err := NewGroupsAPI(ctx, c).Read(split[1])
				if err != nil {
					return err
				}
				return group.Entitlements.readIntoData(d)
			case "user":
				user, err := NewUsersAPI(ctx, c).Read(split[1])
				if err != nil {
					return err
				}
				return user.Entitlements.readIntoData(d)
			case "spn":
				spn, err := NewServicePrincipalsAPI(ctx, c).Read(split[1])
				if err != nil {
					return err
				}
				return spn.Entitlements.readIntoData(d)
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return enforceEntitlement(ctx, d, c)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e entitlements
			e.readIntoData(d)
			return enforceEntitlement(ctx, d, c)
		},
		Schema: entitlementSchema,
	}.ToResource()
}

func createEntitlement(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	groupId := d.Get("group_id").(string)
	userId := d.Get("user_id").(string)
	spnId := d.Get("spn_id").(string)
	if groupId != "" {
		groupsAPI := NewGroupsAPI(ctx, c)
		err := groupsAPI.UpdateEntitlements(groupId, readEntitlementsFromData(d))
		if err != nil {
			return err
		}
		d.SetId("group/" + groupId)
	}
	if userId != "" {
		usersAPI := NewUsersAPI(ctx, c)
		err := usersAPI.UpdateEntitlements(userId, readEntitlementsFromData(d))
		if err != nil {
			return err
		}
		d.SetId("user/" + userId)
	}
	if spnId != "" {
		spnAPI := NewServicePrincipalsAPI(ctx, c)
		err := spnAPI.UpdateEntitlements(spnId, readEntitlementsFromData(d))
		if err != nil {
			return err
		}
		d.SetId("spn/" + spnId)
	}
	return nil
}

func enforceEntitlement(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	split := strings.SplitN(d.Id(), "/", 2)
	if len(split) != 2 {
		return fmt.Errorf("ID must be two elements: %s", d.Id())
	}
	switch strings.ToLower(split[0]) {
	case "group":
		return NewGroupsAPI(ctx, c).UpdateEntitlements(split[1], readEntitlementsFromData(d))
	case "user":
		return NewUsersAPI(ctx, c).UpdateEntitlements(split[1], readEntitlementsFromData(d))
	case "spn":
		return NewServicePrincipalsAPI(ctx, c).UpdateEntitlements(split[1], readEntitlementsFromData(d))
	}
	return nil
}
