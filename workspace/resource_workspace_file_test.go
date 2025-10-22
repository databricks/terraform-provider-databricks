package workspace

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	ws_api "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	dummyWorkspaceFilePath    = "/foo/path.py"
	dummyWorkspaceFilePayload = "YWJjCg=="
)

func TestResourceWorkspaceFileRead(t *testing.T) {
	objectID := int64(12345)
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, dummyWorkspaceFilePath).
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       dummyWorkspaceFilePath,
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		New:      true,
		ID:       dummyWorkspaceFilePath,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             dummyWorkspaceFilePath,
		"path":           dummyWorkspaceFilePath,
		"workspace_path": "/Workspace" + dummyWorkspaceFilePath,
		"object_id":      int(objectID),
	})
}

func TestResourceWorkspaceFileDelete(t *testing.T) {
	path := "/test/path.py"
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Delete(mock.Anything, ws_api.Delete{
					Path:      path,
					Recursive: false,
				}).
				Return(nil)
		},
		Resource: ResourceWorkspaceFile(),
		Delete:   true,
		ID:       path,
	}.ApplyAndExpectData(t, map[string]any{
		"id": path,
	})
}

func TestResourceWorkspaceFileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/test/path").
				Return(nil, apierr.ErrNotFound)
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		Removed:  true,
		ID:       "/test/path",
	}.ApplyNoError(t)
}

func TestResourceWorkspaceFileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/test/path").
				Return(nil, errors.New("Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		ID:       "/test/path",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path", d.Id(), "Id should not be empty for error reads")
}

func TestResourceWorkspaceFileCreate_DirectoryExist(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         dummyWorkspaceFilePayload,
				Path:            dummyWorkspaceFilePath,
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, dummyWorkspaceFilePath).
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       dummyWorkspaceFilePath,
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": dummyWorkspaceFilePath,
	})
}

func TestResourceWorkspaceFileCreate_DirectoryDoesntExist(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         dummyWorkspaceFilePayload,
				Path:            dummyWorkspaceFilePath,
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(errors.New("The parent folder (/foo) does not exist.")).Once()
			workspaceAPI.MkdirsByPath(mock.Anything, "/foo").Return(nil)
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         dummyWorkspaceFilePayload,
				Path:            dummyWorkspaceFilePath,
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, dummyWorkspaceFilePath).
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       dummyWorkspaceFilePath,
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": dummyWorkspaceFilePath,
	})
}

func TestResourceWorkspaceFileCreate_DirectoryCreateError(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         dummyWorkspaceFilePayload,
				Path:            dummyWorkspaceFilePath,
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(errors.New("The parent folder (/foo) does not exist."))
			workspaceAPI.MkdirsByPath(mock.Anything, "/foo").
				Return(errors.New("INVALID_REQUEST: Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
		},
		Create: true,
	}.Apply(t)
	assert.Error(t, err, "Internal error happened")
}

func TestResourceWorkspaceFileCreateSource(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content: "LS0gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKU0VMRUNUIDEwKjIwC" +
					"gotLSBDT01NQU5EIC0tLS0tLS0tLS0KClNFTEVDVCAyMCoxMDAKCi0tIE" +
					"NPTU1BTkQgLS0tLS0tLS0tLQoKCg==",
				Path:            "/Dashboard",
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, "/Dashboard").
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "/Dashboard",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"source": "acceptance/testdata/tf-test-sql.sql",
			"path":   "/Dashboard",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "/Dashboard",
		"path":           "/Dashboard",
		"workspace_path": "/Workspace/Dashboard",
		"object_id":      4567})
}

func TestResourceWorkspaceFileCreateEmptyFileSource(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         "",
				Path:            "/__init__.py",
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, "/__init__.py").
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "/__init__.py",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"source": "acceptance/testdata/empty_file",
			"path":   "/__init__.py",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "/__init__.py",
	})
}

func TestResourceWorkspaceFileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Import(mock.Anything, ws_api.Import{
					Content:         dummyWorkspaceFilePayload,
					Path:            "/path.py",
					Overwrite:       true,
					Format:          ws_api.ImportFormatRaw,
					ForceSendFields: []string{"Content"},
				}).Return(errors.New("Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           "/path.py",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceWorkspaceFileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Delete(mock.Anything, ws_api.Delete{
					Path: "abc",
				}).
				Return(errors.New("Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceWorkspaceFileUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.Import(mock.Anything, ws_api.Import{
				Content:         dummyWorkspaceFilePayload,
				Path:            "abc",
				Overwrite:       true,
				Format:          ws_api.ImportFormatRaw,
				ForceSendFields: []string{"Content"},
			}).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, "abc").
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "abc",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           "/path.py",
		},
		ID:          "abc",
		RequiresNew: true,
		Update:      true,
	}.ApplyNoError(t)
}

func TestResourceWorkspaceFileRead_WorkspacePrefixNormalization(t *testing.T) {
	objectID := int64(12345)
	// Test case 1: Config without /Workspace prefix, API returns with prefix
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/Users/user@example.com/file.py").
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "/Workspace/Users/user@example.com/file.py",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		New:      true,
		ID:       "/Users/user@example.com/file.py",
		State: map[string]any{
			"path": "/Users/user@example.com/file.py",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "/Users/user@example.com/file.py",
		"path":           "/Users/user@example.com/file.py", // Should match configured path
		"workspace_path": "/Workspace/Users/user@example.com/file.py",
		"object_id":      int(objectID),
	})

	// Test case 2: Config with /Workspace prefix, API returns without prefix
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/Workspace/Users/user@example.com/file.py").
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "/Users/user@example.com/file.py",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		New:      true,
		ID:       "/Workspace/Users/user@example.com/file.py",
		State: map[string]any{
			"path": "/Workspace/Users/user@example.com/file.py",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "/Workspace/Users/user@example.com/file.py",
		"path":           "/Workspace/Users/user@example.com/file.py", // Should match configured path
		"workspace_path": "/Workspace/Workspace/Users/user@example.com/file.py",
		"object_id":      int(objectID),
	})
}
