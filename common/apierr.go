package common

// APIErrorBody represents an API error returned by Databricks API.
//
// Deprecated: this class is meant to disappear as the Terraform provider
// progressively moves to use the service clients provided by the SDK. Clients
// should not use this class for any other purpose than testing code that
// mocks the behavior of Databricks services.
type APIErrorBody struct {
	ErrorCode  string        `json:"error_code,omitempty"`
	Message    string        `json:"message,omitempty"`
	Details    []ErrorDetail `json:"details,omitempty"`
	ScimDetail string        `json:"detail,omitempty"`
	ScimStatus string        `json:"status,omitempty"`
	ScimType   string        `json:"scimType,omitempty"`
	API12Error string        `json:"error,omitempty"`
}

// ErrorDetail represents the details of an API error.
//
// Deprecated: this struct is meant to disappear as the Terraform provider
// progressively moves to use the service clients provided by the SDK. Clients
// should not use this struct for any other purpose than testing code that
// mocks the behavior of Databricks services.
type ErrorDetail struct {
	Type     string            `json:"@type,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Domain   string            `json:"domain,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
