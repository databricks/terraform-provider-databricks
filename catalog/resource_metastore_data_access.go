package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataAccessConfigurationsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewDataAccessConfigurationsAPI(ctx context.Context, m any) DataAccessConfigurationsAPI {
	return DataAccessConfigurationsAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

type AwsIamRole struct {
	RoleARN string `json:"role_arn"`
}

type AzureServicePrincipal struct {
	DirectoryID   string `json:"directory_id"`
	ApplicationID string `json:"application_id"`
	ClientSecret  string `json:"client_secret" tf:"sensitive"`
}

type AzureManagedIdentity struct {
	AccessConnectorID string `json:"access_connector_id"`
}

type GcpServiceAccountKey struct {
	Email        string `json:"email"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key" tf:"sensitive"`
}

type DbGcpServiceAccount struct {
	Email string `json:"email,omitempty" tf:"computed"`
}

type DataAccessConfiguration struct {
	ID                string                 `json:"id,omitempty" tf:"computed"`
	Name              string                 `json:"name"`
	ConfigurationType string                 `json:"configuration_type,omitempty" tf:"computed"`
	Aws               *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure             *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI              *AzureManagedIdentity  `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey          *GcpServiceAccountKey  `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DBGcpSA           *DbGcpServiceAccount   `json:"databricks_gcp_service_account,omitempty" tf:"group:access"`
}

var alofCred = []string{"aws_iam_role", "azure_service_principal", "azure_managed_identity", "gcp_service_account_key", "databricks_gcp_service_account"}

func (a DataAccessConfigurationsAPI) Create(metastoreID string, dac *DataAccessConfiguration) error {
	path := fmt.Sprintf("/unity-catalog/metastores/%s/data-access-configurations", metastoreID)
	return a.client.Post(a.context, path, dac, dac)
}

func (a DataAccessConfigurationsAPI) Get(metastoreID, dacID string) (dac DataAccessConfiguration, err error) {
	path := fmt.Sprintf("/unity-catalog/metastores/%s/data-access-configurations/%s", metastoreID, dacID)
	err = a.client.Get(a.context, path, nil, &dac)
	return
}

func (a DataAccessConfigurationsAPI) Delete(metastoreID, dacID string) error {
	path := fmt.Sprintf("/unity-catalog/metastores/%s/data-access-configurations/%s", metastoreID, dacID)
	return a.client.Delete(a.context, path, nil)
}

func SuppressGcpSAKeyDiff(k, old, new string, d *schema.ResourceData) bool {
	//ignore changes in private_key
	return !d.HasChanges("gcp_service_account_key.0.email", "gcp_service_account_key.0.private_key_id")
}

func ResourceMetastoreDataAccess() *schema.Resource {
	s := common.StructToSchema(DataAccessConfiguration{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["metastore_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			}
			m["is_default"] = &schema.Schema{
				// having more than one default DAC per metastore will lead
				// to Terraform re-assigning default_data_access_config_id
				// on every apply.
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["aws_iam_role"].AtLeastOneOf = alofCred
			m["azure_service_principal"].AtLeastOneOf = alofCred
			m["azure_managed_identity"].AtLeastOneOf = alofCred
			m["gcp_service_account_key"].AtLeastOneOf = alofCred
			m["databricks_gcp_service_account"].AtLeastOneOf = alofCred

			// suppress changes for private_key
			m["gcp_service_account_key"].DiffSuppressFunc = SuppressGcpSAKeyDiff
			return m
		})
	p := common.NewPairID("metastore_id", "id")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var dac DataAccessConfiguration
			common.DataToStructPointer(d, s, &dac)
			metastoreID := d.Get("metastore_id").(string)
			if err := NewDataAccessConfigurationsAPI(ctx, c).Create(metastoreID, &dac); err != nil {
				return err
			}
			d.Set("id", dac.ID)
			p.Pack(d)
			if d.Get("is_default").(bool) {
				return NewMetastoresAPI(ctx, c).updateMetastore(metastoreID, map[string]any{
					"default_data_access_config_id": dac.ID,
				})
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreID, dacID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			dac, err := NewDataAccessConfigurationsAPI(ctx, c).Get(metastoreID, dacID)
			if err != nil {
				return err
			}
			metastore, err := NewMetastoresAPI(ctx, c).getMetastore(metastoreID)
			if err != nil {
				return err
			}
			isDefault := metastore.DefaultDacID == dacID
			d.Set("is_default", isDefault)
			return common.StructToData(dac, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreID, dacID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewDataAccessConfigurationsAPI(ctx, c).Delete(metastoreID, dacID)
		},
	}.ToResource()
}
