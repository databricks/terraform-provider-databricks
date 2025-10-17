package mws

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestResourcePASCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockPrivateAccessAPI().EXPECT()
			e.Create(mock.Anything, provisioning.CreatePrivateAccessSettingsRequest{
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
				PrivateAccessLevel:        "ACCOUNT",
			}).Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId: "pas_id",
				AccountId:               "abc",
			}, nil)
			e.GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId:   "pas_id",
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
			}, nil)
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

func TestResourcePASCreateWithoutAccountId(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockPrivateAccessAPI().EXPECT()
			e.Create(mock.Anything, provisioning.CreatePrivateAccessSettingsRequest{
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
				PrivateAccessLevel:        "ACCOUNT",
			}).Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId: "pas_id",
				AccountId:               "abc",
			}, nil)
			e.GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId:   "pas_id",
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
			}, nil)
		},
		Resource:  ResourceMwsPrivateAccessSettings(),
		AccountID: "abc",
		HCL: `
		private_access_settings_name = "pas_name"
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASCreate_PublicAccessDisabled(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockPrivateAccessAPI().EXPECT()
			e.Create(mock.Anything, provisioning.CreatePrivateAccessSettingsRequest{
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
				PrivateAccessLevel:        "ACCOUNT",
				PublicAccessEnabled:       false,
				ForceSendFields:           []string{"PublicAccessEnabled"},
			}).Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId: "pas_id",
				AccountId:               "abc",
			}, nil)
			e.GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId:   "pas_id",
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
				ForceSendFields:           []string{"PublicAccessEnabled"},
			}, nil)
		},
		Resource:  ResourceMwsPrivateAccessSettings(),
		AccountID: "abc",
		HCL: `
		account_id = "abc"
		private_access_settings_name = "pas_name"
		public_access_enabled = false
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().Create(mock.Anything, mock.Anything).Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				PrivateAccessSettingsId:   "pas_id",
				Region:                    "ar",
				PrivateAccessSettingsName: "pas_name",
			}, nil)
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Read:     true,
		New:      true,
		ID:       "abc/pas_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/pas_id", d.Id(), "Id should not be empty")
	assert.Equal(t, "abc", d.Get("account_id"))
	assert.Equal(t, "pas_name", d.Get("private_access_settings_name"))
	assert.Equal(t, "ar", d.Get("region"))
	assert.Equal(t, "pas_id", d.Get("private_access_settings_id"))
}

func TestResourcePAStRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "Item not found",
			})
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Read:     true,
		Removed:  true,
		ID:       "abc/pas_id",
	}.ApplyNoError(t)
}

func TestResourcePAS_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockPrivateAccessAPI().EXPECT()
			e.Replace(mock.Anything, provisioning.ReplacePrivateAccessSettingsRequest{
				PrivateAccessSettingsId: "pas_id",
				CustomerFacingPrivateAccessSettings: provisioning.PrivateAccessSettings{
					AccountId:                 "abc",
					Region:                    "eu-west-1",
					PublicAccessEnabled:       true,
					PrivateAccessLevel:        "ENDPOINT",
					PrivateAccessSettingsId:   "pas_id",
					PrivateAccessSettingsName: "pas_name",
					AllowedVpcEndpointIds:     []string{"a", "b"},
				},
			}).Return(&provisioning.PrivateAccessSettings{
				Region:                    "eu-west-1",
				PublicAccessEnabled:       true,
				PrivateAccessLevel:        "ENDPOINT",
				PrivateAccessSettingsId:   "pas_id",
				PrivateAccessSettingsName: "pas_name",
				AllowedVpcEndpointIds:     []string{"a", "b"},
			}, nil)
			e.GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				Region:                    "eu-west-1",
				PublicAccessEnabled:       true,
				PrivateAccessLevel:        "ENDPOINT",
				PrivateAccessSettingsId:   "pas_id",
				PrivateAccessSettingsName: "pas_name",
				AllowedVpcEndpointIds:     []string{"a", "b"},
			}, nil)
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

func TestResourcePAS_Update_PublicAccessDisabled(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(mac *mocks.MockAccountClient) {
			e := mac.GetMockPrivateAccessAPI().EXPECT()
			e.Replace(mock.Anything, provisioning.ReplacePrivateAccessSettingsRequest{
				PrivateAccessSettingsId: "pas_id",
				CustomerFacingPrivateAccessSettings: provisioning.PrivateAccessSettings{
					AccountId:                 "abc",
					Region:                    "eu-west-1",
					PublicAccessEnabled:       false,
					PrivateAccessLevel:        "ENDPOINT",
					PrivateAccessSettingsId:   "pas_id",
					PrivateAccessSettingsName: "pas_name",
					AllowedVpcEndpointIds:     []string{"a", "b"},
					ForceSendFields:           []string{"PublicAccessEnabled"},
				},
			}).Return(&provisioning.PrivateAccessSettings{
				Region:                    "eu-west-1",
				PublicAccessEnabled:       false,
				PrivateAccessLevel:        "ENDPOINT",
				PrivateAccessSettingsId:   "pas_id",
				PrivateAccessSettingsName: "pas_name",
				AllowedVpcEndpointIds:     []string{"a", "b"},
			}, nil)
			e.GetByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(&provisioning.PrivateAccessSettings{
				Region:                    "eu-west-1",
				PublicAccessEnabled:       false,
				PrivateAccessLevel:        "ENDPOINT",
				PrivateAccessSettingsId:   "pas_id",
				PrivateAccessSettingsName: "pas_name",
				AllowedVpcEndpointIds:     []string{"a", "b"},
				ForceSendFields:           []string{"PublicAccessEnabled"},
			}, nil)
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Update:   true,
		ID:       "abc/pas_id",
		HCL: `
		account_id = "abc"
		private_access_settings_name = "pas_name"
		public_access_enabled = false
		region = "eu-west-1"
		private_access_level = "ENDPOINT"
		allowed_vpc_endpoint_ids = ["a", "b"]
		`,
	}.ApplyNoError(t)
}

func TestResourcePASDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().DeleteByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(nil, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockPrivateAccessAPI().EXPECT().DeleteByPrivateAccessSettingsId(mock.Anything, "pas_id").Return(nil, &apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourceMwsPrivateAccessSettings(),
		Delete:   true,
		ID:       "abc/pas_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/pas_id", d.Id())
}

func TestResourcePASUpdateAccountIdNoDiff(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMwsPrivateAccessSettings(),
		ID:       "abc",
		InstanceState: map[string]string{
			"account_id":                   "foo",
			"private_access_settings_name": "pas_name",
			"public_access_enabled":        "false",
			"region":                       "eu-west-1",
			"private_access_level":         "ENDPOINT",
		},
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"private_access_settings_id": {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
		},
		HCL: `
		private_access_settings_name = "pas_name"
		public_access_enabled = false
		region = "eu-west-1"
		private_access_level = "ENDPOINT"
		`,
	}.ApplyNoError(t)
}
