package sql

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
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

func (a *AlertEntity) toCreateAlertApiObject(s map[string]*schema.Schema, data *schema.ResourceData) (sql.CreateAlert, error) {
	common.DataToStructPointer(data, s, a)

	var ca sql.CreateAlert
	ca.Name = a.Name
	ca.Parent = a.Parent
	ca.QueryId = a.QueryId
	ca.Rearm = a.Rearm
	ca.Options = sql.AlertOptions{
		Column:        a.Options.Column,
		CustomBody:    a.Options.CustomBody,
		CustomSubject: a.Options.CustomSubject,
		Muted:         a.Options.Muted,
		Op:            a.Options.Op,
		Value:         a.Options.Value,
	}

	return ca, nil
}

func (a *AlertEntity) toEditAlertApiObject(s map[string]*schema.Schema, data *schema.ResourceData) (sql.EditAlert, error) {
	common.DataToStructPointer(data, s, a)

	return sql.EditAlert{
		AlertId: data.Id(),
		Name:    a.Name,
		Options: sql.AlertOptions{
			Column:        a.Options.Column,
			CustomBody:    a.Options.CustomBody,
			CustomSubject: a.Options.CustomSubject,
			Muted:         a.Options.Muted,
			Op:            a.Options.Op,
			Value:         a.Options.Value,
		},
		QueryId: a.QueryId,
		Rearm:   a.Rearm,
	}, nil
}

func (a *AlertEntity) fromAPIObject(apiAlert *sql.Alert, s map[string]*schema.Schema, data *schema.ResourceData) error {
	a.Name = apiAlert.Name
	a.Parent = apiAlert.Parent
	a.QueryId = apiAlert.Query.Id
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

func ResourceSqlAlert() *schema.Resource {
	s := common.StructToSchema(AlertEntity{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		options := m["options"].Elem.(*schema.Resource)
		options.Schema["op"].ValidateFunc = validation.StringInSlice([]string{">", ">=", "<", "<=", "==", "!="}, true)
		return m
	})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var a AlertEntity
			ca, err := a.toCreateAlertApiObject(s, data)
			if err != nil {
				return err
			}
			apiAlert, err := w.Alerts.Create(ctx, ca)
			if err != nil {
				return err
			}
			data.SetId(apiAlert.Id)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			apiAlert, err := w.Alerts.GetByAlertId(ctx, data.Id())
			if err != nil {
				return err
			}
			var a AlertEntity
			return a.fromAPIObject(apiAlert, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var a AlertEntity
			ca, err := a.toEditAlertApiObject(s, data)
			if err != nil {
				return err
			}
			return w.Alerts.Update(ctx, ca)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Alerts.DeleteByAlertId(ctx, data.Id())
		},
		Schema: s,
	}.ToResource()
}
