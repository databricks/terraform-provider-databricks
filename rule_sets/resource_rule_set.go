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
			etagKey := fmt.Sprintf("%s:etag", ruleSet.Name)
			d.Set(etagKey, etag)
			d.SetId(ruleSet.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Id()
			etagKey := fmt.Sprintf("%s:etag", name)
			rsApi := NewRuleSetApi(ctx, c)
			readRuleSetRequest := ReadRuleSetRequest{
				Name: name,
				Etag: "",
			}
			if etag, ok := d.GetOk(etagKey); ok {
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
			etagKey := fmt.Sprintf("%s:etag", ruleSet.Name)
			if etag, ok := d.GetOk(etagKey); ok {
				err := rsApi.Update(UpdateRuleSetRequest{
					Name: ruleSet.Name,
					RS: RuleSet{
						Name:       ruleSet.Name,
						Etag:       fmt.Sprintf("%v", etag),
						GrantRules: ruleSet.GrantRules,
					},
				})
				if err != nil && err.Error() == "Conflict with another RuleSet operation" {
					// we need to get and updated
					etag, err = getAndUpdateRuleSet(ruleSet, rsApi)
					d.Set(etagKey, etag)
					return err
				}
				return err
			}
			etag, err := getAndUpdateRuleSet(ruleSet, rsApi)
			if err != nil {
				return err
			}
			d.Set(etagKey, etag)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Id()
			etagKey := fmt.Sprintf("%s:etag", name)
			rsApi := NewRuleSetApi(ctx, c)
			deleteRuleSet := RuleSet{Name: name}
			if etag, ok := d.GetOk(etagKey); ok {
				err := rsApi.Update(UpdateRuleSetRequest{
					Name: name,
					RS: RuleSet{
						Name: name,
						Etag: fmt.Sprintf("%v", etag),
					},
				})
				if err != nil && err.Error() == "Conflict with another RuleSet operation" {
					// we need to get and updated
					etag, err = getAndUpdateRuleSet(deleteRuleSet, rsApi)
					d.Set(etagKey, etag)
					return err
				}
				return err
			}
			etag, err := getAndUpdateRuleSet(deleteRuleSet, rsApi)
			if err != nil {
				return err
			}
			d.Set(etagKey, etag)
			return nil
		},
	}.ToResource()
}
