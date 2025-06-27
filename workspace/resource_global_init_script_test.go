package workspace

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestResourceGlobalInitScriptRead(t *testing.T) {
	scriptID := "1234"
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGlobalInitScriptsAPI().EXPECT().
				GetByScriptId(mock.Anything, scriptID).
				Return(&compute.GlobalInitScriptDetailsWithContent{
					ScriptId:  "1234",
					Name:      "Test",
					Position:  0,
					Enabled:   true,
					CreatedBy: "someuser@domain.com",
					CreatedAt: 1612520583493,
					UpdatedBy: "someuser@domain.com",
					UpdatedAt: 1612520583493,
					Script:    "ZWNobyBoZWxsbw==",
				}, nil)
		},
		Resource: ResourceGlobalInitScript(),
		Read:     true,
		New:      true,
		ID:       scriptID,
	}.ApplyAndExpectData(t, map[string]any{
		"id":       scriptID,
		"name":     "Test",
		"enabled":  true,
		"position": 0,
	})
}

func TestResourceGlobalInitScriptDelete(t *testing.T) {
	scriptID := "1234"
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGlobalInitScriptsAPI().EXPECT().
				DeleteByScriptId(mock.Anything, scriptID).
				Return(nil)
		},
		Resource: ResourceGlobalInitScript(),
		Delete:   true,
		ID:       scriptID,
	}.ApplyAndExpectData(t, map[string]any{
		"id": scriptID,
	})
}

func TestResourceGlobalInitScriptRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGlobalInitScriptsAPI().EXPECT().
				GetByScriptId(mock.Anything, "1234").
				Return(nil, apierr.ErrNotFound)
		},
		Resource: ResourceGlobalInitScript(),
		Read:     true,
		Removed:  true,
		ID:       "1234",
	}.ApplyNoError(t)
}

func TestResourceGlobalInitScriptCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			scriptsAPI := w.GetMockGlobalInitScriptsAPI().EXPECT()
			scriptsAPI.Create(mock.Anything, compute.GlobalInitScriptCreateRequest{
				Name:   "test",
				Script: "ZWNobyBoZWxsbw==",
			}).Return(&compute.CreateResponse{
				ScriptId: "1234",
			}, nil)

			scriptsAPI.GetByScriptId(mock.Anything, "1234").
				Return(&compute.GlobalInitScriptDetailsWithContent{
					ScriptId: "1234",
					Script:   "ZWNobyBoZWxsbw==",
					Name:     "test",
					Position: 0,
				}, nil)
		},
		Create:   true,
		Resource: ResourceGlobalInitScript(),
		State: map[string]any{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "1234",
		"name":           "test",
		"content_base64": "ZWNobyBoZWxsbw==",
		"position":       0,
	})
}

func TestResourceGlobalInitScriptCreateBigPayload(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceGlobalInitScript(),
		Create:   true,
		State: map[string]any{
			"name":           "test",
			"content_base64": base64.StdEncoding.EncodeToString([]byte(strings.Repeat("12", maxScriptSize))),
		},
	}.Apply(t)
	require.Error(t, err)
	assert.Equal(t, "size of the global init script (131072 bytes) exceeds maximal allowed (65536 bytes)", err.Error())
}

func TestResourceGlobalInitScriptUpdateBigPayload(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceGlobalInitScript(),
		Update:   true,
		ID:       "1234",
		State: map[string]any{
			"name":           "test",
			"content_base64": base64.StdEncoding.EncodeToString([]byte(strings.Repeat("12", maxScriptSize))),
		},
	}.Apply(t)
	require.Error(t, err)
	assert.Equal(t, "size of the global init script (131072 bytes) exceeds maximal allowed (65536 bytes)", err.Error())
}

func TestResourceGlobalInitScriptUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			scriptsAPI := w.GetMockGlobalInitScriptsAPI().EXPECT()
			scriptsAPI.Update(mock.Anything, compute.GlobalInitScriptUpdateRequest{
				ScriptId: "1234",
				Name:     "test",
				Position: 0,
				Script:   "ZWNobyBoZWxsbw==",
			}).Return(nil)

			scriptsAPI.GetByScriptId(mock.Anything, "1234").
				Return(&compute.GlobalInitScriptDetailsWithContent{
					ScriptId: "1234",
					Script:   "ZWNobyBoZWxsbw==",
					Position: 0,
					Name:     "test",
				}, nil)
		},
		Update:   true,
		ID:       "1234",
		Resource: ResourceGlobalInitScript(),
		State: map[string]any{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
			"position":       0,
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "1234",
		"name":           "test",
		"content_base64": "ZWNobyBoZWxsbw==",
		"position":       0,
	})
}

func TestResourceGlobalInitScriptReadError(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGlobalInitScriptsAPI().EXPECT().
				GetByScriptId(mock.Anything, "1234").
				Return(nil, errors.New("internal server error"))
		},
		Resource: ResourceGlobalInitScript(),
		Read:     true,
		ID:       "1234",
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, "internal server error", err.Error())
	assert.Equal(t, "1234", d.Id())
}

func TestResourceGlobalInitScriptCreateError(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			scriptsAPI := w.GetMockGlobalInitScriptsAPI().EXPECT()
			scriptsAPI.Create(mock.Anything, compute.GlobalInitScriptCreateRequest{
				Name:   "test",
				Script: "ZWNobyBoZWxsbw==",
			}).Return(nil, errors.New("creation failed"))
		},
		Create:   true,
		Resource: ResourceGlobalInitScript(),
		State: map[string]any{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
		},
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, "creation failed", err.Error())
	assert.Equal(t, "", d.Id())
}

func TestResourceGlobalInitScriptUpdateError(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			scriptsAPI := w.GetMockGlobalInitScriptsAPI().EXPECT()
			scriptsAPI.Update(mock.Anything, compute.GlobalInitScriptUpdateRequest{
				ScriptId: "1234",
				Name:     "test",
				Position: 0,
				Script:   "ZWNobyBoZWxsbw==",
			}).Return(errors.New("update failed"))
		},
		Update:   true,
		ID:       "1234",
		Resource: ResourceGlobalInitScript(),
		State: map[string]any{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
			"position":       0,
		},
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, "update failed", err.Error())
	assert.Equal(t, "1234", d.Id())
}

func TestResourceGlobalInitScriptDeleteError(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockGlobalInitScriptsAPI().EXPECT().
				DeleteByScriptId(mock.Anything, "1234").
				Return(errors.New("delete failed"))
		},
		Resource: ResourceGlobalInitScript(),
		Delete:   true,
		ID:       "1234",
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, "delete failed", err.Error())
	assert.Equal(t, "1234", d.Id())
}
