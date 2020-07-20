package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// IPAccessListsAPI exposes the instance pools api
type IPAccessListsAPI struct {
	Client *DBApiClient
}

// Create creates the IP Access List to given the instance pool configuration
func (a IPAccessListsAPI) Create(ipAddresses []string, label string, listType model.IPAccessListType) (model.IPAccessListStatus, error) {
	cr := model.CreateIPAccessListRequest{}
	cr.IPAddresses = ipAddresses
	cr.Label = label
	cr.ListType = listType

	resp, err := a.Client.performQuery(
		http.MethodPost,
		"/preview/ip-access-lists",
		"2.0",
		nil, cr, nil)

	wrapper := model.IPAccessListStatusWrapper{}
	if err != nil {
		return wrapper.IPAccessList, err
	}

	err = json.Unmarshal(resp, &wrapper)
	return wrapper.IPAccessList, err
}

func (a IPAccessListsAPI) Update(objectID string, label string, listType model.IPAccessListType, ipAddresses []string, enabled bool) (model.IPAccessListStatus, error) {
	ur := model.IPAccessListUpdateRequest{}
	ur.Enabled = enabled
	ur.IPAddresses = ipAddresses
	ur.Label = label
	ur.ListType = listType

	status := model.IPAccessListStatus{}

	resp, err := a.Client.performQuery(
		http.MethodPut,
		"/preview/ip-access-lists/"+objectID,
		"2.0",
		nil, ur, nil)

	if err != nil {
		return status, err
	}
	err = json.Unmarshal(resp, &status)
	return status, err
}

func (a IPAccessListsAPI) Delete(objectID string) error {
	_, err := a.Client.performQuery(
		http.MethodDelete,
		"/preview/ip-access-lists/"+objectID,
		"2.0",
		nil, nil, nil)
	return err
}

func (a IPAccessListsAPI) Read(objectID string) (model.IPAccessListStatus, error) {
	resp, err := a.Client.performQuery(http.MethodGet,
		"/preview/ip-access-lists/"+objectID,
		"2.0", nil, nil, nil)

	wrapper := model.IPAccessListStatusWrapper{}
	if err != nil {
		return wrapper.IPAccessList, err
	}
	err = json.Unmarshal(resp, &wrapper)
	return wrapper.IPAccessList, err
}

func (a IPAccessListsAPI) List() (model.ListIPAccessListsResponse, error) {
	resp, err := a.Client.performQuery(http.MethodGet,
		"/preview/ip-access-lists",
		"2.0", nil, nil, nil)
	listResponse := model.ListIPAccessListsResponse{}
	if err != nil {
		return listResponse, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse, err
}
