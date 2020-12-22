package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMWSPAS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	pasAPI := NewPASAPI(ctx, client)
	pasList, err := pasAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(pasList)

	pas := PAS{
		AccountID: acctID,
		PasName:   qa.RandomName(),
		PasID:     "",
	}
	err = pasAPI.Create(&pas)
	assert.NoError(t, err, err)
	defer func() {
		err = pasAPI.Delete(acctID, pas.PasID)
		assert.NoError(t, err, err)
	}()

	myPAS, err := pasAPI.Read(acctID, pas.PasID)
	assert.NoError(t, err, err)
	t.Log(myPAS)
}

func TestResourcePASCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/private-access-settings",
				ExpectedRequest: PAS{
					AccountID: "abc",
					Region:    "ar",
					PasName:   "pas_name",
					PasID:     "pas_id",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",

				Response: PAS{
					AccountID: "abc",
					Region:    "ar",
					PasName:   "pas_name",
					PasID:     "pas_id",
				},
			},
		},
		Resource: ResourcePAS(),
		HCL: `
		account_id = "abc"
		private_access_settings_name = "pas_name"
		region = "ar"
		private_access_settings_id = "pas_id"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/private-access-settings",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePAS(),
		State: map[string]interface{}{
			"account_id":                   "abc",
			"private_access_settings_name": "pas_name",
			"region":                       "ar",
			"private_access_settings_id":   "pas_id",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourcePASRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: PAS{
					AccountID: "account_id",
					PasID:     "pas_id",
					PasName:   "pas_name",
					Region:    "ar",
				},
			},
		},
		Resource: ResourcePAS(),
		Read:     true,
		New:      true,
		ID:       "abc/pas_id",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/pas_id", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "account_id", d.Get("account_id"))
	assert.Equal(t, "pas_name", d.Get("private_access_settings_name"))
	assert.Equal(t, "ar", d.Get("region"))
	assert.Equal(t, "pas_id", d.Get("private_access_settings_id"))
}

func TestResourcePAStRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourcePAS(),
		Read:     true,
		Removed:  true,
		ID:       "abc/pas_id",
	}.ApplyNoError(t)
}

func TestResourcePAS_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePAS(),
		Read:     true,
		ID:       "abc/pas_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/pas_id", d.Id(), "Id should not be empty for error reads")
}

func TestResourcePASDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: PAS{
					AccountID: "account_id",
					PasID:     "pas_id",
					PasName:   "pas_name",
					Region:    "ar",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Yes, it's not found",
				},
				Status: 404,
			},
		},
		Resource: ResourcePAS(),
		Delete:   true,
		ID:       "abc/pas_id",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePAS(),
		Delete:   true,
		ID:       "abc/pas_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASList(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/private-access-settings",
			Response: []PAS{},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	l, err := NewPASAPI(context.Background(), client).List("abc")
	require.NoError(t, err)
	assert.Len(t, l, 0)
}
