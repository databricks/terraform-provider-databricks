package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// LogDelivery wrapper
type LogDelivery struct {
	LogDeliveryConfiguration LogDeliveryConfiguration `json:"log_delivery_configuration"`
}

// LogDeliveryConfiguration describes log delivery
type LogDeliveryConfiguration struct {
	common.Namespace
	AccountID              string  `json:"account_id" tf:"force_new"`
	ConfigID               string  `json:"config_id,omitempty" tf:"computed,force_new"`
	CredentialsID          string  `json:"credentials_id" tf:"force_new"`
	StorageConfigurationID string  `json:"storage_configuration_id" tf:"force_new"`
	WorkspaceIdsFilter     []int64 `json:"workspace_ids_filter,omitempty" tf:"force_new"`
	ConfigName             string  `json:"config_name,omitempty" tf:"force_new"`
	Status                 string  `json:"status,omitempty" tf:"computed"`
	LogType                string  `json:"log_type" tf:"force_new"`
	OutputFormat           string  `json:"output_format" tf:"force_new"`
	DeliveryPathPrefix     string  `json:"delivery_path_prefix,omitempty" tf:"force_new"`
	DeliveryStartTime      string  `json:"delivery_start_time,omitempty" tf:"computed,force_new"`
}

// LogDeliveryAPI ...
type LogDeliveryAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewLogDeliveryAPI ...
func NewLogDeliveryAPI(ctx context.Context, m any) LogDeliveryAPI {
	return LogDeliveryAPI{m.(*common.DatabricksClient), ctx}
}

// Read reads log delivery configuration
func (a LogDeliveryAPI) Read(accountID, configID string) (LogDeliveryConfiguration, error) {
	var ld LogDelivery
	err := a.client.Get(a.context, fmt.Sprintf("/accounts/%s/log-delivery/%s", accountID, configID), nil, &ld)
	return ld.LogDeliveryConfiguration, err
}

// Create new log delivery configuration
func (a LogDeliveryAPI) Create(ldc LogDeliveryConfiguration) (string, error) {
	var ld LogDelivery
	err := a.client.Post(a.context, fmt.Sprintf("/accounts/%s/log-delivery", ldc.AccountID), LogDelivery{
		LogDeliveryConfiguration: ldc,
	}, &ld)
	// todo: verify with empty response - structs should have empty default strings
	return ld.LogDeliveryConfiguration.ConfigID, err
}

// patch log delivery configuration - i.e. can only enable or disable it
func (a LogDeliveryAPI) Patch(accountID, configID string, status string) error {
	return a.client.Patch(a.context, fmt.Sprintf("/accounts/%s/log-delivery/%s", accountID, configID), map[string]string{
		"status": status,
	})
}

func ResourceMwsLogDelivery() common.Resource {
	p := common.NewPairID("account_id", "config_id")
	s := common.StructToSchema(LogDeliveryConfiguration{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			// nolint
			s["config_name"].ValidateFunc = validation.StringLenBetween(0, 255)
			s["delivery_start_time"].DiffSuppressFunc = func(
				k, old, new string, d *schema.ResourceData) bool {
				return false
			}
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ldc LogDeliveryConfiguration
			common.DataToStructPointer(d, s, &ldc)
			configID, err := NewLogDeliveryAPI(ctx, c).Create(ldc)
			if err != nil {
				return err
			}
			if err = d.Set("config_id", configID); err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ldc LogDeliveryConfiguration
			common.DataToStructPointer(d, s, &ldc)
			err := NewLogDeliveryAPI(ctx, c).Patch(ldc.AccountID, ldc.ConfigID, ldc.Status)
			if err != nil {
				return err
			}
			return common.StructToData(ldc, s, d)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, configID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			ldc, err := NewLogDeliveryAPI(ctx, c).Read(accountID, configID)
			if err != nil {
				return err
			}
			return common.StructToData(ldc, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, configID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewLogDeliveryAPI(ctx, c).Patch(accountID, configID, "DISABLED")
		},
	}
}
