package identity

import (
	"context"
	"sort"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataSourceGroup returns information about group specified by display name
func DataSourceGroup() *schema.Resource {
	type entity struct {
		DisplayName      string   `json:"display_name"`
		Recursive        bool     `json:"recursive,omitempty"`
		Members          []string `json:"members,omitempty" tf:"slice_set,computed"`
		Groups           []string `json:"groups,omitempty" tf:"slice_set,computed"`
		InstanceProfiles []string `json:"instance_profiles,omitempty" tf:"slice_set,computed"`
	}

	s := common.StructToSchema(entity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint once SDKv2 has Diagnostics-returning validators, change
		s["display_name"].ValidateFunc = validation.StringIsNotEmpty
		s["recursive"].Default = true
		addEntitlementsToSchema(&s)
		return s
	})

	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var this entity
			err := common.DataToStructPointer(d, s, &this)
			if err != nil {
				return diag.FromErr(err)
			}
			groupsAPI := NewGroupsAPI(ctx, m)
			group, err := groupsAPI.ReadByDisplayName(this.DisplayName)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(group.ID)
			queue := []ScimGroup{group}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				for _, x := range current.Members {
					this.Members = append(this.Members, x.Value)
				}
				for _, x := range current.Roles {
					this.InstanceProfiles = append(this.InstanceProfiles, x.Value)
				}
				current.Entitlements.readIntoData(d)
				for _, x := range current.Groups {
					this.Groups = append(this.Groups, x.Value)
					if this.Recursive {
						childGroup, err := groupsAPI.Read(x.Value)
						if err != nil {
							return diag.FromErr(err)
						}
						queue = append(queue, childGroup)
					}
				}
			}
			sort.Strings(this.Groups)
			sort.Strings(this.Members)
			sort.Strings(this.InstanceProfiles)
			err = common.StructToData(this, s, d)
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
	}
}
