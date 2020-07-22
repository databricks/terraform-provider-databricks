package service

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestScimGroupAPI_Create(t *testing.T) {
	type args struct {
		Schemas      []model.URN           `json:"schemas,omitempty"`
		DisplayName  string                `json:"displayName,omitempty"`
		Members      []model.ValueListItem `json:"members,omitempty"`
		Entitlements []model.ValueListItem `json:"entitlements,omitempty"`
		Roles        []model.ValueListItem `json:"roles,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		want           model.Group
		wantErr        bool
	}{
		{
			name: "Create Test",
			response: `{
							"id": "101030"
						}`,
			responseStatus: http.StatusOK,
			want: model.Group{
				ID: "101030",
			},
			args: args{
				Schemas:      []model.URN{model.GroupSchema},
				Entitlements: []model.ValueListItem{{Value: "test-entitlement"}},
				DisplayName:  "test group",
				Roles:        []model.ValueListItem{{Value: "test-role"}},
				Members:      []model.ValueListItem{{Value: "test-member"}},
			},
			wantErr: false,
		},
		{
			name:           "Create Test Failure",
			response:       ``,
			want:           model.Group{},
			responseStatus: http.StatusBadRequest,
			args: args{
				Schemas:      []model.URN{model.GroupSchema},
				Entitlements: []model.ValueListItem{{Value: "test-entitlement"}},
				DisplayName:  "test group",
				Roles:        []model.ValueListItem{{Value: "test-role"}},
				Members:      []model.ValueListItem{{Value: "test-member"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/preview/scim/v2/Groups", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Groups().Create(tt.args.DisplayName, []string{tt.args.Members[0].Value}, []string{tt.args.Roles[0].Value}, []string{tt.args.Entitlements[0].Value})
			})
		})
	}
}

func TestScimGroupAPI_GetAdminGroup(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		responseStatus int
		want           model.Group
		wantErr        bool
	}{
		{
			name: "GetAdminGroup Test",
			response: `{
							"resources": [{"id": "101030"}]
						}`,
			responseStatus: http.StatusOK,
			want: model.Group{
				ID: "101030",
			},
			wantErr: false,
		},
		{
			name: "GetAdminGroup no admin failure Test",
			response: `{
							"resources": []
						}`,
			responseStatus: http.StatusOK,
			want: model.Group{
				ID: "101030",
			},
			wantErr: true,
		},
		{
			name:           "GetAdminGroup Test Failure",
			response:       ``,
			want:           model.Group{},
			responseStatus: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, "/api/2.0/preview/scim/v2/Groups?filter=displayName+eq+admins", nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return client.Groups().GetAdminGroup()
			})
		})
	}
}

func TestScimGroupAPI_Patch(t *testing.T) {
	type args model.GroupPatchRequest
	tests := []struct {
		name   string
		params struct {
			groupID    string
			addList    []string
			removeList []string
			path       model.GroupPathType
		}
		response       string
		args           args
		requestURI     string
		responseStatus int
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Patch Test",
			params: struct {
				groupID    string
				addList    []string
				removeList []string
				path       model.GroupPathType
			}{groupID: "my-group-id", addList: []string{"100"}, removeList: []string{"200"}, path: model.GroupMembersPath},
			response: ``,
			args: args{
				Schemas: []model.URN{model.PatchOp},
				Operations: []model.GroupPatchOperations{
					{
						Op:    "add",
						Path:  model.GroupMembersPath,
						Value: []model.ValueListItem{{Value: "100"}},
					},
					{
						Op:   "remove",
						Path: "members[value eq \"200\"]",
					},
				},
			},
			requestURI:     "/api/2.0/preview/scim/v2/Groups/my-group-id",
			responseStatus: http.StatusOK,
			want:           nil,
			wantErr:        false,
		},
		{
			name:     "Patch Test Failure",
			response: ``,
			want:     nil,
			args: args{
				Schemas: []model.URN{model.PatchOp},
				Operations: []model.GroupPatchOperations{
					{
						Op:    "add",
						Path:  model.GroupMembersPath,
						Value: []model.ValueListItem{{Value: "100"}},
					},
					{
						Op:   "remove",
						Path: "members[value eq \"200\"]",
					},
				},
			},
			params: struct {
				groupID    string
				addList    []string
				removeList []string
				path       model.GroupPathType
			}{groupID: "my-group-id", addList: []string{"100"}, removeList: []string{"200"}, path: model.GroupMembersPath},
			requestURI:     "/api/2.0/preview/scim/v2/Groups/my-group-id",
			responseStatus: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		var input args
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, &tt.args, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.Groups().Patch(tt.params.groupID, tt.params.addList, tt.params.removeList, tt.params.path)
			})
		})
	}
}

func TestScimGroupAPI_Delete(t *testing.T) {
	type args struct {
		GroupID string `json:"user_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		requestURI     string
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name:           "Delete test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				GroupID: "10030",
			},
			requestURI: "/api/2.0/preview/scim/v2/Groups/10030",
			want:       nil,
			wantErr:    false,
		},
		{
			name:           "Delete failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				GroupID: "10030",
			},
			requestURI: "/api/2.0/preview/scim/v2/Groups/10030",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodDelete, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
				return nil, client.Groups().Delete(tt.args.GroupID)
			})
		})
	}
}

func TestFmt(t *testing.T) {
	str := fmt.Sprintf("members[value eq \"%s\"]", "1000")
	t.Log(str)
}

func TestAccGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := GetIntegrationDBAPIClient()

	user, err := client.Users().Create("test-acc@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	user2, err := client.Users().Create("test-acc2@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	//Create empty group
	group, err := client.Groups().Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(group.ID)
		assert.NoError(t, err, err)
		err = client.Users().Delete(user.ID)
		assert.NoError(t, err, err)
		err = client.Users().Delete(user2.ID)
		assert.NoError(t, err, err)
	}()

	group, err = client.Groups().Read(group.ID)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(group.ID, []string{user.ID, user2.ID}, nil, model.GroupMembersPath)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(group.ID, nil, []string{user.ID}, model.GroupMembersPath)
	assert.NoError(t, err, err)

	group, err = client.Groups().Read(group.ID)
	assert.NoError(t, err, err)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}

func TestAccGetAdminGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	grp, err := client.Groups().GetAdminGroup()
	assert.NoError(t, err, err)
	assert.NotNil(t, grp)
	assert.True(t, len(grp.ID) > 0)
}

func TestAccReadInheritedRolesFromGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	myTestRole := "arn:aws:iam::123456789012:instance-profile/go-sdk-integeration-testing"
	err := client.InstanceProfiles().Create(myTestRole, true)
	assert.NoError(t, err, err)
	defer func() {
		err := client.InstanceProfiles().Delete(myTestRole)
		assert.NoError(t, err, err)
	}()

	myTestGroup, err := client.Groups().Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(myTestGroup.ID)
		assert.NoError(t, err, err)
	}()

	myTestSubGroup, err := client.Groups().Create("my-test-sub-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(myTestSubGroup.ID)
		assert.NoError(t, err, err)
	}()

	err = client.Groups().Patch(myTestGroup.ID, []string{myTestRole}, nil, model.GroupRolesPath)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(myTestGroup.ID, []string{myTestSubGroup.ID}, nil, model.GroupMembersPath)
	assert.NoError(t, err, err)

	myTestGroupInfo, err := client.Groups().Read(myTestSubGroup.ID)
	assert.NoError(t, err, err)

	assert.True(t, len(myTestGroupInfo.InheritedRoles) > 0)
	assert.True(t, func(roles []model.RoleListItem, testRole string) bool {
		for _, role := range roles {
			if role.Value == testRole {
				return true
			}
		}
		return false
	}(myTestGroupInfo.InheritedRoles, myTestRole))
}
