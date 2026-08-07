package mws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// workspaceAssignmentsCacheKey uniquely identifies a workspace within an account.
// Workspace IDs are account-scoped, not globally unique, so the accountID is
// required to prevent two providers targeting different accounts from sharing
// cached workspace assignments.
type workspaceAssignmentsCacheKey struct {
	accountID   string
	workspaceId int64
}

// globalWorkspaceAssignmentsCache caches ListByWorkspaceId results per
// (accountID, workspaceId) so that N databricks_mws_permission_assignment
// resources sharing the same workspace_id only issue a single API call during
// a terraform plan/apply cycle.
var globalWorkspaceAssignmentsCache = newWorkspaceAssignmentsCache()

func newWorkspaceAssignmentsCache() *common.KeyedCache[workspaceAssignmentsCacheKey, *iam.PermissionAssignments] {
	return common.NewKeyedCache[workspaceAssignmentsCacheKey, *iam.PermissionAssignments]()
}

func getPermissionsByPrincipal(list iam.PermissionAssignments, principalId int64) (res iam.UpdateWorkspaceAssignments, err error) {
	for _, v := range list.PermissionAssignments {
		if v.Principal.PrincipalId != principalId {
			continue
		}
		return iam.UpdateWorkspaceAssignments{Permissions: v.Permissions}, nil
	}
	return res, &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    fmt.Sprintf("%d not found", principalId),
	}
}

func ResourceMwsPermissionAssignment() common.Resource {
	s := common.StructToSchema(iam.UpdateWorkspaceAssignments{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m).AddNewField("workspace_id", &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			}).AddNewField("principal_id", &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			})
			common.CustomizeSchemaPath(m, "permissions").SetRequired().SetSliceSet()
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
			cacheKey := workspaceAssignmentsCacheKey{accountID: c.Config.AccountID, workspaceId: assignment.WorkspaceId}
			defer globalWorkspaceAssignmentsCache.Invalidate(cacheKey)
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
			wsID := common.MustInt64(workspaceId)
			cacheKey := workspaceAssignmentsCacheKey{accountID: c.Config.AccountID, workspaceId: wsID}
			list, err := globalWorkspaceAssignmentsCache.Get(cacheKey, func() (*iam.PermissionAssignments, error) {
				return acc.WorkspaceAssignment.ListByWorkspaceId(ctx, wsID)
			})
			if err != nil {
				return err
			}
			permissions, err := getPermissionsByPrincipal(*list, common.MustInt64(principalId))
			if err != nil {
				return err
			}
			d.Set("workspace_id", wsID)
			d.Set("principal_id", common.MustInt64(principalId))
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
			wsID := common.MustInt64(workspaceId)
			cacheKey := workspaceAssignmentsCacheKey{accountID: c.Config.AccountID, workspaceId: wsID}
			defer globalWorkspaceAssignmentsCache.Invalidate(cacheKey)
			return acc.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, wsID, common.MustInt64(principalId))
		},
	}
}
