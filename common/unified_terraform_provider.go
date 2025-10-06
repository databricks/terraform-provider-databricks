package common

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id" tf:"force_new"`
}

type ProviderConfigData struct {
	WorkspaceID string `json:"workspace_id"`
}
