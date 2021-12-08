package catalog

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataAccessConfigurationsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewDataAccessConfigurationsAPI(ctx context.Context, m interface{}) DataAccessConfigurationsAPI {
	return DataAccessConfigurationsAPI{m.(*common.DatabricksClient), ctx}
}

type AwsIamRole struct {
	RoleARN string `json:"role_arn"`
}

type AzureServicePrincipal struct {
	DirectoryID   string `json:"directory_id"`
	ApplicationID string `json:"application_id"`
	ClientSecret  string `json:"client_secret"`
}

type DataAccessConfiguration struct {
	ID                string                 `json:"id,omitempty" tf:"computed"`
	Name              string                 `json:"name"`
	ConfigurationType string                 `json:"configuration_type,omitempty" tf:"computed"`
	Aws               *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure             *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
}

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

func ResourceDataAccessConfiguration() *schema.Resource {
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
			alof := []string{"aws_iam_role", "azure_service_principal"}
			m["aws_iam_role"].AtLeastOneOf = alof
			m["azure_service_principal"].AtLeastOneOf = alof
			return m
		})
	p := common.NewPairID("metastore_id", "id")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var dac DataAccessConfiguration
			if err := common.DataToStructPointer(d, s, &dac); err != nil {
				return err
			}
			metastoreID := d.Get("metastore_id").(string)
			if err := NewDataAccessConfigurationsAPI(ctx, c).Create(metastoreID, &dac); err != nil {
				return err
			}
			d.Set("id", dac.ID)
			p.Pack(d)
			if d.Get("is_default").(bool) {
				return NewMetastoresAPI(ctx, c).updateMetastore(metastoreID, map[string]interface{}{
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
