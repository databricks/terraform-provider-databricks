package sqlanalytics

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ConfPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GlobalConfig used to generate Terraform resource schema and bind to resource data
type GlobalConfig struct {
	SecurityPolicy          string            `json:"security_policy,omitempty" tf:"default:DATA_ACCESS_CONTROL"`
	DataAccessConfig        map[string]string `json:"data_access_config,omitempty"`
	InstanceProfileARN      string            `json:"instance_profile_arn,omitempty"`
	EnableServerlessCompute bool              `json:"enable_serverless_compute,omitempty" tf:"default:false"`
}

// GlobalConfigForRead used to talk to REST API
type GlobalConfigForRead struct {
	SecurityPolicy          string     `json:"security_policy"`
	DataAccessConfig        []ConfPair `json:"data_access_config"`
	InstanceProfileARN      string     `json:"instance_profile_arn,omitempty"`
	EnableServerlessCompute bool       `json:"enable_serverless_compute,omitempty"`
}

func NewSqlGlobalConfigAPI(ctx context.Context, m interface{}) globalConfigAPI {
	return globalConfigAPI{m.(*common.DatabricksClient), ctx}
}

type globalConfigAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a globalConfigAPI) Set(gc GlobalConfig) error {
	data := map[string]interface{}{
		"security_policy":           gc.SecurityPolicy,
		"enable_serverless_compute": gc.EnableServerlessCompute,
	}
	if a.client.Host == "" {
		err := a.client.Authenticate(a.context)
		if err != nil {
			return err
		}
	}
	if a.client.IsAws() {
		data["instance_profile_arn"] = gc.InstanceProfileARN
	} else if gc.InstanceProfileARN != "" {
		return fmt.Errorf("can't use instance_profile_arn outside of AWS")
	}
	cfg := make([]ConfPair, 0, len(gc.DataAccessConfig))
	for k, v := range gc.DataAccessConfig {
		cfg = append(cfg, ConfPair{Key: k, Value: v})
	}
	data["data_access_config"] = cfg

	return a.client.Put(a.context, "/sql/config/endpoints", data)
}

func (a globalConfigAPI) Get() (GlobalConfig, error) {
	gc := GlobalConfig{}
	gcr := GlobalConfigForRead{}
	if err := a.client.Get(a.context, "/sql/config/endpoints", nil, &gcr); err != nil {
		return gc, err
	}
	gc.InstanceProfileARN = gcr.InstanceProfileARN
	gc.SecurityPolicy = gcr.SecurityPolicy
	gc.EnableServerlessCompute = gcr.EnableServerlessCompute
	gc.DataAccessConfig = make(map[string]string, len(gcr.DataAccessConfig))
	for _, v := range gcr.DataAccessConfig {
		gc.DataAccessConfig[v.Key] = v.Value
	}

	return gc, nil
}

func ResourceSQLGlobalConfig() *schema.Resource {
	s := common.StructToSchema(GlobalConfig{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_arn"].Default = ""
		return m
	})
	setGlobalConfig := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		var gc GlobalConfig
		if err := common.DataToStructPointer(d, s, &gc); err != nil {
			return err
		}
		if err := NewSqlGlobalConfigAPI(ctx, c).Set(gc); err != nil {
			return err
		}
		d.SetId("global")
		return nil
	}
	return common.Resource{
		Create: setGlobalConfig,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			gc, err := NewSqlGlobalConfigAPI(ctx, c).Get()
			if err != nil {
				return err
			}
			err = common.StructToData(gc, s, d)
			return err
		},
		Update: setGlobalConfig,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSqlGlobalConfigAPI(ctx, c).Set(GlobalConfig{SecurityPolicy: "DATA_ACCESS_CONTROL"})
		},
		Schema: s,
	}.ToResource()
}
