package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIPAccessList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPACLCreate,
		Read:   resourceIPACLRead,
		Update: resourceIPACLUpdate,
		Delete: resourceIPACLDelete,

		Schema: map[string]*schema.Schema{
			"preview_ipacl_enabled": {
				Type:     schema.TypeBool,
				Optional: false,
			},
			"ip_acl": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Optional: false,
						},
						"type": {
							Type: schema.TypeString,
							ExactlyOneOf: []string{
								"WHITELIST",
								"BLACKLIST",
							},
						},
						"ip_addresses": {
							Type:     schema.TypeList,
							Optional: false,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.IsIPv4Address,
							},
						},
					},
				},
			},
		},
	}
}

func resourceIPACLCreate(d *schema.ResourceData, m interface{}) error {
	return resourceIPACLRead(d, m)
}

func resourceIPACLRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIPACLUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPACLRead(d, m)
}

func resourceIPACLDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

/* Enabled model
{
  "enableIpAccessLists": "true",
}
*/

/* Status feature
curl -X -n \
 https://<databricks-instance>/api/2.0/preview/workspace-conf?keys=enableIpAccessLists
*/

/* Enable
curl -X PATCH -n \
  https://<databricks-instance>/api/2.0/preview/workspace-conf \
  -d '{
    "enableIpAccessLists": "true",
    }'
*/

/* Disable
curl -X PATCH -n \
  https://<databricks-instance>/api/2.0/preview/workspace-conf \
  -d '{
    "enableIpAccessLists": "false",
    }'
*/

/* List model
label — Label for this list.
list_type — Either WHITELIST (allow list) or BLACKLIST (a block list, which means exclude even if in allow list).
ip_addresses — A JSON array of IP addresses and CIDR ranges, as String values.
{
  "ip_access_list": {
    "list_id": "<list-id>",
    "label": "office",
    "ip_addresses": [
        "1.1.1.1",
        "2.2.2.2/21"
    ],
    "address_count": 2,
    "list_type": "WHITELIST",
    "created_at": 1578423494457,
    "creator_user_id": 6476783916686816,
    "updated_at": 1578423494457,
    "updator_user_id": 6476783916686816,
    "enabled": true
  }
}
*/

/* Add
curl -X POST -n \
  https://<databricks-instance>/api/2.0/preview/ip-access-lists
  -d '{
    "label": "office",
    "list_type": "WHITELIST",
    "ip_addresses": [
        "1.1.1.1",
        "2.2.2.2/21"
      ]
    }'
*/

// For all whitelists and blacklists combined, the API supports a maximum of 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to exceed that number return error 400 with error_code value QUOTA_EXCEEDED.

/* Update
curl -X PUT -n \
  https://<databricks-instance>/api/2.0/preview/ip-access-lists/<list-id>
  -d '{
    "label": "office",
    "list_type": "WHITELIST",
    "ip_addresses": [
        "1.1.1.1",
        "2.2.2.2/21"
      ],
    "enabled": "false"
    }'
*/

/* Delete
curl -X DELETE -n \
  https://<databricks-instance>/api/2.0/preview/ip-access-lists/<list-id>
  -d '{
    "label": "office",
    "list_type": "WHITELIST",
    "ip_addresses": [
        "1.1.1.1",
        "2.2.2.2/21"
      ],
    "enabled": "false"
    }'
*/
