package api

type Alert struct {
	Id      string        `json:"id,omitempty"`
	Name    string        `json:"name,omitempty"`
	Options *AlertOptions `json:"options,omitempty"`
	Parent  string        `json:"parent,omitempty"`
	Query   *Query        `json:"query,omitempty"`
	Rearm   int           `json:"rearm,omitempty"`
}

// Alert configuration options.
type AlertOptions struct {
	Column        string `json:"column"`
	CustomBody    string `json:"custom_body,omitempty"`
	CustomSubject string `json:"custom_subject,omitempty"`
	Muted         bool   `json:"muted,omitempty"`
	Op            string `json:"op"`
	Value         string `json:"value"`
}

type CreateAlert struct {
	Name    string        `json:"name"`
	Options *AlertOptions `json:"options"`
	Parent  string        `json:"parent,omitempty"`
	QueryId string        `json:"query_id"`
	Rearm   int           `json:"rearm,omitempty"`
}
