package repos

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func isOnlyOneGitCredentialError(err error) bool {
	errStr := err.Error()
	return (strings.Contains(errStr, "Only one Git credential is supported ") && strings.Contains(errStr, " at this time")) ||
		strings.Contains(errStr, "Only one credential per provider is allowed")
}

func ResourceGitCredential() common.Resource {
	s := common.StructToSchema(workspace.CreateCredentialsRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["force"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
		s["personal_access_token"].DefaultFunc = schema.MultiEnvDefaultFunc([]string{
			"GITHUB_TOKEN",               // https://registry.terraform.io/providers/integrations/github/latest/docs
			"GITLAB_TOKEN",               // https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs
			"AZDO_PERSONAL_ACCESS_TOKEN", // https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs
		}, nil)
		return s
	})

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var req workspace.CreateCredentialsRequest
			common.DataToStructPointer(d, s, &req)
			resp, err := w.GitCredentials.Create(ctx, req)

			if err != nil {
				if !d.Get("force").(bool) || !isOnlyOneGitCredentialError(err) {
					return err
				}
				creds, err := w.GitCredentials.ListAll(ctx)
				if err != nil {
					return err
				}
				if len(creds) != 1 {
					return fmt.Errorf("list of credentials is either empty or have more than one entry (%d)", len(creds))
				}
				var req workspace.UpdateCredentialsRequest
				common.DataToStructPointer(d, s, &req)
				req.CredentialId = creds[0].CredentialId

				err = w.GitCredentials.Update(ctx, req)
				if err != nil {
					return err
				}
				resp.CredentialId = creds[0].CredentialId
			}
			d.SetId(fmt.Sprintf("%d", resp.CredentialId))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			resp, err := w.GitCredentials.Get(ctx, workspace.GetCredentialsRequest{CredentialId: cred_id})
			if err != nil {
				return err
			}
			d.Set("git_provider", resp.GitProvider)
			d.Set("git_username", resp.GitUsername)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var req workspace.UpdateCredentialsRequest

			common.DataToStructPointer(d, s, &req)
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			req.CredentialId = cred_id
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			if d.HasChange("is_default_for_provider") {
				req.ForceSendFields = append(req.ForceSendFields, "IsDefaultForProvider")
			}
			if d.HasChange("name") {
				req.ForceSendFields = append(req.ForceSendFields, "Name")
			}
			return w.GitCredentials.Update(ctx, req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			return w.GitCredentials.DeleteByCredentialId(ctx, cred_id)
		},
	}
}
