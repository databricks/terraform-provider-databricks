package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestScimUserAPI_Create(t *testing.T) {
	type args struct {
		Schemas      []model.URN                  `json:"schemas,omitempty"`
		UserName     string                       `json:"userName,omitempty"`
		Entitlements []model.EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                       `json:"displayName,omitempty"`
		Roles        []model.RoleListItem         `json:"roles,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		want           model.User
		wantErr        bool
	}{
		{
			name: "Create Test",
			response: `{
							"id": "101030",
							"userName": "test.user@databricks.com"
						}`,
			responseStatus: http.StatusOK,
			want: model.User{
				ID:       "101030",
				UserName: "test.user@databricks.com",
			},
			args: args{
				Schemas:      []model.URN{model.UserSchema},
				UserName:     "test.user@databricks.com",
				Entitlements: []model.EntitlementsListItem{{Value: "test-entitlement"}},
				DisplayName:  "test user",
				Roles:        []model.RoleListItem{{Value: "test-role"}},
			},
			wantErr: false,
		},
		{
			name:           "Create Test Failure",
			response:       ``,
			want:           model.User{},
			responseStatus: http.StatusBadRequest,
			args: args{
				Schemas:      []model.URN{model.UserSchema},
				UserName:     "test.user@databricks.com",
				Entitlements: []model.EntitlementsListItem{{Value: "test-entitlement"}},
				DisplayName:  "test user",
				Roles:        []model.RoleListItem{{Value: "test-role"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/preview/scim/v2/Users", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Users().Create(tt.args.UserName, tt.args.DisplayName, []string{string(tt.args.Entitlements[0].Value)}, []string{tt.args.Roles[0].Value})
			})
		})
	}
}

func TestScimUserAPI_Update(t *testing.T) {
	type args struct {
		Schemas      []model.URN                  `json:"schemas,omitempty"`
		UserName     string                       `json:"userName,omitempty"`
		Entitlements []model.EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                       `json:"displayName,omitempty"`
		Roles        []model.RoleListItem         `json:"roles,omitempty"`
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
			wantURI:        []string{"/api/2.0/preview/scim/v2/Users/101030?", "/api/2.0/preview/scim/v2/Users/101030"},
			args: []interface{}{
				nil,
				&args{
					Schemas:      []model.URN{model.UserSchema},
					UserName:     "test.user@databricks.com",
					Entitlements: []model.EntitlementsListItem{{Value: "test-entitlement"}},
					DisplayName:  "test user",
					Roles:        []model.RoleListItem{{Value: "test-role"}},
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
			wantURI:        []string{"/api/2.0/preview/scim/v2/Users/101030?", "/api/2.0/preview/scim/v2/Users/101030"},
			args: []interface{}{
				nil,
				&args{
					Schemas:      []model.URN{model.UserSchema},
					UserName:     "test.user@databricks.com",
					Entitlements: []model.EntitlementsListItem{{Value: "test-entitlement"}},
					DisplayName:  "test user",
					Roles:        []model.RoleListItem{{Value: "test-role"}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodPut}, tt.wantURI, []interface{}{nil, &args{}}, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Users().Update("101030", tt.args[1].(*args).UserName, tt.args[1].(*args).DisplayName, []string{string(tt.args[1].(*args).Entitlements[0].Value)}, []string{tt.args[1].(*args).Roles[0].Value})
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodDelete, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Users().Delete(tt.args.UserID)
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
			var input model.UserPatchRequest
			expectedPatchRequest := model.UserPatchRequest{
				Schemas: []model.URN{model.PatchOp},
				Operations: []model.UserPatchOperations{
					{
						Op:    "add",
						Value: &model.GroupsValue{Groups: []model.ValueListItem{{Value: tt.args.AdminGroupID}}},
					},
				},
			}
			AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Users().SetUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
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
			requestURI: "/api/2.0/preview/scim/v2/Users/10030?",
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
			requestURI: "/api/2.0/preview/scim/v2/Users/10030?",
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
			requestURI: "/api/2.0/preview/scim/v2/Users/10030?",
			want:       false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input model.UserPatchRequest
			expectedPatchRequest := model.UserPatchRequest{
				Schemas: []model.URN{model.PatchOp},
				Operations: []model.UserPatchOperations{
					{
						Op:    "add",
						Value: &model.GroupsValue{Groups: []model.ValueListItem{{Value: tt.args.AdminGroupID}}},
					},
				},
			}
			AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodGet, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Users().VerifyUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
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
			var input model.UserPatchRequest
			expectedPatchRequest := model.UserPatchRequest{
				Schemas: []model.URN{model.PatchOp},
				Operations: []model.UserPatchOperations{
					{
						Op:   "remove",
						Path: fmt.Sprintf("groups[value eq \"%s\"]", tt.args.AdminGroupID),
					},
				},
			}
			AssertRequestWithMockServer(t, &expectedPatchRequest, http.MethodPatch, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Users().RemoveUserAsAdmin(tt.args.UserID, tt.args.AdminGroupID)
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
		want           model.User
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
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030?", "/api/2.0/preview/scim/v2/Groups/100002?", "/api/2.0/preview/scim/v2/Groups/101355?"},
			want: model.User{
				ID:          "101030",
				DisplayName: "test.user@databricks.com",
				UserName:    "test.user@databricks.com",
				Groups: []model.GroupsListItem{
					{
						Value: "100002",
					},
					{
						Value: "101355",
					},
				},
				Roles: []model.RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
					},
				},
				UnInheritedRoles: []model.RoleListItem{
					{
						Value: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
					},
				},
				InheritedRoles: []model.RoleListItem{
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
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030?"},
			want:    model.User{},
			wantErr: true,
		},
		{
			name: "Read user unmarshal failure",
			response: []string{``,
				`{`,
			},
			responseStatus: []int{http.StatusOK},
			args: []args{
				{
					UserID: "101030",
				},
			},
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030?"},
			want:    model.User{},
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
			wantURI: []string{"/api/2.0/preview/scim/v2/Users/101030?", "/api/2.0/preview/scim/v2/Groups/100002?"},
			want: model.User{
				ID:          "101030",
				DisplayName: "test.user@databricks.com",
				UserName:    "test.user@databricks.com",
				Groups: []model.GroupsListItem{
					{
						Value: "100002",
					},
					{
						Value: "101355",
					},
				},
				Roles: []model.RoleListItem{
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
			AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet, http.MethodGet}, tt.wantURI, []args{input}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Users().Read(tt.args[0].UserID)
			})
		})
	}
}
