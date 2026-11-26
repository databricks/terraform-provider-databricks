package catalog

import (
	"context"
	"net/url"
)

// This file is hand-written (not generated). It extends the generated
// EntityTagAssignmentsAPI to percent-encode the governed-tag key when it travels
// as a URL *path* parameter.
//
// Governed tag keys are hierarchical and may contain "/" (e.g.
// "Field/Shared Technical Services"). The generated impl builds the path as
// /api/2.1/unity-catalog/entity-tag-assignments/{entity_type}/{entity_name}/tags/{tag_key},
// so a raw "/" in the key is treated as a path separator and the request
// resolves to no endpoint (404). Encoding the key ("/" -> %2F) makes it route as
// a single path segment; the server decodes it back to the original key.
//
// These methods shadow the promoted (generated) entityTagAssignmentsImpl methods.
// Only the tag-key path parameter is encoded (entity_type/entity_name are left
// untouched). Create is intentionally NOT overridden: it sends the key in the
// JSON body, where a raw "/" is correct and must be stored verbatim. The TagKey
// field encoded below is `json:"-" url:"-"`, so it is used only in the path —
// never the body.

func (a *EntityTagAssignmentsAPI) Get(ctx context.Context, request GetEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Get(ctx, request)
}

func (a *EntityTagAssignmentsAPI) Delete(ctx context.Context, request DeleteEntityTagAssignmentRequest) error {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Delete(ctx, request)
}

func (a *EntityTagAssignmentsAPI) Update(ctx context.Context, request UpdateEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Update(ctx, request)
}
