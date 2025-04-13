package oauth

import (
	"fmt"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestResourceServicePrincipalFederationPolicyUpdate(t *testing.T) {
	spid := 44
	id := "07cc67f4-45d6-494b-adac-09b5cbc7e2b5"
	structProvider := func() oauth2.FederationPolicy {
		return oauth2.FederationPolicy{
			Description: "My federation policy description.",
			Name:        "Federation Policy Name",
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Audiences:    []string{"databricks"},
				Issuer:       "https://myidp.example.com/oidc",
				JwksJson:     "string",
				JwksUri:      "string",
				Subject:      "subject",
				SubjectClaim: "sub",
			},
		}
	}
	input := structProvider()
	returnValue := structProvider()
	returnValue.Uid = id
	returnValue.OidcPolicy.Audiences = []string{"newvalue"}
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			api := m.GetMockServicePrincipalFederationPolicyAPI()
			api.EXPECT().Update(
				mock.Anything,
				oauth2.UpdateServicePrincipalFederationPolicyRequest{
					PolicyId:           id,
					Policy:             &input,
					ServicePrincipalId: int64(spid),
				}).Return(&returnValue, nil)
			api.EXPECT().GetByServicePrincipalIdAndPolicyId(mock.Anything, int64(spid), id).Return(
				&returnValue, nil)
		},
		Resource:  ResourceServicePrincipalFederationPolicy(),
		Update:    true,
		AccountID: "account_id",
		ID:        id,
		InstanceState: map[string]string{
			"service_principal_id": "44",
			"name":                 "Federation Policy Name",
		},
		HCL: `
			description =           "My federation policy description."
			name =                 "Federation Policy Name"
			service_principal_id = 44
			oidc_policy  {
				issuer = "https://myidp.example.com/oidc"
				audiences = ["databricks"]
				subject= "subject"
				subject_claim = "sub"
				"jwks_uri" = "string"
				"jwks_json" = "string"
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"service_principal_id": spid,
		"description":          "My federation policy description.",
		"name":                 "Federation Policy Name",
		"oidc_policy": []interface{}{
			map[string]any{
				"audiences":     []interface{}{"newvalue"},
				"issuer":        "https://myidp.example.com/oidc",
				"jwks_json":     "string",
				"jwks_uri":      "string",
				"subject":       "subject",
				"subject_claim": "sub",
			},
		},
	})
}

func TestResourceServicePrincipalFederationPolicyCreateWithoutName(t *testing.T) {
	spid := 44
	id := "07cc67f4-45d6-494b-adac-09b5cbc7e2b5"
	accountId := "account_id"
	name := fmt.Sprintf("accounts/%s/servicePrincipals/%d/federationPolicies/%s",
		accountId,
		spid,
		id,
	)
	structProvider := func() oauth2.FederationPolicy {
		return oauth2.FederationPolicy{
			Description: "My federation policy description.",
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Audiences:    []string{"databricks"},
				Issuer:       "https://myidp.example.com/oidc",
				JwksJson:     "string",
				JwksUri:      "string",
				Subject:      "subject",
				SubjectClaim: "sub",
			},
		}
	}
	input := structProvider()
	returnValue := structProvider()
	returnValue.Uid = id
	returnValue.Name = name
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			api := m.GetMockServicePrincipalFederationPolicyAPI()
			api.EXPECT().Create(
				mock.Anything,
				oauth2.CreateServicePrincipalFederationPolicyRequest{
					Policy:             &input,
					ServicePrincipalId: int64(spid),
				}).Return(&returnValue, nil)
			api.EXPECT().GetByServicePrincipalIdAndPolicyId(mock.Anything, int64(spid), id).Return(
				&returnValue, nil)
		},
		Resource:  ResourceServicePrincipalFederationPolicy(),
		Create:    true,
		AccountID: accountId,
		HCL: `
			description =           "My federation policy description."
			service_principal_id = 44
			oidc_policy  {
				issuer = "https://myidp.example.com/oidc"
				audiences = ["databricks"]
				subject= "subject"
				subject_claim = "sub"
				"jwks_uri" = "string"
				"jwks_json" = "string"
			}
		`,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalFederationPolicyCreate(t *testing.T) {
	spid := 44
	accountId := "account_id"
	id := "provided_id"
	name := fmt.Sprintf("accounts/%s/servicePrincipals/%d/federationPolicies/%s",
		accountId,
		spid,
		id,
	)
	structProvider := func() oauth2.FederationPolicy {
		return oauth2.FederationPolicy{
			Description: "My federation policy description.",
			Name:        name,
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Audiences:    []string{"databricks"},
				Issuer:       "https://myidp.example.com/oidc",
				JwksJson:     "string",
				JwksUri:      "string",
				Subject:      "subject",
				SubjectClaim: "sub",
			},
		}
	}
	input := structProvider()
	returnValue := structProvider()
	returnValue.Uid = name
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			api := m.GetMockServicePrincipalFederationPolicyAPI()
			api.EXPECT().Create(
				mock.Anything,
				oauth2.CreateServicePrincipalFederationPolicyRequest{
					Policy:             &input,
					ServicePrincipalId: int64(spid),
				}).Return(&returnValue, nil)
			api.EXPECT().GetByServicePrincipalIdAndPolicyId(mock.Anything, int64(spid), id).Return(
				&returnValue, nil)
		},
		Resource:  ResourceServicePrincipalFederationPolicy(),
		Create:    true,
		AccountID: accountId,
		HCL: fmt.Sprintf(`
			description =           "My federation policy description."
			name =                 "%s"
			service_principal_id = 44
			oidc_policy  {
				issuer = "https://myidp.example.com/oidc"
				audiences = ["databricks"]
				subject= "subject"
				subject_claim = "sub"
				"jwks_uri" = "string"
				"jwks_json" = "string"
			}
		`, name),
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalFederationPolicyRead(t *testing.T) {
	id := "07cc67f4-45d6-494b-adac-09b5cbc7e2b5"
	spid := 42
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			fp := &oauth2.FederationPolicy{
				Uid:         id,
				UpdateTime:  "2019-08-24T14:15:22Z",
				CreateTime:  "2019-08-24T14:15:22Z",
				Description: "My federation policy description.",
				Name:        "Federation Policy Name",
				OidcPolicy: &oauth2.OidcFederationPolicy{
					Audiences:    []string{"databricks"},
					Issuer:       "https://myidp.example.com/oidc",
					JwksJson:     "string",
					JwksUri:      "string",
					Subject:      "subject",
					SubjectClaim: "sub",
				},
			}
			m.GetMockServicePrincipalFederationPolicyAPI().EXPECT().GetByServicePrincipalIdAndPolicyId(mock.Anything, int64(spid), id).Return(
				fp, nil)
		},
		Resource:  ResourceServicePrincipalFederationPolicy(),
		Read:      true,
		New:       true,
		AccountID: "account_id",
		ID:        id,
		HCL: `
			description =           "My federation policy description."
			name =                 "Federation Policy Name"
			service_principal_id = 42
			oidc_policy  {
				issuer = "https://myidp.example.com/oidc"
				audiences = ["databricks"]
				subject= "subject"
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"service_principal_id": spid,
		"description":          "My federation policy description.",
		"name":                 "Federation Policy Name",
		"oidc_policy": []interface{}{
			map[string]any{
				"audiences":     []interface{}{"databricks"},
				"issuer":        "https://myidp.example.com/oidc",
				"jwks_json":     "string",
				"jwks_uri":      "string",
				"subject":       "subject",
				"subject_claim": "sub",
			},
		},
	})
}

func TestResourceServicePrincipalFederationPolicyDelete(t *testing.T) {
	id := "07cc67f4-45d6-494b-adac-09b5cbc7e2b5"
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			m.GetMockServicePrincipalFederationPolicyAPI().EXPECT().DeleteByServicePrincipalIdAndPolicyId(
				mock.Anything,
				int64(123),
				id,
			).Return(nil)
		},
		Resource:  ResourceServicePrincipalFederationPolicy(),
		Delete:    true,
		ID:        id,
		AccountID: "account_id",
		HCL: `
			description =           "My federation policy description."
			name =                 "Federation Policy Name"
			service_principal_id = 123
			oidc_policy  {
				issuer = "https://myidp.example.com/oidc"
				audiences = ["databricks"]
				subject= "subject"
			}
		`,
	}.ApplyNoError(t)
}
