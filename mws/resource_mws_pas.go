package mws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewPASAPI creates VPCEndpointAPI instance from provider meta
func NewPASAPI(ctx context.Context, m interface{}) PASAPI {
	return PASAPI{m.(*common.DatabricksClient), ctx}
}

// PASAPI exposes the PAS API
type PASAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the PAS ceation process
func (a PASAPI) Create(pas *PAS) error {
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings", pas.AccountID)
	return a.client.Post(a.context, pasAPIPath, pas, &pas)
}

// Read returns the PAS object along with metadata and any additional errors
func (a PASAPI) Read(mwsAcctID, pasID string) (PAS, error) {
	var pas PAS
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings/%s", mwsAcctID, pasID)
	err := a.client.Get(a.context, pasAPIPath, nil, &pas)
	return pas, err
}

// Delete deletes the PAS object given a pas id
func (a PASAPI) Delete(mwsAcctID, pasID string) error {
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings/%s", mwsAcctID, pasID)
	if err := a.client.Delete(a.context, pasAPIPath, nil); err != nil {
		return err
	}
	return resource.RetryContext(a.context, 60*time.Second, func() *resource.RetryError {
		pas, err := a.Read(mwsAcctID, pasID)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] private access settings %s/%s is removed.", mwsAcctID, pasID)
			return nil
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		msg := fmt.Errorf("private access settings %s is not removed yet. private access settings Status: %s", pas.PasName, pas.PasStatus)
		log.Printf("[INFO] %s", msg)
		return resource.RetryableError(msg)
	})
}

// List lists all the available PAS objects in the mws account
func (a PASAPI) List(mwsAcctID string) ([]PAS, error) {
	var pasList []PAS
	pasAPIPath := fmt.Sprintf("/accounts/%s/private-access-settings", mwsAcctID)
	err := a.client.Get(a.context, pasAPIPath, nil, &pasList)
	return pasList, err
}

// ResourcePAS ...
func ResourcePAS() *schema.Resource {
	s := internal.StructToSchema(PAS{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["account_id"].MinItems = 1
		s["private_access_settings_name"].ValidateFunc = validation.StringLenBetween(4, 256)
		s["private_access_settings_id"].MinItems = 1
		s["aws_region"].MinItems = 1
		return s
	})
	p := util.NewPairSeparatedID("account_id", "private_access_settings_id", "/")
	return util.CommonResource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var pas PAS
			if err := internal.DataToStructPointer(d, s, &pas); err != nil {
				return err
			}
			if err := NewPASAPI(ctx, c).Create(&pas); err != nil {
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
			pas, err := NewPASAPI(ctx, c).Read(accountID, pasID)
			if err != nil {
				return err
			}
			return internal.StructToData(pas, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, pasID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewPASAPI(ctx, c).Delete(accountID, pasID)
		},
	}.ToResource()
}
