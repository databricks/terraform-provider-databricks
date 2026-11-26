// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domains

import (
	"context"
)

// Manage domains for organizing and discovering data assets.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DomainsService interface {

	// Create a domain. If `domain_id` is omitted, the server generates one.
	CreateDomain(ctx context.Context, request CreateDomainRequest) (*Domain, error)

	// Delete a domain. By default the request fails if the domain still has
	// Glossary pages; set `force` to delete those pages along with the domain.
	DeleteDomain(ctx context.Context, request DeleteDomainRequest) error

	// Get a domain by resource name.
	//
	// Authorization: external callers must have the `MANAGE DISCOVER`
	// permission.
	GetDomain(ctx context.Context, request GetDomainRequest) (*Domain, error)

	// List domains in the account. Set `parent_domain_id` to return only the
	// direct subdomains of a given domain.
	//
	// Authorization: external callers must have the `MANAGE DISCOVER`
	// permission; only domains the caller is authorized to read are returned.
	ListDomains(ctx context.Context, request ListDomainsRequest) (*ListDomainsResponse, error)

	// Update a domain. `update_mask` selects which fields to modify; the domain
	// is identified by its resource `name`.
	UpdateDomain(ctx context.Context, request UpdateDomainRequest) (*Domain, error)
}
