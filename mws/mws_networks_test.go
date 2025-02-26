package mws_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccNetworks(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET") // marker for AWS test env
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_networks" "my_network" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-test-{var.RANDOM}"
			vpc_id       = "vpc-11111111"
			subnet_ids   = [
				"subnet-11111111",
				"subnet-99999999"
			]
			security_group_ids = [
				"sg-99999999"
			]
		}`,
	})
}

type checkResourceActions struct {
	address         string
	expectedActions []tfjson.Action
}

func (c checkResourceActions) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	var resource *tfjson.ResourceChange
	for _, r := range req.Plan.ResourceChanges {
		if r.Address == c.address {
			resource = r
			break
		}
	}
	if resource == nil {
		foundAddresses := make([]string, 0, len(req.Plan.ResourceChanges))
		for _, r := range req.Plan.ResourceChanges {
			foundAddresses = append(foundAddresses, r.Address)
		}
		resp.Error = fmt.Errorf("address %s not found in plan, found changes for the following resources: %s", c.address, strings.Join(foundAddresses, ", "))
		return
	}
	if err := c.checkActions(resource.Change.Actions); err != nil {
		resp.Error = err
	}
}

// checkActions compares the actual actions for a resource to the expected actions.
func (c checkResourceActions) checkActions(actualActions tfjson.Actions) error {
	if len(c.expectedActions) != len(actualActions) {
		return fmt.Errorf("mismatch in number of actions. Expected %d actions, actual %d actions. Expected actions: %v, actual actions: %v", len(c.expectedActions), len(actualActions), c.expectedActions, actualActions)
	}
	for i := range actualActions {
		expectedAction := c.expectedActions[i]
		actualAction := actualActions[i]
		if expectedAction != actualAction {
			return fmt.Errorf("mismatch in action %d. Expected %q, actual %q. Expected actions: %v, actual actions: %v", i, expectedAction, actualAction, c.expectedActions, actualActions)
		}
	}
	return nil
}

func TestMwsAccGcpPscNetworks(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_networks" "my_network" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-test-{var.STICKY_RANDOM}"
			gcp_network_info {
			  network_project_id = "{env.GOOGLE_PROJECT}"
			  vpc_id = "{env.TEST_VPC_ID}"
			  subnet_id = "{env.TEST_SUBNET_ID}"
			  subnet_region = "{env.GOOGLE_REGION}"
              pod_ip_range_name = "pods"
              service_ip_range_name = "svc"
            }
		}`,
	}, acceptance.Step{
		// Changing the name should cause the network to be recreated.
		Template: `
		resource "databricks_mws_networks" "my_network" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-test-new-{var.STICKY_RANDOM}"
			gcp_network_info {
			  network_project_id = "{env.GOOGLE_PROJECT}"
			  vpc_id = "{env.TEST_VPC_ID}"
			  subnet_id = "{env.TEST_SUBNET_ID}"
			  subnet_region = "{env.GOOGLE_REGION}"
              pod_ip_range_name = "pods"
              service_ip_range_name = "svc"
            }
		}`,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				checkResourceActions{"databricks_mws_networks.my_network", []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate}},
			},
		},
	}, acceptance.Step{
		// Removing the pod_ip_range_name and service_ip_range_name fields should
		// 1. plan a resource update instead of a delete-create.
		// 2. result in those fields being removed from the state.
		Template: `
		resource "databricks_mws_networks" "my_network" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-test-new-{var.STICKY_RANDOM}"
			gcp_network_info {
			  network_project_id = "{env.GOOGLE_PROJECT}"
			  vpc_id = "{env.TEST_VPC_ID}"
			  subnet_id = "{env.TEST_SUBNET_ID}"
			  subnet_region = "{env.GOOGLE_REGION}"
            }
		}`,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				checkResourceActions{"databricks_mws_networks.my_network", []tfjson.Action{tfjson.ActionUpdate}},
			},
		},
		Check: func(s *terraform.State) error {
			r := s.RootModule().Resources["databricks_mws_networks.my_network"].Primary
			assert.Empty(t, r.Attributes["gcp_network_info.0.pod_ip_range_name"])
			assert.Empty(t, r.Attributes["gcp_network_info.0.service_ip_range_name"])
			return nil
		},
	})
}
