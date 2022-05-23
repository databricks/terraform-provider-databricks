package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestRecipientCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceRecipient())
}

func TestCreateRecipient(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/recipients",
				ExpectedRequest: RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Tokens:             nil,
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/recipients/a",
				Response: RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Tokens:             nil,
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
		`,
	}.ApplyNoError(t)
}
