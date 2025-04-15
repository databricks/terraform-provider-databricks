package tokens

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/oauth2"
)

type ServicePrincipalSecret struct {
	oauth2.CreateServicePrincipalSecretResponse
	ServicePrincipalId string `json:"service_principal_id" tf:"force_new"`
	Lifetime           string `json:"lifetime,omitempty" tf:"computed,force_new"`
}

func createFailedToConvertServicePrincipalIdToNumericError(err error) error {
	return fmt.Errorf("failed to convert service principal ID to numeric: %w", err)
}

func ResourceServicePrincipalSecret() common.Resource {
	spnSecretSchema := common.StructToSchema(ServicePrincipalSecret{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
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
			ac, err := c.AccountClient()
			if err != nil {
				return err
			}
			spId := d.Get("service_principal_id").(string)
			spIdNumeric, err := strconv.ParseInt(spId, 10, 64)
			if err != nil {
				return createFailedToConvertServicePrincipalIdToNumericError(err)
			}
			lifetime := d.Get("lifetime").(string)
			res, err := ac.ServicePrincipalSecrets.Create(ctx, oauth2.CreateServicePrincipalSecretRequest{
				ServicePrincipalId: spIdNumeric,
				Lifetime:           lifetime,
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
			ac, err := c.AccountClient()
			if err != nil {
				return err
			}
			spId := d.Get("service_principal_id").(string)
			spIdNumeric, err := strconv.ParseInt(spId, 10, 64)
			if err != nil {
				return createFailedToConvertServicePrincipalIdToNumericError(err)
			}
			secrets, err := ac.ServicePrincipalSecrets.ListByServicePrincipalId(ctx, spIdNumeric)
			if err != nil {
				return err
			}
			for _, v := range secrets.Secrets {
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
			ac, err := c.AccountClient()
			if err != nil {
				return err
			}
			spId := d.Get("service_principal_id").(string)
			spIdNumeric, err := strconv.ParseInt(spId, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to convert service principal ID to numeric: %w", err)
			}
			err = ac.ServicePrincipalSecrets.Delete(ctx, oauth2.DeleteServicePrincipalSecretRequest{
				SecretId:           d.Id(),
				ServicePrincipalId: spIdNumeric,
			})
			return common.IgnoreNotFoundError(err)
		},
	}
}
