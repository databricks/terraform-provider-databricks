package model

type GroupMember struct {
	Display string `json:"display,omitempty"`
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type MemberListItem struct {
	Value string `json:"value,omitempty"`
}

type Group struct {
	ID          string        `json:"id,omitempty"`
	Schemas     []URN         `json:"schemas,omitempty"`
	DisplayName string        `json:"displayName,omitempty"`
	Members     []GroupMember `json:"members,omitempty"`
	Groups      []interface{} `json:"groups,omitempty"`
}

type GroupPatchRequest struct {
	Schemas    []URN                  `json:"schemas,omitempty"`
	Operations []GroupPatchOperations `json:"Operations,omitempty"`
}
