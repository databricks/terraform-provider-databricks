package mws

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceMwsCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/credentials",

				Response: []Credentials{
					{
						CredentialsID:   "bcd",
						CredentialsName: "123",
					},
					{
						CredentialsID:   "def",
						CredentialsName: "456",
					},
				},
			},
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsCredentials(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"123": "bcd",
			"456": "def",
		},
	})
}

// TestDataSourceMwsCredentials_AccountLevelNoHookFailure is a regression test
// for https://github.com/databricks/terraform-provider-databricks/issues/5672
// (credentials variant). The post-Read provider_config hook must be skipped
// for this account-only data source; otherwise it would crash trying to
// resolve a workspace_id that does not exist.
func TestDataSourceMwsCredentials_AccountLevelNoHookFailure(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/credentials",
			Response: []Credentials{{CredentialsID: "x", CredentialsName: "n"}},
		},
	})
	assert.NoError(t, err)
	defer server.Close()

	client.SetCachedWorkspaceID(0)
	client.Config.AccountID = "abc"

	r := DataSourceMwsCredentials().ToResource()
	d := r.TestResourceData()
	d.SetId("_")
	ctx := context.WithValue(context.Background(), common.ResourceName, "mws_credentials")
	diags := r.ReadContext(ctx, d, client)

	assert.False(t, diags.HasError(), "post-Read hook must be skipped for account-only data source: %v", diags)
}

func TestDataSourceMwsCredentials_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		AccountID:   "abc",
		Resource:    DataSourceMwsCredentials(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
