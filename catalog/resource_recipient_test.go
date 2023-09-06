package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestRecipientCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceRecipient())
}

func TestCreateRecipient(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/recipients",
				ExpectedRequest: RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Tokens:             nil,
					IpAccessList: &IpAccessList{
						AllowedIpAddresses: []string{"0.0.0.0/0"},
					},
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/recipients/a",
				Response: RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Tokens:             nil,
					IpAccessList: &IpAccessList{
						AllowedIpAddresses: []string{"0.0.0.0/0"},
					},
				},
			},
		},
		Resource: ResourceRecipient(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		authentication_type = "TOKEN"
		sharing_code = "c"
		ip_access_list {
		   allowed_ip_addresses = ["0.0.0.0/0"]
		}
		`,
	}.ApplyNoError(t)
}

func TestCreateRecipient_InvalidAuthType(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceRecipient(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		authentication_type = "temp"
		sharing_code = "c"
		ip_access_list {
		   allowed_ip_addresses = ["0.0.0.0/0"]
		}
		`,
	}.ExpectError(t, "invalid config supplied. "+
		"[authentication_type] expected authentication_type "+
		"to be one of [TOKEN DATABRICKS], got temp")

}
