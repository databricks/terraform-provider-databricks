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

// QueryEntity defines the parameters that can be set in the resource.
type QueryEntity struct {
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
	Text             *QueryParameterText             `json:"text,omitempty"`
	Number           *QueryParameterNumber           `json:"number,omitempty"`
	Enum             *QueryParameterEnum             `json:"enum,omitempty"`
	Query            *QueryParameterQuery            `json:"query,omitempty"`
	Date             *QueryParameterDate             `json:"date,omitempty"`
	DateTime         *QueryParameterDateTime         `json:"datetime,omitempty"`
	DateTimeSec      *QueryParameterDateTimeSec      `json:"datetimesec,omitempty"`
	DateRange        *QueryParameterDateRange        `json:"date_range,omitempty"`
	DateTimeRange    *QueryParameterDateTimeRange    `json:"datetime_range,omitempty"`
	DateTimeSecRange *QueryParameterDateTimeSecRange `json:"datetimesec_range,omitempty"`
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
	// Value iff `multiple == nil`
	Value string `json:"value,omitempty"`
	// Values iff `multiple != nil`
	Values []string `json:"values,omitempty"`

	Options  []string                     `json:"options"`
	Multiple *QueryParameterAllowMultiple `json:"multiple,omitempty"`
}

// QueryParameterQuery ...
type QueryParameterQuery struct {
	// Value iff `multiple == nil`
	Value string `json:"value,omitempty"`
	// Values iff `multiple != nil`
	Values []string `json:"values,omitempty"`

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

// QueryParameterDateRange ...
type QueryParameterDateRange struct {
	Value string `json:"value"`
}

// QueryParameterDateTimeRange ...
type QueryParameterDateTimeRange struct {
	Value string `json:"value"`
}

// QueryParameterDateTimeSecRange ...
type QueryParameterDateTimeSecRange struct {
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

func (q *QueryEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Query, error) {
	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, q); err != nil {
		return nil, err
	}

	// Transform to API object.
	var aq api.Query
	aq.ID = data.Id()
	aq.DataSourceID = q.DataSourceID
	aq.Name = q.Name
	aq.Description = q.Description
	aq.Query = q.Query
	aq.Tags = append([]string{}, q.Tags...)

	if s := q.Schedule; s != nil {
		aq.Schedule = &api.QuerySchedule{
			Interval: s.Interval,
		}
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
					Options:        strings.Join(p.Enum.Options, "\n"),
				}
				if p.Enum.Multiple != nil {
					tmp.Values = p.Enum.Values
					tmp.Multi = p.Enum.Multiple.toAPIObject()
				} else {
					tmp.Values = []string{p.Enum.Value}
				}
				iface = tmp
			case p.Query != nil:
				tmp := api.QueryParameterQuery{
					QueryParameter: ap,
					QueryID:        p.Query.QueryID,
				}
				if p.Query.Multiple != nil {
					tmp.Values = p.Query.Values
					tmp.Multi = p.Query.Multiple.toAPIObject()
				} else {
					tmp.Values = []string{p.Query.Value}
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
			case p.DateRange != nil:
				iface = api.QueryParameterDateRange{
					QueryParameter: ap,
					Value:          p.DateRange.Value,
				}
			case p.DateTimeRange != nil:
				iface = api.QueryParameterDateTimeRange{
					QueryParameter: ap,
					Value:          p.DateTimeRange.Value,
				}
			case p.DateTimeSecRange != nil:
				iface = api.QueryParameterDateTimeSecRange{
					QueryParameter: ap,
					Value:          p.DateTimeSecRange.Value,
				}
			default:
				log.Fatalf("Don't know what to do for QueryParameter...")
			}

			aq.Options.Parameters = append(aq.Options.Parameters, iface)
		}
	}

	return &aq, nil
}

func (q *QueryEntity) fromAPIObject(aq *api.Query, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	q.DataSourceID = aq.DataSourceID
	q.Name = aq.Name
	q.Description = aq.Description
	q.Query = aq.Query
	q.Tags = append([]string{}, aq.Tags...)

	if s := aq.Schedule; s != nil {
		q.Schedule = &QuerySchedule{
			Interval: s.Interval,
		}
	}

	if aq.Options != nil {
		q.Parameter = nil

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
					Options:  strings.Split(apv.Options, "\n"),
					Multiple: newQueryParameterAllowMultiple(apv.Multi),
				}
				if p.Enum.Multiple != nil {
					p.Enum.Values = apv.Values
				} else {
					p.Enum.Value = apv.Values[0]
				}
			case *api.QueryParameterQuery:
				p.Name = apv.Name
				p.Title = apv.Title
				p.Query = &QueryParameterQuery{
					QueryID:  apv.QueryID,
					Multiple: newQueryParameterAllowMultiple(apv.Multi),
				}
				if p.Query.Multiple != nil {
					p.Query.Values = apv.Values
				} else {
					p.Query.Value = apv.Values[0]
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
			case *api.QueryParameterDateRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateRange = &QueryParameterDateRange{
					Value: apv.Value,
				}
			case *api.QueryParameterDateTimeRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeRange = &QueryParameterDateTimeRange{
					Value: apv.Value,
				}
			case *api.QueryParameterDateTimeSecRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeSecRange = &QueryParameterDateTimeSecRange{
					Value: apv.Value,
				}
			default:
				log.Fatalf("Don't know what to do for type: %#v", reflect.TypeOf(apv).String())
			}

			q.Parameter = append(q.Parameter, p)
		}
	}

	// Transform to ResourceData.
	if err := common.StructToData(*q, schema, data); err != nil {
		return err
	}

	return nil
}

// ResourceQuery ...
func ResourceQuery() *schema.Resource {
	s := common.StructToSchema(
		QueryEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			aqNew, err := api.NewWrapper(ctx, c).CreateQuery(aq)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(aqNew.ID)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			aqNew, err := api.NewWrapper(ctx, c).ReadQuery(aq)
			if err != nil {
				return err
			}

			return q.fromAPIObject(aqNew, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			_, err = api.NewWrapper(ctx, c).UpdateQuery(aq)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			return nil
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return api.NewWrapper(ctx, c).DeleteQuery(aq)
		},
		Schema: s,
	}.ToResource()
}
