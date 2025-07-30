package scim

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceGroups returns information about groups that match the specified filter
func DataSourceGroups() common.Resource {
	type groupEntity struct {
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

	type groupsData struct {
		Filter string        `json:"filter,omitempty"`
		Groups []groupEntity `json:"groups,omitempty" tf:"computed"`
	}

	return common.DataResource(groupsData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		response := e.(*groupsData)
		groupsAPI := NewGroupsAPI(ctx, c)

		groupList, err := groupsAPI.Filter(response.Filter)
		if err != nil {
			return err
		}

		// Initialize empty groups slice to ensure we return empty array when no matches
		response.Groups = []groupEntity{}

		for _, group := range groupList.Resources {
			// Convert each Group to the expected entity shape
			entity := groupEntity{
				DisplayName:    group.DisplayName,
				ExternalID:     group.ExternalID,
				AclPrincipalID: fmt.Sprintf("groups/%s", group.DisplayName),
			}

			// Process members to extract different types
			for _, member := range group.Members {
				entity.Members = append(entity.Members, member.Value)
				if strings.HasPrefix(member.Ref, "Users/") {
					entity.Users = append(entity.Users, member.Value)
				}
				if strings.HasPrefix(member.Ref, "Groups/") {
					entity.ChildGroups = append(entity.ChildGroups, member.Value)
				}
				if strings.HasPrefix(member.Ref, "ServicePrincipals/") {
					entity.ServicePrincipals = append(entity.ServicePrincipals, member.Value)
				}
			}

			// Process roles (instance profiles)
			for _, role := range group.Roles {
				entity.InstanceProfiles = append(entity.InstanceProfiles, role.Value)
			}

			// Process parent groups
			for _, parentGroup := range group.Groups {
				entity.Groups = append(entity.Groups, parentGroup.Value)
			}

			// Sort all slices for consistent output
			sort.Strings(entity.Members)
			sort.Strings(entity.Users)
			sort.Strings(entity.ServicePrincipals)
			sort.Strings(entity.ChildGroups)
			sort.Strings(entity.Groups)
			sort.Strings(entity.InstanceProfiles)

			response.Groups = append(response.Groups, entity)
		}

		// Sort groups by display name for consistent output
		sort.Slice(response.Groups, func(i, j int) bool {
			return response.Groups[i].DisplayName < response.Groups[j].DisplayName
		})

		return nil
	})
}
