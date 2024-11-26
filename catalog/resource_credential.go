package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var credentialSchema = common.StructToSchema(catalog.CredentialInfo{},
	func(m map[string]*schema.Schema) map[string]*schema.Schema {
		var alofServiceCreds = []string{"aws_iam_role", "azure_managed_identity", "azure_service_principal"}
		for _, cred := range alofServiceCreds {
			common.CustomizeSchemaPath(m, cred).SetExactlyOneOf(alofServiceCreds)
		}

		for _, required := range []string{"name", "purpose"} {
			common.CustomizeSchemaPath(m, required).SetRequired()
		}

		for _, computed := range []string{"id", "created_at", "created_by", "full_name", "isolation_mode",
			"metastore_id", "owner", "updated_at", "updated_by", "used_for_managed_storage"} {
			common.CustomizeSchemaPath(m, computed).SetComputed()
		}

		common.MustSchemaPath(m, "aws_iam_role", "external_id").Computed = true
		common.MustSchemaPath(m, "aws_iam_role", "unity_catalog_iam_arn").Computed = true
		common.MustSchemaPath(m, "azure_managed_identity", "credential_id").Computed = true
		common.MustSchemaPath(m, "azure_service_principal", "client_secret").Sensitive = true

		m["force_destroy"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
		m["force_update"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
		m["skip_validation"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return old == "false" && new == "true"
			},
		}
		m["credential_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
		return m
	})

func ResourceCredential() common.Resource {
	return common.Resource{
		Schema: credentialSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var create catalog.CreateCredentialRequest
			common.DataToStructPointer(d, credentialSchema, &create)
			cred, err := w.Credentials.CreateCredential(ctx, create)
			if err != nil {
				return err
			}
			d.SetId(cred.Name)

			// Update owner or isolation mode if it is provided
			if !updateRequired(d, []string{"owner", "isolation_mode"}) {
				return nil
			}

			var update catalog.UpdateCredentialRequest
			common.DataToStructPointer(d, credentialSchema, &update)
			update.NameArg = d.Id()
			_, err = w.Credentials.UpdateCredential(ctx, update)
			if err != nil {
				return err
			}

			// Bind the current workspace if the credential is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, cred.Name, catalog.UpdateBindingsSecurableTypeServiceCredential)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			cred, err := w.Credentials.GetCredentialByNameArg(ctx, d.Id())
			if err != nil {
				return err
			}
			d.Set("credential_id", cred.Id)
			return common.StructToData(cred, credentialSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_update").(bool)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateCredRequest catalog.UpdateCredentialRequest
			common.DataToStructPointer(d, credentialSchema, &updateCredRequest)
			updateCredRequest.NameArg = d.Id()
			updateCredRequest.Force = force

			if d.HasChange("owner") {
				_, err = w.Credentials.UpdateCredential(ctx, catalog.UpdateCredentialRequest{
					NameArg: updateCredRequest.NameArg,
					Owner:   updateCredRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}
			if d.HasChange("read_only") {
				updateCredRequest.ForceSendFields = append(updateCredRequest.ForceSendFields, "ReadOnly")
			}

			updateCredRequest.Owner = ""
			_, err = w.Credentials.UpdateCredential(ctx, updateCredRequest)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Credentials.UpdateCredential(ctx, catalog.UpdateCredentialRequest{
						NameArg: updateCredRequest.NameArg,
						Owner:   old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			// Bind the current workspace if the credential is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, updateCredRequest.NameArg, catalog.UpdateBindingsSecurableTypeServiceCredential)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Credentials.DeleteCredential(ctx, catalog.DeleteCredentialRequest{
				NameArg: d.Id(),
				Force:   force,
			})
		},
	}
}
