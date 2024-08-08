package permissions

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

var (
	testAccountId                   = "cb376b18-60fa-4058-b2cd-dd85acf63165"
	testServicePrincipalId          = "1686b74b-a611-4360-8feb-3ef226ad1145"
	testServicePrincipalRuleSetName = fmt.Sprintf("accounts/%s/servicePrincipals/%s/ruleSets/default", testAccountId, testServicePrincipalId)
	ruleSetApiPath                  = "/api/2.0/preview/accounts/access-control/rule-sets"
)

func getResourceName(name string, etag string) string {
	return fmt.Sprintf("%s?etag=%s&name=%s", ruleSetApiPath, url.QueryEscape(etag), url.QueryEscape(name))
}

func TestResourceRuleSetCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: testServicePrincipalRuleSetName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: testServicePrincipalRuleSetName,
						Etag: "etagEx=",
						GrantRules: []iam.GrantRule{
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
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx2=",
					GrantRules: []iam.GrantRule{
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
				Resource: getResourceName(testServicePrincipalRuleSetName, "etagEx2="),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx2=",
					GrantRules: []iam.GrantRule{
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
		Resource: ResourceAccessControlRuleSet(),
		Create:   true,
		HCL: `
		name    = "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default"
		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}
		grant_rules {
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
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
					GrantRules: []iam.GrantRule{
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
		Resource: ResourceAccessControlRuleSet(),
		New:      true,
		Read:     true,
		ID:       testServicePrincipalRuleSetName,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, testServicePrincipalRuleSetName, d.Id(), "Id should not be empty")
	assert.Equal(t, testServicePrincipalRuleSetName, d.Get("name"))
	assert.Equal(t, "etagEx=", d.Get("etag"))
}

func TestResourceRuleSetUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: testServicePrincipalRuleSetName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: testServicePrincipalRuleSetName,
						Etag: "etagEx=",
						GrantRules: []iam.GrantRule{
							{
								Principals: []string{"users/abc@example.com"},
								Role:       "roles/servicePrincipal.manager",
							},
						},
					},
				},
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx2=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, "etagEx2="),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx2=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
		},
		Resource: ResourceAccessControlRuleSet(),
		Update:   true,
		ID:       testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
		HCL: `
		name = "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default"

		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`,
	}.ApplyNoError(t)
}

func TestResourceRuleSetUpdateConflict(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: testServicePrincipalRuleSetName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: testServicePrincipalRuleSetName,
						Etag: "etagEx=",
						GrantRules: []iam.GrantRule{
							{
								Principals: []string{"users/abc@example.com"},
								Role:       "roles/servicePrincipal.manager",
							},
						},
					},
				},
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_CONFLICT",
					Message:   "Conflict with another RuleSet operation",
				},
				Status: 409,
			},
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: ruleSetApiPath,
					Etag: "etagEx2=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"groups/testgroup"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: testServicePrincipalRuleSetName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: testServicePrincipalRuleSetName,
						Etag: "etagEx2=",
						GrantRules: []iam.GrantRule{
							{
								Principals: []string{"users/abc@example.com"},
								Role:       "roles/servicePrincipal.manager",
							},
						},
					},
				},
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx3=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, "etagEx3="),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx3=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
		},
		Resource: ResourceAccessControlRuleSet(),
		Update:   true,
		ID:       testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
		HCL: `
		name = "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default"

		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`,
	}.ApplyNoError(t)
}

func TestResourceRuleSetDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, "etagEx="),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{"users/abc@example.com"},
							Role:       "roles/servicePrincipal.manager",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: testServicePrincipalRuleSetName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: testServicePrincipalRuleSetName,
						Etag: "etagEx=",
					},
				},
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx2=",
				},
			},
		},
		Resource: ResourceAccessControlRuleSet(),
		Delete:   true,
		ID:       testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
	}.ApplyNoError(t)
}
