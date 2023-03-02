package mws

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourcePASCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/private-access-settings",
				ExpectedRequest: PrivateAccessSettings{
					AccountID:          "abc",
					Region:             "ar",
					PasName:            "pas_name",
					PrivateAccessLevel: "ACCOUNT",
				},
				Response: PrivateAccessSettings{
					PasID: "pas_id",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",

				Response: PrivateAccessSettings{
					AccountID: "abc",
					PasID:     "pas_id",
					Region:    "ar",
					PasName:   "pas_name",
				},
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		HCL: `
		account_id = "abc"
		private_access_settings_name = "pas_name"
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/private-access-settings",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		State: map[string]any{
			"account_id":                   "abc",
			"private_access_settings_name": "pas_name",
			"region":                       "ar",
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
				Response: PrivateAccessSettings{
					AccountID: "account_id",
					PasName:   "pas_name",
					Region:    "ar",
				},
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Read:     true,
		New:      true,
		ID:       "abc/pas_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id(), "Id should not be empty")
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
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
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
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Read:     true,
		ID:       "abc/pas_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/pas_id", d.Id(), "Id should not be empty for error reads")
}

func TestResourcePAS_Update(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				ExpectedRequest: PrivateAccessSettings{
					Region:                "eu-west-1",
					PublicAccessEnabled:   true,
					PrivateAccessLevel:    "ENDPOINT",
					AccountID:             "abc",
					PasID:                 "pas_id",
					PasName:               "pas_name",
					AllowedVpcEndpointIDS: []string{"a", "b"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: PrivateAccessSettings{
					Region:              "eu-west-1",
					PublicAccessEnabled: true,
					AccountID:           "abc",
					PasID:               "pas_id",
					PasName:             "pas_name",
				},
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Update:   true,
		ID:       "abc/pas_id",
		HCL: `
		account_id = "abc"
		private_access_settings_name = "pas_name"
		public_access_enabled = true
		region = "eu-west-1"
		private_access_level = "ENDPOINT"
		allowed_vpc_endpoint_ids = ["a", "b"]
		`,
	}.ApplyNoError(t)
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
				Response: PrivateAccessSettings{
					AccountID: "account_id",
					PasID:     "pas_id",
					PasName:   "pas_name",
					Region:    "ar",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Yes, it's not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Delete:   true,
		ID:       "abc/pas_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/private-access-settings/pas_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsPrivateAccessSettings(),
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
			Response: []PrivateAccessSettings{},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	l, err := NewPrivateAccessSettingsAPI(context.Background(), client).List("abc")
	require.NoError(t, err)
	assert.Len(t, l, 0)
}
