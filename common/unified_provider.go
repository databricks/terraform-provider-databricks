package common

import "github.com/databricks/databricks-sdk-go/marshal"

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
}

func (p *ProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, p)
}

func (p ProviderConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(p)
}
