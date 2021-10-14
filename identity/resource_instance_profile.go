package identity

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// InstanceProfileInfo contains the ARN for aws instance profiles
type InstanceProfileInfo struct {
	InstanceProfileArn    string `json:"instance_profile_arn,omitempty"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`
	SkipValidation        bool   `json:"skip_validation,omitempty" tf:"computed"`
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
func (a InstanceProfilesAPI) Create(ipi InstanceProfileInfo) error {
	return a.client.Post(a.context, "/instance-profiles/add", ipi, nil)
}

// Read returns the ARN back if it exists on the Databricks workspace
func (a InstanceProfilesAPI) Read(instanceProfileARN string) (result InstanceProfileInfo, err error) {
	instanceProfiles, err := a.List()
	if err != nil {
		return
	}
	for _, profile := range instanceProfiles {
		if profile.InstanceProfileArn == instanceProfileARN {
			result = profile
			return
		}
	}
	err = common.APIError{
		ErrorCode: "NOT_FOUND",
		Message: fmt.Sprintf("Instance profile with name: %s not found in "+
			"list of instance profiles in the workspace!", instanceProfileARN),
		Resource:   "/api/2.0/instance-profiles/list",
		StatusCode: http.StatusNotFound,
	}
	return
}

// List lists all the instance profiles in the workspace
func (a InstanceProfilesAPI) List() ([]InstanceProfileInfo, error) {
	var instanceProfilesArnList InstanceProfileList
	err := a.client.Get(a.context, "/instance-profiles/list", nil, &instanceProfilesArnList)
	return instanceProfilesArnList.InstanceProfiles, err
}

// Delete deletes the instance profile given an instance profile arn
func (a InstanceProfilesAPI) Delete(instanceProfileARN string) error {
	return a.client.Post(a.context, "/instance-profiles/remove", map[string]interface{}{
		"instance_profile_arn": instanceProfileARN,
	}, nil)
}

// IsRegistered checks if instance profile exists
func (a InstanceProfilesAPI) IsRegistered(arn string) bool {
	if _, err := a.Read(arn); err == nil {
		return true
	}
	return false
}

// Synchronized test helper for working with only single instance profile
func (a InstanceProfilesAPI) Synchronized(arn string, testCallback func() bool) {
	timeout := 30 * time.Minute
	err := resource.RetryContext(a.context, timeout,
		func() *resource.RetryError {
			list, err := a.List()
			if err != nil {
				return resource.NonRetryableError(err)
			}
			currentTest := common.Current.GetOrUnknown(a.context)
			for _, ip := range list {
				if ip.InstanceProfileArn == arn {
					log.Printf("[INFO] %s: Waiting to acquire instance profile", currentTest)
					return resource.RetryableError(fmt.Errorf(
						"%s is registered, waiting to release", arn))
				}
			}
			cbError := resource.RetryContext(a.context, timeout, func() *resource.RetryError {
				if a.IsRegistered(arn) {
					log.Printf("[INFO] %s: Waiting to acquire instance profile", currentTest)
					return resource.RetryableError(fmt.Errorf("%s: Waiting to acquire", currentTest))
				}
				if !testCallback() {
					log.Printf("[INFO] %s: Callback returned false", currentTest)
					return resource.RetryableError(fmt.Errorf("%s: Callback returned false", currentTest))
				}
				log.Printf("[INFO] %s: Successfully tested instance profile", currentTest)
				if _, err = a.Read(arn); err == nil {
					log.Printf("[INFO] %s: Didn't release instance profile", currentTest)
				}
				return nil
			})
			if cbError != nil {
				return resource.RetryableError(cbError)
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
}

// ResourceInstanceProfile manages Instance Profile ARN binding
func ResourceInstanceProfile() *schema.Resource {
	instanceProfileSchema := common.StructToSchema(InstanceProfileInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["instance_profile_arn"].ValidateDiagFunc = ValidInstanceProfile
			return m
		})
	return common.Resource{
		Schema: instanceProfileSchema,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			profile, err := NewInstanceProfilesAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(profile, instanceProfileSchema, d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var profile InstanceProfileInfo
			if err := common.DataToStructPointer(d, instanceProfileSchema, &profile); err != nil {
				return err
			}
			if err := NewInstanceProfilesAPI(ctx, c).Create(profile); err != nil {
				return err
			}
			d.SetId(profile.InstanceProfileArn)
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
	if !strings.HasPrefix(s, "arn:") {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        "Invalid prefix",
			},
		}
	}
	arnSections := strings.SplitN(s, ":", 6)
	if len(arnSections) != 6 {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        "Incorrect number of sections",
			},
		}
	}
	if !strings.HasPrefix(arnSections[5], "instance-profile") {
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
