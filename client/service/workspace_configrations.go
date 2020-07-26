package service

import (
	"encoding/json"
	"net/http"
)

// WorkspaceConfAPI exposes the workspace configurations API
type WorkspaceConfAPI struct {
	Client *DBApiClient
}

// Update will handle creation of new values as well as deletes. Deleting just implies that a value of "" is sent with
// the appropriate key
func (a WorkspaceConfAPI) Update(workspaceConfMap map[string]string) error {
	_, err := a.Client.performQuery(http.MethodPatch, "/preview/workspace-conf", "2.0", nil, workspaceConfMap, nil)
	return err
}

// Read just returns back a map of keys and values which keys are the configuration items and values are the settings
func (a WorkspaceConfAPI) Read(keys string) (map[string]string, error) {
	var wsConfResp map[string]string
	wsConfQuery := struct {
		Keys string `json:"keys,omitempty" url:"keys,omitempty"`
	}{
		Keys: keys,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/preview/workspace-conf", "2.0", nil, wsConfQuery, nil)
	if err != nil {
		return wsConfResp, err
	}
	err = json.Unmarshal(resp, &wsConfResp)
	return wsConfResp, err
}
