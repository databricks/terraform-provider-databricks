package repos

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

func isOnlyOneGitCredentialForProviderError(err error) bool {
	errStr := err.Error()
	return errors.Is(err, apierr.ErrResourceConflict) ||
		(strings.Contains(errStr, "Only one Git credential is supported ") && strings.Contains(errStr, " at this time"))
}

// getPersonalAccessToken returns `personal_access_token_wo` from the raw config if set,
// otherwise falls back to `personal_access_token`. Write-only values are never persisted
// in state, so they must be read from the configuration.
func getPersonalAccessToken(d *schema.ResourceData) string {
	woValue, diags := d.GetRawConfigAt(cty.GetAttrPath("personal_access_token_wo"))
	if !diags.HasError() && woValue.Type().Equals(cty.String) && !woValue.IsNull() && woValue.IsKnown() {
		return woValue.AsString()
	}
	return d.Get("personal_access_token").(string)
}

// suppressDefaultTokenWhenWriteOnly drops a `personal_access_token` sourced from the
// environment (via DefaultFunc) when `personal_access_token_wo` is configured, so the
// token is not persisted in state. Changes that clear the value are not suppressed.
func suppressDefaultTokenWhenWriteOnly(k, old, new string, d *schema.ResourceData) bool {
	if new == "" {
		return false
	}
	pat, diags := d.GetRawConfigAt(cty.GetAttrPath("personal_access_token"))
	if diags.HasError() || !pat.IsNull() {
		return false
	}
	woValue, diags := d.GetRawConfigAt(cty.GetAttrPath("personal_access_token_wo"))
	return !diags.HasError() && !woValue.IsNull()
}

type GitCredentialSchemaStruct struct {
	workspace.CreateCredentialsRequest
	common.Namespace
}

func ResourceGitCredential() common.Resource {
	s := common.StructToSchema(GitCredentialSchemaStruct{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["force"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
		s["personal_access_token"].Sensitive = true
		s["personal_access_token"].DiffSuppressFunc = suppressDefaultTokenWhenWriteOnly
		s["personal_access_token"].DefaultFunc = schema.MultiEnvDefaultFunc([]string{
			"GITHUB_TOKEN",               // https://registry.terraform.io/providers/integrations/github/latest/docs
			"GITLAB_TOKEN",               // https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs
			"AZDO_PERSONAL_ACCESS_TOKEN", // https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs
		}, nil)
		// ConflictsWith is set only on the write-only attribute: on `personal_access_token` it would
		// also trigger for a token sourced from the environment via DefaultFunc.
		s["personal_access_token_wo"] = &schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			WriteOnly:     true,
			Sensitive:     true,
			ConflictsWith: []string{"personal_access_token"},
			RequiredWith:  []string{"personal_access_token_wo_version"},
		}
		s["personal_access_token_wo_version"] = &schema.Schema{
			Type:         schema.TypeInt,
			Optional:     true,
			RequiredWith: []string{"personal_access_token_wo"},
		}
		common.NamespaceCustomizeSchemaMap(s)
		return s
	})

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			var req workspace.CreateCredentialsRequest
			common.DataToStructPointer(d, s, &req)
			req.PersonalAccessToken = getPersonalAccessToken(d)
			resp, err := w.GitCredentials.Create(ctx, req)

			if err != nil {
				if !d.Get("force").(bool) || !isOnlyOneGitCredentialForProviderError(err) {
					return err
				}
				principalId := int64(d.Get("principal_id").(int))
				creds, err := w.GitCredentials.ListAll(ctx, workspace.ListCredentialsRequest{
					PrincipalId: principalId,
				})
				if err != nil {
					return err
				}
				if len(creds) != 1 {
					return fmt.Errorf("list of credentials is either empty or have more than one entry (%d)", len(creds))
				}
				var req workspace.UpdateCredentialsRequest
				common.DataToStructPointer(d, s, &req)
				req.PersonalAccessToken = getPersonalAccessToken(d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			principalId := int64(d.Get("principal_id").(int))
			resp, err := w.GitCredentials.Get(ctx, workspace.GetCredentialsRequest{
				CredentialId: cred_id,
				PrincipalId:  principalId,
			})
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
			req.PersonalAccessToken = getPersonalAccessToken(d)
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			req.CredentialId = cred_id
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			cred_id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			principalId := int64(d.Get("principal_id").(int))
			return w.GitCredentials.Delete(ctx, workspace.DeleteCredentialsRequest{
				CredentialId: cred_id,
				PrincipalId:  principalId,
			})
		},
	}
}
