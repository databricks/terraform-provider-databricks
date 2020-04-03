package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"net/http"
)

// TokensAPI exposes the Secrets API
type GroupsAPI struct {
	Client DBApiClient
}

func (a GroupsAPI) init(client DBApiClient) GroupsAPI {
	a.Client = client
	return a
}

func (a GroupsAPI) Create(groupName string, members []string) (model.Group, error) {
	var group model.Group

	scimGroupRequest := struct {
		Schemas     []model.URN            `json:"schemas,omitempty"`
		DisplayName string                 `json:"displayName,omitempty"`
		Members     []model.MemberListItem `json:"members,omitempty"`
	}{}
	scimGroupRequest.Schemas = []model.URN{model.GroupSchema}
	scimGroupRequest.DisplayName = groupName

	scimGroupRequest.Members = []model.MemberListItem{}
	for _, member := range members {
		scimGroupRequest.Members = append(scimGroupRequest.Members, model.MemberListItem{Value: member})
	}

	resp, err := a.Client.performQuery(http.MethodPost, "/preview/scim/v2/Groups", "2.0", scimHeaders, scimGroupRequest)
	if err != nil {
		return group, err
	}

	err = json.Unmarshal(resp, &group)
	return group, err
}

func (a GroupsAPI) Read(groupID string) (model.Group, error) {
	var group model.Group
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	resp, err := a.Client.performQuery(http.MethodGet, groupPath, "2.0", scimHeaders, nil)
	if err != nil {
		return group, err
	}

	err = json.Unmarshal(resp, &group)
	return group, err
}

func (a GroupsAPI) Update(groupID string, addMembers []string, removeMembers []string) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	var addOperations model.GroupPatchOperations
	var removeOperations model.GroupPatchOperations

	groupPatchRequest := model.GroupPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.GroupPatchOperations{},
	}

	if addMembers == nil && removeMembers == nil {
		return errors.New("Empty members list to add or to remove.")
	}

	if len(addMembers) > 0 {
		addOperations = model.GroupPatchOperations{
			Op: "add",
			Value: &model.MembersValue{
				Members: []model.MemberListItem{},
			},
		}
		for _, addMember := range addMembers {
			addOperations.Value.Members = append(addOperations.Value.Members, model.MemberListItem{Value: addMember})
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, addOperations)
	}

	for _, removeMember := range removeMembers {
		path := fmt.Sprintf("members[value eq \"%s\"]", removeMember)
		removeOperations = model.GroupPatchOperations{
			Op:   "remove",
			Path: path,
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, removeOperations)
	}

	_, err := a.Client.performQuery(http.MethodPatch, groupPath, "2.0", scimHeaders, groupPatchRequest)

	return err
}

func (a GroupsAPI) Delete(groupID string) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)
	_, err := a.Client.performQuery(http.MethodDelete, groupPath, "2.0", scimHeaders, nil)
	return err
}
