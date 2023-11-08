package storage

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	ws_api "github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceFileCreate(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	decodedString, err := base64.StdEncoding.DecodeString("YWJjCg==")
	assert.NoError(t, err)
	reader := io.NopCloser(bytes.NewReader(decodedString))
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				Status:   http.StatusOK,
				Response: nil,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: ws_api.DownloadResponse{
					Contents: reader,
				},
			},
		},
		Resource: ResourceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           path,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceFileCreateSource(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	source := "testdata/tf-test-python.py"
	reader, err := os.Open(source)
	assert.NoError(t, err)
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				Status:   http.StatusOK,
				Response: nil,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: ws_api.DownloadResponse{
					Contents: reader,
				},
			},
		},
		Resource: ResourceFile(),
		State: map[string]any{
			"source": source,
			"path":   path,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceFileCreate_Error(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           path,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceFileRead(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: ws_api.DownloadResponse{},
			},
		},
		Resource: ResourceFile(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, path, d.Get("path"))
}

func TestResourceFileRead_NotFound(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "File not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceFile(),
		Read:     true,
		Removed:  true,
		ID:       path,
	}.ApplyNoError(t)
}

func TestResourceFileRead_Error(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceFile(),
		Read:     true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, path, d.Id(), "Id should not be empty for error reads")
}

func TestResourceFileDelete(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: nil,
			},
		},
		Resource: ResourceFile(),
		Delete:   true,
		New:      true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceFileDelete_Error(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceFile(),
		Delete:   true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, path, d.Id())
}

func TestResourceFileUpdate(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				Status:   http.StatusOK,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: ws_api.DownloadResponse{},
			},
		},
		Resource: ResourceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           path,
		},
		ID:          path,
		RequiresNew: true,
		Update:      true,
	}.ApplyNoError(t)
}

func TestResourceFileBadPrefix(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceFile(),
		Create:   true,
		HCL: `
		path = "Volumes/CatalogName/SchemaName/VolumeName/fileName"
		content_base64 = "YWJjCg=="
		`,
	}.ExpectError(t, "invalid config supplied. [path] Path should start with /Volumes")
}
