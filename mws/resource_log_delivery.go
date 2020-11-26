package mws

import (
	"context"
	"fmt"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// LogDelivery wrapper
type LogDelivery struct {
	LogDeliveryConfiguration LogDeliveryConfiguration `json:"log_delivery_configuration"`
}

// LogDeliveryConfiguration describes log delivery
type LogDeliveryConfiguration struct {
	AccountID              string   `json:"account_id"`
	ConfigID               string   `json:"config_id,omitempty" tf:"computed"`
	CredentialsID          string   `json:"credentials_id"`
	StorageConfigurationID string   `json:"storage_configuration_id"`
	WorkspaceIdsFilter     []string `json:"workspace_ids_filter,omitempty"`
	ConfigName             string   `json:"config_name,omitempty"`
	Status                 string   `json:"status,omitempty" tf:"computed"`
	LogType                string   `json:"log_type"`
	OutputFormat           string   `json:"output_format"`
	DeliveryPathPrefix     string   `json:"delivery_path_prefix,omitempty"`
	DeliveryStartTime      string   `json:"delivery_start_time,omitempty" tf:"computed"`
}

// LogDeliveryAPI ...
type LogDeliveryAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewLogDeliveryAPI ...
func NewLogDeliveryAPI(m interface{}) LogDeliveryAPI {
	return LogDeliveryAPI{m.(*common.DatabricksClient), context.TODO()}
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

// Disable log delivery configuration - e.g. delete it
func (a LogDeliveryAPI) Disable(accountID, configID string) error {
	return a.client.Patch(a.context, fmt.Sprintf("/accounts/%s/log-delivery/%s", accountID, configID), map[string]string{
		"status": "DISABLED",
	})
}

// ResourceLogDelivery ..
func ResourceLogDelivery() *schema.Resource {
	p := util.NewPairID("account_id", "config_id")
	s := internal.StructToSchema(LogDeliveryConfiguration{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			// nolint
			s["config_name"].ValidateFunc = validation.StringLenBetween(0, 255)
			for k, v := range s {
				if v.Computed {
					continue
				}
				s[k].ForceNew = true
			}
			s["delivery_start_time"].DiffSuppressFunc = func(
				k, old, new string, d *schema.ResourceData) bool {
				return false
			}
			return s
		})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		accountID, configID, err := p.Unpack(d)
		if err != nil {
			return diag.FromErr(err)
		}
		ldc, err := NewLogDeliveryAPI(m).Read(accountID, configID)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[DEBUG] Log delivery configuration %s was not found. Removing from state.", configID)
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		if ldc.Status == "DISABLED" {
			log.Printf("[DEBUG] Log delivery configuration %s was disabled. Removing from state.", configID)
			d.SetId("")
			return nil
		}
		err = internal.StructToData(ldc, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:      s,
		ReadContext: readContext,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ldc LogDeliveryConfiguration
			err := internal.DataToStructPointer(d, s, &ldc)
			if err != nil {
				return diag.FromErr(err)
			}
			configID, err := NewLogDeliveryAPI(m).Create(ldc)
			if err != nil {
				return diag.FromErr(err)
			}
			err = d.Set("config_id", configID)
			if err != nil {
				return diag.FromErr(err)
			}
			p.Pack(d)
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			accountID, configID, err := p.Unpack(d)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewLogDeliveryAPI(m).Disable(accountID, configID)
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
	}
}
