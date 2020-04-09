package model

type CommandResults struct {
	ResultType   string      `json:"resultType,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Cause        string      `json:"cause,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Schema       interface{} `json:"schema,omitempty"`
	Truncated    bool        `json:"truncated,omitempty"`
	IsJsonSchema bool        `json:"isJsonSchema,omitempty"`
}

type Command struct {
	ID      string          `json:"id,omitempty"`
	Status  string          `json:"status,omitempty"`
	Results *CommandResults `json:"results,omitempty"`
}

type ExecutionContext struct {
	ContextId string   `json:"contextId,omitempty"`
	ClusterId string   `json:"clusterId,omitempty"`
	Language  Language `json:"language,omitempty"`
}
