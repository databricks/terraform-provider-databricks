package sqlanalytics

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Query ...
type Query struct {
	ID           string           `json:"id,omitempty" tf:"computed"`
	DataSourceID string           `json:"data_source_id"`
	Name         string           `json:"name"`
	Description  string           `json:"description,omitempty"`
	Query        string           `json:"query"`
	Schedule     *QuerySchedule   `json:"schedule,omitempty"`
	Tags         []string         `json:"tags,omitempty"`
	Parameter    []QueryParameter `json:"parameter,omitempty"`
}

// QuerySchedule ...
type QuerySchedule struct {
	Interval int `json:"interval"`
}

// QueryParameter ...
type QueryParameter struct {
	Name  string `json:"name"`
	Title string `json:"title,omitempty"`

	// Type specific structs.
	// Only one of them may be set.
	Text        *QueryParameterText        `json:"text,omitempty"`
	Number      *QueryParameterNumber      `json:"number,omitempty"`
	Enum        *QueryParameterEnum        `json:"enum,omitempty"`
	Query       *QueryParameterQuery       `json:"query,omitempty"`
	Date        *QueryParameterDate        `json:"date,omitempty"`
	DateTime    *QueryParameterDateTime    `json:"datetime,omitempty"`
	DateTimeSec *QueryParameterDateTimeSec `json:"datetimesec,omitempty"`
}

// QueryParameterText ...
type QueryParameterText struct {
	Value string `json:"value"`
}

// QueryParameterNumber ...
type QueryParameterNumber struct {
	Value float64 `json:"value"`
}

// QueryParameterEnum ...
type QueryParameterEnum struct {
	Value    string                       `json:"value"`
	Options  []string                     `json:"options"`
	Multiple *QueryParameterAllowMultiple `json:"multiple,omitempty"`
}

// QueryParameterQuery ...
type QueryParameterQuery struct {
	Value    string                       `json:"value"`
	QueryID  string                       `json:"query_id"`
	Multiple *QueryParameterAllowMultiple `json:"multiple,omitempty"`
}

// QueryParameterDate ...
type QueryParameterDate struct {
	Value string `json:"value"`
}

// QueryParameterDateTime ...
type QueryParameterDateTime struct {
	Value string `json:"value"`
}

// QueryParameterDateTimeSec ...
type QueryParameterDateTimeSec struct {
	Value string `json:"value"`
}

// QueryParameterAllowMultiple ...
type QueryParameterAllowMultiple struct {
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Separator string `json:"separator"`
}

func (q *QueryParameterAllowMultiple) toAPIObject() *api.QueryParameterMultipleValuesOptions {
	return &api.QueryParameterMultipleValuesOptions{
		Prefix:    q.Prefix,
		Suffix:    q.Suffix,
		Separator: q.Separator,
	}
}

func newQueryParameterAllowMultiple(aq *api.QueryParameterMultipleValuesOptions) *QueryParameterAllowMultiple {
	if aq == nil {
		return nil
	}
	return &QueryParameterAllowMultiple{
		Prefix:    aq.Prefix,
		Suffix:    aq.Suffix,
		Separator: aq.Separator,
	}
}

type queryResource struct {
	schema map[string]*schema.Schema
}

func (r *queryResource) toAPIObject(d *schema.ResourceData) (*api.Query, error) {
	var q Query

	// Transform from ResourceData.
	if err := common.DataToStructPointer(d, r.schema, &q); err != nil {
		return nil, err
	}

	// Transform to API object.
	var aq api.Query
	aq.ID = q.ID
	aq.DataSourceID = q.DataSourceID
	aq.Name = q.Name
	aq.Description = q.Description
	aq.Query = q.Query

	if s := q.Schedule; s != nil {
		aq.Schedule = &api.QuerySchedule{
			Interval: s.Interval,
		}
	}

	if len(q.Tags) > 0 {
		aq.Tags = append(aq.Tags, q.Tags...)
	}

	if len(q.Parameter) > 0 {
		aq.Options = &api.QueryOptions{}
		for _, p := range q.Parameter {
			ap := api.QueryParameter{
				Name:  p.Name,
				Title: p.Title,
			}

			var iface interface{}

			switch {
			case p.Text != nil:
				iface = api.QueryParameterText{
					QueryParameter: ap,
					Value:          p.Text.Value,
				}
			case p.Number != nil:
				iface = api.QueryParameterNumber{
					QueryParameter: ap,
					Value:          p.Number.Value,
				}
			case p.Enum != nil:
				tmp := api.QueryParameterEnum{
					QueryParameter: ap,
					Value:          p.Enum.Value,
					Options:        strings.Join(p.Enum.Options, "\n"),
				}
				if p.Enum.Multiple != nil {
					tmp.Multi = p.Enum.Multiple.toAPIObject()
				}
				iface = tmp
			case p.Query != nil:
				tmp := api.QueryParameterQuery{
					QueryParameter: ap,
					Value:          p.Query.Value,
					QueryID:        p.Query.QueryID,
				}
				if p.Query.Multiple != nil {
					tmp.Multi = p.Query.Multiple.toAPIObject()
				}
				iface = tmp
			case p.Date != nil:
				iface = api.QueryParameterDate{
					QueryParameter: ap,
					Value:          p.Date.Value,
				}
			case p.DateTime != nil:
				iface = api.QueryParameterDateTime{
					QueryParameter: ap,
					Value:          p.DateTime.Value,
				}
			case p.DateTimeSec != nil:
				iface = api.QueryParameterDateTimeSec{
					QueryParameter: ap,
					Value:          p.DateTimeSec.Value,
				}
			default:
				log.Fatalf("Don't know what to do for QueryParameter...")
			}

			aq.Options.Parameters = append(aq.Options.Parameters, iface)
		}
	}

	return &aq, nil
}

func (r *queryResource) fromAPIObject(aq *api.Query, d *schema.ResourceData) error {
	var q Query

	// Transform from API object.
	q.ID = aq.ID
	q.DataSourceID = aq.DataSourceID
	q.Name = aq.Name
	q.Description = aq.Description
	q.Query = aq.Query

	if s := aq.Schedule; s != nil {
		q.Schedule = &QuerySchedule{
			Interval: s.Interval,
		}
	}

	if len(aq.Tags) > 0 {
		q.Tags = append(q.Tags, aq.Tags...)
	}

	if aq.Options != nil {
		for _, ap := range aq.Options.Parameters {
			var p QueryParameter
			switch apv := ap.(type) {
			case *api.QueryParameterText:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Text = &QueryParameterText{
					Value: apv.Value,
				}
			case *api.QueryParameterNumber:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Number = &QueryParameterNumber{
					Value: apv.Value,
				}
			case *api.QueryParameterEnum:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Enum = &QueryParameterEnum{
					Value:    apv.Value,
					Options:  strings.Split(apv.Options, "\n"),
					Multiple: newQueryParameterAllowMultiple(apv.Multi),
				}
			case *api.QueryParameterQuery:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Query = &QueryParameterQuery{
					Value:    apv.Value,
					QueryID:  apv.QueryID,
					Multiple: newQueryParameterAllowMultiple(apv.Multi),
				}
			case *api.QueryParameterDate:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Date = &QueryParameterDate{
					Value: apv.Value,
				}
			case *api.QueryParameterDateTime:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTime = &QueryParameterDateTime{
					Value: apv.Value,
				}
			case *api.QueryParameterDateTimeSec:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeSec = &QueryParameterDateTimeSec{
					Value: apv.Value,
				}
			default:
				log.Fatalf("Don't know what to do for type: %#v", reflect.TypeOf(apv).String())
			}

			q.Parameter = append(q.Parameter, p)
		}
	}

	// Transform to ResourceData.
	if err := common.StructToData(q, r.schema, d); err != nil {
		return err
	}

	return nil
}

func (r *queryResource) create(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	aq, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	aqNew, err := w.CreateQuery(aq)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(aqNew, d)
	if err != nil {
		return err
	}

	d.SetId(aqNew.ID)
	return nil
}

func (r *queryResource) read(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	aq, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	aqNew, err := w.ReadQuery(aq)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(aqNew, d)
	if err != nil {
		return err
	}

	return nil
}

func (r *queryResource) update(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	aq, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	aqNew, err := w.UpdateQuery(aq)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(aqNew, d)
	if err != nil {
		return err
	}

	return nil
}

func (r *queryResource) delete(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	aq, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	err = w.DeleteQuery(aq)
	if err != nil {
		return err
	}

	return nil
}

// ResourceQuery ...
func ResourceQuery() *schema.Resource {
	r := queryResource{
		common.StructToSchema(
			Query{},
			func(m map[string]*schema.Schema) map[string]*schema.Schema {
				return m
			}),
	}
	return common.Resource{
		Schema: r.schema,
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
	}.ToResource()
}
