package sharing

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestRecipientCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceRecipient())
}

func TestCreateRecipient(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/recipients",
				ExpectedRequest: sharing.CreateRecipient{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Owner:              "InitialOwner",
					IpAccessList: &sharing.IpAccessList{
						AllowedIpAddresses: []string{"0.0.0.0/0"},
					},
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/recipients/a?",
				Response: sharing.RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Owner:              "InitialOwner",
					Tokens:             nil,
					IpAccessList: &sharing.IpAccessList{
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
		owner = "InitialOwner"
		ip_access_list {
		   allowed_ip_addresses = ["0.0.0.0/0"]
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "a", d.Get("name"))
	assert.Equal(t, "InitialOwner", d.Get("owner"))
	assert.Equal(t, "TOKEN", d.Get("authentication_type"))
	assert.Equal(t, "b", d.Get("comment"))
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

func TestReadRecipient(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/recipients/a?",
				Response: sharing.RecipientInfo{
					Name:               "a",
					Comment:            "b",
					SharingCode:        "c",
					AuthenticationType: "TOKEN",
					Tokens:             nil,
					IpAccessList: &sharing.IpAccessList{
						AllowedIpAddresses: []string{"0.0.0.0/0"},
					},
				},
			},
		},
		Resource: ResourceRecipient(),
		Read:     true,
		ID:       "a",
		HCL: `
		name = "a"
		comment = "b"
		authentication_type = "TOKEN"
		sharing_code = "c"
		ip_access_list {
		   allowed_ip_addresses = ["0.0.0.0/0"]
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "a", d.Get("name"))
	assert.Equal(t, "b", d.Get("comment"))
}

func TestDeleteRecipient(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/recipients/testRecipient?",
			},
		},
		Resource: ResourceRecipient(),
		Delete:   true,
		ID:       "testRecipient",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testRecipient", d.Id())
}

func TestDeleteRecipientError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/recipients/testRecipient?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceRecipient(),
		Delete:   true,
		ID:       "testRecipient",
	}.ExpectError(t, "Something went wrong")
}
