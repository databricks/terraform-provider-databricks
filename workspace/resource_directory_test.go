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

func TestResourceDirectoryRead(t *testing.T) {
	path := "/test/path"
	objectID := int64(12345)
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, path).
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeDirectory,
					Path:       path,
				}, nil)
		},
		Resource: ResourceDirectory(),
		Read:     true,
		New:      true,
		ID:       path,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             path,
		"path":           path,
		"workspace_path": "/Workspace" + path,
		"object_id":      int(objectID),
	})
}

func TestResourceDirectoryDelete(t *testing.T) {
	path := "/test/path"
	delete_recursive := true
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Delete(mock.Anything, ws_api.Delete{
					Path:      path,
					Recursive: delete_recursive,
				}).
				Return(nil)
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
		State: map[string]any{
			"path":             "/foo/path.py",
			"delete_recursive": delete_recursive,
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id": path,
	})
}

func TestResourceDirectoryDelete_NotFound(t *testing.T) {
	path := "/test/path"
	delete_recursive := true
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Delete(mock.Anything, ws_api.Delete{
					Path:      path,
					Recursive: delete_recursive,
				}).
				Return(apierr.ErrNotFound)
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
		State: map[string]any{
			"path":             "/foo/path.py",
			"delete_recursive": delete_recursive,
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id": path,
	})
}

func TestResourceDirectoryRead_NotFound(t *testing.T) {
	path := "/test/path"
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, path).
				Return(nil, apierr.ErrNotFound)
		},
		Resource: ResourceDirectory(),
		Read:     true,
		Removed:  true,
		ID:       path,
	}.ApplyNoError(t)
}

func TestResourceDirectoryRead_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, path).
				Return(nil, errors.New("Internal error happened"))
		},
		Resource: ResourceDirectory(),
		Read:     true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, path, d.Id(), "Id should not be empty for error reads")
}

func TestResourceDirectoryCreate(t *testing.T) {
	path := "/test/path"
	objectID := int64(4567)
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.MkdirsByPath(mock.Anything, path).Return(nil)
			workspaceAPI.GetStatusByPath(mock.Anything, path).
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeDirectory,
					Path:       path,
				}, nil)
		},
		Resource: ResourceDirectory(),
		State: map[string]any{
			"path":             path,
			"delete_recursive": false,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             path,
		"path":           path,
		"workspace_path": "/Workspace" + path,
		"object_id":      int(objectID),
	})
}

func TestResourceDirectoryCreate_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				MkdirsByPath(mock.Anything, path).
				Return(errors.New("Internal error happened"))
		},
		Resource: ResourceDirectory(),
		State: map[string]any{
			"path": path,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceDirectoryDelete_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				Delete(mock.Anything, ws_api.Delete{
					Path:      path,
					Recursive: false,
				}).
				Return(errors.New("Internal error happened"))
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, path, d.Id())
}

func TestResourceDirectoryUpdate(t *testing.T) {
	path := "/test/path"
	objectID := int64(4567)
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			workspaceAPI := w.GetMockWorkspaceAPI().EXPECT()
			workspaceAPI.GetStatusByPath(mock.Anything, path).
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeDirectory,
					Path:       path,
				}, nil)
		},
		Resource: ResourceDirectory(),
		InstanceState: map[string]string{
			"path":             path,
			"delete_recursive": "true",
		},
		State: map[string]any{
			"path": path,
		},
		ID:     path,
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             path,
		"path":           path,
		"workspace_path": "/Workspace" + path,
		"object_id":      int(objectID),
	})
}

func TestResourceDirectoryReadNotDirectory(t *testing.T) {
	path := "/test/path"
	objectID := int64(12345)
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, path).
				Return(&ws_api.ObjectInfo{
					ObjectId:   objectID,
					ObjectType: ws_api.ObjectTypeNotebook,
					Path:       path,
					Language:   ws_api.LanguagePython,
				}, nil)
		},
		Resource: ResourceDirectory(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "different object type")
	assert.Equal(t, "", d.Id(), "Id should be empty for different object type read")
}

func TestDirectoryPathSuppressDiff(t *testing.T) {
	assert.True(t, directoryPathSuppressDiff("", "/TF_DIR_WITH_SLASH", "/TF_DIR_WITH_SLASH/", nil))
	assert.False(t, directoryPathSuppressDiff("", "/new_dir", "/TF_DIR_WITH_SLASH/", nil))
}
