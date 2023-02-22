package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceCredentialsCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/credentials",
				ExpectedRequest: Credentials{
					CredentialsName: "Cross-account ARN",
					AwsCredentials: &AwsCredentials{
						StsRole: &StsRole{
							RoleArn: "arn:aws:iam::098765:role/cross-account",
						},
					},
				},
				Response: Credentials{
					CredentialsID: "cid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: Credentials{
					CredentialsID:   "cid",
					CredentialsName: "Cross-account ARN",
					AwsCredentials: &AwsCredentials{
						StsRole: &StsRole{
							RoleArn: "arn:aws:iam::098765:role/cross-account",
						},
					},
				},
			},
		},
		Resource: ResourceMwsCredentials(),
		State: map[string]any{
			"account_id":       "abc",
			"credentials_name": "Cross-account ARN",
			"role_arn":         "arn:aws:iam::098765:role/cross-account",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceCredentialsCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/credentials",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsCredentials(),
		State: map[string]any{
			"account_id":       "abc",
			"credentials_name": "Cross-account ARN",
			"role_arn":         "arn:aws:iam::098765:role/cross-account",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceCredentialsRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: Credentials{
					CredentialsID:   "cid",
					CredentialsName: "Cross-account ARN",
					AwsCredentials: &AwsCredentials{
						StsRole: &StsRole{
							RoleArn: "arn:aws:iam::098765:role/cross-account",
						},
					},
				},
			},
		},
		Resource: ResourceMwsCredentials(),
		Read:     true,
		ID:       "abc/cid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "cid", d.Get("credentials_id"))
	assert.Equal(t, "Cross-account ARN", d.Get("credentials_name"))
	assert.Equal(t, "", d.Get("external_id"))
	assert.Equal(t, "arn:aws:iam::098765:role/cross-account", d.Get("role_arn"))
}

func TestResourceCredentialsRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsCredentials(),
		Read:     true,
		Removed:  true,
		ID:       "abc/cid",
	}.ApplyNoError(t)
}

func TestResourceCredentialsRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsCredentials(),
		Read:     true,
		ID:       "abc/cid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/cid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceCredentialsDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
			},
		},
		Resource: ResourceMwsCredentials(),
		Delete:   true,
		ID:       "abc/cid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceCredentialsDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsCredentials(),
		Delete:   true,
		ID:       "abc/cid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/cid", d.Id())
}
