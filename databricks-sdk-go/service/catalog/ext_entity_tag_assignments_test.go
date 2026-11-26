package catalog

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
)

// Verifies that the governed-tag key is percent-encoded when it is sent as a URL
// *path* parameter, so a hierarchical key containing "/" routes as a single path
// segment. The qa fixture matches on the raw request URI, so a passing match
// proves the "/" was encoded to %2F.
func TestGetEntityTagAssignmentEncodesPathKey(t *testing.T) {
	requestMocks := qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/entity-tag-assignments/schemas/main.default/tags/cost%2Fcenter?",
			Response: EntityTagAssignment{},
		},
	}
	client, server := requestMocks.Client(t)
	defer server.Close()
	api := &EntityTagAssignmentsAPI{entityTagAssignmentsImpl: entityTagAssignmentsImpl{client: client}}

	_, err := api.Get(context.Background(), GetEntityTagAssignmentRequest{
		EntityType: "schemas",
		EntityName: "main.default",
		TagKey:     "cost/center",
	})
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
}
