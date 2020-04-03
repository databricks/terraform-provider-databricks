package model

type CommandResults struct {
	ResultType   string      `json:"resultType,omitempty"`
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
