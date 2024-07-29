package mws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getPermissionsByPrincipal(list iam.PermissionAssignments, principalId int64) (res iam.UpdateWorkspaceAssignments, err error) {
	for _, v := range list.PermissionAssignments {
		if v.Principal.PrincipalId != principalId {
			continue
		}
		return iam.UpdateWorkspaceAssignments{Permissions: v.Permissions}, nil
	}
	return res, apierr.NotFound(fmt.Sprintf("%d not found", principalId))
}

func ResourceMwsPermissionAssignment() common.Resource {
	s := common.StructToSchema(iam.UpdateWorkspaceAssignments{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["workspace_id"] = &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			}
			m["principal_id"] = &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			}
			return m
		})
	pair := common.NewPairID("workspace_id", "principal_id").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			var assignment iam.UpdateWorkspaceAssignments
			common.DataToStructPointer(d, s, &assignment)
			assignment.PrincipalId = common.GetInt64(d, "principal_id")
			assignment.WorkspaceId = common.GetInt64(d, "workspace_id")
			_, err = acc.WorkspaceAssignment.Update(ctx, assignment)
			if err != nil {
				return err
			}
			pair.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			workspaceId, principalId, err := pair.Unpack(d)
			if err != nil {
				return fmt.Errorf("parse id: %w", err)
			}
			list, err := acc.WorkspaceAssignment.ListByWorkspaceId(ctx, common.MustInt64(workspaceId))
			if err != nil {
				return err
			}
			permissions, err := getPermissionsByPrincipal(*list, common.MustInt64(principalId))
			if err != nil {
				return err
			}
			return common.StructToData(permissions, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			workspaceId, principalId, err := pair.Unpack(d)
			if err != nil {
				return fmt.Errorf("parse id: %w", err)
			}
			return acc.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, common.MustInt64(workspaceId), common.MustInt64(principalId))
		},
	}
}
