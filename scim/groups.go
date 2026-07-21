package scim

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/databricks/terraform-provider-databricks/common"
)

// NewGroupsAPI creates GroupsAPI instance from provider meta.
// apiLevel controls whether account-level or workspace-level SCIM endpoints are used.
// Pass "" to infer from the provider host.
func NewGroupsAPI(ctx context.Context, m any, apiLevel string) GroupsAPI {
	return GroupsAPI{
		client:   m.(*common.DatabricksClient),
		context:  ctx,
		ApiLevel: apiLevel,
	}
}

// GroupsAPI exposes the scim groups API
type GroupsAPI struct {
	client   *common.DatabricksClient
	context  context.Context
	ApiLevel string
}

// Create creates a scim group in the Databricks workspace
func (a GroupsAPI) Create(scimGroupRequest Group) (group Group, err error) {
	scimGroupRequest.Schemas = []URN{GroupSchema}
	err = a.client.Scim(a.context, http.MethodPost, "/preview/scim/v2/Groups", scimGroupRequest, &group, a.ApiLevel)
	return
}

// Read reads and returns a Group object via SCIM api
func (a GroupsAPI) Read(groupID, attributes string) (group Group, err error) {
	key := fmt.Sprintf(
		"/preview/scim/v2/Groups/%v?attributes=%s", groupID, attributes)
	err = a.client.Scim(a.context, http.MethodGet, key, nil, &group, a.ApiLevel)
	return
}

// Filter returns groups matching the filter
func (a GroupsAPI) Filter(filter string, attributes string) (GroupList, error) {
	var groups GroupList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	if attributes != "" {
		req["attributes"] = attributes
	}
	err := a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Groups", req, &groups, a.ApiLevel)
	return groups, err
}

// ListAll retrieves all groups with the given attributes, handling SCIM pagination.
func (a GroupsAPI) ListAll(attributes string) ([]Group, error) {
	startIndex := 1
	var result []Group
	for {
		req := map[string]string{
			"count":      "10000",
			"startIndex": strconv.Itoa(startIndex),
		}
		if attributes != "" {
			req["attributes"] = attributes
		}
		var page GroupList
		err := a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Groups", req, &page, a.ApiLevel)
		if err != nil {
			return nil, err
		}
		result = append(result, page.Resources...)
		if len(page.Resources) == 0 || len(result) >= int(page.TotalResults) {
			break
		}
		startIndex += len(page.Resources)
	}
	return result, nil
}

func (a GroupsAPI) ReadByDisplayName(displayName, attributes string) (group Group, err error) {
	groupList, err := a.Filter(fmt.Sprintf(`displayName eq "%s"`, displayName), attributes)
	if err != nil {
		return
	}
	if len(groupList.Resources) == 0 {
		err = fmt.Errorf("cannot find group: %s", displayName)
		return
	}
	group = groupList.Resources[0]
	return
}

func (a GroupsAPI) Patch(groupID string, r patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch, fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), r, nil, a.ApiLevel)
}

func (a GroupsAPI) UpdateNameAndEntitlements(groupID string, name string, externalID string, e entitlements) error {
	g, err := a.Read(groupID, "displayName,entitlements,groups,members,externalId")
	if err != nil {
		return err
	}
	return a.client.Scim(a.context, http.MethodPut,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID),
		Group{
			DisplayName:  name,
			Entitlements: e,
			Groups:       g.Groups,
			Roles:        g.Roles,
			Members:      g.Members,
			Schemas:      []URN{GroupSchema},
			ExternalID:   externalID,
		}, nil, a.ApiLevel)
}

func (a GroupsAPI) UpdateEntitlements(groupID string, entitlements patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), entitlements, nil, a.ApiLevel)
}

// Delete deletes a group given a group id
func (a GroupsAPI) Delete(groupID string) error {
	return a.client.Scim(a.context, http.MethodDelete,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID),
		nil, nil, a.ApiLevel)
}
