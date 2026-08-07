package mws

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// defaultPrivateEndpointRuleCreateTimeout is the default Create timeout for
// polling a newly-submitted private endpoint rule to a cloud-provisioned
// state. The NCC CreatePrivateEndpointRule API can return immediately with
// connection_state=CREATING while the underlying cloud endpoint is still
// being provisioned; the provider polls GetPrivateEndpointRule until the
// rule reaches PENDING or ESTABLISHED (success) or CREATE_FAILED (failure).
// Users can override via a `timeouts { create = "..." }` block on the
// resource.
const defaultPrivateEndpointRuleCreateTimeout = 30 * time.Minute

// waitForPrivateEndpointRuleCreate polls Get until the rule leaves the
// CREATING state. PENDING and ESTABLISHED are success terminal states; at
// that point vpc_endpoint_id (AWS) / endpoint_name (Azure) / gcp_endpoint
// (GCP) are populated. CREATE_FAILED is terminal failure and surfaces
// ErrorMessage. REJECTED / DISCONNECTED / EXPIRED are not expected on a
// fresh Create but are treated as terminal failures defensively.
func waitForPrivateEndpointRuleCreate(ctx context.Context, acc *databricks.AccountClient, nccId, ruleId string, timeout time.Duration) (*settings.NccPrivateEndpointRule, error) {
	var result *settings.NccPrivateEndpointRule
	err := retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		rule, err := acc.NetworkConnectivity.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		switch rule.ConnectionState {
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStatePending,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateEstablished:
			result = rule
			return nil
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed:
			return retry.NonRetryableError(fmt.Errorf("private endpoint rule %s creation failed: %s", ruleId, rule.ErrorMessage))
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStateRejected,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateExpired:
			return retry.NonRetryableError(fmt.Errorf("private endpoint rule %s reached unexpected terminal state %s during creation", ruleId, rule.ConnectionState))
		default:
			msg := fmt.Sprintf("private endpoint rule %s is in state %s, waiting for cloud-side provisioning", ruleId, rule.ConnectionState)
			log.Printf("[DEBUG] %s", msg)
			return retry.RetryableError(errors.New(msg))
		}
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ResourceMwsNccPrivateEndpointRule() common.Resource {
	s := common.StructToSchema(settings.NccPrivateEndpointRule{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		for _, p := range []string{"endpoint_service", "group_id", "resource_id", "domain_names", "resource_names"} {
			common.CustomizeSchemaPath(m, p).SetOptional()
		}
		for _, p := range []string{"endpoint_service", "group_id", "resource_id"} {
			common.CustomizeSchemaPath(m, p).SetForceNew()
		}

		// Server-set fields. Marking them read-only rejects HCL writes at
		// plan time instead of silently overwriting on the next refresh.
		// psc_endpoint_uri is overwritten from the cloud platform on every
		// Read; the others are simply never read from the request body.
		for _, p := range []string{"rule_id", "endpoint_name", "connection_state", "creation_time", "updated_time", "vpc_endpoint_id", "account_id", "deactivated", "deactivated_at", "error_message"} {
			common.CustomizeSchemaPath(m, p).SetReadOnly()
		}
		common.CustomizeSchemaPath(m, "gcp_endpoint", "psc_endpoint_uri").SetReadOnly()

		common.CustomizeSchemaPath(m, "network_connectivity_config_id").SetRequired().SetForceNew()
		common.CustomizeSchemaPath(m, "enabled").SetOptional().SetComputed()

		// only one of `domain_names`, `resource_names`, `group_id` can be specified, as they are applicable
		// to different type of endpoints
		supportedFields := []string{"group_id", "resource_names", "domain_names"}
		for _, key := range supportedFields {
			conflicts := make([]string, 0, len(supportedFields)-1)
			for _, otherKey := range supportedFields {
				if key != otherKey {
					conflicts = append(conflicts, otherKey)
				}
			}
			common.CustomizeSchemaPath(m, key).SetConflictsWith(conflicts)
		}
		return m
	})
	p := common.NewPairSeparatedID("network_connectivity_config_id", "rule_id", "/")
	return common.Resource{
		Schema: s,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultPrivateEndpointRuleCreateTimeout),
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create settings.CreatePrivateEndpointRuleRequest
			common.DataToStructPointer(d, s, &create.PrivateEndpointRule)
			create.NetworkConnectivityConfigId = d.Get("network_connectivity_config_id").(string)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			rule, err := acc.NetworkConnectivity.CreatePrivateEndpointRule(ctx, create)
			if err != nil {
				return err
			}
			// Pack the ID before polling so a polling failure or context cancellation
			// still leaves a Read-able resource in state for `terraform destroy` to
			// clean up the partially-created rule on the server.
			common.StructToData(rule, s, d)
			p.Pack(d)
			// The NCC API can return immediately with connection_state=CREATING and an
			// empty vpc_endpoint_id / endpoint_name while PLLM provisions the cloud
			// endpoint. Poll until a terminal cloud-provisioned state (PENDING /
			// ESTABLISHED) is reached. When the server returns a terminal state
			// directly, the first Get exits the loop immediately.
			final, err := waitForPrivateEndpointRuleCreate(ctx, acc, create.NetworkConnectivityConfigId, rule.RuleId, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return err
			}
			return common.StructToData(final, s, d)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			rule, err := acc.NetworkConnectivity.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
			if err != nil {
				return err
			}
			return common.StructToData(rule, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}

			// only enabled, domain names & resource names are updatable
			// they do require update_mask to be set
			// resource_names are not applicable to Azure, so we exclude them from the update
			// Update does not poll: the server's update handler does not touch
			// connection_state, so the rule never re-enters CREATING.
			updateMask := []string{}
			updatePrivateEndpointRule := settings.UpdatePrivateEndpointRule{}

			if d.HasChange("enabled") {
				updateMask = append(updateMask, "enabled")
				updatePrivateEndpointRule.Enabled = d.Get("enabled").(bool)
			}
			if d.HasChange("domain_names") {
				updateMask = append(updateMask, "domain_names")
				newDomainNames := []string{}
				for _, v := range d.Get("domain_names").([]any) {
					newDomainNames = append(newDomainNames, v.(string))
				}
				updatePrivateEndpointRule.DomainNames = newDomainNames
			}
			if d.HasChange("resource_names") {
				updateMask = append(updateMask, "resource_names")
				newResourceNames := []string{}
				for _, v := range d.Get("resource_names").([]any) {
					newResourceNames = append(newResourceNames, v.(string))
				}
				updatePrivateEndpointRule.ResourceNames = newResourceNames
			}
			_, err = acc.NetworkConnectivity.UpdatePrivateEndpointRule(ctx, settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: nccId,
				PrivateEndpointRuleId:       ruleId,
				PrivateEndpointRule:         updatePrivateEndpointRule,
				UpdateMask:                  strings.Join(updateMask, ","),
			})
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, err = acc.NetworkConnectivity.DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
			return err
		},
	}
}
