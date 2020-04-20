package model

// CommandResults is the out put when the command finishes in API 1.2
type CommandResults struct {
	ResultType   string      `json:"resultType,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Cause        string      `json:"cause,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Schema       interface{} `json:"schema,omitempty"`
	Truncated    bool        `json:"truncated,omitempty"`
	IsJSONSchema bool        `json:"isJsonSchema,omitempty"`
}

// Command is the struct that contains what the 1.2 api returns for the commands api
type Command struct {
	ID      string          `json:"id,omitempty"`
	Status  string          `json:"status,omitempty"`
	Results *CommandResults `json:"results,omitempty"`
}
