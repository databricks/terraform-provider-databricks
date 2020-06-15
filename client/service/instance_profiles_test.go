package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestInstanceProfilesAPI_Create(t *testing.T) {
	type args struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
		SkipValidation     bool   `json:"skip_validation,omitempty" url:"skip_validation,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Create test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				InstanceProfileArn: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
				SkipValidation:     true,
			},
			wantErr: false,
		},
		{
			name:           "Create faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				InstanceProfileArn: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
				SkipValidation:     true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-profiles/add", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.InstanceProfiles().Create(tt.args.InstanceProfileArn, tt.args.SkipValidation)
			})
		})
	}
}

func TestInstanceProfilesAPI_Delete(t *testing.T) {
	type args struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Delete Test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				InstanceProfileArn: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
			},
			wantErr: false,
		},
		{
			name:           "Delete failure Test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				InstanceProfileArn: "arn:aws:iam::1231231123123:instance-profile/my-instance-profile",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-profiles/remove", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.InstanceProfiles().Delete(tt.args.InstanceProfileArn)
			})
		})
	}
}

func TestInstanceProfilesAPI_List(t *testing.T) {
	type args struct{}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           []model.InstanceProfileInfo
		wantErr        bool
	}{
		{
			name: "List test",
			response: `{
						  "instance_profiles": [{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role1", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role2", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role3", "is_meta_instance_profile": false}]
						}`,
			responseStatus: http.StatusOK,
			args:           args{},
			wantURI:        "/api/2.0/instance-profiles/list?",
			want: []model.InstanceProfileInfo{
				{
					InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role1",
				},
				{
					InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role2",
				},
				{
					InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role3",
				},
			},
			wantErr: false,
		},
		{
			name:           "List failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args:           args{},
			wantURI:        "/api/2.0/instance-profiles/list?",
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.InstanceProfiles().List()
			})
		})
	}
}

func TestInstanceProfilesAPI_Read(t *testing.T) {
	type args struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           string
		wantErr        bool
	}{
		{
			name: "Read test",
			response: `{
						  "instance_profiles": [{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role1", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role2", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role3", "is_meta_instance_profile": false}]
						}`,
			responseStatus: http.StatusOK,
			args: args{
				InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role1",
			},
			wantURI: "/api/2.0/instance-profiles/list?",
			want:    "arn:aws:iam::123456789:instance-profile/datascience-role1",
			wantErr: false,
		},
		{
			name: "Read profile not found failure test",
			response: `{
						  "instance_profiles": [{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role1", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role2", "is_meta_instance_profile": false},
												{"instance_profile_arn":"arn:aws:iam::123456789:instance-profile/datascience-role3", "is_meta_instance_profile": false}]
						}`,
			responseStatus: http.StatusOK,
			args: args{
				InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role4",
			},
			wantURI: "/api/2.0/instance-profiles/list?",
			want:    "",
			wantErr: true,
		},
		{
			name:           "Read list failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args: args{
				InstanceProfileArn: "arn:aws:iam::123456789:instance-profile/datascience-role1",
			},
			wantURI: "/api/2.0/instance-profiles/list?",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.InstanceProfiles().Read(tt.args.InstanceProfileArn)
			})
		})
	}
}
