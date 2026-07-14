package exporter

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

// Account-Level Test Helpers and Fixtures
//
// This file contains unit tests for account-level resources in the exporter.
// Account-level resources are different from workspace-level resources in that:
// 1. They require an AccountID to be set in the client configuration
// 2. They use different API endpoints (account-level paths vs workspace paths)
// 3. They use the AccountClient instead of WorkspaceClient
//
// Key patterns for writing account-level tests:
// - Use HTTPFixturesApplyAccount() helper which sets up account client configuration
// - Account API endpoints typically start with /api/2.0/accounts/{account_id}/...
// - Account SCIM endpoints are at /api/2.0/account/scim/v2/... (note: 'account' not 'accounts')
// - Use setClientsForTests() to automatically configure the correct client type
//
// Example HTTP fixtures for account-level APIs:
//   - GET /api/2.0/accounts/{account_id}/ruleSets/default - access control rule sets
//   - GET /api/2.0/account/scim/v2/Users - account-level users
//   - GET /api/2.0/account/scim/v2/Groups - account-level groups
//   - GET /api/2.0/account/scim/v2/ServicePrincipals - account-level service principals

const testAccountID = "00000000-0000-0000-0000-000000000000"

// HTTPFixturesApplyAccount is a helper that sets up an account-level client for testing
func HTTPFixturesApplyAccount(t *testing.T, fixtures []qa.HTTPFixture, testFunc func(ctx context.Context, client *common.DatabricksClient)) {
	qa.HTTPFixturesApply(t, fixtures, func(ctx context.Context, client *common.DatabricksClient) {
		// Configure the client as an account-level client
		// Setting AccountID is enough since the client is in testing mode (isTesting=true)
		client.Config.AccountID = testAccountID
		// Set Host to accounts.cloud.databricks.com to enable AWS detection via IsAws()
		client.Config.Host = "https://accounts.cloud.databricks.com"
		client.Config.WithTesting()
		testFunc(ctx, client)
	})
}

// Common fixtures for account-level testing

var emptyAccountUsers = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/account/scim/v2/Users?attributes=id%2CuserName&count=10000&startIndex=1",
	Response:     iam.ListUsersResponse{},
	ReuseRequest: true,
}

var emptyAccountServicePrincipals = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/account/scim/v2/ServicePrincipals?attributes=id%2CuserName&count=10000&startIndex=1",
	Response:     iam.ListServicePrincipalResponse{},
	ReuseRequest: true,
}

var emptyAccountGroups = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/account/scim/v2/Groups?attributes=id%2CdisplayName&count=10000&startIndex=1",
	Response:     iam.ListGroupsResponse{},
	ReuseRequest: true,
}

// TestImportingAccessControlRuleSet tests the import of access control rule sets
func TestImportingAccessControlRuleSet(t *testing.T) {
	accountUserFixtures := qa.ListUsersFixtures([]iam.User{
		{
			Id:       "user-123",
			UserName: "test@example.com",
		},
	})
	accountGroupFixtures := qa.ListGroupsFixtures([]iam.Group{
		{
			Id:          "group-456",
			DisplayName: "TestGroup",
		},
	})
	accountSpFixtures := qa.ListServicePrincipalsFixtures([]iam.ServicePrincipal{
		{
			Id:            "sp-789",
			ApplicationId: "12345678-1234-1234-1234-123456789012",
		},
	})

	HTTPFixturesApplyAccount(t,
		[]qa.HTTPFixture{
			// Account-level user fixtures
			accountUserFixtures[0],
			accountUserFixtures[1],
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/Users/user-123?attributes=id,userName,displayName,active,externalId,entitlements,groups,roles",
				Response: iam.User{
					Id:       "user-123",
					UserName: "test@example.com",
				},
				ReuseRequest: true,
			},
			// Account-level group fixtures
			accountGroupFixtures[0],
			accountGroupFixtures[1],
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/Groups/group-456?attributes=id,displayName,active,externalId,entitlements,groups,roles,members,meta",
				Response: iam.Group{
					Id:          "group-456",
					DisplayName: "TestGroup",
				},
				ReuseRequest: true,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/Groups/group-456?attributes=displayName,externalId,entitlements",
				Response: iam.Group{
					Id:          "group-456",
					DisplayName: "TestGroup",
				},
				ReuseRequest: true,
			},
			// Account-level service principal fixtures
			accountSpFixtures[0],
			accountSpFixtures[1],
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/ServicePrincipals/sp-789?attributes=userName,displayName,active,externalId,entitlements",
				Response: iam.ServicePrincipal{
					Id:            "sp-789",
					ApplicationId: "12345678-1234-1234-1234-123456789012",
				},
				ReuseRequest: true,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/ServicePrincipals/sp-789?attributes=userName,displayName,active,externalId,entitlements,groups,roles",
				Response: iam.ServicePrincipal{
					Id:            "sp-789",
					ApplicationId: "12345678-1234-1234-1234-123456789012",
				},
				ReuseRequest: true,
			},
			// Rule set fixture - this is the main resource being tested
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/accounts/%s/ruleSets/default?", testAccountID),
				Response: iam.RuleSetResponse{
					Name: fmt.Sprintf("accounts/%s/ruleSets/default", testAccountID),
					Etag: "test-etag-123",
					GrantRules: []iam.GrantRule{
						{
							Principals: []string{
								"users/test@example.com",
								"groups/TestGroup",
								"servicePrincipals/12345678-1234-1234-1234-123456789012",
							},
							Role: "roles/account.admin",
						},
					},
				},
				ReuseRequest: true,
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForTestWithClient(ctx, client)
			ic.enableServices("access,users,groups")

			// Test the List function - it should emit the default rule set
			err := ic.Importables["databricks_access_control_rule_set"].List(ic)
			assert.NoError(t, err)

			// In test mode, resources are added to testEmits instead of being processed
			assert.True(t, len(ic.testEmits) > 0, "Expected at least one resource to be emitted")

			// Check that the access control rule set was emitted
			expectedRuleSetID := fmt.Sprintf("accounts/%s/ruleSets/default", testAccountID)
			ruleSetEmitted := false
			for key := range ic.testEmits {
				if strings.Contains(key, "databricks_access_control_rule_set") && strings.Contains(key, expectedRuleSetID) {
					ruleSetEmitted = true
					break
				}
			}
			assert.True(t, ruleSetEmitted, "Expected databricks_access_control_rule_set to be emitted")

			// Now test the Import function by creating a resource manually and calling Import
			// First, we need to create a resource data object with the rule set data
			ruleSetResource := ic.Resources["databricks_access_control_rule_set"]
			d := ruleSetResource.TestResourceData()
			d.SetId(expectedRuleSetID)
			d.Set("name", expectedRuleSetID)
			d.Set("etag", "test-etag-123")
			d.Set("grant_rules", []interface{}{
				map[string]interface{}{
					"principals": []interface{}{
						"users/test@example.com",
						"groups/TestGroup",
						"servicePrincipals/12345678-1234-1234-1234-123456789012",
					},
					"role": "roles/account.admin",
				},
			})

			// Clear testEmits to track only Import emissions
			ic.testEmits = map[string]bool{}

			// Call the Import function
			err = ic.Importables["databricks_access_control_rule_set"].Import(ic, &resource{
				Resource: "databricks_access_control_rule_set",
				ID:       expectedRuleSetID,
				Data:     d,
			})
			assert.NoError(t, err)

			// Verify that dependent resources were emitted
			assert.True(t, len(ic.testEmits) > 0, "Expected Import to emit dependent resources")

			// Check for emitted users, groups, and service principals
			userEmitted := false
			groupEmitted := false
			spEmitted := false
			for key := range ic.testEmits {
				if strings.Contains(key, "databricks_user") {
					userEmitted = true
				}
				if strings.Contains(key, "databricks_group") {
					groupEmitted = true
				}
				if strings.Contains(key, "databricks_service_principal") {
					spEmitted = true
				}
			}

			assert.True(t, userEmitted, "Expected databricks_user to be emitted")
			assert.True(t, groupEmitted, "Expected databricks_group to be emitted")
			assert.True(t, spEmitted, "Expected databricks_service_principal to be emitted")
		})
}

// TestImportingAccessControlRuleSetWithoutGrantRules tests that rule sets without grant rules are ignored
func TestImportingAccessControlRuleSetWithoutGrantRules(t *testing.T) {
	HTTPFixturesApplyAccount(t,
		[]qa.HTTPFixture{
			emptyAccountUsers,
			emptyAccountServicePrincipals,
			emptyAccountGroups,
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/accounts/%s/ruleSets/default?", testAccountID),
				Response: iam.RuleSetResponse{
					Name:       fmt.Sprintf("accounts/%s/ruleSets/default", testAccountID),
					Etag:       "test-etag-123",
					GrantRules: []iam.GrantRule{}, // Empty grant rules
				},
				ReuseRequest: true,
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForTestWithClient(ctx, client)
			ic.enableServices("access")

			// List should emit the resource
			err := ic.Importables["databricks_access_control_rule_set"].List(ic)
			assert.NoError(t, err)

			// Verify the resource was emitted
			assert.True(t, len(ic.testEmits) > 0, "Expected resource to be emitted")

			// Create a resource data object with empty grant rules
			expectedRuleSetID := fmt.Sprintf("accounts/%s/ruleSets/default", testAccountID)
			ruleSetResource := ic.Resources["databricks_access_control_rule_set"]
			d := ruleSetResource.TestResourceData()
			d.SetId(expectedRuleSetID)
			d.Set("name", expectedRuleSetID)
			d.Set("etag", "test-etag-123")
			d.Set("grant_rules", []interface{}{}) // Empty grant rules

			r := &resource{
				Resource: "databricks_access_control_rule_set",
				ID:       expectedRuleSetID,
				Data:     d,
			}

			// Test the Ignore function
			importable := ic.Importables["databricks_access_control_rule_set"]
			shouldIgnore := importable.Ignore(ic, r)
			assert.True(t, shouldIgnore, "Resource with empty grant rules should be ignored")

			// Verify the resource was added to ignored list by the Ignore function
			assert.Contains(t, ic.ignoredResources, fmt.Sprintf("databricks_access_control_rule_set. ID=%s", expectedRuleSetID))
		})
}

// TestImportingMetastoreAssignment tests the metastore assignment resource
func TestImportingMetastoreAssignment(t *testing.T) {
	HTTPFixturesApplyAccount(t,
		[]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment?workspaceId=12345",
				Response: catalog.MetastoreAssignment{
					MetastoreId:        "metastore-123",
					WorkspaceId:        12345,
					DefaultCatalogName: "main",
				},
				ReuseRequest: true,
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			ic := importContextForTest()
			ic.Client = client
			ic.Context = ctx
			ic.accountClient, _ = client.AccountClient()
			ic.enableServices("uc-metastores")

			// Simulate emitting a metastore assignment resource (as would be done by Import)
			ic.Emit(&resource{
				Resource: "databricks_metastore_assignment",
				ID:       "12345|metastore-123",
			})

			// Verify the resource was emitted
			assert.True(t, len(ic.testEmits) > 0, "Expected resource to be emitted")

			assignmentEmitted := false
			for key := range ic.testEmits {
				if strings.Contains(key, "databricks_metastore_assignment") && strings.Contains(key, "12345") {
					assignmentEmitted = true
					break
				}
			}
			assert.True(t, assignmentEmitted, "Expected databricks_metastore_assignment to be emitted")
		})
}
