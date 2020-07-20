package model

type IPAccessListType string

const (
	BlackList IPAccessListType = "BLACKLIST"
	WhiteList IPAccessListType = "WHITELIST"
)

type IPAccessListStatus struct {
	ListID        string           `json:"list_id,omitempty"`
	Label         string           `json:"label,omitempty"`
	ListType      IPAccessListType `json:"list_type,omitempty"`
	IPAddresses   []string         `json:"ip_addresses,omitempty"`
	AddressCount  int              `json:"address_count,omitempty"`
	CreatedAt     int64            `json:"created_at,omitempty"`
	CreatorUserID int64            `json:"creator_user_id,omitempty"`
	UpdatedAt     int64            `json:"updated_at,omitempty"`
	UpdatorUserID int64            `json:"updator_user_id,omitempty"`
	Enabled       bool             `json:"enabled,omitempty"`
}

type IPAccessListStatusWrapper struct {
	IPAccessList IPAccessListStatus `json:"ip_access_list,omitempty"`
}

// Add an IP access list
type CreateIPAccessListRequest struct {
	Label       string           `json:"label,omitempty"`
	ListType    IPAccessListType `json:"list_type,omitempty"`
	IPAddresses []string         `json:"ip_addresses,omitempty"`
}

// List: Get all IP access lists
// Norequest needed--just a get
type ListIPAccessListsResponse struct {
	ListIPAccessListsResponse []IPAccessListStatus `json:"ip_access_lists,omitempty"`
}

// Update an IP access list
type IPAccessListUpdateRequest struct {
	Label       string           `json:"label,omitempty"`
	ListType    IPAccessListType `json:"list_type,omitempty"`
	IPAddresses []string         `json:"ip_addresses,omitempty"`
	Enabled     bool             `json:"enabled,omitempty"`
}
