// Package declarative classifies Databricks API errors that carry
// server-provided "declarative_context" metadata, so generated Terraform
// resources can act on server-signaled conditions (e.g. suppress a Delete
// error when a parent will cascade-clean the child).
package declarative

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/apierr"
)

const (
	metadataKey     = "declarative_context"
	managedByParent = "MANAGED_BY_PARENT"
)

// IsDeleteError reports whether err is an actionable Delete error that the
// caller should surface. It returns false when err carries a
// declarative_context marker (today: MANAGED_BY_PARENT) signaling that the
// Delete is safe to disregard — when the resource's lifecycle is owned by
// a parent (e.g. a Lakebase endpoint inside a branch), the server rejects
// the standalone Delete knowing the parent will cascade-clean.
//
// Call sites null out err when this returns false, then fall through to
// the existing error-handling chain:
//
//	if !declarative.IsDeleteError(err) {
//	    err = nil
//	}
//	if err != nil && !apierr.IsMissing(err) { ... }
//
// apierr.IsMissing (404-on-delete) is a separate SDK convention and stays
// at the call site.
func IsDeleteError(err error) bool {
	if err == nil {
		return false
	}
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) || apiErr == nil {
		return true
	}
	info := apiErr.ErrorDetails().ErrorInfo
	if info == nil {
		return true
	}
	return info.Metadata[metadataKey] != managedByParent
}
