package aws

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInstanceProfiles() *schema.Resource {
	type instanceProfileData struct {
		Name    string `json:"name,omitempty" tf:"computed"`
		Arn     string `json:"arn,omitempty" tf:"computed"`
		RoleArn string `json:"role_arn,omitempty" tf:"computed"`
		IsMeta  bool   `json:"is_meta,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		InstanceProfiles []instanceProfileData `json:"instance_profiles,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		instanceProfiles, err := w.InstanceProfiles.ListAll(ctx)
		if err != nil {
			return err
		}
		for _, v := range instanceProfiles {
			arnSlices := strings.Split(v.InstanceProfileArn, "/")
			name := arnSlices[len(arnSlices)-1]
			data.InstanceProfiles = append(data.InstanceProfiles, instanceProfileData{
				Name:    name,
				Arn:     v.InstanceProfileArn,
				RoleArn: v.IamRoleArn,
				IsMeta:  v.IsMetaInstanceProfile,
			})
		}
		return nil
	})
}
