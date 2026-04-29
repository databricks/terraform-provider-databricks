package mws

import (
	"testing"

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

func TestDataSourceMwsCredentials_NoProviderConfigInSchema(t *testing.T) {
	s := DataSourceMwsCredentials().Schema
	if _, ok := s["provider_config"]; ok {
		t.Fatalf("databricks_mws_credentials is account-only and must not expose provider_config")
	}
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
