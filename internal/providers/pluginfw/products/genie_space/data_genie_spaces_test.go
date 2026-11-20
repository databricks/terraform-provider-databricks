package genie_space

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceGenieSpaces_Metadata(t *testing.T) {
	d := DataSourceGenieSpaces()
	resp := &datasource.MetadataResponse{}
	d.Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "databricks"}, resp)
	assert.Equal(t, "databricks_genie_spaces", resp.TypeName)
}

func TestDataSourceGenieSpaces_Schema(t *testing.T) {
	d := DataSourceGenieSpaces()
	resp := &datasource.SchemaResponse{}
	d.Schema(context.Background(), datasource.SchemaRequest{}, resp)
	s := resp.Schema

	titleAttr, ok := s.Attributes["title_contains"].(schema.StringAttribute)
	require.True(t, ok, "title_contains must be a string attribute")
	assert.True(t, titleAttr.Optional, "title_contains should be optional")

	spacesAttr, ok := s.Attributes["spaces"].(schema.ListNestedAttribute)
	require.True(t, ok, "spaces must be a list-nested attribute")
	assert.True(t, spacesAttr.Computed, "spaces should be computed")
	assert.Contains(t, spacesAttr.NestedObject.Attributes, "space_id")
	assert.Contains(t, spacesAttr.NestedObject.Attributes, "title")
	assert.Contains(t, spacesAttr.NestedObject.Attributes, "warehouse_id")

	pcAttr, ok := s.Attributes["provider_config"].(schema.SingleNestedAttribute)
	require.True(t, ok, "provider_config must be a single nested attribute")
	assert.True(t, pcAttr.Optional, "provider_config should be optional")
}

// pagedGenieClient is a fake genieListClient that serves a pre-canned slice
// of pages, one per call. After the last page it errors if called again so
// tests fail loudly when ListSpaces is called more often than expected.
type pagedGenieClient struct {
	pages []*dashboards.GenieListSpacesResponse
	calls int
}

func (p *pagedGenieClient) ListSpaces(ctx context.Context, req dashboards.GenieListSpacesRequest) (*dashboards.GenieListSpacesResponse, error) {
	if p.calls >= len(p.pages) {
		return nil, errors.New("ListSpaces called more times than there are pre-canned pages")
	}
	page := p.pages[p.calls]
	p.calls++
	// Sanity: every call beyond the first must echo back the previous page's
	// next_page_token. We don't enforce the exact value here because the
	// test layout is small enough to inspect call count, but capture mismatch
	// to keep the helper honest.
	if p.calls > 1 {
		prev := p.pages[p.calls-2]
		if req.PageToken != prev.NextPageToken {
			return nil, errors.New("page_token did not match previous next_page_token")
		}
	}
	return page, nil
}

func TestListAllGenieSpaces_SinglePage(t *testing.T) {
	client := &pagedGenieClient{
		pages: []*dashboards.GenieListSpacesResponse{
			{
				Spaces: []dashboards.GenieSpace{
					{SpaceId: "s1", Title: "First"},
					{SpaceId: "s2", Title: "Second"},
				},
			},
		},
	}
	spaces, err := listAllGenieSpaces(context.Background(), client)
	require.NoError(t, err)
	assert.Len(t, spaces, 2)
	assert.Equal(t, 1, client.calls)
}

func TestListAllGenieSpaces_PaginatesAcrossPages(t *testing.T) {
	client := &pagedGenieClient{
		pages: []*dashboards.GenieListSpacesResponse{
			{
				Spaces:        []dashboards.GenieSpace{{SpaceId: "s1", Title: "First"}},
				NextPageToken: "tok-2",
			},
			{
				Spaces:        []dashboards.GenieSpace{{SpaceId: "s2", Title: "Second"}},
				NextPageToken: "tok-3",
			},
			{
				Spaces: []dashboards.GenieSpace{{SpaceId: "s3", Title: "Third"}},
			},
		},
	}
	spaces, err := listAllGenieSpaces(context.Background(), client)
	require.NoError(t, err)
	require.Len(t, spaces, 3)
	assert.Equal(t, "s1", spaces[0].SpaceId)
	assert.Equal(t, "s2", spaces[1].SpaceId)
	assert.Equal(t, "s3", spaces[2].SpaceId)
	assert.Equal(t, 3, client.calls)
}

func TestListAllGenieSpaces_PropagatesError(t *testing.T) {
	wantErr := errors.New("boom")
	client := &erroringGenieClient{err: wantErr}
	spaces, err := listAllGenieSpaces(context.Background(), client)
	assert.Nil(t, spaces)
	assert.ErrorIs(t, err, wantErr)
}

type erroringGenieClient struct{ err error }

func (e *erroringGenieClient) ListSpaces(ctx context.Context, req dashboards.GenieListSpacesRequest) (*dashboards.GenieListSpacesResponse, error) {
	return nil, e.err
}

func TestFilterSpacesByTitleContains_CaseInsensitiveSubstring(t *testing.T) {
	all := []dashboards.GenieSpace{
		{Title: "Sales Genie"},
		{Title: "Marketing analytics"},
		{Title: "Customer Service GENIE"},
	}
	matched := filterSpacesByTitleContains(all, "genie")
	require.Len(t, matched, 2)
	assert.Equal(t, "Sales Genie", matched[0].Title)
	assert.Equal(t, "Customer Service GENIE", matched[1].Title)
}

func TestFilterSpacesByTitleContains_EmptyQueryMatchesAll(t *testing.T) {
	all := []dashboards.GenieSpace{
		{Title: "Alpha"}, {Title: "Beta"},
	}
	matched := filterSpacesByTitleContains(all, "")
	assert.Len(t, matched, 2)
}

func TestFilterSpacesByTitleContains_NoMatch(t *testing.T) {
	all := []dashboards.GenieSpace{{Title: "Alpha"}}
	assert.Empty(t, filterSpacesByTitleContains(all, "zzz"))
}
