package sharing

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestProviderCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceProvider())
}

func TestCreateProvider(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/providers",
				ExpectedRequest: ProviderInfo{
					Name:                "a",
					Comment:             "b",
					AuthenticationType:  "TOKEN",
					RecipientProfileStr: "c",
				},
				Response: ProviderInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/providers/a",
				Response: ProviderInfo{
					Name:                "a",
					Comment:             "b",
					AuthenticationType:  "TOKEN",
					RecipientProfileStr: "c",
				},
			},
		},
		Resource: ResourceProvider(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		authentication_type = "TOKEN"
		recipient_profile_str = "c"
		`,
	}.ApplyNoError(t)
}

func TestCreateProvider_InvalidAuthType(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceProvider(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		authentication_type = "temp"
		recipient_profile_str = "{\"shareCredentialsVersion\":1,\"bearerToken\":\"a\",\"endpoint\":\"b\",\"expirationTime\":\"c\"}"
		`,
	}.ExpectError(t, "invalid config supplied. "+
		"[authentication_type] expected authentication_type "+
		"to be one of [TOKEN], got temp")
}
