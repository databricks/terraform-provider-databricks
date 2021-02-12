package workspace

import (
	"encoding/base64"
	"net/http"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceGlobalInitScriptRead(t *testing.T) {
	scriptID := "1234"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/global-init-scripts/1234",
				Response: GlobalInitScriptInfo{
					ScriptID:      "1234",
					Name:          "Test",
					Position:      0,
					Enabled:       true,
					CreatedBy:     "someuser@domain.com",
					CreatedAt:     1612520583493,
					UpdatedBy:     "someuser@domain.com",
					UpdatedAt:     1612520583493,
					ContentBase64: "ZWNobyBoZWxsbw==",
				},
			},
		},
		Resource: ResourceGlobalInitScript(),
		Read:     true,
		New:      true,
		ID:       scriptID,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, scriptID, d.Id())
	assert.Equal(t, "Test", d.Get("name"))
	assert.Equal(t, true, d.Get("enabled"))
	assert.Equal(t, 0, d.Get("position"))
}

func TestResourceGlobalInitScriptDelete(t *testing.T) {
	scriptID := "1234"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodDelete,
				Resource:        "/api/2.0/global-init-scripts/" + scriptID,
				Status:          http.StatusOK,
				ExpectedRequest: map[string]string{"script_id": scriptID},
			},
		},
		Resource: ResourceGlobalInitScript(),
		Delete:   true,
		ID:       scriptID,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, scriptID, d.Id())
}

func TestResourceGlobalInitScriptRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/global-init-scripts/1234",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "The global unit script with ID 1234 does not exist.",
				},
				Status: 404,
			},
		},
		Resource: ResourceGlobalInitScript(),
		Read:     true,
		Removed:  true,
		ID:       "1234",
	}.ApplyNoError(t)
}

func TestResourceGlobalInitScriptCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/global-init-scripts",
				ExpectedRequest: GlobalInitScriptPayload{
					Name:          "test",
					ContentBase64: "ZWNobyBoZWxsbw==",
				},
				Response: globalInitScriptCreateResponse{
					ScriptID: "1234",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts/1234",
				ReuseRequest: true,
				Response: GlobalInitScriptInfo{
					ScriptID:      "1234",
					ContentBase64: "ZWNobyBoZWxsbw==",
					Position:      0,
					Name:          "test",
				},
			},
		},
		Create:   true,
		Resource: ResourceGlobalInitScript(),
		State: map[string]interface{}{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "1234", d.Id())
	assert.Equal(t, 0, d.Get("position"))
}

func TestResourceGlobalInitScriptCreateBigPayload(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Create:   true,
		Resource: ResourceGlobalInitScript(),
		State: map[string]interface{}{
			"name":           "test",
			"content_base64": base64.StdEncoding.EncodeToString([]byte(strings.Repeat("12", maxScriptSize))),
		},
	}.Apply(t)
	require.Error(t, err, err)
	assert.Equal(t, "Size of the global init script (131072 bytes) exceeds maximal allowed (65536 bytes)", err.Error())
}

func TestResourceGlobalInitScriptUpdateBigPayload(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Update:   true,
		Resource: ResourceGlobalInitScript(),
		ID:       "1234",
		State: map[string]interface{}{
			"name":           "test",
			"content_base64": base64.StdEncoding.EncodeToString([]byte(strings.Repeat("12", maxScriptSize))),
		},
	}.Apply(t)
	require.Error(t, err, err)
	assert.Equal(t, "Size of the global init script (131072 bytes) exceeds maximal allowed (65536 bytes)", err.Error())
}

func TestResourceGlobalInitScriptUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/global-init-scripts/1234",
				ExpectedRequest: GlobalInitScriptPayload{
					Name:          "test",
					Position:      0,
					ContentBase64: "ZWNobyBoZWxsbw==",
				},
				Response: globalInitScriptCreateResponse{
					ScriptID: "1234",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts/1234",
				ReuseRequest: true,
				Response: GlobalInitScriptInfo{
					ScriptID:      "1234",
					ContentBase64: "ZWNobyBoZWxsbw==",
					Position:      0,
					Name:          "test",
				},
			},
		},
		Update:   true,
		ID:       "1234",
		Resource: ResourceGlobalInitScript(),
		State: map[string]interface{}{
			"name":           "test",
			"content_base64": "ZWNobyBoZWxsbw==",
			"position":       0,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "1234", d.Id())
	assert.Equal(t, 0, d.Get("position"))
}
