package api

import "encoding/json"

// Visualization ...
type Visualization struct {
	// Visualizations evolved from having integer IDs to string UUIDs.
	// This type supports either in support of a transition without breakage.
	ID stringOrInt `json:"id,omitempty"`

	QueryID     string          `json:"query_id,omitempty"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Options     json.RawMessage `json:"options,omitempty"`
	QueryPlan   json.RawMessage `json:"query_plan,omitempty"`
}
