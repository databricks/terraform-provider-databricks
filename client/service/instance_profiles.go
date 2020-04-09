package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"log"
	"net/http"
)

type InstanceProfilesAPI struct {
	Client DBApiClient
}

func (a InstanceProfilesAPI) init(client DBApiClient) InstanceProfilesAPI {
	a.Client = client
	return a
}

func (a InstanceProfilesAPI) Create(instanceProfileARN string, skipValidation bool) error {
	addInstanceProfileRequest := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
		SkipValidation     bool   `json:"skip_validation,omitempty" url:"skip_validation,omitempty"`
	}{
		InstanceProfileArn: instanceProfileARN,
		SkipValidation:     skipValidation,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/add", "2.0", nil, addInstanceProfileRequest)
	return err
}

func (a InstanceProfilesAPI) Read(instanceProfileARN string) (string, error) {
	var response string
	instanceProfiles, err := a.List()
	if err != nil {
		return response, err
	}
	for _, profile := range instanceProfiles {
		if profile.InstanceProfileArn == instanceProfileARN {
			response = profile.InstanceProfileArn
			return response, nil
		}
	}

	return response, errors.New(fmt.Sprintf("Instance profile with name: %s not found in "+
		"list of instance profiles in the workspace!", instanceProfileARN))
}

func (a InstanceProfilesAPI) List() ([]model.InstanceProfileInfo, error) {

	var instanceProfilesArnList struct {
		InstanceProfiles []model.InstanceProfileInfo `json:"instance_profiles,omitempty" url:"instance_profiles,omitempty"`
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/instance-profiles/list", "2.0", nil, nil)
	if err != nil {
		return instanceProfilesArnList.InstanceProfiles, err
	}
	log.Println(string(resp))

	err = json.Unmarshal(resp, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	deleteInstanceProfileRequest := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
	}{
		InstanceProfileArn: instanceProfileARN,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/remove", "2.0", nil, deleteInstanceProfileRequest)
	return err
}
