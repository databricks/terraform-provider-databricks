package sql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
	RunAsRole    string           `json:"run_as_role,omitempty"`
}

// QuerySchedule ...
type QuerySchedule struct {
	Continuous *QueryScheduleContinuous `json:"continuous,omitempty"`
	Daily      *QueryScheduleDaily      `json:"daily,omitempty"`
	Weekly     *QueryScheduleWeekly     `json:"weekly,omitempty"`
}

// QueryScheduleContinuous ...
type QueryScheduleContinuous struct {
	IntervalSeconds int    `json:"interval_seconds"`
	UntilDate       string `json:"until_date,omitempty"`
}

// QueryScheduleDaily ...
type QueryScheduleDaily struct {
	IntervalDays int    `json:"interval_days"`
	TimeOfDay    string `json:"time_of_day"`
	UntilDate    string `json:"until_date,omitempty"`
}

// QueryScheduleWeekly ...
type QueryScheduleWeekly struct {
	IntervalWeeks int    `json:"interval_weeks"`
	DayOfWeek     string `json:"day_of_week"`
	TimeOfDay     string `json:"time_of_day"`
	UntilDate     string `json:"until_date,omitempty"`
}

// QueryParameter ...
type QueryParameter struct {
	Name  string `json:"name"`
	Title string `json:"title,omitempty"`

	// Type specific structs.
	// Only one of them may be set.
	Text             *QueryParameterText          `json:"text,omitempty"`
	Number           *QueryParameterNumber        `json:"number,omitempty"`
	Enum             *QueryParameterEnum          `json:"enum,omitempty"`
	Query            *QueryParameterQuery         `json:"query,omitempty"`
	Date             *QueryParameterDateLike      `json:"date,omitempty"`
	DateTime         *QueryParameterDateLike      `json:"datetime,omitempty"`
	DateTimeSec      *QueryParameterDateLike      `json:"datetimesec,omitempty"`
	DateRange        *QueryParameterDateRangeLike `json:"date_range,omitempty"`
	DateTimeRange    *QueryParameterDateRangeLike `json:"datetime_range,omitempty"`
	DateTimeSecRange *QueryParameterDateRangeLike `json:"datetimesec_range,omitempty"`
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

// QueryParameterDateLike ...
type QueryParameterDateLike struct {
	Value string `json:"value"`
}

// QueryParameterDateRangeLike ...
type QueryParameterDateRangeLike struct {
	Value string             `json:"value,omitempty"`
	Range *api.DateTimeRange `json:"range,omitempty"`
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

const secondsInDay = 24 * 60 * 60
const secondsInWeek = 7 * secondsInDay

func (q *QueryEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Query, error) {
	// Extract from ResourceData.
	common.DataToStructPointer(data, schema, q)

	// Transform to API object.
	var aq api.Query
	aq.ID = data.Id()
	aq.DataSourceID = q.DataSourceID
	aq.Name = q.Name
	aq.Description = q.Description
	aq.Query = q.Query
	aq.Tags = append([]string{}, q.Tags...)

	if s := q.Schedule; s != nil {
		if sp := s.Continuous; sp != nil {
			aq.Schedule = &api.QuerySchedule{
				Interval: sp.IntervalSeconds,
			}
			if sp.UntilDate != "" {
				aq.Schedule.Until = &sp.UntilDate
			}
		}
		if sp := s.Daily; sp != nil {
			aq.Schedule = &api.QuerySchedule{
				Interval: sp.IntervalDays * secondsInDay,
				Time:     &sp.TimeOfDay,
			}
			if sp.UntilDate != "" {
				aq.Schedule.Until = &sp.UntilDate
			}
		}
		if sp := s.Weekly; sp != nil {
			aq.Schedule = &api.QuerySchedule{
				Interval:  sp.IntervalWeeks * secondsInWeek,
				DayOfWeek: &sp.DayOfWeek,
				Time:      &sp.TimeOfDay,
			}
			if sp.UntilDate != "" {
				aq.Schedule.Until = &sp.UntilDate
			}
		}
	}

	if len(q.Parameter) > 0 {
		aq.Options = &api.QueryOptions{}
		for _, p := range q.Parameter {
			ap := api.QueryParameter{
				Name:  p.Name,
				Title: p.Title,
			}

			var iface any

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
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: ap,
						StringValue:    p.DateRange.Value,
						RangeValue:     p.DateRange.Range,
					},
				}
			case p.DateTimeRange != nil:
				iface = api.QueryParameterDateTimeRange{
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: ap,
						StringValue:    p.DateTimeRange.Value,
						RangeValue:     p.DateTimeRange.Range,
					},
				}
			case p.DateTimeSecRange != nil:
				iface = api.QueryParameterDateTimeSecRange{
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: ap,
						StringValue:    p.DateTimeSecRange.Value,
						RangeValue:     p.DateTimeSecRange.Range,
					},
				}
			default:
				log.Fatalf("Don't know what to do for QueryParameter...")
			}

			aq.Options.Parameters = append(aq.Options.Parameters, iface)
		}
	}

	if q.RunAsRole != "" {
		if aq.Options == nil {
			aq.Options = &api.QueryOptions{}
		}
		aq.Options.RunAsRole = q.RunAsRole
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
		// Set `schedule` to non-empty value to ensure it's picked up by `StructToSchema`.
		// If it is not yet set in `schema.ResourceData`, then `StructToSchema` mistakingly
		// interprets the server side value as a default and skips the field.
		// This means, however, that if the schedule is configured out-of-band (e.g. manually),
		// running `terraform apply` again won't remove the setting.
		data.Set("schedule", []any{
			map[string][]any{},
		})

		q.Schedule = &QuerySchedule{}
		switch {
		case s.Interval%secondsInWeek == 0:
			q.Schedule.Weekly = &QueryScheduleWeekly{
				IntervalWeeks: s.Interval / secondsInWeek,
			}
			if s.DayOfWeek != nil {
				q.Schedule.Weekly.DayOfWeek = *s.DayOfWeek
			}
			if s.Time != nil {
				q.Schedule.Weekly.TimeOfDay = *s.Time
			}
			if s.Until != nil {
				q.Schedule.Weekly.UntilDate = *s.Until
			}
		case s.Interval%secondsInDay == 0:
			q.Schedule.Daily = &QueryScheduleDaily{
				IntervalDays: s.Interval / secondsInDay,
			}
			if s.Time != nil {
				q.Schedule.Daily.TimeOfDay = *s.Time
			}
			if s.Until != nil {
				q.Schedule.Daily.UntilDate = *s.Until
			}
		default:
			q.Schedule.Continuous = &QueryScheduleContinuous{
				IntervalSeconds: s.Interval,
			}
			if s.Until != nil {
				q.Schedule.Continuous.UntilDate = *s.Until
			}
		}
	} else {
		// Overwrite `schedule` in case it's not set on the server side.
		// This would have been skipped by `common.StructToData` because of slice emptiness.
		// Ideally, the reflection code also sets empty values, but we'd risk
		// clobbering values we actually want to keep around in existing code.
		data.Set("schedule", nil)
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
				p.Date = &QueryParameterDateLike{
					Value: apv.Value,
				}
			case *api.QueryParameterDateTime:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTime = &QueryParameterDateLike{
					Value: apv.StringValue,
				}
			case *api.QueryParameterDateTimeSec:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeSec = &QueryParameterDateLike{
					Value: apv.Value,
				}
			case *api.QueryParameterDateRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateRange = &QueryParameterDateRangeLike{
					Value: apv.StringValue,
					Range: apv.RangeValue,
				}
			case *api.QueryParameterDateTimeRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeRange = &QueryParameterDateRangeLike{
					Value: apv.StringValue,
					Range: apv.RangeValue,
				}
			case *api.QueryParameterDateTimeSecRange:
				p.Name = apv.Name
				p.Title = apv.Title
				p.DateTimeSecRange = &QueryParameterDateRangeLike{
					Value: apv.StringValue,
					Range: apv.RangeValue,
				}
			default:
				log.Fatalf("Don't know what to do for type: %#v", reflect.TypeOf(apv).String())
			}

			q.Parameter = append(q.Parameter, p)
		}

		q.RunAsRole = aq.Options.RunAsRole
	}

	// Transform to ResourceData.
	return common.StructToData(*q, schema, data)
}

// NewQueryAPI ...
func NewQueryAPI(ctx context.Context, m any) QueryAPI {
	return QueryAPI{m.(*common.DatabricksClient), ctx}
}

// QueryAPI ...
type QueryAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create ...
func (a QueryAPI) Create(q *api.Query) error {
	err := a.client.Post(a.context, "/preview/sql/queries", q, &q)
	if err != nil {
		return err
	}

	// New queries are created with a table visualization by default.
	// We don't manage that visualization here, so immediately remove it.
	for _, rv := range q.Visualizations {
		var v api.Visualization
		err = json.Unmarshal(rv, &v)
		if err != nil {
			return err
		}
		// This is a best effort -- don't fail if it doesn't work.
		err = NewVisualizationAPI(a.context, a.client).Delete(v.ID.String())
		if err != nil {
			log.Printf("[WARN] Unable to delete automatically created visualization for query %s (%s)", q.ID, v.ID)
		}
	}
	q.Visualizations = []json.RawMessage{}
	return nil
}

// Read ...
func (a QueryAPI) Read(queryID string) (*api.Query, error) {
	var q api.Query
	err := a.client.Get(a.context, fmt.Sprintf("/preview/sql/queries/%s", queryID), nil, &q)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

// Update ...
func (a QueryAPI) Update(queryID string, q *api.Query) error {
	return a.client.Post(a.context, fmt.Sprintf("/preview/sql/queries/%s", queryID), q, nil)
}

// Delete ...
func (a QueryAPI) Delete(queryID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/preview/sql/queries/%s", queryID), nil)
}

func ResourceSqlQuery() *schema.Resource {
	s := common.StructToSchema(
		QueryEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			schedule := m["schedule"].Elem.(*schema.Resource)

			// Make different query schedule types mutually exclusive.
			{
				ns := []string{"continuous", "daily", "weekly"}
				for _, n1 := range ns {
					for _, n2 := range ns {
						if n1 == n2 {
							continue
						}
						schedule.Schema[n1].ConflictsWith = append(schedule.Schema[n1].ConflictsWith, fmt.Sprintf("schedule.0.%s", n2))
					}
				}
			}

			// Validate week of day in weekly schedule.
			// Manually verified that this is case sensitive.
			weekly := schedule.Schema["weekly"].Elem.(*schema.Resource)
			weekly.Schema["day_of_week"].ValidateFunc = validation.StringInSlice([]string{
				"Sunday",
				"Monday",
				"Tuesday",
				"Wednesday",
				"Thursday",
				"Friday",
				"Saturday",
			}, false)

			m["run_as_role"].ValidateFunc = validation.StringInSlice([]string{"viewer", "owner"}, false)
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			err = NewQueryAPI(ctx, c).Create(aq)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(aq.ID)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			aq, err := NewQueryAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			var q QueryEntity
			return q.fromAPIObject(aq, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var q QueryEntity
			aq, err := q.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return NewQueryAPI(ctx, c).Update(data.Id(), aq)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			return NewQueryAPI(ctx, c).Delete(data.Id())
		},
		Schema: s,
	}.ToResource()
}
