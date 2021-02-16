package workspace

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// NewGlobalInitScriptsAPI creates GlobalInitScriptsAPI instance from provider meta
func NewGlobalInitScriptsAPI(ctx context.Context, m interface{}) GlobalInitScriptsAPI {
	// TODO: context.WithValue
	return GlobalInitScriptsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// GlobalInitScriptsAPI exposes the Global Init Scripts API: https://docs.databricks.com/dev-tools/api/latest/global-init-scripts.html#
type GlobalInitScriptsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// GlobalInitScriptInfo contains information about registered global init script
type GlobalInitScriptInfo struct {
	ScriptID      string `json:"script_id"`
	Name          string `json:"name"`
	Position      int32  `json:"position,omitempty" tf:"computed"`
	Enabled       bool   `json:"enabled,omitempty"`
	CreatedBy     string `json:"creator_user_name,omitempty"`
	CreatedAt     int64  `json:"created_at_timestamp,omitempty"`
	UpdatedBy     string `json:"updater_user_name,omitempty"`
	UpdatedAt     int64  `json:"updated_at_timestamp,omitempty"`
	ContentBase64 string `json:"script,omitempty"`
}

// GlobalInitScriptPayload contains information about registered global init script
type GlobalInitScriptPayload struct {
	Name          string `json:"name"`
	Position      int32  `json:"position"`
	Enabled       bool   `json:"enabled,omitempty"`
	ContentBase64 string `json:"script"`
}

type globalInitScriptCreateResponse struct {
	ScriptID string `json:"script_id"`
}

type globalInitScriptListResponse struct {
	Scripts []GlobalInitScriptInfo `json:"scripts"`
}

// List returns a list of registered global init scripts
func (a GlobalInitScriptsAPI) List() ([]GlobalInitScriptInfo, error) {
	var giss globalInitScriptListResponse
	err := a.client.Get(a.context, "/global-init-scripts", nil, &giss)
	return giss.Scripts, err
}

// Get returns information about specific global init scripts
func (a GlobalInitScriptsAPI) Get(scriptID string) (initScript GlobalInitScriptInfo, err error) {
	err = a.client.Get(a.context, "/global-init-scripts/"+scriptID, nil, &initScript)
	return
}

// Delete deletes specific global init scripts
func (a GlobalInitScriptsAPI) Delete(scriptID string) error {
	request := map[string]string{
		"script_id": scriptID,
	}
	return a.client.Delete(a.context, "/global-init-scripts/"+scriptID, request)
}

// Create creates the global init script from the given payload.
func (a GlobalInitScriptsAPI) Create(payload GlobalInitScriptPayload) (string, error) {
	var response globalInitScriptCreateResponse
	err := a.client.Post(a.context, "/global-init-scripts", payload, &response)
	return response.ScriptID, err
}

// Update updates the specific global init script from the given payload.
func (a GlobalInitScriptsAPI) Update(scriptID string, payload GlobalInitScriptPayload) error {
	return a.client.Patch(a.context, "/global-init-scripts/"+scriptID, payload)
}
