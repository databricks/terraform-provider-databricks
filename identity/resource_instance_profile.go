package identity

import (
	"fmt"
	"log"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// InstanceProfileInfo contains the ARN for aws instance profiles
type InstanceProfileInfo struct {
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
}

// InstanceProfileList ...
type InstanceProfileList struct {
	InstanceProfiles []InstanceProfileInfo `json:"instance_profiles,omitempty"`
}

// NewInstanceProfilesAPI creates InstanceProfilesAPI instance from provider meta
func NewInstanceProfilesAPI(m interface{}) InstanceProfilesAPI {
	return InstanceProfilesAPI{C: m.(*common.DatabricksClient)}
}

// InstanceProfilesAPI exposes the instance profiles api on the AWS deployment of Databricks
type InstanceProfilesAPI struct {
	C *common.DatabricksClient
}

// Create creates an instance profile record on Databricks
func (a InstanceProfilesAPI) Create(instanceProfileARN string, skipValidation bool) error {
	return a.C.Post("/instance-profiles/add", map[string]interface{}{
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
	return response, common.APIError{
		ErrorCode: "NOT_FOUND",
		Message: fmt.Sprintf("Instance profile with name: %s not found in "+
			"list of instance profiles in the workspace!", instanceProfileARN),
		Resource:   "/api/2.0/instance-profiles/list",
		StatusCode: http.StatusNotFound,
	}
}

// List lists all the instance profiles in the workspace
func (a InstanceProfilesAPI) List() ([]InstanceProfileInfo, error) {
	var instanceProfilesArnList InstanceProfileList
	err := a.C.Get("/instance-profiles/list", nil, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

// Delete deletes the instance profile given an instance profile arn
func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	return a.C.Post("/instance-profiles/remove", map[string]interface{}{
		"instance_profile_arn": instanceProfileARN,
	}, nil)
}

func ResourceInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceProfileCreate,
		Read:   resourceInstanceProfileRead,
		Delete: resourceInstanceProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"instance_profile_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skip_validation": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInstanceProfileCreate(d *schema.ResourceData, m interface{}) error {
	instanceProfileArn := d.Get("instance_profile_arn").(string)
	skipValidation := d.Get("skip_validation").(bool)
	err := NewInstanceProfilesAPI(m).Create(instanceProfileArn, skipValidation)
	if err != nil {
		return err
	}
	d.SetId(instanceProfileArn)
	err = d.Set("skip_validation", skipValidation)
	if err != nil {
		return err
	}

	return resourceInstanceProfileRead(d, m)
}

func resourceInstanceProfileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	profile, err := NewInstanceProfilesAPI(m).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("instance_profile_arn", profile)
	return err
}

func resourceInstanceProfileDelete(d *schema.ResourceData, m interface{}) error {
	return NewInstanceProfilesAPI(m).Delete(d.Id())
}
