package tokens

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
)

type ServicePrincipalSecret struct {
	oauth2.CreateServicePrincipalSecretResponse
	common.Namespace
	ServicePrincipalId string `json:"service_principal_id" tf:"force_new"`
	Lifetime           string `json:"lifetime,omitempty" tf:"computed,force_new"`
}

func createFailedToConvertServicePrincipalIdToNumericError(err error) error {
	return fmt.Errorf("failed to convert service principal ID to numeric: %w", err)
}

func ResourceServicePrincipalSecret() common.Resource {
	spnSecretSchema := common.StructToSchema(ServicePrincipalSecret{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["time_rotating"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			}
			m["id"].Computed = true
			m["create_time"].Computed = true
			m["expire_time"].Computed = true
			m["update_time"].Computed = true
			m["secret_hash"].Computed = true
			m["secret"].Computed = true
			m["secret"].Sensitive = true
			m["status"].Computed = true
			return m
		})
	return common.Resource{
		Schema: spnSecretSchema,
		CanSkipReadAfterCreateAndUpdate: func(d *schema.ResourceData) bool {
			return true
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			spId := d.Get("service_principal_id").(string)
			spIdNumeric, err := strconv.ParseInt(spId, 10, 64)
			if err != nil {
				return createFailedToConvertServicePrincipalIdToNumericError(err)
			}
			lifetime := d.Get("lifetime").(string)
			var res *oauth2.CreateServicePrincipalSecretResponse

			err = c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				var err error
				res, err = acc.ServicePrincipalSecrets.Create(ctx, oauth2.CreateServicePrincipalSecretRequest{
					ServicePrincipalId: strconv.FormatInt(spIdNumeric, 10),
					Lifetime:           lifetime,
				})
				return err
			}, func(w *databricks.WorkspaceClient) error {
				var err error
				res, err = w.ServicePrincipalSecretsProxy.Create(ctx, oauth2.CreateServicePrincipalSecretRequest{
					ServicePrincipalId: strconv.FormatInt(spIdNumeric, 10),
					Lifetime:           lifetime,
				})
				return err
			})
			if err != nil {
				return err
			}
			err = common.StructToData(*res, spnSecretSchema, d)
			if err != nil {
				return err
			}
			d.Set("lifetime", lifetime)
			d.Set("service_principal_id", spId)
			d.SetId(res.Id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			spId := d.Get("service_principal_id").(string)
			var secrets []oauth2.SecretInfo
			err := c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				var err error
				secrets, err = acc.ServicePrincipalSecrets.ListAll(ctx, oauth2.ListServicePrincipalSecretsRequest{
					ServicePrincipalId: spId,
				})
				return err
			}, func(w *databricks.WorkspaceClient) error {
				var err error
				secrets, err = w.ServicePrincipalSecretsProxy.ListAll(ctx, oauth2.ListServicePrincipalSecretsRequest{
					ServicePrincipalId: spId,
				})
				return err
			})
			if err != nil {
				return err
			}
			for _, v := range secrets {
				if v.Id != d.Id() {
					continue
				}
				// check if the token is expired, although in practice API just doesn't return expired tokens
				if v.ExpireTime != "" {
					expireTime, err := time.Parse(time.RFC3339, v.ExpireTime)
					if err != nil {
						return fmt.Errorf("failed to parse expire time: %w", err)
					}
					if time.Now().After(expireTime) {
						log.Printf("[INFO] service principal secret with id %s is expired, recreating it", d.Id())
						d.SetId("")
						return nil
					}
				}
				// copy fields that aren't part of the result
				secret := d.Get("secret").(string)
				lifetime := d.Get("lifetime").(string)
				common.StructToData(v, spnSecretSchema, d)
				d.Set("secret", secret)
				d.Set("lifetime", lifetime)
				d.Set("service_principal_id", spId)
				return nil
			}
			// recreate it if not found
			log.Printf("[INFO] service principal secret with id %s not found, recreating it", d.Id())
			d.SetId("")
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			spId := d.Get("service_principal_id").(string)
			spIdNumeric, err := strconv.ParseInt(spId, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to convert service principal ID to numeric: %w", err)
			}
			err = c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				return acc.ServicePrincipalSecrets.Delete(ctx, oauth2.DeleteServicePrincipalSecretRequest{
					SecretId:           d.Id(),
					ServicePrincipalId: strconv.FormatInt(spIdNumeric, 10),
				})
			}, func(w *databricks.WorkspaceClient) error {
				return w.ServicePrincipalSecretsProxy.Delete(ctx, oauth2.DeleteServicePrincipalSecretRequest{
					SecretId:           d.Id(),
					ServicePrincipalId: strconv.FormatInt(spIdNumeric, 10),
				})
			})
			return common.IgnoreNotFoundError(err)
		},
	}
}
