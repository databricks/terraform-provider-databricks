package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceGroupRoleCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequestWithValue("add", "roles", "arn:aws:iam::000000000000:role/test-role"),
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []ComplexValue{
						{
							Value: "arn:aws:iam::000000000000:role/test-role",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		State: map[string]any{
			"group_id": "abc",
			"role":     "arn:aws:iam::000000000000:role/test-role",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "abc|arn:aws:iam::000000000000:role/test-role"})
}

func TestResourceGroupRoleCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupRole(),
		State: map[string]any{
			"group_id": "abc",
			"role":     "arn:aws:iam::000000000000:role/test-role",
		},
		Create: true,
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupRoleRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []ComplexValue{
						{
							Value: "arn:aws:iam::000000000000:role/test-role",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyAndExpectData(t, map[string]any{"id": "abc|arn:aws:iam::000000000000:role/test-role"})
}

func TestResourceGroupRoleRead_NoRole(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyNoError(t)
}

func TestResourceGroupRoleRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyNoError(t)
}

func TestResourceGroupRoleRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupRoleDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequest(
					"remove",
					`roles[value eq "arn:aws:iam::000000000000:role/test-role"]`),
			},
		},
		Resource: ResourceGroupRole(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyNoError(t)
}

func TestResourceGroupRoleDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupRole(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ExpectError(t, "Internal error happened")
}
