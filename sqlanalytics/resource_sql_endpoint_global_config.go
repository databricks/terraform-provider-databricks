package sqlanalytics

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SQLEndpointConfPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SQLEndpointGlobalConfig ...
type SQLEndpointGlobalConfig struct {
	SecurityPolicy          string            `json:"security_policy,omitempty" tf:"default:DATA_ACCESS_CONTROL"`
	DataAccessConfig        map[string]string `json:"data_access_config,omitempty"`
	InstanceProfileARN      string            `json:"instance_profile_arn,omitempty"`
	EnableServerlessCompute bool              `json:"enable_serverless_compute,omitempty"`
}

// SQLEndpointGlobalConfigForRead ...
type SQLEndpointGlobalConfigForRead struct {
	SecurityPolicy          string                `json:"security_policy"`
	DataAccessConfig        []SQLEndpointConfPair `json:"data_access_config"`
	InstanceProfileARN      string                `json:"instance_profile_arn,omitempty"`
	EnableServerlessCompute bool                  `json:"enable_serverless_compute,omitempty"`
}

// NewSQLEndpointsGlobalConfigAPI ...
func NewSQLEndpointsGlobalConfigAPI(ctx context.Context, m interface{}) SQLEndpointsGlobalConfigAPI {
	return SQLEndpointsGlobalConfigAPI{m.(*common.DatabricksClient), ctx}
}

// SQLEndpointsAPI ...
type SQLEndpointsGlobalConfigAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Set ...
func (a SQLEndpointsGlobalConfigAPI) Set(gc SQLEndpointGlobalConfig) error {
	data := map[string]interface{}{
		"security_policy":           gc.SecurityPolicy,
		"enable_serverless_compute": gc.EnableServerlessCompute,
	}
	if a.client.IsAws() {
		data["instance_profile_arn"] = gc.InstanceProfileARN
	} else if gc.InstanceProfileARN != "" {
		return fmt.Errorf("can't use instance_profile_arn outside of AWS")
	}
	cfg := make([]SQLEndpointConfPair, len(gc.DataAccessConfig))
	for k, v := range gc.DataAccessConfig {
		cfg = append(cfg, SQLEndpointConfPair{Key: k, Value: v})
	}
	data["data_access_config"] = cfg

	//log.Printf("[DEBUG] data=%v, isAWS? %v isAzure? %v isGcp? %v host: %v", data, a.client.IsAws(), a.client.IsAzure(), a.client.IsGcp(), a.client.Host)

	return a.client.Put(a.context, "/sql/config/endpoints", data)
}

func (a SQLEndpointsGlobalConfigAPI) Get() (SQLEndpointGlobalConfig, error) {
	gc := SQLEndpointGlobalConfig{}
	gcr := SQLEndpointGlobalConfigForRead{}
	if err := a.client.Get(a.context, "/sql/config/endpoints", nil, &gcr); err != nil {
		return gc, err
	}
	gc.InstanceProfileARN = gcr.InstanceProfileARN
	gc.SecurityPolicy = gcr.SecurityPolicy
	gc.EnableServerlessCompute = gcr.EnableServerlessCompute
	for _, v := range gcr.DataAccessConfig {
		gc.DataAccessConfig[v.Key] = v.Value
	}

	return gc, nil
}

// ResourceSQLEndpointGlobalConfig ...
func ResourceSQLEndpointGlobalConfig() *schema.Resource {
	s := common.StructToSchema(SQLEndpointGlobalConfig{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		// m[security_policy].ValidateDiagFunc = validation // NONE, DATA_ACCESS_CONTROL, PASSTHROUGH
		return m
	})

	set_func := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		var gc SQLEndpointGlobalConfig
		if err := common.DataToStructPointer(d, s, &gc); err != nil {
			return err
		}
		if err := NewSQLEndpointsGlobalConfigAPI(ctx, c).Set(gc); err != nil {
			return err
		}
		d.SetId("global")
		return nil
	}

	return common.Resource{
		Create: set_func,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			gc, err := NewSQLEndpointsGlobalConfigAPI(ctx, c).Get()
			if err != nil {
				return err
			}
			err = common.StructToData(gc, s, d)
			return err
		},
		Update: set_func,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSQLEndpointsGlobalConfigAPI(ctx, c).Set(SQLEndpointGlobalConfig{SecurityPolicy: "DATA_ACCESS_CONTROL"})
		},
		Schema: s,
	}.ToResource()
}
