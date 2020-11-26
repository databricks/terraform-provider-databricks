package access

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
)

// ipAccessListsAPI exposes the IP ACL Lists api
type ipAccessListsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewIPAccessListsAPI(m interface{}) ipAccessListsAPI {
	return ipAccessListsAPI{
		client:  m.(*common.DatabricksClient),
		context: context.TODO(),
	}
}

// Create creates the IP Access List to given the instance pool configuration
func (a ipAccessListsAPI) Create(ipAddresses []string, label string, listType ipAccessListType) (status ipAccessListStatus, err error) {
	cr := createIPAccessListRequest{}
	cr.IPAddresses = ipAddresses
	cr.Label = label
	cr.ListType = listType

	wrapper := ipAccessListStatusWrapper{}
	err = a.client.Post(a.context, "/ip-access-lists", cr, &wrapper)

	status = wrapper.IPAccessList
	return
}

func (a ipAccessListsAPI) Update(objectID string, label string, listType ipAccessListType, ipAddresses []string, enabled bool) (err error) {
	ur := ipAccessListUpdateRequest{}
	ur.Enabled = enabled
	ur.IPAddresses = ipAddresses
	ur.Label = label
	ur.ListType = listType

	err = a.client.Put(a.context, "/ip-access-lists/"+objectID, ur)

	return
}

func (a ipAccessListsAPI) Delete(objectID string) (err error) {
	err = a.client.Delete(a.context, "/ip-access-lists/"+objectID, nil)
	return
}

func (a ipAccessListsAPI) Read(objectID string) (status ipAccessListStatus, err error) {
	wrapper := ipAccessListStatusWrapper{}
	err = a.client.Get(a.context, "/ip-access-lists/"+objectID, nil, &wrapper)
	status = wrapper.IPAccessList
	return
}

func (a ipAccessListsAPI) List() (listResponse listIPAccessListsResponse, err error) {
	listResponse = listIPAccessListsResponse{}
	err = a.client.Get(a.context, "/ip-access-lists", &listResponse, nil)

	return
}
