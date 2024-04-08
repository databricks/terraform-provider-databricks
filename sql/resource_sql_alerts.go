package sql

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type AlertOptions struct {
	Column           string `json:"column"`
	Op               string `json:"op"`
	Value            string `json:"value"`
	Muted            bool   `json:"muted,omitempty"`
	CustomBody       string `json:"custom_body,omitempty"`
	CustomSubject    string `json:"custom_subject,omitempty"`
	EmptyResultState string `json:"empty_result_state,omitempty"`
}

type AlertEntity struct {
	Name      string        `json:"name"`
	QueryId   string        `json:"query_id"`
	Rearm     int           `json:"rearm,omitempty"`
	Options   *AlertOptions `json:"options"`
	Parent    string        `json:"parent,omitempty" tf:"suppress_diff,force_new"`
	CreatedAt string        `json:"created_at,omitempty" tf:"computed"`
	UpdatedAt string        `json:"updated_at,omitempty" tf:"computed"`
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
	// This is a workaround for Go SDK problem, will be fixed there.
	var err error
	if a.Options.EmptyResultState != "" {
		err = ca.Options.EmptyResultState.Set(a.Options.EmptyResultState)
	}
	return ca, err
}

func (a *AlertEntity) toEditAlertApiObject(s map[string]*schema.Schema, data *schema.ResourceData) (sql.EditAlert, error) {
	common.DataToStructPointer(data, s, a)

	ea := sql.EditAlert{
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
	}

	var err error
	if a.Options.EmptyResultState != "" {
		err = ea.Options.EmptyResultState.Set(a.Options.EmptyResultState)
	}
	return ea, err
}

func (a *AlertEntity) fromAPIObject(apiAlert *sql.Alert, s map[string]*schema.Schema, data *schema.ResourceData) error {
	a.Name = apiAlert.Name
	a.Parent = apiAlert.Parent
	if apiAlert.Query != nil {
		a.QueryId = apiAlert.Query.Id
	} else {
		log.Printf("[WARN] Query object is nil in alert '%s' (id: %s) ", apiAlert.Name, apiAlert.Id)
	}
	a.Rearm = apiAlert.Rearm
	a.CreatedAt = apiAlert.CreatedAt
	a.UpdatedAt = apiAlert.UpdatedAt

	if apiAlert.Options != nil {
		a.Options = &AlertOptions{
			Column:           apiAlert.Options.Column,
			Op:               apiAlert.Options.Op,
			Muted:            apiAlert.Options.Muted,
			CustomBody:       apiAlert.Options.CustomBody,
			CustomSubject:    apiAlert.Options.CustomSubject,
			EmptyResultState: apiAlert.Options.EmptyResultState.String(),
		}

		// value can be a string or a float64 - unfortunately this can't be encoded in OpenAPI yet
		switch value := apiAlert.Options.Value.(type) {
		case string:
			a.Options.Value = value
		case float64:
			a.Options.Value = strconv.FormatFloat(value, 'f', 0, 64)
		case bool:
			a.Options.Value = strconv.FormatBool(value)
		default:
			return fmt.Errorf("unexpected type for value: %T", value)
		}
	} else {
		log.Printf("[WARN] Options object is nil in alert '%s' (id: %s) ", apiAlert.Name, apiAlert.Id)
		a.Options = &AlertOptions{}
	}

	return common.StructToData(a, s, data)
}

func ResourceSqlAlert() common.Resource {
	s := common.StructToSchema(AlertEntity{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.MustSchemaPath(m, "options", "op").ValidateFunc = validation.StringInSlice([]string{">", ">=", "<", "<=", "==", "!="}, true)
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
				log.Printf("[WARN] error getting alert by ID: %v", err)
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
	}
}
