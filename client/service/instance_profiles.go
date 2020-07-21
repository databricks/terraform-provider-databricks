package service

import (
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// InstanceProfilesAPI exposes the instance profiles api on the AWS deployment of Databricks
type InstanceProfilesAPI struct {
	client *DatabricksClient
}

// Create creates an instance profile record on Databricks
func (a InstanceProfilesAPI) Create(instanceProfileARN string, skipValidation bool) error {
	return a.client.post("/instance-profiles/add", map[string]interface{}{
		"instance_profile_arn": instanceProfileARN,
		"skip_validation":      skipValidation,
	}, nil)
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
	return response, APIError{
		ErrorCode: "NOT_FOUND",
		Message: fmt.Sprintf("Instance profile with name: %s not found in "+
			"list of instance profiles in the workspace!", instanceProfileARN),
		Resource:   "/api/2.0/instance-profiles/list",
		StatusCode: http.StatusNotFound,
	}
}

// List lists all the instance profiles in the workspace
func (a InstanceProfilesAPI) List() ([]model.InstanceProfileInfo, error) {
	var instanceProfilesArnList struct {
		InstanceProfiles []model.InstanceProfileInfo `json:"instance_profiles,omitempty" url:"instance_profiles,omitempty"`
	}
	err := a.client.get("/instance-profiles/list", nil, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

// Delete deletes the instance profile given an instance profile arn
func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	return a.client.post("/instance-profiles/remove", map[string]interface{}{
		"instance_profile_arn": instanceProfileARN,
	}, nil)
}
