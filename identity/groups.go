package identity

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
)

// NewGroupsAPI creates GroupsAPI instance from provider meta
func NewGroupsAPI(ctx context.Context, m interface{}) GroupsAPI {
	return GroupsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// GroupsAPI exposes the scim groups API
type GroupsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a scim group in the Databricks workspace
func (a GroupsAPI) Create(groupName string, members []string, roles []string, entitlements []string) (group ScimGroup, err error) {
	scimGroupRequest := struct {
		Schemas      []URN           `json:"schemas,omitempty"`
		DisplayName  string          `json:"displayName,omitempty"`
		Members      []ValueListItem `json:"members,omitempty"`
		Entitlements []ValueListItem `json:"entitlements,omitempty"`
		Roles        []ValueListItem `json:"roles,omitempty"`
	}{}
	scimGroupRequest.Schemas = []URN{GroupSchema}
	scimGroupRequest.DisplayName = groupName

	scimGroupRequest.Members = []ValueListItem{}
	for _, member := range members {
		scimGroupRequest.Members = append(scimGroupRequest.Members, ValueListItem{Value: member})
	}

	scimGroupRequest.Roles = []ValueListItem{}
	for _, role := range roles {
		scimGroupRequest.Roles = append(scimGroupRequest.Roles, ValueListItem{Value: role})
	}

	scimGroupRequest.Entitlements = []ValueListItem{}
	for _, entitlement := range entitlements {
		scimGroupRequest.Entitlements = append(scimGroupRequest.Entitlements, ValueListItem{Value: entitlement})
	}
	err = a.client.Scim(a.context, http.MethodPost, "/preview/scim/v2/Groups", scimGroupRequest, &group)
	return
}

// Read reads and returns a Group object via SCIM api
func (a GroupsAPI) Read(groupID string) (group ScimGroup, err error) {
	err = a.client.Scim(a.context, http.MethodGet, fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), nil, &group)
	if err != nil {
		return
	}
	return
}

// Filter returns groups matching the filter
func (a GroupsAPI) Filter(filter string) (GroupList, error) {
	var groups GroupList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	err := a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Groups", req, &groups)
	return groups, err
}

// PatchR ...
func (a GroupsAPI) PatchR(groupID string, r patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch, fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), r, nil)
}

// Patch applys a patch request for a group given a path attribute
func (a GroupsAPI) Patch(groupID string, addList []string, removeList []string, path GroupPathType) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	var addOperations GroupPatchOperations
	var removeOperations GroupPatchOperations

	groupPatchRequest := GroupPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []GroupPatchOperations{},
	}

	if addList == nil && removeList == nil {
		return errors.New("empty members list to add or to remove")
	}

	if len(addList) > 0 {
		addOperations = GroupPatchOperations{
			Op:    "add",
			Path:  path,
			Value: []ValueListItem{},
		}
		for _, addItem := range addList {
			addOperations.Value = append(addOperations.Value, ValueListItem{Value: addItem})
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, addOperations)
	}

	for _, removeItem := range removeList {
		path := fmt.Sprintf("%s[value eq \"%s\"]", string(path), removeItem)
		removeOperations = GroupPatchOperations{
			Op:   "remove",
			Path: GroupPathType(path),
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, removeOperations)
	}
	return a.client.Scim(a.context, http.MethodPatch, groupPath, groupPatchRequest, nil)
}

// Delete deletes a group given a group id
func (a GroupsAPI) Delete(groupID string) error {
	return a.client.Scim(a.context, http.MethodDelete,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID),
		nil, nil)
}
