package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// InstanceProfilesAPI exposes the instance profiles api on the AWS deployment of Databricks
type InstanceProfilesAPI struct {
	Client *DBApiClient
}

// Create creates an instance profile record on Databricks
func (a InstanceProfilesAPI) Create(instanceProfileARN string, skipValidation bool) error {
	addInstanceProfileRequest := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
		SkipValidation     bool   `json:"skip_validation,omitempty" url:"skip_validation,omitempty"`
	}{
		InstanceProfileArn: instanceProfileARN,
		SkipValidation:     skipValidation,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/add", "2.0", nil, addInstanceProfileRequest, nil)
	return err
}

// Read returns the ARN back if it exists on the Databricks workspace
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

	return response, fmt.Errorf("Instance profile with name: %s not found in "+
		"list of instance profiles in the workspace!", instanceProfileARN)
}

// List lists all the instance profiles in the workspace
func (a InstanceProfilesAPI) List() ([]model.InstanceProfileInfo, error) {
	var instanceProfilesArnList struct {
		InstanceProfiles []model.InstanceProfileInfo `json:"instance_profiles,omitempty" url:"instance_profiles,omitempty"`
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/instance-profiles/list", "2.0", nil, nil, nil)
	if err != nil {
		return instanceProfilesArnList.InstanceProfiles, err
	}

	err = json.Unmarshal(resp, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

// Delete deletes the instance profile given an instance profile arn
func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	deleteInstanceProfileRequest := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
	}{
		InstanceProfileArn: instanceProfileARN,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/remove", "2.0", nil, deleteInstanceProfileRequest, nil)
	return err
}
