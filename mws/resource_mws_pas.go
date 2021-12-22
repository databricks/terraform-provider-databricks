package mws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewPrivateAccessSettingsAPI creates VPCEndpointAPI instance from provider meta
func NewPrivateAccessSettingsAPI(ctx context.Context, m interface{}) PrivateAccessSettingsAPI {
	return PrivateAccessSettingsAPI{m.(*common.DatabricksClient), ctx}
}

// PrivateAccessSettingsAPI exposes the PAS API
type PrivateAccessSettingsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the PAS ceation process
func (a PrivateAccessSettingsAPI) Create(pas *PrivateAccessSettings) error {
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings", pas.AccountID)
	return a.client.Post(a.context, pasAPIPath, pas, &pas)
}

// Read returns the PAS object along with metadata and any additional errors
func (a PrivateAccessSettingsAPI) Read(mwsAcctID, pasID string) (PrivateAccessSettings, error) {
	var pas PrivateAccessSettings
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings/%s", mwsAcctID, pasID)
	err := a.client.Get(a.context, pasAPIPath, nil, &pas)
	return pas, err
}

func (a PrivateAccessSettingsAPI) Update(pas *PrivateAccessSettings) error {
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings/%s", pas.AccountID, pas.PasID)
	return a.client.Put(a.context, pasAPIPath, pas)
}

// Delete deletes the PAS object given a pas id
func (a PrivateAccessSettingsAPI) Delete(mwsAcctID, pasID string) error {
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings/%s", mwsAcctID, pasID)
	if err := a.client.Delete(a.context, pasAPIPath, nil); err != nil {
		return err
	}
	return nil
}

// List lists all the available PAS objects in the mws account
func (a PrivateAccessSettingsAPI) List(mwsAcctID string) ([]PrivateAccessSettings, error) {
	var pasList []PrivateAccessSettings
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings", mwsAcctID)
	err := a.client.Get(a.context, pasAPIPath, nil, &pasList)
	return pasList, err
}

// ResourcePrivateAccessSettings ...
func ResourcePrivateAccessSettings() *schema.Resource {
	s := common.StructToSchema(PrivateAccessSettings{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["private_access_settings_name"].ValidateFunc = validation.StringLenBetween(4, 256)
		return s
	})
	p := common.NewPairSeparatedID("account_id", "private_access_settings_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var pas PrivateAccessSettings
			common.DataToStructPointer(d, s, &pas)
			if err := NewPrivateAccessSettingsAPI(ctx, c).Create(&pas); err != nil {
				return err
			}
			d.Set("private_access_settings_id", pas.PasID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, pasID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			pas, err := NewPrivateAccessSettingsAPI(ctx, c).Read(accountID, pasID)
			if err != nil {
				return err
			}
			return common.StructToData(pas, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, pasID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			var pas PrivateAccessSettings
			common.DataToStructPointer(d, s, &pas)
			pas.PasID = pasID
			return NewPrivateAccessSettingsAPI(ctx, c).Update(&pas)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, pasID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewPrivateAccessSettingsAPI(ctx, c).Delete(accountID, pasID)
		},
	}.ToResource()
}
