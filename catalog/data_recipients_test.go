package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestRecipientsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/recipients",
				Response: Recipients{
					Recipients: []RecipientInfo{
						{
							Name:               "a",
							Comment:            "c",
							SharingCode:        "sc",
							AuthenticationType: "TOKEN",
							Tokens:             nil,
						},
						{
							Name:               "b",
							Comment:            "c",
							SharingCode:        "sc",
							AuthenticationType: "TOKEN",
							Tokens:             nil,
						},
					},
				},
			},
		},
		Resource:    DataSourceRecipients(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestRecipientsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceRecipients(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
