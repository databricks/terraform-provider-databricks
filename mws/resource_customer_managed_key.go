package mws

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AwsKeyInfo has information about the KMS key for BYOK
type AwsKeyInfo struct {
	KeyArn    string `json:"key_arn"`
	KeyAlias  string `json:"key_alias"`
	KeyRegion string `json:"key_region,omitempty" tf:"computed"`
}

// CustomerManagedKey contains key information and metadata for BYOK for E2
type CustomerManagedKey struct {
	CustomerManagedKeyID string      `json:"customer_managed_key_id,omitempty" tf:"computed"`
	AwsKeyInfo           *AwsKeyInfo `json:"aws_key_info"`
	AccountID            string      `json:"account_id"`
	CreationTime         int64       `json:"creation_time,omitempty" tf:"computed"`
}

// NewCustomerManagedKeysAPI creates CustomerManagedKeysAPI instance from provider meta
func NewCustomerManagedKeysAPI(m interface{}) CustomerManagedKeysAPI {
	return CustomerManagedKeysAPI{m.(*common.DatabricksClient), context.TODO()}
}

// CustomerManagedKeysAPI exposes the mws customerManagedKeys API
type CustomerManagedKeysAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a CustomerManagedKeysAPI) Create(cmk CustomerManagedKey) (k CustomerManagedKey, err error) {
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", cmk.AccountID)
	err = a.client.Post(a.context, customerManagedKeysAPIPath, cmk, &k)
	return
}

// Read returns the customer managed key object along with metadata
func (a CustomerManagedKeysAPI) Read(
	mwsAcctID, customerManagedKeyID string) (k CustomerManagedKey, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeyID), nil, &k)
	return
}

// Delete deletes the customer managed key object given a network id
func (a CustomerManagedKeysAPI) Delete(mwsAcctID, customerManagedKeyID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeyID), nil)
}

// List lists all the available customer managed key objects in the mws account
func (a CustomerManagedKeysAPI) List(mwsAcctID string) (kl []CustomerManagedKey, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID), nil, &kl)
	return
}

var customerManagedKeySchema = resourceMWSCustomerManagedKeysSchema()

func ResourceCustomerManagedKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomerManagedKeyCreate,
		ReadContext:   resourceCustomerManagedKeyRead,
		DeleteContext: resourceCustomerManagedKeyDelete,
		Schema:        customerManagedKeySchema,
	}
}

func resourceMWSCustomerManagedKeysSchema() map[string]*schema.Schema {
	return internal.StructToSchema(CustomerManagedKey{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["aws_key_info"].ForceNew = true
		s["account_id"].ForceNew = true
		return s
	})
}

func resourceCustomerManagedKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	var customerMangedKey CustomerManagedKey
	err := internal.DataToStructPointer(d, customerManagedKeySchema, &customerMangedKey)
	if err != nil {
		return diag.FromErr(err)
	}
	customerManagedKeyData, err := customerMangedKeyApi.Create(customerMangedKey)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(packMWSAccountID(PackagedMWSIds{
		MwsAcctID:  customerMangedKey.AccountID,
		ResourceID: customerManagedKeyData.CustomerManagedKeyID,
	}))
	return resourceCustomerManagedKeyRead(ctx, d, m)
}

func resourceCustomerManagedKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	id := d.Id()
	packagedId, err := UnpackMWSAccountID(id)
	if err != nil {
		return diag.FromErr(err)
	}
	customerManagedKey, err := customerMangedKeyApi.Read(packagedId.MwsAcctID, packagedId.ResourceID)
	if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
		log.Printf("Missing customer managed key with id: %s.", d.Id())
		d.SetId("")
		return nil
	}
	err = internal.StructToData(customerManagedKey, customerManagedKeySchema, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceCustomerManagedKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	id := d.Id()
	packagedId, err := UnpackMWSAccountID(id)
	if err != nil {
		return diag.FromErr(err)
	}
	err = customerMangedKeyApi.Delete(packagedId.MwsAcctID, packagedId.ResourceID)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
