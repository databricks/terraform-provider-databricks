package permissions

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RuleSetApi struct {
	client *common.DatabricksClient
	ctx    context.Context
}

func NewRuleSetApi(ctx context.Context, m interface{}) RuleSetApi {
	return RuleSetApi{m.(*common.DatabricksClient), ctx}
}

func (a RuleSetApi) Read(getRuleSetReq iam.GetRuleSetRequest) (*iam.RuleSetResponse, error) {
	if a.client.Config.AccountID != "" {
		accountClient, err := a.client.AccountClient()
		if err != nil {
			return nil, err
		}
		return accountClient.AccessControl.GetRuleSet(a.ctx, getRuleSetReq)
	}
	workspaceClient, err := a.client.WorkspaceClient()
	if err != nil {
		return nil, err
	}
	return workspaceClient.AccountAccessControlProxy.GetRuleSet(a.ctx, getRuleSetReq)
}

func (a RuleSetApi) Update(updateRuleSetReq iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error) {
	if a.client.Config.AccountID != "" {
		accountClient, err := a.client.AccountClient()
		if err != nil {
			return nil, err
		}
		return accountClient.AccessControl.UpdateRuleSet(a.ctx, updateRuleSetReq)
	}
	workspaceClient, err := a.client.WorkspaceClient()
	if err != nil {
		return nil, err
	}
	return workspaceClient.AccountAccessControlProxy.UpdateRuleSet(a.ctx, updateRuleSetReq)
}

func ResourceRuleSet() *schema.Resource {
	s := common.StructToSchema(
		iam.RuleSetUpdateRequest{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].Required = true
			m["etag"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}
			m["grant_rules"].Type = schema.TypeSet
			common.MustSchemaPath(m, "grant_rules", "principals").Type = schema.TypeSet

			return m
		})
	getAndUpdateRuleSet := func(ruleSetUpdateReq iam.UpdateRuleSetRequest, rsApi RuleSetApi) (string, error) {
		ruleSetGetRes, err := rsApi.Read(iam.GetRuleSetRequest{
			Name: ruleSetUpdateReq.Name,
			Etag: "",
		})
		if err != nil {
			return "", err
		}
		ruleSetUpdateReq.RuleSet.Etag = ruleSetGetRes.Etag
		ruleSetUpdateRes, err := rsApi.Update(ruleSetUpdateReq)
		if err != nil {
			return "", err
		}
		return ruleSetUpdateRes.Etag, nil
	}
	handleConflictAndUpdate := func(ruleSetUpdateReq iam.UpdateRuleSetRequest, rsApi RuleSetApi) (string, error) {
		ruleSetUpdateRes, err := rsApi.Update(ruleSetUpdateReq)
		if err != nil {
			if strings.EqualFold(err.Error(), "Conflict with another RuleSet operation") {
				// we need to get and update
				etag, err := getAndUpdateRuleSet(ruleSetUpdateReq, rsApi)
				return etag, err
			}
			return "", err
		}
		return ruleSetUpdateRes.Etag, err
	}
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ruleSetUpdateReq iam.UpdateRuleSetRequest
			common.DataToStructPointer(d, s, &ruleSetUpdateReq.RuleSet)
			ruleSetUpdateReq.Name = ruleSetUpdateReq.RuleSet.Name
			rsApi := NewRuleSetApi(ctx, c)
			etag, err := getAndUpdateRuleSet(ruleSetUpdateReq, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			d.SetId(ruleSetUpdateReq.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Id()
			rsApi := NewRuleSetApi(ctx, c)
			getRuleSetReq := iam.GetRuleSetRequest{
				Name: name,
				Etag: "",
			}
			if etag, ok := d.GetOk("etag"); ok {
				getRuleSetReq = iam.GetRuleSetRequest{
					Name: name,
					Etag: fmt.Sprintf("%v", etag),
				}
			}
			data, err := rsApi.Read(getRuleSetReq)
			if err != nil {
				return err
			}
			return common.StructToData(data, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ruleSetUpdateReq iam.UpdateRuleSetRequest
			common.DataToStructPointer(d, s, &ruleSetUpdateReq.RuleSet)
			ruleSetUpdateReq.Name = ruleSetUpdateReq.RuleSet.Name
			rsApi := NewRuleSetApi(ctx, c)
			if etag, ok := d.GetOk("etag"); ok {
				ruleSetUpdateReq.RuleSet.Etag = fmt.Sprintf("%v", etag)
				updatedEtag, err := handleConflictAndUpdate(ruleSetUpdateReq, rsApi)
				if err != nil {
					return err
				}
				d.Set("etag", updatedEtag)
				return nil
			}
			etag, err := getAndUpdateRuleSet(ruleSetUpdateReq, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// we remove all grant rules. Account admins will still be able to update rule set
			name := d.Id()
			rsApi := NewRuleSetApi(ctx, c)
			updateRuleSetReq := iam.UpdateRuleSetRequest{
				Name: name,
				RuleSet: iam.RuleSetUpdateRequest{
					Name: name,
					Etag: "",
				},
			}
			if etag, ok := d.GetOk("etag"); ok {
				updateRuleSetReq.RuleSet.Etag = fmt.Sprintf("%v", etag)
				updatedEtag, err := handleConflictAndUpdate(updateRuleSetReq, rsApi)
				if err != nil {
					return err
				}
				d.Set("etag", updatedEtag)
				return nil
			}
			etag, err := getAndUpdateRuleSet(updateRuleSetReq, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			return nil
		},
	}.ToResource()
}
