package storage

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/databricks/databricks-sdk-go/service/files"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceFileCreate(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				ExpectedHeaders: map[string]string{"Content-Type": "application/octet-stream"},
				ExpectedBody:    []byte("abc\n"),
				Status:          http.StatusNoContent,
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				ResponseHeaders: map[string]string{
					"Content-Length": "4",
					"Content-Type":   "text/plain",
					"Last-Modified":  "Wed, 21 Oct 2015 07:28:00 GMT",
				},
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				ResponseHeaders: map[string]string{
					"Content-Length": "4",
					"Content-Type":   "text/plain",
					"Last-Modified":  "Wed, 21 Oct 2015 07:28:00 GMT",
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
	assert.Equal(t, "Wed, 21 Oct 2015 07:28:00 GMT", d.Get("modification_time"))
	assert.Equal(t, 4, d.Get("file_size"))
}

func TestResourceFileCreateSource(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	source := "testdata/tf-test-python.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName",
				Status:   http.StatusOK,
				Response: nil,
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
				},
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
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
				Response: apierr.APIError{
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
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
				},
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
				},
			},
		},
		Resource: ResourceFile(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceFileRead_NotFound(t *testing.T) {
	path := "/Volumes/CatalogName/SchemaName/VolumeName/fileName"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "HEAD",
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Status:   404,
			},
		},
		Resource: ResourceFile(),
		Read:     true,
		Removed:  true,
		ID:       path,
	}.ApplyNoError(t)
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
				Response: apierr.APIError{
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
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
				},
			},
			{
				Method:   http.MethodHead,
				Resource: "/api/2.0/fs/files/Volumes/CatalogName/SchemaName/VolumeName/fileName?",
				Response: files.GetMetadataResponse{
					LastModified:  "Wed, 21 Oct 2015 07:28:00 GMT",
					ContentLength: 1024,
				},
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
