package access

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewPermissionAssignmentAPI(ctx context.Context, m any) PermissionAssignmentAPI {
	return PermissionAssignmentAPI{m.(*common.DatabricksClient), ctx}
}

// permAssignmentEntry holds the cached list result for one workspace.
type permAssignmentEntry struct {
	mu          sync.RWMutex
	initialized bool
	list        permissionAssignmentResponse
	// sg deduplicates concurrent in-flight API calls when the cache is cold.
	// All goroutines that arrive while a fetch is in-flight join the same call
	// and are woken simultaneously when it completes, rather than serialising
	// through a write lock.
	sg singleflight.Group
}

// permAssignmentCache caches the permission assignments list per workspace host
// so that N databricks_permission_assignment resources in the same workspace
// only issue a single API call during a terraform plan/apply cycle.
type permAssignmentCache struct {
	mu    sync.Mutex
	cache map[string]*permAssignmentEntry
}

func newPermAssignmentCache() *permAssignmentCache {
	return &permAssignmentCache{
		cache: make(map[string]*permAssignmentEntry),
	}
}

func (c *permAssignmentCache) getOrCreate(host string) *permAssignmentEntry {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.cache[host]; ok {
		return entry
	}
	entry := &permAssignmentEntry{}
	c.cache[host] = entry
	return entry
}

func (c *permAssignmentCache) list(api PermissionAssignmentAPI) (permissionAssignmentResponse, error) {
	host := api.client.Config.Host
	entry := c.getOrCreate(host)

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
			return &l, nil
		}
		entry.mu.RUnlock()

		l, err := api.List()
		if err != nil {
			return nil, err
		}
		entry.mu.Lock()
		entry.list = l
		entry.initialized = true
		entry.mu.Unlock()
		return &l, nil
	})
	if err != nil {
		return permissionAssignmentResponse{}, err
	}
	return *v.(*permissionAssignmentResponse), nil
}

func (c *permAssignmentCache) invalidate(host string) {
	c.mu.Lock()
	entry, ok := c.cache[host]
	c.mu.Unlock()
	if ok {
		entry.mu.Lock()
		entry.initialized = false
		entry.list = permissionAssignmentResponse{}
		entry.mu.Unlock()
	}
}

var globalPermAssignmentCache = newPermAssignmentCache()

type PermissionAssignmentAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

type Permissions struct {
	Permissions []string `json:"permissions"`
}

func (a PermissionAssignmentAPI) CreateOrUpdate(assignment permissionAssignmentEntity) (principalInfo, error) {
	if assignment.PrincipalId != 0 {
		var resp permissionAssignmentResponseItem
		path := fmt.Sprintf("/api/2.0/preview/permissionassignments/principals/%d", assignment.PrincipalId)
		err := a.client.Do(a.context, http.MethodPut, path, nil, nil,
			Permissions{Permissions: assignment.Permissions}, &resp)
		if err == nil && resp.Error != "" {
			err = errors.New(resp.Error)
		}
		return resp.Principal, err
	} else {
		var principal permissionAssignmentResponse
		request := permissionAssignmentRequest{
			PermissionAssignments: []permissionAssignmentRequestItem{
				{
					principalInfo: principalInfo{
						ServicePrincipalName: assignment.ServicePrincipalName,
						UserName:             assignment.UserName,
						GroupName:            assignment.GroupName,
					},
					Permissions: assignment.Permissions,
				},
			},
		}
		err := a.client.Post(a.context, "/preview/permissionassignments", request, &principal)
		if err != nil {
			return principalInfo{}, err
		}
		if len(principal.PermissionAssignments) == 0 {
			return principalInfo{}, fmt.Errorf("no permission assignment found")
		}
		if principal.PermissionAssignments[0].Error != "" {
			return principalInfo{}, errors.New(principal.PermissionAssignments[0].Error)
		}
		return principal.PermissionAssignments[0].Principal, nil
	}
}

func (a PermissionAssignmentAPI) Remove(principalId string) error {
	path := fmt.Sprintf("/preview/permissionassignments/principals/%s", principalId)
	return a.client.Delete(a.context, path, nil)
}

type principalInfo struct {
	DisplayName          string `json:"display_name,omitempty"`
	PrincipalID          int64  `json:"principal_id,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
}

type permissionAssignmentRequestItem struct {
	principalInfo
	Permissions []string `json:"permissions"`
}

type permissionAssignmentRequest struct {
	PermissionAssignments []permissionAssignmentRequestItem `json:"permission_assignments"`
}

type permissionAssignmentResponseItem struct {
	Permissions []string `json:"permissions"`
	Principal   principalInfo
	Error       string `json:"error,omitempty"`
}

type permissionAssignmentResponse struct {
	PermissionAssignments []permissionAssignmentResponseItem `json:"permission_assignments"`
}

func (l permissionAssignmentResponse) ForPrincipal(principalId int64) (res permissionAssignmentEntity, err error) {
	for _, v := range l.PermissionAssignments {
		if v.Principal.PrincipalID != principalId {
			continue
		}
		return permissionAssignmentEntity{
			PrincipalId:          v.Principal.PrincipalID,
			ServicePrincipalName: v.Principal.ServicePrincipalName,
			UserName:             v.Principal.UserName,
			GroupName:            v.Principal.GroupName,
			Permissions:          v.Permissions,
			DisplayName:          v.Principal.DisplayName,
		}, nil
	}
	return res, &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    fmt.Sprintf("%d not found", principalId),
	}
}

func (a PermissionAssignmentAPI) List() (list permissionAssignmentResponse, err error) {
	err = a.client.Get(a.context, "/preview/permissionassignments", nil, &list)
	return
}

type permissionAssignmentEntity struct {
	PrincipalId          int64    `json:"principal_id,omitempty" tf:"computed,force_new"`
	ServicePrincipalName string   `json:"service_principal_name,omitempty" tf:"computed,force_new"`
	UserName             string   `json:"user_name,omitempty" tf:"computed,force_new"`
	GroupName            string   `json:"group_name,omitempty" tf:"computed,force_new"`
	Permissions          []string `json:"permissions" tf:"slice_as_set"`
	DisplayName          string   `json:"display_name" tf:"computed"`
	common.Namespace
}

// ResourcePermissionAssignment performs of users to a workspace
// from a workspace context, though it requires additional set
// data resource for "workspace account scim", whicl will be added later.
func ResourcePermissionAssignment() common.Resource {
	s := common.StructToSchema(permissionAssignmentEntity{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		fields := []string{"principal_id", "service_principal_name", "user_name", "group_name"}
		for _, field := range fields {
			s[field].ExactlyOneOf = fields
		}
		common.CustomizeSchemaPath(s, "display_name").SetReadOnly()
		common.NamespaceCustomizeSchemaMap(s)
		return s
	})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			defer globalPermAssignmentCache.invalidate(c.Config.Host)
			var assignment permissionAssignmentEntity
			common.DataToStructPointer(d, s, &assignment)
			api := NewPermissionAssignmentAPI(ctx, c)
			// We need this because assignment by name doesn't work for admins, so we need to
			// first assign them as users.  And then reassign them as admins.
			shouldReassignAdmin := false
			if assignment.PrincipalId == 0 && slices.Contains(assignment.Permissions, "ADMIN") {
				shouldReassignAdmin = true
				assignment.Permissions = []string{"USER"}
			}
			principal, err := api.CreateOrUpdate(assignment)
			if err != nil {
				return err
			}
			d.SetId(strconv.FormatInt(principal.PrincipalID, 10))
			if shouldReassignAdmin {
				common.DataToStructPointer(d, s, &assignment)
				assignment.PrincipalId = principal.PrincipalID
				_, err := api.CreateOrUpdate(assignment)
				if err != nil {
					log.Printf("[WARN] error reassigning admin permissions: %v", err)
					api.Remove(d.Id())
				}
				return err
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			list, err := globalPermAssignmentCache.list(NewPermissionAssignmentAPI(ctx, c))
			if err != nil {
				return err
			}
			data, err := list.ForPrincipal(common.MustInt64(d.Id()))
			if err != nil {
				return err
			}
			return common.StructToData(data, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			defer globalPermAssignmentCache.invalidate(c.Config.Host)
			var assignment permissionAssignmentEntity
			common.DataToStructPointer(d, s, &assignment)
			api := NewPermissionAssignmentAPI(ctx, c)
			_, err = api.CreateOrUpdate(assignment)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			defer globalPermAssignmentCache.invalidate(c.Config.Host)
			return NewPermissionAssignmentAPI(ctx, c).Remove(d.Id())
		},
	}
}
