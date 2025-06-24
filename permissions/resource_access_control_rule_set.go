package permissions

import (
	"context"
	"errors"
	"net/http"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAccessControlRuleSet() common.Resource {
	s := common.StructToSchema(
		iam.RuleSetUpdateRequest{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["etag"].Required = false
			m["etag"].Computed = true
			m["grant_rules"].Type = schema.TypeSet
			common.MustSchemaPath(m, "grant_rules", "principals").Type = schema.TypeSet
			m["name"].ForceNew = true

			return m
		})
	readFromWsOrAcc := func(ctx context.Context, c *common.DatabricksClient, getRuleSetReq iam.GetRuleSetRequest) (*iam.RuleSetResponse, error) {
		if c.Config.AccountID != "" {
			accountClient, err := c.AccountClient()
			if err != nil {
				return nil, err
			}
			return accountClient.AccessControl.GetRuleSet(ctx, getRuleSetReq)
		}
		workspaceClient, err := c.WorkspaceClient()
		if err != nil {
			return nil, err
		}
		return workspaceClient.AccountAccessControlProxy.GetRuleSet(ctx, getRuleSetReq)
	}
	updateThroughWsOrAcc := func(ctx context.Context, c *common.DatabricksClient, updateRuleSetReq iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error) {
		if c.Config.AccountID != "" {
			accountClient, err := c.AccountClient()
			if err != nil {
				return nil, err
			}
			return accountClient.AccessControl.UpdateRuleSet(ctx, updateRuleSetReq)
		}
		workspaceClient, err := c.WorkspaceClient()
		if err != nil {
			return nil, err
		}
		return workspaceClient.AccountAccessControlProxy.UpdateRuleSet(ctx, updateRuleSetReq)
	}
	fetchLatestEtag := func(ctx context.Context, c *common.DatabricksClient, name string) (string, error) {
		ruleSetGetRes, err := readFromWsOrAcc(ctx, c, iam.GetRuleSetRequest{
			Name: name,
			Etag: "",
		})
		if err != nil {
			return "", err
		}
		return ruleSetGetRes.Etag, nil
	}
	updateRuleSet := func(ctx context.Context, c *common.DatabricksClient,
		ruleSetUpdateReq iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error) {
		latestEtag, err := fetchLatestEtag(ctx, c, ruleSetUpdateReq.Name)
		if err != nil {
			return nil, err
		}
		ruleSetUpdateReq.RuleSet.Etag = latestEtag
		ruleSetUpdateRes, err := updateThroughWsOrAcc(ctx, c, ruleSetUpdateReq)
		if err != nil {
			return nil, err
		}
		return ruleSetUpdateRes, nil
	}
	handleConflictAndUpdate := func(ctx context.Context, c *common.DatabricksClient,
		ruleSetUpdateReq iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error) {
		ruleSetUpdateRes, err := updateThroughWsOrAcc(ctx, c, ruleSetUpdateReq)
		if err != nil {
			var aerr *apierr.APIError
			if !errors.As(err, &aerr) {
				return nil, err
			}
			if aerr.StatusCode == http.StatusConflict {
				if aerr.ErrorCode == "RESOURCE_CONFLICT" {
					// we need to get and update
					etag, err := updateRuleSet(ctx, c, ruleSetUpdateReq)
					return etag, err
				}
			}
			return nil, err
		}
		return ruleSetUpdateRes, err
	}
	return common.Resource{
		Schema: s,
		CanSkipReadAfterCreateAndUpdate: func(_ *schema.ResourceData) bool {
			return true
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ruleSetUpdateReq iam.UpdateRuleSetRequest
			common.DataToStructPointer(d, s, &ruleSetUpdateReq.RuleSet)
			ruleSetUpdateReq.Name = ruleSetUpdateReq.RuleSet.Name
			response, err := updateRuleSet(ctx, c, ruleSetUpdateReq)
			if err != nil {
				return err
			}
			// Store empty etag in state as we will always use the latest etag
			response.Etag = ""
			err = common.StructToData(response, s, d)
			if err != nil {
				return err
			}
			d.SetId(ruleSetUpdateReq.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			data, err := readFromWsOrAcc(ctx, c, iam.GetRuleSetRequest{
				Name: d.Id(),
				Etag: "",
			})
			if err != nil {
				return err
			}
			// Store empty etag in state as we will always use the latest etag
			data.Etag = ""
			return common.StructToData(data, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// Fetch the latest Etag
			latestEtag, err := fetchLatestEtag(ctx, c, d.Id())
			if err != nil {
				return err
			}

			var ruleSetUpdateReq iam.UpdateRuleSetRequest
			common.DataToStructPointer(d, s, &ruleSetUpdateReq.RuleSet)
			ruleSetUpdateReq.Name = ruleSetUpdateReq.RuleSet.Name
			ruleSetUpdateReq.RuleSet.Etag = latestEtag
			response, err := handleConflictAndUpdate(ctx, c, ruleSetUpdateReq)
			if err != nil {
				return err
			}
			// Storing empty etag in state as we will always use the latest etag
			response.Etag = ""
			err = common.StructToData(response, s, d)
			if err != nil {
				return err
			}
			d.SetId(ruleSetUpdateReq.Name)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// Fetch the latest Etag
			latestEtag, err := fetchLatestEtag(ctx, c, d.Id())
			if err != nil {
				return err
			}

			// we remove all grant rules. Account admins will still be able to update rule set
			_, err = handleConflictAndUpdate(ctx, c, iam.UpdateRuleSetRequest{
				Name: d.Id(),
				RuleSet: iam.RuleSetUpdateRequest{
					Name: d.Id(),
					Etag: latestEtag,
				},
			})
			if err != nil {
				return err
			}
			return nil
		},
	}
}
