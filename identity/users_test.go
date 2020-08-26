package identity

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestScimUserAPI_Create(t *testing.T) {
	type args struct {
		Schemas      []URN                  `json:"schemas,omitempty"`
		UserName     string                 `json:"userName,omitempty"`
		Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                 `json:"displayName,omitempty"`
		Roles        []RoleListItem         `json:"roles,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		want           User
		wantErr        bool
	}{
		{
			name: "Create Test",
			response: `{
							"id": "101030",
							"userName": "test.user@databricks.com"
						}`,
			responseStatus: http.StatusOK,
			want: User{
				ID:       "101030",
				UserName: "test.user@databricks.com",
			},
			args: args{
				Schemas:      []URN{UserSchema},
				UserName:     "test.user@databricks.com",
				Entitlements: []EntitlementsListItem{{Value: "test-entitlement"}},
				DisplayName:  "test user",
				Roles:        []RoleListItem{{Value: "test-role"}},
			},
			wantErr: false,
		},
		{
			name:           "Create Test Failure",
			response:       ``,
			want:           User{},
			responseStatus: http.StatusBadRequest,
			args: args{
				Schemas:      []URN{UserSchema},
				UserName:     "test.user@databricks.com",
				Entitlements: []EntitlementsListItem{{Value: "test-entitlement"}},
				DisplayName:  "test user",
				Roles:        []RoleListItem{{Value: "test-role"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/preview/scim/v2/Users", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewUsersAPI(client).Create(tt.args.UserName, tt.args.DisplayName, []string{string(tt.args.Entitlements[0].Value)}, []string{tt.args.Roles[0].Value})
			})
		})
	}
}

func TestScimUserAPI_Update(t *testing.T) {
	type args struct {
		Schemas      []URN                  `json:"schemas,omitempty"`
		UserName     string                 `json:"userName,omitempty"`
		Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                 `json:"displayName,omitempty"`
		Roles        []RoleListItem         `json:"roles,omitempty"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		args           []interface{}
		wantURI        []string
		wantErr        bool
	}{
		{
			name: "Update Test",
			response: []string{`{
									"id": "101030",
									"userName": "test.user@databricks.com"
								}`,
				"",
			},
			responseStatus: []int{http.StatusOK, http.StatusOK},
			wantURI:        []string{"/api/2.0/preview/scim/v2/Users/101030", "/api/2.0/preview/scim/v2/Users/101030"},
			args: []interface{}{
				nil,
				&args{
					Schemas:      []URN{UserSchema},
					UserName:     "test.user@databricks.com",
					Entitlements: []EntitlementsListItem{{Value: "test-entitlement"}},
					DisplayName:  "test user",
					Roles:        []RoleListItem{{Value: "test-role"}},
				},
			},
			wantErr: false,
		},
		{
			name: "Update failure test",
			response: []string{``,
				"",
			},
			responseStatus: []int{http.StatusBadRequest, http.StatusOK},
			wantURI:        []string{"/api/2.0/preview/scim/v2/Users/101030", "/api/2.0/preview/scim/v2/Users/101030"},
			args: []interface{}{
				nil,
				&args{
					Schemas:      []URN{UserSchema},
					UserName:     "test.user@databricks.com",
					Entitlements: []EntitlementsListItem{{Value: "test-entitlement"}},
					DisplayName:  "test user",
					Roles:        []RoleListItem{{Value: "test-role"}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertMultipleRequestsWithMockServer(t, tt.args,
				[]string{http.MethodGet, http.MethodPut}, tt.wantURI,
				[]interface{}{nil, &args{}}, tt.response, tt.responseStatus,
				nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
					return nil, NewUsersAPI(client).Update("101030", tt.args[1].(*args).UserName,
						tt.args[1].(*args).DisplayName,
						[]string{string(tt.args[1].(*args).Entitlements[0].Value)},
						[]string{tt.args[1].(*args).Roles[0].Value})
				})
		})
	}
}

func TestScimUserAPI_Delete(t *testing.T) {
	type args struct {
		UserID string `json:"user_id,omitempty"`
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
				UserID: "10030",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    false,
		},
		{
			name:           "Delete failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				UserID: "10030",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodDelete, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewUsersAPI(client).Delete(tt.args.UserID)
			})
		})
	}
}

func TestScimUserAPI_SetUserAsAdmin(t *testing.T) {
	type args struct {
		UserID       string `json:"user_id,omitempty"`
		AdminGroupID string `json:"admin_group_id,omitempty"`
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
			name:           "SetUserAsAdmin test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    false,
		},
		{
			name:           "SetUserAsAdmin failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input UserPatchRequest
			expectedPatchRequest := UserPatchRequest{
				Schemas: []URN{PatchOp},
				Operations: []UserPatchOperations{
					{
						Op:    "add",
						Value: &GroupsValue{Groups: []ValueListItem{{Value: tt.args.AdminGroupID}}},
					},
				},
			}
			qa.AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewUsersAPI(client).SetUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
			})
		})
	}
}

func TestScimUserAPI_VerifyUserAsAdmin(t *testing.T) {
	type args struct {
		UserID       string `json:"user_id,omitempty"`
		AdminGroupID string `json:"admin_group_id,omitempty"`
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
			name: "VerifyUserAsAdmin true test",
			response: `{
					"groups":[
						{
							"display":"admins",
							"value":"100002",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/100002"
						},
						{
							"display":"test-create-group",
							"value":"101355",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/101355"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-instance-profile"
						}
					],
					"id":"101030",
					"userName":"test.user@databricks.com",
					"displayName":"test.user@databricks.com"
				}`,
			responseStatus: http.StatusOK,
			args: args{
				UserID:       "10030",
				AdminGroupID: "100002",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       true,
			wantErr:    false,
		},
		{
			name: "VerifyUserAsAdmin false test",
			response: `{
					"groups":[
						{
							"display":"admins",
							"value":"100052",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/100002"
						},
						{
							"display":"test-create-group",
							"value":"101355",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/101355"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-instance-profile"
						}
					],
					"id":"101030",
					"userName":"test.user@databricks.com",
					"displayName":"test.user@databricks.com"
				}`,
			responseStatus: http.StatusOK,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       false,
			wantErr:    false,
		},
		{
			name:           "VerifyUserAsAdmin failure test",
			response:       "{}",
			responseStatus: http.StatusBadRequest,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input UserPatchRequest
			expectedPatchRequest := UserPatchRequest{
				Schemas: []URN{PatchOp},
				Operations: []UserPatchOperations{
					{
						Op:    "add",
						Value: &GroupsValue{Groups: []ValueListItem{{Value: tt.args.AdminGroupID}}},
					},
				},
			}
			qa.AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodGet, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewUsersAPI(client).VerifyUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
			})
		})
	}
}

func TestScimUserAPI_RemoveUserAsAdmin(t *testing.T) {
	type args struct {
		UserID       string `json:"user_id,omitempty"`
		AdminGroupID string `json:"admin_group_id,omitempty"`
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
			name:           "RemoveUserAsAdmin test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    false,
		},
		{
			name:           "RemoveUserAsAdmin failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				UserID:       "10030",
				AdminGroupID: "10000",
			},
			requestURI: "/api/2.0/preview/scim/v2/Users/10030",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input UserPatchRequest
			expectedPatchRequest := UserPatchRequest{
				Schemas: []URN{PatchOp},
				Operations: []UserPatchOperations{
					{
						Op:   "remove",
						Path: fmt.Sprintf("groups[value eq \"%s\"]", tt.args.AdminGroupID),
					},
				},
			}
			qa.AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewUsersAPI(client).RemoveUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
			})
		})
	}
}

func TestScimUserAPI_Read(t *testing.T) {
	type args struct {
		UserID string `json:"user_id"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		args           []args
		wantURI        []string
		want           User
		wantErr        bool
	}{
		{
			name: "Read test",
			response: []string{`{
					"groups":[
						{
							"display":"admins",
							"value":"100002",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/100002"
						},
						{
							"display":"test-create-group",
							"value":"101355",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/101355"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-instance-profile"
						}
					],
					"id":"101030",
					"userName":"test.user@databricks.com",
					"displayName":"test.user@databricks.com"
				}`,
				`{
					"schemas":[
						"urn:ietf:params:scim:schemas:core:2.0:Group"
					],
					"id":"100002",
					"displayName":"admins",
					"members":[
						{
							"value":"100000"
						},
						{
							"value":"100001"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-inherited-profile1"
						}
					]
				}`,
				`{
					"schemas":[
						"urn:ietf:params:scim:schemas:core:2.0:Group"
					],
					"id":"101355",
					"displayName":"test-create-group",
					"members":[
						{
							"value":"100000"
						},
						{
							"value":"100001"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-inherited-profile2"
						}
					]
				}`,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK},
			args: []args{
				{
					UserID: "101030",
				},
			},
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030", "/api/2.0/preview/scim/v2/Groups/100002", "/api/2.0/preview/scim/v2/Groups/101355"},
			want: User{
				ID:          "101030",
				DisplayName: "test.user@databricks.com",
				UserName:    "test.user@databricks.com",
				Groups: []GroupsListItem{
					{
						Value: "100002",
					},
					{
						Value: "101355",
					},
				},
				Roles: []RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
					},
				},
				UnInheritedRoles: []RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
					},
				},
				InheritedRoles: []RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-inherited-profile1",
					},
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-inherited-profile2",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Read user failure",
			response: []string{``,
				``,
			},
			responseStatus: []int{http.StatusBadRequest},
			args: []args{
				{
					UserID: "101030",
				},
			},
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030"},
			want:    User{},
			wantErr: true,
		},
		{
			name: "Read user unmarshal failure",
			response: []string{
				`{`, //``,
			},
			responseStatus: []int{http.StatusOK},
			args: []args{
				{
					UserID: "101030",
				},
			},
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030"},
			want:    User{},
			wantErr: true,
		},
		{
			name: "Read user first group failure no inherited and non inherited roles",
			response: []string{`{
					"groups":[
						{
							"display":"admins",
							"value":"100002",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/100002"
						},
						{
							"display":"test-create-group",
							"value":"101355",
							"$ref":"https://test.databricks.com/api/2.0/scim/v2/Groups/101355"
						}
					],
					"roles":[
						{
							"value":"arn:aws:iam::1231231123123:instance-profile/my-instance-profile"
						}
					],
					"id":"101030",
					"userName":"test.user@databricks.com",
					"displayName":"test.user@databricks.com"
				}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest},
			args: []args{
				{
					UserID: "101030",
				},
			},
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030", "/api/2.0/preview/scim/v2/Groups/100002"},
			want: User{
				ID:          "101030",
				DisplayName: "test.user@databricks.com",
				UserName:    "test.user@databricks.com",
				Groups: []GroupsListItem{
					{
						Value: "100002",
					},
					{
						Value: "101355",
					},
				},
				Roles: []RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet, http.MethodGet}, tt.wantURI, []args{input}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewUsersAPI(client).Read(tt.args[0].UserID)
			})
		})
	}
}

func TestAccCreateUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()

	user, err := NewUsersAPI(client).Create("testuser@databricks.com", "Display Name", nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := NewUsersAPI(client).Delete(idToDelete)
		assert.NoError(t, err, err)
	}()

	user, err = NewUsersAPI(client).Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	err = NewUsersAPI(client).Update(user.ID, "newtestuser@databricks.com", "Test User", []string{string(AllowClusterCreateEntitlement)}, nil)
	//t.Log(user)
	assert.NoError(t, err, err)
}

func TestAccCreateAdminUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	user, err := NewUsersAPI(client).Create(
		fmt.Sprintf("terraform+%s@databricks.com", randomName),
		"Terra "+randomName, nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := NewUsersAPI(client).Delete(idToDelete)
		assert.NoError(t, err, err)
	}()
	log.Println(idToDelete)

	user, err = NewUsersAPI(client).Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	group, err := NewGroupsAPI(client).GetAdminGroup()
	assert.NoError(t, err, err)

	adminGroupID := group.ID

	err = NewUsersAPI(client).SetUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err := NewUsersAPI(client).VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == true)
	log.Println(userIsAdmin)

	err = NewUsersAPI(client).RemoveUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err = NewUsersAPI(client).VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == false)
	log.Println(userIsAdmin)
}

func TestAccRoleDifferences(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	user, err := NewUsersAPI(client).Create(
		fmt.Sprintf("terraform+%s@databricks.com", randomName),
		"Terra "+randomName, nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID

	user, err = NewUsersAPI(client).Read(idToDelete)
	assert.NoError(t, err, err)
	t.Log(user.Roles)
	t.Log(user.Groups)
	t.Log(user.InheritedRoles)
	t.Log(user.UnInheritedRoles)

	err = NewUsersAPI(client).Delete(idToDelete)
	assert.NoError(t, err, err)
}
