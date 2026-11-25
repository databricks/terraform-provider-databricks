package tags

import (
	"context"
	"net/url"
)

// This file is hand-written (not generated). It extends the generated tag APIs
// to percent-encode the governed-tag key when it travels as a URL *path*
// parameter.
//
// Governed tag keys are hierarchical and may contain "/" (e.g.
// "Field/Shared Technical Services"). The generated impls interpolate the key
// straight into the request path:
//   - TagPolicies:                   /api/2.1/tag-policies/{tag_key}
//   - WorkspaceEntityTagAssignments: /api/2.0/entity-tag-assignments/{entity_type}/{entity_id}/tags/{tag_key}
// so a raw "/" is treated as a path separator and the request resolves to no
// endpoint (404). Encoding the key ("/" -> %2F) makes it route as a single path
// segment; the server decodes it back to the original key.
//
// These methods shadow the promoted (generated) impl methods. Only the tag-key
// path parameter is encoded (entity_type/entity_id are left untouched). The
// Create* methods are intentionally NOT overridden: they send the key in the
// JSON body, where a raw "/" is correct and must be stored verbatim. The TagKey
// fields encoded below are `json:"-" url:"-"`, so they are used only in the
// path — never the body.

func (a *TagPoliciesAPI) GetTagPolicy(ctx context.Context, request GetTagPolicyRequest) (*TagPolicy, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.tagPoliciesImpl.GetTagPolicy(ctx, request)
}

func (a *TagPoliciesAPI) DeleteTagPolicy(ctx context.Context, request DeleteTagPolicyRequest) error {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.tagPoliciesImpl.DeleteTagPolicy(ctx, request)
}

func (a *TagPoliciesAPI) UpdateTagPolicy(ctx context.Context, request UpdateTagPolicyRequest) (*TagPolicy, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.tagPoliciesImpl.UpdateTagPolicy(ctx, request)
}

func (a *WorkspaceEntityTagAssignmentsAPI) GetTagAssignment(ctx context.Context, request GetTagAssignmentRequest) (*TagAssignment, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.workspaceEntityTagAssignmentsImpl.GetTagAssignment(ctx, request)
}

func (a *WorkspaceEntityTagAssignmentsAPI) DeleteTagAssignment(ctx context.Context, request DeleteTagAssignmentRequest) error {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.workspaceEntityTagAssignmentsImpl.DeleteTagAssignment(ctx, request)
}

func (a *WorkspaceEntityTagAssignmentsAPI) UpdateTagAssignment(ctx context.Context, request UpdateTagAssignmentRequest) (*TagAssignment, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.workspaceEntityTagAssignmentsImpl.UpdateTagAssignment(ctx, request)
}
