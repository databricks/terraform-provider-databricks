package model

// TokenRequest asks for a token
type TokenRequest struct {
	LifetimeSeconds int32  `json:"lifetime_seconds,omitempty"`
	Comment         string `json:"comment,omitempty"`
}

// TokenResponse is a struct that contains information about token that is created from the create tokens api
type TokenResponse struct {
	TokenValue string     `json:"token_value,omitempty"`
	TokenInfo  *TokenInfo `json:"token_info,omitempty"`
}

// TokenInfo is a struct that contains metadata about a given token
type TokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}
