package sql

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type AlertOptions struct {
	Column        string `json:"column"`
	Op            string `json:"op"`
	Value         string `json:"value"`
	Muted         bool   `json:"muted,omitempty" tf:"suppress_diff"`
	CustomBody    string `json:"custom_body,omitempty"`
	CustomSubject string `json:"custom_subject,omitempty"`
}

type AlertEntity struct {
	Name    string        `json:"name"`
	QueryId string        `json:"query_id"`
	Rearm   int           `json:"rearm,omitempty"`
	Options *AlertOptions `json:"options"`
	Parent  string        `json:"parent,omitempty" tf:"suppress_diff,force_new"`
}

func (a *AlertEntity) toCreateAlertApiObject(s map[string]*schema.Schema, data *schema.ResourceData) (*api.CreateAlert, error) {
	common.DataToStructPointer(data, s, a)

	var ca api.CreateAlert
	ca.Name = a.Name
	ca.Parent = a.Parent
	ca.QueryId = a.QueryId
	ca.Rearm = a.Rearm
	ca.Options = &api.AlertOptions{
		Column:        a.Options.Column,
		CustomBody:    a.Options.CustomBody,
		CustomSubject: a.Options.CustomSubject,
		Muted:         a.Options.Muted,
		Op:            a.Options.Op,
		Value:         a.Options.Value,
	}

	return &ca, nil
}

func (a *AlertEntity) fromAPIObject(apiAlert *api.Alert, s map[string]*schema.Schema, data *schema.ResourceData) error {
	a.Name = apiAlert.Name
	a.Parent = apiAlert.Parent
	a.QueryId = apiAlert.Query.ID
	a.Rearm = apiAlert.Rearm
	a.Options = &AlertOptions{
		Column:        apiAlert.Options.Column,
		Op:            apiAlert.Options.Op,
		Value:         apiAlert.Options.Value,
		Muted:         apiAlert.Options.Muted,
		CustomBody:    apiAlert.Options.CustomBody,
		CustomSubject: apiAlert.Options.CustomSubject,
	}

	return common.StructToData(a, s, data)
}
func NewAlertApi(ctx context.Context, m any) AlertApi {
	return AlertApi{m.(*common.DatabricksClient), ctx}
}

type AlertApi struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a AlertApi) Create(q *api.CreateAlert) (*api.Alert, error) {
	var r api.Alert
	err := a.client.Post(a.context, "/preview/sql/alerts", q, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a AlertApi) Get(id string) (*api.Alert, error) {
	var alert api.Alert
	path := fmt.Sprintf("/preview/sql/alerts/%v", id)
	err := a.client.Get(a.context, path, nil, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (a AlertApi) Delete(id string) error {
	path := fmt.Sprintf("/preview/sql/alerts/%v", id)
	return a.client.Get(a.context, path, nil, nil)
}

func (a AlertApi) Update(id string, request *api.CreateAlert) error {
	path := fmt.Sprintf("/preview/sql/alerts/%v", id)
	return a.client.Put(a.context, path, request)
}

func ResourceSqlAlert() *schema.Resource {
	s := common.StructToSchema(AlertEntity{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		options := m["options"].Elem.(*schema.Resource)
		options.Schema["op"].ValidateFunc = validation.StringInSlice([]string{">", ">=", "<", "<=", "==", "!="}, true)
		return m
	})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var a AlertEntity
			ca, err := a.toCreateAlertApiObject(s, data)
			if err != nil {
				return err
			}

			apiAlert, err := NewAlertApi(ctx, c).Create(ca)
			if err != nil {
				return err
			}

			data.SetId(apiAlert.Id)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			apiAlert, err := NewAlertApi(ctx, c).Get(data.Id())
			if err != nil {
				return err
			}

			var a AlertEntity
			return a.fromAPIObject(apiAlert, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var a AlertEntity
			ca, err := a.toCreateAlertApiObject(s, data)
			if err != nil {
				return err
			}

			return NewAlertApi(ctx, c).Update(data.Id(), ca)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			return NewAlertApi(ctx, c).Delete(data.Id())
		},
		Schema: s,
	}.ToResource()
}
