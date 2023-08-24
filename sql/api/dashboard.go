package api

import "encoding/json"

// Dashboard ...
type Dashboard struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Tags      []string          `json:"tags,omitempty"`
	Widgets   []json.RawMessage `json:"widgets,omitempty"`
	Parent    string            `json:"parent,omitempty"`
	CreatedAt string            `json:"created_at,omitempty"`
	UpdatedAt string            `json:"updated_at,omitempty"`
}
