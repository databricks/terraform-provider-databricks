package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceEntitlements() *schema.Resource {
	type entity struct {
		GroupId string `json:"group_id,omitempty" tf:"force_new"`
		UserId  string `json:"user_id,omitempty" tf:"force_new"`
		SpnId   string `json:"service_principal_id,omitempty" tf:"force_new"`
	}
	entitlementSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(&m)
			alof := []string{"group_id", "user_id", "service_principal_id"}
			for _, field := range alof {
				m[field].AtLeastOneOf = alof
			}
			return m
		})
	addEntitlementsToSchema(&entitlementSchema)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return patchEntitlements(ctx, d, c, "replace")
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements: %s", d.Id())
			}
			switch strings.ToLower(split[0]) {
			case "group":
				group, err := NewGroupsAPI(ctx, c).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				group.Entitlements.generateEmpty(d)
				return group.Entitlements.readIntoData(d)
			case "user":
				user, err := NewUsersAPI(ctx, c).Read(split[1], "entitlements")
				if err != nil {
					return err
				}
				user.Entitlements.generateEmpty(d)
				return user.Entitlements.readIntoData(d)
			case "spn":
				spn, err := NewServicePrincipalsAPI(ctx, c).Read(split[1])
				if err != nil {
					return err
				}
				spn.Entitlements.generateEmpty(d)
				return spn.Entitlements.readIntoData(d)
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return enforceEntitlements(ctx, d, c)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return patchEntitlements(ctx, d, c, "remove")
		},
		Schema: entitlementSchema,
	}.ToResource()
}

func patchEntitlements(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient, op string) error {
	groupId := d.Get("group_id").(string)
	userId := d.Get("user_id").(string)
	spnId := d.Get("service_principal_id").(string)
	noEntitlementMessage := "invalidPath No such attribute with the name : entitlements in the current resource"
	request := PatchRequestComplexValue([]patchOperation{
		{
			op,
			"entitlements",
			readEntitlementsFromData(d),
		},
	})
	if groupId != "" {
		groupsAPI := NewGroupsAPI(ctx, c)
		err := groupsAPI.UpdateEntitlements(groupId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
		d.SetId("group/" + groupId)
	}
	if userId != "" {
		usersAPI := NewUsersAPI(ctx, c)
		err := usersAPI.UpdateEntitlements(userId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
		d.SetId("user/" + userId)
	}
	if spnId != "" {
		spnAPI := NewServicePrincipalsAPI(ctx, c)
		err := spnAPI.UpdateEntitlements(spnId, request)
		if err != nil && !strings.Contains(err.Error(), noEntitlementMessage) {
			return err
		}
		d.SetId("spn/" + spnId)
	}
	return nil
}

func enforceEntitlements(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	split := strings.SplitN(d.Id(), "/", 2)
	if len(split) != 2 {
		return fmt.Errorf("ID must be two elements: %s", d.Id())
	}
	identity := strings.ToLower(split[0])
	id := strings.ToLower(split[1])
	request := PatchRequestComplexValue(
		[]patchOperation{
			{
				"remove", "entitlements", generateFullEntitlements(),
			},
			{
				"add", "entitlements", readEntitlementsFromData(d),
			},
		},
	)
	switch identity {
	case "group":
		NewGroupsAPI(ctx, c).UpdateEntitlements(id, request)
	case "user":
		NewUsersAPI(ctx, c).UpdateEntitlements(id, request)
	case "spn":
		NewServicePrincipalsAPI(ctx, c).UpdateEntitlements(id, request)
	}
	return nil
}
