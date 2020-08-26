package identity

import (
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestScimGroupAPI_Create(t *testing.T) {
	type args struct {
		Schemas      []URN           `json:"schemas,omitempty"`
		DisplayName  string          `json:"displayName,omitempty"`
		Members      []ValueListItem `json:"members,omitempty"`
		Entitlements []ValueListItem `json:"entitlements,omitempty"`
		Roles        []ValueListItem `json:"roles,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		want           Group
		wantErr        bool
	}{
		{
			name: "Create Test",
			response: `{
							"id": "101030"
						}`,
			responseStatus: http.StatusOK,
			want: Group{
				ID: "101030",
			},
			args: args{
				Schemas:      []URN{GroupSchema},
				Entitlements: []ValueListItem{{Value: "test-entitlement"}},
				DisplayName:  "test group",
				Roles:        []ValueListItem{{Value: "test-role"}},
				Members:      []ValueListItem{{Value: "test-member"}},
			},
			wantErr: false,
		},
		{
			name:           "Create Test Failure",
			response:       ``,
			want:           Group{},
			responseStatus: http.StatusBadRequest,
			args: args{
				Schemas:      []URN{GroupSchema},
				Entitlements: []ValueListItem{{Value: "test-entitlement"}},
				DisplayName:  "test group",
				Roles:        []ValueListItem{{Value: "test-role"}},
				Members:      []ValueListItem{{Value: "test-member"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/preview/scim/v2/Groups", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewGroupsAPI(client).Create(tt.args.DisplayName, []string{tt.args.Members[0].Value}, []string{tt.args.Roles[0].Value}, []string{tt.args.Entitlements[0].Value})
			})
		})
	}
}

func TestScimGroupAPI_GetAdminGroup(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		responseStatus int
		want           Group
		wantErr        bool
	}{
		{
			name: "GetAdminGroup Test",
			response: `{
							"resources": [{"id": "101030"}]
						}`,
			responseStatus: http.StatusOK,
			want: Group{
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
			want: Group{
				ID: "101030",
			},
			wantErr: true,
		},
		{
			name:           "GetAdminGroup Test Failure",
			response:       ``,
			want:           Group{},
			responseStatus: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, nil, http.MethodGet, "/api/2.0/preview/scim/v2/Groups?filter=displayName+eq+admins", nil, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewGroupsAPI(client).GetAdminGroup()
			})
		})
	}
}

func TestScimGroupAPI_Patch(t *testing.T) {
	type args GroupPatchRequest
	tests := []struct {
		name   string
		params struct {
			groupID    string
			addList    []string
			removeList []string
			path       GroupPathType
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
				path       GroupPathType
			}{groupID: "my-group-id", addList: []string{"100"}, removeList: []string{"200"}, path: GroupMembersPath},
			response: ``,
			args: args{
				Schemas: []URN{PatchOp},
				Operations: []GroupPatchOperations{
					{
						Op:    "add",
						Path:  GroupMembersPath,
						Value: []ValueListItem{{Value: "100"}},
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
				Schemas: []URN{PatchOp},
				Operations: []GroupPatchOperations{
					{
						Op:    "add",
						Path:  GroupMembersPath,
						Value: []ValueListItem{{Value: "100"}},
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
				path       GroupPathType
			}{groupID: "my-group-id", addList: []string{"100"}, removeList: []string{"200"}, path: GroupMembersPath},
			requestURI:     "/api/2.0/preview/scim/v2/Groups/my-group-id",
			responseStatus: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		var input args
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewGroupsAPI(client).Patch(tt.params.groupID, tt.params.addList, tt.params.removeList, tt.params.path)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodDelete, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewGroupsAPI(client).Delete(tt.args.GroupID)
			})
		})
	}
}

func TestAccGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	user, err := NewUsersAPI(client).Create("test-acc@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	user2, err := NewUsersAPI(client).Create("test-acc2@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	//Create empty group
	group, err := NewGroupsAPI(client).Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(group.ID)
		assert.NoError(t, err, err)
		err = NewUsersAPI(client).Delete(user.ID)
		assert.NoError(t, err, err)
		err = NewUsersAPI(client).Delete(user2.ID)
		assert.NoError(t, err, err)
	}()

	group, err = NewGroupsAPI(client).Read(group.ID)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(group.ID, []string{user.ID, user2.ID}, nil, GroupMembersPath)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(group.ID, nil, []string{user.ID}, GroupMembersPath)
	assert.NoError(t, err, err)

	group, err = NewGroupsAPI(client).Read(group.ID)
	assert.NoError(t, err, err)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}

func TestAccGetAdminGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	grp, err := NewGroupsAPI(client).GetAdminGroup()
	assert.NoError(t, err, err)
	assert.NotNil(t, grp)
	assert.True(t, len(grp.ID) > 0)
}

func TestAwsAccReadInheritedRolesFromGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	// TODO: pass IAM role with ENV variable
	myTestRole := "arn:aws:iam::123456789012:instance-profile/go-sdk-integeration-testing"
	err := NewInstanceProfilesAPI(client).Create(myTestRole, true)
	assert.NoError(t, err, err)
	defer func() {
		err := NewInstanceProfilesAPI(client).Delete(myTestRole)
		assert.NoError(t, err, err)
	}()

	myTestGroup, err := NewGroupsAPI(client).Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(myTestGroup.ID)
		assert.NoError(t, err, err)
	}()

	myTestSubGroup, err := NewGroupsAPI(client).Create("my-test-sub-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(myTestSubGroup.ID)
		assert.NoError(t, err, err)
	}()

	err = NewGroupsAPI(client).Patch(myTestGroup.ID, []string{myTestRole}, nil, GroupRolesPath)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(myTestGroup.ID, []string{myTestSubGroup.ID}, nil, GroupMembersPath)
	assert.NoError(t, err, err)

	myTestGroupInfo, err := NewGroupsAPI(client).Read(myTestSubGroup.ID)
	assert.NoError(t, err, err)

	assert.True(t, len(myTestGroupInfo.InheritedRoles) > 0)
	assert.True(t, func(roles []RoleListItem, testRole string) bool {
		for _, role := range roles {
			if role.Value == testRole {
				return true
			}
		}
		return false
	}(myTestGroupInfo.InheritedRoles, myTestRole))
}
