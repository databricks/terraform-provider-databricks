package storage

import (
	"bytes"
	"context"
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDbfsCreateChunksContentsAndClosesHandle(t *testing.T) {
	contents := bytes.Repeat([]byte("a"), 1_000_003)
	client, _, err := qa.StrictHTTPFixtureClient(t, []qa.HTTPFixture{
		{
			Method:          "POST",
			Resource:        "/api/2.0/dbfs/create",
			ExpectedRequest: createHandle{Path: "/fixture", Overwrite: true},
			Response:        handleResponse{Handle: 42},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
			ExpectedRequest: addBlock{
				Data:   base64.StdEncoding.EncodeToString(contents[:1_000_000]),
				Handle: 42,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
			ExpectedRequest: addBlock{
				Data:   base64.StdEncoding.EncodeToString(contents[1_000_000:]),
				Handle: 42,
			},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.0/dbfs/close",
			ExpectedRequest: handleResponse{Handle: 42},
		},
	})
	require.NoError(t, err)

	err = NewDbfsAPI(context.Background(), client).Create("/fixture", contents, true)
	require.NoError(t, err)
}

func TestDbfsReadAndRecursiveListAtHTTPBoundary(t *testing.T) {
	firstChunk := bytes.Repeat([]byte("a"), 1_000_000)
	lastChunk := []byte("xyz")
	client, _, err := qa.StrictHTTPFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&path=%2Ffixture",
			Response: ReadResponse{
				BytesRead: int64(len(firstChunk)),
				Data:      base64.StdEncoding.EncodeToString(firstChunk),
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&offset=1000000&path=%2Ffixture",
			Response: ReadResponse{
				BytesRead: int64(len(lastChunk)),
				Data:      base64.StdEncoding.EncodeToString(lastChunk),
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=%2Froot",
			Response: FileList{Files: []FileInfo{
				{Path: "/root/file", IsDir: false},
				{Path: "/root/subdir", IsDir: true},
			}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=%2Froot%2Fsubdir",
			Response: FileList{Files: []FileInfo{
				{Path: "/root/subdir/nested", IsDir: false},
			}},
		},
	})
	require.NoError(t, err)

	api := NewDbfsAPI(context.Background(), client)
	contents, err := api.Read("/fixture")
	require.NoError(t, err)
	assert.Equal(t, append(firstChunk, lastChunk...), contents)

	files, err := api.List("/root", true)
	require.NoError(t, err)
	assert.Equal(t, []FileInfo{
		{Path: "/root/file", IsDir: false},
		{Path: "/root/subdir/nested", IsDir: false},
	}, files)
}

func TestCreateFileFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: createHandle{
				Path:      "/create-fails",
				Overwrite: true,
			},
			Status: 404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/create-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot create handle: fails")
	})
}

func TestCreateFile_AddBlockFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: createHandle{
				Path:      "/add-fails",
				Overwrite: true,
			},
			Response: handleResponse{123},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
			ExpectedRequest: addBlock{
				Data:   "YWJj",
				Handle: 123,
			},
			Status: 404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/add-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot add block: fails")
	})
}

func TestCreateFile_CloseFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/close-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot close handle: fails")
	})
}

func TestDbfsListRecursiveFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=abc",
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=sub",
			Response: FileList{
				Files: []FileInfo{
					{
						Path:  "bcd",
						IsDir: true,
					},
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/dbfs/list?path=bcd",
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		_, err := a.List("abc", true)
		assert.EqualError(t, err, "cannot list abc: fails")
		_, err = a.List("sub", true)
		assert.EqualError(t, err, "cannot list subfolder: cannot list bcd: fails")
		_, err = a.List("bcd", false)
		assert.EqualError(t, err, "cannot list bcd: fails")
	})
}

func TestDbfsReadFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny: true,
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "fails",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		_, err := a.Read("abc")
		assert.EqualError(t, err, "cannot read abc: fails")
	})
}
