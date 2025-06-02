package scim

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type groupData struct {
	entitlements
	DisplayName       string   `json:"display_name"`
	Recursive         bool     `json:"recursive,omitempty"`
	Members           []string `json:"members,omitempty" tf:"slice_set,computed"`
	Users             []string `json:"users,omitempty" tf:"slice_set,computed"`
	ServicePrincipals []string `json:"service_principals,omitempty" tf:"slice_set,computed"`
	ChildGroups       []string `json:"child_groups,omitempty" tf:"slice_set,computed"`
	Groups            []string `json:"groups,omitempty" tf:"slice_set,computed"`
	InstanceProfiles  []string `json:"instance_profiles,omitempty" tf:"slice_set,computed"`
	ExternalID        string   `json:"external_id,omitempty" tf:"computed"`
	AclPrincipalID    string   `json:"acl_principal_id,omitempty" tf:"computed"`
}

// DataSourceGroup returns information about group specified by display name
func DataSourceGroup() common.Resource {
	s := common.StructToSchema(groupData{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint once SDKv2 has Diagnostics-returning validators, change
		s["display_name"].ValidateFunc = validation.StringIsNotEmpty
		s["recursive"].Default = true
		s["members"].Deprecated = "Please use `users`, `service_principals`, and `child_groups` instead"
		return s
	})

	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			var this groupData
			common.DataToStructPointer(d, s, &this)
			groupsAPI := NewGroupsAPI(ctx, m)
			groupAttributes := "members,roles,entitlements,externalId"
			group, err := groupsAPI.ReadByDisplayName(this.DisplayName, groupAttributes)
			if err != nil {
				return err
			}
			d.SetId(group.ID)
			queue := []Group{group}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				for _, x := range current.Members {
					this.Members = append(this.Members, x.Value)
					if strings.HasPrefix(x.Ref, "Users/") {
						this.Users = append(this.Users, x.Value)
					}
					if strings.HasPrefix(x.Ref, "Groups/") {
						this.ChildGroups = append(this.ChildGroups, x.Value)
					}
					if strings.HasPrefix(x.Ref, "ServicePrincipals/") {
						this.ServicePrincipals = append(this.ServicePrincipals, x.Value)
					}
				}
				for _, x := range current.Roles {
					this.InstanceProfiles = append(this.InstanceProfiles, x.Value)
				}
				this.entitlements = fromComplexValueList(ctx, current.Entitlements)
				for _, x := range current.Groups {
					this.Groups = append(this.Groups, x.Value)
					if this.Recursive {
						childGroup, err := groupsAPI.Read(x.Value, groupAttributes)
						if err != nil {
							return err
						}
						queue = append(queue, childGroup)
					}
				}
			}
			this.ExternalID = group.ExternalID
			this.AclPrincipalID = fmt.Sprintf("groups/%s", group.DisplayName)
			sort.Strings(this.Groups)
			sort.Strings(this.Members)
			sort.Strings(this.Users)
			sort.Strings(this.ChildGroups)
			sort.Strings(this.ServicePrincipals)
			sort.Strings(this.InstanceProfiles)
			err = common.StructToData(this, s, d)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
