package service

import (
	"encoding/json"
	"net/http"
)

// WorkspaceConfAPI exposes the workspace configurations API
type WorkspaceConfAPI struct {
	Client *DBApiClient
}

// Update will handle creation of new values as well as deletes. Deleting just implies that a value of "" or
// the appropriate disable string like "false" is sent with the appropriate key
// TODO: map[string]string is the only thing accepted by the API currently.  If you send in another type, you get the response
// {
//    "error_code": "BAD_REQUEST",
//    "message": "Values must be strings"
//}
// This is the case for any key tested.  It would be worth finding any internal documentation detailing workspace-conf
func (a WorkspaceConfAPI) Update(workspaceConfMap map[string]string) (err error) {
	_, err = a.Client.performQuery(http.MethodPatch, "/preview/workspace-conf", "2.0", nil, workspaceConfMap, nil)
	return
}

// Read just returns back a map of keys and values which keys are the configuration items and values are the settings
func (a WorkspaceConfAPI) Read(keys string) (wsConfResp map[string]string, err error) {
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
	return
}
