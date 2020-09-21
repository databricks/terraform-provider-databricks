package mws

import (
	"fmt"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewCustomerManagedKeysAPI creates CustomerManagedKeysAPI instance from provider meta
func NewCustomerManagedKeysAPI(m interface{}) CustomerManagedKeysAPI {
	return CustomerManagedKeysAPI{client: m.(*common.DatabricksClient)}
}

// CustomerManagedKeysAPI exposes the mws customerManagedKeys API
type CustomerManagedKeysAPI struct {
	client *common.DatabricksClient
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a CustomerManagedKeysAPI) Create(cmk CustomerManagedKey) (k CustomerManagedKey, err error) {
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", cmk.AccountID)
	err = a.client.Post(customerManagedKeysAPIPath, cmk, &k)
	return
}

// Read returns the customer managed key object along with metadata
func (a CustomerManagedKeysAPI) Read(
	mwsAcctID, customerManagedKeyID string) (k CustomerManagedKey, err error) {
	err = a.client.Get(fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeyID), nil, &k)
	return
}

// Delete deletes the customer managed key object given a network id
func (a CustomerManagedKeysAPI) Delete(mwsAcctID, customerManagedKeyID string) error {
	return a.client.Delete(fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeyID), nil)
}

// List lists all the available customer managed key objects in the mws account
func (a CustomerManagedKeysAPI) List(mwsAcctID string) (kl []CustomerManagedKey, err error) {
	err = a.client.Get(fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID), nil, &kl)
	return
}

var customerManagedKeySchema = resourceMWSCustomerManagedKeysSchema()

func ResourceCustomerManagedKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceCustomerManagedKeyCreate,
		Read:   resourceCustomerManagedKeyRead,
		Delete: resourceCustomerManagedKeyDelete,
		Schema: customerManagedKeySchema,
	}
}

func resourceMWSCustomerManagedKeysSchema() map[string]*schema.Schema {
	return internal.StructToSchema(CustomerManagedKey{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["aws_key_info"].ForceNew = true
		s["account_id"].ForceNew = true
		return s
	})
}

func resourceCustomerManagedKeyCreate(d *schema.ResourceData, m interface{}) error {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	var customerMangedKey CustomerManagedKey
	err := internal.DataToStructPointer(d, customerManagedKeySchema, &customerMangedKey)
	if err != nil {
		return err
	}
	customerManagedKeyData, err := customerMangedKeyApi.Create(customerMangedKey)
	if err != nil {
		return err
	}
	d.SetId(packMWSAccountID(PackagedMWSIds{
		MwsAcctID:  customerMangedKey.AccountID,
		ResourceID: customerManagedKeyData.CustomerManagedKeyID,
	}))
	return resourceCustomerManagedKeyRead(d, m)
}

func resourceCustomerManagedKeyRead(d *schema.ResourceData, m interface{}) error {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	id := d.Id()
	packagedId, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	customerManagedKey, err := customerMangedKeyApi.Read(packagedId.MwsAcctID, packagedId.ResourceID)
	if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
		log.Printf("Missing customer managed key with id: %s.", d.Id())
		d.SetId("")
		return nil
	}
	err = internal.StructToData(customerManagedKey, customerManagedKeySchema, d)
	return err
}

func resourceCustomerManagedKeyDelete(d *schema.ResourceData, m interface{}) error {
	customerMangedKeyApi := NewCustomerManagedKeysAPI(m)
	id := d.Id()
	packagedId, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = customerMangedKeyApi.Delete(packagedId.MwsAcctID, packagedId.ResourceID)
	return err
}
