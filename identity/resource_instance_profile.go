package identity

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
func NewInstanceProfilesAPI(ctx context.Context, m interface{}) InstanceProfilesAPI {
	return InstanceProfilesAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// InstanceProfilesAPI exposes the instance profiles api on the AWS deployment of Databricks
type InstanceProfilesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates an instance profile record on Databricks
func (a InstanceProfilesAPI) Create(instanceProfileARN string, skipValidation bool) error {
	return a.client.Post("/instance-profiles/add", map[string]interface{}{
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
	err := a.client.Get("/instance-profiles/list", nil, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

// Delete deletes the instance profile given an instance profile arn
func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	return a.client.Post("/instance-profiles/remove", map[string]interface{}{
		"instance_profile_arn": instanceProfileARN,
	}, nil)
}

// Synchronized test helper for working with only single instance profile
func (a InstanceProfilesAPI) Synchronized(arn string, cb func()) {
	err := resource.RetryContext(a.context, 10*time.Minute, func() *resource.RetryError {
		list, err := a.List()
		if err != nil {
			return resource.NonRetryableError(err)
		}
		for _, ip := range list {
			if ip.InstanceProfileArn == arn {
				return resource.RetryableError(fmt.Errorf(
					"%s is registered, waiting to release", arn))
			}
		}
		cb()
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// ResourceInstanceProfile manages Instance Profile ARN binding
func ResourceInstanceProfile() *schema.Resource {
	return util.CommonResource{
		Schema: map[string]*schema.Schema{
			"instance_profile_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,

				ValidateDiagFunc: ValidInstanceProfile,
			},
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			profile, err := NewInstanceProfilesAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return d.Set("instance_profile_arn", profile)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ipa := d.Get("instance_profile_arn").(string)
			if err := NewInstanceProfilesAPI(ctx, c).Create(ipa, false); err != nil {
				return err
			}
			d.SetId(ipa)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewInstanceProfilesAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}

// ValidInstanceProfile validate if it's valid instance profile ARN
func ValidInstanceProfile(v interface{}, c cty.Path) diag.Diagnostics {
	s, ok := v.(string)
	if !ok {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        "Not a string",
			},
		}
	}
	instanceProfileArn, err := arn.Parse(s)
	if err != nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        err.Error(),
			},
		}
	}
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        fmt.Sprintf("Not an instance profile ARN: %s", v),
			},
		}
	}
	return nil
}
