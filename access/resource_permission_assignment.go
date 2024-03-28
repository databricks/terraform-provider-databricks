package access

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewPermissionAssignmentAPI(ctx context.Context, m any) PermissionAssignmentAPI {
	return PermissionAssignmentAPI{m.(*common.DatabricksClient), ctx}
}

type PermissionAssignmentAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

type Permissions struct {
	Permissions []string `json:"permissions"`
}

func (a PermissionAssignmentAPI) CreateOrUpdate(principalId int64, r Permissions) error {
	path := fmt.Sprintf("/preview/permissionassignments/principals/%d", principalId)
	return a.client.Put(a.context, path, r)
}

func (a PermissionAssignmentAPI) Remove(principalId string) error {
	path := fmt.Sprintf("/preview/permissionassignments/principals/%s", principalId)
	return a.client.Delete(a.context, path, nil)
}

type Principal struct {
	DisplayName          string `json:"display_name"`
	PrincipalID          int64  `json:"principal_id"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
}

type PermissionAssignment struct {
	Permissions []string `json:"permissions"`
	Principal   Principal
}

type PermissionAssignmentList struct {
	PermissionAssignments []PermissionAssignment `json:"permission_assignments"`
}

func (l PermissionAssignmentList) ForPrincipal(principalId int64) (res Permissions, err error) {
	for _, v := range l.PermissionAssignments {
		if v.Principal.PrincipalID != principalId {
			continue
		}
		return Permissions{v.Permissions}, nil
	}
	return res, apierr.NotFound(fmt.Sprintf("%d not found", principalId))
}

func (a PermissionAssignmentAPI) List() (list PermissionAssignmentList, err error) {
	err = a.client.Get(a.context, "/preview/permissionassignments", nil, &list)
	return
}

// ResourcePermissionAssignment performs of users to a workspace
// from a workspace context, though it requires additional set
// data resource for "workspace account scim", whicl will be added later.
func ResourcePermissionAssignment() common.Resource {
	type entity struct {
		PrincipalId int64    `json:"principal_id"`
		Permissions []string `json:"permissions" tf:"slice_as_set"`
	}
	s := common.StructToSchema(entity{}, common.NoCustomize)
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var assignment entity
			common.DataToStructPointer(d, s, &assignment)
			api := NewPermissionAssignmentAPI(ctx, c)
			err := api.CreateOrUpdate(assignment.PrincipalId, Permissions{assignment.Permissions})
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%d", assignment.PrincipalId))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			list, err := NewPermissionAssignmentAPI(ctx, c).List()
			if err != nil {
				return err
			}
			data := entity{
				PrincipalId: common.MustInt64(d.Id()),
			}
			permissions, err := list.ForPrincipal(data.PrincipalId)
			if err != nil {
				return err
			}
			data.Permissions = permissions.Permissions
			return common.StructToData(data, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewPermissionAssignmentAPI(ctx, c).Remove(d.Id())
		},
	}
}
