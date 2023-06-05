package rule_sets

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	TEST_ACCOUNT_ID                 = "cb376b18-60fa-4058-b2cd-dd85acf63165"
	TEST_SERVICE_PRINCIPAL_ID       = "1686b74b-a611-4360-8feb-3ef226ad1145"
	SERVICE_PRINCIPAL_RULE_SET_NAME = fmt.Sprintf("accounts/%s/servicePrincipals/%s/ruleSets/default", TEST_ACCOUNT_ID, TEST_SERVICE_PRINCIPAL_ID)
)

func TestResourceRuleSetCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=&name=%s", url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/accounts/access-control/rule-sets",
				ExpectedRequest: UpdateRuleSetRequest{
					Name: SERVICE_PRINCIPAL_RULE_SET_NAME,
					RS: RuleSet{
						Name: SERVICE_PRINCIPAL_RULE_SET_NAME,
						Etag: "etagEx=",
						GrantRules: []GrantRule{
							{
								Principals: []string{"users/abc@example.com"},
								Role:       "roles/servicePrincipal.manager",
							},
							{
								Principals: []string{"groups/new_group"},
								Role:       "roles/servicePrincipal.user",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=&name=%s", url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx2=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
						{
							Principals: []string{"groups/new_group"},
							Role:       "roles/servicePrincipal.user",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=%s&name=%s", url.QueryEscape("etagEx2="), url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx2=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
						{
							Principals: []string{"groups/new_group"},
							Role:       "roles/servicePrincipal.user",
						},
					},
				},
			},
		},
		Resource: ResourceRuleSet(),
		Create:   true,
		HCL: `
		name    = "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default"
		grant_rule {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}
		grant_rule {
			principals = [
				"groups/new_group"
			]
			role = "roles/servicePrincipal.user"
		}
		`,
	}.ApplyNoError(t)
}

func TestResourceRuleSetRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=&name=%s", url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
						{
							Principals: []string{"groups/new_group"},
							Role:       "roles/servicePrincipal.user",
						},
					},
				},
			},
		},
		Resource: ResourceRuleSet(),
		New:      true,
		Read:     true,
		ID:       SERVICE_PRINCIPAL_RULE_SET_NAME,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, SERVICE_PRINCIPAL_RULE_SET_NAME, d.Id(), "Id should not be empty")
	assert.Equal(t, SERVICE_PRINCIPAL_RULE_SET_NAME, d.Get("name"))
	assert.Equal(t, "etagEx=", d.Get("etag"))
}

func TestResourceRuleSetUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/accounts/access-control/rule-sets",
				ExpectedRequest: UpdateRuleSetRequest{
					Name: SERVICE_PRINCIPAL_RULE_SET_NAME,
					RS: RuleSet{
						Name: SERVICE_PRINCIPAL_RULE_SET_NAME,
						Etag: "etagEx=",
						GrantRules: []GrantRule{
							{
								Principals: []string{"users/abc@example.com"},
								Role:       "roles/servicePrincipal.manager",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=&name=%s", url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx2=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets?etag=%s&name=%s", url.QueryEscape("etagEx2="), url.QueryEscape(SERVICE_PRINCIPAL_RULE_SET_NAME)),
				Response: RuleSet{
					Name:        SERVICE_PRINCIPAL_RULE_SET_NAME,
					Description: "",
					Etag:        "etagEx2=",
					GrantRules: []GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
		},
		Resource: ResourceRuleSet(),
		Update:   true,
		ID:       SERVICE_PRINCIPAL_RULE_SET_NAME,
		InstanceState: map[string]string{
			"name": SERVICE_PRINCIPAL_RULE_SET_NAME,
			"etag": "etagEx=",
		},
		HCL: `
		name = "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default"

		grant_rule {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`,
	}.ApplyNoError(t)
}
