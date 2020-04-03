package model

//go:generate easytags $GOFILE

type TokenResponse struct {
	TokenValue string     `json:"token_value,omitempty"`
	TokenInfo  *TokenInfo `json:"token_info,omitempty"`
}

type TokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

//
//{
//"token_value":"dapideadbeefdeadbeefdeadbeefdeadbeef",
//"token_info": {
//"token_id":"5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
//"creation_time":1513120516294,
//"expiry_time":1513120616294,
//"comment":"this is an example token"
//}
//}
