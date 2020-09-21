package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Pair defines an ID pair
type Pair struct {
	Left, Right string
}

// NewPairID creates new ID pair
func NewPairID(left, right string) *Pair {
	return &Pair{left, right}
}

// Schema of paired fields
func (p *Pair) Schema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{}
	s[p.Left] = &schema.Schema{Type: schema.TypeString, ForceNew: true, Required: true}
	s[p.Right] = &schema.Schema{Type: schema.TypeString, ForceNew: true, Required: true}
	return s
}

// Unpack ID into two strings and set data
func (p *Pair) Unpack(d *schema.ResourceData) (string, string, error) {
	parts := strings.SplitN(d.Id(), "|", 2)
	if parts[0] == "" {
		return "", "", fmt.Errorf("%s cannot be empty", p.Left)
	}
	if parts[1] == "" {
		return "", "", fmt.Errorf("%s cannot be empty", p.Right)
	}
	err := d.Set(p.Left, parts[0])
	if err != nil {
		return "", "", err
	}
	err = d.Set(p.Right, parts[1])
	if err != nil {
		return "", "", err
	}
	return parts[0], parts[1], nil
}

// ReadContext helper function
func (p *Pair) ReadContext(d *schema.ResourceData, do func(left, right string) error) diag.Diagnostics {
	left, right, err := p.Unpack(d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = do(left, right)
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// Pack data attributes to ID
func (p *Pair) Pack(d *schema.ResourceData) {
	d.SetId(fmt.Sprintf("%s|%s", d.Get(p.Left), d.Get(p.Right)))
}

// ResourceGroupInstanceProfile defines group role resource
func ResourceGroupInstanceProfile() *schema.Resource {
	p := NewPairID("group_id", "instance_profile_id")
	s := p.Schema()
	// nolint temporary disable
	s["instance_profile_id"].ValidateFunc = ValidateInstanceProfileARN
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return p.ReadContext(d, func(groupID, roleARN string) error {
			group, err := NewGroupsAPI(m).Read(groupID)
			if err == nil && !group.HasRole(roleARN) {
				return common.APIError{ErrorCode: "NOT_FOUND", StatusCode: 404}
			}
			return err
		})
	}
	return &schema.Resource{
		Schema:      s,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			groupID := d.Get("group_id").(string)
			roleARN := d.Get("instance_profile_id").(string)
			err := NewGroupsAPI(m).PatchR(groupID, scimPatchRequest("add", "roles", roleARN))
			if err != nil {
				return diag.FromErr(err)
			}
			p.Pack(d)
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			groupID, roleARN, err := p.Unpack(d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewGroupsAPI(m).PatchR(groupID, scimPatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
	}
}

// ValidateInstanceProfileARN is a ValidateFunc that ensures the role id is a valid aws iam instance profile arn
func ValidateInstanceProfileARN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)

	if v == "" {
		return nil, []error{fmt.Errorf("%s is empty got: %s, must be an aws instance profile arn", key, v)}
	}

	// Parse and verify instance profiles
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return nil, []error{fmt.Errorf("%s is invalid got: %s received error: %w", key, v, err)}
	}
	// Verify instance profile resource type, Resource gets parsed as instance-profile/<profile-name>
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return nil, []error{fmt.Errorf("%s must be an instance profile resource, got: %s in %s",
			key, instanceProfileArn.Resource, v)}
	}
	return nil, nil
}
