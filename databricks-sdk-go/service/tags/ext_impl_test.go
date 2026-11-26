package tags

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
)

// These tests verify that the governed-tag key is percent-encoded when it is
// sent as a URL *path* parameter, so a hierarchical key containing "/" (e.g.
// "Field/Shared Technical Services") routes as a single path segment rather
// than being split into multiple segments (which resolves to no endpoint).
// The qa fixture matches on the raw request URI, so a passing match proves the
// "/" was encoded to %2F.

func TestGetTagPolicyEncodesPathKey(t *testing.T) {
	requestMocks := qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/tag-policies/Field%2FShared%20Technical%20Services?",
			Response: TagPolicy{TagKey: "Field/Shared Technical Services"},
		},
	}
	client, server := requestMocks.Client(t)
	defer server.Close()
	api := &TagPoliciesAPI{tagPoliciesImpl: tagPoliciesImpl{client: client}}

	resp, err := api.GetTagPolicy(context.Background(), GetTagPolicyRequest{
		TagKey: "Field/Shared Technical Services",
	})
	if err != nil {
		t.Fatalf("GetTagPolicy returned error: %v", err)
	}
	if resp.TagKey != "Field/Shared Technical Services" {
		t.Fatalf("unexpected tag_key in response: %q", resp.TagKey)
	}
}

func TestDeleteTagPolicyEncodesPathKey(t *testing.T) {
	requestMocks := qa.HTTPFixtures{
		{
			Method:   "DELETE",
			Resource: "/api/2.1/tag-policies/Field%2FShared%20Technical%20Services?",
			Response: map[string]any{},
		},
	}
	client, server := requestMocks.Client(t)
	defer server.Close()
	api := &TagPoliciesAPI{tagPoliciesImpl: tagPoliciesImpl{client: client}}

	err := api.DeleteTagPolicy(context.Background(), DeleteTagPolicyRequest{
		TagKey: "Field/Shared Technical Services",
	})
	if err != nil {
		t.Fatalf("DeleteTagPolicy returned error: %v", err)
	}
}

func TestGetTagAssignmentEncodesPathKey(t *testing.T) {
	requestMocks := qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/entity-tag-assignments/dashboards/123/tags/team%2Fdata-eng?",
			Response: TagAssignment{},
		},
	}
	client, server := requestMocks.Client(t)
	defer server.Close()
	api := &WorkspaceEntityTagAssignmentsAPI{workspaceEntityTagAssignmentsImpl: workspaceEntityTagAssignmentsImpl{client: client}}

	_, err := api.GetTagAssignment(context.Background(), GetTagAssignmentRequest{
		EntityType: "dashboards",
		EntityId:   "123",
		TagKey:     "team/data-eng",
	})
	if err != nil {
		t.Fatalf("GetTagAssignment returned error: %v", err)
	}
}
