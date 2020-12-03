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
func (a IPAccessListsAPI) Create(ipAddresses []string, label string, listType ipAccessListType) (status ipAccessListStatus, err error) {
	cr := createIPAccessListRequest{}
	cr.IPAddresses = ipAddresses
	cr.Label = label
	cr.ListType = listType

	wrapper := ipAccessListStatusWrapper{}
	err = a.Client.Post("/ip-access-lists", cr, &wrapper)

	status = wrapper.IPAccessList
	return
}

func (a IPAccessListsAPI) Update(objectID string, label string, listType ipAccessListType, ipAddresses []string, enabled bool) (err error) {
	ur := ipAccessListUpdateRequest{}
	ur.Enabled = enabled
	ur.IPAddresses = ipAddresses
	ur.Label = label
	ur.ListType = listType

	err = a.Client.Put("/ip-access-lists/"+objectID, ur)

	return
}

func (a IPAccessListsAPI) Delete(objectID string) (err error) {
	err = a.Client.Delete("/ip-access-lists/"+objectID, nil)
	return
}

func (a IPAccessListsAPI) Read(objectID string) (status ipAccessListStatus, err error) {
	wrapper := ipAccessListStatusWrapper{}
	err = a.Client.Get("/ip-access-lists/"+objectID, nil, &wrapper)
	status = wrapper.IPAccessList
	return
}

func (a IPAccessListsAPI) List() (listResponse listIPAccessListsResponse, err error) {
	listResponse = listIPAccessListsResponse{}
	err = a.Client.Get("/ip-access-lists", &listResponse, nil)

	return
}
