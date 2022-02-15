package repos

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GitCredentialsAPI exposes the Git Credentials API
type GitCredentialsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// GitCredentialsAPI creates GitCredentialsAPI instance from provider meta
func NewGitCredentialsAPI(ctx context.Context, m interface{}) GitCredentialsAPI {
	return GitCredentialsAPI{m.(*common.DatabricksClient), ctx}
}

type GitCredentialRequest struct {
	PAT      string `json:"personal_access_token" tf:"sensitive"`
	Provider string `json:"git_provider" tf:"suppress_diff"`
	UserName string `json:"git_username"`
}

type GitCredentialResponse struct {
	ID       int64  `json:"credential_id"`
	UserName string `json:"git_username"`
	Provider string `json:"git_provider"`
}

type GitCredentialList struct {
	Credentials []GitCredentialResponse `json:"credentials,omitempty"`
}

// ID returns job id as string
func (r GitCredentialResponse) GitCredentialID() string {
	return fmt.Sprintf("%d", r.ID)
}

func (a GitCredentialsAPI) Delete(id string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/git-credentials/%s", id), nil)
}

func (a GitCredentialsAPI) List() ([]GitCredentialResponse, error) {
	resp := GitCredentialList{}
	err := a.client.Get(a.context, "/git-credentials", nil, &resp)
	return resp.Credentials, err
}

func (a GitCredentialsAPI) Read(id string) (resp GitCredentialResponse, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/git-credentials/%s", id), nil, &resp)
	return
}

func (a GitCredentialsAPI) Create(req GitCredentialRequest) (resp GitCredentialResponse, err error) {
	err = a.client.Post(a.context, "/git-credentials", &req, &resp)
	return
}

func (a GitCredentialsAPI) Update(id string, req GitCredentialRequest) error {
	return a.client.Patch(a.context, fmt.Sprintf("/git-credentials/%s", id), &req)
}

func ResourceGitCredential() *schema.Resource {
	s := common.StructToSchema(GitCredentialRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["force"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
		return s
	})

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			api := NewGitCredentialsAPI(ctx, c)
			var req GitCredentialRequest
			common.DataToStructPointer(d, s, &req)
			resp, err := api.Create(req)

			if err != nil {
				if !d.Get("force").(bool) || !strings.HasPrefix(err.Error(), "Only one Git credential is supported at this time") {
					return err
				}
				creds, err := api.List()
				if err != nil {
					return err
				}
				if len(creds) != 1 {
					return fmt.Errorf("list of credentials is either empty or have more than one entry (%d)", len(creds))
				}
				resp = creds[0]
				err = api.Update(resp.GitCredentialID(), req)
				if err != nil {
					return err
				}
			}
			d.SetId(resp.GitCredentialID())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			resp, err := NewGitCredentialsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("git_provider", resp.Provider)
			d.Set("git_username", resp.UserName)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var req GitCredentialRequest
			common.DataToStructPointer(d, s, &req)
			return NewGitCredentialsAPI(ctx, c).Update(d.Id(), req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGitCredentialsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
