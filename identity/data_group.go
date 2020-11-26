package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataSourceGroup returns information about group specified by display name
func DataSourceGroup() *schema.Resource {
	type entity struct {
		DisplayName             string   `json:"display_name"`
		Recursive               bool     `json:"recursive,omitempty"`
		Members                 []string `json:"members,omitempty" tf:"slice_set,computed"`
		Groups                  []string `json:"groups,omitempty" tf:"slice_set,computed"`
		InstanceProfiles        []string `json:"instance_profiles,omitempty" tf:"slice_set,computed"`
		AllowClusterCreate      bool     `json:"allow_cluster_create,omitempty" tf:"computed"`
		AllowInstancePoolCreate bool     `json:"allow_instance_pool_create,omitempty" tf:"computed"`
	}

	s := internal.StructToSchema(entity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint once SDKv2 has Diagnostics-returning validators, change
		s["display_name"].ValidateFunc = validation.StringIsNotEmpty
		s["recursive"].Default = true
		return s
	})

	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var this entity
			err := internal.DataToStructPointer(d, s, &this)
			if err != nil {
				return diag.FromErr(err)
			}
			groupsAPI := NewGroupsAPI(ctx, m)
			groupList, err := groupsAPI.Filter(fmt.Sprintf("displayName eq %s", this.DisplayName))
			if err != nil {
				return diag.FromErr(err)
			}
			if len(groupList.Resources) == 0 {
				return diag.FromErr(fmt.Errorf("Cannot find group %s", this.DisplayName))
			}
			d.SetId(groupList.Resources[0].ID)
			queue := []ScimGroup{groupList.Resources[0]}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				for _, x := range current.Members {
					this.Members = append(this.Members, x.Value)
				}
				for _, x := range current.Roles {
					this.InstanceProfiles = append(this.InstanceProfiles, x.Value)
				}
				for _, x := range current.Entitlements {
					switch x.Value {
					case AllowClusterCreateEntitlement:
						this.AllowClusterCreate = true
					case AllowInstancePoolCreateEntitlement:
						this.AllowInstancePoolCreate = true
					}
				}
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
			err = internal.StructToData(this, s, d)
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
	}
}
