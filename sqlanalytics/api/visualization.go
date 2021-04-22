package api

import "encoding/json"

// Visualization ...
type Visualization struct {
	ID          int             `json:"id,omitempty"`
	QueryID     string          `json:"query_id,omitempty"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Options     json.RawMessage `json:"options,omitempty"`
}
