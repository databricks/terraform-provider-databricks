package workspace

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	ws_api "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceDirectory(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/a/b/c").
				Return(&ws_api.ObjectInfo{
					ObjectId:   987,
					ObjectType: ws_api.ObjectTypeDirectory,
					Path:       "/a/b/c",
				}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "/a/b/c",
		"object_id":      987,
		"workspace_path": "/Workspace/a/b/c",
	})
}

func TestDataSourceDirectory_NotDirectory(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/a/b/c").
				Return(&ws_api.ObjectInfo{
					ObjectId:   987,
					Language:   ws_api.LanguagePython,
					ObjectType: ws_api.ObjectTypeNotebook,
					Path:       "/a/b/c",
				}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.ExpectError(t, "'/a/b/c' isn't a directory")
}

func TestDataSourceDirectory_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceAPI().EXPECT().
				GetStatusByPath(mock.Anything, "/a/b/c").
				Return(nil, errors.New("i'm a teapot"))
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.ExpectError(t, "i'm a teapot")
}
