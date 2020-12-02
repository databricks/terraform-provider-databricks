package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"

	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

func TestMwsAccCreds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	credsAPI := NewCredentialsAPI(context.Background(), client)
	credsList, err := credsAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(credsList)

	myCreds, err := credsAPI.Create(acctID, "sri-mws-terraform-automation-role", "arn:aws:iam::997819999999:role/sri-e2-terraform-automation-role")
	assert.NoError(t, err, err)

	myCredsFull, err := credsAPI.Read(acctID, myCreds.CredentialsID)
	assert.NoError(t, err, err)
	t.Log(myCredsFull.AwsCredentials.StsRole.ExternalID)

	defer func() {
		err = credsAPI.Delete(acctID, myCreds.CredentialsID)
		assert.NoError(t, err, err)
	}()
}

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
		Resource: ResourceCredentials(),
		State: map[string]interface{}{
			"account_id":       "abc",
			"credentials_name": "Cross-account ARN",
			"role_arn":         "arn:aws:iam::098765:role/cross-account",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceCredentialsCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/credentials",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceCredentials(),
		State: map[string]interface{}{
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
		Resource: ResourceCredentials(),
		Read:     true,
		ID:       "abc/cid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "cid", d.Get("credentials_id"))
	assert.Equal(t, "Cross-account ARN", d.Get("credentials_name"))
	assert.Equal(t, "", d.Get("external_id"))
	assert.Equal(t, "arn:aws:iam::098765:role/cross-account", d.Get("role_arn"))
}

func TestResourceCredentialsRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceCredentials(),
		Read:     true,
		ID:       "abc/cid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceCredentialsRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceCredentials(),
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
		Resource: ResourceCredentials(),
		Delete:   true,
		ID:       "abc/cid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceCredentialsDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/credentials/cid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceCredentials(),
		Delete:   true,
		ID:       "abc/cid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/cid", d.Id())
}
