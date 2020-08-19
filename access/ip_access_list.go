package access

import "github.com/databrickslabs/databricks-terraform/common"

// IPAccessListsAPI exposes the IP ACL Lists api
type IPAccessListsAPI struct {
	Client *common.DatabricksClient
}

func NewIPAccessListsAPI(m interface{}) IPAccessListsAPI {
	return IPAccessListsAPI{Client: m.(*common.DatabricksClient)}
}

// Create creates the IP Access List to given the instance pool configuration
func (a IPAccessListsAPI) Create(ipAddresses []string, label string, listType IPAccessListType) (status IPAccessListStatus, err error) {
	cr := CreateIPAccessListRequest{}
	cr.IPAddresses = ipAddresses
	cr.Label = label
	cr.ListType = listType

	wrapper := IPAccessListStatusWrapper{}
	err = a.Client.Post("/preview/ip-access-lists", cr, &wrapper)

	status = wrapper.IPAccessList
	return
}

func (a IPAccessListsAPI) Update(objectID string, label string, listType IPAccessListType, ipAddresses []string, enabled bool) (err error) {
	ur := IPAccessListUpdateRequest{}
	ur.Enabled = enabled
	ur.IPAddresses = ipAddresses
	ur.Label = label
	ur.ListType = listType

	err = a.Client.Put("/preview/ip-access-lists/"+objectID, ur)

	return
}

func (a IPAccessListsAPI) Delete(objectID string) (err error) {
	err = a.Client.Delete("/preview/ip-access-lists/"+objectID, nil)
	return
}

func (a IPAccessListsAPI) Read(objectID string) (status IPAccessListStatus, err error) {
	wrapper := IPAccessListStatusWrapper{}
	err = a.Client.Get("/preview/ip-access-lists/"+objectID, nil, &wrapper)
	status = wrapper.IPAccessList
	return
}

func (a IPAccessListsAPI) List() (listResponse ListIPAccessListsResponse, err error) {
	listResponse = ListIPAccessListsResponse{}
	err = a.Client.Get("/preview/ip-access-lists", &listResponse, nil)

	return
}
