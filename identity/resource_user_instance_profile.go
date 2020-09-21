package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUserInstanceProfile binds user and instance profile
func ResourceUserInstanceProfile() *schema.Resource {
	p := NewPairID("user_id", "instance_profile_id")
	s := p.Schema()
	// nolint temporary disable
	s["instance_profile_id"].ValidateFunc = ValidateInstanceProfileARN

	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return p.ReadContext(d, func(userID, roleARN string) error {
			user, err := NewUsersAPI(m).Read(userID)
			if err == nil && !user.HasRole(roleARN) {
				return common.APIError{ErrorCode: "NOT_FOUND", StatusCode: 404}
			}
			return err
		})
	}
	return &schema.Resource{
		Schema:      s,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			userID := d.Get("user_id").(string)
			roleARN := d.Get("instance_profile_id").(string)
			err := NewUsersAPI(m).PatchR(userID, scimPatchRequest("add", "roles", roleARN))
			if err != nil {
				return diag.FromErr(err)
			}
			p.Pack(d)
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			userID, roleARN, err := p.Unpack(d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewUsersAPI(m).PatchR(userID, scimPatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
