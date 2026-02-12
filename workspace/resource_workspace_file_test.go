package workspace

import (
	"bytes"
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
	dummyWorkspaceFilePath          = "/foo/path.py"
	dummyWorkspaceFilePayloadBinary = []byte("abc\n")
	dummyWorkspaceFilePayloadBase64 = "YWJjCg=="
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
			workspaceAPI.Upload(mock.Anything, dummyWorkspaceFilePath,
				bytes.NewReader(dummyWorkspaceFilePayloadBinary),
				mock.AnythingOfType("func(*workspace.Import)")).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, dummyWorkspaceFilePath).
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       dummyWorkspaceFilePath,
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayloadBase64,
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
			workspaceAPI.Upload(mock.Anything, dummyWorkspaceFilePath,
				bytes.NewReader(dummyWorkspaceFilePayloadBinary),
				mock.AnythingOfType("func(*workspace.Import)")).Return(
				errors.New("The parent folder (/foo) does not exist.")).Once()
			workspaceAPI.MkdirsByPath(mock.Anything, "/foo").Return(nil)
			workspaceAPI.Upload(mock.Anything, dummyWorkspaceFilePath, bytes.NewReader(dummyWorkspaceFilePayloadBinary),
				mock.AnythingOfType("func(*workspace.Import)")).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, dummyWorkspaceFilePath).
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       dummyWorkspaceFilePath,
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayloadBase64,
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
			workspaceAPI.Upload(mock.Anything, "/foo/path.py",
				bytes.NewReader(dummyWorkspaceFilePayloadBinary),
				mock.AnythingOfType("func(*workspace.Import)")).Return(errors.New("The parent folder (/foo) does not exist."))
			workspaceAPI.MkdirsByPath(mock.Anything, "/foo").
				Return(errors.New("INVALID_REQUEST: Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayloadBase64,
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
			workspaceAPI.Upload(mock.Anything, "/Dashboard",
				bytes.NewReader([]byte(`-- Databricks notebook source
SELECT 10*20

-- COMMAND ----------

SELECT 20*100

-- COMMAND ----------


`)), mock.AnythingOfType("func(*workspace.Import)")).Return(nil)
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
			workspaceAPI.Upload(mock.Anything, "/__init__.py",
				bytes.NewReader([]byte("")),
				mock.AnythingOfType("func(*workspace.Import)")).Return(nil)
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
				Upload(mock.Anything, "/path.py",
					bytes.NewReader(dummyWorkspaceFilePayloadBinary),
					mock.AnythingOfType("func(*workspace.Import)")).Return(errors.New("Internal error happened"))
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayloadBase64,
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
			workspaceAPI.Upload(mock.Anything, "abc",
				bytes.NewReader(dummyWorkspaceFilePayloadBinary),
				mock.AnythingOfType("func(*workspace.Import)")).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, "abc").
				Return(&ws_api.ObjectInfo{
					ObjectId:   4567,
					ObjectType: ws_api.ObjectTypeFile,
					Path:       "abc",
				}, nil)
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayloadBase64,
			"path":           "/path.py",
		},
		ID:          "abc",
		RequiresNew: true,
		Update:      true,
	}.ApplyNoError(t)
}
