package common

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Get on path
func (c *DatabricksClient) Get(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodGet, path, request, response, c.addApiPrefix)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPost, path, request, response, c.addApiPrefix)
}

// Delete on path
func (c *DatabricksClient) Delete(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodDelete, path, request, nil, c.addApiPrefix)
}

// Patch on path
func (c *DatabricksClient) Patch(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPatch, path, request, nil, c.addApiPrefix)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPut, path, request, nil, c.addApiPrefix)
}

type ApiVersion string

const (
	API_1_2 ApiVersion = "1.2"
	API_2_0 ApiVersion = "2.0"
	API_2_1 ApiVersion = "2.1"
)

func (c *DatabricksClient) addApiPrefix(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}
	ctx := r.Context()
	av, ok := ctx.Value(Api).(ApiVersion)
	if !ok {
		av = API_2_0
	}
	r.URL.Path = fmt.Sprintf("/api/%s%s", av, r.URL.Path)
	return nil
}

// scimPathVisitorFactory is a separate method for the sake of unit tests
func (c *DatabricksClient) scimVisitor(r *http.Request) error {
	r.Header.Set("Content-Type", "application/scim+json; charset=utf-8")
	if c.Config.IsAccountClient() && c.Config.AccountID != "" {
		// until `/preview` is there for workspace scim,
		// `/api/2.0` is added by completeUrl visitor
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "/api/2.0/preview",
			fmt.Sprintf("/api/2.0/accounts/%s", c.Config.AccountID))
	}
	return nil
}

// Scim sets SCIM headers
func (c *DatabricksClient) Scim(ctx context.Context, method, path string, request any, response any) error {
	return c.Do(ctx, method, path, request, response, c.addApiPrefix, c.scimVisitor)
}
