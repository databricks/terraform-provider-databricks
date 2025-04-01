package aws

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// InstanceProfileInfo contains the ARN for aws instance profiles
type InstanceProfileInfo struct {
	InstanceProfileArn    string `json:"instance_profile_arn"`
	IamRoleArn            string `json:"iam_role_arn,omitempty"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`
	SkipValidation        bool   `json:"skip_validation,omitempty" tf:"computed"`
}

// InstanceProfileList ...
type InstanceProfileList struct {
	InstanceProfiles []InstanceProfileInfo `json:"instance_profiles,omitempty"`
}

// NewInstanceProfilesAPI creates InstanceProfilesAPI instance from provider meta
func NewInstanceProfilesAPI(ctx context.Context, m any) InstanceProfilesAPI {
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
	err = &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message: fmt.Sprintf("Instance profile with name: %s not found in "+
			"list of instance profiles in the workspace!", instanceProfileARN),
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
	return a.client.Post(a.context, "/instance-profiles/remove", map[string]any{
		"instance_profile_arn": instanceProfileARN,
	}, nil)
}

// Update updates the IAM role ARN of an existing instance profile
func (a InstanceProfilesAPI) Update(ipi InstanceProfileInfo) error {
	data := map[string]any{
		"instance_profile_arn": ipi.InstanceProfileArn,
		"iam_role_arn":         ipi.InstanceProfileArn,
	}
	if ipi.IamRoleArn != "" {
		data["iam_role_arn"] = ipi.IamRoleArn
	}
	return a.client.Post(a.context, "/instance-profiles/edit", data, nil)
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
func ResourceInstanceProfile() common.Resource {
	instanceProfileSchema := common.StructToSchema(InstanceProfileInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["instance_profile_arn"].ValidateDiagFunc = ValidArn
			m["iam_role_arn"].ValidateDiagFunc = ValidArn
			m["skip_validation"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if old == "false" && new == "true" {
					return true
				}
				return false
			}
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
			common.DataToStructPointer(d, instanceProfileSchema, &profile)
			if err := NewInstanceProfilesAPI(ctx, c).Create(profile); err != nil {
				return err
			}
			d.SetId(profile.InstanceProfileArn)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewInstanceProfilesAPI(ctx, c).Delete(d.Id())
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var profile InstanceProfileInfo
			common.DataToStructPointer(d, instanceProfileSchema, &profile)
			return NewInstanceProfilesAPI(ctx, c).Update(profile)
		},
	}
}

// ValidArn validate if it's valid instance profile or role ARN
func ValidArn(v any, c cty.Path) diag.Diagnostics {
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
	var arnType string
	switch c[0].(cty.GetAttrStep).Name {
	case "instance_profile_arn", "instance_profile_id":
		arnType = "instance-profile"
	case "iam_role_arn":
		arnType = "role"
	default:
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Unknown attribute",
				Detail:        "ARN type associated with attribute is not known",
			},
		}
	}
	if s == "" && arnType == "role" {
		return nil
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
	if !strings.HasPrefix(arnSections[5], arnType) {
		return diag.Diagnostics{
			diag.Diagnostic{
				AttributePath: c,
				Summary:       "Invalid ARN",
				Detail:        fmt.Sprintf("Not a %s ARN: %s", arnType, v),
			},
		}
	}
	return nil
}
