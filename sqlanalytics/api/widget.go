package api

import "encoding/json"

// Widget ...
type Widget struct {
	ID int `json:"id,omitempty"`

	// Widgets are part of a dashboard.
	DashboardID string `json:"dashboard_id"`

	// They are either linked to a query visualization or embed a piece of Markdown text.
	// These fields are mutually exclusive and must be `null` if they don't apply.
	VisualizationID *int    `json:"visualization_id"`
	Text            *string `json:"text"`

	// This field is no longer in use, but is still required as part of the schema.
	// It's OK that the field value is 0 everywhere.
	Width int `json:"width"`

	Options struct {
		ParameterMapping map[string]WidgetParameterMapping `json:"parameterMappings"`
		Position         *WidgetPosition                   `json:"position,omitempty"`
	} `json:"options"`

	// Fields below are set only when retrieving an existing widget.
	Visualization json.RawMessage `json:"visualization,omitempty"`
}

// WidgetPosition ...
type WidgetPosition struct {
	AutoHeight bool `json:"autoHeight"`
	SizeX      int  `json:"sizeX,omitempty"`
	SizeY      int  `json:"sizeY,omitempty"`
	PosX       int  `json:"col"`
	PosY       int  `json:"row"`
}

// WidgetParameterMapping ...
type WidgetParameterMapping struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	MapTo string `json:"mapTo,omitempty"`

	// The type of the value depends on the type of the parameter referred to by `name`.
	Value interface{} `json:"value"`

	// This title overrides the title given to this parameter by the query, if specified.
	Title string `json:"title,omitempty"`
}
