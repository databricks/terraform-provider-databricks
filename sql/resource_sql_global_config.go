package sql

import (
	"context"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	GlobalSqlConfigResourceID = "global"
)

type confPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type repeatedEndpointConfPairs struct {
	ConfigPairs []confPair `json:"configuration_pairs,omitempty"`
}

// GlobalConfig used to generate Terraform resource schema and bind to resource data
type GlobalConfig struct {
	SecurityPolicy          string            `json:"security_policy,omitempty" tf:"default:DATA_ACCESS_CONTROL"`
	DataAccessConfig        map[string]string `json:"data_access_config,omitempty"`
	InstanceProfileARN      string            `json:"instance_profile_arn,omitempty"`
	GoogleServiceAccount    string            `json:"google_service_account,omitempty"`
	EnableServerlessCompute bool              `json:"enable_serverless_compute,omitempty" tf:"computed"`
	SqlConfigParams         map[string]string `json:"sql_config_params,omitempty"`
}

// GlobalConfigForRead used to talk to REST API
type GlobalConfigForRead struct {
	SecurityPolicy             string                     `json:"security_policy"`
	DataAccessConfig           []confPair                 `json:"data_access_config"`
	InstanceProfileARN         string                     `json:"instance_profile_arn,omitempty"`
	GoogleServiceAccount       string                     `json:"google_service_account,omitempty"`
	EnableServerlessCompute    bool                       `json:"enable_serverless_compute,omitempty"`
	SqlConfigurationParameters *repeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`
	ForceSendFields            []string                   `json:"-"`
}

func (g GlobalConfigForRead) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(g)
}

func (g *GlobalConfigForRead) UnmarshalJSON(bs []byte) error {
	return marshal.Unmarshal(bs, g)
}

func NewSqlGlobalConfigAPI(ctx context.Context, m any) globalConfigAPI {
	return globalConfigAPI{m.(*common.DatabricksClient), ctx}
}

type globalConfigAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a globalConfigAPI) Set(gc GlobalConfig, d *schema.ResourceData) error {
	data := GlobalConfigForRead{
		SecurityPolicy:          gc.SecurityPolicy,
		EnableServerlessCompute: gc.EnableServerlessCompute,
	}
	if a.client.Config.Host == "" {
		err := a.client.Config.EnsureResolved()
		if err != nil {
			return err
		}
	}
	if gc.InstanceProfileARN != "" {
		data.InstanceProfileARN = gc.InstanceProfileARN
	}
	if gc.GoogleServiceAccount != "" {
		data.GoogleServiceAccount = gc.GoogleServiceAccount
	}
	cfg := make([]confPair, 0, len(gc.DataAccessConfig))
	for k, v := range gc.DataAccessConfig {
		cfg = append(cfg, confPair{Key: k, Value: v})
	}
	data.DataAccessConfig = cfg
	if len(gc.SqlConfigParams) > 0 {
		sql_params := repeatedEndpointConfPairs{}
		sql_params.ConfigPairs = make([]confPair, 0, len(gc.SqlConfigParams))
		for k, v := range gc.SqlConfigParams {
			sql_params.ConfigPairs = append(sql_params.ConfigPairs, confPair{Key: k, Value: v})
		}
		data.SqlConfigurationParameters = &sql_params
	}
	common.SetForceSendFields(&data, d, []string{"enable_serverless_compute"})

	return a.client.Put(a.context, "/sql/config/warehouses", data)
}

func (a globalConfigAPI) Get() (GlobalConfig, error) {
	gc := GlobalConfig{}
	gcr := GlobalConfigForRead{}
	if err := a.client.Get(a.context, "/sql/config/warehouses", nil, &gcr); err != nil {
		return gc, err
	}
	gc.InstanceProfileARN = gcr.InstanceProfileARN
	gc.GoogleServiceAccount = gcr.GoogleServiceAccount
	gc.SecurityPolicy = gcr.SecurityPolicy
	gc.EnableServerlessCompute = gcr.EnableServerlessCompute
	gc.DataAccessConfig = make(map[string]string, len(gcr.DataAccessConfig))
	for _, v := range gcr.DataAccessConfig {
		gc.DataAccessConfig[v.Key] = v.Value
	}

	return gc, nil
}

func ResourceSqlGlobalConfig() common.Resource {
	s := common.StructToSchema(GlobalConfig{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["enable_serverless_compute"].Deprecated = "This field is intended as an internal API " +
			"and may be removed from the Databricks Terraform provider in the future"
		m["instance_profile_arn"].ConflictsWith = []string{"google_service_account"}
		m["google_service_account"].ConflictsWith = []string{"instance_profile_arn"}
		return m
	})
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	setGlobalConfig := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		var gc GlobalConfig
		common.DataToStructPointer(d, s, &gc)
		if err := NewSqlGlobalConfigAPI(ctx, c).Set(gc, d); err != nil {
			return err
		}
		d.SetId(GlobalSqlConfigResourceID)
		return nil
	}
	return common.Resource{
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			// enable_serverless_compute is an optional boolean parameter which may be specified as `false`.
			if _, ok := d.GetOkExists("enable_serverless_compute"); !ok {
				// Read the current global config and use the current value of enable_serverless_compute as
				// the default value if not specified.
				gc, err := NewSqlGlobalConfigAPI(ctx, newClient).Get()
				if err != nil {
					return err
				}
				d.Set("enable_serverless_compute", gc.EnableServerlessCompute)
			}
			return setGlobalConfig(ctx, d, newClient)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			gc, err := NewSqlGlobalConfigAPI(ctx, newClient).Get()
			if err != nil {
				return err
			}
			err = common.StructToData(gc, s, d)
			return err
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return setGlobalConfig(ctx, d, newClient)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return NewSqlGlobalConfigAPI(ctx, newClient).Set(GlobalConfig{
				SecurityPolicy:          "DATA_ACCESS_CONTROL",
				EnableServerlessCompute: d.Get("enable_serverless_compute").(bool),
			}, d)
		},
		Schema: s,
	}
}
