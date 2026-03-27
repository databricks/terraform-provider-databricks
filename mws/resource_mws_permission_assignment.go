package mws

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// workspaceAssignmentsEntry holds the cached ListByWorkspaceId result for one workspace.
type workspaceAssignmentsEntry struct {
	mu          sync.RWMutex
	initialized bool
	list        *iam.PermissionAssignments
	// sg deduplicates concurrent in-flight API calls when the cache is cold.
	// All goroutines that arrive while a fetch is in-flight join the same call
	// and are woken simultaneously when it completes, rather than serialising
	// through a write lock.
	sg singleflight.Group
}

// workspaceAssignmentsCache caches ListByWorkspaceId results per workspace ID so
// that N databricks_mws_permission_assignment resources sharing the same
// workspace_id only issue a single API call during a terraform plan/apply cycle.
type workspaceAssignmentsCache struct {
	mu    sync.Mutex
	cache map[int64]*workspaceAssignmentsEntry
}

func newWorkspaceAssignmentsCache() *workspaceAssignmentsCache {
	return &workspaceAssignmentsCache{
		cache: make(map[int64]*workspaceAssignmentsEntry),
	}
}

func (c *workspaceAssignmentsCache) getOrCreate(workspaceId int64) *workspaceAssignmentsEntry {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.cache[workspaceId]; ok {
		return entry
	}
	entry := &workspaceAssignmentsEntry{}
	c.cache[workspaceId] = entry
	return entry
}

func (c *workspaceAssignmentsCache) list(ctx context.Context, api iam.WorkspaceAssignmentInterface, workspaceId int64) (*iam.PermissionAssignments, error) {
	entry := c.getOrCreate(workspaceId)

	// Fast path: warm cache. Many goroutines can hold a read-lock simultaneously.
	entry.mu.RLock()
	if entry.initialized {
		l := entry.list
		entry.mu.RUnlock()
		return l, nil
	}
	entry.mu.RUnlock()

	// Slow path: cache is cold. Use singleflight so exactly one API call is made
	// regardless of how many goroutines arrive concurrently; all share the result.
	v, err, _ := entry.sg.Do("fetch", func() (interface{}, error) {
		// Double-check now that we are the singleflight leader.
		entry.mu.RLock()
		if entry.initialized {
			l := entry.list
			entry.mu.RUnlock()
			return l, nil
		}
		entry.mu.RUnlock()

		list, err := api.ListByWorkspaceId(ctx, workspaceId)
		if err != nil {
			return nil, err
		}
		entry.mu.Lock()
		entry.list = list
		entry.initialized = true
		entry.mu.Unlock()
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return v.(*iam.PermissionAssignments), nil
}

func (c *workspaceAssignmentsCache) invalidate(workspaceId int64) {
	c.mu.Lock()
	entry, ok := c.cache[workspaceId]
	c.mu.Unlock()
	if ok {
		entry.mu.Lock()
		entry.initialized = false
		entry.list = nil
		entry.mu.Unlock()
	}
}

var globalWorkspaceAssignmentsCache = newWorkspaceAssignmentsCache()

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
			_, err = acc.WorkspaceAssignment.Update(ctx, assignment)
			if err != nil {
				return err
			}
			globalWorkspaceAssignmentsCache.invalidate(assignment.WorkspaceId)
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
			list, err := globalWorkspaceAssignmentsCache.list(ctx, acc.WorkspaceAssignment, common.MustInt64(workspaceId))
			if err != nil {
				return err
			}
			permissions, err := getPermissionsByPrincipal(*list, common.MustInt64(principalId))
			if err != nil {
				return err
			}
			d.Set("workspace_id", common.MustInt64(workspaceId))
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
			err = acc.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, common.MustInt64(workspaceId), common.MustInt64(principalId))
			globalWorkspaceAssignmentsCache.invalidate(common.MustInt64(workspaceId))
			return err
		},
	}
}
