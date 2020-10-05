package identity

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				ExpectedRequest: map[string]interface{}{
					"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
					"skip_validation":      true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"skip_validation":      true,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"skip_validation":      true,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstanceProfileCreate_Error_InvalidARN(t *testing.T) {
	d, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "abc",
			"skip_validation":      true,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Illegal instance profile abc: arn: invalid prefix")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstanceProfileRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Get("instance_profile_arn"))
	assert.Equal(t, false, d.Get("skip_validation"))
}

func TestResourceInstanceProfileRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceInstanceProfileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty for error reads")
}

func TestResourceInstanceProfileDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/remove",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Delete:   true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/remove",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		Delete:   true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-profiles/add", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewInstanceProfilesAPI(client).Create(tt.args.InstanceProfileArn, tt.args.SkipValidation)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/instance-profiles/remove", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewInstanceProfilesAPI(client).Delete(tt.args.InstanceProfileArn)
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
		want           []InstanceProfileInfo
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
			wantURI:        "/api/2.0/instance-profiles/list",
			want: []InstanceProfileInfo{
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
			wantURI:        "/api/2.0/instance-profiles/list",
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewInstanceProfilesAPI(client).List()
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
			wantURI: "/api/2.0/instance-profiles/list",
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
			wantURI: "/api/2.0/instance-profiles/list",
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
			wantURI: "/api/2.0/instance-profiles/list",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewInstanceProfilesAPI(client).Read(tt.args.InstanceProfileArn)
			})
		})
	}
}

func TestAwsAccInstanceProfiles(t *testing.T) {
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	client := common.NewClientFromEnvironment()
	instanceProfilesAPI := NewInstanceProfilesAPI(client)
	defer func() {
		err := instanceProfilesAPI.Delete(arn)
		assert.NoError(t, err, err)
	}()
	err := instanceProfilesAPI.Create(arn, true)
	assert.NoError(t, err, err)

	arnSearch, err := instanceProfilesAPI.Read(arn)
	assert.NoError(t, err, err)
	assert.True(t, len(arnSearch) > 0)
}
