package permissions

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"

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
		},
		Resource: ResourceAccessControlRuleSet(),
		Create:   true,
		HCL: fmt.Sprintf(`
		name    = "%s"
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
		`, testServicePrincipalRuleSetName),
	}.ApplyAndExpectData(t, map[string]any{
		"name": testServicePrincipalRuleSetName,
		"etag": "",
	})
}

func TestResourceRuleSetRead(t *testing.T) {
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
	}.ApplyAndExpectData(t, map[string]any{
		"name": testServicePrincipalRuleSetName,
		"etag": "",
		"id":   testServicePrincipalRuleSetName,
	})
}

func TestResourceRuleSetUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
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
		},
		Resource: ResourceAccessControlRuleSet(),
		Update:   true,
		ID:       testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
		HCL: fmt.Sprintf(`
		name = "%s"

		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`, testServicePrincipalRuleSetName),
	}.ApplyAndExpectData(t, map[string]any{
		"name": testServicePrincipalRuleSetName,
		"etag": "",
		"id":   testServicePrincipalRuleSetName,
	})
}

func TestResourceRuleSetUpdateName(t *testing.T) {
	newName := "accounts/cb376b18-60fa-4058-b2cd-dd85acf63165/servicePrincipals/1686b74b-a611-4360-8feb-3ef226ad1145/ruleSets/default-2"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
				},
			},
			{
				Method:   "PUT",
				Resource: ruleSetApiPath,
				ExpectedRequest: iam.UpdateRuleSetRequest{
					Name: newName,
					RuleSet: iam.RuleSetUpdateRequest{
						Name: newName,
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
					Name: newName,
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
		Resource:    ResourceAccessControlRuleSet(),
		Update:      true,
		RequiresNew: true,
		ID:          testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
		HCL: fmt.Sprintf(`
		name = "%s"

		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`, newName),
	}.ApplyAndExpectData(t, map[string]any{
		"name": newName,
		"etag": "",
		"id":   newName,
	})
}

func TestResourceRuleSetUpdateConflict(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: getResourceName(testServicePrincipalRuleSetName, ""),
				Response: iam.RuleSetResponse{
					Name: testServicePrincipalRuleSetName,
					Etag: "etagEx=",
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
					Etag: "etagEx=",
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
		},
		Resource: ResourceAccessControlRuleSet(),
		Update:   true,
		ID:       testServicePrincipalRuleSetName,
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
		HCL: fmt.Sprintf(`
		name = "%s"

		grant_rules {
			principals = [
				"users/abc@example.com"
			]
			role = "roles/servicePrincipal.manager"
		}`, testServicePrincipalRuleSetName),
	}.ApplyAndExpectData(t, map[string]any{
		"name": testServicePrincipalRuleSetName,
		"etag": "",
		"id":   testServicePrincipalRuleSetName,
	})
}

func TestResourceRuleSetDelete(t *testing.T) {
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
		State: map[string]any{
			"name": testServicePrincipalRuleSetName,
		},
		InstanceState: map[string]string{
			"name": testServicePrincipalRuleSetName,
			"etag": "etagEx=",
		},
	}.ApplyNoError(t)
}
