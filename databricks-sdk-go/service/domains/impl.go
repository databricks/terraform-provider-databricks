// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domains

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Domains API methods
type domainsImpl struct {
	client *client.DatabricksClient
}

func (a *domainsImpl) CreateDomain(ctx context.Context, request CreateDomainRequest) (*Domain, error) {
	var domain Domain
	path := "/api/2.0/domains"
	queryParams := make(map[string]any)

	if request.DomainId != "" || slices.Contains(request.ForceSendFields, "DomainId") {
		queryParams["domain_id"] = request.DomainId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Domain, &domain)
	return &domain, err
}

func (a *domainsImpl) DeleteDomain(ctx context.Context, request DeleteDomainRequest) error {
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *domainsImpl) GetDomain(ctx context.Context, request GetDomainRequest) (*Domain, error) {
	var domain Domain
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &domain)
	return &domain, err
}

// List domains in the account. Set `parent_domain_id` to return only the direct
// subdomains of a given domain.
//
// Authorization: external callers must have the `MANAGE DISCOVER` permission;
// only domains the caller is authorized to read are returned.
func (a *domainsImpl) ListDomains(ctx context.Context, request ListDomainsRequest) listing.Iterator[Domain] {

	getNextPage := func(ctx context.Context, req ListDomainsRequest) (*ListDomainsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDomains(ctx, req)
	}
	getItems := func(resp *ListDomainsResponse) []Domain {
		return resp.Domains
	}
	getNextReq := func(resp *ListDomainsResponse) *ListDomainsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List domains in the account. Set `parent_domain_id` to return only the direct
// subdomains of a given domain.
//
// Authorization: external callers must have the `MANAGE DISCOVER` permission;
// only domains the caller is authorized to read are returned.
func (a *domainsImpl) ListDomainsAll(ctx context.Context, request ListDomainsRequest) ([]Domain, error) {
	iterator := a.ListDomains(ctx, request)
	return listing.ToSlice[Domain](ctx, iterator)
}

func (a *domainsImpl) internalListDomains(ctx context.Context, request ListDomainsRequest) (*ListDomainsResponse, error) {
	var listDomainsResponse ListDomainsResponse
	path := "/api/2.0/domains"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDomainsResponse)
	return &listDomainsResponse, err
}

func (a *domainsImpl) UpdateDomain(ctx context.Context, request UpdateDomainRequest) (*Domain, error) {
	var domain Domain
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Domain, &domain)
	return &domain, err
}
