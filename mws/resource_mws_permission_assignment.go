package mws

import (
	"context"
	"errors"
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

func (a PermissionAssignmentAPI) CreateOrUpdate(workspaceId, principalId int64, r Permissions) error {
	if a.client.Config.AccountID == "" {
		return errors.New("must have `account_id` on provider")
	}
	path := fmt.Sprintf(
		"/accounts/%s/workspaces/%d/permissionassignments/principals/%d",
		a.client.Config.AccountID, workspaceId, principalId)
	return a.client.Put(a.context, path, r)
}

func (a PermissionAssignmentAPI) Remove(workspaceId, principalId string) error {
	if a.client.Config.AccountID == "" {
		return errors.New("must have `account_id` on provider")
	}
	path := fmt.Sprintf(
		"/accounts/%s/workspaces/%s/permissionassignments/principals/%s",
		a.client.Config.AccountID, workspaceId, principalId)
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

func (a PermissionAssignmentAPI) List(workspaceId int64) (list PermissionAssignmentList, err error) {
	if a.client.Config.AccountID == "" {
		return list, errors.New("must have `account_id` on provider")
	}
	path := fmt.Sprintf("/accounts/%s/workspaces/%d/permissionassignments",
		a.client.Config.AccountID, workspaceId)
	err = a.client.Get(a.context, path, nil, &list)
	return
}

func ResourceMwsPermissionAssignment() common.Resource {
	type entity struct {
		WorkspaceId int64    `json:"workspace_id"`
		PrincipalId int64    `json:"principal_id"`
		Permissions []string `json:"permissions" tf:"slice_as_set"`
	}
	s := common.StructToSchema(entity{},
		common.NoCustomize)
	pair := common.NewPairID("workspace_id", "principal_id").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var assignment entity
			common.DataToStructPointer(d, s, &assignment)
			api := NewPermissionAssignmentAPI(ctx, c)
			err := api.CreateOrUpdate(assignment.WorkspaceId, assignment.PrincipalId,
				Permissions{assignment.Permissions})
			if err != nil {
				return err
			}
			pair.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId, principalId, err := pair.Unpack(d)
			if err != nil {
				return fmt.Errorf("parse id: %w", err)
			}
			list, err := NewPermissionAssignmentAPI(ctx, c).List(common.MustInt64(workspaceId))
			if err != nil {
				return err
			}
			permissions, err := list.ForPrincipal(common.MustInt64(principalId))
			if err != nil {
				return err
			}
			return common.StructToData(permissions, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			workspaceId, principalId, err := pair.Unpack(d)
			if err != nil {
				return fmt.Errorf("parse id: %w", err)
			}
			return NewPermissionAssignmentAPI(ctx, c).Remove(workspaceId, principalId)
		},
	}
}
