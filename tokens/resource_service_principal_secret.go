package tokens

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ServicePrincipalSecret struct {
	ID     string `json:"id,omitempty"`
	Secret string `json:"secret,omitempty" tf:"computed,sensitive"`
	Status string `json:"status,omitempty" tf:"computed"`
}

type ListServicePrincipalSecrets struct {
	Secrets []ServicePrincipalSecret `json:"secrets"`
}

// NewServicePrincipalSecretAPI creates ServicePrincipalSecretAPI instance from provider meta
func NewServicePrincipalSecretAPI(ctx context.Context, m any) ServicePrincipalSecretAPI {
	return ServicePrincipalSecretAPI{m.(*common.DatabricksClient), ctx}
}

// ServicePrincipalSecretAPI exposes the API to create client secrets
type ServicePrincipalSecretAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a ServicePrincipalSecretAPI) createServicePrincipalSecret(spnID string) (secret *ServicePrincipalSecret, err error) {
	path := fmt.Sprintf("/accounts/%s/servicePrincipals/%s/credentials/secrets", a.client.Config.AccountID, spnID)
	err = a.client.Post(a.context, path, map[string]any{}, &secret)
	return
}

func (a ServicePrincipalSecretAPI) listServicePrincipalSecrets(spnID string) (secrets ListServicePrincipalSecrets, err error) {
	path := fmt.Sprintf("/accounts/%s/servicePrincipals/%s/credentials/secrets", a.client.Config.AccountID, spnID)
	err = a.client.Get(a.context, path, nil, &secrets)
	return
}

func (a ServicePrincipalSecretAPI) deleteServicePrincipalSecret(spnID, secretID string) error { // FIXME
	path := fmt.Sprintf("/accounts/%s/servicePrincipals/%s/credentials/secrets/%s", a.client.Config.AccountID, spnID, secretID)
	return a.client.Delete(a.context, path, nil)
}

func ResourceServicePrincipalSecret() common.Resource {
	spnSecretSchema := common.StructToSchema(ServicePrincipalSecret{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["id"].Computed = true
			m["service_principal_id"] = &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			}
			return m
		})
	return common.Resource{
		Schema: spnSecretSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.AccountID == "" {
				return errors.New("must have `account_id` on provider")
			}
			idSeen := map[string]bool{}
			api := NewServicePrincipalSecretAPI(ctx, c)
			spnID := d.Get("service_principal_id").(string)
			secrets, err := api.listServicePrincipalSecrets(spnID)
			if err != nil {
				return err
			}
			for _, v := range secrets.Secrets {
				idSeen[v.ID] = true
			}
			secret, err := api.createServicePrincipalSecret(spnID)
			if err != nil {
				return err
			}
			secrets, err = api.listServicePrincipalSecrets(spnID)
			if err != nil {
				return err
			}
			// ugly hack because rpc does not return ID of created secret
			for _, v := range secrets.Secrets {
				if len(idSeen) > 0 && idSeen[v.ID] {
					continue
				}
				d.SetId(v.ID)
			}
			return d.Set("secret", secret.Secret)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.AccountID == "" {
				return errors.New("must have `account_id` on provider")
			}
			api := NewServicePrincipalSecretAPI(ctx, c)
			spnID := d.Get("service_principal_id").(string)
			secrets, err := api.listServicePrincipalSecrets(spnID)
			if err != nil {
				return err
			}
			for _, v := range secrets.Secrets {
				if v.ID != d.Id() {
					continue
				}
				return d.Set("status", v.Status)
			}
			return &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "client secret not found",
			}
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.AccountID == "" {
				return errors.New("must have `account_id` on provider")
			}
			api := NewServicePrincipalSecretAPI(ctx, c)
			spnID := d.Get("service_principal_id").(string)
			return api.deleteServicePrincipalSecret(spnID, d.Id())
		},
	}
}
