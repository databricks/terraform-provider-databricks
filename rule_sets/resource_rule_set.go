package rule_sets

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GrantRule struct {
	Principals []string `json:"principals" tf:"slice_set"`
	Role       string   `json:"role"`
}

type RuleSet struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty" tf:"optional"`
	Etag        string      `json:"etag" tf:"optional,computed"`
	GrantRules  []GrantRule `json:"grant_rules" tf:"slice_set,alias:grant_rule"`
}

type ReadRuleSetRequest struct {
	Name string `url:"name"`
	Etag string `url:"etag"`
}

type UpdateRuleSetRequest struct {
	Name string  `json:"name"`
	RS   RuleSet `json:"rule_set"`
}

type RuleSetApi struct {
	client *common.DatabricksClient
	ctx    context.Context
}

func NewRuleSetApi(ctx context.Context, m interface{}) RuleSetApi {
	return RuleSetApi{m.(*common.DatabricksClient), ctx}
}

func (a RuleSetApi) Read(readRuleSetReq ReadRuleSetRequest) (ruleSet RuleSet, err error) {
	err = a.client.Get(a.ctx, "/preview/accounts/access-control/rule-sets", readRuleSetReq, &ruleSet)
	return
}

func (a RuleSetApi) Update(updateRuleSetReq UpdateRuleSetRequest) error {
	return a.client.Put(a.ctx, "/preview/accounts/access-control/rule-sets", updateRuleSetReq)
}

func ResourceRuleSet() *schema.Resource {
	ruleSetSchema := common.StructToSchema(RuleSet{}, func(m map[string]*schema.Schema) map[string]*schema.Schema { return m })
	getAndUpdateRuleSet := func(ruleSet RuleSet, rsApi RuleSetApi) (string, error) {
		rsGet, err := rsApi.Read(ReadRuleSetRequest{
			Name: ruleSet.Name,
			Etag: "",
		})
		if err != nil {
			return "", err
		}
		err = rsApi.Update(UpdateRuleSetRequest{
			Name: ruleSet.Name,
			RS: RuleSet{
				Name:       ruleSet.Name,
				Etag:       rsGet.Etag,
				GrantRules: ruleSet.GrantRules,
			},
		})
		if err != nil {
			return "", err
		}
		// Get again for new etag
		rsGet, err = rsApi.Read(ReadRuleSetRequest{
			Name: ruleSet.Name,
			Etag: "",
		})
		if err != nil {
			return "", err
		}
		return rsGet.Etag, nil
	}
	handleConflictAndUpdate := func(ruleSet RuleSet, rsApi RuleSetApi) (string, error) {
		err := rsApi.Update(UpdateRuleSetRequest{
			Name: ruleSet.Name,
			RS: RuleSet{
				Name:       ruleSet.Name,
				Etag:       ruleSet.Etag,
				GrantRules: ruleSet.GrantRules,
			},
		})
		if err != nil {
			if err.Error() == "Conflict with another RuleSet operation" {
				// we need to get and update
				etag, err := getAndUpdateRuleSet(ruleSet, rsApi)
				return etag, err
			}
			return "", err
		}
		// get new etag and update
		updatedRuleSet, err := rsApi.Read(ReadRuleSetRequest{
			Name: ruleSet.Name,
			Etag: "",
		})
		return updatedRuleSet.Etag, err
	}
	return common.Resource{
		Schema: ruleSetSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ruleSet RuleSet
			common.DataToStructPointer(d, ruleSetSchema, &ruleSet)
			rsApi := NewRuleSetApi(ctx, c)
			etag, err := getAndUpdateRuleSet(ruleSet, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			d.SetId(ruleSet.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Id()
			rsApi := NewRuleSetApi(ctx, c)
			readRuleSetRequest := ReadRuleSetRequest{
				Name: name,
				Etag: "",
			}
			if etag, ok := d.GetOk("etag"); ok {
				readRuleSetRequest = ReadRuleSetRequest{
					Name: name,
					Etag: fmt.Sprintf("%v", etag),
				}
			}
			data, err := rsApi.Read(readRuleSetRequest)
			if err != nil {
				return err
			}
			return common.StructToData(data, ruleSetSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ruleSet RuleSet
			common.DataToStructPointer(d, ruleSetSchema, &ruleSet)
			rsApi := NewRuleSetApi(ctx, c)
			if etag, ok := d.GetOk("etag"); ok {
				ruleSet.Etag = fmt.Sprintf("%v", etag)
				updatedEtag, err := handleConflictAndUpdate(ruleSet, rsApi)
				if err != nil {
					return err
				}
				d.Set("etag", updatedEtag)
				return nil
			}
			etag, err := getAndUpdateRuleSet(ruleSet, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Id()
			rsApi := NewRuleSetApi(ctx, c)
			deleteRuleSet := RuleSet{Name: name}
			if etag, ok := d.GetOk("etag"); ok {
				deleteRuleSet.Etag = fmt.Sprintf("%v", etag)
				updatedEtag, err := handleConflictAndUpdate(deleteRuleSet, rsApi)
				if err != nil {
					return err
				}
				d.Set("etag", updatedEtag)
				return nil
			}
			etag, err := getAndUpdateRuleSet(deleteRuleSet, rsApi)
			if err != nil {
				return err
			}
			d.Set("etag", etag)
			return nil
		},
	}.ToResource()
}
