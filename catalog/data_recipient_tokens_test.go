package catalog

import (
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestRecipientTokensData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/recipients/a",
				Response: RecipientInfo{
					Name:               "a",
					Comment:            "c",
					SharingCode:        "sc",
					AuthenticationType: "TOKEN",
					Tokens: []Token{
						{
							Id:             "a",
							CreatedAt:      0,
							CreatedBy:      "",
							ActivationUrl:  "",
							ExpirationTime: 0,
							UpdatedAt:      0,
							UpdatedBy:      "",
						},
						{
							Id:             "b",
							CreatedAt:      12,
							CreatedBy:      "",
							ActivationUrl:  "",
							ExpirationTime: 0,
							UpdatedAt:      0,
							UpdatedBy:      "",
						},
						{
							Id:             "c",
							CreatedAt:      31,
							CreatedBy:      "",
							ActivationUrl:  "",
							ExpirationTime: 0,
							UpdatedAt:      0,
							UpdatedBy:      "",
						},
					},
				},
			},
		},
		Resource:    DataSourceRecipientTokens(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "a"
		`,
	}.Apply(t)
	assert.NoError(t, err, "no error")
	var recipientData recipientsData
	common.DataToStructPointer(d, DataSourceRecipientTokens().Schema, &recipientData)
	assert.Equalf(t, 3, len(recipientData.Tokens), "There should be 3 tokens")
	assert.Equalf(t, "c", recipientData.Tokens[0].Id, "First token id should be c")
	assert.Equalf(t, "b", recipientData.Tokens[1].Id, "Second token id should be b")
	assert.Equalf(t, "a", recipientData.Tokens[2].Id, "Third token id should be a")
}

func TestRecipientTokensData_Latest(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/recipients/a",
				Response: RecipientInfo{
					Name:               "a",
					Comment:            "c",
					SharingCode:        "sc",
					AuthenticationType: "TOKEN",
					Tokens: []Token{
						{
							Id:             "a",
							CreatedAt:      0,
							CreatedBy:      "",
							ActivationUrl:  "",
							ExpirationTime: 0,
							UpdatedAt:      0,
							UpdatedBy:      "",
						},
						{
							Id:             "b",
							CreatedAt:      2,
							CreatedBy:      "",
							ActivationUrl:  "",
							ExpirationTime: 0,
							UpdatedAt:      0,
							UpdatedBy:      "",
						},
					},
				},
			},
		},
		Resource:    DataSourceRecipientTokens(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "a"
		latest = true
		`,
	}.Apply(t)
	assert.NoError(t, err, "no error")
	var recipientData recipientsData
	common.DataToStructPointer(d, DataSourceRecipientTokens().Schema, &recipientData)
	assert.Equalf(t, recipientData.Tokens[0].Id, "b", "Id should be b and sorted descending")
}

func TestRecipientTokensData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceRecipientTokens(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
